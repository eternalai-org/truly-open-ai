package mysql_models

import "gorm.io/gorm"

type AgentInfo struct {
	gorm.Model
	TwitterUsername   string
	FarcasterID       string `gorm:"index"`
	FarcasterUsername string `gorm:"index"`
	AgentName         string
	TokenName         string
	TokenSymbol       string
	TokenAddress      string
}

func (AgentInfo) TableName() string {
	return "agent_infos"
}
