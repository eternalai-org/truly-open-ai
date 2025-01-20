package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/rapid"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetListAgentTwitterPost(ctx context.Context, networkID uint64, page, limit int) ([]*models.AgentTwitterPost, uint, error) {
	joinFilters := map[string][]interface{}{
		`join agent_infos on agent_infos.id = agent_twitter_posts.agent_info_id`: {},
	}
	selected := []string{
		`agent_twitter_posts.*`,
	}

	filters := map[string][]interface{}{
		`agent_twitter_posts.reply_post_id is not null and agent_twitter_posts.reply_post_id != '' and agent_twitter_posts.reply_post_at is not null`: {},
		`agent_twitter_posts.post_type in (?)`: {[]models.AgentSnapshotPostActionType{models.AgentSnapshotPostActionTypeReply, models.AgentSnapshotPostActionTypeTweet, models.AgentSnapshotPostActionTypeTweetV2}},
	}
	if networkID > 0 {
		filters["agent_infos.network_id = ?"] = []interface{}{networkID}
	}

	keys, err := s.dao.FindAgentTwitterPostJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"AgentInfo":             {},
			"AgentInfo.TwitterInfo": {},
			"AgentInfo.TokenInfo":   {},
		},
		[]string{"agent_twitter_posts.reply_post_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) GetAgentTwitterPostDetail(ctx context.Context, postID uint) (*models.AgentTwitterPost, error) {
	post, err := s.dao.FirstAgentTwitterPostByID(daos.GetDBMainCtx(ctx),
		postID,
		map[string][]interface{}{
			"AgentInfo": {},
		},
		false,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return post, nil
}

func (s *Service) GetListAgentEaiTopup(ctx context.Context, agentID string, typeStr string, page, limit int) ([]*models.AgentEaiTopup, uint, error) {
	if typeStr == "" {
		typeStr = string(models.AgentEaiTopupTypeDeposit)
	}
	joinFilters := map[string][]interface{}{
		`join agent_infos on agent_eai_topups.agent_info_id = agent_infos.id`: {},
	}
	selected := []string{
		"agent_eai_topups.*",
	}
	ms, err := s.dao.FindAgentEaiTopupJoinSelect(daos.GetDBMainCtx(ctx),
		selected, joinFilters,
		map[string][]interface{}{
			"agent_infos.agent_id = ?":    {agentID},
			"agent_eai_topups.type = ?":   {typeStr},
			"agent_eai_topups.status = ?": {models.AgentEaiTopupStatusDone},
		},
		map[string][]interface{}{
			"AgentInfo": {},
		},
		[]string{"agent_eai_topups.id desc"},
		page,
		limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, 0, nil
}

func (s *Service) AgentRequestTwitterShareCode(ctx context.Context, topupAddress string) (string, string, error) {
	if topupAddress == "" {
		return "", "", errs.NewError(errs.ErrBadRequest)
	}

	agentInfo, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"eth_address = ?": {strings.ToLower(topupAddress)},
			"network_id = ?":  {models.BASE_CHAIN_ID},
		},
		map[string][]interface{}{},
		[]string{},
	)

	if err != nil {
		return "", "", errs.NewError(errs.ErrBadRequest)
	}

	if agentInfo == nil {
		return "", "", errs.NewError(errs.ErrBadRequest)
	}

	authSecretCode := uuid.NewString()
	authPublicCode := helpers.RandomReferralCode(6)
	err = s.SetRedisCachedWithKey(
		fmt.Sprintf("AgentRequestTwitterShareCode_%s", helpers.GenerateMD5(authSecretCode)),
		&authPublicCode,
		20*time.Minute,
	)
	if err != nil {
		return "", "", errs.NewError(err)
	}
	//save db
	err = s.dao.Save(daos.GetDBMainCtx(ctx), &models.AuthCode{
		ETHAddress: strings.ToLower(topupAddress),
		PublicCode: authPublicCode,
		SecretCode: helpers.GenerateMD5(authSecretCode),
		Expired:    time.Now().Add(20 * time.Minute),
	})

	if err != nil {
		return "", "", errs.NewError(err)
	}
	return authSecretCode, authPublicCode, nil
}

