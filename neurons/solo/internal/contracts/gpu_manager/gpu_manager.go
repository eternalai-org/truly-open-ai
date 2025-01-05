// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gpu_manager

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

// IGPUManagerModel is an auto generated low-level Go binding around an user-defined struct.
type IGPUManagerModel struct {
	MinimumFee *big.Int
	Tier       uint32
}

// IGPUManagerUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type IGPUManagerUnstakeRequest struct {
	Stake    *big.Int
	UnlockAt *big.Int
}

// GpuManagerMetaData contains all meta data concerning the GpuManager contract.
var GpuManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameModelAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinFeeToUseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_modelCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"_models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_promptScheduler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinerUnstakeRequests\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unstakeAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"internalType\":\"structIGPUManager.UnstakeRequest[]\",\"name\":\"unstakeRequests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getModelInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"internalType\":\"structIGPUManager.Model\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wEAIAmt\",\"type\":\"uint256\"}],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"modelCollection_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minerMinimumStake_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksPerEpoch_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerEpoch_\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unstakeDelayTime_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"penaltyDuration_\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"finePercentage_\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minFeeToUse_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newPercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"}],\"name\":\"setMinFeeToUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinerMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPromptScheduler\",\"type\":\"address\"}],\"name\":\"setPromptSchedulerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"delayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakeDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"validateMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"minersRequired\",\"type\":\"uint256\"}],\"name\":\"validateModelAndChooseRandomMiner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GpuManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use GpuManagerMetaData.ABI instead.
var GpuManagerABI = GpuManagerMetaData.ABI

// GpuManager is an auto generated Go binding around an Ethereum contract.
type GpuManager struct {
	GpuManagerCaller     // Read-only binding to the contract
	GpuManagerTransactor // Write-only binding to the contract
	GpuManagerFilterer   // Log filterer for contract events
}

// GpuManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GpuManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GpuManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GpuManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GpuManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GpuManagerSession struct {
	Contract     *GpuManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GpuManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GpuManagerCallerSession struct {
	Contract *GpuManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GpuManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GpuManagerTransactorSession struct {
	Contract     *GpuManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GpuManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GpuManagerRaw struct {
	Contract *GpuManager // Generic contract binding to access the raw methods on
}

// GpuManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GpuManagerCallerRaw struct {
	Contract *GpuManagerCaller // Generic read-only contract binding to access the raw methods on
}

// GpuManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GpuManagerTransactorRaw struct {
	Contract *GpuManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGpuManager creates a new instance of GpuManager, bound to a specific deployed contract.
func NewGpuManager(address common.Address, backend bind.ContractBackend) (*GpuManager, error) {
	contract, err := bindGpuManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GpuManager{GpuManagerCaller: GpuManagerCaller{contract: contract}, GpuManagerTransactor: GpuManagerTransactor{contract: contract}, GpuManagerFilterer: GpuManagerFilterer{contract: contract}}, nil
}

// NewGpuManagerCaller creates a new read-only instance of GpuManager, bound to a specific deployed contract.
func NewGpuManagerCaller(address common.Address, caller bind.ContractCaller) (*GpuManagerCaller, error) {
	contract, err := bindGpuManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GpuManagerCaller{contract: contract}, nil
}

// NewGpuManagerTransactor creates a new write-only instance of GpuManager, bound to a specific deployed contract.
func NewGpuManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GpuManagerTransactor, error) {
	contract, err := bindGpuManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GpuManagerTransactor{contract: contract}, nil
}

// NewGpuManagerFilterer creates a new log filterer instance of GpuManager, bound to a specific deployed contract.
func NewGpuManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GpuManagerFilterer, error) {
	contract, err := bindGpuManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GpuManagerFilterer{contract: contract}, nil
}

