package workerhub

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IWorkerHubAssignmentInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignmentInfo struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Value        *big.Int
	Input        []byte
	ModelAddress common.Address
	Creator      common.Address
	ExpiredAt    *big.Int
}

// IWorkerHubInferenceInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubInferenceInfo struct {
	InferenceId      *big.Int
	Input            []byte
	Output           []byte
	Value            *big.Int
	DisputingAddress common.Address
	ModelAddress     common.Address
	ExpiredAt        *big.Int
	Status           uint8
	Creator          common.Address
}

// IWorkerHubWorkerInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubWorkerInfo struct {
	WorkerAddress    common.Address
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}

// WorkerHubMetaData contains all meta data concerning the WorkerHub contract.
var WorkerHubMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"assignmentNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"worker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"disapprovalCount\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blocksPerEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimReward\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputeInfer\",\"inputs\":[{\"name\":\"_assignmentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disputingTimeLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disqualificationPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feePercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finePercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInferences\",\"inputs\":[{\"name\":\"_inferenceIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIWorkerHub.InferenceInfo[]\",\"components\":[{\"name\":\"inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"output\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"disputingAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiredAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIWorkerHub.InferenceStatus\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinerAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinerAddressesOfModel\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMiners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"components\":[{\"name\":\"workerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMiningAssignments\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"components\":[{\"name\":\"assignmentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiredAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMintingAssignmentsOfInference\",\"inputs\":[{\"name\":\"_inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"components\":[{\"name\":\"assignmentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiredAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getModelAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNOMiner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddressesOfModel\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"components\":[{\"name\":\"workerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseMinerStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"increaseValidatorStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"infer\",\"inputs\":[{\"name\":\"_input\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_creator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"inferenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feePercentage\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minerMinimumStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validatorMinimumStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_miningTimeLimit\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"_minerRequirement\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_blocksPerEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardPerEpochBasedOnPerf\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardPerEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_unstakeDelayTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"_penaltyDuration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"_finePercentage\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAssignmentPending\",\"inputs\":[{\"name\":\"_assignmentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"joinForMinting\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinForValidating\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumTier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minerMinimumStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minerRequirement\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minerUnstakeRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlockAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"miners\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"miningTimeLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"models\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"minimumFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multiplier\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"penaltyDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerMiner\",\"inputs\":[{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registerModel\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_minimumFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerValidator\",\"inputs\":[{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveInference\",\"inputs\":[{\"name\":\"_inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"restakeForMiner\",\"inputs\":[{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardInEpoch\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"perfReward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epochReward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalTaskCompleted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMiner\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardPerEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardToClaim\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlocksPerEpoch\",\"inputs\":[{\"name\":\"_blocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinePercentage\",\"inputs\":[{\"name\":\"_finePercentage\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNewRewardInEpoch\",\"inputs\":[{\"name\":\"_newRewardAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPenaltyDuration\",\"inputs\":[{\"name\":\"_penaltyDuration\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnstakDelayTime\",\"inputs\":[{\"name\":\"_newUnstakeDelayTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashMiner\",\"inputs\":[{\"name\":\"_miner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isFined\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitSolution\",\"inputs\":[{\"name\":\"_assigmentId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topUpInfer\",\"inputs\":[{\"name\":\"_inferenceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterMiner\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterModel\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeDelayTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeForMiner\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeForValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMiningTimeLimit\",\"inputs\":[{\"name\":\"_miningTimeLimit\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateModelMinimumFee\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateModelTier\",\"inputs\":[{\"name\":\"_model\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tier\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatingTimeLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorDisputed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorMinimumStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorUnstakeRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlockAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BlocksPerEpoch\",\"inputs\":[{\"name\":\"oldBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinePercentageUpdated\",\"inputs\":[{\"name\":\"oldPercent\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newPercent\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FraudulentMinerPenalized\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"treasury\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fine\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InferenceDisputation\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assigmentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InferenceStatusUpdate\",\"inputs\":[{\"name\":\"inferenceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerDeactivated\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"modelAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activeTime\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerExtraStake\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerJoin\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerRegistration\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerUnregistration\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinerUnstake\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MiningTimeLimitUpdate\",\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelMinimumFeeUpdate\",\"inputs\":[{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minimumFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelRegistration\",\"inputs\":[{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"minimumFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelTierUpdate\",\"inputs\":[{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tier\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModelUnregistration\",\"inputs\":[{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewAssignment\",\"inputs\":[{\"name\":\"assignmentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"inferenceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"expiredAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewInference\",\"inputs\":[{\"name\":\"inferenceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PenaltyDurationUpdated\",\"inputs\":[{\"name\":\"oldDuration\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"newDuration\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Restake\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"restake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"model\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardClaim\",\"inputs\":[{\"name\":\"worker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardPerEpoch\",\"inputs\":[{\"name\":\"oldReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SolutionSubmission\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assigmentId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TopUpInfer\",\"inputs\":[{\"name\":\"inferenceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFee\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mingingFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"treasury\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"protocolFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeDelayTime\",\"inputs\":[{\"name\":\"oldDelayTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDelayTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorExtraStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJoin\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistration\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnregistration\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnstake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressSet_DuplicatedValue\",\"inputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressSet_ValueNotFound\",\"inputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedTransfer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InferMustBeSolvingState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBlockValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInferenceStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMiner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidModel\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinerInDeactivationTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MiningSessionEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MiningSessionNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughMiners\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NullStake\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakeTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillBeingLocked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Uint256Set_DuplicatedValue\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatingSessionNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorInDeactivationTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroValue\",\"inputs\":[]}]",
}

// WorkerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerHubMetaData.ABI instead.
var WorkerHubABI = WorkerHubMetaData.ABI

