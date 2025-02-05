package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type AgentStore struct {
	gorm.Model
	Name               string
	Description        string `gorm:"type:text"`
	OwnerAddress       string
	AuthenUrl          string `gorm:"type:longtext"`
	AgentStoreMissions []*AgentStoreMission
}

type AgentStoreMission struct {
	gorm.Model
	AgentStoreID uint `gorm:"index"`
	Name         string
	Description  string           `gorm:"type:text"`
	UserPrompt   string           `gorm:"type:longtext"`
	Price        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	OwnerAddress string
	ToolList     string  `gorm:"type:longtext"`
	Rating       float64 `gorm:"type:decimal(5,2);default:0"`
	NumRating    uint
	NumUsed      uint
	Icon         string `gorm:"type:text"`
}

type AgentStoreInstall struct {
	gorm.Model
	AgentStoreID   uint `gorm:"index"`
	AgentStore     *AgentStore
	AgentInfoID    uint   `gorm:"index"`
	CallbackParams string `gorm:"type:longtext"` //{"user_id" : "123", "authen_token" : "xxx",...}
}