// bindGpuManager binds a generic wrapper to an already deployed contract.
func bindGpuManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GpuManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GpuManager *GpuManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GpuManager.Contract.GpuManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GpuManager *GpuManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.Contract.GpuManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GpuManager *GpuManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GpuManager.Contract.GpuManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GpuManager *GpuManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GpuManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GpuManager *GpuManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GpuManager *GpuManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GpuManager.Contract.contract.Transact(opts, method, params...)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerSession) BlocksPerEpoch() (*big.Int, error) {
	return _GpuManager.Contract.BlocksPerEpoch(&_GpuManager.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _GpuManager.Contract.BlocksPerEpoch(&_GpuManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GpuManager *GpuManagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GpuManager *GpuManagerSession) CurrentEpoch() (*big.Int, error) {
	return _GpuManager.Contract.CurrentEpoch(&_GpuManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GpuManager *GpuManagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _GpuManager.Contract.CurrentEpoch(&_GpuManager.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GpuManager *GpuManagerCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GpuManager *GpuManagerSession) FinePercentage() (uint16, error) {
	return _GpuManager.Contract.FinePercentage(&_GpuManager.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GpuManager *GpuManagerCallerSession) FinePercentage() (uint16, error) {
	return _GpuManager.Contract.FinePercentage(&_GpuManager.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GpuManager *GpuManagerCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GpuManager *GpuManagerSession) LastBlock() (*big.Int, error) {
	return _GpuManager.Contract.LastBlock(&_GpuManager.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) LastBlock() (*big.Int, error) {
	return _GpuManager.Contract.LastBlock(&_GpuManager.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GpuManager *GpuManagerCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GpuManager *GpuManagerSession) MaximumTier() (uint16, error) {
	return _GpuManager.Contract.MaximumTier(&_GpuManager.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GpuManager *GpuManagerCallerSession) MaximumTier() (uint16, error) {
	return _GpuManager.Contract.MaximumTier(&_GpuManager.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GpuManager *GpuManagerCaller) MinFeeToUse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_minFeeToUse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GpuManager *GpuManagerSession) MinFeeToUse() (*big.Int, error) {
	return _GpuManager.Contract.MinFeeToUse(&_GpuManager.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) MinFeeToUse() (*big.Int, error) {
	return _GpuManager.Contract.MinFeeToUse(&_GpuManager.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GpuManager *GpuManagerCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GpuManager *GpuManagerSession) MinerMinimumStake() (*big.Int, error) {
	return _GpuManager.Contract.MinerMinimumStake(&_GpuManager.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _GpuManager.Contract.MinerMinimumStake(&_GpuManager.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x54eb2d2a.
//
// Solidity: function _minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_GpuManager *GpuManagerCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_minerUnstakeRequests", arg0)

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

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x54eb2d2a.
//
// Solidity: function _minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_GpuManager *GpuManagerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _GpuManager.Contract.MinerUnstakeRequests(&_GpuManager.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x54eb2d2a.
//
// Solidity: function _minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_GpuManager *GpuManagerCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _GpuManager.Contract.MinerUnstakeRequests(&_GpuManager.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xb1a976ef.
//
// Solidity: function _miners(address ) view returns(uint256 stake, uint32 modelId, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_GpuManager *GpuManagerCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_miners", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		ModelId          uint32
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ModelId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0xb1a976ef.
//
// Solidity: function _miners(address ) view returns(uint256 stake, uint32 modelId, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_GpuManager *GpuManagerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _GpuManager.Contract.Miners(&_GpuManager.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xb1a976ef.
//
// Solidity: function _miners(address ) view returns(uint256 stake, uint32 modelId, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_GpuManager *GpuManagerCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _GpuManager.Contract.Miners(&_GpuManager.CallOpts, arg0)
}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GpuManager *GpuManagerCaller) ModelCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_modelCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GpuManager *GpuManagerSession) ModelCollection() (common.Address, error) {
	return _GpuManager.Contract.ModelCollection(&_GpuManager.CallOpts)
}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GpuManager *GpuManagerCallerSession) ModelCollection() (common.Address, error) {
	return _GpuManager.Contract.ModelCollection(&_GpuManager.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x55f89085.
//
// Solidity: function _models(uint32 ) view returns(uint256 minimumFee, uint32 tier)
func (_GpuManager *GpuManagerCaller) Models(opts *bind.CallOpts, arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_models", arg0)

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

// Models is a free data retrieval call binding the contract method 0x55f89085.
//
// Solidity: function _models(uint32 ) view returns(uint256 minimumFee, uint32 tier)
func (_GpuManager *GpuManagerSession) Models(arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _GpuManager.Contract.Models(&_GpuManager.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x55f89085.
//
// Solidity: function _models(uint32 ) view returns(uint256 minimumFee, uint32 tier)
func (_GpuManager *GpuManagerCallerSession) Models(arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _GpuManager.Contract.Models(&_GpuManager.CallOpts, arg0)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GpuManager *GpuManagerCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GpuManager *GpuManagerSession) PenaltyDuration() (*big.Int, error) {
	return _GpuManager.Contract.PenaltyDuration(&_GpuManager.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GpuManager *GpuManagerCallerSession) PenaltyDuration() (*big.Int, error) {
	return _GpuManager.Contract.PenaltyDuration(&_GpuManager.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GpuManager *GpuManagerCaller) PromptScheduler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_promptScheduler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GpuManager *GpuManagerSession) PromptScheduler() (common.Address, error) {
	return _GpuManager.Contract.PromptScheduler(&_GpuManager.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GpuManager *GpuManagerCallerSession) PromptScheduler() (common.Address, error) {
	return _GpuManager.Contract.PromptScheduler(&_GpuManager.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0xa662f84d.
//
// Solidity: function _rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_GpuManager *GpuManagerCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_rewardInEpoch", arg0)

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

// RewardInEpoch is a free data retrieval call binding the contract method 0xa662f84d.
//
// Solidity: function _rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_GpuManager *GpuManagerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _GpuManager.Contract.RewardInEpoch(&_GpuManager.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0xa662f84d.
//
// Solidity: function _rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_GpuManager *GpuManagerCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _GpuManager.Contract.RewardInEpoch(&_GpuManager.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerSession) RewardPerEpoch() (*big.Int, error) {
	return _GpuManager.Contract.RewardPerEpoch(&_GpuManager.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _GpuManager.Contract.RewardPerEpoch(&_GpuManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GpuManager *GpuManagerCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GpuManager *GpuManagerSession) Treasury() (common.Address, error) {
	return _GpuManager.Contract.Treasury(&_GpuManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GpuManager *GpuManagerCallerSession) Treasury() (common.Address, error) {
	return _GpuManager.Contract.Treasury(&_GpuManager.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GpuManager *GpuManagerCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GpuManager *GpuManagerSession) UnstakeDelayTime() (*big.Int, error) {
	return _GpuManager.Contract.UnstakeDelayTime(&_GpuManager.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GpuManager *GpuManagerCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _GpuManager.Contract.UnstakeDelayTime(&_GpuManager.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GpuManager *GpuManagerCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GpuManager *GpuManagerSession) WEAIToken() (common.Address, error) {
	return _GpuManager.Contract.WEAIToken(&_GpuManager.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GpuManager *GpuManagerCallerSession) WEAIToken() (common.Address, error) {
	return _GpuManager.Contract.WEAIToken(&_GpuManager.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_GpuManager *GpuManagerCaller) GetAllMinerUnstakeRequests(opts *bind.CallOpts) (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getAllMinerUnstakeRequests")

	outstruct := new(struct {
		UnstakeAddresses []common.Address
		UnstakeRequests  []IGPUManagerUnstakeRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnstakeAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UnstakeRequests = *abi.ConvertType(out[1], new([]IGPUManagerUnstakeRequest)).(*[]IGPUManagerUnstakeRequest)

	return *outstruct, err

}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_GpuManager *GpuManagerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	return _GpuManager.Contract.GetAllMinerUnstakeRequests(&_GpuManager.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_GpuManager *GpuManagerCallerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	return _GpuManager.Contract.GetAllMinerUnstakeRequests(&_GpuManager.CallOpts)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GpuManager *GpuManagerCaller) GetMinFeeToUse(opts *bind.CallOpts, modelId uint32) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getMinFeeToUse", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GpuManager *GpuManagerSession) GetMinFeeToUse(modelId uint32) (*big.Int, error) {
	return _GpuManager.Contract.GetMinFeeToUse(&_GpuManager.CallOpts, modelId)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) GetMinFeeToUse(modelId uint32) (*big.Int, error) {
	return _GpuManager.Contract.GetMinFeeToUse(&_GpuManager.CallOpts, modelId)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GpuManager *GpuManagerCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GpuManager *GpuManagerSession) GetMinerAddresses() ([]common.Address, error) {
	return _GpuManager.Contract.GetMinerAddresses(&_GpuManager.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GpuManager *GpuManagerCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _GpuManager.Contract.GetMinerAddresses(&_GpuManager.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GpuManager *GpuManagerCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, modelId uint32) ([]common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getMinerAddressesOfModel", modelId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GpuManager *GpuManagerSession) GetMinerAddressesOfModel(modelId uint32) ([]common.Address, error) {
	return _GpuManager.Contract.GetMinerAddressesOfModel(&_GpuManager.CallOpts, modelId)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GpuManager *GpuManagerCallerSession) GetMinerAddressesOfModel(modelId uint32) ([]common.Address, error) {
	return _GpuManager.Contract.GetMinerAddressesOfModel(&_GpuManager.CallOpts, modelId)
}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GpuManager *GpuManagerCaller) GetModelIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getModelIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GpuManager *GpuManagerSession) GetModelIds() ([]*big.Int, error) {
	return _GpuManager.Contract.GetModelIds(&_GpuManager.CallOpts)
}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GpuManager *GpuManagerCallerSession) GetModelIds() ([]*big.Int, error) {
	return _GpuManager.Contract.GetModelIds(&_GpuManager.CallOpts)
}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GpuManager *GpuManagerCaller) GetModelInfo(opts *bind.CallOpts, modelId uint32) (IGPUManagerModel, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getModelInfo", modelId)

	if err != nil {
		return *new(IGPUManagerModel), err
	}

	out0 := *abi.ConvertType(out[0], new(IGPUManagerModel)).(*IGPUManagerModel)

	return out0, err

}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GpuManager *GpuManagerSession) GetModelInfo(modelId uint32) (IGPUManagerModel, error) {
	return _GpuManager.Contract.GetModelInfo(&_GpuManager.CallOpts, modelId)
}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GpuManager *GpuManagerCallerSession) GetModelInfo(modelId uint32) (IGPUManagerModel, error) {
	return _GpuManager.Contract.GetModelInfo(&_GpuManager.CallOpts, modelId)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GpuManager *GpuManagerCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GpuManager *GpuManagerSession) GetNOMiner() (*big.Int, error) {
	return _GpuManager.Contract.GetNOMiner(&_GpuManager.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) GetNOMiner() (*big.Int, error) {
	return _GpuManager.Contract.GetNOMiner(&_GpuManager.CallOpts)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GpuManager *GpuManagerCaller) Multiplier(opts *bind.CallOpts, miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "multiplier", miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GpuManager *GpuManagerSession) Multiplier(miner common.Address) (*big.Int, error) {
	return _GpuManager.Contract.Multiplier(&_GpuManager.CallOpts, miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GpuManager *GpuManagerCallerSession) Multiplier(miner common.Address) (*big.Int, error) {
	return _GpuManager.Contract.Multiplier(&_GpuManager.CallOpts, miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GpuManager *GpuManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GpuManager *GpuManagerSession) Owner() (common.Address, error) {
	return _GpuManager.Contract.Owner(&_GpuManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GpuManager *GpuManagerCallerSession) Owner() (common.Address, error) {
	return _GpuManager.Contract.Owner(&_GpuManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GpuManager *GpuManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GpuManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GpuManager *GpuManagerSession) Paused() (bool, error) {
	return _GpuManager.Contract.Paused(&_GpuManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GpuManager *GpuManagerCallerSession) Paused() (bool, error) {
	return _GpuManager.Contract.Paused(&_GpuManager.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GpuManager *GpuManagerTransactor) ClaimReward(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "claimReward", miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GpuManager *GpuManagerSession) ClaimReward(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.ClaimReward(&_GpuManager.TransactOpts, miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GpuManager *GpuManagerTransactorSession) ClaimReward(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.ClaimReward(&_GpuManager.TransactOpts, miner)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "forceChangeModelForMiner", miner, modelId)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GpuManager *GpuManagerSession) ForceChangeModelForMiner(miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.ForceChangeModelForMiner(&_GpuManager.TransactOpts, miner, modelId)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactorSession) ForceChangeModelForMiner(miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.ForceChangeModelForMiner(&_GpuManager.TransactOpts, miner, modelId)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GpuManager *GpuManagerTransactor) IncreaseMinerStake(opts *bind.TransactOpts, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "increaseMinerStake", wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GpuManager *GpuManagerSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.IncreaseMinerStake(&_GpuManager.TransactOpts, wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GpuManager *GpuManagerTransactorSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.IncreaseMinerStake(&_GpuManager.TransactOpts, wEAIAmt)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GpuManager *GpuManagerTransactor) Initialize(opts *bind.TransactOpts, wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "initialize", wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GpuManager *GpuManagerSession) Initialize(wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.Initialize(&_GpuManager.TransactOpts, wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GpuManager *GpuManagerTransactorSession) Initialize(wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.Initialize(&_GpuManager.TransactOpts, wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GpuManager *GpuManagerTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GpuManager *GpuManagerSession) JoinForMinting() (*types.Transaction, error) {
	return _GpuManager.Contract.JoinForMinting(&_GpuManager.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GpuManager *GpuManagerTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _GpuManager.Contract.JoinForMinting(&_GpuManager.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "registerMiner", tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterMiner(&_GpuManager.TransactOpts, tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerTransactorSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterMiner(&_GpuManager.TransactOpts, tier)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactor) RegisterMiner0(opts *bind.TransactOpts, tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "registerMiner0", tier, modelId)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GpuManager *GpuManagerSession) RegisterMiner0(tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterMiner0(&_GpuManager.TransactOpts, tier, modelId)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactorSession) RegisterMiner0(tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterMiner0(&_GpuManager.TransactOpts, tier, modelId)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerTransactor) RegisterModel(opts *bind.TransactOpts, modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "registerModel", modelId, tier, minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerSession) RegisterModel(modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterModel(&_GpuManager.TransactOpts, modelId, tier, minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerTransactorSession) RegisterModel(modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.RegisterModel(&_GpuManager.TransactOpts, modelId, tier, minimumFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GpuManager *GpuManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GpuManager *GpuManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _GpuManager.Contract.RenounceOwnership(&_GpuManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GpuManager *GpuManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GpuManager.Contract.RenounceOwnership(&_GpuManager.TransactOpts)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.RestakeForMiner(&_GpuManager.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GpuManager *GpuManagerTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.RestakeForMiner(&_GpuManager.TransactOpts, tier)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GpuManager *GpuManagerTransactor) RewardToClaim(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "rewardToClaim", miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GpuManager *GpuManagerSession) RewardToClaim(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.RewardToClaim(&_GpuManager.TransactOpts, miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GpuManager *GpuManagerTransactorSession) RewardToClaim(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.RewardToClaim(&_GpuManager.TransactOpts, miner)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GpuManager *GpuManagerTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, blocks *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setBlocksPerEpoch", blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GpuManager *GpuManagerSession) SetBlocksPerEpoch(blocks *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetBlocksPerEpoch(&_GpuManager.TransactOpts, blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GpuManager *GpuManagerTransactorSession) SetBlocksPerEpoch(blocks *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetBlocksPerEpoch(&_GpuManager.TransactOpts, blocks)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GpuManager *GpuManagerTransactor) SetFinePercentage(opts *bind.TransactOpts, newPercentage uint16) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setFinePercentage", newPercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GpuManager *GpuManagerSession) SetFinePercentage(newPercentage uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.SetFinePercentage(&_GpuManager.TransactOpts, newPercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GpuManager *GpuManagerTransactorSession) SetFinePercentage(newPercentage uint16) (*types.Transaction, error) {
	return _GpuManager.Contract.SetFinePercentage(&_GpuManager.TransactOpts, newPercentage)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GpuManager *GpuManagerTransactor) SetMinFeeToUse(opts *bind.TransactOpts, minFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setMinFeeToUse", minFee)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GpuManager *GpuManagerSession) SetMinFeeToUse(minFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetMinFeeToUse(&_GpuManager.TransactOpts, minFee)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GpuManager *GpuManagerTransactorSession) SetMinFeeToUse(minFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetMinFeeToUse(&_GpuManager.TransactOpts, minFee)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GpuManager *GpuManagerTransactor) SetMinerMinimumStake(opts *bind.TransactOpts, _minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setMinerMinimumStake", _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GpuManager *GpuManagerSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetMinerMinimumStake(&_GpuManager.TransactOpts, _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GpuManager *GpuManagerTransactorSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetMinerMinimumStake(&_GpuManager.TransactOpts, _minerMinimumStake)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GpuManager *GpuManagerTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setNewRewardInEpoch", newReward)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GpuManager *GpuManagerSession) SetNewRewardInEpoch(newReward *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetNewRewardInEpoch(&_GpuManager.TransactOpts, newReward)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GpuManager *GpuManagerTransactorSession) SetNewRewardInEpoch(newReward *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetNewRewardInEpoch(&_GpuManager.TransactOpts, newReward)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GpuManager *GpuManagerTransactor) SetPenaltyDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setPenaltyDuration", duration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GpuManager *GpuManagerSession) SetPenaltyDuration(duration *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetPenaltyDuration(&_GpuManager.TransactOpts, duration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GpuManager *GpuManagerTransactorSession) SetPenaltyDuration(duration *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetPenaltyDuration(&_GpuManager.TransactOpts, duration)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GpuManager *GpuManagerTransactor) SetPromptSchedulerAddress(opts *bind.TransactOpts, newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setPromptSchedulerAddress", newPromptScheduler)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GpuManager *GpuManagerSession) SetPromptSchedulerAddress(newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.SetPromptSchedulerAddress(&_GpuManager.TransactOpts, newPromptScheduler)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GpuManager *GpuManagerTransactorSession) SetPromptSchedulerAddress(newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.SetPromptSchedulerAddress(&_GpuManager.TransactOpts, newPromptScheduler)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GpuManager *GpuManagerTransactor) SetUnstakeDelayTime(opts *bind.TransactOpts, delayTime *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setUnstakeDelayTime", delayTime)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GpuManager *GpuManagerSession) SetUnstakeDelayTime(delayTime *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetUnstakeDelayTime(&_GpuManager.TransactOpts, delayTime)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GpuManager *GpuManagerTransactorSession) SetUnstakeDelayTime(delayTime *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.SetUnstakeDelayTime(&_GpuManager.TransactOpts, delayTime)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GpuManager *GpuManagerTransactor) SetWEAIAddress(opts *bind.TransactOpts, wEAIToken common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "setWEAIAddress", wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GpuManager *GpuManagerSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.SetWEAIAddress(&_GpuManager.TransactOpts, wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GpuManager *GpuManagerTransactorSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.SetWEAIAddress(&_GpuManager.TransactOpts, wEAIToken)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GpuManager *GpuManagerTransactor) SlashMiner(opts *bind.TransactOpts, miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "slashMiner", miner, isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GpuManager *GpuManagerSession) SlashMiner(miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GpuManager.Contract.SlashMiner(&_GpuManager.TransactOpts, miner, isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GpuManager *GpuManagerTransactorSession) SlashMiner(miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GpuManager.Contract.SlashMiner(&_GpuManager.TransactOpts, miner, isFined)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GpuManager *GpuManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GpuManager *GpuManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.TransferOwnership(&_GpuManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GpuManager *GpuManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.TransferOwnership(&_GpuManager.TransactOpts, newOwner)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GpuManager *GpuManagerTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GpuManager *GpuManagerSession) UnregisterMiner() (*types.Transaction, error) {
	return _GpuManager.Contract.UnregisterMiner(&_GpuManager.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GpuManager *GpuManagerTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _GpuManager.Contract.UnregisterMiner(&_GpuManager.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactor) UnregisterModel(opts *bind.TransactOpts, modelId uint32) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "unregisterModel", modelId)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GpuManager *GpuManagerSession) UnregisterModel(modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.UnregisterModel(&_GpuManager.TransactOpts, modelId)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GpuManager *GpuManagerTransactorSession) UnregisterModel(modelId uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.UnregisterModel(&_GpuManager.TransactOpts, modelId)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GpuManager *GpuManagerTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GpuManager *GpuManagerSession) UnstakeForMiner() (*types.Transaction, error) {
	return _GpuManager.Contract.UnstakeForMiner(&_GpuManager.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GpuManager *GpuManagerTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _GpuManager.Contract.UnstakeForMiner(&_GpuManager.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GpuManager *GpuManagerTransactor) UpdateEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "updateEpoch")
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GpuManager *GpuManagerSession) UpdateEpoch() (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateEpoch(&_GpuManager.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GpuManager *GpuManagerTransactorSession) UpdateEpoch() (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateEpoch(&_GpuManager.TransactOpts)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "updateModelMinimumFee", modelId, minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerSession) UpdateModelMinimumFee(modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateModelMinimumFee(&_GpuManager.TransactOpts, modelId, minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GpuManager *GpuManagerTransactorSession) UpdateModelMinimumFee(modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateModelMinimumFee(&_GpuManager.TransactOpts, modelId, minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GpuManager *GpuManagerTransactor) UpdateModelTier(opts *bind.TransactOpts, modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "updateModelTier", modelId, tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GpuManager *GpuManagerSession) UpdateModelTier(modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateModelTier(&_GpuManager.TransactOpts, modelId, tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GpuManager *GpuManagerTransactorSession) UpdateModelTier(modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GpuManager.Contract.UpdateModelTier(&_GpuManager.TransactOpts, modelId, tier)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GpuManager *GpuManagerTransactor) ValidateMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "validateMiner", miner)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GpuManager *GpuManagerSession) ValidateMiner(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.ValidateMiner(&_GpuManager.TransactOpts, miner)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GpuManager *GpuManagerTransactorSession) ValidateMiner(miner common.Address) (*types.Transaction, error) {
	return _GpuManager.Contract.ValidateMiner(&_GpuManager.TransactOpts, miner)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GpuManager *GpuManagerTransactor) ValidateModelAndChooseRandomMiner(opts *bind.TransactOpts, modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GpuManager.contract.Transact(opts, "validateModelAndChooseRandomMiner", modelId, minersRequired)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GpuManager *GpuManagerSession) ValidateModelAndChooseRandomMiner(modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.ValidateModelAndChooseRandomMiner(&_GpuManager.TransactOpts, modelId, minersRequired)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GpuManager *GpuManagerTransactorSession) ValidateModelAndChooseRandomMiner(modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GpuManager.Contract.ValidateModelAndChooseRandomMiner(&_GpuManager.TransactOpts, modelId, minersRequired)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GpuManager *GpuManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GpuManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GpuManager *GpuManagerSession) Receive() (*types.Transaction, error) {
	return _GpuManager.Contract.Receive(&_GpuManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GpuManager *GpuManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _GpuManager.Contract.Receive(&_GpuManager.TransactOpts)
}

// GpuManagerBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the GpuManager contract.
type GpuManagerBlocksPerEpochIterator struct {
	Event *GpuManagerBlocksPerEpoch // Event containing the contract specifics and raw log

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
func (it *GpuManagerBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerBlocksPerEpoch)
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
		it.Event = new(GpuManagerBlocksPerEpoch)
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
func (it *GpuManagerBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerBlocksPerEpoch represents a BlocksPerEpoch event raised by the GpuManager contract.
type GpuManagerBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_GpuManager *GpuManagerFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*GpuManagerBlocksPerEpochIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &GpuManagerBlocksPerEpochIterator{contract: _GpuManager.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_GpuManager *GpuManagerFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *GpuManagerBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerBlocksPerEpoch)
				if err := _GpuManager.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseBlocksPerEpoch(log types.Log) (*GpuManagerBlocksPerEpoch, error) {
	event := new(GpuManagerBlocksPerEpoch)
	if err := _GpuManager.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the GpuManager contract.
type GpuManagerFinePercentageUpdatedIterator struct {
	Event *GpuManagerFinePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *GpuManagerFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerFinePercentageUpdated)
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
		it.Event = new(GpuManagerFinePercentageUpdated)
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
func (it *GpuManagerFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerFinePercentageUpdated represents a FinePercentageUpdated event raised by the GpuManager contract.
type GpuManagerFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_GpuManager *GpuManagerFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*GpuManagerFinePercentageUpdatedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &GpuManagerFinePercentageUpdatedIterator{contract: _GpuManager.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_GpuManager *GpuManagerFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *GpuManagerFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerFinePercentageUpdated)
				if err := _GpuManager.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseFinePercentageUpdated(log types.Log) (*GpuManagerFinePercentageUpdated, error) {
	event := new(GpuManagerFinePercentageUpdated)
	if err := _GpuManager.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the GpuManager contract.
type GpuManagerFraudulentMinerPenalizedIterator struct {
	Event *GpuManagerFraudulentMinerPenalized // Event containing the contract specifics and raw log

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
func (it *GpuManagerFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerFraudulentMinerPenalized)
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
		it.Event = new(GpuManagerFraudulentMinerPenalized)
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
func (it *GpuManagerFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the GpuManager contract.
type GpuManagerFraudulentMinerPenalized struct {
	Miner    common.Address
	ModelId  uint32
	Treasury common.Address
	Fine     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x396ee931f435c63405d255f5e0d31a0d1a1f6b57d59ef9559155464a15b13593.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, uint32 indexed modelId, address indexed treasury, uint256 fine)
func (_GpuManager *GpuManagerFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelId []uint32, treasury []common.Address) (*GpuManagerFraudulentMinerPenalizedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelIdRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerFraudulentMinerPenalizedIterator{contract: _GpuManager.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x396ee931f435c63405d255f5e0d31a0d1a1f6b57d59ef9559155464a15b13593.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, uint32 indexed modelId, address indexed treasury, uint256 fine)
func (_GpuManager *GpuManagerFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *GpuManagerFraudulentMinerPenalized, miner []common.Address, modelId []uint32, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelIdRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerFraudulentMinerPenalized)
				if err := _GpuManager.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
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

// ParseFraudulentMinerPenalized is a log parse operation binding the contract event 0x396ee931f435c63405d255f5e0d31a0d1a1f6b57d59ef9559155464a15b13593.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, uint32 indexed modelId, address indexed treasury, uint256 fine)
func (_GpuManager *GpuManagerFilterer) ParseFraudulentMinerPenalized(log types.Log) (*GpuManagerFraudulentMinerPenalized, error) {
	event := new(GpuManagerFraudulentMinerPenalized)
	if err := _GpuManager.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GpuManager contract.
type GpuManagerInitializedIterator struct {
	Event *GpuManagerInitialized // Event containing the contract specifics and raw log

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
func (it *GpuManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerInitialized)
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
		it.Event = new(GpuManagerInitialized)
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
func (it *GpuManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerInitialized represents a Initialized event raised by the GpuManager contract.
type GpuManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GpuManager *GpuManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*GpuManagerInitializedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GpuManagerInitializedIterator{contract: _GpuManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GpuManager *GpuManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GpuManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerInitialized)
				if err := _GpuManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseInitialized(log types.Log) (*GpuManagerInitialized, error) {
	event := new(GpuManagerInitialized)
	if err := _GpuManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinFeeToUseUpdatedIterator is returned from FilterMinFeeToUseUpdated and is used to iterate over the raw logs and unpacked data for MinFeeToUseUpdated events raised by the GpuManager contract.
type GpuManagerMinFeeToUseUpdatedIterator struct {
	Event *GpuManagerMinFeeToUseUpdated // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinFeeToUseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinFeeToUseUpdated)
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
		it.Event = new(GpuManagerMinFeeToUseUpdated)
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
func (it *GpuManagerMinFeeToUseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinFeeToUseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinFeeToUseUpdated represents a MinFeeToUseUpdated event raised by the GpuManager contract.
type GpuManagerMinFeeToUseUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeToUseUpdated is a free log retrieval operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_GpuManager *GpuManagerFilterer) FilterMinFeeToUseUpdated(opts *bind.FilterOpts) (*GpuManagerMinFeeToUseUpdatedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinFeeToUseUpdatedIterator{contract: _GpuManager.contract, event: "MinFeeToUseUpdated", logs: logs, sub: sub}, nil
}

// WatchMinFeeToUseUpdated is a free log subscription operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_GpuManager *GpuManagerFilterer) WatchMinFeeToUseUpdated(opts *bind.WatchOpts, sink chan<- *GpuManagerMinFeeToUseUpdated) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinFeeToUseUpdated)
				if err := _GpuManager.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
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

// ParseMinFeeToUseUpdated is a log parse operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_GpuManager *GpuManagerFilterer) ParseMinFeeToUseUpdated(log types.Log) (*GpuManagerMinFeeToUseUpdated, error) {
	event := new(GpuManagerMinFeeToUseUpdated)
	if err := _GpuManager.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the GpuManager contract.
type GpuManagerMinerDeactivatedIterator struct {
	Event *GpuManagerMinerDeactivated // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerDeactivated)
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
		it.Event = new(GpuManagerMinerDeactivated)
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
func (it *GpuManagerMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerDeactivated represents a MinerDeactivated event raised by the GpuManager contract.
type GpuManagerMinerDeactivated struct {
	Miner      common.Address
	ModelId    uint32
	ActiveTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x6e4a7233a3b583018e3a3d018e76ad619bab8ad6e8fe05e12cb83ec1fa75d85e.
//
// Solidity: event MinerDeactivated(address indexed miner, uint32 indexed modelId, uint40 activeTime)
func (_GpuManager *GpuManagerFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelId []uint32) (*GpuManagerMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerDeactivatedIterator{contract: _GpuManager.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x6e4a7233a3b583018e3a3d018e76ad619bab8ad6e8fe05e12cb83ec1fa75d85e.
//
// Solidity: event MinerDeactivated(address indexed miner, uint32 indexed modelId, uint40 activeTime)
func (_GpuManager *GpuManagerFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerDeactivated, miner []common.Address, modelId []uint32) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerDeactivated)
				if err := _GpuManager.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
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

// ParseMinerDeactivated is a log parse operation binding the contract event 0x6e4a7233a3b583018e3a3d018e76ad619bab8ad6e8fe05e12cb83ec1fa75d85e.
//
// Solidity: event MinerDeactivated(address indexed miner, uint32 indexed modelId, uint40 activeTime)
func (_GpuManager *GpuManagerFilterer) ParseMinerDeactivated(log types.Log) (*GpuManagerMinerDeactivated, error) {
	event := new(GpuManagerMinerDeactivated)
	if err := _GpuManager.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the GpuManager contract.
type GpuManagerMinerExtraStakeIterator struct {
	Event *GpuManagerMinerExtraStake // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerExtraStake)
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
		it.Event = new(GpuManagerMinerExtraStake)
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
func (it *GpuManagerMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerExtraStake represents a MinerExtraStake event raised by the GpuManager contract.
type GpuManagerMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_GpuManager *GpuManagerFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*GpuManagerMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerExtraStakeIterator{contract: _GpuManager.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_GpuManager *GpuManagerFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerExtraStake)
				if err := _GpuManager.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseMinerExtraStake(log types.Log) (*GpuManagerMinerExtraStake, error) {
	event := new(GpuManagerMinerExtraStake)
	if err := _GpuManager.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the GpuManager contract.
type GpuManagerMinerJoinIterator struct {
	Event *GpuManagerMinerJoin // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerJoin)
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
		it.Event = new(GpuManagerMinerJoin)
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
func (it *GpuManagerMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerJoin represents a MinerJoin event raised by the GpuManager contract.
type GpuManagerMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_GpuManager *GpuManagerFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*GpuManagerMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerJoinIterator{contract: _GpuManager.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_GpuManager *GpuManagerFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerJoin)
				if err := _GpuManager.contract.UnpackLog(event, "MinerJoin", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseMinerJoin(log types.Log) (*GpuManagerMinerJoin, error) {
	event := new(GpuManagerMinerJoin)
	if err := _GpuManager.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the GpuManager contract.
type GpuManagerMinerRegistrationIterator struct {
	Event *GpuManagerMinerRegistration // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerRegistration)
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
		it.Event = new(GpuManagerMinerRegistration)
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
func (it *GpuManagerMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerRegistration represents a MinerRegistration event raised by the GpuManager contract.
type GpuManagerMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_GpuManager *GpuManagerFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*GpuManagerMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerRegistrationIterator{contract: _GpuManager.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_GpuManager *GpuManagerFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerRegistration)
				if err := _GpuManager.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseMinerRegistration(log types.Log) (*GpuManagerMinerRegistration, error) {
	event := new(GpuManagerMinerRegistration)
	if err := _GpuManager.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the GpuManager contract.
type GpuManagerMinerUnregistrationIterator struct {
	Event *GpuManagerMinerUnregistration // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerUnregistration)
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
		it.Event = new(GpuManagerMinerUnregistration)
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
func (it *GpuManagerMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerUnregistration represents a MinerUnregistration event raised by the GpuManager contract.
type GpuManagerMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_GpuManager *GpuManagerFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*GpuManagerMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerUnregistrationIterator{contract: _GpuManager.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_GpuManager *GpuManagerFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerUnregistration)
				if err := _GpuManager.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseMinerUnregistration(log types.Log) (*GpuManagerMinerUnregistration, error) {
	event := new(GpuManagerMinerUnregistration)
	if err := _GpuManager.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the GpuManager contract.
type GpuManagerMinerUnstakeIterator struct {
	Event *GpuManagerMinerUnstake // Event containing the contract specifics and raw log

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
func (it *GpuManagerMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerMinerUnstake)
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
		it.Event = new(GpuManagerMinerUnstake)
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
func (it *GpuManagerMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerMinerUnstake represents a MinerUnstake event raised by the GpuManager contract.
type GpuManagerMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_GpuManager *GpuManagerFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*GpuManagerMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerMinerUnstakeIterator{contract: _GpuManager.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_GpuManager *GpuManagerFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *GpuManagerMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerMinerUnstake)
				if err := _GpuManager.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseMinerUnstake(log types.Log) (*GpuManagerMinerUnstake, error) {
	event := new(GpuManagerMinerUnstake)
	if err := _GpuManager.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the GpuManager contract.
type GpuManagerModelMinimumFeeUpdateIterator struct {
	Event *GpuManagerModelMinimumFeeUpdate // Event containing the contract specifics and raw log

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
func (it *GpuManagerModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerModelMinimumFeeUpdate)
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
		it.Event = new(GpuManagerModelMinimumFeeUpdate)
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
func (it *GpuManagerModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the GpuManager contract.
type GpuManagerModelMinimumFeeUpdate struct {
	ModelId    uint32
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x32fdbd4cff3135e1bb0ae98bb593ee0c78a48a5e92e80ccf8a8ab6e72b21ffb9.
//
// Solidity: event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, modelId []uint32) (*GpuManagerModelMinimumFeeUpdateIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerModelMinimumFeeUpdateIterator{contract: _GpuManager.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x32fdbd4cff3135e1bb0ae98bb593ee0c78a48a5e92e80ccf8a8ab6e72b21ffb9.
//
// Solidity: event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *GpuManagerModelMinimumFeeUpdate, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerModelMinimumFeeUpdate)
				if err := _GpuManager.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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

// ParseModelMinimumFeeUpdate is a log parse operation binding the contract event 0x32fdbd4cff3135e1bb0ae98bb593ee0c78a48a5e92e80ccf8a8ab6e72b21ffb9.
//
// Solidity: event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*GpuManagerModelMinimumFeeUpdate, error) {
	event := new(GpuManagerModelMinimumFeeUpdate)
	if err := _GpuManager.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the GpuManager contract.
type GpuManagerModelRegistrationIterator struct {
	Event *GpuManagerModelRegistration // Event containing the contract specifics and raw log

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
func (it *GpuManagerModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerModelRegistration)
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
		it.Event = new(GpuManagerModelRegistration)
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
func (it *GpuManagerModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerModelRegistration represents a ModelRegistration event raised by the GpuManager contract.
type GpuManagerModelRegistration struct {
	ModelId    uint32
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0xbf8d4447fa6c121c179656152534cb5032c1ce50f747e90c56580bec25583d81.
//
// Solidity: event ModelRegistration(uint32 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) FilterModelRegistration(opts *bind.FilterOpts, modelId []uint32, tier []uint16) (*GpuManagerModelRegistrationIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "ModelRegistration", modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerModelRegistrationIterator{contract: _GpuManager.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0xbf8d4447fa6c121c179656152534cb5032c1ce50f747e90c56580bec25583d81.
//
// Solidity: event ModelRegistration(uint32 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *GpuManagerModelRegistration, modelId []uint32, tier []uint16) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "ModelRegistration", modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerModelRegistration)
				if err := _GpuManager.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
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

// ParseModelRegistration is a log parse operation binding the contract event 0xbf8d4447fa6c121c179656152534cb5032c1ce50f747e90c56580bec25583d81.
//
// Solidity: event ModelRegistration(uint32 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_GpuManager *GpuManagerFilterer) ParseModelRegistration(log types.Log) (*GpuManagerModelRegistration, error) {
	event := new(GpuManagerModelRegistration)
	if err := _GpuManager.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the GpuManager contract.
type GpuManagerModelTierUpdateIterator struct {
	Event *GpuManagerModelTierUpdate // Event containing the contract specifics and raw log

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
func (it *GpuManagerModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerModelTierUpdate)
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
		it.Event = new(GpuManagerModelTierUpdate)
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
func (it *GpuManagerModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerModelTierUpdate represents a ModelTierUpdate event raised by the GpuManager contract.
type GpuManagerModelTierUpdate struct {
	ModelId uint32
	Tier    uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x4ecbcd19e308970fa368644f223de37bf9800e203349b5873d83970277c30356.
//
// Solidity: event ModelTierUpdate(uint32 indexed modelId, uint32 tier)
func (_GpuManager *GpuManagerFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, modelId []uint32) (*GpuManagerModelTierUpdateIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "ModelTierUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerModelTierUpdateIterator{contract: _GpuManager.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x4ecbcd19e308970fa368644f223de37bf9800e203349b5873d83970277c30356.
//
// Solidity: event ModelTierUpdate(uint32 indexed modelId, uint32 tier)
func (_GpuManager *GpuManagerFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *GpuManagerModelTierUpdate, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "ModelTierUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerModelTierUpdate)
				if err := _GpuManager.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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

// ParseModelTierUpdate is a log parse operation binding the contract event 0x4ecbcd19e308970fa368644f223de37bf9800e203349b5873d83970277c30356.
//
// Solidity: event ModelTierUpdate(uint32 indexed modelId, uint32 tier)
func (_GpuManager *GpuManagerFilterer) ParseModelTierUpdate(log types.Log) (*GpuManagerModelTierUpdate, error) {
	event := new(GpuManagerModelTierUpdate)
	if err := _GpuManager.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the GpuManager contract.
type GpuManagerModelUnregistrationIterator struct {
	Event *GpuManagerModelUnregistration // Event containing the contract specifics and raw log

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
func (it *GpuManagerModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerModelUnregistration)
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
		it.Event = new(GpuManagerModelUnregistration)
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
func (it *GpuManagerModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerModelUnregistration represents a ModelUnregistration event raised by the GpuManager contract.
type GpuManagerModelUnregistration struct {
	ModelId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x543408e7ce45c07531e494b8909d4d1b9dea7a8d8f5907b4673949a90fc56ba2.
//
// Solidity: event ModelUnregistration(uint32 indexed modelId)
func (_GpuManager *GpuManagerFilterer) FilterModelUnregistration(opts *bind.FilterOpts, modelId []uint32) (*GpuManagerModelUnregistrationIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "ModelUnregistration", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerModelUnregistrationIterator{contract: _GpuManager.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x543408e7ce45c07531e494b8909d4d1b9dea7a8d8f5907b4673949a90fc56ba2.
//
// Solidity: event ModelUnregistration(uint32 indexed modelId)
func (_GpuManager *GpuManagerFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *GpuManagerModelUnregistration, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "ModelUnregistration", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerModelUnregistration)
				if err := _GpuManager.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
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

// ParseModelUnregistration is a log parse operation binding the contract event 0x543408e7ce45c07531e494b8909d4d1b9dea7a8d8f5907b4673949a90fc56ba2.
//
// Solidity: event ModelUnregistration(uint32 indexed modelId)
func (_GpuManager *GpuManagerFilterer) ParseModelUnregistration(log types.Log) (*GpuManagerModelUnregistration, error) {
	event := new(GpuManagerModelUnregistration)
	if err := _GpuManager.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GpuManager contract.
type GpuManagerOwnershipTransferredIterator struct {
	Event *GpuManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GpuManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerOwnershipTransferred)
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
		it.Event = new(GpuManagerOwnershipTransferred)
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
func (it *GpuManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerOwnershipTransferred represents a OwnershipTransferred event raised by the GpuManager contract.
type GpuManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GpuManager *GpuManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GpuManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerOwnershipTransferredIterator{contract: _GpuManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GpuManager *GpuManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GpuManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerOwnershipTransferred)
				if err := _GpuManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseOwnershipTransferred(log types.Log) (*GpuManagerOwnershipTransferred, error) {
	event := new(GpuManagerOwnershipTransferred)
	if err := _GpuManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GpuManager contract.
type GpuManagerPausedIterator struct {
	Event *GpuManagerPaused // Event containing the contract specifics and raw log

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
func (it *GpuManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerPaused)
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
		it.Event = new(GpuManagerPaused)
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
func (it *GpuManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerPaused represents a Paused event raised by the GpuManager contract.
type GpuManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GpuManager *GpuManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*GpuManagerPausedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GpuManagerPausedIterator{contract: _GpuManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GpuManager *GpuManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GpuManagerPaused) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerPaused)
				if err := _GpuManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParsePaused(log types.Log) (*GpuManagerPaused, error) {
	event := new(GpuManagerPaused)
	if err := _GpuManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the GpuManager contract.
type GpuManagerPenaltyDurationUpdatedIterator struct {
	Event *GpuManagerPenaltyDurationUpdated // Event containing the contract specifics and raw log

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
func (it *GpuManagerPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerPenaltyDurationUpdated)
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
		it.Event = new(GpuManagerPenaltyDurationUpdated)
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
func (it *GpuManagerPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the GpuManager contract.
type GpuManagerPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_GpuManager *GpuManagerFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*GpuManagerPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &GpuManagerPenaltyDurationUpdatedIterator{contract: _GpuManager.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_GpuManager *GpuManagerFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *GpuManagerPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerPenaltyDurationUpdated)
				if err := _GpuManager.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParsePenaltyDurationUpdated(log types.Log) (*GpuManagerPenaltyDurationUpdated, error) {
	event := new(GpuManagerPenaltyDurationUpdated)
	if err := _GpuManager.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the GpuManager contract.
type GpuManagerRestakeIterator struct {
	Event *GpuManagerRestake // Event containing the contract specifics and raw log

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
func (it *GpuManagerRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerRestake)
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
		it.Event = new(GpuManagerRestake)
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
func (it *GpuManagerRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerRestake represents a Restake event raised by the GpuManager contract.
type GpuManagerRestake struct {
	Miner   common.Address
	ModelId uint32
	Restake *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0xd71961af2f46a633dc473cc0dda9e08783282fdb38c8f90482a143eb63b039e0.
//
// Solidity: event Restake(address indexed miner, uint32 indexed modelId, uint256 restake)
func (_GpuManager *GpuManagerFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, modelId []uint32) (*GpuManagerRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "Restake", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerRestakeIterator{contract: _GpuManager.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0xd71961af2f46a633dc473cc0dda9e08783282fdb38c8f90482a143eb63b039e0.
//
// Solidity: event Restake(address indexed miner, uint32 indexed modelId, uint256 restake)
func (_GpuManager *GpuManagerFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *GpuManagerRestake, miner []common.Address, modelId []uint32) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "Restake", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerRestake)
				if err := _GpuManager.contract.UnpackLog(event, "Restake", log); err != nil {
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

// ParseRestake is a log parse operation binding the contract event 0xd71961af2f46a633dc473cc0dda9e08783282fdb38c8f90482a143eb63b039e0.
//
// Solidity: event Restake(address indexed miner, uint32 indexed modelId, uint256 restake)
func (_GpuManager *GpuManagerFilterer) ParseRestake(log types.Log) (*GpuManagerRestake, error) {
	event := new(GpuManagerRestake)
	if err := _GpuManager.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the GpuManager contract.
type GpuManagerRewardClaimIterator struct {
	Event *GpuManagerRewardClaim // Event containing the contract specifics and raw log

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
func (it *GpuManagerRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerRewardClaim)
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
		it.Event = new(GpuManagerRewardClaim)
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
func (it *GpuManagerRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerRewardClaim represents a RewardClaim event raised by the GpuManager contract.
type GpuManagerRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_GpuManager *GpuManagerFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*GpuManagerRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &GpuManagerRewardClaimIterator{contract: _GpuManager.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_GpuManager *GpuManagerFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *GpuManagerRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerRewardClaim)
				if err := _GpuManager.contract.UnpackLog(event, "RewardClaim", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseRewardClaim(log types.Log) (*GpuManagerRewardClaim, error) {
	event := new(GpuManagerRewardClaim)
	if err := _GpuManager.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the GpuManager contract.
type GpuManagerRewardPerEpochIterator struct {
	Event *GpuManagerRewardPerEpoch // Event containing the contract specifics and raw log

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
func (it *GpuManagerRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerRewardPerEpoch)
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
		it.Event = new(GpuManagerRewardPerEpoch)
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
func (it *GpuManagerRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerRewardPerEpoch represents a RewardPerEpoch event raised by the GpuManager contract.
type GpuManagerRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_GpuManager *GpuManagerFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*GpuManagerRewardPerEpochIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &GpuManagerRewardPerEpochIterator{contract: _GpuManager.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_GpuManager *GpuManagerFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *GpuManagerRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerRewardPerEpoch)
				if err := _GpuManager.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseRewardPerEpoch(log types.Log) (*GpuManagerRewardPerEpoch, error) {
	event := new(GpuManagerRewardPerEpoch)
	if err := _GpuManager.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GpuManager contract.
type GpuManagerUnpausedIterator struct {
	Event *GpuManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *GpuManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerUnpaused)
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
		it.Event = new(GpuManagerUnpaused)
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
func (it *GpuManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerUnpaused represents a Unpaused event raised by the GpuManager contract.
type GpuManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GpuManager *GpuManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GpuManagerUnpausedIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GpuManagerUnpausedIterator{contract: _GpuManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GpuManager *GpuManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GpuManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerUnpaused)
				if err := _GpuManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseUnpaused(log types.Log) (*GpuManagerUnpaused, error) {
	event := new(GpuManagerUnpaused)
	if err := _GpuManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GpuManagerUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the GpuManager contract.
type GpuManagerUnstakeDelayTimeIterator struct {
	Event *GpuManagerUnstakeDelayTime // Event containing the contract specifics and raw log

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
func (it *GpuManagerUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GpuManagerUnstakeDelayTime)
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
		it.Event = new(GpuManagerUnstakeDelayTime)
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
func (it *GpuManagerUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GpuManagerUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GpuManagerUnstakeDelayTime represents a UnstakeDelayTime event raised by the GpuManager contract.
type GpuManagerUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_GpuManager *GpuManagerFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*GpuManagerUnstakeDelayTimeIterator, error) {

	logs, sub, err := _GpuManager.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &GpuManagerUnstakeDelayTimeIterator{contract: _GpuManager.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_GpuManager *GpuManagerFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *GpuManagerUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _GpuManager.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GpuManagerUnstakeDelayTime)
				if err := _GpuManager.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
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
func (_GpuManager *GpuManagerFilterer) ParseUnstakeDelayTime(log types.Log) (*GpuManagerUnstakeDelayTime, error) {
	event := new(GpuManagerUnstakeDelayTime)
	if err := _GpuManager.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
