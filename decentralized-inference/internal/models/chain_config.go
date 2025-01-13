package models

import "github.com/kamva/mgm/v3"

type ChainType string

const (
	ChainTypeEth    ChainType = "eth"
	ChainTypeZkSync ChainType = "zksync"
)

type ChainConfig struct {
	mgm.DefaultModel  `bson:",inline"`
	ChainID           string            `bson:"chain_id" json:"chain_id"`
	Type              ChainType         `bson:"type" json:"type"`
	Name              string            `bson:"name" json:"name"`
	Explorer          string            `bson:"explorer" json:"explorer"`
	PaymasterAddress  string            `bson:"paymaster_address" json:"paymaster_address"`
	PaymasterFeeZero  bool              `bson:"paymaster_fee_zero" json:"paymaster_fee_zero"`
	PaymasterToken    string            `bson:"paymaster_token" json:"paymaster_token"`
	WorkerHubAddress  string            `bson:"worker_hub_address" json:"worker_hub_address"`
	ListRPC           []string          `bson:"list_rpc" json:"list_rpc"`
	GasLimit          uint64            `bson:"gas_limit" json:"gas_limit"`
	SupportModelNames map[string]string `json:"support_model_names" bson:"support_model_names"`
	SupportStoreRaw   bool              `bson:"support_store_raw" json:"support_store_raw"`
}
