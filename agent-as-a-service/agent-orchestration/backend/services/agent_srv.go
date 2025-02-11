package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	sync "sync"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/openai"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
	openai2 "github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

func (s *Service) UpdateAgentScanEventSuccess(ctx context.Context, agentInfoID uint, lastTimeEvent *time.Time, lastId string) error {
	if lastTimeEvent == nil {
		lastTimeEvent = helpers.TimeNow()
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where(
			"id = ?", agentInfoID,
		).
		Updates(
			map[string]interface{}{
				"scan_latest_time": lastTimeEvent,
				"scan_latest_id":   lastId,
				"scan_error":       "OK",
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateAgentScanEventError(ctx context.Context, agentInfoID uint, errData error) error {
	if strings.Contains(errData.Error(), "not found") {
		return nil
	}
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AgentInfo{}).
		Where(
			"id = ?", agentInfoID,
		).
		Updates(
			map[string]interface{}{
				"scan_error": errData.Error(),
			},
		).Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GenerateTipAddress(ctx context.Context, agentInfoID uint) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent.TipBtcAddress == "" {
		address, err := s.CreateBTCAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_btc_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	if agent.TipEthAddress == "" {
		address, err := s.CreateETHAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_eth_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	if agent.TipSolAddress == "" {
		address, err := s.CreateSOLAddress(ctx)
		if err != nil {
			return errs.NewError(err)
		}
		err = daos.GetDBMainCtx(ctx).Model(agent).UpdateColumn("tip_sol_address", address).Error
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) TwitterOauthCallbackV1(ctx context.Context, callbackUrl, address, code, agentID, clientID string) error {
	if agentID == "" {
		return s.TwitterOauthCallbackForRelink(ctx, callbackUrl, address, code, clientID)
	} else if agentID == "0" {
		return s.TwitterOauthCallbackForApiSubscription(ctx, callbackUrl, address, code, clientID)
	} else if agentID == "1" {
		return s.TwitterOauthCallbackForCreateAgent(ctx, callbackUrl, address, code, clientID)
	}

	agentInfo, err := s.SyncAgentInfoDetailByAgentID(ctx, agentID)
	if err != nil {
		return errs.NewError(err)
	}

	if agentInfo != nil {
		isFirstLinked := false
		// isAdvance := false
		oauthClientId := ""
		oauthClientSecret := ""
		if strings.EqualFold(clientID, s.conf.Twitter.OauthClientId) {
			oauthClientId = s.conf.Twitter.OauthClientId
			oauthClientSecret = s.conf.Twitter.OauthClientSecret
		} else {
			// isAdvance = true
			oauthClientId = agentInfo.OauthClientId
			oauthClientSecret = agentInfo.OauthClientSecret
		}

		respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKey(
			oauthClientId, oauthClientSecret,
			code, callbackUrl, address, agentID)
		if err != nil {
			return errs.NewError(err)
		}

		if respOauth != nil && respOauth.AccessToken != "" {
			twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
			if err != nil {
				return errs.NewError(err)
			}

			user, err := s.GetUser(daos.GetDBMainCtx(ctx), agentInfo.NetworkID, address, true)
			if err != nil {
				return errs.NewError(err)
			}

			user.TwitterID = twitterUser.ID
			user.TwitterAvatar = twitterUser.ProfileImageURL
			user.TwitterName = twitterUser.Name
			user.TwitterUsername = twitterUser.UserName

			err = s.dao.Save(daos.GetDBMainCtx(ctx), user)
			if err != nil {
				return errs.NewError(err)
			}

			//
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
				isFirstLinked = true
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
			//

			updateFields := map[string]interface{}{
				"twitter_info_id":  twitterInfo.ID,
				"twitter_id":       twitterInfo.TwitterID,
				"twitter_username": twitterInfo.TwitterUsername,
				"scan_enabled":     true,
				"reply_enabled":    true,
			}

			err = daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
				updateFields,
			).Error
			if err != nil {
				return errs.NewError(err)
			}

			if isFirstLinked {
				// off default follow
				// go s.FollowListDefaultTwitters(ctx, agentInfo.ID)
				go s.AgentCreateMissionDefault(context.Background(), agentInfo.ID)
			}

		}
	}
	return nil
}

func (s *Service) AgentCreateMissionDefault(ctx context.Context, agentInfoID uint) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	switch agentInfo.AgentType {
	case models.AgentInfoAgentTypeEliza, models.AgentInfoAgentTypeZerepy:
		{
			return nil
		}
	}
	switch agentInfo.NetworkID {
	case models.SHARDAI_CHAIN_ID:
		{
			_ = daos.GetDBMainCtx(ctx).Exec(
				`
				INSERT INTO agent_snapshot_missions (created_at, updated_at, deleted_at, network_id, agent_info_id, user_prompt, interval_sec, reply_enabled, enabled, tool_set)
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Check and follow Twitter accounts that look interesting to you.',
					7200,
					1,
					1,
					'follow'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'follow'
					)
				and agent_infos.id = ?
						`,
				agentInfo.ID,
			).Error
		}
	case models.LOCAL_CHAIN_ID:
		{
		}
	default:
		{
			_ = daos.GetDBMainCtx(ctx).Exec(
				`
				INSERT INTO agent_snapshot_missions (created_at, updated_at, deleted_at, network_id, agent_info_id, user_prompt, interval_sec, reply_enabled, enabled, tool_set)
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Provide a single message to join the following conversation. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don''t include any instructions or extra words, just the raw message ready to post.',
					7200,
					1,
					1,
					'reply_mentions'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'reply_mentions'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Browse Twitter and choose ONE post and reply it. Keep the reply concise (under 128 chars), NO hashtags, links or emojis, and don''t include any instructions or extra words, just the raw reply ready to post. IMPORTANT: Immediately stop after replying one post. DO NOT REPLY YOUR OWN TWEET.',
					7200,
					1,
					1,
					'reply_non_mentions'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'reply_non_mentions'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Check and follow Twitter accounts that look interesting to you.',
					7200,
					1,
					1,
					'follow'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'follow'
					)
				and agent_infos.id = ?
				union
				select now(),
					now(),
					null,
					agent_infos.network_id,
					agent_infos.id,
					'Choose ONE topic that you like or dislike, and tweet about it with your own perspective.',
					7200,
					1,
					1,
					'post_search'
				from agent_infos
				where not exists(
						select 1
						from agent_snapshot_missions
						where agent_snapshot_missions.agent_info_id = agent_infos.id
						and agent_snapshot_missions.tool_set = 'post_search'
					)
				and agent_infos.id = ?;
						`,
				agentInfo.ID,
				agentInfo.ID,
				agentInfo.ID,
				agentInfo.ID,
			).Error
		}
	}
	return nil
}

func (s *Service) TwitterOauthCallbackForInternalData(ctx context.Context, callbackUrl, code string) error {
	respOauth, err := s.twitterAPI.TwitterOauthCallbackForInternalData(
		s.conf.Twitter.OauthClientIdForTwitterData, s.conf.Twitter.OauthClientSecretForTwitterData, code, callbackUrl)
	if err != nil {
		return errs.NewError(err)
	}

	if respOauth != nil && respOauth.AccessToken != "" {
		twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
		if err != nil {
			return errs.NewError(err)
		}

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
		twitterInfo.OauthClientId = s.conf.Twitter.OauthClientIdForTwitterData
		twitterInfo.OauthClientSecret = s.conf.Twitter.OauthClientSecretForTwitterData
		twitterInfo.Description = twitterUser.Description
		twitterInfo.RefreshError = "OK"

		expiredAt := time.Now().Add(time.Second * time.Duration(respOauth.ExpiresIn-(60*20)))
		twitterInfo.ExpiredAt = &expiredAt
		err = s.dao.Save(daos.GetDBMainCtx(ctx), twitterInfo)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) TwitterOauthCallbackForRelink(ctx context.Context, callbackUrl, address, code, clientID string) error {
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForRelink(
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

			agentInfos, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`twitter_id = ?`: {twitterUser.ID},
				},
				map[string][]interface{}{},
				[]string{},
				0, 100,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfos != nil {
				for _, agentInfo := range agentInfos {
					updateFields := map[string]interface{}{
						"twitter_info_id":  twitterInfo.ID,
						"twitter_id":       twitterInfo.TwitterID,
						"twitter_username": twitterInfo.TwitterUsername,
						"scan_enabled":     true,
					}

					err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						updateFields,
					).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}

		}

	}

	return nil
}

func (s *Service) TwitterOauthCallbackForCreateAgent(ctx context.Context, callbackUrl, address, code, clientID string) error {
	oauthClientId := s.conf.Twitter.OauthClientId
	oauthClientSecret := s.conf.Twitter.OauthClientSecret

	respOauth, err := s.twitterAPI.GetTwitterOAuthTokenWithKeyForCreateAgent(
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

			agentInfos, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					`tmp_twitter_id = ?`:                    {twitterUser.ID},
					`(twitter_id is null or twitter_id="")`: {},
				},
				map[string][]interface{}{},
				[]string{},
				0, 100,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfos != nil {
				user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.GENERTAL_NETWORK_ID, strings.ToLower(address), false)
				if err != nil {
					return errs.NewError(err)
				}

				if user != nil {
					user.TwitterID = twitterInfo.TwitterID
					user.TwitterName = twitterInfo.TwitterName
					user.TwitterUsername = twitterInfo.TwitterUsername
					user.TwitterAvatar = twitterInfo.TwitterAvatar
				}

				for _, agentInfo := range agentInfos {
					updateFields := map[string]interface{}{
						"creator":      strings.ToLower(address),
						"scan_enabled": true,
					}

					err := daos.GetDBMainCtx(ctx).Model(agentInfo).Updates(
						updateFields,
					).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}

		}

	}

	return nil
}

func (s *Service) CreateUpdateUserTwitter(tx *gorm.DB, userTwitterID string) (*models.TwitterUser, error) {
	tweetUser, err := s.dao.FirstTwitterUser(tx,
		map[string][]interface{}{
			"twitter_id = ?": {userTwitterID},
		}, map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	info, err := s.twitterAPI.GetTwitterUserInfoID(userTwitterID)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if tweetUser == nil {
		tweetUser = &models.TwitterUser{
			TwitterID:       info.TwitterID,
			TwitterUsername: info.TwitterUsername,
			Name:            info.Name,
			ProfileUrl:      info.ProfileUrl,
			FollowersCount:  info.FollowersCount,
			FollowingsCount: info.FollowingsCount,
			IsBlueVerified:  info.IsBlueVerified,
			JoinedAt:        info.CreatedAt,
		}
	} else {
		tweetUser.TwitterUsername = info.TwitterUsername
		tweetUser.Name = info.Name
		tweetUser.ProfileUrl = info.ProfileUrl
	}

	err = s.dao.Save(tx, tweetUser)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweetUser, nil
}

func (s *Service) CreateUpdateUserTwitterByUserName(tx *gorm.DB, username string) (*models.TwitterUser, error) {
	tweetUser, err := s.dao.FirstTwitterUser(tx,
		map[string][]interface{}{
			"twitter_username = ?": {username},
		}, map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if tweetUser == nil {
		info, err := s.twitterAPI.GetTwitterByUserName(username)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if info != nil {
			tweetUser = &models.TwitterUser{
				TwitterID:       info.ID,
				TwitterUsername: info.UserName,
				Name:            info.Name,
				ProfileUrl:      info.ProfileImageURL,
			}

			err = s.dao.Create(tx, tweetUser)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}
	}
	return tweetUser, nil
}

func (s *Service) JobUpdateTwitterAccessToken(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobUpdateTwitterAccessToken",
		func() error {
			var retErr error
			twitterInfos, err := s.dao.FindTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"expired_at <= now()":                             {},
					"expired_at >= adddate(now(), interval -24 hour)": {},
				},
				map[string][]interface{}{},
				[]string{
					"updated_at asc",
				}, 0, 20,
			)
			if err != nil {
				return errs.NewError(err)
			}
			for _, twitterInfo := range twitterInfos {
				err := s.UpdateTwitterAccessToken(ctx, twitterInfo.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterInfo.ID))
				}
			}
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) UpdateTwitterAccessToken(ctx context.Context, twitterInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("UpdateTwitterAccessToken_%d", twitterInfoID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterInfo, err := s.dao.FirstTwitterInfoByID(tx, twitterInfoID, map[string][]interface{}{}, true)
					if err != nil {
						return errs.NewError(err)
					}
					var authInfo *twitter.TwitterTokenResponse

					if twitterInfo.OauthClientId != "" && twitterInfo.OauthClientSecret != "" {
						authInfo, err = s.twitterAPI.GetTwitterAccessTokenWithKey(twitterInfo.OauthClientId,
							twitterInfo.OauthClientSecret, twitterInfo.RefreshToken)
						// if err != nil {
						// 	return errs.NewError(err)
						// }
					} else {
						authInfo, err = s.twitterAPI.GetTwitterAccessToken(twitterInfo.RefreshToken)
						// if err != nil {
						// 	return errs.NewError(err)
						// }
					}

					if authInfo != nil && err == nil {
						twitterInfo.AccessToken = authInfo.AccessToken
						twitterInfo.RefreshToken = authInfo.RefreshToken
						twitterInfo.ExpiresIn = authInfo.ExpiresIn
						twitterInfo.Scope = authInfo.Scope
						twitterInfo.TokenType = authInfo.TokenType
						expiredAt := time.Now().Add(time.Second * time.Duration(authInfo.ExpiresIn-(60*20)))
						twitterInfo.ExpiredAt = &expiredAt
						twitterInfo.RefreshError = "OK"
						err = s.dao.Save(tx, twitterInfo)
						if err != nil {
							return errs.NewError(err)
						}
					} else {
						twitterInfo.RefreshError = err.Error()
						err = s.dao.Save(tx, twitterInfo)
						if err != nil {
							return errs.NewError(err)
						}
					}
					return nil
				},
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
	_ = s.UpdateAgentTwitterInfo(ctx, twitterInfoID)
	return nil
}

