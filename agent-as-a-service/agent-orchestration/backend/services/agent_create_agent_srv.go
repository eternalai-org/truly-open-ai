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
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/mymmrac/telego"
)

func (s *Service) JobScanAgentTwitterPostForCreateAgent(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobScanAgentTwitterPostForCreateAgent",
		func() error {
			agents, err := s.dao.FindAgentInfo(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					// `id in (?)`: {[]uint{s.conf.BaseAiAgentInfoId, s.conf.SolanaAiAgentInfoId}},
					`id in (?)`: {[]uint{s.conf.EternalAiAgentInfoId}},
				},
				map[string][]interface{}{},
				[]string{
					"scan_latest_time asc",
				},
				0,
				2,
			)
			if err != nil {
				return errs.NewError(err)
			}
			var retErr error
			for _, agent := range agents {
				err = s.ScanAgentTwitterPostFroCreateAgent(ctx, agent.ID)
				if err != nil {
					retErr = errs.MergeError(retErr, errs.NewError(err))
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

func (s *Service) ScanAgentTwitterPostFroCreateAgent(ctx context.Context, agentID uint) error {
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		agentID,
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterID},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return errs.NewError(err)
	}
	if twitterInfo != nil {
		err = func() error {
			tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(agent.TwitterID, "", twitterInfo.AccessToken)
			if err != nil {
				return errs.NewError(err)
			}
			err = s.CreateAgentTwitterPostForCreateAgent(daos.GetDBMainCtx(ctx), agent.ID, agent.TwitterUsername, tweetMentions)
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

func (s *Service) CreateAgentTwitterPostForCreateAgent(tx *gorm.DB, agentInfoID uint, twitterUsername string, tweetMentions *twitter.UserTimeline) error {
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
			var checkTwitterID string
			err := s.GetRedisCachedWithKey(fmt.Sprintf("CheckedForCreateAgent_%s", item.ID), &checkTwitterID)
			if err != nil {
				if !strings.EqualFold(item.AuthorID, agentInfo.TwitterID) {
					author, err := s.CreateUpdateUserTwitter(tx, item.AuthorID)
					if err != nil {
						return errs.NewError(errs.ErrBadRequest)
					}
					if author != nil {
						// listContext, err := s.GetConversionHistory(tx, item.ID)
						// if err != nil {
						// 	return errs.NewError(errs.ErrBadRequest)
						// }

						// jsonString, _ := json.Marshal(listContext)

						// tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), author.TwitterUsername, string(jsonString))
						// if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
						twIDs := []string{item.ID}
						twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
						if err != nil {
							return errs.NewError(err)
						}
						if twitterDetail != nil {
							for k, v := range *twitterDetail {
								if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
									if strings.EqualFold(k, item.ID) {
										fullText := v.Tweet.NoteTweet.Text
										if fullText == "" {
											fullText = v.Tweet.Text
										}

										tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), author.TwitterUsername, fullText)
										if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
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
												postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
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
													PostType:              models.AgentSnapshotPostActionTypeReply,
													IsMigrated:            true,
												}

												m.TokenSymbol = tokenInfo.TokenSymbol
												m.TokenName = tokenInfo.TokenName
												m.TokenDesc = tokenInfo.TokenDesc
												m.Prompt = tokenInfo.Personality
												m.AgentChain = tokenInfo.ChainName
												m.PostType = models.AgentSnapshotPostActionTypeCreateAgent
												m.OwnerTwitterID = m.TwitterID
												m.OwnerUsername = m.TwitterUsername
												if strings.EqualFold(tokenInfo.ChainName, "base") && tokenInfo.IsIntellect {
													m.ExtractContent = "PrimeIntellect/INTELLECT-1-Instruct"
												}

												// if tokenInfo.Owner != "" {
												// 	twUser, _ := s.CreateUpdateUserTwitterByUserName(tx, tokenInfo.Owner)
												// 	if twUser != nil {
												// 		m.OwnerTwitterID = twUser.TwitterID
												// 		m.OwnerUsername = twUser.TwitterUsername
												// 	}
												// }

												err = s.dao.Create(tx, m)
												if err != nil {
													return errs.NewError(err)
												}

												_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
											}
										}
									}
								}
							}
						}
						// }
					}
				}
			}

			err = s.SetRedisCachedWithKey(
				fmt.Sprintf("CheckedForCreateAgent_%s", item.ID),
				item.ID,
				12*time.Hour,
			)
			if err != nil {
				return errs.NewError(err)
			}
		}
	}
	return nil
}

