// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gpumanager

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

// GPUManagerMetaData contains all meta data concerning the GPUManager contract.
var GPUManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameModelAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinFeeToUseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_modelCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"_models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_promptScheduler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinerUnstakeRequests\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unstakeAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"internalType\":\"structIGPUManager.UnstakeRequest[]\",\"name\":\"unstakeRequests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"getModelInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"internalType\":\"structIGPUManager.Model\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wEAIAmt\",\"type\":\"uint256\"}],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"modelCollection_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minerMinimumStake_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksPerEpoch_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerEpoch_\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unstakeDelayTime_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"penaltyDuration_\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"finePercentage_\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minFeeToUse_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"isActiveModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newPercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"}],\"name\":\"setMinFeeToUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinerMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPromptScheduler\",\"type\":\"address\"}],\"name\":\"setPromptSchedulerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"delayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakeDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"validateMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"minersRequired\",\"type\":\"uint256\"}],\"name\":\"validateModelAndChooseRandomMiner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GPUManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use GPUManagerMetaData.ABI instead.
var GPUManagerABI = GPUManagerMetaData.ABI

// GPUManager is an auto generated Go binding around an Ethereum contract.
type GPUManager struct {
	GPUManagerCaller     // Read-only binding to the contract
	GPUManagerTransactor // Write-only binding to the contract
	GPUManagerFilterer   // Log filterer for contract events
}

// GPUManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GPUManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPUManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GPUManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPUManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GPUManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPUManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GPUManagerSession struct {
	Contract     *GPUManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GPUManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GPUManagerCallerSession struct {
	Contract *GPUManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GPUManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GPUManagerTransactorSession struct {
	Contract     *GPUManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GPUManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GPUManagerRaw struct {
	Contract *GPUManager // Generic contract binding to access the raw methods on
}

// GPUManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GPUManagerCallerRaw struct {
	Contract *GPUManagerCaller // Generic read-only contract binding to access the raw methods on
}

// GPUManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GPUManagerTransactorRaw struct {
	Contract *GPUManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGPUManager creates a new instance of GPUManager, bound to a specific deployed contract.
func NewGPUManager(address common.Address, backend bind.ContractBackend) (*GPUManager, error) {
	contract, err := bindGPUManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GPUManager{GPUManagerCaller: GPUManagerCaller{contract: contract}, GPUManagerTransactor: GPUManagerTransactor{contract: contract}, GPUManagerFilterer: GPUManagerFilterer{contract: contract}}, nil
}

// NewGPUManagerCaller creates a new read-only instance of GPUManager, bound to a specific deployed contract.
func NewGPUManagerCaller(address common.Address, caller bind.ContractCaller) (*GPUManagerCaller, error) {
	contract, err := bindGPUManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GPUManagerCaller{contract: contract}, nil
}

// NewGPUManagerTransactor creates a new write-only instance of GPUManager, bound to a specific deployed contract.
func NewGPUManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GPUManagerTransactor, error) {
	contract, err := bindGPUManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GPUManagerTransactor{contract: contract}, nil
}

// NewGPUManagerFilterer creates a new log filterer instance of GPUManager, bound to a specific deployed contract.
func NewGPUManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GPUManagerFilterer, error) {
	contract, err := bindGPUManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GPUManagerFilterer{contract: contract}, nil
}

