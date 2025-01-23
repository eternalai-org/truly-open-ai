package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type (
	DepositStatus string
	PackageType   string
)

const (
	DepositStatusPending DepositStatus = "pending"
	DepositStatusDone    DepositStatus = "done"

	PackageTypeFree  PackageType = "free"
	PackageTypeBasic PackageType = "basic"
	PackageTypePro   PackageType = "pro"
)

type ApiSubscriptionPackage struct {
	gorm.Model
	Name        string
	Description string           `gorm:"type:text"`
	Price       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	NetworkID   uint64
	NumToken    uint64
	Type        PackageType
	DurationDay uint
}

type ApiSubscriptionKey struct {
	gorm.Model
	NetworkID      uint64 `gorm:"unique_index:api_sub_key_main_uidx"`
	UserAddress    string `gorm:"unique_index:api_sub_key_main_uidx"`
	TwitterID      string `gorm:"unique_index:api_sub_key_main_uidx"`
	TwitterInfoID  uint
	TwitterInfo    *TwitterInfo
	ApiKey         string `gorm:"unique_index"`
	PackageID      uint
	Package        ApiSubscriptionPackage
	QuotaRemaining uint64
	StartedAt      *time.Time
	ExpiresAt      *time.Time
	DepositAddress string `gorm:"index"`
}

type ApiSubscriptionHistory struct {
	gorm.Model
	NetworkID      uint64
	UserAddress    string
	ApiKey         string
	PackageID      uint
	Package        ApiSubscriptionPackage
	DepositAddress string `gorm:"index"`
	DepositStatus  DepositStatus
	TxHash         string
	EventId        string `gorm:"unique_index"`
	NumToken       uint64
	StartedAt      *time.Time
	ExpiresAt      *time.Time
}

type ApiSubscriptionUsageLog struct {
	gorm.Model
	ApiKey   string
	Endpoint string
	NumToken uint64
}