func (s *Service) UpdateAgentTwitterInfo(ctx context.Context, twitterInfoID uint) error {
	agents, err := s.dao.FindAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_info_id = ?": {twitterInfoID},
		},
		map[string][]interface{}{
			"TwitterInfo": {},
		},
		[]string{},
		0,
		999999,
	)
	if err != nil {
		return errs.NewError(err)
	}
	for _, agent := range agents {
		func(agent *models.AgentInfo) error {
			if agent.TwitterInfo != nil {
				userMe, err := helpers.GetTwitterUserMe(agent.TwitterInfo.AccessToken)
				if err != nil {
					return errs.NewError(err)
				}
				if userMe != nil && userMe.Data.UserName != "" && userMe.Data.Name != "" {
					err = daos.GetDBMainCtx(ctx).Model(agent).
						UpdateColumn("twitter_username", userMe.Data.UserName).
						UpdateColumn("agent_name", userMe.Data.Name).
						UpdateColumn("twitter_verified", userMe.Data.Verified).Error
					if err != nil {
						return errs.NewError(err)
					}
					err = daos.GetDBMainCtx(ctx).Model(agent.TwitterInfo).
						UpdateColumn("twitter_username", userMe.Data.UserName).
						UpdateColumn("twitter_name", userMe.Data.Name).
						UpdateColumn("twitter_avatar", userMe.Data.ProfileImageURL).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		}(agent)
	}
	return nil
}

func (s *Service) GetTwitterVerified(tx *gorm.DB, agentInfoID uint) (bool, error) {
	m, err := s.dao.FirstAgentInfoByID(
		tx,
		agentInfoID,
		map[string][]interface{}{
			"TwitterInfo": {},
		},
		false,
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	if m == nil {
		return false, errs.NewError(errs.ErrBadRequest)
	}
	return m.TwitterVerified, nil
}

func (s *Service) GetTwitterPostMaxChars(tx *gorm.DB, agentInfoID uint) (uint, error) {
	verified, err := s.GetTwitterVerified(tx, agentInfoID)
	if err != nil {
		return 0, errs.NewError(err)
	}
	if verified {
		return 4000, nil
	}
	return 280, nil
}

func (s *Service) GetAgentTokenInfoFromContractAddress(ctx context.Context, tokenAddress string) (string, string, error) {
	tokenAddress = strings.ToLower(tokenAddress)
	if tokenAddress != "" && tokenAddress != "no" && tokenAddress != "yes" && tokenAddress != "pending" {
		tokenMetaData, err := s.blockchainUtils.SolanaTokenMetaData(tokenAddress)
		if err != nil {
			return "", "", nil
		}
		return tokenMetaData.Name, tokenMetaData.Symbol, nil
	}
	return "", "", nil
}

func (s *Service) UpdateAgentFarcasterInfo(ctx context.Context, agentID string, fID string, fUsername string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo != nil {
				agentInfo, _ = s.dao.FirstAgentInfoByID(tx, agentInfo.ID, map[string][]interface{}{}, true)
				agentInfo.FarcasterID = fID
				agentInfo.FarcasterUsername = fUsername
				err = s.dao.Save(tx, agentInfo)
				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) PreviewAgentSystemPromp(ctx context.Context, personality, question string) (string, error) {
	aiStr, err := s.openais["Agent"].TestAgentPersinality(personality, question, s.conf.AgentOffchainChatUrl)
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) RetrieveKnowledge(agentModel string, messages []openai2.ChatCompletionMessage,
	knowledgeBases []*models.KnowledgeBase, topK *int, threshold *float64) (string, error) {
	if len(knowledgeBases) == 0 {
		return "", errs.NewError(errors.New("knowledge bases is empty"))
	}
	if agentModel == "" {
		agentModel = "Llama3.3"
	}
	systemPrompt := openai.GetSystemPromptFromLLMMessage(messages)
	isKbAgent := false
	if systemPrompt == "" {
		isKbAgent = true
		systemPrompt = "You are a helpful assistant."
	}
	_ = isKbAgent

	userPromptInput := openai.LastUserPrompt(messages)
	retrieveQuery := userPromptInput
	retrieveQueryFromLLM, _ := s.GenerateKnowledgeQuery(systemPrompt, userPromptInput)
	if retrieveQueryFromLLM != nil {
		retrieveQuery = *retrieveQueryFromLLM
	}

	topKQuery := 5
	if topK != nil {
		topKQuery = *topK
	}
	th := 0.2
	if threshold != nil {
		th = *threshold
	}

	request := serializers.RetrieveKnowledgeBaseRequest{
		Query: retrieveQuery,
		TopK:  topKQuery,
		Kb: []string{
			knowledgeBases[0].KbId,
		},
		Threshold: th,
	}

	// retry
	var (
		body string
		err  error
	)
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		body, err = helpers.CurlURLString(
			s.conf.KnowledgeBaseConfig.QueryServiceUrl,
			"POST",
			map[string]string{},
			&request,
		)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	response := &serializers.RetrieveKnowledgeBaseResponse{}
	err = json.Unmarshal([]byte(body), response)
	if err != nil {
		return "", errs.NewError(err)
	}

	searchResult := []string{}
	for _, item := range response.Result {
		searchResult = append(searchResult, item.Content)
	}

	answerPrmptPrefix := ""
	for _, item := range searchResult {
		answerPrmptPrefix += fmt.Sprintf("- %v\n", item)
	}

	answerPrmptPrefix += ". \n\nUsing the above information to address the user's input: " + userPromptInput
	payloadAgentChat := []openai2.ChatCompletionMessage{
		{
			Role:    openai2.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	for i, item := range messages {
		if item.Role == openai2.ChatMessageRoleSystem {
			continue
		}
		if i == len(messages)-1 {
			continue
		}

		payloadAgentChat = append(payloadAgentChat, openai2.ChatCompletionMessage{
			Role:    openai2.ChatMessageRoleUser,
			Content: item.Content,
		})
	}

	// add answer prompt
	payloadAgentChat = append(payloadAgentChat, openai2.ChatCompletionMessage{
		Role:    openai2.ChatMessageRoleUser,
		Content: answerPrmptPrefix,
	})

	messageCallLLM, _ := json.Marshal(&payloadAgentChat)
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
	}

	stringResp, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), agentModel, url)
	if err != nil {
		return "", errs.NewError(err)
	}
	return stringResp, nil
}

func (s *Service) GenerateKnowledgeQuery(systemPrompt, textUserInput string) (*string, error) {
	baseModel := "Llama3.3"
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
	}

	generateQueryPrefix := `Generate a concise and effective search query to retrieve relevant information from the database. Ensure the query is clear, simple, and optimized for accurate results based on the input question:
%v
Respond in stringified JSON format with the following structure:
{
  "query": "<generated_query>"
}`
	userPrompt := fmt.Sprintf(generateQueryPrefix, textUserInput)
	messages := []openai2.ChatCompletionMessage{
		{
			Role:    openai2.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai2.ChatMessageRoleUser,
			Content: userPrompt,
		},
	}

	maxRetry := 10
	messageCallLLM, _ := json.Marshal(&messages)

	var queryStringResp *string
	for i := 1; i <= maxRetry; i++ {
		if i > 1 {
			time.Sleep(time.Second)
		}

		stringResp, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), baseModel, url)
		if err != nil || stringResp == "" {
			continue
		}
		// find { and } in stringResp
		startIndex := strings.Index(stringResp, "{")
		endIndex := strings.LastIndex(stringResp, "}")
		if startIndex == -1 || endIndex == -1 {
			continue
		}
		queryStringRespJson := stringResp[startIndex : endIndex+1]
		queryStringRespMap := map[string]string{}
		err = json.Unmarshal([]byte(queryStringRespJson), &queryStringRespMap)
		if err != nil {
			continue
		}
		if _, ok := queryStringRespMap["query"]; !ok {
			continue
		}

		val := queryStringRespMap["query"]
		queryStringResp = &val
		break
	}

	return queryStringResp, nil
}

