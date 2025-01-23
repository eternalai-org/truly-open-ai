package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) CreateAgentInternalAction(ctx context.Context, req *serializers.AdminAgentActionReq) error {
	agent, err := s.dao.FirstAgentInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"network_id = ?":        {req.ChainID},
			"agent_contract_id = ?": {req.AgentContractId},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent == nil || agent.TwitterInfoID == 0 {
		return errs.NewError(fmt.Errorf("agent is not found"))
	}
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agent, err := s.dao.FirstAgentInfoByID(
				tx,
				agent.ID,
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			status := models.AgentSnapshotPostActionStatusNew
			if req.IsTesting {
				status = models.AgentSnapshotPostActionStatusTesting
			}
			var snapshotPostID, snapshotMissionID uint
			var toolSet models.ToolsetType
			if req.RefID != "" {
				snapshotPost, _ := s.dao.FirstAgentSnapshotPost(
					tx,
					map[string][]interface{}{
						"infer_tx_hash = ?": {req.RefID},
					},
					map[string][]interface{}{
						"AgentSnapshotMission": {},
					},
					[]string{},
				)
				if snapshotPost != nil {
					snapshotPostID = snapshotPost.ID
					snapshotMissionID = snapshotPost.AgentSnapshotMissionID
					if snapshotPost.AgentSnapshotMission != nil {
						toolSet = snapshotPost.AgentSnapshotMission.ToolSet
					}
				}
			} else if req.MissionID > 0 {
				mission, _ := s.dao.FirstAgentSnapshotMissionByID(
					tx, req.MissionID,
					map[string][]interface{}{},
					false,
				)
				if mission != nil {
					snapshotMissionID = mission.ID
					toolSet = mission.ToolSet
				}
			}
			switch req.ActionType {
			case models.AgentSnapshotPostActionTypeFollow,
				models.AgentSnapshotPostActionTypeTweet,
				models.AgentSnapshotPostActionTypeReply,
				models.AgentSnapshotPostActionTypeRetweet,
				models.AgentSnapshotPostActionTypeQuoteTweet,
				models.AgentSnapshotPostActionTypeInscribeTweet,
				models.AgentSnapshotPostActionTypeTweetV2,
				models.AgentSnapshotPostActionTypeTweetMulti,
				models.AgentSnapshotPostActionTypeReplyMulti,
				models.AgentSnapshotPostActionTypeReplyMultiUnlimited:
				{
					content := req.ActionInput.Content
					if content == "" {
						content = req.ActionInput.Comment
					}
					tweetId := req.ActionInput.Twid
					if tweetId == "" {
						tweetId = req.ActionInput.TweetId
					}
					var conversationId string
					if req.ConversationId != "" {
						conversationId = req.ConversationId
					} else {
						if tweetId != "" {
							conversationId = s.GetConversationIdByTweetID(daos.GetDBMainCtx(ctx), tweetId)
						}
					}
					content = strings.TrimSpace(content)
					var followerCount uint
					var targetTwitterId string
					targetUsername := strings.TrimPrefix(req.ActionInput.TargetUsername, "@")
					if targetUsername != "" {
					} else if tweetId != "" {
						if content == "" {
							return errs.NewError(fmt.Errorf("content is missing"))
						}
					}
					err = s.dao.Create(
						tx,
						&models.AgentSnapshotPostAction{
							NetworkID:              agent.NetworkID,
							AgentInfoID:            agent.ID,
							AgentSnapshotPostID:    snapshotPostID,
							AgentSnapshotMissionID: snapshotMissionID,
							AgentTwitterId:         agent.TwitterID,
							Type:                   req.ActionType,
							Tweetid:                tweetId,
							TargetUsername:         targetUsername,
							TargetTwitterId:        targetTwitterId,
							Content:                content,
							Status:                 status,
							FollowerCount:          followerCount,
							ScheduleAt:             helpers.TimeNow(),
							ReqRefID:               req.RefID,
							ToolSet:                toolSet,
							TokenImageUrl:          req.ActionInput.ImageUrl,
							ConversationId:         conversationId,
							InscribeTxHash:         req.InscribeTxHash,
							BitcoinTxHash:          req.BitcoinTxHash,
							Price:                  req.ActionInput.Price,
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
				}
			case models.AgentSnapshotPostActionTypeCreateToken:
				{
					err = s.dao.Create(
						tx,
						&models.AgentSnapshotPostAction{
							NetworkID:              agent.NetworkID,
							AgentInfoID:            agent.ID,
							AgentSnapshotPostID:    snapshotPostID,
							AgentSnapshotMissionID: snapshotMissionID,
							AgentTwitterId:         agent.TwitterID,
							Type:                   req.ActionType,
							Tweetid:                req.ActionInput.Twid,
							TargetUsername:         "",
							TargetTwitterId:        "",
							Content:                req.ActionInput.Content,
							Description:            req.ActionInput.Description,
							TokenName:              req.ActionInput.Name,
							TokenSymbol:            req.ActionInput.Symbol,
							Status:                 status,
							FollowerCount:          0,
							ScheduleAt:             helpers.TimeNow(),
							ReqRefID:               req.RefID,
							ToolSet:                toolSet,
							InscribeTxHash:         req.InscribeTxHash,
							BitcoinTxHash:          req.BitcoinTxHash,
							Price:                  req.ActionInput.Price,
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
				}
			default:
				{
					content := req.ActionInput.Content
					if content == "" {
						content = req.ActionInput.Comment
					}
					tweetId := req.ActionInput.Twid
					if tweetId == "" {
						tweetId = req.ActionInput.TweetId
					}
					content = strings.TrimSpace(content)
					var followerCount uint
					var targetTwitterId string
					targetUsername := strings.TrimPrefix(req.ActionInput.TargetUsername, "@")
					if targetUsername != "" {
					} else if tweetId != "" {
						if content == "" {
							return errs.NewError(fmt.Errorf("content is missing"))
						}
					}
					err = s.dao.Create(
						tx,
						&models.AgentSnapshotPostAction{
							NetworkID:              agent.NetworkID,
							AgentInfoID:            agent.ID,
							AgentSnapshotPostID:    snapshotPostID,
							AgentSnapshotMissionID: snapshotMissionID,
							AgentTwitterId:         agent.TwitterID,
							Type:                   req.ActionType,
							Tweetid:                tweetId,
							TargetUsername:         targetUsername,
							TargetTwitterId:        targetTwitterId,
							Content:                content,
							TokenImageUrl:          req.ActionInput.ImageUrl,
							Status:                 models.AgentSnapshotPostActionStatusInvalid,
							FollowerCount:          followerCount,
							ScheduleAt:             helpers.TimeNow(),
							TokenName:              req.ActionInput.Name,
							TokenSymbol:            req.ActionInput.Symbol,
							ReqRefID:               req.RefID,
							ToolSet:                toolSet,
							InscribeTxHash:         req.InscribeTxHash,
							BitcoinTxHash:          req.BitcoinTxHash,
							Price:                  req.ActionInput.Price,
						},
					)
					if err != nil {
						return errs.NewError(err)
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

func (s *Service) CreateAgentInternalActionByRefID(ctx context.Context, refID string, req *serializers.AdminAgentActionByRefReq) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			var snapshotPostID, snapshotMissionID uint
			var toolSet models.ToolsetType
			if refID != "" {
				snapshotPost, _ := s.dao.FirstAgentSnapshotPost(
					tx,
					map[string][]interface{}{
						"infer_tx_hash = ?": {refID},
					},
					map[string][]interface{}{
						"AgentSnapshotMission": {},
						"AgentInfo":            {},
					},
					[]string{},
				)
				if snapshotPost != nil && snapshotPost.AgentInfo != nil && snapshotPost.AgentSnapshotMission != nil &&
					snapshotPost.AgentSnapshotMission.ToolSet != models.ToolsetTypeTradeAnalyticsOnTwitter {
					snapshotPostID = snapshotPost.ID
					snapshotMissionID = snapshotPost.AgentSnapshotMissionID
					status := models.AgentSnapshotPostActionStatusNew
					if snapshotPost.AgentSnapshotMission != nil {
						toolSet = snapshotPost.AgentSnapshotMission.ToolSet
						if snapshotPost.AgentSnapshotMission.IsTesting {
							status = models.AgentSnapshotPostActionStatusTesting
						}
					}

					content := strings.TrimSpace(req.Reason)
					err := s.dao.Create(
						tx,
						&models.AgentSnapshotPostAction{
							NetworkID:              snapshotPost.AgentInfo.NetworkID,
							AgentInfoID:            snapshotPost.AgentInfo.ID,
							AgentSnapshotPostID:    snapshotPostID,
							AgentSnapshotMissionID: snapshotMissionID,
							AgentTwitterId:         snapshotPost.AgentInfo.TwitterID,
							Type:                   req.ActionType,
							Content:                content,
							Status:                 status,
							ScheduleAt:             helpers.TimeNow(),
							ReqRefID:               refID,
							ToolSet:                toolSet,
						},
					)
					if err != nil {
						return errs.NewError(err)
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

func (s *Service) TweetByToken(ctx context.Context, agentInfoID uint, req *serializers.AdminTweetReq) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentInfoID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if agent == nil || agent.TwitterInfoID == 0 {
		return errs.NewError(errs.ErrBadRequest)
	}
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agent, err := s.dao.FirstAgentInfoByID(
				tx,
				agent.ID,
				map[string][]interface{}{
					"TwitterInfo": {},
				},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agent != nil && agent.TwitterInfo != nil {
				refId, err := helpers.PostTweetByToken(agent.TwitterInfo.AccessToken, req.Text, "")
				if err != nil {
					return errs.NewError(err)
				}
				//create
				err = s.dao.Create(tx, &models.TwitterTweet{
					TwitterID: agent.TwitterInfo.TwitterID,
					TweetID:   refId,
					FullText:  req.Text,
					PostedAt:  time.Now(),
				})
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
}
