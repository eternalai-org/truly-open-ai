package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type AgentInfraStatus string

const (
	AgentInfraStatusNew     AgentInfraStatus = "new"
	AgentInfraStatusActived AgentInfraStatus = "actived"
)

type AgentInfra struct {
	gorm.Model
	InfraId      string `gorm:"unique_index"`
	OwnerID      uint
	Owner        *User
	OwnerAddress string
	Name         string
	Description  string `gorm:"type:text"`
	Icon         string `gorm:"type:text"`
	Docs         string `gorm:"type:longtext"`
	Status       AgentInfraStatus
	ApiUrl       string
	Price        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type AgentInfraInstall struct {
	gorm.Model
	Code         string `gorm:"unique_index"`
	AgentInfraID uint   `gorm:"index"`
	AgentInfra   *AgentInfra
	UserID       uint `gorm:"index"`
	User         *User
}

type AgentInfraLog struct {
	gorm.Model
	AgentInfraInstallID uint
	UserID              uint
	AgentInfraID        uint
	Price               numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	UrlPath             string           `gorm:"type:text"`
	Status              int
}
