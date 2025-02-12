package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	NetworkID       uint64           `gorm:"unique_index:user_main_uidx"`
	Address         string           `gorm:"unique_index:user_main_uidx"`
	Username        string           `gorm:"unique_index"`
	SubscriptionNum uint             `gorm:"default:0"`
	Price30d        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Price30dUpdated uint64           `gorm:"default:0"`
	Price90d        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Price90dUpdated uint64           `gorm:"default:0"`
	Description     string           `gorm:"type:text"`
	Social          string           `gorm:"type:text"`
	ImageURL        string           `gorm:"type:text"`
	TwitterID       string           `gorm:"index"`
	TwitterAvatar   string
	TwitterUsername string
	TwitterName     string
	EthAddress      string           `gorm:"index"`
	TronAddress     string           `gorm:"index"`
	SolAddress      string           `gorm:"index"`
	EaiBalance      numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`

	Mentions  int64
	Likes     int64
	Followers int16
	Following int16

	Subscribed   bool             `gorm:"-"`
	TotalLike    uint             `gorm:"-"`
	TotalPost    uint             `gorm:"-"`
	TotalMessage uint             `gorm:"-"`
	TipPayment   numeric.BigFloat `gorm:"-"`
	TipReceive   numeric.BigFloat `gorm:"-"`
}

type TokenHolder struct {
	gorm.Model
	NetworkID       uint64 `gorm:"unique_index:token_holder_main_uidx"`
	ContractAddress string `gorm:"unique_index:token_holder_main_uidx"`
	Address         string `gorm:"unique_index:token_holder_main_uidx"`
	Balance         string `gorm:"default:0"`
	LastBlockNumber uint

	UserName            string           `gorm:"-"`
	ImageURL            string           `gorm:"-"`
	MemeName            string           `gorm:"-"`
	MemeTicker          string           `gorm:"-"`
	MemeImage           string           `gorm:"-"`
	MemePrice           numeric.BigFloat `gorm:"-"`
	MemePriceUsd        numeric.BigFloat `gorm:"-"`
	MemeBaseTokenSymbol string           `gorm:"-"`
}

type (
	UserTransactionType   string
	UserTransactionStatus string
)

const (
	UserTransactionTypeDeposit              UserTransactionType = "deposit"
	UserTransactionTypeAgentStoreFee        UserTransactionType = "agent_store_fee"
	UserTransactionTypeTriggerFee           UserTransactionType = "trigger_fee"
	UserTransactionTypeTriggerRefundFee     UserTransactionType = "trigger_refund_fee"
	UserTransactionTypeUserAgentInfraFee    UserTransactionType = "user_agent_infra_fee"
	UserTransactionTypeCreatorAgentInfraFee UserTransactionType = "creator_agent_infra_fee"

	UserTransactionStatusDone      UserTransactionStatus = "done"
	UserTransactionStatusCancelled UserTransactionStatus = "cancelled"
)

type UserTransaction struct {
	gorm.Model
	NetworkID   uint64
	UserID      uint `gorm:"index"`
	User        *User
	EventId     string `gorm:"unique_index"`
	Type        UserTransactionType
	FromAddress string
	ToAddress   string
	TxHash      string
	Amount      numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status      UserTransactionStatus
	Error       string `gorm:"type:longtext"`
}
