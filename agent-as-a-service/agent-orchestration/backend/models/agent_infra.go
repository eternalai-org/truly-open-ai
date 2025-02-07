package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type AgentInfra struct {
	gorm.Model
	OwnerAddress string
	Name         string
	Description  string           `gorm:"type:text"`
	EaiBalance   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Icon         string           `gorm:"type:text"`
}

type (
	AgentInfraTransactionType   string
	AgentInfraTransactionStatus string
)

const (
	AgentInfraTransactionTypeFee AgentInfraTransactionType = "fee"

	AgentInfraTransactionStatusDone AgentInfraTransactionStatus = "done"
)

type AgentInfraTransaction struct {
	gorm.Model
	NetworkID    uint64
	AgentInfraID uint   `gorm:"index"`
	EventId      string `gorm:"unique_index"`
	Type         AgentInfraTransactionType
	Amount       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Toolset      string
	Status       AgentInfraTransactionStatus
}
