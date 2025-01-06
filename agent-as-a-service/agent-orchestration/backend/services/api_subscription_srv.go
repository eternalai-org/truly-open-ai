package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) GenerateAPIKey() (string, error) {
	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		return "", errs.NewError(err)
	}

	return hex.EncodeToString(key), nil
}

func (s *Service) CreateApiTokenUsage(ctx context.Context, networkID uint64, apiKey string, endPoint string, numToken uint) error {
	obj, err := s.dao.FirstApiSubscriptionKey(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"api_key = ?": {apiKey},
	}, map[string][]interface{}{}, []string{})
	if err != nil {
		return errs.NewError(err)
	}
	if obj == nil {
		return errs.NewError(errs.ErrBadRequest)
	}
	if networkID > 0 && obj.NetworkID != networkID {
		return errs.NewError(errs.ErrBadRequest)
	}
	if obj.QuotaRemaining < uint64(numToken) {
		return errs.NewError(errs.ErrApiKeyRateLimited)
	}
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			obj, err = s.dao.FirstApiSubscriptionKeyByID(tx, obj.ID, map[string][]interface{}{}, true)

			if err != nil {
				return errs.NewError(err)
			}
			if obj.QuotaRemaining < uint64(numToken) {
				return errs.NewError(errs.ErrApiKeyRateLimited)
			}
			obj.QuotaRemaining = obj.QuotaRemaining - uint64(numToken)
			err = s.dao.Save(tx, obj)
			if err != nil {
				return errs.NewError(err)
			}
			usage := &models.ApiSubscriptionUsageLog{
				ApiKey:   obj.ApiKey,
				Endpoint: endPoint,
				NumToken: uint64(numToken),
			}
			err = s.dao.Create(
				tx,
				usage,
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) RefundApiTokenUsage(ctx context.Context, apiKey string, numToken uint) error {
	obj, err := s.dao.FirstApiSubscriptionKey(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"api_key = ?": {apiKey},
	}, map[string][]interface{}{}, []string{})
	if err != nil {
		return errs.NewError(err)
	}
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			obj, err = s.dao.FirstApiSubscriptionKeyByID(tx, obj.ID, map[string][]interface{}{}, true)

			if err != nil {
				return errs.NewError(err)
			}
			obj.QuotaRemaining = obj.QuotaRemaining + uint64(numToken)
			err = s.dao.Save(tx, obj)
			if err != nil {
				return errs.NewError(err)
			}
			usage := &models.ApiSubscriptionUsageLog{
				ApiKey:   obj.ApiKey,
				Endpoint: "refund",
				NumToken: uint64(numToken),
			}
			err = s.dao.Create(
				tx,
				usage,
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateApiSubscriptionKeyForTest(ctx context.Context, address, twitterID string) (*models.ApiSubscriptionKey, error) {
	userSubApi, err := s.dao.FirstApiSubscriptionKey(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"user_address = ?": {strings.ToLower(address)},
		},
		map[string][]interface{}{}, []string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if userSubApi == nil {
		apiKeyStr, err := s.GenerateAPIKey()
		if err != nil {
			return nil, errs.NewError(err)
		}
		lstFreePackages, err := s.dao.FindApiSubscriptionPackage(daos.GetDBMainCtx(ctx), map[string][]interface{}{
			"type = ?": {models.PackageTypeFree},
		}, map[string][]interface{}{}, []string{}, 0, 100)
		if err != nil {
			return nil, errs.NewError(err)
		}
		now := time.Now()
		for _, v := range lstFreePackages {
			expiresAt := now.Add(time.Duration(v.DurationDay) * 24 * time.Hour)
			depositAddress, err := s.CreateETHAddress(ctx)
			if err != nil {
				return nil, errs.NewError(err)
			}
			userSubApi = &models.ApiSubscriptionKey{
				NetworkID:      v.NetworkID,
				UserAddress:    strings.ToLower(address),
				TwitterID:      twitterID,
				ApiKey:         apiKeyStr,
				PackageID:      v.ID,
				QuotaRemaining: v.NumToken,
				StartedAt:      &now,
				ExpiresAt:      &expiresAt,
				DepositAddress: depositAddress,
			}
			err = s.dao.Create(daos.GetDBMainCtx(ctx), userSubApi)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}

	}
	return userSubApi, nil
}

func (s *Service) GetApiUsage(ctx context.Context, apiKey string, page, limit int) ([]*models.ApiSubscriptionUsageLog, uint, error) {
	res, count, err := s.dao.FindApiSubscriptionUsageLog4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {apiKey},
		},
		map[string][]interface{}{}, []string{"id desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return res, count, nil
}

func (s *Service) TwitterOauthCallbackForApiSubscription(ctx context.Context, callbackUrl, address, code, clientID string) error {
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForDeveloper(
		oauthClientId, oauthClientSecret,
		code, callbackUrl, address)
	if err != nil {
		return errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return errs.NewError(err)
		}

		if twitterUser != nil {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{}, false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			if twitterInfo == nil {
				twitterInfo = &models.TwitterInfo{
					TwitterID: twitterUser.ID,
				}
			}
			twitterInfo.TwitterAvatar = twitterUser.ProfileImageURL
			twitterInfo.TwitterName = twitterUser.Name
			twitterInfo.TwitterUsername = twitterUser.UserName
			twitterInfo.AccessToken = respOauth.AccessToken
			twitterInfo.RefreshToken = respOauth.RefreshToken
			twitterInfo.ExpiresIn = respOauth.ExpiresIn
			twitterInfo.Scope = respOauth.Scope
			twitterInfo.TokenType = respOauth.TokenType
			twitterInfo.OauthClientId = oauthClientId
			twitterInfo.OauthClientSecret = oauthClientSecret
			twitterInfo.Description = twitterUser.Description
			twitterInfo.RefreshError = "OK"

			expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
			twitterInfo.ExpiredAt = &expiredAt
			err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
			if err != nil {
				return errs.NewError(err)
			}

			userSubApi, err := s.dao.FirstApiSubscriptionKey(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"user_address = ?": {strings.ToLower(address)},
				},
				map[string][]interface{}{}, []string{},
			)
			if err != nil {
				return errs.NewError(err)
			}

			if userSubApi == nil {

				lstFreePackages, err := s.dao.FindApiSubscriptionPackage(daos.GetDBMainCtx(ctx), map[string][]interface{}{
					"type = ?": {models.PackageTypeFree},
				}, map[string][]interface{}{}, []string{}, 0, 100)
				if err != nil {
					return errs.NewError(err)
				}
				now := time.Now()
				for _, v := range lstFreePackages {
					apiKeyStr, err := s.GenerateAPIKey()
					if err != nil {
						return errs.NewError(err)
					}
					expiresAt := now.Add(time.Duration(v.DurationDay) * 24 * time.Hour)
					depositAddress, err := s.CreateETHAddress(ctx)
					if err != nil {
						return errs.NewError(err)
					}
					userSubApi = &models.ApiSubscriptionKey{
						NetworkID:      v.NetworkID,
						UserAddress:    strings.ToLower(address),
						TwitterInfoID:  twitterInfo.ID,
						TwitterID:      twitterInfo.TwitterID,
						ApiKey:         apiKeyStr,
						PackageID:      v.ID,
						QuotaRemaining: v.NumToken,
						StartedAt:      &now,
						ExpiresAt:      &expiresAt,
						DepositAddress: depositAddress,
					}
					err = s.dao.Create(daos.GetDBMainCtx(ctx), userSubApi)
					if err != nil {
						return errs.NewError(err)
					}
				}

			}
		}
	}

	return nil
}
func (s *Service) GetApiPackages(ctx context.Context) ([]*models.ApiSubscriptionPackage, error) {
	packages, err := s.dao.FindApiSubscriptionPackage(daos.GetDBMainCtx(ctx), map[string][]interface{}{}, map[string][]interface{}{}, []string{"id asc"}, 0, 1000)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return packages, nil
}