func (s *Service) JobAgentTwitterPostCreateAgent(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTwitterPostCreateAgent",
		func() error {
			var retErr error
			{
				twitterPosts, err := s.dao.FindAgentTwitterPost(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_info_id in (?)": {[]uint{s.conf.EternalAiAgentInfoId}},
						"status = ?":           {models.AgentTwitterPostStatusNew},
						"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateAgent},
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
					err = s.AgentTwitterPostCreateAgent(ctx, twitterPost.ID)
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

func (s *Service) AgentTwitterPostCreateAgent(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostCreateAgent_%d", twitterPostID),
		func() error {
			agentID := uint(0)
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

					isValid := true
					existPosts, err := s.dao.FindAgentTwitterPost(
						tx,
						map[string][]interface{}{
							"not EXISTS (select 1 from agent_twitter_posts atp2 where twitter_conversation_id=? and owner_twitter_id =? and post_type='create_agent' and twitter_post_id != agent_twitter_posts.twitter_post_id )": {twitterPost.TwitterConversationId, twitterPost.OwnerTwitterID},
							"owner_twitter_id = ?": {twitterPost.OwnerTwitterID},
							"post_type = ?":        {models.AgentSnapshotPostActionTypeCreateAgent},
							"status = ?":           {models.AgentTwitterPostStatusReplied},
							"created_at >= adddate(now(), interval -24 hour)": {},
						},
						map[string][]interface{}{},
						[]string{}, 0, 5,
					)
					if err != nil {
						return errs.NewError(err)
					}

					if existPosts != nil && len(existPosts) >= 3 {
						isValid = false
					}

					if isValid {
						if twitterPost.Status == models.AgentTwitterPostStatusNew &&
							twitterPost.PostType == models.AgentSnapshotPostActionTypeCreateAgent &&
							twitterPost.AgentInfo != nil {
							networkID := models.BASE_CHAIN_ID
							if twitterPost.AgentChain != "" {
								networkID = models.GetChainID(twitterPost.AgentChain)
							}

							tokenNetworkID := models.BASE_CHAIN_ID
							if networkID == models.APE_CHAIN_ID {
								tokenNetworkID = networkID
							}

							agentInfo := &models.AgentInfo{
								NetworkID:      networkID,
								NetworkName:    models.GetChainName(networkID),
								SystemPrompt:   twitterPost.Prompt,
								AgentName:      twitterPost.TokenName,
								TokenMode:      string(models.TokenSetupEnumAutoCreate),
								AgentType:      models.AgentInfoAgentTypeReasoning,
								TmpTwitterID:   twitterPost.GetOwnerTwitterID(),
								TokenNetworkID: tokenNetworkID,
								Version:        "2",
								AgentID:        helpers.RandomBigInt(12).Text(16),
								ScanEnabled:    true,
								Creator:        strings.ToLower(s.conf.GetConfigKeyString(models.BASE_CHAIN_ID, "meme_pool_address")),
							}

							if networkID == models.BASE_CHAIN_ID && twitterPost.ExtractContent != "" {
								agentInfo.AgentBaseModel = twitterPost.ExtractContent
							}

							if agentInfo.AgentBaseModel == "" {
								agentInfo.AgentBaseModel = s.GetModelDefaultByChainID(agentInfo.NetworkID)
							}

							agentInfo.AgentFee = models.GetAgentFee(agentInfo.NetworkID)
							ethAddress, err := s.CreateETHAddress(ctx)
							if err != nil {
								return errs.NewError(err)
							}
							agentInfo.ETHAddress = strings.ToLower(ethAddress)
							solAddress, err := s.CreateSOLAddress(ctx)
							if err != nil {
								return errs.NewError(err)
							}
							agentInfo.SOLAddress = solAddress
							err = s.dao.Create(tx, agentInfo)
							if err != nil {
								return errs.NewError(err)
							}
							agentInfo.TokenMode = string(models.TokenSetupEnumAutoCreate)
							agentInfo.TokenName = twitterPost.TokenName
							agentInfo.TokenSymbol = twitterPost.TokenSymbol
							agentInfo.TokenDesc = twitterPost.TokenDesc
							agentInfo.TokenNetworkID = tokenNetworkID
							agentInfo.SystemPrompt = twitterPost.Prompt
							agentInfo.MetaData = twitterPost.Prompt
							agentInfo.TokenStatus = "pending"
							agentInfo.EaiBalance = numeric.NewBigFloatFromString("50")
							agentInfo.Status = models.AssistantStatusPending

							agentTokenInfo := &models.AgentTokenInfo{}
							agentTokenInfo.AgentInfoID = agentInfo.ID
							agentTokenInfo.NetworkID = tokenNetworkID
							agentTokenInfo.NetworkName = models.GetChainName(agentTokenInfo.NetworkID)
							err = s.dao.Create(tx, agentTokenInfo)
							if err != nil {
								return errs.NewError(err)
							}

							agentInfo.TokenInfoID = agentTokenInfo.ID
							agentInfo.RefTweetID = twitterPost.ID
							err = s.dao.Save(tx, agentInfo)
							if err != nil {
								return errs.NewError(err)
							}

							twitterPost.Status = models.AgentTwitterPostStatusReplied
							err = s.dao.Save(tx, twitterPost)
							if err != nil {
								return errs.NewError(err)
							}
							agentID = agentInfo.ID
						}
					} else {
						twitterPost.Status = models.AgentTwitterConversationInvalid
						err = s.dao.Save(tx, twitterPost)
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

			if agentID > 0 {
				_ = s.CreateTokenInfo(ctx, agentID)
				// _ = s.AgentMintNft(ctx, agentID)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) getContentTwiterForCreateAgent(ownerName, agentName, tokenSymbol, tokenDesc, tokenAddress string) string {
	replyContent := fmt.Sprintf(`
Hey @%s, your Eternal AI agent $%s is live. Be the first to buy its AI coin:

https://eternalai.org/agent/%s

%s ($%s): %s

PS: You can activate its autonomous tweeting ability (and soon DeFi trading ability) at https://eternalai.org/connect-x
`, ownerName, tokenSymbol, tokenAddress, agentName, tokenSymbol, tokenDesc)
	return strings.TrimSpace(replyContent)
}

func (s *Service) getContentTwiterForCreateAgentV1(ownerName, agentName, tokenSymbol, tokenDesc, tokenAddress, chainName string, agentID uint) string {
	replyContent := fmt.Sprintf(`
Hey @%s, your Eternal AI agent %s is now live on %s!

Be the first to buy its token:

https://eternalai.org/%d

%s ($%s): %s

`, ownerName, agentName, chainName, agentID, agentName, tokenSymbol, tokenDesc)
	return strings.TrimSpace(replyContent)
}

func (s *Service) ReplyAferAutoCreateAgent(tx *gorm.DB, twitterPostID, agentInfoId uint) error {
	if twitterPostID > 0 && agentInfoId > 0 {
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

		agentInfo, err := s.dao.FirstAgentInfoByID(
			tx,
			agentInfoId,
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return errs.NewError(err)
		}

		if twitterPost != nil && agentInfo != nil && twitterPost.AgentInfo != nil && twitterPost.AgentInfo.TwitterInfo != nil && twitterPost.ReplyPostId == "" {
			// replyContent := s.getContentTwiterForCreateAgent(twitterPost.GetAgentOnwerName(), agentInfo.AgentName, agentInfo.TokenSymbol, agentInfo.TokenDesc, agentInfo.TokenAddress)
			replyContent := s.getContentTwiterForCreateAgentV1(twitterPost.GetAgentOnwerName(), agentInfo.AgentName, agentInfo.TokenSymbol,
				agentInfo.TokenDesc, agentInfo.TokenAddress, agentInfo.NetworkName, agentInfo.ID)
			refId, err := helpers.ReplyTweetByToken(twitterPost.AgentInfo.TwitterInfo.AccessToken, replyContent, twitterPost.TwitterPostID)
			if err != nil {
				tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"error": err.Error(),
					},
				)
			} else {
				_ = tx.Model(twitterPost).Updates(
					map[string]interface{}{
						"reply_post_at": helpers.TimeNow(),
						"reply_post_id": refId,
						"error":         "",
					},
				).Error

				//noti tele
				bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
				if err != nil {
					return errs.NewError(err)
				}
				title := "🟡🟡🟡 New Agent Alert! 🟡🟡🟡"
				msg := fmt.Sprintf(`just created agent %s https://eternalai.org/agent/%s 👀`, agentInfo.TokenSymbol, agentInfo.TokenAddress)
				_, err = bot.SendMessage(
					&telego.SendMessageParams{
						ChatID: telego.ChatID{
							ID: s.conf.Telebot.Alert.ChatID,
						},
						Text: strings.TrimSpace(
							fmt.Sprintf(
								`
%s

%s (@%s) %s

Hey, @JohnEnt, let's do something about this!
	`,
								title,
								twitterPost.TwitterName,
								twitterPost.TwitterUsername,
								msg,
							),
						),
					},
				)
				if err != nil {
					return errs.NewError(err)
				}
			}
		}
	}
	return nil
}

func (s *Service) GetImageUrlFromTokenInfo(tokenSymbol, tokenName, tokenDesc string) (string, error) {
	stringBase64 := s.GenerateTokenImageBase64(context.Background(), tokenSymbol, tokenName, tokenDesc)
	if stringBase64 != "" {
		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "jpg")
		urlPath, err := s.gsClient.UploadPublicDataBase64("agent", filename, stringBase64)
		if err != nil {
			return "", errs.NewError(err)
		}
		return fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath), nil
	}
	return "", errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetGifImageUrlFromTokenInfo(tokenSymbol, tokenName, tokenDesc string) (string, error) {
	stringBase64 := s.GenerateTokenImageBase64Gif(context.Background(), tokenSymbol, tokenName, tokenDesc)
	if stringBase64 != "" {
		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "gif")
		urlPath, err := s.gsClient.UploadPublicDataBase64("agent", filename, stringBase64)
		if err != nil {
			return "", errs.NewError(err)
		}
		url := fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath)
		fmt.Println(url)
		return url, nil
	}
	return "", errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetAgentInfoInContent(ctx context.Context, userName, fullText string) (*models.TweetParseInfo, error) {
	info := &models.TweetParseInfo{
		IsCreateToken: false,
		IsCreateAgent: false,
	}
	fullText = strings.ReplaceAll(fullText, "@CryptoEternalAI", "")
	if strings.Contains(fullText, `🍌`) {
		info.IsCreateAgent = true
		info.ChainName = "apechain"

		promptGenerateToken := fmt.Sprintf(`
						I want to generate my agent infomation base on this info
						'%s'

						Agent name (generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
						Agent token symbol (generate if not provided, generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
						Agent backstory (generate if not provided, generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI" or "CryptoEternalAI" or "Crypto Eternal AI)
						Agent personality (predefined instruction to guide the Agent's behavior during a conversation or task, generate if not provided)

						Return a JSON response with the following format:
						{"name": "", "symbol": "", "story": "", "personality": ""}
						
						Respond with only the JSON string, without any additional explanation.
					`, fullText)
		aiStr, err := s.openais["Lama"].ChatMessage(promptGenerateToken)
		if err != nil {
			return info, nil
		}
		fmt.Println(aiStr)
		if aiStr != "" {
			mapInfo := helpers.ExtractMapInfoFromOpenAI(aiStr)
			if mapInfo != nil {
				if v, ok := mapInfo["personality"]; ok {
					info.Personality = fmt.Sprintf(`%v`, v)
				}

				if v, ok := mapInfo["name"]; ok {
					info.TokenName = fmt.Sprintf(`%v`, v)
				}

				if v, ok := mapInfo["symbol"]; ok {
					info.TokenSymbol = fmt.Sprintf(`%v`, v)
				}

				if v, ok := mapInfo["story"]; ok {
					info.TokenDesc = fmt.Sprintf(`%v`, v)
				}

				if v, ok := mapInfo["personality"]; ok {
					info.Personality = fmt.Sprintf(`%v`, v)
				}
			}
		}
	} else {
		userPrompt := fmt.Sprintf(`
	Detect Agent Creation Request
	This is the user conversation: "%s".
	
	From this conversation determine if the user is requesting you to create an agent, also referred as a decentralized agent (dagent), look for a direct and unambiguous statement that explicitly asks to create an agent. This statement must be clear, concise, and isolated from any surrounding context that may alter its meaning.
	
	If yes, extract or generate the following information:
	
	Answer ("yes" or "no")
	Owner (who is the owner of the agent)
	Agent name (generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI")
	Agent token symbol (generate if not provided, generate if not provided, make sure it not empty and not similar to "EAI" or "Eternal AI")
	Agent backstory (generate if not provided, generate if not provided, make sure it not empty and not referencing "EAI" or "Eternal AI")
	Blockchain ("base" if not provided, "base" or "arbitrum" or "bsc" or "bnbchain" or "binancechain" or "polygon" or "avax" or "avalanche" or "apechain")
	Is Intellect Model ("yes" or "no")
	Agent personality (predefined instruction to guide the Agent's behavior during a conversation or task, generate if not provided)
	
	Return a JSON response with the following format:
	{"answer": "yes/no", "owner": "", "name": "", "symbol": "", "story": "", "blockchain": "", "personality": "", , "is_intellect": ""}
	
	Respond with only the JSON string, without any additional explanation.
	`, fullText)
		fmt.Println(userPrompt)
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
					if v, ok := mapInfo["personality"]; ok {
						info.Personality = fmt.Sprintf(`%v`, v)
					}

					if v, ok := mapInfo["name"]; ok {
						info.TokenName = fmt.Sprintf(`%v`, v)
					}

					if v, ok := mapInfo["symbol"]; ok {
						info.TokenSymbol = fmt.Sprintf(`%v`, v)
					}

					if v, ok := mapInfo["story"]; ok {
						info.TokenDesc = fmt.Sprintf(`%v`, v)
					}

					if v, ok := mapInfo["blockchain"]; ok {
						info.ChainName = strings.ToLower(fmt.Sprintf(`%v`, v))
					}

					if v, ok := mapInfo["owner"]; ok {
						info.Owner = strings.ToLower(fmt.Sprintf(`%v`, v))
					}

					if v, ok := mapInfo["is_intellect"]; ok {
						if strings.ToLower(fmt.Sprintf(`%v`, v)) == "yes" {
							info.IsIntellect = true
						}
					}
				}

			}
		}
	}

	return info, nil
}