func (s *Service) PreviewAgentSystemPrompV1(ctx context.Context,
	messages string, agentId *uint, kbIdFromKnowledegeBase *string, modelNameFromRequest *string) (string, error) {
	var agentInfo *models.AgentInfo
	baseModel := "NousResearch/Hermes-3-Llama-3.1-70B-FP8"
	if agentId != nil {
		agentInfo, _ = s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), *agentId, map[string][]interface{}{}, false)
	}
	llmMessage := []openai2.ChatCompletionMessage{}
	err := json.Unmarshal([]byte(messages), &llmMessage)
	if err != nil {
		return "", errs.NewError(errors.New("invalid message request"))
	}

	systemContent := openai.GetSystemPromptFromLLMMessage(llmMessage)
	url := s.conf.AgentOffchainChatUrl
	if s.conf.KnowledgeBaseConfig.DirectServiceUrl != "" {
		url = s.conf.KnowledgeBaseConfig.DirectServiceUrl
		if agentInfo != nil && agentInfo.AgentBaseModel != "" {
			baseModel = agentInfo.AgentBaseModel
		}
	}
	if modelNameFromRequest != nil && *modelNameFromRequest != "" {
		baseModel = *modelNameFromRequest
	}

	{
		if s.conf.KnowledgeBaseConfig.EnableSimulation && agentInfo != nil {
			// get last knowledge base of agent
			var knowledgeBaseUse *models.KnowledgeBase
			agentInfoKnowledgeBase, _ := s.dao.FirstAgentInfoKnowledgeBaseByAgentInfoID(
				daos.GetDBMainCtx(ctx),
				agentInfo.ID,
				map[string][]interface{}{
					"KnowledgeBase": {}, // must preload
				},
				[]string{"id desc"},
			)
			if agentInfoKnowledgeBase != nil && agentInfoKnowledgeBase.KnowledgeBase != nil {
				knowledgeBaseUse = agentInfoKnowledgeBase.KnowledgeBase
			}

			if kbIdFromKnowledegeBase != nil {
				knowledgeBaseFromQuery, _ := s.dao.FirstKnowledgeBase(daos.GetDBMainCtx(ctx), map[string][]interface{}{
					"kb_id = ?": {*kbIdFromKnowledegeBase},
				}, map[string][]interface{}{}, []string{"id desc"}, false)
				if knowledgeBaseFromQuery != nil {
					knowledgeBaseUse = knowledgeBaseFromQuery
				}
			}
			if knowledgeBaseUse != nil {
				retrieveKnowledgeBaseResponse, err := s.RetrieveKnowledge(baseModel, llmMessage, []*models.KnowledgeBase{
					knowledgeBaseUse,
				}, nil, nil)

				if err != nil {
					return "", err
				}

				return retrieveKnowledgeBaseResponse, nil
			}
		}
	}

	llmMessage = openai.UpdateSystemPromptInLLMRequest(llmMessage, systemContent)
	messageCallLLM, _ := json.Marshal(&llmMessage)
	aiStr, err := s.openais["Agent"].CallDirectlyEternalLLM(string(messageCallLLM), baseModel, url)
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) AgentChatSupport(ctx context.Context, msg string) (string, error) {
	aiStr, err := s.openais["Lama"].ChatMessage(msg)
	if err != nil {
		return "", errs.NewError(err)
	}
	return aiStr, nil
}

