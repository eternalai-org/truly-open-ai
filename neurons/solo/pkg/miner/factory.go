package miner

import (
	"context"
	"errors"

	"solo/chains/base_new"

	"solo/chains/abstract_testnet"
	"solo/chains/base"
	interCommon "solo/chains/common"
	"solo/config"
	"solo/internal/contracts/gpu_manager"
	"solo/internal/contracts/staking_hub"
	"solo/internal/port"
	"solo/internal/usecase"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

func NewMiner(cnf *config.Config) (port.ITaskWatcher, error) {
	ctx := context.Background()
	var err error
	defer func() {
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("chainFactory",
				zap.String("chain", cnf.ChainID),
				zap.Error(err),
			)
		} else {
			logger.GetLoggerInstanceFromContext(ctx).Info("chainFactory", zap.String("chain", cnf.ChainID))
		}
	}()

	cm, err := interCommon.NewCommon(ctx, cnf)
	if err != nil {
		return nil, err
	}

	switch cnf.ChainID {
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

		taskWatcher := usecase.NewTasksWatcher(c, s, cm, cnf)
		return taskWatcher, nil
	case "8453": // new
		c, err := base_new.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}

		sthub, err := gpu_manager.NewGpuManager(cm.GetStakingHubAddress(), cm.GetClient())
		if err != nil {
			return nil, err
		}

		s := abstract_testnet.NewStaking(cm, sthub)

		taskWatcher := usecase.NewTasksWatcher(c, s, cm, cnf)
		return taskWatcher, nil

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

		taskWatcher := usecase.NewTasksWatcher(c, s, cm, cnf)
		return taskWatcher, nil
	default:
		// not support
		err = errors.New("not support")
	}

	return nil, err
}
