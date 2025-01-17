package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type LaunchpadStatus string
type LaunchpadTier string

const (
	LaunchpadStatusNew          LaunchpadStatus = "new"
	LaunchpadStatusRunning      LaunchpadStatus = "running"
	LaunchpadStatusEnd          LaunchpadStatus = "end"
	LaunchpadStatusFailed       LaunchpadStatus = "failed"
	LaunchpadStatusDone         LaunchpadStatus = "done"
	LaunchpadStatusCancelled    LaunchpadStatus = "cancelled"
	LaunchpadStatusTokenError   LaunchpadStatus = "token_error"
	LaunchpadStatusTokenCreated LaunchpadStatus = "token_created"
	LaunchpadStatusSettleError  LaunchpadStatus = "settle_error"
	LaunchpadStatusSettled      LaunchpadStatus = "settled"
	LaunchpadStatusTge          LaunchpadStatus = "tge"
	LaunchpadStatusPoolError    LaunchpadStatus = "pool_error"

	LaunchpadTier1 LaunchpadTier = "Tier 1"
	LaunchpadTier2 LaunchpadTier = "Tier 2"
	LaunchpadTier3 LaunchpadTier = "Tier 3"
)

type Launchpad struct {
	gorm.Model
	NetworkID              uint64
	TwitterPostID          uint   `gorm:"unique_index"`
	TweetId                string `gorm:"unique_index"`
	ReplyTweetId           string
	Name                   string
	Description            string `gorm:"type:text"`
	TwitterId              string
	TwitterUsername        string
	TwitterName            string
	Address                string
	LastScanID             string
	Status                 LaunchpadStatus `gorm:"index"`
	AgentSnapshotMissionID uint
	AgentSnapshotMission   *AgentSnapshotMission
	StartAt                *time.Time
	EndAt                  *time.Time `gorm:"index"`
	FinishedAt             *time.Time
	FundBalance            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalBalance           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SettleFundTxHash       string
	CancelFundTxHash       string
	AddLiquidityTxHash     string
	TokenAddress           string `gorm:"index"`
	TokenName              string
	TokenSymbol            string
	TokenImageUrl          string
	TotalSupply            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	DeployTokenTxHash      string
	TgeBalance             numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	MaxFundBalance         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RefundBalance          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type LaunchpadMemberStatus string

const (
	LaunchpadMemberStatusNew         LaunchpadMemberStatus = "new"
	LaunchpadMemberStatusTgeDone     LaunchpadMemberStatus = "tge_done"
	LaunchpadMemberStatusTgeError    LaunchpadMemberStatus = "tge_error"
	LaunchpadMemberStatusDone        LaunchpadMemberStatus = "done"
	LaunchpadMemberStatusRefundError LaunchpadMemberStatus = "refund_error"
)

type LaunchpadMember struct {
	gorm.Model
	NetworkID            uint64
	UserAddress          string `gorm:"unique_index:lp_member_main_idx"`
	TwitterID            string `gorm:"unique_index:lp_member_main_idx"`
	LaunchpadID          uint   `gorm:"unique_index:lp_member_main_idx"`
	Launchpad            *Launchpad
	TweetID              string
	TweetContent         string `gorm:"type:longtext"`
	Tier                 LaunchpadTier
	ReplyContent         string `gorm:"type:longtext"`
	ReplyPostAt          *time.Time
	ReplyPostID          string
	Error                string           `gorm:"type:longtext"`
	FundBalance          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	MaxFundBalance       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalBalance         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RefundBalance        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RefundFeeBalance     numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenBalance         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenTransferTxHash  string
	RefundTransferTxHash string
	Status               LaunchpadMemberStatus `gorm:"default:'new'"`
}

type LaunchpadTransactionType string
type LaunchpadTransactionStatus string

const (
	LaunchpadTransactionTypeDeposit LaunchpadTransactionType = "deposit"
	LaunchpadTransactionTypeRefund  LaunchpadTransactionType = "refund"

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
	Fee         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status      LaunchpadTransactionStatus
}
