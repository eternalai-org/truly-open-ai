package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
)

func (s *Service) InfraTwitterAppAuthenInstall(ctx context.Context, address string, installUri string, installCode string, signature string) (string, error) {
	err := func() error {
		if address == "" || installCode == "" || signature == "" {
			return errs.NewError(errs.ErrBadRequest)
		}
		err := helpers.ValidateMessageSignature(
			installCode,
			address,
			signature,
		)
		if err != nil {
			return errs.NewError(err)
		}
		infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
			daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"install_code = ?": {installCode},
			},
			map[string][]interface{}{}, []string{},
		)
		if err != nil {
			return errs.NewError(err)
		}
		if infraTwitterApp == nil {
			infraTwitterApp = &models.InfraTwitterApp{
				Address:     address,
				InstallCode: installCode,
				Signature:   signature,
			}
		} else {
			if !strings.EqualFold(infraTwitterApp.Address, address) {
				return errs.NewError(errs.ErrBadRequest)
			}
		}
		err = s.dao.Save(daos.GetDBMainCtx(ctx), infraTwitterApp)
		if err != nil {
			return errs.NewError(err)
		}
		return nil
	}()
	if err != nil {
		return helpers.BuildUri(
			installUri,
			map[string]string{
				"install_code": installCode,
				"error":        err.Error(),
			},
		), nil
	}
	redirectUri := helpers.BuildUri(
		s.conf.InfraTwitterApp.RedirectUri,
		map[string]string{
			"install_code": installCode,
			"install_uri":  installUri,
		},
	)
	return helpers.BuildUri(
		"https://twitter.com/i/oauth2/authorize",
		map[string]string{
			"client_id":             s.conf.InfraTwitterApp.OauthClientId,
			"state":                 "state",
			"response_type":         "code",
			"code_challenge":        "challenge",
			"code_challenge_method": "plain",
			"scope":                 "offline.access+tweet.read+tweet.write+users.read+follows.write+like.write+like.read+users.read",
			"redirect_uri":          redirectUri,
		},
	), nil
}

func (s *Service) InfraTwitterAppAuthenCallback(ctx context.Context, installUri string, installCode string, code string) (string, error) {
	if installCode == "" || code == "" {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	address, err := func() (string, error) {
		redirectUri := helpers.BuildUri(
			s.conf.InfraTwitterApp.RedirectUri,
			map[string]string{
				"install_code": installCode,
				"install_uri":  installUri,
			},
		)
		respOauth, err := s.twitterAPI.TwitterOauthCallbackForSampleApp(
			s.conf.InfraTwitterApp.OauthClientId, s.conf.InfraTwitterApp.OauthClientSecret, code, redirectUri)
		if err != nil {
			return "", errs.NewError(err)
		}
		if respOauth != nil && respOauth.AccessToken != "" {
			twitterUser, err := s.twitterAPI.GetTwitterMe(respOauth.AccessToken)
			if err != nil {
				return "", errs.NewError(err)
			}
			twitterInfo, err := s.dao.FirstTwitterInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {twitterUser.ID},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return "", errs.NewError(err)
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
				return "", errs.NewError(err)
			}
			infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"install_code = ?": {installCode},
				},
				map[string][]interface{}{}, []string{},
			)
			if err != nil {
				return "", errs.NewError(err)
			}
			if infraTwitterApp == nil {
				return "", errs.NewError(errs.ErrBadRequest)
			}
			return infraTwitterApp.Address, nil
		}
		return "", errs.NewError(errs.ErrBadRequest)
	}()
	if err != nil {
		return helpers.BuildUri(
			installUri,
			map[string]string{
				"install_code": installCode,
				"error":        err.Error(),
			},
		), nil
	}
	params := map[string]string{
		"address": address,
	}
	returnData := base64.StdEncoding.EncodeToString([]byte(helpers.ConvertJsonString(params)))
	return helpers.BuildUri(
		installUri,
		map[string]string{
			"install_code": installCode,
			"return_data":  returnData,
		},
	), nil
}

