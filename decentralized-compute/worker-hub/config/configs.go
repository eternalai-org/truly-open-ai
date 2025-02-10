package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	IPFSPrefix = "ipfs://"
)

type Config struct {
	Rpc                      string
	PubSubURL                string
	Account                  string
	StakingHubAddress        string
	WorkerHubAddress         string
	ApiUrl                   string
	ApiKey                   string
	LighthouseKey            string
	ModelAddress             string
	ChainID                  string
	Erc20Address             string
	DebugMode                bool
	ClusterID                string
	ModelName                string
	Platform                 string
	ModelCollectionAddress   string
	ModelLoadBalancerAddress string
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := godotenv.Overload(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cfg.PubSubURL = os.Getenv("PUBSUB_URL")
	cfg.Rpc = os.Getenv("CHAIN_RPC")
	cfg.Account = os.Getenv("ACCOUNT_PRIV")
	cfg.StakingHubAddress = os.Getenv("STAKING_HUB_ADDRESS")
	cfg.WorkerHubAddress = os.Getenv("WORKER_HUB_ADDRESS")
	cfg.ApiUrl = os.Getenv("API_URL")
	cfg.ApiKey = os.Getenv("API_KEY")
	cfg.LighthouseKey = os.Getenv("LIGHT_HOUSE_API_KEY")
	cfg.ModelAddress = os.Getenv("MODEL_ADDRESS")
	cfg.ChainID = os.Getenv("CHAIN_ID")
	cfg.ClusterID = os.Getenv("CLUSTER_ID")
	cfg.ModelName = os.Getenv("MODEL_NAME")
	cfg.Erc20Address = os.Getenv("ERC20_ADDRESS")
	cfg.Platform = os.Getenv("PLATFORM")
	dmode := os.Getenv("DEBUG_MODE")
	modelCollectionAddress := os.Getenv("COLLECTION_ADDRESS")
	modelLoadBalancerAddress := os.Getenv("MODEL_LOAD_BALANCER_ADDRESS")
	if dmode != "" {
		dmodeBool, errP := strconv.ParseBool(dmode)
		if errP == nil {
			cfg.DebugMode = dmodeBool
		}
	}

	cfg.ModelLoadBalancerAddress = modelLoadBalancerAddress
	cfg.ModelCollectionAddress = modelCollectionAddress
	return cfg, nil
}

func (cfg *Config) Verify() error {
	// validate
	/*
		if cfg.LighthouseKey == "" {
			return errors.New("Lighthouse key is missing. Let's configure it now.")
		}*/

	if cfg.ApiUrl == "" {
		return errors.New("API URL is missing. Let's configure it now.")
	}

	/*
		if cfg.ApiKey == "" {
			return errors.New("API KEY is missing. Let's configure it now.")
		}*/
	return nil
}
