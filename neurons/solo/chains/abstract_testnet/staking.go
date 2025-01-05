package abstract_testnet

import (
	"context"
	"strconv"

	"solo/pkg/logger"

	"go.uber.org/zap"

	"solo/internal/contracts/gpu_manager"
	"solo/internal/port"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/core/types"
)

type staking struct {
	stakingHub *gpu_manager.GpuManager
	common     port.ICommon
}

func NewStaking(common port.ICommon, stakingHub *gpu_manager.GpuManager) port.IStaking {
	return &staking{
		common:     common,
		stakingHub: stakingHub,
	}
}

func (s *staking) IsStaked() (bool, error) {
	workerInfo, err := s.stakingHub.Miners(nil, s.common.GetWalletAddres())
	if err != nil {
		return false, err
	}

	minStake, err := s.stakingHub.MinerMinimumStake(nil)
	if err != nil {
		return false, err
	}

	if workerInfo.Stake.Cmp(minStake) < 0 {
		return false, nil
	}

	pendingUnstake, err := s.stakingHub.MinerUnstakeRequests(nil, s.common.GetWalletAddres())
	if err != nil {
		return false, err
	}

	_ = pendingUnstake
	/*
		tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
		if tskw.status.pendingUnstakeAmount.Cmp(new(big.Int).SetUint64(0)) > 0 {
			tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)
		}

		tskw.status.assignModel = workerInfo.ModelAddress.Hex()
		tskw.status.stakedAmount = workerInfo.Stake*/

	return true, nil
}

func (s *staking) StakeForWorker() error {
	ctx := context.Background()
	balance, err := s.common.GetErc20contract().BalanceOf(nil, s.common.GetWalletAddres())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return err
	}

	err = eth.ApproveERC20(
		ctx, s.common.GetClient(), s.common.GetPrivateKey(), s.common.GetStakingHubAddress(),
		s.common.GetErc20contractAddress(), int64(s.common.GetGasLimit()),
	)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return err
	}

	minStake, err := s.stakingHub.MinerMinimumStake(nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.Error(err))

		return err
	}

	// auth.Value = minStake
	tx := new(types.Transaction)
	if s.common.GetConfig().ClusterID != "" {
		modelID, err1 := strconv.Atoi(s.common.GetConfig().ClusterID)
		if err1 == nil {
			tx, err = s.stakingHub.RegisterMiner0(auth, 1, uint32(modelID))
			if err != nil {

				logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
					zap.String("worker_address", s.common.GetWalletAddres().Hex()),
					zap.Int("model_id", modelID),
					zap.String("balance", balance.String()),
					zap.String("min_stake", minStake.String()),
					zap.Error(err))
				return err
			}

			logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
				zap.String("worker_address", s.common.GetWalletAddres().Hex()),
				zap.Int("model_id", modelID),
				zap.String("balance", balance.String()),
				zap.String("min_stake", minStake.String()),
				zap.Any("tx", tx.Hash().Hex()),
			)
		}
	} else {
		tx, err = s.stakingHub.RegisterMiner(auth, 1)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
				zap.String("worker_address", s.common.GetWalletAddres().Hex()),
				zap.String("balance", balance.String()),
				zap.String("min_stake", minStake.String()),
				zap.Error(err))
			return err
		}

		logger.GetLoggerInstanceFromContext(ctx).Info("StakeForWorker",
			zap.String("worker_address", s.common.GetWalletAddres().Hex()),
			zap.String("balance", balance.String()),
			zap.String("min_stake", minStake.String()),
			zap.Any("tx", tx.Hash().Hex()),
		)
	}

	// TODO - here
	_ = tx
	_ = balance
	_ = minStake
	return nil
}

func (s *staking) JoinForMinting() error {
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, s.common.GetClient(), s.common.GetPrivateKey(), int64(s.common.GetGasLimit()))
	if err != nil {
		return err
	}

	tx, err := s.stakingHub.JoinForMinting(auth)
	if err != nil {
		return err
	}

	_ = tx
	return nil
}
