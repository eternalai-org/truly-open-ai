package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type LaunchpadStatus string
type LaunchpadTier string

const (
	LaunchpadStatusNew       LaunchpadStatus = "new"
	LaunchpadStatusRunning   LaunchpadStatus = "running"
	LaunchpadStatusEnd       LaunchpadStatus = "end"
	LaunchpadStatusDone      LaunchpadStatus = "done"
	LaunchpadStatusCancelled LaunchpadStatus = "cancelled"

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
	FundBalance            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalBalance           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SettleFundTxHash       string
	CancelFundTxHash       string
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
	FundBalance  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalBalance numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type LaunchpadTransactionType string
type LaunchpadTransactionStatus string

const (
	LaunchpadTransactionTypeDeposit LaunchpadTransactionType = "deposit"

	LaunchpadTransactionStatusDone LaunchpadTransactionStatus = "done"
)

type LaunchpadTransaction struct {
	gorm.Model
	NetworkID   uint64
	EventId     string `gorm:"unique_index"`
	TxHash      string
	Type        LaunchpadTransactionType
	LaunchpadID uint
	UserAddress string
	Amount      numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status      LaunchpadTransactionStatus
}