// WorkerHub is an auto generated Go binding around an Ethereum contract.
type WorkerHub struct {
	WorkerHubCaller     // Read-only binding to the contract
	WorkerHubTransactor // Write-only binding to the contract
	WorkerHubFilterer   // Log filterer for contract events
}

// WorkerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerHubSession struct {
	Contract     *WorkerHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerHubCallerSession struct {
	Contract *WorkerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WorkerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerHubTransactorSession struct {
	Contract     *WorkerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WorkerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerHubRaw struct {
	Contract *WorkerHub // Generic contract binding to access the raw methods on
}

// WorkerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerHubCallerRaw struct {
	Contract *WorkerHubCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerHubTransactorRaw struct {
	Contract *WorkerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerHub creates a new instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHub(address common.Address, backend bind.ContractBackend) (*WorkerHub, error) {
	contract, err := bindWorkerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerHub{WorkerHubCaller: WorkerHubCaller{contract: contract}, WorkerHubTransactor: WorkerHubTransactor{contract: contract}, WorkerHubFilterer: WorkerHubFilterer{contract: contract}}, nil
}

// NewWorkerHubCaller creates a new read-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubCaller(address common.Address, caller bind.ContractCaller) (*WorkerHubCaller, error) {
	contract, err := bindWorkerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubCaller{contract: contract}, nil
}

// NewWorkerHubTransactor creates a new write-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerHubTransactor, error) {
	contract, err := bindWorkerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTransactor{contract: contract}, nil
}

// NewWorkerHubFilterer creates a new log filterer instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerHubFilterer, error) {
	contract, err := bindWorkerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerHubFilterer{contract: contract}, nil
}

