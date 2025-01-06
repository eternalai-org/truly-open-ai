package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/mymmrac/telego"
)

func (s *Service) JobAgentTeleAlertTopupNotAction(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobAgentTeleAlertTopupNotAction",
		func() error {
			var retErr error
			{
				topups, err := s.dao.FindAgentEaiTopup(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"agent_eai_topups.created_at >= adddate(now(), interval -4 hour)": {},
						"agent_eai_topups.created_at < adddate(now(), interval -2 hour)":  {},
						`not exists(
							select 1
							from agent_snapshot_post_actions
							where agent_snapshot_post_actions.agent_info_id = agent_eai_topups.agent_info_id
							and agent_snapshot_post_actions.status = 'done'
						)`: {},
					},
					map[string][]interface{}{},
					[]string{},
					0,
					999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, topup := range topups {
					err = s.AgentTeleAlertTopupNotActionByID(ctx, topup.AgentInfoID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
				}
			}
			{
				agents, err := s.dao.FindAgentInfoJoin(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"join twitter_infos on twitter_infos.id = agent_infos.twitter_info_id": {},
					},
					map[string][]interface{}{
						"agent_infos.eai_balance >= agent_infos.agent_fee":             {},
						"twitter_infos.created_at >= adddate(now(), interval -4 hour)": {},
						"twitter_infos.created_at < adddate(now(), interval -2 hour)":  {},
						`not exists(
							select 1
							from agent_snapshot_post_actions
							where agent_snapshot_post_actions.agent_info_id = agent_infos.id
							and agent_snapshot_post_actions.status = 'done'
						)`: {},
					},
					map[string][]interface{}{},
					[]string{},
					0,
					999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, agent := range agents {
					err = s.AgentTeleAlertTopupNotActionByID(ctx, agent.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
					}
				}
			}
			{
				twitterInfos, err := s.dao.FindTwitterInfo(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"created_at >= adddate(now(), interval -3 hour)": {},
						`created_at < adddate(now(), interval -1 hour)
						or exists(
								select 1
								from agent_infos
								where agent_infos.twitter_info_id = twitter_infos.id
								and agent_infos.token_address != ''
								and agent_infos.token_supply != ''
							)`: {},
					},
					map[string][]interface{}{},
					[]string{},
					0,
					999999,
				)
				if err != nil {
					return errs.NewError(err)
				}
				for _, twitterInfo := range twitterInfos {
					err = s.AgentTeleAlertNewAgentTrakerByID(ctx, twitterInfo.ID)
					if err != nil {
						retErr = errs.MergeError(retErr, errs.NewError(err))
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

func (s *Service) AgentTeleAlertTopupNotActionByID(ctx context.Context, agentInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTeleAlertTopupNotActionByID_%d", agentInfoID),
		func() error {
			var rs bool
			err := s.RedisCached(
				fmt.Sprintf("AgentTeleAlertTopupNotActionByID_%d", agentInfoID),
				true,
				6*time.Hour,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						agent, err := s.dao.FirstAgentInfoByID(
							daos.GetDBMainCtx(ctx),
							agentInfoID,
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agent != nil {
							bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
							if err != nil {
								return errs.NewError(err)
							}
							post, err := s.dao.FirstAgentSnapshotPost(
								daos.GetDBMainCtx(ctx),
								map[string][]interface{}{
									"agent_info_id = ?": {agent.ID},
								},
								map[string][]interface{}{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							balanceStr := fmt.Sprintf(
								`account balance is %s EAI`,
								numeric.BigFloat2Text(&agent.EaiBalance.Float),
							)
							triggerResult := "trigger interval is not running"
							if post != nil {
								triggerResult = "trigger interval is running"
							}
							action, err := s.dao.FirstAgentSnapshotPostAction(
								daos.GetDBMainCtx(ctx),
								map[string][]interface{}{
									"agent_info_id = ?": {agent.ID},
								},
								map[string][]interface{}{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							actionResult := "trigger result is not received"
							if action != nil {
								actionResult = "trigger result is received"
							}
							title := "💀💀💀 Not Action Alert! 💀💀💀"
							msg := fmt.Sprintf(
								`%s (@%s - %s - %s) deposited money but haven't seen any action yet (note: %s, %s, %s)`,
								agent.AgentName,
								agent.TwitterUsername,
								agent.NetworkName,
								agent.AgentContractID,
								balanceStr,
								triggerResult,
								actionResult,
							)
							_, err = bot.SendMessage(
								&telego.SendMessageParams{
									ChatID: telego.ChatID{
										ID: s.conf.Telebot.Alert.ChatID,
									},
									MessageThreadID: s.conf.Telebot.Alert.MessageThreadID,
									Text: strings.TrimSpace(
										fmt.Sprintf(
											`
%s

%s

Hey, @zoro_521, let's check about this!
				`,
											title,
											msg,
										),
									),
								},
							)
							if err != nil {
								return errs.NewError(err)
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
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentTeleAlertNewAgentTrakerByID(ctx context.Context, twitterInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTeleAlertNewAgentTrakerByID_%d", twitterInfoID),
		func() error {
			var rs bool
			err := s.RedisCached(
				fmt.Sprintf("AgentTeleAlertNewAgentTrakerByID_%d", twitterInfoID),
				true,
				6*time.Hour,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						twitterInfo, err := s.dao.FirstTwitterInfoByID(
							daos.GetDBMainCtx(ctx),
							twitterInfoID,
							map[string][]interface{}{},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if twitterInfo != nil {
							agent, err := s.dao.FirstAgentInfo(
								daos.GetDBMainCtx(ctx),
								map[string][]interface{}{
									"twitter_info_id = ?": {twitterInfo.ID},
								},
								map[string][]interface{}{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if agent != nil {
								fwId := agent.AgentID
								if agent.TokenAddress != "" && agent.TokenSymbol != "" {
									fwId = agent.TokenAddress
								}
								bot, err := telego.NewBot(s.conf.Telebot.Tracker.Botkey, telego.WithDefaultDebugLogger())
								if err != nil {
									return errs.NewError(err)
								}

								_, err = bot.SendMessage(
									&telego.SendMessageParams{
										ChatID: telego.ChatID{
											ID: s.conf.Telebot.Tracker.ChatID,
										},
										MessageThreadID: s.conf.Telebot.Tracker.MessageThreadID,
										Text: strings.TrimSpace(
											fmt.Sprintf(
												`
Welcome %s to the world!

X: https://twitter.com/%s
Network: %s

https://eternalai.org/agent/%s
					`,
												twitterInfo.TwitterName,
												twitterInfo.TwitterUsername,
												agent.NetworkName,
												fwId,
											),
										),
									},
								)
								if err != nil {
									return errs.NewError(err)
								}
							} else {
								return errs.NewError(errs.ErrBadRequest)
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
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentTeleAlertByID(ctx context.Context, agentInfoID uint, refID string, amount *big.Float, networkID uint64) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTeleAlertByID_%d_%s", agentInfoID, refID),
		func() error {
			var rs bool
			err := s.RedisCached(
				fmt.Sprintf("AgentTeleAlertByID_%d_%s", agentInfoID, refID),
				true,
				1*time.Hour,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						agent, err := s.dao.FirstAgentInfoByID(
							daos.GetDBMainCtx(ctx),
							agentInfoID,
							map[string][]interface{}{
								"TwitterInfo": {},
							},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agent.TwitterInfo != nil {
							bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
							if err != nil {
								return errs.NewError(err)
							}
							title := "💰💰💰 New Topped Up Alert! 💰💰💰"
							msg := fmt.Sprintf("just topped up %s EAI tokens on %s! 👀", numeric.BigFloat2Text(amount), models.GetChainName(networkID))
							if numeric.BigFloat2Text(amount) == "0" {
								title = "📝📝📝 New Agent Alert! 📝📝📝"
								msg = fmt.Sprintf("just registered on %s! 👀", agent.NetworkName)
							}
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
											agent.TwitterInfo.TwitterName,
											agent.TwitterInfo.TwitterUsername,
											msg,
										),
									),
								},
							)
							if err != nil {
								return errs.NewError(err)
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
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentTeleAlertNewTokenByID(ctx context.Context, agentInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTeleAlertNewTokenByID_%d", agentInfoID),
		func() error {
			var rs bool
			err := s.RedisCached(
				fmt.Sprintf("AgentTeleAlertNewTokenByID_%d", agentInfoID),
				true,
				1*time.Hour,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						agent, err := s.dao.FirstAgentInfoByID(
							daos.GetDBMainCtx(ctx),
							agentInfoID,
							map[string][]interface{}{
								"TwitterInfo": {},
							},
							false,
						)
						if err != nil {
							return errs.NewError(err)
						}
						if agent.TwitterInfo != nil && agent.TokenSymbol != "" && agent.TokenAddress != "" && agent.TokenAddress != "pending" {
							bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
							if err != nil {
								return errs.NewError(err)
							}
							title := "🟡🟡🟡 New Pumpfun Token Alert! 🟡🟡🟡"
							msg := fmt.Sprintf(`just created token %s https://pump.fun/coin/%s 👀`, agent.TokenSymbol, agent.TokenAddress)
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
											agent.TwitterInfo.TwitterName,
											agent.TwitterInfo.TwitterUsername,
											msg,
										),
									),
								},
							)
							if err != nil {
								return errs.NewError(err)
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
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) AgentDailyReport(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"AgentDailyReport",
		func() error {
			var rs bool
			err := s.RedisCached(
				"AgentDailyReport",
				true,
				1*time.Hour,
				&rs,
				func() (interface{}, error) {
					err := func() error {
						msg, err := s.dao.AgentInfoGetReportDaily(daos.GetDBMainCtx(ctx))
						if err != nil {
							return errs.NewError(err)
						}
						msg = fmt.Sprintf("🚨 Report!\n\n%s\n\nHey, @tygenz", msg)
						bot, err := telego.NewBot(s.conf.Telebot.Alert.Botkey, telego.WithDefaultDebugLogger())
						if err != nil {
							return errs.NewError(err)
						}
						_, err = bot.SendMessage(
							&telego.SendMessageParams{
								ChatID: telego.ChatID{
									ID: s.conf.Telebot.Alert.ChatID,
								},
								Text: msg,
							},
						)
						if err != nil {
							return errs.NewError(err)
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
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