func (s *Service) InfraTwitterAppExecuteRequest(ctx context.Context, address string, ipfsReq string) (string, error) {
	rawReq, _, err := lighthouse.DownloadDataSimple(ipfsReq)
	if err != nil {
		return "", errs.NewError(err)
	}
	resp, err := func() (interface{}, error) {
		infraTwitterApp, err := s.dao.FirstInfraTwitterApp(
			daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"address = ?": {address},
			},
			map[string][]interface{}{
				"TwitterInfo": {},
			},
			[]string{},
		)
		if err != nil {
			return "", errs.NewError(err)
		}
		var reqMethod struct {
			Method string `json:"method"`
		}
		err = json.Unmarshal(rawReq, &reqMethod)
		if err != nil {
			return nil, errs.NewError(err)
		}
		switch reqMethod.Method {
		case "getUserById":
			{
				var req struct {
					Params struct {
						Id string `json:"id"`
					} `json:"params"`
				}
				err := json.Unmarshal(rawReq, &req)
				if err != nil {
					return nil, errs.NewError(err)
				}
				user, err := s.GetTwitterUserByID(ctx, req.Params.Id)
				if err != nil {
					return nil, errs.NewError(err)
				}
				return user, nil
			}
		case "getUserByUsername":
			{
				var req struct {
					Params struct {
						Username string `json:"username"`
					} `json:"params"`
				}
				err := json.Unmarshal(rawReq, &req)
				if err != nil {
					return nil, errs.NewError(err)
				}
				resp, err := s.GetTwitterUserByUsername(ctx, req.Params.Username)
				if err != nil {
					return nil, errs.NewError(err)
				}
				return resp, nil
			}
		case "seachUserByQuery":
			{
				var req struct {
					Params struct {
						Username string `json:"username"`
					} `json:"params"`
				}
				err := json.Unmarshal(rawReq, &req)
				if err != nil {
					return nil, errs.NewError(err)
				}
				resp, err := s.SeachTwitterUserByQuery(ctx, req.Params.Username)
				if err != nil {
					return nil, errs.NewError(err)
				}
				return resp, nil
			}
		case "getUserTweets":
			{
				var req struct {
					Params struct {
						TwitterID       string `json:"twitter_id"`
						PaginationToken string `json:"pagination_token"`
						MaxResults      int    `json:"max_results"`
					} `json:"params"`
				}
				err := json.Unmarshal(rawReq, &req)
				if err != nil {
					return nil, errs.NewError(err)
				}
				resp, err := s.GetListUserTweets(ctx, req.Params.TwitterID, req.Params.PaginationToken, req.Params.MaxResults)
				if err != nil {
					return nil, errs.NewError(err)
				}
				return resp, nil
			}
		case "tweet":
			{
				if infraTwitterApp == nil {
					return "", errs.NewError(errs.ErrBadRequest)
				}
				var req struct {
					Params struct {
						Content string `json:"content"`
					} `json:"params"`
				}
				err := json.Unmarshal(rawReq, &req)
				if err != nil {
					return nil, errs.NewError(err)
				}
				tweetId, err := helpers.ReplyTweetByToken(infraTwitterApp.TwitterInfo.AccessToken, req.Params.Content, "", "")
				if err != nil {
					return nil, errs.NewError(err)
				}
				return tweetId, nil
			}
		default:
			{
				return nil, errs.NewError(errs.ErrBadRequest)
			}
		}
	}()
	var rawData string
	if err != nil {
		rawData = helpers.ConvertJsonString(&serializers.Resp{Error: errs.NewError(err)})
	} else {
		rawData = helpers.ConvertJsonString(&serializers.Resp{Result: resp})
	}
	ipfsHash, err := s.IpfsUploadDataForName(ctx, "data", []byte(rawData))
	if err != nil {
		return "", errs.NewError(err)
	}
	return ipfsHash, nil
}
