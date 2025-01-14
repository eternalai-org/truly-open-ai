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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/jinzhu/gorm"
)

func (s *Service) ScanAgentTwitterPostFroCreateLaunchpad(ctx context.Context) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		s.conf.LaunchpadAgentInfoId,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {agent.TwitterID},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if twitterInfo != nil {
		err = func() error {
			tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(twitterInfo.TwitterID, "", twitterInfo.AccessToken)
			if err != nil {
				return errs.NewError(err)
			}
			err = s.CreateAgentTwitterPostForCreateLaunchpad(daos.GetDBMainCtx(ctx), agent.ID, agent.TwitterUsername, tweetMentions)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		}()
		if err != nil {
			s.UpdateAgentScanEventError(ctx, agent.ID, err)
			return err
		} else {
			err = s.UpdateAgentScanEventSuccess(ctx, agent.ID, nil, "")
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) CreateAgentTwitterPostForCreateLaunchpad(tx *gorm.DB, agentInfoID uint, twitterUsername string, tweetMentions *twitter.UserTimeline) error {
	if tweetMentions != nil {
		agentInfo, err := s.dao.FirstAgentInfoByID(
			tx,
			agentInfoID,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if agentInfo == nil {
			return errs.NewError(errs.ErrBadRequest)
		}
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(errs.ErrBadRequest)
		}
		for _, item := range tweetMentions.Tweets {
			if !strings.EqualFold(item.AuthorID, agentInfo.TwitterID) {
				var rs bool
				err := s.RedisCached(
					fmt.Sprintf("CreateAgentTwitterPostForCreateLaunchpad_%s", item.ID),
					true,
					12*time.Hour,
					&rs,
					func() (interface{}, error) {
						err := func() error {
							author, err := s.CreateUpdateUserTwitter(tx, item.AuthorID)
							if err != nil {
								return errs.NewError(errs.ErrBadRequest)
							}
							if author != nil {
								twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, []string{item.ID})
								if err != nil {
									return errs.NewError(err)
								}
								if twitterDetail != nil {
									for k, v := range *twitterDetail {
										if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
											if strings.EqualFold(k, item.ID) {
												existPosts, err := s.dao.FirstAgentTwitterPost(
													tx,
													map[string][]interface{}{
														"twitter_post_id = ?": {v.Tweet.ID},
													},
													map[string][]interface{}{},
													[]string{},
												)
												if err != nil {
													return errs.NewError(err)
												}
												if existPosts == nil {
													fullText := v.Tweet.NoteTweet.Text
													if fullText == "" {
														fullText = v.Tweet.Text
													}
													var prjDesc string
													if v.Tweet.ConversationID == "" || v.Tweet.ConversationID == v.Tweet.ID {
														prjInfo, err := s.GetAgentCreateLaunchpad(context.Background(), v.User.UserName, fullText)
														if err != nil {
															return errs.NewError(err)
														}
														prjDesc = prjInfo.Description
													}
													postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
													if prjDesc != "" {
														m := &models.AgentTwitterPost{
															NetworkID:             agentInfo.NetworkID,
															AgentInfoID:           agentInfo.ID,
															TwitterID:             v.User.ID,
															TwitterUsername:       v.User.UserName,
															TwitterName:           v.User.Name,
															TwitterPostID:         v.Tweet.ID,
															Content:               fullText,
															Status:                models.AgentTwitterPostStatusNew,
															PostAt:                postedAt,
															TwitterConversationId: v.Tweet.ConversationID,
															PostType:              models.AgentSnapshotPostActionTypeCreateLaunchpad,
															IsMigrated:            true,
															TokenDesc:             prjDesc,
														}
														m.OwnerTwitterID = m.TwitterID
														m.OwnerUsername = m.TwitterUsername
														err = s.dao.Create(tx, m)
														if err != nil {
															return errs.NewError(err)
														}
													} else {
														m := &models.AgentTwitterPost{
															NetworkID:             agentInfo.NetworkID,
															AgentInfoID:           agentInfo.ID,
															TwitterID:             v.User.ID,
															TwitterUsername:       v.User.UserName,
															TwitterName:           v.User.Name,
															TwitterPostID:         v.Tweet.ID,
															Content:               fullText,
															Status:                models.AgentTwitterPostStatusInvalid,
															PostAt:                postedAt,
															TwitterConversationId: v.Tweet.ConversationID,
															PostType:              models.AgentSnapshotPostActionTypeUnknown,
															IsMigrated:            true,
														}
														m.OwnerTwitterID = m.TwitterID
														m.OwnerUsername = m.TwitterUsername
														err = s.dao.Create(tx, m)
														if err != nil {
															return errs.NewError(err)
														}
													}
													_, _ = s.CreateUpdateUserTwitter(tx, v.User.ID)
												}
											}
										}
									}
								}
							}
							return nil
						}()
						if err != nil {
							return false, errs.NewError(err)
						}
						return true, nil
					},
				)
				if err != nil {
					return errs.NewError(err)
				}
			}
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) JobAgentTwitterPostCreateLaunchpad(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateLaunchpad",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.LaunchpadAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateLaunchpad},
					},
					map[string][]interface{}{},
					[]string{
						"post_at desc",
					},
					0,
					5,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterPost := range twitterPosts {
					err = s.AgentTwitterPostCreateLaunchpad(ctx, twitterPost.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewErrorWithId(err, twitterPost.ID))
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

func (s *Service) AgentTwitterPostCreateLaunchpad(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostCreateAgent_%d", twitterPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
						tx,
						twitterPostID,
						map[string][]interface{}{
							"AgentInfo":             {},
							"AgentInfo.TwitterInfo": {},
						},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if twitterPost.Status == models.AgentTwitterPostStatusNew &&
						twitterPost.PostType == models.AgentSnapshotPostActionTypeCreateLaunchpad {
						solAddress, err := s.CreateSOLAddress(ctx)
						if err != nil {
							return errs.NewError(err)
						}
						lp := &models.Launchpad{
							TwitterPostID:   twitterPost.ID,
							TweetId:         twitterPost.TwitterPostID,
							TwitterId:       twitterPost.TwitterID,
							TwitterUsername: twitterPost.TwitterUsername,
							TwitterName:     twitterPost.TwitterName,
							Address:         solAddress,
							Description:     twitterPost.Content,
							Name:            twitterPost.TokenDesc,
						}
						err = s.dao.Create(tx, lp)
						if err != nil {
							return errs.NewError(err)
						}
						twitterPost.Status = models.AgentTwitterPostStatusReplied
						err = s.dao.Save(tx, twitterPost)
						if err != nil {
							return errs.NewError(err)
						}
						err = s.ReplyAferAutoCreateLaunchpad(tx, twitterPost.ID, lp.ID)
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
	return nil
}

func (s *Service) GetAgentCreateLaunchpad(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{}
	fullText = strings.ReplaceAll(fullText, "@"+userName, "")
	userPrompt := fmt.Sprintf(`
Detect Fundraising Requests
This is the user conversation: %s

From this conversation determine if the user is requesting assistance with fundraising, look for a direct and unambiguous statement that explicitly asks to assistance with fundraising. This statement must be clear, concise, and isolated from any surrounding context that may alter its meaning.

If yes, extract or generate the following information:

Answer ("yes" or "no")
Project information (generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI")

Return a JSON response with the following format:
{"answer": "yes/no", "project-information": ""}

Respond with only the JSON string, without any additional explanation.
	`, fullText)
	aiStr, err := s.openais["Lama"].ChatMessage(strings.TrimSpace(userPrompt))
	if err != nil {
		return info, nil
	}
	if aiStr != "" {
		mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
		if mapInfo != nil {
			answer := "no"
			if v, ok := mapInfo["answer"]; ok {
				answer = fmt.Sprintf(`%v`, v)
			}
			if strings.EqualFold(answer, "yes") {
				info.IsCreateAgent = true
				if v, ok := mapInfo["project-information"]; ok {
					info.Description = fmt.Sprintf(`%v`, v)
				}
			}

		}
	}
	return info, nil
}

func (s *Service) ReplyAferAutoCreateLaunchpad(tx *gorm.DB, twitterPostID, launchpadId uint) error {
	if twitterPostID > 0 && launchpadId > 0 {
		twitterPost, err := s.dao.FirstAgentTwitterPostByID(
			tx,
			twitterPostID,
			map[string][]interface{}{
				"AgentInfo.TwitterInfo": {},
			},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		launchpad, err := s.dao.FirstLaunchpadByID(
			tx,
			launchpadId,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}
		if twitterPost != nil && launchpad != nil && twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil && twitterPost.ReplyPostId == "" {
			replyContent := fmt.Sprintf(`
We're excited to announce new raise fund project %s, empowering decentralized AI innovation with community-owned compute power.

Receiving fund address here: %s

Whitelist applications are now open!
			`, launchpad.Name, launchpad.Address)
			replyContent = strings.TrimSpace(replyContent)
			refId, err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, replyContent, twitterPost.TwitterPostID, "")
			if err != nil {
				_ = tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"error":         err.Error(),
					},
				)
			} else {
				_ = tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"reply_content": replyContent,
						"reply_post_at": helpers.TimeNow(),
						"reply_post_id": refId,
						"error":         "",
					},
				).Error
			}
		}
	}
	return nil
}
