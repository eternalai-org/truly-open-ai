package telegram

import (
	"context"
	"fmt"

	"agent-battle/pkg/logger"
	"agent-battle/pkg/utils"

	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

const DefaultTelegramChannel = -0

var MapSlackHookToTelegramChannelId = map[string]int64{}

type telegramClient struct {
	teleBot *tgbotapi.Bot
}

type Attachment struct {
	AuthorIcon string  `json:"author_icon"`
	AuthorLink string  `json:"author_link"`
	AuthorName string  `json:"author_name"`
	Color      string  `json:"color"`
	Fallback   string  `json:"fallback"`
	Fields     []Field `json:"fields"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	ImageURL   string  `json:"image_url"`
	Pretext    string  `json:"pretext"`
	Text       string  `json:"text"`
	ThumbURL   string  `json:"thumb_url"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Ts         int     `json:"ts"`
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Message struct {
	Attachments          []Attachment `json:"attachments"`
	ChannelId            int64        `json:"channel_id"`
	Text                 string                      `json:"text"`
	FileUploadParameters *slack.FileUploadParameters `json:"file_upload_parameters,omitempty"`
}

type ITelegramClient interface {
	SendMessageToTelegramChannel(ctx context.Context, payload *Message) error
}

func (client *telegramClient) SendMessageToTelegramChannel(ctx context.Context, payload *Message) error {
	if client.teleBot == nil {
		return nil
	}
	buildText := payload.Text
	if len(payload.Attachments) > 0 {
		buildText = ""
		att := payload.Attachments[0]
		if att.Title != "" {
			buildText = fmt.Sprintf("<strong>%s</strong>", att.Title)
		}
		buildText = buildText + "\n<blockquote>"
		for _, v := range att.Fields {
			if v.Title != "" {
				title := fmt.Sprintf("<b>* %s</b>", v.Title)
				buildText = buildText + "\n" + title + ": " + v.Value
			} else {
				buildText = buildText + "\n" + v.Value
			}
		}
		buildText = buildText + "</blockquote>"
	}

	if payload.FileUploadParameters != nil {
		_, err := client.teleBot.SendDocument(ctx, &tgbotapi.SendDocumentParams{
			ChatID: getTelegramChannelId(payload.ChannelId),
			Document: &models.InputFileUpload{
				Filename: payload.FileUploadParameters.Filename,
				Data:     payload.FileUploadParameters.Reader,
			},
		})
		if err != nil {
			return err
		}
	} else {
		_, err := client.teleBot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:    getTelegramChannelId(payload.ChannelId),
			Text:      buildText,
			ParseMode: "HTML",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getTelegramChannelId(channelId int64) int64 {
	if utils.IsEnvProduction() {
		return channelId
	}

	return DefaultTelegramChannel
}

func New() ITelegramClient {
	teleBot, err := tgbotapi.New(viper.GetString("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		logger.AtLog.Fatal(err)
	}
	return &telegramClient{teleBot: teleBot}
}