func (s *Service) GetExtractDataFromPost(content string) (string, error) {
	extractData := ""
	link, isTwitterPost := helpers.ExtractLinks(content)
	if link != "" {
		if isTwitterPost {
			twiterPostIDArry := strings.Split(link, "/")
			if len(twiterPostIDArry) > 0 {
				twitterPostID := twiterPostIDArry[len(twiterPostIDArry)-1]
				fmt.Println(twitterPostID)
				tweetDetail, err := s.rapid.GetTweetDetailByID(twitterPostID)
				if err != nil {
					return extractData, nil
				}
				if tweetDetail != nil {
					extractData = tweetDetail.FullText
				}
			}
		} else {
			webContent := helpers.ContentHtmlByUrl(link)
			if webContent == "" {
				webContent = helpers.RodContentHtmlByUrl(link)
			}
			webContent, err := s.blockchainUtils.CleanHtml(webContent)
			if err != nil {
				return extractData, nil
			}

			if webContent != "" {
				summary, err := s.openais["Lama"].SummaryWebContent(webContent)
				if err != nil {
					return extractData, nil
				}
				extractData = summary
			}
		}
	}
	return extractData, nil
}

func (s *Service) CheckAgentIsReadyToRunTwinTraining(agentInfo *models.AgentInfo) (float64, bool, error) {
	needBalance := 0.0
	switch agentInfo.NetworkID {
	case models.SHARDAI_CHAIN_ID:
		needBalance = 5999.9
	case models.ETHEREUM_CHAIN_ID:
		needBalance = 299.9
	case models.SOLANA_CHAIN_ID:
		needBalance = 174.9
	default:
		needBalance = 0.99
	}

	if agentInfo.TwinTwitterUsernames != "" {
		arr := strings.Split(agentInfo.TwinTwitterUsernames, ",")
		// 300 EAI for each twitter username
		needBalance += float64(len(arr)) * 300
	}

	agentEaiBalance, _ := agentInfo.EaiBalance.Float64()
	if agentEaiBalance >= needBalance {
		return needBalance, true, nil
	}

	return needBalance, false, nil
}

