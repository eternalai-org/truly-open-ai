package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

var KeyConfigNameWalletDeploy = "wallet_deploy"
var KeyConfigNameModelId = "knowledge_base_model_id"
var KeyConfigNameKnowledgeBaseWorkerHubAddress = "know_ledge_base_worker_hub_address"
var KeyConfigNameKnowledgeBaseTokenContractAddress = "knowledge_base_token_contract_address"

type AppConfig struct {
	gorm.Model
	NetworkID uint64 `gorm:"unique_index:app_configs_main_uidx"`
	Name      string `gorm:"unique_index:app_configs_main_uidx"`
	Value     string
	Public    bool `gorm:"default:0"`
	TxHash    string
}

type BlockScanInfo struct {
	gorm.Model
	Type            string
	NetworkID       uint64
	Duration        uint   `gorm:"default:0"`
	LastBlockNumber int64  `gorm:"default:1"`
	NumBlocks       int64  `gorm:"default:50"`
	LastBlockError  string `gorm:"type:longtext"`
	Enabled         bool   `gorm:"default:1"`
	ContractAddrs   string `gorm:"type:longtext"`
}

type TokenPrice struct {
	gorm.Model
	NetworkID    uint64
	Symbol       string
	Price        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Last24hPrice numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type BTCL1InscribeTx struct {
	gorm.Model
	TxHash         string
	InscribeTxHash string
	Error          string `gorm:"type:longtext"`
}