func (s *Service) AgentVerifyShareTwitter(ctx context.Context, authSecretCode string, link string) (bool, error) {
	var authPublicCode string
	authCode, err := s.dao.FirstAuthCode(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"secret_code = ? ": {helpers.GenerateMD5(authSecretCode)},
			"expired >= now()": nil,
		}, map[string][]interface{}{},
		false,
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	if authCode != nil {
		authPublicCode = authCode.PublicCode
	}

	if authPublicCode == "" {
		return false, errs.NewError(errs.ErrBadRequest)
	}

	var twitterUser *rapid.TwitterDetail
	if link != "" {
		tweetID := helpers.ExtractTweetID(link)
		if tweetID != "" {
			twitterUser, err = s.rapid.GetTwitterUserFromTweetID(tweetID, authPublicCode)
			if err != nil {
				return false, errs.NewError(errs.ErrBadRequest)
			}
		}
	} else {
		twitterUser, err = s.rapid.IsPostForFaucetAgent(authPublicCode)
		if err != nil {
			return false, errs.NewError(err)
		}
	}

	if twitterUser == nil || twitterUser.TwitterID == "" {
		return false, errs.NewError(errs.ErrTwitterIdNotFound)
	}

	agentInfo, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"eth_address = ?": {authCode.ETHAddress},
		},
		map[string][]interface{}{},
		[]string{},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	if agentInfo == nil {
		return false, errs.NewError(errs.ErrAgentNotFound)
	}

	historyFaucet, err := s.dao.FirstAgentEaiTopup(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"event_id = ?": {strings.ToLower(authCode.ETHAddress)},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return false, errs.NewError(err)
	}

	if historyFaucet == nil {
		faucetAmount := s.conf.GetConfigKeyString(models.GENERTAL_NETWORK_ID, "agent_faucet_amount")
		historyFaucet = &models.AgentEaiTopup{
			NetworkID:   agentInfo.NetworkID,
			AgentInfoID: agentInfo.ID,
			EventId:     strings.ToLower(authCode.ETHAddress),
			Type:        models.AgentEaiTopupTypeFaucet,
			Amount:      numeric.NewBigFloatFromString(faucetAmount),
			Status:      models.AgentEaiTopupStatusDone,
			ToAddress:   authCode.ETHAddress,
			Toolset:     "faucet",
		}

		err = s.dao.Create(daos.GetDBMainCtx(ctx), historyFaucet)
		if err != nil {
			return false, errs.NewError(err)
		}

		err = daos.GetDBMainCtx(ctx).Model(agentInfo).
			UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", historyFaucet.Amount)).
			UpdateColumn("is_faucet", true).
			Error

		if err != nil {
			return false, errs.NewError(err)
		}
		// _ = s.AgentMintNft(ctx, agentInfo.ID)
	}
	return true, nil
}

// /brains
func (s *Service) GetAgentBrainHistory(ctx context.Context, agentID string, postID uint, page, limit int) ([]*models.AgentSnapshotPost, uint, error) {
	joinFilters := map[string][]interface{}{
		`
			join agent_infos on agent_snapshot_posts.agent_info_id = agent_infos.id and agent_infos.agent_id = ?
		`: {agentID},
	}

	selected := []string{
		"agent_snapshot_posts.*",
	}

	filters := map[string][]interface{}{
		`agent_snapshot_posts.status = ?`: {models.AgentSnapshotPostStatusInferResolved},
		`EXISTS (select 1 from agent_snapshot_post_actions where agent_snapshot_post_id = agent_snapshot_posts.id and status="done")`: {},
	}
	if postID > 0 {
		filters["agent_snapshot_posts.id = ?"] = []interface{}{postID}
	}

	posts, err := s.dao.FindAgentSnapshotPostJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters, filters,
		map[string][]interface{}{
			"AgentInfo":             {},
			"AgentInfo.TwitterInfo": {},
		},
		[]string{"agent_snapshot_posts.created_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return posts, 0, nil
}

func (s *Service) GetBrainDetailByTweetIDBK(ctx context.Context, tweetID string) (*models.AgentSnapshotPost, error) {
	actionPostAction, err := s.dao.FirstAgentSnapshotPostAction(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"ref_id = ?": {tweetID},
		},
		map[string][]interface{}{
			"AgentSnapshotPost": {},
		},
		[]string{},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}
	if actionPostAction != nil && actionPostAction.AgentSnapshotPost != nil {
		return actionPostAction.AgentSnapshotPost, nil
	}
	return nil, nil
}

