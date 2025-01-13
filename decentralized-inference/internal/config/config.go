package config

import (
	"encoding/json"
	"os"
)

const (
	EAI_OldWH_ChainID      = "43338" // workerHub version 0 , other chain workerHub v1
	BaseChainID            = "8453"
	BaseChainIDInt         = 8453
	BitAiChainID           = "222671"
	DAGIChainID            = "222672"
	HermesChain            = "45762"
	ArbitrumChainID        = "42161"
	DuckChainID            = "5545"
	PolygonChainID         = "137"
	ZkSyncChainID          = "324"
	ZkSyncChainIDInt       = 324
	EthereumChainID        = "1"
	BscChainID             = "56"
	AbstractTestnetChainID = "11124"
	SubtensorEVMChainID    = "964"
	SubtensorEVMChainIDInt = 964
	SolanaChainID          = "1111"
	SolanaModelID          = "990001"
	IPFSPrefix             = "ipfs://"

	GasLimitDefault         = uint64(10000000)
	BaseFeeWiggleMultiplier = 2
)

var config *Config

func GetConfig() *Config {
	return config
}

type Config struct {
	Server  ServerConfig `json:"server"`
	Mongodb MongoConfig  `json:"mongodb"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type MongoConfig struct {
	Uri string `json:"uri"`
	Db  string `json:"db"`
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
