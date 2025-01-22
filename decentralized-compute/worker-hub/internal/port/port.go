package port

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

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
	SetTask(task *model.Task)
	GetInferenceByMiner() ([]*big.Int, error)
	GetInferenceInfo(opt *bind.CallOpts, inferID uint64) (*model.InferInfo, error)
}

type ICommon interface {
	GetWalletAddres() common.Address
	GetStakingHubAddress() common.Address
	GetWorkerHubAddress() common.Address
	GetModelCollectionAddress() common.Address
	GetModelLoadBalancerAddress() common.Address
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
	Erc20Balance() (*big.Int, error)
}

type INewChainFlow interface {
	GetPendingTasks(ctx context.Context, startBlock, endBlock uint64, out chan *model.Task) error
	SubmitTask(ctx context.Context, inferenceID *big.Int, result []byte) (*types.Transaction, error)
	GetInferenceByMiner() ([]*big.Int, error)
	GetInferenceInfo(opt *bind.CallOpts, inferID uint64) (*model.InferInfo, error)
}

type IMiner interface {
	GetPendingTasks(ctx context.Context)
	ExecueteTasks(ctx context.Context)
	Verify() bool
	StakeForWorker() (*types.Transaction, error)
	JoinForMinting() (*types.Transaction, error)
	MakeVerify() (*types.Transaction, *types.Transaction, error)
	Info() (*model.MinerInfo, error)
	ClaimReward() (*types.Transaction, error)
	GetChainCommon() ICommon
	GetCluster() ICluster
	GetConfig() *config.Config
}

type IStaking interface {
	IsStaked() (bool, error)
	StakeForWorker() (*types.Transaction, error)
	JoinForMinting() (*types.Transaction, error)
	ClaimReward() (*types.Transaction, error)
	RewardToClaim(opts *bind.CallOpts) (*big.Int, error)
}

type ICluster interface {
	CreateCluster(version int, minHardware int, modelName, modelType string) (*types.Transaction, *big.Int, error)
	CreateAGroupOfCluster(groupName string, clusterIDs []*big.Int) (*types.Transaction, error)
	AddClustersToGroup(groupName string, clusterIDs []*big.Int) (*types.Transaction, error)
	RemoveClustersFromGroup(groupName string, clusterIDs []*big.Int) (*types.Transaction, error)
}

type ICMDLocalChain interface {
	GetPrivateKey() string
	SetGasPrice(gp *big.Int)
	SetGasLimit(gl uint64)
	DeployContracts(rpc, chainID, prvkey string) (*model.LocalChain, error)
	DeployContract(rpc, chainID, prvkey, contractName string) (*model.LocalChain, error)
	MintWrappedEAI(rpc, chainID, mintAmount, prvkey string) (*types.Transaction, error)
	SetWEAIForStakingHub(client *ethclient.Client, prvkey string) (*types.Transaction, error)
	CreateMinerAddress(rpc, chainID, prvkey string) (*string, *string, error)
	StartHardHat() error
	StartOllama() error
	CreateInfer(prompt []model.LLMInferMessage) (*types.Transaction, *uint64, *string, error)
	StartMinerLogic() error
	CreateConfigENV(minerAddress string, index int) error
	DeployContractLogic() error
	ReadLocalChainCnf() *model.LocalChain
}