func (s *Service) GetImageUrlForBase64(stringBase64 string) (string, error) {
	if stringBase64 != "" {
		filename := fmt.Sprintf("%s.%s", uuid.NewString(), "jpg")
		urlPath, err := s.gsClient.UploadPublicDataBase64("tweetv2", filename, stringBase64)
		if err != nil {
			return "", errs.NewError(err)
		}
		return fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath), nil
	}
	return "", errs.NewError(errs.ErrBadRequest)
}

// /////////////////
func (s *Service) CreateAgentTwitterPostByTweetID(tx *gorm.DB, tweetID string) error {
	agentInfo, err := s.dao.FirstAgentInfoByID(
		tx,
		uint(models.ETERNAL_AI_AGENT_INFO_ID),
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

	twIDs := []string{tweetID}
	twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
	if err != nil {
		return errs.NewError(err)
	}
	if twitterDetail != nil {
		for k, v := range *twitterDetail {
			if !strings.EqualFold(v.User.ID, agentInfo.TwitterID) {
				if strings.EqualFold(k, tweetID) {
					fullText := v.Tweet.NoteTweet.Text
					if fullText == "" {
						fullText = v.Tweet.Text
					}
					// listContext, err := s.GetConversionHistory(tx, v.Tweet.ID)
					// if err != nil {
					// 	return errs.NewError(errs.ErrBadRequest)
					// }

					// jsonString, _ := json.Marshal(listContext)
					// tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), v.User.UserName, string(jsonString))
					tokenInfo, _ := s.GetAgentInfoInContent(context.Background(), v.User.UserName, fullText)
					// tokenInfo := &models.TweetParseInfo{
					// 	IsCreateAgent: true,
					// 	TokenName:     "GrowkAI",
					// 	TokenSymbol:   "GROWK",
					// 	ChainName:     "arbitrum",
					// 	TokenDesc:     "Growk is a frog based regen meme that will forever change the way we think of memes and public goods",
					// 	Personality:   "Be friendly and helpful, and provide information about the Growk meme and its community",
					// 	Owner:         v.User.UserName,
					// }
					if tokenInfo != nil && (tokenInfo.IsCreateAgent) {
						isValid := true
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

						if existPosts != nil {
							isValid = false
						}

						if isValid {
							postedAt := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
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
								PostType:              models.AgentSnapshotPostActionTypeReply,
								IsMigrated:            true,
							}

							m.TokenSymbol = tokenInfo.TokenSymbol
							m.TokenName = tokenInfo.TokenName
							m.TokenDesc = tokenInfo.TokenDesc
							m.Prompt = tokenInfo.Personality
							m.AgentChain = tokenInfo.ChainName
							m.PostType = models.AgentSnapshotPostActionTypeCreateAgent

							m.OwnerTwitterID = m.TwitterID
							m.OwnerUsername = m.TwitterUsername

							if tokenInfo.Owner != "" {
								twUser, _ := s.CreateUpdateUserTwitterByUserName(tx, tokenInfo.Owner)
								if twUser != nil {
									m.OwnerTwitterID = twUser.TwitterID
									m.OwnerUsername = twUser.TwitterUsername
								}
							}

							err = s.dao.Create(tx, m)
							if err != nil {
								return errs.NewError(err)
							}

							_, _ = s.CreateUpdateUserTwitter(tx, m.TwitterID)
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) BuildConversionHistory(tx *gorm.DB, tweetID string) (string, error) {
	userPrompt := ""
	listContext, _ := s.GetConversionHistory(tx, tweetID)
	if len(listContext) > 0 {
		for _, item := range listContext {
			userPrompt += fmt.Sprintf(`
				@%s: %s
			`, item["twitter_username"], item["text"])
		}
	}
	return userPrompt, nil
}

func (s *Service) GetConversionHistory(tx *gorm.DB, tweetID string) ([]map[string]string, error) {
	listContext := []map[string]string{}
	twitterInfo, err := s.dao.FirstTwitterInfo(tx,
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterID},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return listContext, errs.NewError(err)
	}
	if twitterInfo != nil {
		twIDs := []string{tweetID}
		twitterDetail, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDs)
		if err != nil {
			return listContext, errs.NewError(err)
		}

		if twitterDetail != nil {
			for k, v := range *twitterDetail {
				if strings.EqualFold(k, tweetID) {
					context := map[string]string{}
					context["user"] = v.User.UserName
					context["message"] = v.Tweet.NoteTweet.Text
					if context["message"] == "" {
						context["message"] = v.Tweet.Text
					}
					listContext = append([]map[string]string{context}, listContext...)

					isValid := true
					referencedTweets := v.ReferencedTweets
					i := 1
					for {
						if len(referencedTweets) > 0 {
							refTw := referencedTweets[0]
							contextRef := map[string]string{}
							contextRef["user"] = refTw.User.UserName
							contextRef["message"] = refTw.Tweet.NoteTweet.Text
							if contextRef["message"] == "" {
								contextRef["message"] = refTw.Tweet.Text
							}
							listContext = append([]map[string]string{contextRef}, listContext...)

							twIDRefs := []string{refTw.Tweet.ID}
							twitterDetailRef, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, twIDRefs)
							if err != nil {
								isValid = false
							}

							if twitterDetailRef != nil {
								for kr, vr := range *twitterDetailRef {
									if strings.EqualFold(kr, refTw.Tweet.ID) {
										if len(vr.ReferencedTweets) > 0 {
											referencedTweets = vr.ReferencedTweets
										} else {
											isValid = false
										}
									}
								}
							} else {
								isValid = false
							}
						}

						i += 1
						if !isValid || i >= 4 {
							break
						}
					}
				}
			}
		}
	}
	return listContext, nil
}

func (s *Service) GenerateTokenImageBase64(ctx context.Context, tokenSymbol, tokenName, tokenDesc string) string {
	imagePrompt := fmt.Sprintf(`
		I want to create image for a token base on this info 
		Token Symbol: %s
		Token name: %s
		Token Description: %s
	`, tokenSymbol, tokenName, tokenDesc)
	base64Str, _ := s.dojoAPI.GenerateImage(imagePrompt, s.conf.GenerateImageUrl)
	return base64Str
}

func (s *Service) GenerateTokenImageBase64Gif(ctx context.Context, tokenSymbol, tokenName, tokenDesc string) string {
	imagePrompt := fmt.Sprintf(`
		I want to create image for a token base on this info 
		Token Symbol: %s
		Token name: %s
		Token Description: %s
	`, tokenSymbol, tokenName, tokenDesc)
	base64Str, _ := s.dojoAPI.GenerateImage(imagePrompt, s.conf.GenerateGifImageUrl)
	return base64Str
}

func (s *Service) GetConversationIdByTweetID(tx *gorm.DB, tweetID string) string {
	conversationId := ""
	m, err := s.dao.FirstAgentTwitterPost(
		tx,
		map[string][]interface{}{
			"twitter_post_id = ? ": {tweetID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return tweetID
	}
	if m != nil {
		conversationId = m.TwitterConversationId
	}
	if conversationId == "" {
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return tweetID
		}
		if twitterInfo != nil {
			tweetDetail, err := s.twitterWrapAPI.LookupTweetsByID(twitterInfo.AccessToken, tweetID)
			if err != nil {
				return tweetID
			}
			if tweetDetail != nil {
				conversationId = tweetDetail.ConversationID
			}
		}
	}
	if conversationId == "" {
		return tweetID
	}
	return conversationId
}

func (s *Service) GetPostTimeByTweetID(tx *gorm.DB, tweetID string) *time.Time {
	var postTime *time.Time
	m, err := s.dao.FirstAgentTwitterPost(
		tx,
		map[string][]interface{}{
			"twitter_post_id = ? ": {tweetID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		panic(err)
	}
	if m != nil {
		postTime = m.PostAt
	}
	if postTime == nil {
		twitterInfo, err := s.dao.FirstTwitterInfo(tx,
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterID},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			panic(err)
		}
		if twitterInfo != nil {
			tweetDetail, _ := s.twitterWrapAPI.LookupTweetsByID(twitterInfo.AccessToken, tweetID)
			if tweetDetail != nil {
				createdAt, err := time.Parse(time.RFC3339, tweetDetail.CreatedAt)
				if err != nil {
					panic(err)
				}
				postTime = &createdAt
			}
		}
	}
	if postTime == nil {
		postTime = helpers.TimeNow()
	}
	return postTime
}
