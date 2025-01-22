package factory

import (
	"context"
	"solo/chains/local_v1"

	"solo/chains/base_new"
	"solo/chains/local"

	"solo/chains/abstract_testnet"
	"solo/chains/base"
	interCommon "solo/chains/common"
	"solo/config"
	"solo/internal/contracts/gpu_manager"
	"solo/internal/contracts/staking_hub"
	staking_hub_v1 "solo/internal/contracts/v1/staking_hub"
	"solo/internal/port"
	"solo/internal/usecase"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

func NewMiner(cnf *config.Config) (port.IMiner, error) {
	ctx := context.Background()
	var err error
	defer func() {
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("chainFactory",
				zap.String("chain", cnf.ChainID),
				zap.Error(err),
			)
		} else {
			// logger.GetLoggerInstanceFromContext(ctx).Info("chainFactory", zap.String("chain", cnf.ChainID))
		}
	}()

	cm, err := interCommon.NewCommon(ctx, cnf)
	if err != nil {
		return nil, err
	}

	switch cnf.ChainID {
	case "31337":

		c, err := local_v1.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}
		sthub, err := staking_hub_v1.NewStakingHub(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}
		s := local_v1.NewStaking(cm, sthub)
		// cluster, _ := &base.
		miner := usecase.NewMiner(c, s, cm, cnf, nil)
		return miner, nil

	case "8453___": // old
		c, err := base.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}
		sthub, err := staking_hub.NewStakingHub(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}
		s := base.NewStaking(cm, sthub)
		// cluster, _ := &base.
		miner := usecase.NewMiner(c, s, cm, cnf, nil)
		return miner, nil
	case "8453": // new
		c, err := base_new.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}

		sthub, err := gpu_manager.NewGpuManager(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}

		s := base_new.NewStaking(cm, sthub)
		cluster, err := base_new.NewCluster(cm)
		if err != nil {
			return nil, err
		}

		miner := usecase.NewMiner(c, s, cm, cnf, cluster)
		return miner, nil
	case "11124":
		c, err := abstract_testnet.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}
		sthub, err := gpu_manager.NewGpuManager(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}
		s := abstract_testnet.NewStaking(cm, sthub)
		cluster, err := abstract_testnet.NewCluster(cm)
		if err != nil {
			return nil, err
		}
		miner := usecase.NewMiner(c, s, cm, cnf, cluster)
		return miner, nil
	default: // localhost

		c, err := local.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}
		sthub, err := gpu_manager.NewGpuManager(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}
		s := local.NewStaking(cm, sthub)
		cluster, err := local.NewCluster(cm)
		if err != nil {
			return nil, err
		}
		miner := usecase.NewMiner(c, s, cm, cnf, cluster)
		return miner, nil
	}

	return nil, err
}
