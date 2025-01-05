package config

import (
	"context"
	"errors"
	"os"
	"strconv"

	"solo/pkg/logger"

	"go.uber.org/zap"

	"github.com/joho/godotenv"
)

const (
	IPFSPrefix = "ipfs://"
)

type Config struct {
	Rpc               string
	Account           string
	StakingHubAddress string
	WorkerHubAddress  string
	ApiUrl            string
	ApiKey            string
	LighthouseKey     string
	ModelAddress      string
	ChainID           string
	Erc20Address      string
	DebugMode         bool
	ClusterID         string
	ModelName         string
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

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
	dmode := os.Getenv("DEBUG_MODE")
	if dmode != "" {
		dmodeBool, errP := strconv.ParseBool(dmode)
		if errP == nil {
			cfg.DebugMode = dmodeBool
		}
	}

	logger.GetLoggerInstanceFromContext(context.Background()).Info("ReadConfig",
		zap.Any("cfg", cfg),
	)
	// validate
	if cfg.LighthouseKey == "" {
		return nil, errors.New("Lighthouse key is missing. Let's configure it now.")
	}

	if cfg.ApiUrl == "" {
		return nil, errors.New("API URL is missing. Let's configure it now.")
	}

	if cfg.ApiKey == "" {
		return nil, errors.New("API KEY is missing. Let's configure it now.")
	}
	return cfg, nil
}
