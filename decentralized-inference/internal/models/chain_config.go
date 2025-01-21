package models

import (
	"encoding/json"
	"github.com/kamva/mgm/v3"
	"math/rand"
)

type ChainType string

const (
	ChainTypeEth    ChainType = "eth"
	ChainTypeZkSync ChainType = "zksync"
)

type ChainConfig struct {
	mgm.DefaultModel     `bson:",inline"`
	ChainID              string            `bson:"chain_id" json:"chain_id"`
	Type                 ChainType         `bson:"type" json:"type"`
	Name                 string            `bson:"name" json:"name"`
	Explorer             string            `bson:"explorer" json:"explorer"`
	PaymasterAddress     string            `bson:"paymaster_address" json:"paymaster_address"`
	PaymasterFeeZero     bool              `bson:"paymaster_fee_zero" json:"paymaster_fee_zero"`
	PaymasterToken       string            `bson:"paymaster_token" json:"paymaster_token"`
	WorkerHubAddress     string            `bson:"worker_hub_address" json:"worker_hub_address"`
	AgentContractAddress string            `bson:"agent_contract_address" json:"agent_contract_address"`
	ListRPC              []string          `bson:"list_rpc" json:"list_rpc"`
	GasLimit             uint64            `bson:"gas_limit" json:"gas_limit"`
	SupportModelNames    map[string]string `json:"support_model_names" bson:"support_model_names"`
	SupportStoreRaw      bool              `bson:"support_store_raw" json:"support_store_raw"`
	BackwardBlockNumber  uint64            `bson:"backward_block_number" json:"backward_block_number"` // need check duplicate event
}

func (ChainConfig) CollectionName() string {
	return "chain_config"
}

var CheckValidRpc func(rpc string) (bool, error)

func (chain *ChainConfig) GetRPC() string {

	var validRpc []string
	for _, rpc := range chain.ListRPC {
		valid, _ := CheckValidRpc(rpc)
		if valid {
			validRpc = append(validRpc, rpc)
		}
	}
	rpc := ""
	if len(validRpc) > 0 {
		rpc = validRpc[rand.Intn(len(validRpc))]
	}
	return rpc
}
func (chain *ChainConfig) MakeCopy() *ChainConfig {
	if chain == nil {
		return nil
	}
	data, _ := json.Marshal(chain)
	NewChainConfig := &ChainConfig{}
	json.Unmarshal(data, NewChainConfig)
	return NewChainConfig
}

type AppConfig struct {
	mgm.DefaultModel   `json:"-" bson:",inline"`
	ModelToChain       map[string]string `json:"model_to_chain" bson:"model_to_chain"`
	AIModelDescription map[string]string `json:"ai_model_description" bson:"ai_model_description"`
	AIModelNameDetail  map[string]string `json:"ai_model_name_detail" bson:"ai_model_name_detail"`
}

func (AppConfig) CollectionName() string {
	return "app_config"
}
