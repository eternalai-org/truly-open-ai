// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prompt_scheduler

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

// ISchedulerInference is an auto generated low-level Go binding around an user-defined struct.
type ISchedulerInference struct {
	Value          *big.Int
	ModelId        uint32
	SubmitTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Input          []byte
	Output         []byte
}

// PromptSchedulerMetaData contains all meta data concerning the PromptScheduler contract.
var PromptSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssignedWorker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitTimeout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"AppendToBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_batchPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gpuManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_inferenceCounter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastBatchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerValidatorFeeRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_submitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"}],\"name\":\"getBatchInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getInferenceByMiner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIScheduler.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minerRequirement_\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"submitDuration_\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minerValidatorFeeRatio_\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"batchPeriod_\",\"type\":\"uint40\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"solution\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PromptSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use PromptSchedulerMetaData.ABI instead.
var PromptSchedulerABI = PromptSchedulerMetaData.ABI

// PromptScheduler is an auto generated Go binding around an Ethereum contract.
type PromptScheduler struct {
	PromptSchedulerCaller     // Read-only binding to the contract
	PromptSchedulerTransactor // Write-only binding to the contract
	PromptSchedulerFilterer   // Log filterer for contract events
}

// PromptSchedulerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PromptSchedulerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PromptSchedulerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PromptSchedulerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PromptSchedulerSession struct {
	Contract     *PromptScheduler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PromptSchedulerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PromptSchedulerCallerSession struct {
	Contract *PromptSchedulerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PromptSchedulerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PromptSchedulerTransactorSession struct {
	Contract     *PromptSchedulerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PromptSchedulerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PromptSchedulerRaw struct {
	Contract *PromptScheduler // Generic contract binding to access the raw methods on
}

// PromptSchedulerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PromptSchedulerCallerRaw struct {
	Contract *PromptSchedulerCaller // Generic read-only contract binding to access the raw methods on
}

// PromptSchedulerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PromptSchedulerTransactorRaw struct {
	Contract *PromptSchedulerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPromptScheduler creates a new instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptScheduler(address common.Address, backend bind.ContractBackend) (*PromptScheduler, error) {
	contract, err := bindPromptScheduler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PromptScheduler{PromptSchedulerCaller: PromptSchedulerCaller{contract: contract}, PromptSchedulerTransactor: PromptSchedulerTransactor{contract: contract}, PromptSchedulerFilterer: PromptSchedulerFilterer{contract: contract}}, nil
}

// NewPromptSchedulerCaller creates a new read-only instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerCaller(address common.Address, caller bind.ContractCaller) (*PromptSchedulerCaller, error) {
	contract, err := bindPromptScheduler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerCaller{contract: contract}, nil
}

// NewPromptSchedulerTransactor creates a new write-only instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerTransactor(address common.Address, transactor bind.ContractTransactor) (*PromptSchedulerTransactor, error) {
	contract, err := bindPromptScheduler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerTransactor{contract: contract}, nil
}

// NewPromptSchedulerFilterer creates a new log filterer instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerFilterer(address common.Address, filterer bind.ContractFilterer) (*PromptSchedulerFilterer, error) {
	contract, err := bindPromptScheduler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerFilterer{contract: contract}, nil
}

