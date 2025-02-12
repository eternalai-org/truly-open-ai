package services

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

func (s *Service) SampleTwitterAppAuthenInstall(ctx context.Context, installCode string, installUri string) (string, error) {
	if installCode == "" {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	redirectUri := helpers.BuildUri(
		s.conf.SampleTwitterApp.RedirectUri,
		map[string]string{
			"install_code": installCode,
			"install_uri":  installUri,
		},
	)
	return helpers.BuildUri(
		"https://twitter.com/i/oauth2/authorize",
		map[string]string{
			"client_id":             s.conf.SampleTwitterApp.OauthClientId,
			"state":                 "state",
			"response_type":         "code",
			"code_challenge":        "challenge",
			"code_challenge_method": "plain",
			"scope":                 "offline.access+tweet.read+tweet.write+users.read+follows.write+like.write+like.read+users.read",
			"redirect_uri":          redirectUri,
		},
	), nil
}

func (s *Service) SampleTwitterAppAuthenCallback(ctx context.Context, installCode string, installUri string, code string) (string, error) {
	if installCode == "" || code == "" {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	apiKey, err := func() (string, error) {
		redirectUri := helpers.BuildUri(
			s.conf.SampleTwitterApp.RedirectUri,
			map[string]string{
				"install_code": installCode,
				"install_uri":  installUri,
			},
		)
		respOauth, err := s.twitterAPI.TwitterOauthCallbackForSampleApp(
			s.conf.SampleTwitterApp.OauthClientId, s.conf.SampleTwitterApp.OauthClientSecret, code, redirectUri)
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
			sampleTwitterApp, err := s.dao.FirstSampleTwitterApp(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"install_code = ?": {installCode},
				},
				map[string][]interface{}{}, []string{},
			)
			if err != nil {
				return "", errs.NewError(err)
			}
			if sampleTwitterApp == nil {
				sampleTwitterApp = &models.SampleTwitterApp{
					InstallCode:   installCode,
					ApiKey:        helpers.RandomStringWithLength(64),
					TwitterInfoID: twitterInfo.ID,
				}
				err = s.dao.Create(daos.GetDBMainCtx(ctx), sampleTwitterApp)
				if err != nil {
					return "", errs.NewError(err)
				}
			}
			return sampleTwitterApp.ApiKey, nil
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
		"api_key": apiKey,
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

func (s *Service) SampleTwitterAppGetBTCPrice(ctx context.Context) string {
	btcPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), "BTC")
	return numeric.BigFloat2Text(btcPrice)
}

func (s *Service) SampleTwitterAppTweetMessage(ctx context.Context, apiKey string, content string) (string, error) {
	if apiKey == "" || content == "" {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	sampleTwitterApp, err := s.dao.FirstSampleTwitterApp(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"api_key = ?": {apiKey},
		},
		map[string][]interface{}{
			"TwitterInfo": {},
		},
		[]string{},
	)
	if err != nil {
		return "", errs.NewError(err)
	}
	if sampleTwitterApp == nil {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	userMe, err := helpers.GetTwitterUserMe(sampleTwitterApp.TwitterInfo.AccessToken)
	if err != nil {
		return "", errs.NewError(err)
	}
	maxChars := 280
	if userMe.Data.Verified {
		maxChars = 4000
	}
	contentLines := helpers.SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(content, maxChars)
	if len(contentLines) <= 0 {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	refId, err := helpers.PostTweetByToken(sampleTwitterApp.TwitterInfo.AccessToken, contentLines[0], "")
	if err != nil {
		return "", errs.NewError(err)
	}
	for i := 1; i < len(contentLines); i++ {
		refIdTmp, _ := helpers.ReplyTweetByToken(sampleTwitterApp.TwitterInfo.AccessToken, contentLines[i], refId, "")
		if refIdTmp != "" {
			refId = refIdTmp
		}
	}
	return refId, nil
}
