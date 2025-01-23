package config

import (
	"encoding/json"
	"fmt"
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
	Server         ServerConfig `json:"server"`
	Mongodb        MongoConfig  `json:"mongodb"`
	FilePathInfer  string       `json:"file_path_infer"`
	SubmitFilePath bool         `json:"submit_file_path"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type MongoConfig struct {
	Uri string `json:"uri"`
	Db  string `json:"db"`
}

func (c ChatConfig) VerifyBeforeChat() error {
	if c.AgentContractAddress == "" {
		return fmt.Errorf("Agent contract address is empty")
	}
	if c.Rpc == "" {
		return fmt.Errorf("Chain RPC is empty")
	}
	if c.PrivateKey == "" {
		return fmt.Errorf("Private key to create infer onchain is empty")
	}
	if c.AgentID == "" {
		return fmt.Errorf("Agent ID is empty")
	}
	if c.ModelName == "" {
		return fmt.Errorf("Model name is empty")
	}

	return nil
}

type ChatConfig struct {
	ServerBaseUrl string `json:"server_base_url"`
	Contracts     struct {
		CollectionAddress          string `json:"collectionAddress"`
		DaoTokenAddress            string `json:"daoTokenAddress"`
		HybridModelAddress         string `json:"hybridModelAddress"`
		StakingHubAddress          string `json:"stakingHubAddress"`
		SystemPromptManagerAddress string `json:"systemPromptManagerAddress"`
		TreasuryAddress            string `json:"treasuryAddress"`
		WorkerHubAddress           string `json:"workerHubAddress"`
		WrappedTokenAddress        string `json:"wrappedTokenAddress"`
	} `json:"contracts"`
	Rpc               string `json:"rpc"`
	ModelName         string `json:"model_name"`
	UseExternalRunPod bool   `json:"use_external_run_pod"`
	RunPodInternal    string `json:"run_pod_internal"`
	RunPodExternal    string `json:"run_pod_external"`
	RunPodApiKey      string `json:"run_pod_api_key"`
	ModelId           string `json:"model_id"`
	ChainId           string `json:"chain_id"`
	PrivateKey        string `json:"private_key"`
	Miners            struct {
		X38Adec2Ab9C10D353A8B9F985Aa185B0Dc086864 struct {
			Address    string `json:"address"`
			PrivateKey string `json:"private_key"`
		} `json:"0x38adec2ab9c10d353a8b9f985aa185b0dc086864"`
		X78D65914C6A66Fbd4Ae59Ffe763B72643D33B6Bb struct {
			Address    string `json:"address"`
			PrivateKey string `json:"private_key"`
		} `json:"0x78d65914c6a66fbd4ae59ffe763b72643d33b6bb"`
		Xe046B727A1B76505105E4Fdead71687E10177388 struct {
			Address    string `json:"address"`
			PrivateKey string `json:"private_key"`
		} `json:"0xe046b727a1b76505105e4fdead71687e10177388"`
	} `json:"miners"`
	Platform             string `json:"platform"`
	AgentID              string `json:"agent_id"`
	AgentContractAddress string `json:"agent_contract_address"`
}

func Load() (*Config, error) {
	file, err := os.Open("decentralized-inference/config.json")
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