// bindPromptScheduler binds a generic wrapper to an already deployed contract.
func bindPromptScheduler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PromptSchedulerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptScheduler *PromptSchedulerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptScheduler.Contract.PromptSchedulerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptScheduler *PromptSchedulerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.Contract.PromptSchedulerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptScheduler *PromptSchedulerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptScheduler.Contract.PromptSchedulerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptScheduler *PromptSchedulerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptScheduler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptScheduler *PromptSchedulerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptScheduler *PromptSchedulerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptScheduler.Contract.contract.Transact(opts, method, params...)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCaller) BatchPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_batchPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerSession) BatchPeriod() (*big.Int, error) {
	return _PromptScheduler.Contract.BatchPeriod(&_PromptScheduler.CallOpts)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCallerSession) BatchPeriod() (*big.Int, error) {
	return _PromptScheduler.Contract.BatchPeriod(&_PromptScheduler.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) GpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_gpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) GpuManager() (common.Address, error) {
	return _PromptScheduler.Contract.GpuManager(&_PromptScheduler.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) GpuManager() (common.Address, error) {
	return _PromptScheduler.Contract.GpuManager(&_PromptScheduler.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerCaller) InferenceCounter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_inferenceCounter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) InferenceCounter() (uint64, error) {
	return _PromptScheduler.Contract.InferenceCounter(&_PromptScheduler.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerCallerSession) InferenceCounter() (uint64, error) {
	return _PromptScheduler.Contract.InferenceCounter(&_PromptScheduler.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCaller) LastBatchTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_lastBatchTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerSession) LastBatchTimestamp() (*big.Int, error) {
	return _PromptScheduler.Contract.LastBatchTimestamp(&_PromptScheduler.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCallerSession) LastBatchTimestamp() (*big.Int, error) {
	return _PromptScheduler.Contract.LastBatchTimestamp(&_PromptScheduler.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerSession) MinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.MinerRequirement(&_PromptScheduler.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCallerSession) MinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.MinerRequirement(&_PromptScheduler.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerCaller) MinerValidatorFeeRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_minerValidatorFeeRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerSession) MinerValidatorFeeRatio() (uint16, error) {
	return _PromptScheduler.Contract.MinerValidatorFeeRatio(&_PromptScheduler.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerCallerSession) MinerValidatorFeeRatio() (uint16, error) {
	return _PromptScheduler.Contract.MinerValidatorFeeRatio(&_PromptScheduler.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerCaller) SubmitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_submitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerSession) SubmitDuration() (*big.Int, error) {
	return _PromptScheduler.Contract.SubmitDuration(&_PromptScheduler.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerCallerSession) SubmitDuration() (*big.Int, error) {
	return _PromptScheduler.Contract.SubmitDuration(&_PromptScheduler.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) WEAIToken() (common.Address, error) {
	return _PromptScheduler.Contract.WEAIToken(&_PromptScheduler.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) WEAIToken() (common.Address, error) {
	return _PromptScheduler.Contract.WEAIToken(&_PromptScheduler.CallOpts)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerCaller) GetBatchInfo(opts *bind.CallOpts, modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getBatchInfo", modelId, batchId)

	if err != nil {
		return *new(*big.Int), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _PromptScheduler.Contract.GetBatchInfo(&_PromptScheduler.CallOpts, modelId, batchId)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerCallerSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _PromptScheduler.Contract.GetBatchInfo(&_PromptScheduler.CallOpts, modelId, batchId)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerCaller) GetInferenceByMiner(opts *bind.CallOpts, miner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getInferenceByMiner", miner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _PromptScheduler.Contract.GetInferenceByMiner(&_PromptScheduler.CallOpts, miner)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerCallerSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _PromptScheduler.Contract.GetInferenceByMiner(&_PromptScheduler.CallOpts, miner)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerCaller) GetInferenceInfo(opts *bind.CallOpts, inferId uint64) (ISchedulerInference, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getInferenceInfo", inferId)

	if err != nil {
		return *new(ISchedulerInference), err
	}

	out0 := *abi.ConvertType(out[0], new(ISchedulerInference)).(*ISchedulerInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _PromptScheduler.Contract.GetInferenceInfo(&_PromptScheduler.CallOpts, inferId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerCallerSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _PromptScheduler.Contract.GetInferenceInfo(&_PromptScheduler.CallOpts, inferId)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCaller) GetMinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getMinerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerSession) GetMinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.GetMinerRequirement(&_PromptScheduler.CallOpts)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCallerSession) GetMinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.GetMinerRequirement(&_PromptScheduler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) Owner() (common.Address, error) {
	return _PromptScheduler.Contract.Owner(&_PromptScheduler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) Owner() (common.Address, error) {
	return _PromptScheduler.Contract.Owner(&_PromptScheduler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerSession) Paused() (bool, error) {
	return _PromptScheduler.Contract.Paused(&_PromptScheduler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerCallerSession) Paused() (bool, error) {
	return _PromptScheduler.Contract.Paused(&_PromptScheduler.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerSession) Version() (string, error) {
	return _PromptScheduler.Contract.Version(&_PromptScheduler.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerCallerSession) Version() (string, error) {
	return _PromptScheduler.Contract.Version(&_PromptScheduler.CallOpts)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactor) Infer(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "infer", modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer(&_PromptScheduler.TransactOpts, modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactorSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer(&_PromptScheduler.TransactOpts, modelId, input, creator, flag)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactor) Infer0(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "infer0", modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer0(&_PromptScheduler.TransactOpts, modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactorSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer0(&_PromptScheduler.TransactOpts, modelId, input, creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerTransactor) Initialize(opts *bind.TransactOpts, wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "initialize", wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Initialize(&_PromptScheduler.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Initialize(&_PromptScheduler.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerSession) Pause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Pause(&_PromptScheduler.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Pause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Pause(&_PromptScheduler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptScheduler.Contract.RenounceOwnership(&_PromptScheduler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptScheduler.Contract.RenounceOwnership(&_PromptScheduler.TransactOpts)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerTransactor) SetWEAIAddress(opts *bind.TransactOpts, wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "setWEAIAddress", wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetWEAIAddress(&_PromptScheduler.TransactOpts, wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetWEAIAddress(&_PromptScheduler.TransactOpts, wEAIToken)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerTransactor) SubmitSolution(opts *bind.TransactOpts, inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "submitSolution", inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SubmitSolution(&_PromptScheduler.TransactOpts, inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SubmitSolution(&_PromptScheduler.TransactOpts, inferId, solution)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.TransferOwnership(&_PromptScheduler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.TransferOwnership(&_PromptScheduler.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerSession) Unpause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Unpause(&_PromptScheduler.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Unpause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Unpause(&_PromptScheduler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerSession) Receive() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Receive(&_PromptScheduler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Receive() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Receive(&_PromptScheduler.TransactOpts)
}

// PromptSchedulerAppendToBatchIterator is returned from FilterAppendToBatch and is used to iterate over the raw logs and unpacked data for AppendToBatch events raised by the PromptScheduler contract.
type PromptSchedulerAppendToBatchIterator struct {
	Event *PromptSchedulerAppendToBatch // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerAppendToBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerAppendToBatch)
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
		it.Event = new(PromptSchedulerAppendToBatch)
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
func (it *PromptSchedulerAppendToBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerAppendToBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerAppendToBatch represents a AppendToBatch event raised by the PromptScheduler contract.
type PromptSchedulerAppendToBatch struct {
	BatchId uint64
	ModelId uint32
	InferId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAppendToBatch is a free log retrieval operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) FilterAppendToBatch(opts *bind.FilterOpts, batchId []uint64, modelId []uint32, inferId []uint64) (*PromptSchedulerAppendToBatchIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerAppendToBatchIterator{contract: _PromptScheduler.contract, event: "AppendToBatch", logs: logs, sub: sub}, nil
}

// WatchAppendToBatch is a free log subscription operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) WatchAppendToBatch(opts *bind.WatchOpts, sink chan<- *PromptSchedulerAppendToBatch, batchId []uint64, modelId []uint32, inferId []uint64) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerAppendToBatch)
				if err := _PromptScheduler.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
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

// ParseAppendToBatch is a log parse operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) ParseAppendToBatch(log types.Log) (*PromptSchedulerAppendToBatch, error) {
	event := new(PromptSchedulerAppendToBatch)
	if err := _PromptScheduler.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the PromptScheduler contract.
type PromptSchedulerInferenceStatusUpdateIterator struct {
	Event *PromptSchedulerInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerInferenceStatusUpdate)
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
		it.Event = new(PromptSchedulerInferenceStatusUpdate)
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
func (it *PromptSchedulerInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the PromptScheduler contract.
type PromptSchedulerInferenceStatusUpdate struct {
	InferenceId uint64
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []uint64) (*PromptSchedulerInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerInferenceStatusUpdateIterator{contract: _PromptScheduler.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *PromptSchedulerInferenceStatusUpdate, inferenceId []uint64) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerInferenceStatusUpdate)
				if err := _PromptScheduler.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) ParseInferenceStatusUpdate(log types.Log) (*PromptSchedulerInferenceStatusUpdate, error) {
	event := new(PromptSchedulerInferenceStatusUpdate)
	if err := _PromptScheduler.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PromptScheduler contract.
type PromptSchedulerInitializedIterator struct {
	Event *PromptSchedulerInitialized // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerInitialized)
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
		it.Event = new(PromptSchedulerInitialized)
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
func (it *PromptSchedulerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerInitialized represents a Initialized event raised by the PromptScheduler contract.
type PromptSchedulerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PromptScheduler *PromptSchedulerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PromptSchedulerInitializedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerInitializedIterator{contract: _PromptScheduler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PromptScheduler *PromptSchedulerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PromptSchedulerInitialized) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerInitialized)
				if err := _PromptScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseInitialized(log types.Log) (*PromptSchedulerInitialized, error) {
	event := new(PromptSchedulerInitialized)
	if err := _PromptScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the PromptScheduler contract.
type PromptSchedulerNewAssignmentIterator struct {
	Event *PromptSchedulerNewAssignment // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerNewAssignment)
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
		it.Event = new(PromptSchedulerNewAssignment)
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
func (it *PromptSchedulerNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerNewAssignment represents a NewAssignment event raised by the PromptScheduler contract.
type PromptSchedulerNewAssignment struct {
	InferenceId uint64
	Miner       common.Address
	ExpiredAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) FilterNewAssignment(opts *bind.FilterOpts, inferenceId []uint64, miner []common.Address) (*PromptSchedulerNewAssignmentIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerNewAssignmentIterator{contract: _PromptScheduler.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *PromptSchedulerNewAssignment, inferenceId []uint64, miner []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerNewAssignment)
				if err := _PromptScheduler.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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

// ParseNewAssignment is a log parse operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) ParseNewAssignment(log types.Log) (*PromptSchedulerNewAssignment, error) {
	event := new(PromptSchedulerNewAssignment)
	if err := _PromptScheduler.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the PromptScheduler contract.
type PromptSchedulerNewInferenceIterator struct {
	Event *PromptSchedulerNewInference // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerNewInference)
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
		it.Event = new(PromptSchedulerNewInference)
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
func (it *PromptSchedulerNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerNewInference represents a NewInference event raised by the PromptScheduler contract.
type PromptSchedulerNewInference struct {
	InferenceId uint64
	Creator     common.Address
	ModelId     uint32
	Value       *big.Int
	Input       []byte
	Flag        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []uint64, creator []common.Address, modelId []uint32) (*PromptSchedulerNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerNewInferenceIterator{contract: _PromptScheduler.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *PromptSchedulerNewInference, inferenceId []uint64, creator []common.Address, modelId []uint32) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerNewInference)
				if err := _PromptScheduler.contract.UnpackLog(event, "NewInference", log); err != nil {
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

// ParseNewInference is a log parse operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) ParseNewInference(log types.Log) (*PromptSchedulerNewInference, error) {
	event := new(PromptSchedulerNewInference)
	if err := _PromptScheduler.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PromptScheduler contract.
type PromptSchedulerOwnershipTransferredIterator struct {
	Event *PromptSchedulerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerOwnershipTransferred)
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
		it.Event = new(PromptSchedulerOwnershipTransferred)
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
func (it *PromptSchedulerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerOwnershipTransferred represents a OwnershipTransferred event raised by the PromptScheduler contract.
type PromptSchedulerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptScheduler *PromptSchedulerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PromptSchedulerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerOwnershipTransferredIterator{contract: _PromptScheduler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptScheduler *PromptSchedulerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PromptSchedulerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerOwnershipTransferred)
				if err := _PromptScheduler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseOwnershipTransferred(log types.Log) (*PromptSchedulerOwnershipTransferred, error) {
	event := new(PromptSchedulerOwnershipTransferred)
	if err := _PromptScheduler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PromptScheduler contract.
type PromptSchedulerPausedIterator struct {
	Event *PromptSchedulerPaused // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerPaused)
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
		it.Event = new(PromptSchedulerPaused)
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
func (it *PromptSchedulerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerPaused represents a Paused event raised by the PromptScheduler contract.
type PromptSchedulerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) FilterPaused(opts *bind.FilterOpts) (*PromptSchedulerPausedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerPausedIterator{contract: _PromptScheduler.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PromptSchedulerPaused) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerPaused)
				if err := _PromptScheduler.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParsePaused(log types.Log) (*PromptSchedulerPaused, error) {
	event := new(PromptSchedulerPaused)
	if err := _PromptScheduler.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the PromptScheduler contract.
type PromptSchedulerSolutionSubmissionIterator struct {
	Event *PromptSchedulerSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerSolutionSubmission)
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
		it.Event = new(PromptSchedulerSolutionSubmission)
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
func (it *PromptSchedulerSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerSolutionSubmission represents a SolutionSubmission event raised by the PromptScheduler contract.
type PromptSchedulerSolutionSubmission struct {
	Miner   common.Address
	InferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, inferId []*big.Int) (*PromptSchedulerSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerSolutionSubmissionIterator{contract: _PromptScheduler.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *PromptSchedulerSolutionSubmission, miner []common.Address, inferId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerSolutionSubmission)
				if err := _PromptScheduler.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) ParseSolutionSubmission(log types.Log) (*PromptSchedulerSolutionSubmission, error) {
	event := new(PromptSchedulerSolutionSubmission)
	if err := _PromptScheduler.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the PromptScheduler contract.
type PromptSchedulerStreamedDataIterator struct {
	Event *PromptSchedulerStreamedData // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerStreamedData)
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
		it.Event = new(PromptSchedulerStreamedData)
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
func (it *PromptSchedulerStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerStreamedData represents a StreamedData event raised by the PromptScheduler contract.
type PromptSchedulerStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_PromptScheduler *PromptSchedulerFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*PromptSchedulerStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerStreamedDataIterator{contract: _PromptScheduler.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_PromptScheduler *PromptSchedulerFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *PromptSchedulerStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerStreamedData)
				if err := _PromptScheduler.contract.UnpackLog(event, "StreamedData", log); err != nil {
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

// ParseStreamedData is a log parse operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_PromptScheduler *PromptSchedulerFilterer) ParseStreamedData(log types.Log) (*PromptSchedulerStreamedData, error) {
	event := new(PromptSchedulerStreamedData)
	if err := _PromptScheduler.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PromptScheduler contract.
type PromptSchedulerUnpausedIterator struct {
	Event *PromptSchedulerUnpaused // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerUnpaused)
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
		it.Event = new(PromptSchedulerUnpaused)
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
func (it *PromptSchedulerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerUnpaused represents a Unpaused event raised by the PromptScheduler contract.
type PromptSchedulerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PromptSchedulerUnpausedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerUnpausedIterator{contract: _PromptScheduler.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PromptSchedulerUnpaused) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerUnpaused)
				if err := _PromptScheduler.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseUnpaused(log types.Log) (*PromptSchedulerUnpaused, error) {
	event := new(PromptSchedulerUnpaused)
	if err := _PromptScheduler.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
