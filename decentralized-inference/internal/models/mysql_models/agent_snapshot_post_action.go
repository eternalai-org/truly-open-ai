package mysql_models

import (
	"gorm.io/gorm"
	"time"
)

const (
	AGENT_SNAPSHOT_STATUS_DONE = "done"
)

type AgentSnapshotPostAction struct {
	gorm.Model
	NetworkID              uint64
	AgentInfoID            uint `gorm:"index"`
	AgentSnapshotMissionID uint `gorm:"index"`
	AgentSnapshotPostID    uint `gorm:"index"`
	AgentTwitterId         string
	Type                   string
	TargetUsername         string
	TargetTwitterId        string `gorm:"index"`
	Tweetid                string `gorm:"index"`
	Content                string `gorm:"type:longtext"`
	RefId                  string `gorm:"type:longtext"`
	RefIds                 string `gorm:"type:longtext"`
	Error                  string
	FollowerCount          uint       `gorm:"default:0"`
	IsApproved             bool       `gorm:"default:0"`
	ScheduleAt             *time.Time `gorm:"index"`
	ExecutedAt             *time.Time
	IsMigrated             bool `gorm:"default:0"`
	TokenName              string
	TokenSymbol            string
	TokenAddress           string
	TokenHash              string
	TokenInferID           string
	TokenImageUrl          string
	TokenTweet             string
}
