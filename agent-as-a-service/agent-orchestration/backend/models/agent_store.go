package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type (
	AgentStoreType   string
	AgentStoreStatus string

	AgentStoreInstallStatus string
	AgentStoreInstallType   string
)

const (
	AgentStoreInstallStatusNew  AgentStoreInstallStatus = "new"
	AgentStoreInstallStatusDone AgentStoreInstallStatus = "done"
	AgentStoreInstallTypeAgent  AgentStoreInstallType   = "agent"
	AgentStoreInstallTypeUser   AgentStoreInstallType   = "user"

	AgentStoreTypeStore AgentStoreType = "store"
	AgentStoreTypeInfra AgentStoreType = "infra"

	AgentStoreStatusNew     AgentStoreStatus = "new"
	AgentStoreStatusActived AgentStoreStatus = "actived"
)

type AgentStore struct {
	gorm.Model
	StoreId            string `gorm:"unique_index"`
	Type               AgentStoreType
	Name               string
	Description        string `gorm:"type:text"`
	OwnerAddress       string
	OwnerID            uint
	Owner              *User
	AuthenUrl          string `gorm:"type:longtext"`
	Docs               string `gorm:"type:longtext"`
	ApiUrl             string `gorm:"type:longtext"`
	Icon               string `gorm:"type:text"`
	Status             AgentStoreStatus
	Price              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	AgentStoreMissions []*AgentStoreMission
}

type AgentStoreMission struct {
	gorm.Model
	AgentStoreID uint `gorm:"index"`
	Name         string
	Description  string           `gorm:"type:text"`
	UserPrompt   string           `gorm:"type:longtext"`
	Price        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	OwnerAddress string
	ToolList     string  `gorm:"type:longtext"`
	Rating       float64 `gorm:"type:decimal(5,2);default:0"`
	NumRating    uint
	NumUsed      uint
	Icon         string `gorm:"type:text"`
}

type AgentStoreInstall struct {
	gorm.Model
	Code           string `gorm:"unique_index"`
	UserID         uint   `gorm:"index"`
	User           *User
	AgentStoreID   uint `gorm:"index"`
	AgentInfoID    uint `gorm:"index"`
	AgentStore     *AgentStore
	CallbackParams string `gorm:"type:longtext"` //{"user_id" : "123", "authen_token" : "xxx",...}
	Status         AgentStoreInstallStatus
	Type           AgentStoreInstallType
}

type (
	AgentStoreTransactionType   string
	AgentStoreTransactionStatus string
)

const (
	AgentStoreTransactionTypeFee AgentStoreTransactionType = "fee"

	AgentStoreTransactionStatusDone AgentStoreTransactionStatus = "done"
)

type AgentStoreTransaction struct {
	gorm.Model
	NetworkID    uint64
	AgentStoreID uint   `gorm:"index"`
	EventId      string `gorm:"unique_index"`
	Type         AgentStoreTransactionType
	Amount       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Toolset      string
	Status       AgentStoreTransactionStatus
}

type AgentStoreLog struct {
	gorm.Model
	AgentStoreInstallID uint
	UserID              uint
	AgentStoreID        uint
	Price               numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	UrlPath             string           `gorm:"type:text"`
	Status              int
}