func (s *Service) GetApiKeyInfo(ctx context.Context, apiKey string, address string) ([]*models.ApiSubscriptionKey, error) {
	filters := map[string][]interface{}{}
	if apiKey != "" {
		filters["api_key = ?"] = []interface{}{apiKey}
	}
	if address != "" {
		filters["user_address = ?"] = []interface{}{strings.ToLower(address)}
	}
	obj, err := s.dao.FindApiSubscriptionKey(daos.GetDBMainCtx(ctx), filters, map[string][]interface{}{
		"Package": {},
	}, []string{}, 0, 100)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return obj, nil
}

func (s *Service) ProcessDeposit(ctx context.Context, depositNetworkID uint64, eventID string, txHash string, toAddress string, eventValue *big.Int) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			key, _ := s.dao.FirstApiSubscriptionKey(
				tx,
				map[string][]interface{}{
					"deposit_address = ?": {strings.ToLower(toAddress)},
				},
				map[string][]interface{}{},
				[]string{},
			)
			if key != nil {
				m, err := s.dao.FirstApiSubscriptionHistory(
					tx,
					map[string][]interface{}{
						"event_id = ?": {eventID},
					},
					map[string][]interface{}{},
					[]string{},
				)
				if err != nil {
					return errs.NewError(err)
				}
				if m == nil {
					m = &models.ApiSubscriptionHistory{
						NetworkID:      depositNetworkID,
						EventId:        eventID,
						UserAddress:    key.UserAddress,
						ApiKey:         key.ApiKey,
						DepositAddress: key.DepositAddress,
						DepositStatus:  models.DepositStatusDone,
						TxHash:         txHash,
					}
					err = s.dao.Create(
						tx,
						m,
					)
					if err != nil {
						return errs.NewError(err)
					}
					//get packages
					pack, err := s.dao.FirstApiSubscriptionPackage(tx, map[string][]interface{}{
						"network_id = ?": {key.NetworkID},
						"price = ?":      {numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(eventValue, 18))},
					}, map[string][]interface{}{}, []string{})
					if err != nil {
						return errs.NewError(err)
					}
					if pack != nil {
						key, _ = s.dao.FirstApiSubscriptionKeyByID(tx, key.ID, map[string][]interface{}{}, true)
						if key != nil {
							key.PackageID = pack.ID
							key.QuotaRemaining += pack.NumToken
							newExpired := key.ExpiresAt.Add(time.Duration(pack.DurationDay) * 24 * time.Hour)
							key.ExpiresAt = &newExpired
							err = s.dao.Save(tx, key)
							if err != nil {
								return errs.NewError(err)
							}
						}
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}