func (s *Service) GetBrainDetailByTweetID(ctx context.Context, tweetID string) (*models.AgentSnapshotPost, error) {
	filters := map[string][]interface{}{
		`agent_snapshot_posts.status = ?`: {models.AgentSnapshotPostStatusInferResolved},
		`EXISTS (select * from agent_snapshot_post_actions where agent_snapshot_post_id = agent_snapshot_posts.id and status="done")`: {},
	}
	if tweetID != "" {
		filters["agent_snapshot_posts.id = ?"] = []interface{}{tweetID}
	}

	posts, err := s.dao.FirstAgentSnapshotPost(daos.GetDBMainCtx(ctx),
		filters,
		map[string][]interface{}{
			"AgentInfo":               {},
			"AgentInfo.TwitterInfo":   {},
			"AgentSnapshotPostAction": {`status = "done"`},
		},
		[]string{},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return posts, nil
}

func (s *Service) GetListAgentTwitterLatestPost(
	ctx context.Context,
	networkID uint64,
	agentInfoIDs []uint,
	page, limit int,
) (map[uint]*models.AgentTwitterPost, error) {
	joinFilters := map[string][]interface{}{
		`join agent_infos on agent_infos.id = agent_twitter_posts.agent_info_id`: {},
	}
	selected := []string{
		`agent_twitter_posts.*`,
	}

	filters := map[string][]interface{}{
		`agent_twitter_posts.reply_post_id is not null and agent_twitter_posts.reply_post_id != ''`: {},
	}
	if networkID > 0 {
		filters["agent_infos.network_id = ?"] = []interface{}{networkID}
	}
	if len(agentInfoIDs) > 0 {
		filters["agent_twitter_posts.agent_info_id IN (?)"] = []interface{}{agentInfoIDs}
	}

	subQuery := `
		SELECT agent_twitter_posts.agent_info_id, MAX(agent_twitter_posts.post_at) AS latest_post_at
		FROM agent_twitter_posts
		JOIN agent_infos ON agent_infos.id = agent_twitter_posts.agent_info_id
		WHERE agent_twitter_posts.reply_post_id IS NOT NULL AND agent_twitter_posts.reply_post_id != ''
		GROUP BY agent_twitter_posts.agent_info_id
	`

	joinFilters[`JOIN (`+subQuery+`) latest_posts ON agent_twitter_posts.agent_info_id = latest_posts.agent_info_id AND agent_twitter_posts.post_at = latest_posts.latest_post_at`] = []interface{}{}

	keys, err := s.dao.FindAgentTwitterPostJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"AgentInfo":             {},
			"AgentInfo.TwitterInfo": {},
			"AgentInfo.TokenInfo":   {},
		},
		[]string{"agent_twitter_posts.reply_post_at desc"}, page, limit,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}
	resp := map[uint]*models.AgentTwitterPost{}
	for _, v := range keys {
		resp[v.AgentInfoID] = v
	}
	return resp, nil
}

func (s *Service) GetAgentSnapshotMissionConfigs(ctx context.Context, networkID uint64, platform string) ([]*models.AgentSnapshotMissionConfigs, error) {
	if platform == "" {
		platform = string(models.PlatformTypeTwitter)
	}
	filters := map[string][]interface{}{
		`network_id = ?`: {networkID},
		`platform = ?`:   {platform},
		`is_testing = ?`: {false},
	}
	posts, err := s.dao.FindAgentSnapshotMissionConfigs(daos.GetDBMainCtx(ctx),
		filters,
		map[string][]interface{}{},
		[]string{}, 0, 100,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return posts, nil
}

func (s *Service) GetAgentSnapshotMissionTokens(ctx context.Context) ([]configs.MissionTokensConfig, error) {
	return configs.GetMissionTokenConfig()
}

func (s *Service) GetMapAgentSnapshotMissionTokens(ctx context.Context) (map[string]configs.MissionTokensConfig, error) {
	resp := map[string]configs.MissionTokensConfig{}

	arrToken, err := configs.GetMissionTokenConfig()
	if err != nil {
		return resp, errs.NewError(err)
	}

	for _, item := range arrToken {
		resp[item.Symbol] = item
	}

	return resp, nil
}
