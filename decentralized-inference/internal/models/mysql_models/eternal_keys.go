package mysql_models

import (
	"gorm.io/gorm"
)

type EternalKeys struct {
	gorm.Model

	IdentityId        string  `json:"identity_id"`
	TwitterName       string  `json:"twitter_name" `
	TwitterAvatar     string  `json:"twitter_avatar" `
	TwitterUsername   string  `json:"twitter_username"`
	TokenAddress      string  `json:"token_address" gorm:"token_address"`
	AssistantID       string  `json:"assistant_id" gorm:"assistant_id"`
	AssistantOpenAiID string  `json:"assistant_open_ai_id" gorm:"assistant_open_ai_id"`
	ShareRequired     float64 `json:"share_required" gorm:"share_required"`
	Supply            float64 `json:"supply" gorm:"supply"`
}

func (EternalKeys) TableName() string {
	return "eternal_keys"
}