// bindWorkerHub binds a generic wrapper to an already deployed contract.
func bindWorkerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.WorkerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transact(opts, method, params...)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) AssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHub *WorkerHubCaller) Assignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignments", arg0)

	outstruct := new(struct {
		InferenceId      *big.Int
		Output           []byte
		Worker           common.Address
		DisapprovalCount uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Output = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Worker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.DisapprovalCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHub *WorkerHubSession) Assignments(arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHub *WorkerHubCallerSession) Assignments(arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.BlocksPerEpoch(&_WorkerHub.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.BlocksPerEpoch(&_WorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.CurrentEpoch(&_WorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.CurrentEpoch(&_WorkerHub.CallOpts)
}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) DisputingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "disputingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubSession) DisputingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.DisputingTimeLimit(&_WorkerHub.CallOpts)
}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) DisputingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.DisputingTimeLimit(&_WorkerHub.CallOpts)
}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) DisqualificationPercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "disqualificationPercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) DisqualificationPercentage() (uint16, error) {
	return _WorkerHub.Contract.DisqualificationPercentage(&_WorkerHub.CallOpts)
}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) DisqualificationPercentage() (uint16, error) {
	return _WorkerHub.Contract.DisqualificationPercentage(&_WorkerHub.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FeePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FeePercentage() (uint16, error) {
	return _WorkerHub.Contract.FeePercentage(&_WorkerHub.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FeePercentage() (uint16, error) {
	return _WorkerHub.Contract.FeePercentage(&_WorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FinePercentage() (uint16, error) {
	return _WorkerHub.Contract.FinePercentage(&_WorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FinePercentage() (uint16, error) {
	return _WorkerHub.Contract.FinePercentage(&_WorkerHub.CallOpts)
}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHub *WorkerHubCaller) GetInferences(opts *bind.CallOpts, _inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getInferences", _inferenceIds)

	if err != nil {
		return *new([]IWorkerHubInferenceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubInferenceInfo)).(*[]IWorkerHubInferenceInfo)

	return out0, err

}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHub *WorkerHubSession) GetInferences(_inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	return _WorkerHub.Contract.GetInferences(&_WorkerHub.CallOpts, _inferenceIds)
}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHub *WorkerHubCallerSession) GetInferences(_inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	return _WorkerHub.Contract.GetInferences(&_WorkerHub.CallOpts, _inferenceIds)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddresses(&_WorkerHub.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddresses(&_WorkerHub.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinerAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCaller) GetMiners(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetMiners(&_WorkerHub.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCallerSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetMiners(&_WorkerHub.CallOpts)
}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubCaller) GetMiningAssignments(opts *bind.CallOpts) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMiningAssignments")

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubSession) GetMiningAssignments() ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMiningAssignments(&_WorkerHub.CallOpts)
}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubCallerSession) GetMiningAssignments() ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMiningAssignments(&_WorkerHub.CallOpts)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubCaller) GetMintingAssignmentsOfInference(opts *bind.CallOpts, _inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMintingAssignmentsOfInference", _inferenceId)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMintingAssignmentsOfInference(&_WorkerHub.CallOpts, _inferenceId)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHub *WorkerHubCallerSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMintingAssignmentsOfInference(&_WorkerHub.CallOpts, _inferenceId)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetModelAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getModelAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetModelAddresses(&_WorkerHub.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetModelAddresses(&_WorkerHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHub.Contract.GetNOMiner(&_WorkerHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHub.Contract.GetNOMiner(&_WorkerHub.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetValidatorAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getValidatorAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddresses(&_WorkerHub.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddresses(&_WorkerHub.CallOpts)
}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetValidatorAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getValidatorAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubSession) GetValidatorAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetValidatorAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCaller) GetValidators(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubSession) GetValidators() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetValidators(&_WorkerHub.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCallerSession) GetValidators() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetValidators(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) InferenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "inferenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubCaller) IsAssignmentPending(opts *bind.CallOpts, _assignmentId *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "isAssignmentPending", _assignmentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHub.Contract.IsAssignmentPending(&_WorkerHub.CallOpts, _assignmentId)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHub.Contract.IsAssignmentPending(&_WorkerHub.CallOpts, _assignmentId)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubSession) LastBlock() (*big.Int, error) {
	return _WorkerHub.Contract.LastBlock(&_WorkerHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) LastBlock() (*big.Int, error) {
	return _WorkerHub.Contract.LastBlock(&_WorkerHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubSession) MaximumTier() (uint16, error) {
	return _WorkerHub.Contract.MaximumTier(&_WorkerHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) MaximumTier() (uint16, error) {
	return _WorkerHub.Contract.MaximumTier(&_WorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinerMinimumStake(&_WorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinerMinimumStake(&_WorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubSession) MinerRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinerRequirement(&_WorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCallerSession) MinerRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinerRequirement(&_WorkerHub.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.MinerUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.MinerUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "miners", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Miners(&_WorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Miners(&_WorkerHub.CallOpts, arg0)
}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) MiningTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "miningTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubSession) MiningTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.MiningTimeLimit(&_WorkerHub.CallOpts)
}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) MiningTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.MiningTimeLimit(&_WorkerHub.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "models", arg0)

	outstruct := new(struct {
		MinimumFee *big.Int
		Tier       uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCallerSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubCaller) Multiplier(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "multiplier", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.Multiplier(&_WorkerHub.CallOpts, _miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.Multiplier(&_WorkerHub.CallOpts, _miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHub.Contract.PenaltyDuration(&_WorkerHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHub.Contract.PenaltyDuration(&_WorkerHub.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHub *WorkerHubCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "rewardInEpoch", arg0)

	outstruct := new(struct {
		PerfReward         *big.Int
		EpochReward        *big.Int
		TotalTaskCompleted *big.Int
		TotalMiner         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerfReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalTaskCompleted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalMiner = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHub *WorkerHubSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHub.Contract.RewardInEpoch(&_WorkerHub.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHub *WorkerHubCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHub.Contract.RewardInEpoch(&_WorkerHub.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.RewardPerEpoch(&_WorkerHub.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.RewardPerEpoch(&_WorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubSession) Treasury() (common.Address, error) {
	return _WorkerHub.Contract.Treasury(&_WorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) Treasury() (common.Address, error) {
	return _WorkerHub.Contract.Treasury(&_WorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHub.Contract.UnstakeDelayTime(&_WorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHub.Contract.UnstakeDelayTime(&_WorkerHub.CallOpts)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) ValidatingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingTimeLimit(&_WorkerHub.CallOpts)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingTimeLimit(&_WorkerHub.CallOpts)
}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHub *WorkerHubCaller) ValidatorDisputed(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorDisputed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHub *WorkerHubSession) ValidatorDisputed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _WorkerHub.Contract.ValidatorDisputed(&_WorkerHub.CallOpts, arg0, arg1)
}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) ValidatorDisputed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _WorkerHub.Contract.ValidatorDisputed(&_WorkerHub.CallOpts, arg0, arg1)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) ValidatorMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorMinimumStake(&_WorkerHub.CallOpts)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorMinimumStake(&_WorkerHub.CallOpts)
}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCaller) ValidatorUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubSession) ValidatorUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.ValidatorUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCallerSession) ValidatorUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.ValidatorUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubSession) Validators(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Validators(&_WorkerHub.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCallerSession) Validators(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Validators(&_WorkerHub.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCallerSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubTransactor) ClaimReward(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "claimReward", _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ClaimReward(&_WorkerHub.TransactOpts, _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubTransactorSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ClaimReward(&_WorkerHub.TransactOpts, _miner)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactor) DisputeInfer(opts *bind.TransactOpts, _assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "disputeInfer", _assignmentId)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubSession) DisputeInfer(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.DisputeInfer(&_WorkerHub.TransactOpts, _assignmentId)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactorSession) DisputeInfer(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.DisputeInfer(&_WorkerHub.TransactOpts, _assignmentId)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubTransactor) IncreaseMinerStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "increaseMinerStake")
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubSession) IncreaseMinerStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinerStake(&_WorkerHub.TransactOpts)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) IncreaseMinerStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinerStake(&_WorkerHub.TransactOpts)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubTransactor) IncreaseValidatorStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "increaseValidatorStake")
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubSession) IncreaseValidatorStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseValidatorStake(&_WorkerHub.TransactOpts)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) IncreaseValidatorStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseValidatorStake(&_WorkerHub.TransactOpts)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactor) Infer(opts *bind.TransactOpts, _input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "infer", _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xab121bf3.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactor) Initialize(opts *bind.TransactOpts, _treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "initialize", _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xab121bf3.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubSession) Initialize(_treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xab121bf3.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) Initialize(_treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForMinting(&_WorkerHub.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForMinting(&_WorkerHub.TransactOpts)
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHub *WorkerHubTransactor) JoinForValidating(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "joinForValidating")
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHub *WorkerHubSession) JoinForValidating() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForValidating(&_WorkerHub.TransactOpts)
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHub *WorkerHubTransactorSession) JoinForValidating() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForValidating(&_WorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerMiner", tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMiner(&_WorkerHub.TransactOpts, tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMiner(&_WorkerHub.TransactOpts, tier)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactor) RegisterValidator(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerValidator", tier)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubSession) RegisterValidator(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterValidator(&_WorkerHub.TransactOpts, tier)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterValidator(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterValidator(&_WorkerHub.TransactOpts, tier)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactor) ResolveInference(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "resolveInference", _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactorSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RestakeForMiner(&_WorkerHub.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RestakeForMiner(&_WorkerHub.TransactOpts, tier)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubTransactor) RewardToClaim(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "rewardToClaim", _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RewardToClaim(&_WorkerHub.TransactOpts, _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RewardToClaim(&_WorkerHub.TransactOpts, _miner)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setBlocksPerEpoch", _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetBlocksPerEpoch(&_WorkerHub.TransactOpts, _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetBlocksPerEpoch(&_WorkerHub.TransactOpts, _blocks)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactor) SetFinePercentage(opts *bind.TransactOpts, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setFinePercentage", _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFinePercentage(&_WorkerHub.TransactOpts, _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFinePercentage(&_WorkerHub.TransactOpts, _finePercentage)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, _newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setNewRewardInEpoch", _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetNewRewardInEpoch(&_WorkerHub.TransactOpts, _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetNewRewardInEpoch(&_WorkerHub.TransactOpts, _newRewardAmount)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubTransactor) SetPenaltyDuration(opts *bind.TransactOpts, _penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setPenaltyDuration", _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetPenaltyDuration(&_WorkerHub.TransactOpts, _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetPenaltyDuration(&_WorkerHub.TransactOpts, _penaltyDuration)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubTransactor) SetUnstakDelayTime(opts *bind.TransactOpts, _newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setUnstakDelayTime", _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetUnstakDelayTime(&_WorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetUnstakDelayTime(&_WorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubTransactor) SlashMiner(opts *bind.TransactOpts, _miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "slashMiner", _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.SlashMiner(&_WorkerHub.TransactOpts, _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubTransactorSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.SlashMiner(&_WorkerHub.TransactOpts, _miner, _isFined)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) SubmitSolution(opts *bind.TransactOpts, _assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "submitSolution", _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubTransactor) TopUpInfer(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "topUpInfer", _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.TopUpInfer(&_WorkerHub.TransactOpts, _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.TopUpInfer(&_WorkerHub.TransactOpts, _inferenceId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMiner(&_WorkerHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMiner(&_WorkerHub.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterModel(opts *bind.TransactOpts, _model common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterModel", _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterModel(&_WorkerHub.TransactOpts, _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterModel(&_WorkerHub.TransactOpts, _model)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterValidator")
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterValidator(&_WorkerHub.TransactOpts)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterValidator(&_WorkerHub.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForMiner(&_WorkerHub.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForMiner(&_WorkerHub.TransactOpts)
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHub *WorkerHubTransactor) UnstakeForValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unstakeForValidator")
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHub *WorkerHubSession) UnstakeForValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForValidator(&_WorkerHub.TransactOpts)
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnstakeForValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForValidator(&_WorkerHub.TransactOpts)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHub *WorkerHubTransactor) UpdateMiningTimeLimit(opts *bind.TransactOpts, _miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "updateMiningTimeLimit", _miningTimeLimit)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHub *WorkerHubSession) UpdateMiningTimeLimit(_miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateMiningTimeLimit(&_WorkerHub.TransactOpts, _miningTimeLimit)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHub *WorkerHubTransactorSession) UpdateMiningTimeLimit(_miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateMiningTimeLimit(&_WorkerHub.TransactOpts, _miningTimeLimit)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, _model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "updateModelMinimumFee", _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelMinimumFee(&_WorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactorSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelMinimumFee(&_WorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubTransactor) UpdateModelTier(opts *bind.TransactOpts, _model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "updateModelTier", _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelTier(&_WorkerHub.TransactOpts, _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubTransactorSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelTier(&_WorkerHub.TransactOpts, _model, _tier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// WorkerHubBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the WorkerHub contract.
type WorkerHubBlocksPerEpochIterator struct {
	Event *WorkerHubBlocksPerEpoch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubBlocksPerEpoch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubBlocksPerEpoch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubBlocksPerEpoch represents a BlocksPerEpoch event raised by the WorkerHub contract.
type WorkerHubBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHub *WorkerHubFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*WorkerHubBlocksPerEpochIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubBlocksPerEpochIterator{contract: _WorkerHub.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHub *WorkerHubFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubBlocksPerEpoch)
				if err := _WorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlocksPerEpoch is a log parse operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHub *WorkerHubFilterer) ParseBlocksPerEpoch(log types.Log) (*WorkerHubBlocksPerEpoch, error) {
	event := new(WorkerHubBlocksPerEpoch)
	if err := _WorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the WorkerHub contract.
type WorkerHubFinePercentageUpdatedIterator struct {
	Event *WorkerHubFinePercentageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubFinePercentageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubFinePercentageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubFinePercentageUpdated represents a FinePercentageUpdated event raised by the WorkerHub contract.
type WorkerHubFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHub *WorkerHubFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*WorkerHubFinePercentageUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubFinePercentageUpdatedIterator{contract: _WorkerHub.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHub *WorkerHubFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubFinePercentageUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinePercentageUpdated is a log parse operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHub *WorkerHubFilterer) ParseFinePercentageUpdated(log types.Log) (*WorkerHubFinePercentageUpdated, error) {
	event := new(WorkerHubFinePercentageUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the WorkerHub contract.
type WorkerHubFraudulentMinerPenalizedIterator struct {
	Event *WorkerHubFraudulentMinerPenalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubFraudulentMinerPenalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubFraudulentMinerPenalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the WorkerHub contract.
type WorkerHubFraudulentMinerPenalized struct {
	Miner        common.Address
	ModelAddress common.Address
	Treasury     common.Address
	Fine         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHub *WorkerHubFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (*WorkerHubFraudulentMinerPenalizedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubFraudulentMinerPenalizedIterator{contract: _WorkerHub.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHub *WorkerHubFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *WorkerHubFraudulentMinerPenalized, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubFraudulentMinerPenalized)
				if err := _WorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFraudulentMinerPenalized is a log parse operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHub *WorkerHubFilterer) ParseFraudulentMinerPenalized(log types.Log) (*WorkerHubFraudulentMinerPenalized, error) {
	event := new(WorkerHubFraudulentMinerPenalized)
	if err := _WorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInferenceDisputationIterator is returned from FilterInferenceDisputation and is used to iterate over the raw logs and unpacked data for InferenceDisputation events raised by the WorkerHub contract.
type WorkerHubInferenceDisputationIterator struct {
	Event *WorkerHubInferenceDisputation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubInferenceDisputationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInferenceDisputation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubInferenceDisputation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubInferenceDisputationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInferenceDisputationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInferenceDisputation represents a InferenceDisputation event raised by the WorkerHub contract.
type WorkerHubInferenceDisputation struct {
	Validator   common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceDisputation is a free log retrieval operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) FilterInferenceDisputation(opts *bind.FilterOpts, validator []common.Address, assigmentId []*big.Int) (*WorkerHubInferenceDisputationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "InferenceDisputation", validatorRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubInferenceDisputationIterator{contract: _WorkerHub.contract, event: "InferenceDisputation", logs: logs, sub: sub}, nil
}

// WatchInferenceDisputation is a free log subscription operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) WatchInferenceDisputation(opts *bind.WatchOpts, sink chan<- *WorkerHubInferenceDisputation, validator []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "InferenceDisputation", validatorRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInferenceDisputation)
				if err := _WorkerHub.contract.UnpackLog(event, "InferenceDisputation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInferenceDisputation is a log parse operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) ParseInferenceDisputation(log types.Log) (*WorkerHubInferenceDisputation, error) {
	event := new(WorkerHubInferenceDisputation)
	if err := _WorkerHub.contract.UnpackLog(event, "InferenceDisputation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdateIterator struct {
	Event *WorkerHubInferenceStatusUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInferenceStatusUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubInferenceStatusUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdate struct {
	InferenceId *big.Int
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []*big.Int) (*WorkerHubInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubInferenceStatusUpdateIterator{contract: _WorkerHub.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubInferenceStatusUpdate, inferenceId []*big.Int) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInferenceStatusUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) ParseInferenceStatusUpdate(log types.Log) (*WorkerHubInferenceStatusUpdate, error) {
	event := new(WorkerHubInferenceStatusUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerHub contract.
type WorkerHubInitializedIterator struct {
	Event *WorkerHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInitialized represents a Initialized event raised by the WorkerHub contract.
type WorkerHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerHubInitializedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerHubInitializedIterator{contract: _WorkerHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerHubInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInitialized)
				if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) ParseInitialized(log types.Log) (*WorkerHubInitialized, error) {
	event := new(WorkerHubInitialized)
	if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the WorkerHub contract.
type WorkerHubMinerDeactivatedIterator struct {
	Event *WorkerHubMinerDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerDeactivated represents a MinerDeactivated event raised by the WorkerHub contract.
type WorkerHubMinerDeactivated struct {
	Miner        common.Address
	ModelAddress common.Address
	ActiveTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHub *WorkerHubFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address) (*WorkerHubMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerDeactivatedIterator{contract: _WorkerHub.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHub *WorkerHubFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerDeactivated, miner []common.Address, modelAddress []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerDeactivated)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerDeactivated is a log parse operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHub *WorkerHubFilterer) ParseMinerDeactivated(log types.Log) (*WorkerHubMinerDeactivated, error) {
	event := new(WorkerHubMinerDeactivated)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the WorkerHub contract.
type WorkerHubMinerExtraStakeIterator struct {
	Event *WorkerHubMinerExtraStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerExtraStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerExtraStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerExtraStake represents a MinerExtraStake event raised by the WorkerHub contract.
type WorkerHubMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerExtraStakeIterator{contract: _WorkerHub.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerExtraStake)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerExtraStake is a log parse operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseMinerExtraStake(log types.Log) (*WorkerHubMinerExtraStake, error) {
	event := new(WorkerHubMinerExtraStake)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the WorkerHub contract.
type WorkerHubMinerJoinIterator struct {
	Event *WorkerHubMinerJoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerJoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerJoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerJoin represents a MinerJoin event raised by the WorkerHub contract.
type WorkerHubMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerJoinIterator{contract: _WorkerHub.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerJoin)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerJoin is a log parse operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) ParseMinerJoin(log types.Log) (*WorkerHubMinerJoin, error) {
	event := new(WorkerHubMinerJoin)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the WorkerHub contract.
type WorkerHubMinerRegistrationIterator struct {
	Event *WorkerHubMinerRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerRegistration represents a MinerRegistration event raised by the WorkerHub contract.
type WorkerHubMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*WorkerHubMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerRegistrationIterator{contract: _WorkerHub.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerRegistration is a log parse operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseMinerRegistration(log types.Log) (*WorkerHubMinerRegistration, error) {
	event := new(WorkerHubMinerRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the WorkerHub contract.
type WorkerHubMinerUnregistrationIterator struct {
	Event *WorkerHubMinerUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerUnregistration represents a MinerUnregistration event raised by the WorkerHub contract.
type WorkerHubMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerUnregistrationIterator{contract: _WorkerHub.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerUnregistration is a log parse operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) ParseMinerUnregistration(log types.Log) (*WorkerHubMinerUnregistration, error) {
	event := new(WorkerHubMinerUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the WorkerHub contract.
type WorkerHubMinerUnstakeIterator struct {
	Event *WorkerHubMinerUnstake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMinerUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerUnstake represents a MinerUnstake event raised by the WorkerHub contract.
type WorkerHubMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerUnstakeIterator{contract: _WorkerHub.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerUnstake)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerUnstake is a log parse operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) ParseMinerUnstake(log types.Log) (*WorkerHubMinerUnstake, error) {
	event := new(WorkerHubMinerUnstake)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMiningTimeLimitUpdateIterator is returned from FilterMiningTimeLimitUpdate and is used to iterate over the raw logs and unpacked data for MiningTimeLimitUpdate events raised by the WorkerHub contract.
type WorkerHubMiningTimeLimitUpdateIterator struct {
	Event *WorkerHubMiningTimeLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubMiningTimeLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMiningTimeLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubMiningTimeLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubMiningTimeLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMiningTimeLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMiningTimeLimitUpdate represents a MiningTimeLimitUpdate event raised by the WorkerHub contract.
type WorkerHubMiningTimeLimitUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMiningTimeLimitUpdate is a free log retrieval operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHub *WorkerHubFilterer) FilterMiningTimeLimitUpdate(opts *bind.FilterOpts) (*WorkerHubMiningTimeLimitUpdateIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WorkerHubMiningTimeLimitUpdateIterator{contract: _WorkerHub.contract, event: "MiningTimeLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchMiningTimeLimitUpdate is a free log subscription operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHub *WorkerHubFilterer) WatchMiningTimeLimitUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubMiningTimeLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMiningTimeLimitUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningTimeLimitUpdate is a log parse operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHub *WorkerHubFilterer) ParseMiningTimeLimitUpdate(log types.Log) (*WorkerHubMiningTimeLimitUpdate, error) {
	event := new(WorkerHubMiningTimeLimitUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the WorkerHub contract.
type WorkerHubModelMinimumFeeUpdateIterator struct {
	Event *WorkerHubModelMinimumFeeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelMinimumFeeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubModelMinimumFeeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the WorkerHub contract.
type WorkerHubModelMinimumFeeUpdate struct {
	Model      common.Address
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelMinimumFeeUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelMinimumFeeUpdateIterator{contract: _WorkerHub.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubModelMinimumFeeUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelMinimumFeeUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelMinimumFeeUpdate is a log parse operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*WorkerHubModelMinimumFeeUpdate, error) {
	event := new(WorkerHubModelMinimumFeeUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the WorkerHub contract.
type WorkerHubModelRegistrationIterator struct {
	Event *WorkerHubModelRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubModelRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelRegistration represents a ModelRegistration event raised by the WorkerHub contract.
type WorkerHubModelRegistration struct {
	Model      common.Address
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, tier []uint16) (*WorkerHubModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelRegistrationIterator{contract: _WorkerHub.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubModelRegistration, model []common.Address, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelRegistration is a log parse operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) ParseModelRegistration(log types.Log) (*WorkerHubModelRegistration, error) {
	event := new(WorkerHubModelRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the WorkerHub contract.
type WorkerHubModelTierUpdateIterator struct {
	Event *WorkerHubModelTierUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelTierUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubModelTierUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelTierUpdate represents a ModelTierUpdate event raised by the WorkerHub contract.
type WorkerHubModelTierUpdate struct {
	Model common.Address
	Tier  uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHub *WorkerHubFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelTierUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelTierUpdateIterator{contract: _WorkerHub.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHub *WorkerHubFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubModelTierUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelTierUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelTierUpdate is a log parse operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHub *WorkerHubFilterer) ParseModelTierUpdate(log types.Log) (*WorkerHubModelTierUpdate, error) {
	event := new(WorkerHubModelTierUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the WorkerHub contract.
type WorkerHubModelUnregistrationIterator struct {
	Event *WorkerHubModelUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubModelUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelUnregistration represents a ModelUnregistration event raised by the WorkerHub contract.
type WorkerHubModelUnregistration struct {
	Model common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) FilterModelUnregistration(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelUnregistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelUnregistrationIterator{contract: _WorkerHub.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubModelUnregistration, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelUnregistration is a log parse operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) ParseModelUnregistration(log types.Log) (*WorkerHubModelUnregistration, error) {
	event := new(WorkerHubModelUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the WorkerHub contract.
type WorkerHubNewAssignmentIterator struct {
	Event *WorkerHubNewAssignment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewAssignment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubNewAssignment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewAssignment represents a NewAssignment event raised by the WorkerHub contract.
type WorkerHubNewAssignment struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	ExpiredAt    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) FilterNewAssignment(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubNewAssignmentIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewAssignmentIterator{contract: _WorkerHub.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *WorkerHubNewAssignment, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewAssignment)
				if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAssignment is a log parse operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) ParseNewAssignment(log types.Log) (*WorkerHubNewAssignment, error) {
	event := new(WorkerHubNewAssignment)
	if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the WorkerHub contract.
type WorkerHubNewInferenceIterator struct {
	Event *WorkerHubNewInference // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewInference)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubNewInference)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewInference represents a NewInference event raised by the WorkerHub contract.
type WorkerHubNewInference struct {
	InferenceId *big.Int
	Model       common.Address
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewInferenceIterator{contract: _WorkerHub.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerHubNewInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewInference)
				if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewInference is a log parse operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseNewInference(log types.Log) (*WorkerHubNewInference, error) {
	event := new(WorkerHubNewInference)
	if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerHub contract.
type WorkerHubOwnershipTransferredIterator struct {
	Event *WorkerHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerHub contract.
type WorkerHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubOwnershipTransferredIterator{contract: _WorkerHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubOwnershipTransferred)
				if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerHubOwnershipTransferred, error) {
	event := new(WorkerHubOwnershipTransferred)
	if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerHub contract.
type WorkerHubPausedIterator struct {
	Event *WorkerHubPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubPaused represents a Paused event raised by the WorkerHub contract.
type WorkerHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerHubPausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubPausedIterator{contract: _WorkerHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerHubPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubPaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) ParsePaused(log types.Log) (*WorkerHubPaused, error) {
	event := new(WorkerHubPaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the WorkerHub contract.
type WorkerHubPenaltyDurationUpdatedIterator struct {
	Event *WorkerHubPenaltyDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubPenaltyDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubPenaltyDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the WorkerHub contract.
type WorkerHubPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHub *WorkerHubFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*WorkerHubPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubPenaltyDurationUpdatedIterator{contract: _WorkerHub.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHub *WorkerHubFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubPenaltyDurationUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePenaltyDurationUpdated is a log parse operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHub *WorkerHubFilterer) ParsePenaltyDurationUpdated(log types.Log) (*WorkerHubPenaltyDurationUpdated, error) {
	event := new(WorkerHubPenaltyDurationUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the WorkerHub contract.
type WorkerHubRestakeIterator struct {
	Event *WorkerHubRestake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRestake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubRestake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRestake represents a Restake event raised by the WorkerHub contract.
type WorkerHubRestake struct {
	Miner   common.Address
	Restake *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHub *WorkerHubFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, model []common.Address) (*WorkerHubRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRestakeIterator{contract: _WorkerHub.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHub *WorkerHubFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *WorkerHubRestake, miner []common.Address, model []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRestake)
				if err := _WorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestake is a log parse operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHub *WorkerHubFilterer) ParseRestake(log types.Log) (*WorkerHubRestake, error) {
	event := new(WorkerHubRestake)
	if err := _WorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the WorkerHub contract.
type WorkerHubRewardClaimIterator struct {
	Event *WorkerHubRewardClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRewardClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubRewardClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRewardClaim represents a RewardClaim event raised by the WorkerHub contract.
type WorkerHubRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*WorkerHubRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRewardClaimIterator{contract: _WorkerHub.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *WorkerHubRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRewardClaim)
				if err := _WorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaim is a log parse operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseRewardClaim(log types.Log) (*WorkerHubRewardClaim, error) {
	event := new(WorkerHubRewardClaim)
	if err := _WorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the WorkerHub contract.
type WorkerHubRewardPerEpochIterator struct {
	Event *WorkerHubRewardPerEpoch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRewardPerEpoch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubRewardPerEpoch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRewardPerEpoch represents a RewardPerEpoch event raised by the WorkerHub contract.
type WorkerHubRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHub *WorkerHubFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*WorkerHubRewardPerEpochIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubRewardPerEpochIterator{contract: _WorkerHub.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHub *WorkerHubFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRewardPerEpoch)
				if err := _WorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardPerEpoch is a log parse operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHub *WorkerHubFilterer) ParseRewardPerEpoch(log types.Log) (*WorkerHubRewardPerEpoch, error) {
	event := new(WorkerHubRewardPerEpoch)
	if err := _WorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the WorkerHub contract.
type WorkerHubSolutionSubmissionIterator struct {
	Event *WorkerHubSolutionSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubSolutionSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubSolutionSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubSolutionSubmission represents a SolutionSubmission event raised by the WorkerHub contract.
type WorkerHubSolutionSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubSolutionSubmissionIterator{contract: _WorkerHub.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubSolutionSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubSolutionSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSolutionSubmission is a log parse operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) ParseSolutionSubmission(log types.Log) (*WorkerHubSolutionSubmission, error) {
	event := new(WorkerHubSolutionSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubTopUpInferIterator is returned from FilterTopUpInfer and is used to iterate over the raw logs and unpacked data for TopUpInfer events raised by the WorkerHub contract.
type WorkerHubTopUpInferIterator struct {
	Event *WorkerHubTopUpInfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubTopUpInferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubTopUpInfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubTopUpInfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubTopUpInferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubTopUpInferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubTopUpInfer represents a TopUpInfer event raised by the WorkerHub contract.
type WorkerHubTopUpInfer struct {
	InferenceId *big.Int
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTopUpInfer is a free log retrieval operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterTopUpInfer(opts *bind.FilterOpts, inferenceId []*big.Int, creator []common.Address) (*WorkerHubTopUpInferIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTopUpInferIterator{contract: _WorkerHub.contract, event: "TopUpInfer", logs: logs, sub: sub}, nil
}

// WatchTopUpInfer is a free log subscription operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchTopUpInfer(opts *bind.WatchOpts, sink chan<- *WorkerHubTopUpInfer, inferenceId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubTopUpInfer)
				if err := _WorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTopUpInfer is a log parse operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseTopUpInfer(log types.Log) (*WorkerHubTopUpInfer, error) {
	event := new(WorkerHubTopUpInfer)
	if err := _WorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubTransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the WorkerHub contract.
type WorkerHubTransferFeeIterator struct {
	Event *WorkerHubTransferFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubTransferFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubTransferFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubTransferFee represents a TransferFee event raised by the WorkerHub contract.
type WorkerHubTransferFee struct {
	Miner       common.Address
	MingingFee  *big.Int
	Treasury    common.Address
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHub *WorkerHubFilterer) FilterTransferFee(opts *bind.FilterOpts, miner []common.Address, treasury []common.Address) (*WorkerHubTransferFeeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "TransferFee", minerRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTransferFeeIterator{contract: _WorkerHub.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHub *WorkerHubFilterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *WorkerHubTransferFee, miner []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "TransferFee", minerRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubTransferFee)
				if err := _WorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFee is a log parse operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHub *WorkerHubFilterer) ParseTransferFee(log types.Log) (*WorkerHubTransferFee, error) {
	event := new(WorkerHubTransferFee)
	if err := _WorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerHub contract.
type WorkerHubUnpausedIterator struct {
	Event *WorkerHubUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubUnpaused represents a Unpaused event raised by the WorkerHub contract.
type WorkerHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerHubUnpausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubUnpausedIterator{contract: _WorkerHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubUnpaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) ParseUnpaused(log types.Log) (*WorkerHubUnpaused, error) {
	event := new(WorkerHubUnpaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the WorkerHub contract.
type WorkerHubUnstakeDelayTimeIterator struct {
	Event *WorkerHubUnstakeDelayTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubUnstakeDelayTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubUnstakeDelayTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubUnstakeDelayTime represents a UnstakeDelayTime event raised by the WorkerHub contract.
type WorkerHubUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHub *WorkerHubFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*WorkerHubUnstakeDelayTimeIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &WorkerHubUnstakeDelayTimeIterator{contract: _WorkerHub.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHub *WorkerHubFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *WorkerHubUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubUnstakeDelayTime)
				if err := _WorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeDelayTime is a log parse operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHub *WorkerHubFilterer) ParseUnstakeDelayTime(log types.Log) (*WorkerHubUnstakeDelayTime, error) {
	event := new(WorkerHubUnstakeDelayTime)
	if err := _WorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorExtraStakeIterator is returned from FilterValidatorExtraStake and is used to iterate over the raw logs and unpacked data for ValidatorExtraStake events raised by the WorkerHub contract.
type WorkerHubValidatorExtraStakeIterator struct {
	Event *WorkerHubValidatorExtraStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubValidatorExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorExtraStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubValidatorExtraStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubValidatorExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorExtraStake represents a ValidatorExtraStake event raised by the WorkerHub contract.
type WorkerHubValidatorExtraStake struct {
	Validator common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorExtraStake is a free log retrieval operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorExtraStake(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorExtraStakeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorExtraStakeIterator{contract: _WorkerHub.contract, event: "ValidatorExtraStake", logs: logs, sub: sub}, nil
}

// WatchValidatorExtraStake is a free log subscription operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorExtraStake, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorExtraStake)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorExtraStake is a log parse operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorExtraStake(log types.Log) (*WorkerHubValidatorExtraStake, error) {
	event := new(WorkerHubValidatorExtraStake)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorJoinIterator is returned from FilterValidatorJoin and is used to iterate over the raw logs and unpacked data for ValidatorJoin events raised by the WorkerHub contract.
type WorkerHubValidatorJoinIterator struct {
	Event *WorkerHubValidatorJoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubValidatorJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorJoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubValidatorJoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubValidatorJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorJoin represents a ValidatorJoin event raised by the WorkerHub contract.
type WorkerHubValidatorJoin struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJoin is a free log retrieval operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorJoin(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorJoinIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorJoin", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorJoinIterator{contract: _WorkerHub.contract, event: "ValidatorJoin", logs: logs, sub: sub}, nil
}

// WatchValidatorJoin is a free log subscription operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorJoin(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorJoin, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorJoin", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorJoin)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorJoin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorJoin is a log parse operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorJoin(log types.Log) (*WorkerHubValidatorJoin, error) {
	event := new(WorkerHubValidatorJoin)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the WorkerHub contract.
type WorkerHubValidatorRegistrationIterator struct {
	Event *WorkerHubValidatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubValidatorRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorRegistration represents a ValidatorRegistration event raised by the WorkerHub contract.
type WorkerHubValidatorRegistration struct {
	Validator common.Address
	Tier      uint16
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorRegistration(opts *bind.FilterOpts, validator []common.Address, tier []uint16) (*WorkerHubValidatorRegistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorRegistrationIterator{contract: _WorkerHub.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorRegistration, validator []common.Address, tier []uint16) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistration is a log parse operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorRegistration(log types.Log) (*WorkerHubValidatorRegistration, error) {
	event := new(WorkerHubValidatorRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorUnregistrationIterator is returned from FilterValidatorUnregistration and is used to iterate over the raw logs and unpacked data for ValidatorUnregistration events raised by the WorkerHub contract.
type WorkerHubValidatorUnregistrationIterator struct {
	Event *WorkerHubValidatorUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubValidatorUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubValidatorUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubValidatorUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorUnregistration represents a ValidatorUnregistration event raised by the WorkerHub contract.
type WorkerHubValidatorUnregistration struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnregistration is a free log retrieval operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorUnregistration(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorUnregistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorUnregistrationIterator{contract: _WorkerHub.contract, event: "ValidatorUnregistration", logs: logs, sub: sub}, nil
}

// WatchValidatorUnregistration is a free log subscription operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorUnregistration, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorUnregistration is a log parse operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorUnregistration(log types.Log) (*WorkerHubValidatorUnregistration, error) {
	event := new(WorkerHubValidatorUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorUnstakeIterator is returned from FilterValidatorUnstake and is used to iterate over the raw logs and unpacked data for ValidatorUnstake events raised by the WorkerHub contract.
type WorkerHubValidatorUnstakeIterator struct {
	Event *WorkerHubValidatorUnstake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerHubValidatorUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerHubValidatorUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerHubValidatorUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorUnstake represents a ValidatorUnstake event raised by the WorkerHub contract.
type WorkerHubValidatorUnstake struct {
	Validator common.Address
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnstake is a free log retrieval operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorUnstake(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorUnstakeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorUnstake", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorUnstakeIterator{contract: _WorkerHub.contract, event: "ValidatorUnstake", logs: logs, sub: sub}, nil
}

// WatchValidatorUnstake is a free log subscription operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorUnstake(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorUnstake, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorUnstake", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorUnstake)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorUnstake is a log parse operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorUnstake(log types.Log) (*WorkerHubValidatorUnstake, error) {
	event := new(WorkerHubValidatorUnstake)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
