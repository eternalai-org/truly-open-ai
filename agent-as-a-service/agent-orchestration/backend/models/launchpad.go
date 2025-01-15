package models

import (
	"github.com/jinzhu/gorm"
)

type LaunchpadStatus string

const (
	LaunchpadStatusNew LaunchpadStatus = "new"
	LaunchpadStatusEnd LaunchpadStatus = "end"
)

type Launchpad struct {
	gorm.Model
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
}

type LaunchpadMember struct {
	gorm.Model
	UserAddress  string
	LaunchpadID  uint
	Launchpad    *Launchpad
	TweetID      string
	TweetContent string `gorm:"type:longtext"`
}
