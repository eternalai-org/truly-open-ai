package models

import "github.com/kamva/mgm/v3"

type ContractSyncState struct {
	mgm.DefaultModel `bson:",inline"`
	ContractAddress  string `json:"contract_address" bson:"contract_address"`
	Job              string `json:"job" bson:"job"`
	LastSyncedBlock  uint64 `json:"last_synced_block" bson:"last_synced_block"`
	ResyncToBlock    uint64 `json:"resync_to_block" bson:"resync_to_block"`
	ResyncFromBlock  uint64 `json:"resync_from_block" bson:"resync_from_block"`
	ClearDataAndSync bool   `json:"clear_data_and_sync" bson:"clear_data_and_sync"`
}

func (ContractSyncState) CollectionName() string {
	return "contract_sync_state"
}
