package config

import (
	"encoding/json"
	"os"
)

const (
	EAI_OldWH_ChainID       = "43338" // workerHub version 0 , other chain workerHub v1
	BaseChainID             = "8453"
	BaseChainIDInt          = 8453
	BitAiChainID            = "222671"
	DAGIChainID             = "222672"
	HermesChain             = "45762"
	ArbitrumChainID         = "42161"
	DuckChainID             = "5545"
	PolygonChainID          = "137"
	ZkSyncChainID           = "324"
	ZkSyncChainIDInt        = 324
	EthereumChainID         = "1"
	BscChainID              = "56"
	AbstractTestnetChainID  = "11124"
	SubtensorEVMChainID     = "964"
	SubtensorEVMChainIDInt  = 964
	SolanaChainID           = "1111"
	SolanaModelID           = "990001"
	IPFSPrefix              = "ipfs://"
	FilePrefix              = "file://"
	GasLimitDefault         = uint64(10000000)
	BaseFeeWiggleMultiplier = 2
)

var config *Config

func GetConfig() *Config {
	return config
}

type Config struct {
	Server        ServerConfig `json:"server"`
	Mongodb       MongoConfig  `json:"mongodb"`
	FilePathInfer string       `json:"file_path_infer"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type MongoConfig struct {
	Uri string `json:"uri"`
	Db  string `json:"db"`
}

type ChatConfig struct {
	ServerBaseUrl                  string `json:"server_base_url"`
	ChainRpc                       string `json:"chain_rpc"`
	Dagent721ContractAddress       string `json:"dagent721_contract_address"`
	AgentID                        string `json:"agent_id"`
	PromptSchedulerContractAddress string `json:"prompt_scheduler_contract_address"`
	InferWalletKey                 string `json:"infer_wallet_key"`
}

func Load() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		panic("config.json not found, please copy config.json.example to config.json and fill in the values")
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		panic("config.json is invalid, please check the values")
	}
	config = &v
	return config, nil
}

func init() {
	if _, err := Load(); err != nil {
		panic(err)
	}
}
