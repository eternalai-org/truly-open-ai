package models

import (
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type AgentWallet struct {
	gorm.Model
	NetworkID   uint64 `gorm:"unique_index:agent_wallets_main_idx"`
	AgentInfoID uint   `gorm:"unique_index:agent_wallets_main_idx"`
	AgentInfo   *AgentInfo
	Address     string
	CdpWalletID string
}

type AgentWalletActionStatus string

const (
	AgentWalletActionStatusNew   AgentWalletActionStatus = "new"
	AgentWalletActionStatusDone  AgentWalletActionStatus = "done"
	AgentWalletActionStatusError AgentWalletActionStatus = "error"
)

type AgentWalletAction struct {
	gorm.Model
	NetworkID              uint64
	AgentInfoID            uint `gorm:"index"`
	AgentWalletID          uint `gorm:"index"`
	ActionType             string
	ActionInput            string `gorm:"type:longtext"`
	ActionOutput           string `gorm:"type:longtext"`
	Status                 AgentWalletActionStatus
	RefID                  string
	Toolset                string
	AgentSnapshotMissionID uint `gorm:"index"`
	AgentSnapshotMission   *AgentSnapshotMission
	AgentSnapshotPostID    uint `gorm:"index"`
	AgentSnapshotPost      *AgentSnapshotPost
	Side                   string
	Mint                   string
	AmountIn               numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	AmountOut              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TxHash                 string
}

type AgentWalletTrade struct {
	gorm.Model
	EventId                string `gorm:"unique_index"`
	NetworkID              uint64
	AgentInfoID            uint `gorm:"index"`
	AgentWalletID          uint `gorm:"index"`
	TradedAt               *time.Time
	TokenIn                string
	TokenOut               string
	AmountIn               numeric.BigInt
	AmountOut              numeric.BigInt
	TxHash                 string
	RefID                  string
	AgentSnapshotMissionID uint `gorm:"index"`
	AgentSnapshotMission   *AgentSnapshotMission
	AgentSnapshotPostID    uint `gorm:"index"`
	AgentSnapshotPost      *AgentSnapshotPost
}

type AgentWalletTradeSum struct {
	Mint       string
	BuyAmount  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SellAmount numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}
