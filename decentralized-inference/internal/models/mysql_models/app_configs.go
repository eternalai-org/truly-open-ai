package mysql_models

import (
	"gorm.io/gorm"
)

var KeyConfigNameWalletDeploy = "wallet_deploy"
var KeyConfigNameModelId = "knowledge_base_model_id"
var KeyConfigNameKnowledgeBaseWorkerHubAddress = "knowledge_base_worker_hub_address"
var KeyConfigNameKnowledgeBaseTokenContractAddress = "knowledge_base_token_contract_address"

type AppConfig struct {
	gorm.Model
	NetworkID uint64 `gorm:"unique_index:app_configs_main_uidx"`
	Name      string `gorm:"unique_index:app_configs_main_uidx"`
	Value     string
	Public    bool `gorm:"default:0"`
	TxHash    string
}
