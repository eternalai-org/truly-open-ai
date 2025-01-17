package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/mymmrac/telego"
)

func (s *Service) RunTeleBotJob(ctx context.Context) error {
	err := s.JobEnabledDB(ctx)
	if err != nil {
		panic(err)
	}

	bot, err := telego.NewBot(s.conf.Telebot.TradeAnalytics.Botkey, telego.WithDefaultDebugLogger())
	if err != nil {
		return errs.NewError(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			m, err := s.dao.FirstAgentTeleMsg(
				daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"message_id = ?": {update.Message.MessageID},
				},
				map[string][]interface{}{},
				[]string{},
			)
			if err != nil {
				return errs.NewError(err)
			}
			if m == nil {
				agentSnapshotMission, err := s.dao.FirstAgentSnapshotMission(
					daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						`tele_chat_id = ?`: {update.Message.Chat.ID},
					},
					map[string][]interface{}{},
					[]string{},
				)
				if err != nil {
					return errs.NewError(err)
				}

				if agentSnapshotMission != nil {
					msgDate := time.Unix(update.Message.Date, 0)
					m = &models.AgentTeleMsg{
						MessageID:              fmt.Sprintf(`%d`, update.Message.MessageID),
						MessageDate:            &msgDate,
						Content:                update.Message.Text,
						ChatID:                 fmt.Sprintf(`%d`, update.Message.Chat.ID),
						ChatUsername:           update.Message.Chat.Username,
						AgentInfoID:            agentSnapshotMission.AgentInfoID,
						AgentSnapshotMissionID: agentSnapshotMission.ID,
						Status:                 models.TeleMsgStatusNew,
					}
					err = s.dao.Create(daos.GetDBMainCtx(ctx), m)
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) SendTeleMsgToChatID(ctx context.Context, content, chatID string) (string, error) {
	bot, err := telego.NewBot(s.conf.Telebot.TradeAnalytics.Botkey, telego.WithDefaultDebugLogger())
	if err != nil {
		return "", errs.NewError(err)
	}

	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return "", errs.NewError(err)
	}

	resp, err := bot.SendMessage(
		&telego.SendMessageParams{
			ChatID: telego.ChatID{
				ID: i,
			},
			Text: strings.TrimSpace(content),
		},
	)
	if err != nil {
		return "", errs.NewError(err)
	}

	if resp != nil {
		return fmt.Sprintf(`%d`, resp.MessageID), nil
	}
	return "", nil
}

func (s *Service) SendTeleMsgToKBChannel(ctx context.Context, content, chatID string) (string, error) {
	bot, err := telego.NewBot(s.conf.KnowledgeBaseConfig.KBTelegramKey, telego.WithDefaultDebugLogger())
	if err != nil {
		return "", errs.NewError(err)
	}
	defer bot.Close()
	i, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return "", errs.NewError(err)
	}

	resp, err := bot.SendMessage(
		&telego.SendMessageParams{
			ChatID: telego.ChatID{
				ID: i,
			},
			Text: strings.TrimSpace(content),
		},
	)
	if err != nil {
		return "", errs.NewError(err)
	}

	if resp != nil {
		return fmt.Sprintf(`%d`, resp.MessageID), nil
	}
	return "", nil
}
