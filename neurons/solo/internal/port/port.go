package port

import (
	"context"
	"math/big"

	"solo/config"

	"solo/internal/contracts/erc20"
	"solo/internal/model"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

type IChain interface {
	GetPendingTasks(ctx context.Context, fromblock, toBlock uint64, out chan *model.Task) error
	SubmitTask(ctx context.Context, assigmentID *big.Int, result []byte) (*types.Transaction, error)
}

type ICommon interface {
	GetWalletAddres() common.Address
	GetStakingHubAddress() common.Address
	GetWorkerHubAddress() common.Address
	GetErc20contractAddress() common.Address
	CurrentBlock() uint64
	FromBlock(uint64) uint64
	ToBlock() uint64
	GetErc20contract() *erc20.Erc20
	GetClient() *ethclient.Client
	GetPrivateKey() string
	GetGasLimit() uint64
	GetModelAddress() string
	GetConfig() *config.Config
}

type ITaskWatcher interface {
	GetPendingTasks(ctx context.Context)
	ExecueteTasks(ctx context.Context)
	Verify() bool
	MakeVerify() error
}

type IStaking interface {
	IsStaked() (bool, error)
	StakeForWorker() error
	JoinForMinting() error
}