// bindGPUManager binds a generic wrapper to an already deployed contract.
func bindGPUManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GPUManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPUManager *GPUManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPUManager.Contract.GPUManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPUManager *GPUManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.Contract.GPUManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPUManager *GPUManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPUManager.Contract.GPUManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPUManager *GPUManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPUManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPUManager *GPUManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPUManager *GPUManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPUManager.Contract.contract.Transact(opts, method, params...)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerSession) BlocksPerEpoch() (*big.Int, error) {
	return _GPUManager.Contract.BlocksPerEpoch(&_GPUManager.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xb2424e3f.
//
// Solidity: function _blocksPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _GPUManager.Contract.BlocksPerEpoch(&_GPUManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GPUManager *GPUManagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GPUManager *GPUManagerSession) CurrentEpoch() (*big.Int, error) {
	return _GPUManager.Contract.CurrentEpoch(&_GPUManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x4c98e243.
//
// Solidity: function _currentEpoch() view returns(uint40)
func (_GPUManager *GPUManagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _GPUManager.Contract.CurrentEpoch(&_GPUManager.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GPUManager *GPUManagerCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GPUManager *GPUManagerSession) FinePercentage() (uint16, error) {
	return _GPUManager.Contract.FinePercentage(&_GPUManager.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x92cdf038.
//
// Solidity: function _finePercentage() view returns(uint16)
func (_GPUManager *GPUManagerCallerSession) FinePercentage() (uint16, error) {
	return _GPUManager.Contract.FinePercentage(&_GPUManager.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GPUManager *GPUManagerCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GPUManager *GPUManagerSession) LastBlock() (*big.Int, error) {
	return _GPUManager.Contract.LastBlock(&_GPUManager.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0xf712b279.
//
// Solidity: function _lastBlock() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) LastBlock() (*big.Int, error) {
	return _GPUManager.Contract.LastBlock(&_GPUManager.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GPUManager *GPUManagerCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GPUManager *GPUManagerSession) MaximumTier() (uint16, error) {
	return _GPUManager.Contract.MaximumTier(&_GPUManager.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x88f12044.
//
// Solidity: function _maximumTier() view returns(uint16)
func (_GPUManager *GPUManagerCallerSession) MaximumTier() (uint16, error) {
	return _GPUManager.Contract.MaximumTier(&_GPUManager.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GPUManager *GPUManagerCaller) MinFeeToUse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_minFeeToUse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GPUManager *GPUManagerSession) MinFeeToUse() (*big.Int, error) {
	return _GPUManager.Contract.MinFeeToUse(&_GPUManager.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0xf6a74d05.
//
// Solidity: function _minFeeToUse() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) MinFeeToUse() (*big.Int, error) {
	return _GPUManager.Contract.MinFeeToUse(&_GPUManager.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GPUManager *GPUManagerCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GPUManager *GPUManagerSession) MinerMinimumStake() (*big.Int, error) {
	return _GPUManager.Contract.MinerMinimumStake(&_GPUManager.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0xc5fc548d.
//
// Solidity: function _minerMinimumStake() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _GPUManager.Contract.MinerMinimumStake(&_GPUManager.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x54eb2d2a.
//
// Solidity: function _minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_GPUManager *GPUManagerCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_minerUnstakeRequests", arg0)

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
func (_GPUManager *GPUManagerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _GPUManager.Contract.MinerUnstakeRequests(&_GPUManager.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x54eb2d2a.
//
// Solidity: function _minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_GPUManager *GPUManagerCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _GPUManager.Contract.MinerUnstakeRequests(&_GPUManager.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xb1a976ef.
//
// Solidity: function _miners(address ) view returns(uint256 stake, uint32 modelId, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_GPUManager *GPUManagerCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_miners", arg0)

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
func (_GPUManager *GPUManagerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _GPUManager.Contract.Miners(&_GPUManager.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xb1a976ef.
//
// Solidity: function _miners(address ) view returns(uint256 stake, uint32 modelId, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_GPUManager *GPUManagerCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	ModelId          uint32
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _GPUManager.Contract.Miners(&_GPUManager.CallOpts, arg0)
}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GPUManager *GPUManagerCaller) ModelCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_modelCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GPUManager *GPUManagerSession) ModelCollection() (common.Address, error) {
	return _GPUManager.Contract.ModelCollection(&_GPUManager.CallOpts)
}

// ModelCollection is a free data retrieval call binding the contract method 0x1c49c2d6.
//
// Solidity: function _modelCollection() view returns(address)
func (_GPUManager *GPUManagerCallerSession) ModelCollection() (common.Address, error) {
	return _GPUManager.Contract.ModelCollection(&_GPUManager.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x55f89085.
//
// Solidity: function _models(uint32 ) view returns(uint256 minimumFee, uint32 tier)
func (_GPUManager *GPUManagerCaller) Models(opts *bind.CallOpts, arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_models", arg0)

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
func (_GPUManager *GPUManagerSession) Models(arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _GPUManager.Contract.Models(&_GPUManager.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x55f89085.
//
// Solidity: function _models(uint32 ) view returns(uint256 minimumFee, uint32 tier)
func (_GPUManager *GPUManagerCallerSession) Models(arg0 uint32) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _GPUManager.Contract.Models(&_GPUManager.CallOpts, arg0)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GPUManager *GPUManagerCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GPUManager *GPUManagerSession) PenaltyDuration() (*big.Int, error) {
	return _GPUManager.Contract.PenaltyDuration(&_GPUManager.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0xab692134.
//
// Solidity: function _penaltyDuration() view returns(uint40)
func (_GPUManager *GPUManagerCallerSession) PenaltyDuration() (*big.Int, error) {
	return _GPUManager.Contract.PenaltyDuration(&_GPUManager.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GPUManager *GPUManagerCaller) PromptScheduler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_promptScheduler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GPUManager *GPUManagerSession) PromptScheduler() (common.Address, error) {
	return _GPUManager.Contract.PromptScheduler(&_GPUManager.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_GPUManager *GPUManagerCallerSession) PromptScheduler() (common.Address, error) {
	return _GPUManager.Contract.PromptScheduler(&_GPUManager.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0xa662f84d.
//
// Solidity: function _rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_GPUManager *GPUManagerCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_rewardInEpoch", arg0)

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
func (_GPUManager *GPUManagerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _GPUManager.Contract.RewardInEpoch(&_GPUManager.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0xa662f84d.
//
// Solidity: function _rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_GPUManager *GPUManagerCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _GPUManager.Contract.RewardInEpoch(&_GPUManager.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerSession) RewardPerEpoch() (*big.Int, error) {
	return _GPUManager.Contract.RewardPerEpoch(&_GPUManager.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x62423112.
//
// Solidity: function _rewardPerEpoch() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _GPUManager.Contract.RewardPerEpoch(&_GPUManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GPUManager *GPUManagerCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GPUManager *GPUManagerSession) Treasury() (common.Address, error) {
	return _GPUManager.Contract.Treasury(&_GPUManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0xe319a3d9.
//
// Solidity: function _treasury() view returns(address)
func (_GPUManager *GPUManagerCallerSession) Treasury() (common.Address, error) {
	return _GPUManager.Contract.Treasury(&_GPUManager.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GPUManager *GPUManagerCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GPUManager *GPUManagerSession) UnstakeDelayTime() (*big.Int, error) {
	return _GPUManager.Contract.UnstakeDelayTime(&_GPUManager.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0x72b1f3e4.
//
// Solidity: function _unstakeDelayTime() view returns(uint40)
func (_GPUManager *GPUManagerCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _GPUManager.Contract.UnstakeDelayTime(&_GPUManager.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GPUManager *GPUManagerCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GPUManager *GPUManagerSession) WEAIToken() (common.Address, error) {
	return _GPUManager.Contract.WEAIToken(&_GPUManager.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_GPUManager *GPUManagerCallerSession) WEAIToken() (common.Address, error) {
	return _GPUManager.Contract.WEAIToken(&_GPUManager.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_GPUManager *GPUManagerCaller) GetAllMinerUnstakeRequests(opts *bind.CallOpts) (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getAllMinerUnstakeRequests")

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
func (_GPUManager *GPUManagerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	return _GPUManager.Contract.GetAllMinerUnstakeRequests(&_GPUManager.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_GPUManager *GPUManagerCallerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IGPUManagerUnstakeRequest
}, error) {
	return _GPUManager.Contract.GetAllMinerUnstakeRequests(&_GPUManager.CallOpts)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GPUManager *GPUManagerCaller) GetMinFeeToUse(opts *bind.CallOpts, modelId uint32) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getMinFeeToUse", modelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GPUManager *GPUManagerSession) GetMinFeeToUse(modelId uint32) (*big.Int, error) {
	return _GPUManager.Contract.GetMinFeeToUse(&_GPUManager.CallOpts, modelId)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0x963a0278.
//
// Solidity: function getMinFeeToUse(uint32 modelId) view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) GetMinFeeToUse(modelId uint32) (*big.Int, error) {
	return _GPUManager.Contract.GetMinFeeToUse(&_GPUManager.CallOpts, modelId)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GPUManager *GPUManagerCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GPUManager *GPUManagerSession) GetMinerAddresses() ([]common.Address, error) {
	return _GPUManager.Contract.GetMinerAddresses(&_GPUManager.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_GPUManager *GPUManagerCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _GPUManager.Contract.GetMinerAddresses(&_GPUManager.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GPUManager *GPUManagerCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, modelId uint32) ([]common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getMinerAddressesOfModel", modelId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GPUManager *GPUManagerSession) GetMinerAddressesOfModel(modelId uint32) ([]common.Address, error) {
	return _GPUManager.Contract.GetMinerAddressesOfModel(&_GPUManager.CallOpts, modelId)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0xa5f85cc8.
//
// Solidity: function getMinerAddressesOfModel(uint32 modelId) view returns(address[])
func (_GPUManager *GPUManagerCallerSession) GetMinerAddressesOfModel(modelId uint32) ([]common.Address, error) {
	return _GPUManager.Contract.GetMinerAddressesOfModel(&_GPUManager.CallOpts, modelId)
}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GPUManager *GPUManagerCaller) GetModelIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getModelIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GPUManager *GPUManagerSession) GetModelIds() ([]*big.Int, error) {
	return _GPUManager.Contract.GetModelIds(&_GPUManager.CallOpts)
}

// GetModelIds is a free data retrieval call binding the contract method 0x84881115.
//
// Solidity: function getModelIds() view returns(uint256[])
func (_GPUManager *GPUManagerCallerSession) GetModelIds() ([]*big.Int, error) {
	return _GPUManager.Contract.GetModelIds(&_GPUManager.CallOpts)
}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GPUManager *GPUManagerCaller) GetModelInfo(opts *bind.CallOpts, modelId uint32) (IGPUManagerModel, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getModelInfo", modelId)

	if err != nil {
		return *new(IGPUManagerModel), err
	}

	out0 := *abi.ConvertType(out[0], new(IGPUManagerModel)).(*IGPUManagerModel)

	return out0, err

}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GPUManager *GPUManagerSession) GetModelInfo(modelId uint32) (IGPUManagerModel, error) {
	return _GPUManager.Contract.GetModelInfo(&_GPUManager.CallOpts, modelId)
}

// GetModelInfo is a free data retrieval call binding the contract method 0x77495c20.
//
// Solidity: function getModelInfo(uint32 modelId) view returns((uint256,uint32))
func (_GPUManager *GPUManagerCallerSession) GetModelInfo(modelId uint32) (IGPUManagerModel, error) {
	return _GPUManager.Contract.GetModelInfo(&_GPUManager.CallOpts, modelId)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GPUManager *GPUManagerCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GPUManager *GPUManagerSession) GetNOMiner() (*big.Int, error) {
	return _GPUManager.Contract.GetNOMiner(&_GPUManager.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) GetNOMiner() (*big.Int, error) {
	return _GPUManager.Contract.GetNOMiner(&_GPUManager.CallOpts)
}

// IsActiveModel is a free data retrieval call binding the contract method 0xbce2845a.
//
// Solidity: function isActiveModel(uint32 modelId) view returns(bool)
func (_GPUManager *GPUManagerCaller) IsActiveModel(opts *bind.CallOpts, modelId uint32) (bool, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "isActiveModel", modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveModel is a free data retrieval call binding the contract method 0xbce2845a.
//
// Solidity: function isActiveModel(uint32 modelId) view returns(bool)
func (_GPUManager *GPUManagerSession) IsActiveModel(modelId uint32) (bool, error) {
	return _GPUManager.Contract.IsActiveModel(&_GPUManager.CallOpts, modelId)
}

// IsActiveModel is a free data retrieval call binding the contract method 0xbce2845a.
//
// Solidity: function isActiveModel(uint32 modelId) view returns(bool)
func (_GPUManager *GPUManagerCallerSession) IsActiveModel(modelId uint32) (bool, error) {
	return _GPUManager.Contract.IsActiveModel(&_GPUManager.CallOpts, modelId)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GPUManager *GPUManagerCaller) Multiplier(opts *bind.CallOpts, miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "multiplier", miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GPUManager *GPUManagerSession) Multiplier(miner common.Address) (*big.Int, error) {
	return _GPUManager.Contract.Multiplier(&_GPUManager.CallOpts, miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address miner) view returns(uint256)
func (_GPUManager *GPUManagerCallerSession) Multiplier(miner common.Address) (*big.Int, error) {
	return _GPUManager.Contract.Multiplier(&_GPUManager.CallOpts, miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPUManager *GPUManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPUManager *GPUManagerSession) Owner() (common.Address, error) {
	return _GPUManager.Contract.Owner(&_GPUManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPUManager *GPUManagerCallerSession) Owner() (common.Address, error) {
	return _GPUManager.Contract.Owner(&_GPUManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GPUManager *GPUManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GPUManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GPUManager *GPUManagerSession) Paused() (bool, error) {
	return _GPUManager.Contract.Paused(&_GPUManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GPUManager *GPUManagerCallerSession) Paused() (bool, error) {
	return _GPUManager.Contract.Paused(&_GPUManager.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GPUManager *GPUManagerTransactor) ClaimReward(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "claimReward", miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GPUManager *GPUManagerSession) ClaimReward(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.ClaimReward(&_GPUManager.TransactOpts, miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address miner) returns()
func (_GPUManager *GPUManagerTransactorSession) ClaimReward(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.ClaimReward(&_GPUManager.TransactOpts, miner)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "forceChangeModelForMiner", miner, modelId)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GPUManager *GPUManagerSession) ForceChangeModelForMiner(miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.ForceChangeModelForMiner(&_GPUManager.TransactOpts, miner, modelId)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x49f5ef62.
//
// Solidity: function forceChangeModelForMiner(address miner, uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactorSession) ForceChangeModelForMiner(miner common.Address, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.ForceChangeModelForMiner(&_GPUManager.TransactOpts, miner, modelId)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GPUManager *GPUManagerTransactor) IncreaseMinerStake(opts *bind.TransactOpts, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "increaseMinerStake", wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GPUManager *GPUManagerSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.IncreaseMinerStake(&_GPUManager.TransactOpts, wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_GPUManager *GPUManagerTransactorSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.IncreaseMinerStake(&_GPUManager.TransactOpts, wEAIAmt)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GPUManager *GPUManagerTransactor) Initialize(opts *bind.TransactOpts, wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "initialize", wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GPUManager *GPUManagerSession) Initialize(wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.Initialize(&_GPUManager.TransactOpts, wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// Initialize is a paid mutator transaction binding the contract method 0x04bb771f.
//
// Solidity: function initialize(address wEAIToken_, address modelCollection_, address treasury_, uint256 minerMinimumStake_, uint256 blocksPerEpoch_, uint256 rewardPerEpoch_, uint40 unstakeDelayTime_, uint40 penaltyDuration_, uint16 finePercentage_, uint256 minFeeToUse_) returns()
func (_GPUManager *GPUManagerTransactorSession) Initialize(wEAIToken_ common.Address, modelCollection_ common.Address, treasury_ common.Address, minerMinimumStake_ *big.Int, blocksPerEpoch_ *big.Int, rewardPerEpoch_ *big.Int, unstakeDelayTime_ *big.Int, penaltyDuration_ *big.Int, finePercentage_ uint16, minFeeToUse_ *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.Initialize(&_GPUManager.TransactOpts, wEAIToken_, modelCollection_, treasury_, minerMinimumStake_, blocksPerEpoch_, rewardPerEpoch_, unstakeDelayTime_, penaltyDuration_, finePercentage_, minFeeToUse_)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GPUManager *GPUManagerTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GPUManager *GPUManagerSession) JoinForMinting() (*types.Transaction, error) {
	return _GPUManager.Contract.JoinForMinting(&_GPUManager.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_GPUManager *GPUManagerTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _GPUManager.Contract.JoinForMinting(&_GPUManager.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "registerMiner", tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterMiner(&_GPUManager.TransactOpts, tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerTransactorSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterMiner(&_GPUManager.TransactOpts, tier)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactor) RegisterMiner0(opts *bind.TransactOpts, tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "registerMiner0", tier, modelId)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GPUManager *GPUManagerSession) RegisterMiner0(tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterMiner0(&_GPUManager.TransactOpts, tier, modelId)
}

// RegisterMiner0 is a paid mutator transaction binding the contract method 0x70423c2a.
//
// Solidity: function registerMiner(uint16 tier, uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactorSession) RegisterMiner0(tier uint16, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterMiner0(&_GPUManager.TransactOpts, tier, modelId)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerTransactor) RegisterModel(opts *bind.TransactOpts, modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "registerModel", modelId, tier, minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerSession) RegisterModel(modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterModel(&_GPUManager.TransactOpts, modelId, tier, minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0x88184775.
//
// Solidity: function registerModel(uint32 modelId, uint16 tier, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerTransactorSession) RegisterModel(modelId uint32, tier uint16, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.RegisterModel(&_GPUManager.TransactOpts, modelId, tier, minimumFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPUManager *GPUManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPUManager *GPUManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPUManager.Contract.RenounceOwnership(&_GPUManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPUManager *GPUManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPUManager.Contract.RenounceOwnership(&_GPUManager.TransactOpts)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.RestakeForMiner(&_GPUManager.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_GPUManager *GPUManagerTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.RestakeForMiner(&_GPUManager.TransactOpts, tier)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GPUManager *GPUManagerTransactor) RewardToClaim(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "rewardToClaim", miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GPUManager *GPUManagerSession) RewardToClaim(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.RewardToClaim(&_GPUManager.TransactOpts, miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address miner) returns(uint256)
func (_GPUManager *GPUManagerTransactorSession) RewardToClaim(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.RewardToClaim(&_GPUManager.TransactOpts, miner)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GPUManager *GPUManagerTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, blocks *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setBlocksPerEpoch", blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GPUManager *GPUManagerSession) SetBlocksPerEpoch(blocks *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetBlocksPerEpoch(&_GPUManager.TransactOpts, blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 blocks) returns()
func (_GPUManager *GPUManagerTransactorSession) SetBlocksPerEpoch(blocks *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetBlocksPerEpoch(&_GPUManager.TransactOpts, blocks)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GPUManager *GPUManagerTransactor) SetFinePercentage(opts *bind.TransactOpts, newPercentage uint16) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setFinePercentage", newPercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GPUManager *GPUManagerSession) SetFinePercentage(newPercentage uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.SetFinePercentage(&_GPUManager.TransactOpts, newPercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 newPercentage) returns()
func (_GPUManager *GPUManagerTransactorSession) SetFinePercentage(newPercentage uint16) (*types.Transaction, error) {
	return _GPUManager.Contract.SetFinePercentage(&_GPUManager.TransactOpts, newPercentage)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GPUManager *GPUManagerTransactor) SetMinFeeToUse(opts *bind.TransactOpts, minFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setMinFeeToUse", minFee)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GPUManager *GPUManagerSession) SetMinFeeToUse(minFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetMinFeeToUse(&_GPUManager.TransactOpts, minFee)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 minFee) returns()
func (_GPUManager *GPUManagerTransactorSession) SetMinFeeToUse(minFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetMinFeeToUse(&_GPUManager.TransactOpts, minFee)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GPUManager *GPUManagerTransactor) SetMinerMinimumStake(opts *bind.TransactOpts, _minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setMinerMinimumStake", _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GPUManager *GPUManagerSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetMinerMinimumStake(&_GPUManager.TransactOpts, _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_GPUManager *GPUManagerTransactorSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetMinerMinimumStake(&_GPUManager.TransactOpts, _minerMinimumStake)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GPUManager *GPUManagerTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, newReward *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setNewRewardInEpoch", newReward)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GPUManager *GPUManagerSession) SetNewRewardInEpoch(newReward *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetNewRewardInEpoch(&_GPUManager.TransactOpts, newReward)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 newReward) returns()
func (_GPUManager *GPUManagerTransactorSession) SetNewRewardInEpoch(newReward *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetNewRewardInEpoch(&_GPUManager.TransactOpts, newReward)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GPUManager *GPUManagerTransactor) SetPenaltyDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setPenaltyDuration", duration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GPUManager *GPUManagerSession) SetPenaltyDuration(duration *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetPenaltyDuration(&_GPUManager.TransactOpts, duration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 duration) returns()
func (_GPUManager *GPUManagerTransactorSession) SetPenaltyDuration(duration *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetPenaltyDuration(&_GPUManager.TransactOpts, duration)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GPUManager *GPUManagerTransactor) SetPromptSchedulerAddress(opts *bind.TransactOpts, newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setPromptSchedulerAddress", newPromptScheduler)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GPUManager *GPUManagerSession) SetPromptSchedulerAddress(newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.SetPromptSchedulerAddress(&_GPUManager.TransactOpts, newPromptScheduler)
}

// SetPromptSchedulerAddress is a paid mutator transaction binding the contract method 0x00f19f45.
//
// Solidity: function setPromptSchedulerAddress(address newPromptScheduler) returns()
func (_GPUManager *GPUManagerTransactorSession) SetPromptSchedulerAddress(newPromptScheduler common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.SetPromptSchedulerAddress(&_GPUManager.TransactOpts, newPromptScheduler)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GPUManager *GPUManagerTransactor) SetUnstakeDelayTime(opts *bind.TransactOpts, delayTime *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setUnstakeDelayTime", delayTime)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GPUManager *GPUManagerSession) SetUnstakeDelayTime(delayTime *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetUnstakeDelayTime(&_GPUManager.TransactOpts, delayTime)
}

// SetUnstakeDelayTime is a paid mutator transaction binding the contract method 0x466ca9f9.
//
// Solidity: function setUnstakeDelayTime(uint40 delayTime) returns()
func (_GPUManager *GPUManagerTransactorSession) SetUnstakeDelayTime(delayTime *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.SetUnstakeDelayTime(&_GPUManager.TransactOpts, delayTime)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GPUManager *GPUManagerTransactor) SetWEAIAddress(opts *bind.TransactOpts, wEAIToken common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "setWEAIAddress", wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GPUManager *GPUManagerSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.SetWEAIAddress(&_GPUManager.TransactOpts, wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_GPUManager *GPUManagerTransactorSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.SetWEAIAddress(&_GPUManager.TransactOpts, wEAIToken)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GPUManager *GPUManagerTransactor) SlashMiner(opts *bind.TransactOpts, miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "slashMiner", miner, isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GPUManager *GPUManagerSession) SlashMiner(miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GPUManager.Contract.SlashMiner(&_GPUManager.TransactOpts, miner, isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address miner, bool isFined) returns()
func (_GPUManager *GPUManagerTransactorSession) SlashMiner(miner common.Address, isFined bool) (*types.Transaction, error) {
	return _GPUManager.Contract.SlashMiner(&_GPUManager.TransactOpts, miner, isFined)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPUManager *GPUManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPUManager *GPUManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.TransferOwnership(&_GPUManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPUManager *GPUManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.TransferOwnership(&_GPUManager.TransactOpts, newOwner)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GPUManager *GPUManagerTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GPUManager *GPUManagerSession) UnregisterMiner() (*types.Transaction, error) {
	return _GPUManager.Contract.UnregisterMiner(&_GPUManager.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_GPUManager *GPUManagerTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _GPUManager.Contract.UnregisterMiner(&_GPUManager.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactor) UnregisterModel(opts *bind.TransactOpts, modelId uint32) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "unregisterModel", modelId)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GPUManager *GPUManagerSession) UnregisterModel(modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.UnregisterModel(&_GPUManager.TransactOpts, modelId)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0x781f1453.
//
// Solidity: function unregisterModel(uint32 modelId) returns()
func (_GPUManager *GPUManagerTransactorSession) UnregisterModel(modelId uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.UnregisterModel(&_GPUManager.TransactOpts, modelId)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GPUManager *GPUManagerTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GPUManager *GPUManagerSession) UnstakeForMiner() (*types.Transaction, error) {
	return _GPUManager.Contract.UnstakeForMiner(&_GPUManager.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_GPUManager *GPUManagerTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _GPUManager.Contract.UnstakeForMiner(&_GPUManager.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GPUManager *GPUManagerTransactor) UpdateEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "updateEpoch")
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GPUManager *GPUManagerSession) UpdateEpoch() (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateEpoch(&_GPUManager.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_GPUManager *GPUManagerTransactorSession) UpdateEpoch() (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateEpoch(&_GPUManager.TransactOpts)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "updateModelMinimumFee", modelId, minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerSession) UpdateModelMinimumFee(modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateModelMinimumFee(&_GPUManager.TransactOpts, modelId, minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0x13ee7dbc.
//
// Solidity: function updateModelMinimumFee(uint32 modelId, uint256 minimumFee) returns()
func (_GPUManager *GPUManagerTransactorSession) UpdateModelMinimumFee(modelId uint32, minimumFee *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateModelMinimumFee(&_GPUManager.TransactOpts, modelId, minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GPUManager *GPUManagerTransactor) UpdateModelTier(opts *bind.TransactOpts, modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "updateModelTier", modelId, tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GPUManager *GPUManagerSession) UpdateModelTier(modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateModelTier(&_GPUManager.TransactOpts, modelId, tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0xfdf22bc8.
//
// Solidity: function updateModelTier(uint32 modelId, uint32 tier) returns()
func (_GPUManager *GPUManagerTransactorSession) UpdateModelTier(modelId uint32, tier uint32) (*types.Transaction, error) {
	return _GPUManager.Contract.UpdateModelTier(&_GPUManager.TransactOpts, modelId, tier)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GPUManager *GPUManagerTransactor) ValidateMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "validateMiner", miner)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GPUManager *GPUManagerSession) ValidateMiner(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.ValidateMiner(&_GPUManager.TransactOpts, miner)
}

// ValidateMiner is a paid mutator transaction binding the contract method 0xdfecce6f.
//
// Solidity: function validateMiner(address miner) returns()
func (_GPUManager *GPUManagerTransactorSession) ValidateMiner(miner common.Address) (*types.Transaction, error) {
	return _GPUManager.Contract.ValidateMiner(&_GPUManager.TransactOpts, miner)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GPUManager *GPUManagerTransactor) ValidateModelAndChooseRandomMiner(opts *bind.TransactOpts, modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GPUManager.contract.Transact(opts, "validateModelAndChooseRandomMiner", modelId, minersRequired)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GPUManager *GPUManagerSession) ValidateModelAndChooseRandomMiner(modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.ValidateModelAndChooseRandomMiner(&_GPUManager.TransactOpts, modelId, minersRequired)
}

// ValidateModelAndChooseRandomMiner is a paid mutator transaction binding the contract method 0xe13f220e.
//
// Solidity: function validateModelAndChooseRandomMiner(uint32 modelId, uint256 minersRequired) returns(address, uint256)
func (_GPUManager *GPUManagerTransactorSession) ValidateModelAndChooseRandomMiner(modelId uint32, minersRequired *big.Int) (*types.Transaction, error) {
	return _GPUManager.Contract.ValidateModelAndChooseRandomMiner(&_GPUManager.TransactOpts, modelId, minersRequired)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GPUManager *GPUManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPUManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GPUManager *GPUManagerSession) Receive() (*types.Transaction, error) {
	return _GPUManager.Contract.Receive(&_GPUManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GPUManager *GPUManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _GPUManager.Contract.Receive(&_GPUManager.TransactOpts)
}

// GPUManagerBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the GPUManager contract.
type GPUManagerBlocksPerEpochIterator struct {
	Event *GPUManagerBlocksPerEpoch // Event containing the contract specifics and raw log

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
func (it *GPUManagerBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerBlocksPerEpoch)
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
		it.Event = new(GPUManagerBlocksPerEpoch)
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
func (it *GPUManagerBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerBlocksPerEpoch represents a BlocksPerEpoch event raised by the GPUManager contract.
type GPUManagerBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_GPUManager *GPUManagerFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*GPUManagerBlocksPerEpochIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &GPUManagerBlocksPerEpochIterator{contract: _GPUManager.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_GPUManager *GPUManagerFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *GPUManagerBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerBlocksPerEpoch)
				if err := _GPUManager.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseBlocksPerEpoch(log types.Log) (*GPUManagerBlocksPerEpoch, error) {
	event := new(GPUManagerBlocksPerEpoch)
	if err := _GPUManager.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the GPUManager contract.
type GPUManagerFinePercentageUpdatedIterator struct {
	Event *GPUManagerFinePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *GPUManagerFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerFinePercentageUpdated)
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
		it.Event = new(GPUManagerFinePercentageUpdated)
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
func (it *GPUManagerFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerFinePercentageUpdated represents a FinePercentageUpdated event raised by the GPUManager contract.
type GPUManagerFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_GPUManager *GPUManagerFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*GPUManagerFinePercentageUpdatedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &GPUManagerFinePercentageUpdatedIterator{contract: _GPUManager.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_GPUManager *GPUManagerFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *GPUManagerFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerFinePercentageUpdated)
				if err := _GPUManager.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseFinePercentageUpdated(log types.Log) (*GPUManagerFinePercentageUpdated, error) {
	event := new(GPUManagerFinePercentageUpdated)
	if err := _GPUManager.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the GPUManager contract.
type GPUManagerFraudulentMinerPenalizedIterator struct {
	Event *GPUManagerFraudulentMinerPenalized // Event containing the contract specifics and raw log

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
func (it *GPUManagerFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerFraudulentMinerPenalized)
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
		it.Event = new(GPUManagerFraudulentMinerPenalized)
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
func (it *GPUManagerFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the GPUManager contract.
type GPUManagerFraudulentMinerPenalized struct {
	Miner    common.Address
	ModelId  uint32
	Treasury common.Address
	Fine     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x396ee931f435c63405d255f5e0d31a0d1a1f6b57d59ef9559155464a15b13593.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, uint32 indexed modelId, address indexed treasury, uint256 fine)
func (_GPUManager *GPUManagerFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelId []uint32, treasury []common.Address) (*GPUManagerFraudulentMinerPenalizedIterator, error) {

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

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelIdRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerFraudulentMinerPenalizedIterator{contract: _GPUManager.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x396ee931f435c63405d255f5e0d31a0d1a1f6b57d59ef9559155464a15b13593.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, uint32 indexed modelId, address indexed treasury, uint256 fine)
func (_GPUManager *GPUManagerFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *GPUManagerFraudulentMinerPenalized, miner []common.Address, modelId []uint32, treasury []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelIdRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerFraudulentMinerPenalized)
				if err := _GPUManager.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseFraudulentMinerPenalized(log types.Log) (*GPUManagerFraudulentMinerPenalized, error) {
	event := new(GPUManagerFraudulentMinerPenalized)
	if err := _GPUManager.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GPUManager contract.
type GPUManagerInitializedIterator struct {
	Event *GPUManagerInitialized // Event containing the contract specifics and raw log

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
func (it *GPUManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerInitialized)
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
		it.Event = new(GPUManagerInitialized)
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
func (it *GPUManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerInitialized represents a Initialized event raised by the GPUManager contract.
type GPUManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GPUManager *GPUManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*GPUManagerInitializedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GPUManagerInitializedIterator{contract: _GPUManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GPUManager *GPUManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GPUManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerInitialized)
				if err := _GPUManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseInitialized(log types.Log) (*GPUManagerInitialized, error) {
	event := new(GPUManagerInitialized)
	if err := _GPUManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinFeeToUseUpdatedIterator is returned from FilterMinFeeToUseUpdated and is used to iterate over the raw logs and unpacked data for MinFeeToUseUpdated events raised by the GPUManager contract.
type GPUManagerMinFeeToUseUpdatedIterator struct {
	Event *GPUManagerMinFeeToUseUpdated // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinFeeToUseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinFeeToUseUpdated)
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
		it.Event = new(GPUManagerMinFeeToUseUpdated)
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
func (it *GPUManagerMinFeeToUseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinFeeToUseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinFeeToUseUpdated represents a MinFeeToUseUpdated event raised by the GPUManager contract.
type GPUManagerMinFeeToUseUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeToUseUpdated is a free log retrieval operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_GPUManager *GPUManagerFilterer) FilterMinFeeToUseUpdated(opts *bind.FilterOpts) (*GPUManagerMinFeeToUseUpdatedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinFeeToUseUpdatedIterator{contract: _GPUManager.contract, event: "MinFeeToUseUpdated", logs: logs, sub: sub}, nil
}

// WatchMinFeeToUseUpdated is a free log subscription operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_GPUManager *GPUManagerFilterer) WatchMinFeeToUseUpdated(opts *bind.WatchOpts, sink chan<- *GPUManagerMinFeeToUseUpdated) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinFeeToUseUpdated)
				if err := _GPUManager.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinFeeToUseUpdated(log types.Log) (*GPUManagerMinFeeToUseUpdated, error) {
	event := new(GPUManagerMinFeeToUseUpdated)
	if err := _GPUManager.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the GPUManager contract.
type GPUManagerMinerDeactivatedIterator struct {
	Event *GPUManagerMinerDeactivated // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerDeactivated)
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
		it.Event = new(GPUManagerMinerDeactivated)
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
func (it *GPUManagerMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerDeactivated represents a MinerDeactivated event raised by the GPUManager contract.
type GPUManagerMinerDeactivated struct {
	Miner      common.Address
	ModelId    uint32
	ActiveTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x6e4a7233a3b583018e3a3d018e76ad619bab8ad6e8fe05e12cb83ec1fa75d85e.
//
// Solidity: event MinerDeactivated(address indexed miner, uint32 indexed modelId, uint40 activeTime)
func (_GPUManager *GPUManagerFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelId []uint32) (*GPUManagerMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerDeactivatedIterator{contract: _GPUManager.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x6e4a7233a3b583018e3a3d018e76ad619bab8ad6e8fe05e12cb83ec1fa75d85e.
//
// Solidity: event MinerDeactivated(address indexed miner, uint32 indexed modelId, uint40 activeTime)
func (_GPUManager *GPUManagerFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerDeactivated, miner []common.Address, modelId []uint32) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerDeactivated)
				if err := _GPUManager.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerDeactivated(log types.Log) (*GPUManagerMinerDeactivated, error) {
	event := new(GPUManagerMinerDeactivated)
	if err := _GPUManager.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the GPUManager contract.
type GPUManagerMinerExtraStakeIterator struct {
	Event *GPUManagerMinerExtraStake // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerExtraStake)
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
		it.Event = new(GPUManagerMinerExtraStake)
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
func (it *GPUManagerMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerExtraStake represents a MinerExtraStake event raised by the GPUManager contract.
type GPUManagerMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_GPUManager *GPUManagerFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*GPUManagerMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerExtraStakeIterator{contract: _GPUManager.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_GPUManager *GPUManagerFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerExtraStake)
				if err := _GPUManager.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerExtraStake(log types.Log) (*GPUManagerMinerExtraStake, error) {
	event := new(GPUManagerMinerExtraStake)
	if err := _GPUManager.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the GPUManager contract.
type GPUManagerMinerJoinIterator struct {
	Event *GPUManagerMinerJoin // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerJoin)
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
		it.Event = new(GPUManagerMinerJoin)
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
func (it *GPUManagerMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerJoin represents a MinerJoin event raised by the GPUManager contract.
type GPUManagerMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_GPUManager *GPUManagerFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*GPUManagerMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerJoinIterator{contract: _GPUManager.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_GPUManager *GPUManagerFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerJoin)
				if err := _GPUManager.contract.UnpackLog(event, "MinerJoin", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerJoin(log types.Log) (*GPUManagerMinerJoin, error) {
	event := new(GPUManagerMinerJoin)
	if err := _GPUManager.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the GPUManager contract.
type GPUManagerMinerRegistrationIterator struct {
	Event *GPUManagerMinerRegistration // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerRegistration)
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
		it.Event = new(GPUManagerMinerRegistration)
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
func (it *GPUManagerMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerRegistration represents a MinerRegistration event raised by the GPUManager contract.
type GPUManagerMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_GPUManager *GPUManagerFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*GPUManagerMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerRegistrationIterator{contract: _GPUManager.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_GPUManager *GPUManagerFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerRegistration)
				if err := _GPUManager.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerRegistration(log types.Log) (*GPUManagerMinerRegistration, error) {
	event := new(GPUManagerMinerRegistration)
	if err := _GPUManager.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the GPUManager contract.
type GPUManagerMinerUnregistrationIterator struct {
	Event *GPUManagerMinerUnregistration // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerUnregistration)
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
		it.Event = new(GPUManagerMinerUnregistration)
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
func (it *GPUManagerMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerUnregistration represents a MinerUnregistration event raised by the GPUManager contract.
type GPUManagerMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_GPUManager *GPUManagerFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*GPUManagerMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerUnregistrationIterator{contract: _GPUManager.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_GPUManager *GPUManagerFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerUnregistration)
				if err := _GPUManager.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerUnregistration(log types.Log) (*GPUManagerMinerUnregistration, error) {
	event := new(GPUManagerMinerUnregistration)
	if err := _GPUManager.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the GPUManager contract.
type GPUManagerMinerUnstakeIterator struct {
	Event *GPUManagerMinerUnstake // Event containing the contract specifics and raw log

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
func (it *GPUManagerMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerMinerUnstake)
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
		it.Event = new(GPUManagerMinerUnstake)
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
func (it *GPUManagerMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerMinerUnstake represents a MinerUnstake event raised by the GPUManager contract.
type GPUManagerMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_GPUManager *GPUManagerFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*GPUManagerMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerMinerUnstakeIterator{contract: _GPUManager.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_GPUManager *GPUManagerFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *GPUManagerMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerMinerUnstake)
				if err := _GPUManager.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseMinerUnstake(log types.Log) (*GPUManagerMinerUnstake, error) {
	event := new(GPUManagerMinerUnstake)
	if err := _GPUManager.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the GPUManager contract.
type GPUManagerModelMinimumFeeUpdateIterator struct {
	Event *GPUManagerModelMinimumFeeUpdate // Event containing the contract specifics and raw log

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
func (it *GPUManagerModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerModelMinimumFeeUpdate)
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
		it.Event = new(GPUManagerModelMinimumFeeUpdate)
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
func (it *GPUManagerModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the GPUManager contract.
type GPUManagerModelMinimumFeeUpdate struct {
	ModelId    uint32
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x32fdbd4cff3135e1bb0ae98bb593ee0c78a48a5e92e80ccf8a8ab6e72b21ffb9.
//
// Solidity: event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee)
func (_GPUManager *GPUManagerFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, modelId []uint32) (*GPUManagerModelMinimumFeeUpdateIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerModelMinimumFeeUpdateIterator{contract: _GPUManager.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x32fdbd4cff3135e1bb0ae98bb593ee0c78a48a5e92e80ccf8a8ab6e72b21ffb9.
//
// Solidity: event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee)
func (_GPUManager *GPUManagerFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *GPUManagerModelMinimumFeeUpdate, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerModelMinimumFeeUpdate)
				if err := _GPUManager.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*GPUManagerModelMinimumFeeUpdate, error) {
	event := new(GPUManagerModelMinimumFeeUpdate)
	if err := _GPUManager.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the GPUManager contract.
type GPUManagerModelRegistrationIterator struct {
	Event *GPUManagerModelRegistration // Event containing the contract specifics and raw log

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
func (it *GPUManagerModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerModelRegistration)
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
		it.Event = new(GPUManagerModelRegistration)
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
func (it *GPUManagerModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerModelRegistration represents a ModelRegistration event raised by the GPUManager contract.
type GPUManagerModelRegistration struct {
	ModelId    uint32
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0xbf8d4447fa6c121c179656152534cb5032c1ce50f747e90c56580bec25583d81.
//
// Solidity: event ModelRegistration(uint32 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_GPUManager *GPUManagerFilterer) FilterModelRegistration(opts *bind.FilterOpts, modelId []uint32, tier []uint16) (*GPUManagerModelRegistrationIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "ModelRegistration", modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerModelRegistrationIterator{contract: _GPUManager.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0xbf8d4447fa6c121c179656152534cb5032c1ce50f747e90c56580bec25583d81.
//
// Solidity: event ModelRegistration(uint32 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_GPUManager *GPUManagerFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *GPUManagerModelRegistration, modelId []uint32, tier []uint16) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "ModelRegistration", modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerModelRegistration)
				if err := _GPUManager.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseModelRegistration(log types.Log) (*GPUManagerModelRegistration, error) {
	event := new(GPUManagerModelRegistration)
	if err := _GPUManager.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the GPUManager contract.
type GPUManagerModelTierUpdateIterator struct {
	Event *GPUManagerModelTierUpdate // Event containing the contract specifics and raw log

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
func (it *GPUManagerModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerModelTierUpdate)
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
		it.Event = new(GPUManagerModelTierUpdate)
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
func (it *GPUManagerModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerModelTierUpdate represents a ModelTierUpdate event raised by the GPUManager contract.
type GPUManagerModelTierUpdate struct {
	ModelId uint32
	Tier    uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x4ecbcd19e308970fa368644f223de37bf9800e203349b5873d83970277c30356.
//
// Solidity: event ModelTierUpdate(uint32 indexed modelId, uint32 tier)
func (_GPUManager *GPUManagerFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, modelId []uint32) (*GPUManagerModelTierUpdateIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "ModelTierUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerModelTierUpdateIterator{contract: _GPUManager.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x4ecbcd19e308970fa368644f223de37bf9800e203349b5873d83970277c30356.
//
// Solidity: event ModelTierUpdate(uint32 indexed modelId, uint32 tier)
func (_GPUManager *GPUManagerFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *GPUManagerModelTierUpdate, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "ModelTierUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerModelTierUpdate)
				if err := _GPUManager.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseModelTierUpdate(log types.Log) (*GPUManagerModelTierUpdate, error) {
	event := new(GPUManagerModelTierUpdate)
	if err := _GPUManager.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the GPUManager contract.
type GPUManagerModelUnregistrationIterator struct {
	Event *GPUManagerModelUnregistration // Event containing the contract specifics and raw log

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
func (it *GPUManagerModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerModelUnregistration)
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
		it.Event = new(GPUManagerModelUnregistration)
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
func (it *GPUManagerModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerModelUnregistration represents a ModelUnregistration event raised by the GPUManager contract.
type GPUManagerModelUnregistration struct {
	ModelId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x543408e7ce45c07531e494b8909d4d1b9dea7a8d8f5907b4673949a90fc56ba2.
//
// Solidity: event ModelUnregistration(uint32 indexed modelId)
func (_GPUManager *GPUManagerFilterer) FilterModelUnregistration(opts *bind.FilterOpts, modelId []uint32) (*GPUManagerModelUnregistrationIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "ModelUnregistration", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerModelUnregistrationIterator{contract: _GPUManager.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x543408e7ce45c07531e494b8909d4d1b9dea7a8d8f5907b4673949a90fc56ba2.
//
// Solidity: event ModelUnregistration(uint32 indexed modelId)
func (_GPUManager *GPUManagerFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *GPUManagerModelUnregistration, modelId []uint32) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "ModelUnregistration", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerModelUnregistration)
				if err := _GPUManager.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseModelUnregistration(log types.Log) (*GPUManagerModelUnregistration, error) {
	event := new(GPUManagerModelUnregistration)
	if err := _GPUManager.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GPUManager contract.
type GPUManagerOwnershipTransferredIterator struct {
	Event *GPUManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GPUManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerOwnershipTransferred)
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
		it.Event = new(GPUManagerOwnershipTransferred)
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
func (it *GPUManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerOwnershipTransferred represents a OwnershipTransferred event raised by the GPUManager contract.
type GPUManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPUManager *GPUManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GPUManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerOwnershipTransferredIterator{contract: _GPUManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPUManager *GPUManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GPUManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerOwnershipTransferred)
				if err := _GPUManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseOwnershipTransferred(log types.Log) (*GPUManagerOwnershipTransferred, error) {
	event := new(GPUManagerOwnershipTransferred)
	if err := _GPUManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GPUManager contract.
type GPUManagerPausedIterator struct {
	Event *GPUManagerPaused // Event containing the contract specifics and raw log

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
func (it *GPUManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerPaused)
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
		it.Event = new(GPUManagerPaused)
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
func (it *GPUManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerPaused represents a Paused event raised by the GPUManager contract.
type GPUManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GPUManager *GPUManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*GPUManagerPausedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GPUManagerPausedIterator{contract: _GPUManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GPUManager *GPUManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GPUManagerPaused) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerPaused)
				if err := _GPUManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParsePaused(log types.Log) (*GPUManagerPaused, error) {
	event := new(GPUManagerPaused)
	if err := _GPUManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the GPUManager contract.
type GPUManagerPenaltyDurationUpdatedIterator struct {
	Event *GPUManagerPenaltyDurationUpdated // Event containing the contract specifics and raw log

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
func (it *GPUManagerPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerPenaltyDurationUpdated)
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
		it.Event = new(GPUManagerPenaltyDurationUpdated)
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
func (it *GPUManagerPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the GPUManager contract.
type GPUManagerPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_GPUManager *GPUManagerFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*GPUManagerPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &GPUManagerPenaltyDurationUpdatedIterator{contract: _GPUManager.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_GPUManager *GPUManagerFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *GPUManagerPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerPenaltyDurationUpdated)
				if err := _GPUManager.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParsePenaltyDurationUpdated(log types.Log) (*GPUManagerPenaltyDurationUpdated, error) {
	event := new(GPUManagerPenaltyDurationUpdated)
	if err := _GPUManager.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the GPUManager contract.
type GPUManagerRestakeIterator struct {
	Event *GPUManagerRestake // Event containing the contract specifics and raw log

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
func (it *GPUManagerRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerRestake)
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
		it.Event = new(GPUManagerRestake)
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
func (it *GPUManagerRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerRestake represents a Restake event raised by the GPUManager contract.
type GPUManagerRestake struct {
	Miner   common.Address
	ModelId uint32
	Restake *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0xd71961af2f46a633dc473cc0dda9e08783282fdb38c8f90482a143eb63b039e0.
//
// Solidity: event Restake(address indexed miner, uint32 indexed modelId, uint256 restake)
func (_GPUManager *GPUManagerFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, modelId []uint32) (*GPUManagerRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "Restake", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerRestakeIterator{contract: _GPUManager.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0xd71961af2f46a633dc473cc0dda9e08783282fdb38c8f90482a143eb63b039e0.
//
// Solidity: event Restake(address indexed miner, uint32 indexed modelId, uint256 restake)
func (_GPUManager *GPUManagerFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *GPUManagerRestake, miner []common.Address, modelId []uint32) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "Restake", minerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerRestake)
				if err := _GPUManager.contract.UnpackLog(event, "Restake", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseRestake(log types.Log) (*GPUManagerRestake, error) {
	event := new(GPUManagerRestake)
	if err := _GPUManager.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the GPUManager contract.
type GPUManagerRewardClaimIterator struct {
	Event *GPUManagerRewardClaim // Event containing the contract specifics and raw log

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
func (it *GPUManagerRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerRewardClaim)
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
		it.Event = new(GPUManagerRewardClaim)
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
func (it *GPUManagerRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerRewardClaim represents a RewardClaim event raised by the GPUManager contract.
type GPUManagerRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_GPUManager *GPUManagerFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*GPUManagerRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &GPUManagerRewardClaimIterator{contract: _GPUManager.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_GPUManager *GPUManagerFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *GPUManagerRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerRewardClaim)
				if err := _GPUManager.contract.UnpackLog(event, "RewardClaim", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseRewardClaim(log types.Log) (*GPUManagerRewardClaim, error) {
	event := new(GPUManagerRewardClaim)
	if err := _GPUManager.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the GPUManager contract.
type GPUManagerRewardPerEpochIterator struct {
	Event *GPUManagerRewardPerEpoch // Event containing the contract specifics and raw log

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
func (it *GPUManagerRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerRewardPerEpoch)
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
		it.Event = new(GPUManagerRewardPerEpoch)
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
func (it *GPUManagerRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerRewardPerEpoch represents a RewardPerEpoch event raised by the GPUManager contract.
type GPUManagerRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_GPUManager *GPUManagerFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*GPUManagerRewardPerEpochIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &GPUManagerRewardPerEpochIterator{contract: _GPUManager.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_GPUManager *GPUManagerFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *GPUManagerRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerRewardPerEpoch)
				if err := _GPUManager.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseRewardPerEpoch(log types.Log) (*GPUManagerRewardPerEpoch, error) {
	event := new(GPUManagerRewardPerEpoch)
	if err := _GPUManager.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GPUManager contract.
type GPUManagerUnpausedIterator struct {
	Event *GPUManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *GPUManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerUnpaused)
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
		it.Event = new(GPUManagerUnpaused)
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
func (it *GPUManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerUnpaused represents a Unpaused event raised by the GPUManager contract.
type GPUManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GPUManager *GPUManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GPUManagerUnpausedIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GPUManagerUnpausedIterator{contract: _GPUManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GPUManager *GPUManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GPUManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerUnpaused)
				if err := _GPUManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseUnpaused(log types.Log) (*GPUManagerUnpaused, error) {
	event := new(GPUManagerUnpaused)
	if err := _GPUManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPUManagerUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the GPUManager contract.
type GPUManagerUnstakeDelayTimeIterator struct {
	Event *GPUManagerUnstakeDelayTime // Event containing the contract specifics and raw log

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
func (it *GPUManagerUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPUManagerUnstakeDelayTime)
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
		it.Event = new(GPUManagerUnstakeDelayTime)
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
func (it *GPUManagerUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPUManagerUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPUManagerUnstakeDelayTime represents a UnstakeDelayTime event raised by the GPUManager contract.
type GPUManagerUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_GPUManager *GPUManagerFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*GPUManagerUnstakeDelayTimeIterator, error) {

	logs, sub, err := _GPUManager.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &GPUManagerUnstakeDelayTimeIterator{contract: _GPUManager.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_GPUManager *GPUManagerFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *GPUManagerUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _GPUManager.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPUManagerUnstakeDelayTime)
				if err := _GPUManager.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
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
func (_GPUManager *GPUManagerFilterer) ParseUnstakeDelayTime(log types.Log) (*GPUManagerUnstakeDelayTime, error) {
	event := new(GPUManagerUnstakeDelayTime)
	if err := _GPUManager.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