func (s *Service) JobAgentTwinTrain(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobAgentTwinTrain",
		func() error {
			// Count pending twin training
			twinTrainingAgents, err := s.dao.FindAgentInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twin_twitter_usernames != '' and twin_status = ?": {models.TwinStatusRunning},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				999999,
			)
			if err == nil && len(twinTrainingAgents) > 5 {
				logger.Info("twin_training_jobs", "twin training is running maximum ===> skip", zap.Any("len_training_agents", len(twinTrainingAgents)))
				return nil
			}

			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					//"agent_id != ''":                                   {},
					//"agent_contract_id = ?":                            {""},
					//"agent_nft_minted = ?":                             {false},
					`twin_twitter_usernames != '' and twin_status = ?`: {models.TwinStatusPending},
					"scan_enabled = ?": {true},
				},
				map[string][]interface{}{},
				[]string{
					"id asc",
				},
				0,
				999999,
			)

			if err != nil {
				return errs.NewError(err)
			}

			var retErr error
			wg := sync.WaitGroup{}
			for _, agent := range agents {
				wg.Add(1)
				go func(_agent *models.AgentInfo) {
					defer wg.Done()
					err = s.AgentTwinTrain(ctx, _agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
				}(agent)
			}
			wg.Wait()
			return retErr
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateUpdateAgentSnapshotMission(ctx context.Context, agentID string, authHeader string, req []*serializers.AgentSnapshotMissionInfo) (*models.AgentInfo, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				defataulUserPromp := ""
				listID := []uint{}

				for _, item := range req {
					if item.ID > 0 {
						listID = append(listID, item.ID)
					}
				}

				listTestToolSet := strings.Split(s.conf.ListTestToolSet, ",")
				if len(listID) > 0 {
					err = tx.Where("agent_info_id = ? and id not in (?) and (mission_store_id = 0 or mission_store_id is NULL) and tool_set not in (?)", agentInfo.ID, listID, listTestToolSet).Delete(&models.AgentSnapshotMission{}).Error
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					err = tx.Where("agent_info_id = ? and (mission_store_id = 0 or mission_store_id is NULL) and tool_set not in (?)", agentInfo.ID, listTestToolSet).Delete(&models.AgentSnapshotMission{}).Error
					if err != nil {
						return errs.NewError(err)
					}
				}

				for _, item := range req {
					if defataulUserPromp == "" {
						defataulUserPromp = item.UserPrompt
					}

					mission := &models.AgentSnapshotMission{}
					if item.ID > 0 {
						mission, err = s.dao.FirstAgentSnapshotMissionByID(tx, item.ID,
							map[string][]interface{}{}, false,
						)
						if err != nil {
							return errs.NewError(err)
						}
					}
					mission.NetworkID = agentInfo.NetworkID
					mission.AgentInfoID = agentInfo.ID
					mission.UserPrompt = item.UserPrompt
					mission.IntervalSec = item.Interval
					mission.ToolSet = item.ToolSet
					mission.Enabled = true
					mission.ReplyEnabled = true
					mission.AgentType = item.AgentType
					mission.UserTwitterIds = item.UserTwitterIDs
					mission.Tokens = item.Tokens
					mission.AgentBaseModel = item.AgentBaseModel
					mission.Topics = item.Topics
					mission.IsBingSearch = item.IsBingSearch
					mission.IsTwitterSearch = item.IsTwitterSearch
					mission.RewardAmount = item.RewardAmount
					mission.RewardUser = item.RewardUser
					mission.MinTokenHolding = item.MinTokenHolding
					mission.LookupInterval = item.LookupInterval
					//farcaster
					if mission.ToolSet == models.ToolsetTypePostFarcaster {
						toolList := fmt.Sprintf(s.conf.ToolLists.FarcasterPost, agentInfo.FarcasterID, authHeader, agentInfo.AgentID)
						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeReplyMentionsFarcaster {
						toolList := fmt.Sprintf(s.conf.ToolLists.FarcasterReply, agentInfo.FarcasterID, authHeader, agentInfo.AgentID)
						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeTradeNews {
						toolList := fmt.Sprintf(s.conf.ToolLists.TradeNews, s.conf.InternalApiKey, agentInfo.ID)
						mission.ToolList = toolList
						mission.UserPrompt = "Analyze the coin price fluctuations in the past 24 hours, suggest which coin to buy or sell and post it on twitter"
					} else if mission.ToolSet == models.ToolsetTypeTradeAnalytics || mission.ToolSet == models.ToolsetTypeTradeAnalyticsOnTwitter {
						toolList := s.conf.ToolLists.TradeAnalytic
						if item.Tokens == "" {
							return errs.NewError(errs.ErrTokenNotFound)
						}
						mission.UserPrompt = fmt.Sprintf(`Conduct a technical analysis of $%s price data. Based on your findings, provide a recommended buy price and sell price to maximize potential returns.`, item.Tokens)
						toolList = strings.ReplaceAll(toolList, "{api_key}", s.conf.InternalApiKey)
						toolList = strings.ReplaceAll(toolList, "{token_symbol}", item.Tokens)

						mission.ToolList = toolList
					} else if mission.ToolSet == models.ToolsetTypeLuckyMoneys {
						if item.RewardAmount.Cmp(big.NewFloat(0)) <= 0 || item.RewardUser <= 0 {
							return errs.NewError(errs.ErrBadRequest)
						}
					} else if item.AgentStoreMissionID > 0 {
						agentStoreMission, err := s.dao.FirstAgentStoreMissionByID(tx, item.AgentStoreMissionID, map[string][]interface{}{}, false)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreMission == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
							tx,
							map[string][]interface{}{
								"agent_store_id = ?": {agentStoreMission.AgentStoreID},
								"agent_info_id = ?":  {mission.AgentInfoID},
							},
							map[string][]interface{}{},
							[]string{"id desc"},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agentStoreInstall == nil {
							return errs.NewError(errs.ErrBadRequest)
						}
						mission.AgentStoreMissionID = agentStoreMission.ID
						mission.AgentStoreID = agentStoreMission.AgentStoreID
						mission.ReactMaxSteps = 5
						mission.ToolSet = "mission_store"
					} else if item.ToolList != "" {
						mission.ToolList = item.ToolList
					}
					if mission.ToolList != "" {
						mission.ReactMaxSteps = 5
					}
					//
					err = s.dao.Save(tx, mission)
					if err != nil {
						return errs.NewError(err)
					}
				}
				if defataulUserPromp != "" {
					updateAgentFields := map[string]interface{}{
						"user_prompt": defataulUserPromp,
					}

					err = tx.Model(agentInfo).Updates(updateAgentFields).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}

			return nil
		},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}
	return s.GetAgentInfoDetailByAgentID(ctx, agentID)
}

func (s *Service) DeleteAgentSnapshotMission(ctx context.Context, missionID uint, userAddress string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			missionInfo, err := s.dao.FirstAgentSnapshotMissionByID(tx,
				missionID,
				map[string][]interface{}{
					"AgentInfo": {},
				}, false,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if missionInfo != nil && missionInfo.AgentInfo != nil && strings.EqualFold(missionInfo.AgentInfo.Creator, userAddress) {
				err = tx.Delete(missionInfo).Error
				if err != nil {
					return errs.NewError(err)
				}
			}

			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) FollowListDefaultTwitters(ctx context.Context, agentID uint) error {
	listFollow, err := s.dao.GetListTwitterDefaultFollow(daos.GetDBMainCtx(ctx))
	if err != nil {
		return errs.NewError(err)
	}
	for _, v := range listFollow {
		agent, err := s.dao.FirstAgentInfoByID(
			daos.GetDBMainCtx(ctx),
			agentID,
			map[string][]interface{}{
				"TwitterInfo": {},
			},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if agent != nil {
			helpers.TwitterFollowUserCreate(agent.TwitterInfo.AccessToken, agent.TwitterID, v)
			time.Sleep(20 * time.Second)
		}
	}
	return nil
}

func (s *Service) UpdateTwinStatus(ctx context.Context, req *serializers.UpdateTwinStatusRequest) (*models.AgentInfo, error) {
	agentIDInt, err := strconv.Atoi(req.AgentID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agentInfoEntity, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"id = ?": {agentIDInt},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	agentInfoEntity.TwinStatus = models.TwinStatus(req.TwinStatus)
	agentInfoEntity.KnowledgeBaseID = req.KnowledgeBaseID
	agentInfoEntity.SystemPrompt = req.SystemPrompt
	agentInfoEntity.TwinTrainingProgress = req.TwinTrainingProgress
	if agentInfoEntity.TwinTrainingProgress > 100 {
		agentInfoEntity.TwinTrainingProgress = 100
	}
	if agentInfoEntity.TwinTrainingProgress < 0 {
		agentInfoEntity.TwinTrainingProgress = 0
	}

	if req.TwinStatus == string(models.TwinStatusDoneError) || req.TwinStatus == string(models.TwinStatusDoneSuccess) {
		endAt := time.Now().UTC()
		agentInfoEntity.TwinTrainingMessage = req.TwinTrainingMessage
		agentInfoEntity.TwinEndTrainingAt = &endAt
	}

	err = s.dao.Save(daos.GetDBMainCtx(ctx), agentInfoEntity)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if req.TwinStatus == string(models.TwinStatusDoneError) {
		eventId := fmt.Sprintf("twin_train_refund_%d", agentInfoEntity.ID)
		checkRefunded, err := s.dao.FirstAgentEaiTopup(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"event_id = ?": {eventId},
			},
			map[string][]interface{}{},
			[]string{},
		)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewError(err)
		}
		if checkRefunded != nil {
			return agentInfoEntity, nil
		}

		// increase balance and add refund history
		arr := strings.Split(agentInfoEntity.TwinTwitterUsernames, ",")
		twinFee := numeric.NewFloatFromString(fmt.Sprintf("%v", float64(len(arr))*300))
		_ = daos.WithTransaction(daos.GetDBMainCtx(ctx), func(dbTx *gorm.DB) error {
			err = daos.GetDBMainCtx(ctx).
				Model(agentInfoEntity).
				Updates(
					map[string]interface{}{
						"eai_balance": gorm.Expr("eai_balance + ?", numeric.NewBigFloatFromFloat(twinFee)),
					},
				).
				Error
			if err != nil {
				return errs.NewError(err)
			}

			return s.dao.Create(
				daos.GetDBMainCtx(ctx),
				&models.AgentEaiTopup{
					NetworkID:      agentInfoEntity.NetworkID,
					EventId:        fmt.Sprintf("twin_train_refund_%d", agentInfoEntity.ID),
					AgentInfoID:    agentInfoEntity.ID,
					Type:           models.AgentEaiTopupTypeRefundTrainFail,
					Amount:         numeric.NewBigFloatFromFloat(twinFee),
					Status:         models.AgentEaiTopupStatusDone,
					DepositAddress: agentInfoEntity.ETHAddress,
					ToAddress:      agentInfoEntity.ETHAddress,
					Toolset:        "twin_train_refund",
				},
			)
		})
	}

	return agentInfoEntity, nil
}

func (s *Service) UnlinkAgentTwitterInfo(ctx context.Context, agentID string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				updateFields := map[string]interface{}{
					"twitter_info_id":  0,
					"twitter_id":       "",
					"twitter_username": "",
				}

				err := tx.Model(agentInfo).Updates(
					updateFields,
				).Error
				if err != nil {
					return errs.NewError(err)
				}

			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) AgentChats(ctx context.Context, agentID string, messages serializers.AgentChatMessageReq) (*openai.ChatResponse, error) {
	var aiStr *openai.ChatResponse
	agentInfo, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"agent_id = ?": {agentID},
		},
		map[string][]interface{}{},
		[]string{},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	if agentInfo != nil {
		aiStr, err = s.openais["Agent"].AgentChats(agentInfo.GetSystemPrompt(), s.conf.AgentOffchainChatUrl, messages)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return aiStr, nil
}

func (s *Service) PauseAgent(ctx context.Context, agentID string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if agentInfo != nil {
				isPause := !agentInfo.ReplyEnabled

				updateFields := map[string]interface{}{
					"reply_enabled": isPause,
				}

				err := tx.Model(agentInfo).Updates(
					updateFields,
				).Error
				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) UpdateAgentExternalInfo(ctx context.Context, agentID string, req *serializers.AgentExternalInfoReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfo, err := s.dao.FirstAgentInfo(tx,
				map[string][]interface{}{
					"agent_id = ?": {agentID},
				},
				map[string][]interface{}{},
				[]string{},
			)

			if err != nil {
				return errs.NewError(err)
			}
			if agentInfo != nil {
				externalInfo, err := s.dao.FirstAgentExternalInfo(tx,
					map[string][]interface{}{
						"agent_info_id = ?": {agentInfo.ID},
						"network_id = ?":    {agentInfo.NetworkID},
						"type = ?":          {req.Type},
					},
					map[string][]interface{}{},
					[]string{})
				if err != nil {
					return errs.NewError(err)
				}
				if externalInfo != nil {
					externalInfo.ExternalID = req.ExternalID
					externalInfo.ExternalUsername = req.ExternalUsername
					externalInfo.ExternalName = req.ExternalName
					err = s.dao.Save(tx, externalInfo)
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					err = s.dao.Create(tx, &models.AgentExternalInfo{
						NetworkID:        agentInfo.NetworkID,
						Type:             models.ExternalAgentType(req.Type),
						AgentInfoID:      agentInfo.ID,
						ExternalID:       req.ExternalID,
						ExternalUsername: req.ExternalUsername,
						ExternalName:     req.ExternalName,
					})
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}
