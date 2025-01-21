package mysql_models

import (
	"gorm.io/gorm"
)

type AgentSnapshotMissions struct {
	gorm.Model
	NetworkID   uint64
	AgentInfoID uint `gorm:"index"`
	//AgentInfo    *AgentInfo
	UserPrompt   string `gorm:"type:longtext"`
	IntervalSec  int    `gorm:"default:0"`
	Enabled      bool   `gorm:"default:0"`
	ReplyEnabled bool   `gorm:"default:0"`
	IsTesting    bool   `gorm:"default:0"`
	ToolSet      string
}
