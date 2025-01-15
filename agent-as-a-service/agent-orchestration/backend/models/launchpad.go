package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type LaunchpadStatus string
type LaunchpadTier string

const (
	LaunchpadStatusNew     LaunchpadStatus = "new"
	LaunchpadStatusRunning LaunchpadStatus = "running"
	LaunchpadStatusEnd     LaunchpadStatus = "end"

	LaunchpadTier1 LaunchpadTier = "Tier 1"
	LaunchpadTier2 LaunchpadTier = "Tier 2"
	LaunchpadTier3 LaunchpadTier = "Tier 3"
)

type Launchpad struct {
	gorm.Model
	NetworkID              uint64
	TwitterPostID          uint   `gorm:"unique_index"`
	TweetId                string `gorm:"unique_index"`
	Name                   string
	Description            string `gorm:"type:text"`
	TwitterId              string
	TwitterUsername        string
	TwitterName            string
	Address                string
	LastScanID             string
	Status                 LaunchpadStatus
	AgentSnapshotMissionID uint
	AgentSnapshotMission   *AgentSnapshotMission
	StartAt                *time.Time
	EndAt                  *time.Time
}

type LaunchpadMember struct {
	gorm.Model
	UserAddress  string `gorm:"unique_index:lp_member_main_idx"`
	TwitterID    string `gorm:"unique_index:lp_member_main_idx"`
	LaunchpadID  uint   `gorm:"unique_index:lp_member_main_idx"`
	Launchpad    *Launchpad
	TweetID      string
	TweetContent string `gorm:"type:longtext"`
	Tier         string
}
