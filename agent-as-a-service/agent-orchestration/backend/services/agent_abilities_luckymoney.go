package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) JobLuckyMoneyCollectPost(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobLuckyMoneyCollectPost",
		func() error {
			var retErr error
			actions, err := s.dao.FindAgentSnapshotPostActionJoin(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"join agent_snapshot_missions on agent_snapshot_missions.id = agent_snapshot_post_actions.agent_snapshot_mission_id": {},
					"join agent_infos on agent_infos.id = agent_snapshot_post_actions.agent_info_id":                                     {},
				},
				map[string][]interface{}{
					"agent_snapshot_post_actions.status = ?":   {models.AgentSnapshotPostActionStatusDone},
					"agent_snapshot_post_actions.tool_set = ?": {models.ToolsetTypeLuckyMoneys},
					// `(
					// 	agent_infos.agent_type = 1
					// 	and agent_infos.reply_enabled = true
					// )`: {},
					// `(
					// 	agent_snapshot_missions.enabled = 1
					// 	and agent_snapshot_missions.reply_enabled = 1
					// 	and agent_snapshot_missions.interval_sec > 0
					// 	and agent_snapshot_missions.is_testing = 0
					// 	and agent_snapshot_missions.deleted_at is null
					// )`: {},
				},
				map[string][]interface{}{},
				[]string{
					"rand()",
				},
				0,
				100,
			)
			if err != nil {
				return errs.NewError(err)
			}

			for _, action := range actions {
				err = s.LuckyMoneyCollectPost(ctx, action.ID)
				if err != nil {
					return errs.NewError(err)
				}

				err = s.LuckyMoneyValidateRewardUser(ctx, action.ID)
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) LuckyMoneyCollectPost(ctx context.Context, missionID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("LuckyMoneyCollectPost_%d", missionID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					snapshotPostAction, err := s.dao.FirstAgentSnapshotPostActionByID(
						tx,
						missionID,
						map[string][]interface{}{
							"AgentSnapshotMission": {},
							"AgentSnapshotPost":    {},
							"AgentInfo":            {},
						},
						false,
					)

					if err != nil {
						return errs.NewError(err)
					}

					if snapshotPostAction.AgentInfo != nil &&
						snapshotPostAction.AgentSnapshotPost != nil && snapshotPostAction.AgentSnapshotMission != nil &&
						snapshotPostAction.Status == models.AgentSnapshotPostActionStatusDone {
						twitterInfo, err := s.dao.FirstTwitterInfo(tx,
							map[string][]interface{}{
								"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
							},
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}

						if twitterInfo != nil {
							query := fmt.Sprintf("conversation_id:%s", snapshotPostAction.RefId)
							tweetRecentSearch, err := s.twitterWrapAPI.SearchRecentTweet(query, "", twitterInfo.AccessToken, 50)
							if err != nil {
								return errs.NewTwitterError(err)
							}

							for _, v := range tweetRecentSearch.LookUps {
								existPosts, err := s.dao.FirstAbilityLuckyMoney(
									tx,
									map[string][]interface{}{
										"tweet_id = ?": {v.Tweet.ID},
									},
									map[string][]interface{}{},
									[]string{},
								)
								if err != nil {
									return errs.NewError(err)
								}
								if existPosts == nil {
									postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
									fullText := v.Tweet.NoteTweet.Text
									if fullText == "" {
										fullText = v.Tweet.Text
									}

									rewardAmount := models.QuoBigFloats(&snapshotPostAction.AgentSnapshotMission.RewardAmount.Float, big.NewFloat(float64(snapshotPostAction.AgentSnapshotMission.RewardUser)))
									etherAddress := helpers.ExtractEtherAddress(fullText)
									tokenBalance := big.NewFloat(0)
									if snapshotPostAction.AgentInfo.TokenAddress != "" && etherAddress != "" {
										balance, err := s.GetEthereumClient(ctx, snapshotPostAction.AgentInfo.TokenNetworkID).Erc20Balance(snapshotPostAction.AgentInfo.TokenAddress, etherAddress)
										if err != nil {
											fmt.Println(err.Error())
										}
										tokenBalance = models.ConvertWeiToBigFloat(balance, 18)
									}

									m := &models.AbilityLuckyMoney{
										NetworkID:                 snapshotPostAction.AgentInfo.NetworkID,
										AgentInfoID:               snapshotPostAction.AgentInfo.ID,
										AgentSnapshotMissionID:    snapshotPostAction.AgentSnapshotMissionID,
										AgentSnapshotPostID:       snapshotPostAction.AgentSnapshotPostID,
										AgentSnapshotPostActionID: snapshotPostAction.ID,
										TwitterID:                 v.User.ID,
										TwitterUsername:           v.User.UserName,
										TwitterName:               v.User.Name,
										TweetID:                   v.Tweet.ID,
										Content:                   fullText,
										Status:                    models.LuckyMoneyStatusNew,
										TweetAt:                   postedAt,
										UserAddress:               etherAddress,
										RewardAmount:              numeric.NewBigFloatFromFloat(rewardAmount),
										TokenBalance:              numeric.BigFloat{*tokenBalance},
									}

									if etherAddress == "" || tokenBalance.Cmp(&snapshotPostAction.AgentSnapshotMission.MinTokenHolding.Float) < 0 {
										m.Status = models.LuckyMoneyStatusInvalid
									}

									err = s.dao.Create(tx, m)
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
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) LuckyMoneyValidateRewardUser(ctx context.Context, actionID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("LuckyMoneyPayToUser_%d", actionID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					snapshotPostAction, err := s.dao.FirstAgentSnapshotPostActionByID(
						tx,
						actionID,
						map[string][]interface{}{},
						false,
					)

					if err != nil {
						return errs.NewError(err)
					}

					userPosts, err := s.dao.FindAbilityLuckyMoney(
						tx,
						map[string][]interface{}{
							"agent_snapshot_post_action_id = ? ": {actionID},
							"status in (?)":                      {[]models.LuckyMoneyStatus{models.LuckyMoneyStatusNew, models.LuckyMoneyStatusDone, models.LuckyMoneyStatusProcessing}},
						},
						map[string][]interface{}{},
						[]string{"tweet_at asc"}, 0, snapshotPostAction.RewardUser,
					)

					if err != nil {
						return errs.NewError(err)
					}

					if userPosts != nil {
						for _, item := range userPosts {
							if item.Status == models.LuckyMoneyStatusNew {
								item.Status = models.LuckyMoneyStatusProcessing
								err = s.dao.Save(tx, item)
								if err != nil {
									return errs.NewError(err)
								}
							}
						}

						if len(userPosts) == snapshotPostAction.RewardUser {
							_ = daos.GetDBMainCtx(ctx).Model(snapshotPostAction).Updates(
								map[string]interface{}{
									"status": models.AgentSnapshotPostActionStatusPaid,
								},
							).Error
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
	return nil
}

func (s *Service) JobLuckyMoneyProcessUserReward(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobLuckyMoneyCollectPost",
		func() error {
			var retErr error
			actions, err := s.dao.FindAbilityLuckyMoney(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"status = ?": {models.LuckyMoneyStatusProcessing},
				},
				map[string][]interface{}{},
				[]string{},
				0,
				100,
			)
			if err != nil {
				return errs.NewError(err)
			}

			for _, action := range actions {
				err = s.LuckyMoneyProcessUserReward(ctx, action.ID)
				if err != nil {
					return errs.NewError(err)
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

func (s *Service) LuckyMoneyProcessUserReward(ctx context.Context, luckyID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("LuckyMoneyProcessUserReward_%d", luckyID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					m, err := s.dao.FirstAbilityLuckyMoneyByID(
						tx,
						luckyID,
						map[string][]interface{}{},
						true,
					)

					if err != nil {
						return errs.NewError(err)
					}

					if m != nil && m.Status == models.LuckyMoneyStatusProcessing {
						err = tx.Model(&models.AgentInfo{}).
							Where("id = ?", m.AgentInfoID).
							UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", m.RewardAmount)).
							Error
						if err != nil {
							return errs.NewError(err)
						}

						agentInfo, err := s.dao.FirstAgentInfoByID(
							tx,
							m.AgentInfoID,
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}

						if agentInfo.EaiBalance.Float.Cmp(big.NewFloat(0)) < 0 {
							return errs.NewError(errs.ErrBadRequest)
						}

						luckyMoneyAdminAddress := strings.ToLower(s.conf.LuckyMoneyAdminAddress)
						txHash, err := s.GetEthereumClient(ctx, models.ETERNAL_AI_CHAIN_ID).Transfer(
							s.GetAddressPrk(luckyMoneyAdminAddress), m.UserAddress,
							models.ConvertBigFloatToWei(&m.RewardAmount.Float, 18).String(),
							false, true,
						)

						if err != nil {
							m.Error = err.Error()
						} else {
							m.TxHash = txHash
							m.Status = models.LuckyMoneyStatusDone
						}
						err = s.dao.Save(tx, m)
						if err != nil {
							return errs.NewError(err)
						}

						if m.Status == models.LuckyMoneyStatusDone {
							_ = s.dao.Create(
								tx,
								&models.AgentEaiTopup{
									NetworkID:      agentInfo.NetworkID,
									EventId:        fmt.Sprintf("agent_lucky_money_%d", m.ID),
									AgentInfoID:    agentInfo.ID,
									Type:           models.AgentEaiTopupTypeSpent,
									Amount:         m.RewardAmount,
									Status:         models.AgentEaiTopupStatusDone,
									DepositAddress: m.UserAddress,
									ToAddress:      m.UserAddress,
									Toolset:        string(models.ToolsetTypeLuckyMoneys),
								},
							)
						}
						//reply to user
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
	return nil
}

func (s *Service) JobLuckyMoneyActionExecuted(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobLuckyMoneyActionExecuted",
		func() error {
			var retErr error
			{
				ms, err := s.dao.FindAgentSnapshotPost(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"toolset = ?": {models.ToolsetTypeLuckyMoneys},
						"status = ?":  {models.AgentSnapshotPostStatusInferSubmitted},
					},
					map[string][]interface{}{},
					[]string{
						"updated_at asc",
					}, 0, 999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, m := range ms {
					err := s.LuckyMoneyActionExecuted(ctx, m.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, m.ID))
					}
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

func (s *Service) LuckyMoneyActionExecuted(ctx context.Context, snapshotPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("LuckyMoneyActionExecuted_%d", snapshotPostID),
		func() error {
			err := func() error {
				err := daos.WithTransaction(
					daos.GetDBMainCtx(ctx),
					func(tx *gorm.DB) error {
						snapshotPost, err := s.dao.FirstAgentSnapshotPostByID(
							tx,
							snapshotPostID,
							map[string][]interface{}{},
							true,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if snapshotPost != nil &&
							models.ToolsetType(snapshotPost.Toolset) == models.ToolsetTypeLuckyMoneys {
							randomTime := helpers.TimeNow().Add(time.Minute * time.Duration(helpers.RandomInt(5, 40)))
							content, err := s.LuckyMoneyGetPostContent(tx, snapshotPost.AgentInfoID, snapshotPost.AgentSnapshotMissionID)
							if err != nil {
								return errs.NewError(err)
							}
							if content == "" {
								return errs.NewError(errs.ErrBadContent)
							}

							agentInfo, err := s.dao.FirstAgentInfoByID(tx, snapshotPost.AgentInfoID, map[string][]interface{}{}, false)
							if err != nil {
								return errs.NewError(err)
							}

							mission, err := s.dao.FirstAgentSnapshotMissionByID(tx, snapshotPost.AgentSnapshotMissionID, map[string][]interface{}{}, false)
							if err != nil {
								return errs.NewError(err)
							}

							status := models.AgentSnapshotPostActionStatusNew
							if mission != nil {
								if mission.IsTesting {
									status = models.AgentSnapshotPostActionStatusTesting
								}
							}

							action := &models.AgentSnapshotPostAction{
								NetworkID:              snapshotPost.NetworkID,
								AgentInfoID:            snapshotPost.AgentInfoID,
								AgentSnapshotPostID:    snapshotPostID,
								AgentSnapshotMissionID: snapshotPost.AgentSnapshotMissionID,
								AgentTwitterId:         agentInfo.TwitterID,
								ScheduleAt:             &randomTime,
								ReqRefID:               snapshotPost.InferTxHash,
								ToolSet:                models.ToolsetType(snapshotPost.Toolset),
								Content:                content,
								Status:                 status,
								Type:                   models.AgentSnapshotPostActionTypeTweet,
								RewardUser:             mission.RewardUser,
								RewardAmount:           mission.RewardAmount,
							}

							err = s.dao.Create(tx, action)
							if err != nil {
								return errs.NewError(err)
							}

							snapshotPost.Status = models.AgentSnapshotPostStatusInferResolved
							err = s.dao.Save(tx, snapshotPost)
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
			}()
			return err
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) LuckyMoneyGetPostContent(tx *gorm.DB, agentInfoID, missionID uint) (string, error) {
	postContent := ""
	agentInfo, err := s.dao.FirstAgentInfoByID(tx, agentInfoID, map[string][]interface{}{}, false)
	if err != nil {
		return postContent, errs.NewError(err)
	}

	missionInfo, err := s.dao.FirstAgentSnapshotMissionByID(tx, missionID, map[string][]interface{}{}, false)
	if err != nil {
		return postContent, errs.NewError(err)
	}

	if agentInfo != nil && missionInfo != nil {
		rewardAmount, _ := missionInfo.RewardAmount.Float64()
		minTokenHolding, _ := missionInfo.MinTokenHolding.Float64()
		userPrompt := fmt.Sprintf(`
		Rewrite this for a Twitter post: 
		"Lucky Money Giveaway! 💸 Total %d tokens $EAI up for grabs! First %d comments with an Ethereum address (holding min %d tokens $%s) win! 🚀 Fastest fingers only!"

		Return a JSON response with the following format:

		{"content": ""}
		
		Respond with only the JSON string, without any additional explanation.
		`, rewardAmount, missionInfo.RewardUser, minTokenHolding, agentInfo.TokenSymbol)
		fmt.Println(userPrompt)

		aiStr, err := s.openais["Lama"].ChatMessageWithSystemPromp(strings.TrimSpace(userPrompt), agentInfo.GetSystemPrompt())
		if err != nil {
			return postContent, nil
		}

		if aiStr != "" {
			mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
			if mapInfo != nil {
				if v, ok := mapInfo["content"]; ok {
					postContent = fmt.Sprintf(`%v`, v)
				}
			}
		}
	}
	return postContent, nil
}

func (s *Service) TestUtil() {
	// etherAddress := helpers.ExtractEtherAddress("yo 0x7c9d59cD31F27c7cBEEde2567c9fa377537bdDE0 😄😁😋")
	// fmt.Println(etherAddress)
	twIDs := []string{"1878850921289130415"}
	twitterInfo, _ := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(context.Background()),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterID},
		},
		map[string][]interface{}{},
		false,
	)
	twitterDetail, _ := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
	if twitterDetail != nil {
		for _, v := range *twitterDetail {
			fmt.Println(v.Tweet.Text)
			fulltext, ismention := s.TweetIsMentionNBS(v.Tweet, "NOBULLSHIT_EXE")
			fmt.Println(fulltext)
			fmt.Println(ismention)
		}
	}
}
