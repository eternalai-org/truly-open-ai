// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// WorkerhubContractMetaData contains all meta data concerning the WorkerhubContract contract.
var WorkerhubContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssignedWorker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitTimeout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"AppendToBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_batchPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gpuManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_inferenceCounter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastBatchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerValidatorFeeRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_submitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"}],\"name\":\"getBatchInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getInferenceByMiner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIScheduler.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minerRequirement_\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"submitDuration_\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minerValidatorFeeRatio_\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"batchPeriod_\",\"type\":\"uint40\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"submitDuration\",\"type\":\"uint40\"}],\"name\":\"setSubmitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"solution\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WorkerhubContractABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerhubContractMetaData.ABI instead.
var WorkerhubContractABI = WorkerhubContractMetaData.ABI

// WorkerhubContract is an auto generated Go binding around an Ethereum contract.
type WorkerhubContract struct {
	WorkerhubContractCaller     // Read-only binding to the contract
	WorkerhubContractTransactor // Write-only binding to the contract
	WorkerhubContractFilterer   // Log filterer for contract events
}

// WorkerhubContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerhubContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerhubContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerhubContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerhubContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerhubContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerhubContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerhubContractSession struct {
	Contract     *WorkerhubContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WorkerhubContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerhubContractCallerSession struct {
	Contract *WorkerhubContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WorkerhubContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerhubContractTransactorSession struct {
	Contract     *WorkerhubContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WorkerhubContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerhubContractRaw struct {
	Contract *WorkerhubContract // Generic contract binding to access the raw methods on
}

// WorkerhubContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerhubContractCallerRaw struct {
	Contract *WorkerhubContractCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerhubContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerhubContractTransactorRaw struct {
	Contract *WorkerhubContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerhubContract creates a new instance of WorkerhubContract, bound to a specific deployed contract.
func NewWorkerhubContract(address common.Address, backend bind.ContractBackend) (*WorkerhubContract, error) {
	contract, err := bindWorkerhubContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContract{WorkerhubContractCaller: WorkerhubContractCaller{contract: contract}, WorkerhubContractTransactor: WorkerhubContractTransactor{contract: contract}, WorkerhubContractFilterer: WorkerhubContractFilterer{contract: contract}}, nil
}

// NewWorkerhubContractCaller creates a new read-only instance of WorkerhubContract, bound to a specific deployed contract.
func NewWorkerhubContractCaller(address common.Address, caller bind.ContractCaller) (*WorkerhubContractCaller, error) {
	contract, err := bindWorkerhubContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractCaller{contract: contract}, nil
}

// NewWorkerhubContractTransactor creates a new write-only instance of WorkerhubContract, bound to a specific deployed contract.
func NewWorkerhubContractTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerhubContractTransactor, error) {
	contract, err := bindWorkerhubContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractTransactor{contract: contract}, nil
}

// NewWorkerhubContractFilterer creates a new log filterer instance of WorkerhubContract, bound to a specific deployed contract.
func NewWorkerhubContractFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerhubContractFilterer, error) {
	contract, err := bindWorkerhubContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractFilterer{contract: contract}, nil
}

// bindWorkerhubContract binds a generic wrapper to an already deployed contract.
func bindWorkerhubContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerhubContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerhubContract *WorkerhubContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerhubContract.Contract.WorkerhubContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerhubContract *WorkerhubContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.WorkerhubContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerhubContract *WorkerhubContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.WorkerhubContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerhubContract *WorkerhubContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerhubContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerhubContract *WorkerhubContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerhubContract *WorkerhubContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.contract.Transact(opts, method, params...)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractCaller) BatchPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_batchPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractSession) BatchPeriod() (*big.Int, error) {
	return _WorkerhubContract.Contract.BatchPeriod(&_WorkerhubContract.CallOpts)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractCallerSession) BatchPeriod() (*big.Int, error) {
	return _WorkerhubContract.Contract.BatchPeriod(&_WorkerhubContract.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_WorkerhubContract *WorkerhubContractCaller) GpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_gpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_WorkerhubContract *WorkerhubContractSession) GpuManager() (common.Address, error) {
	return _WorkerhubContract.Contract.GpuManager(&_WorkerhubContract.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_WorkerhubContract *WorkerhubContractCallerSession) GpuManager() (common.Address, error) {
	return _WorkerhubContract.Contract.GpuManager(&_WorkerhubContract.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_WorkerhubContract *WorkerhubContractCaller) InferenceCounter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_inferenceCounter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_WorkerhubContract *WorkerhubContractSession) InferenceCounter() (uint64, error) {
	return _WorkerhubContract.Contract.InferenceCounter(&_WorkerhubContract.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_WorkerhubContract *WorkerhubContractCallerSession) InferenceCounter() (uint64, error) {
	return _WorkerhubContract.Contract.InferenceCounter(&_WorkerhubContract.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractCaller) LastBatchTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_lastBatchTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractSession) LastBatchTimestamp() (*big.Int, error) {
	return _WorkerhubContract.Contract.LastBatchTimestamp(&_WorkerhubContract.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_WorkerhubContract *WorkerhubContractCallerSession) LastBatchTimestamp() (*big.Int, error) {
	return _WorkerhubContract.Contract.LastBatchTimestamp(&_WorkerhubContract.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractSession) MinerRequirement() (uint8, error) {
	return _WorkerhubContract.Contract.MinerRequirement(&_WorkerhubContract.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractCallerSession) MinerRequirement() (uint8, error) {
	return _WorkerhubContract.Contract.MinerRequirement(&_WorkerhubContract.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_WorkerhubContract *WorkerhubContractCaller) MinerValidatorFeeRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_minerValidatorFeeRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_WorkerhubContract *WorkerhubContractSession) MinerValidatorFeeRatio() (uint16, error) {
	return _WorkerhubContract.Contract.MinerValidatorFeeRatio(&_WorkerhubContract.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_WorkerhubContract *WorkerhubContractCallerSession) MinerValidatorFeeRatio() (uint16, error) {
	return _WorkerhubContract.Contract.MinerValidatorFeeRatio(&_WorkerhubContract.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_WorkerhubContract *WorkerhubContractCaller) SubmitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_submitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_WorkerhubContract *WorkerhubContractSession) SubmitDuration() (*big.Int, error) {
	return _WorkerhubContract.Contract.SubmitDuration(&_WorkerhubContract.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_WorkerhubContract *WorkerhubContractCallerSession) SubmitDuration() (*big.Int, error) {
	return _WorkerhubContract.Contract.SubmitDuration(&_WorkerhubContract.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_WorkerhubContract *WorkerhubContractCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_WorkerhubContract *WorkerhubContractSession) WEAIToken() (common.Address, error) {
	return _WorkerhubContract.Contract.WEAIToken(&_WorkerhubContract.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_WorkerhubContract *WorkerhubContractCallerSession) WEAIToken() (common.Address, error) {
	return _WorkerhubContract.Contract.WEAIToken(&_WorkerhubContract.CallOpts)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_WorkerhubContract *WorkerhubContractCaller) GetBatchInfo(opts *bind.CallOpts, modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "getBatchInfo", modelId, batchId)

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
func (_WorkerhubContract *WorkerhubContractSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _WorkerhubContract.Contract.GetBatchInfo(&_WorkerhubContract.CallOpts, modelId, batchId)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_WorkerhubContract *WorkerhubContractCallerSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _WorkerhubContract.Contract.GetBatchInfo(&_WorkerhubContract.CallOpts, modelId, batchId)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_WorkerhubContract *WorkerhubContractCaller) GetInferenceByMiner(opts *bind.CallOpts, miner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "getInferenceByMiner", miner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_WorkerhubContract *WorkerhubContractSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _WorkerhubContract.Contract.GetInferenceByMiner(&_WorkerhubContract.CallOpts, miner)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_WorkerhubContract *WorkerhubContractCallerSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _WorkerhubContract.Contract.GetInferenceByMiner(&_WorkerhubContract.CallOpts, miner)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_WorkerhubContract *WorkerhubContractCaller) GetInferenceInfo(opts *bind.CallOpts, inferId uint64) (ISchedulerInference, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "getInferenceInfo", inferId)

	if err != nil {
		return *new(ISchedulerInference), err
	}

	out0 := *abi.ConvertType(out[0], new(ISchedulerInference)).(*ISchedulerInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_WorkerhubContract *WorkerhubContractSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _WorkerhubContract.Contract.GetInferenceInfo(&_WorkerhubContract.CallOpts, inferId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_WorkerhubContract *WorkerhubContractCallerSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _WorkerhubContract.Contract.GetInferenceInfo(&_WorkerhubContract.CallOpts, inferId)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractCaller) GetMinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "getMinerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractSession) GetMinerRequirement() (uint8, error) {
	return _WorkerhubContract.Contract.GetMinerRequirement(&_WorkerhubContract.CallOpts)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_WorkerhubContract *WorkerhubContractCallerSession) GetMinerRequirement() (uint8, error) {
	return _WorkerhubContract.Contract.GetMinerRequirement(&_WorkerhubContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerhubContract *WorkerhubContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerhubContract *WorkerhubContractSession) Owner() (common.Address, error) {
	return _WorkerhubContract.Contract.Owner(&_WorkerhubContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerhubContract *WorkerhubContractCallerSession) Owner() (common.Address, error) {
	return _WorkerhubContract.Contract.Owner(&_WorkerhubContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerhubContract *WorkerhubContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerhubContract *WorkerhubContractSession) Paused() (bool, error) {
	return _WorkerhubContract.Contract.Paused(&_WorkerhubContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerhubContract *WorkerhubContractCallerSession) Paused() (bool, error) {
	return _WorkerhubContract.Contract.Paused(&_WorkerhubContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerhubContract *WorkerhubContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerhubContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerhubContract *WorkerhubContractSession) Version() (string, error) {
	return _WorkerhubContract.Contract.Version(&_WorkerhubContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerhubContract *WorkerhubContractCallerSession) Version() (string, error) {
	return _WorkerhubContract.Contract.Version(&_WorkerhubContract.CallOpts)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_WorkerhubContract *WorkerhubContractTransactor) Infer(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "infer", modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_WorkerhubContract *WorkerhubContractSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Infer(&_WorkerhubContract.TransactOpts, modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_WorkerhubContract *WorkerhubContractTransactorSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Infer(&_WorkerhubContract.TransactOpts, modelId, input, creator, flag)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_WorkerhubContract *WorkerhubContractTransactor) Infer0(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "infer0", modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_WorkerhubContract *WorkerhubContractSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Infer0(&_WorkerhubContract.TransactOpts, modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_WorkerhubContract *WorkerhubContractTransactorSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Infer0(&_WorkerhubContract.TransactOpts, modelId, input, creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_WorkerhubContract *WorkerhubContractTransactor) Initialize(opts *bind.TransactOpts, wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "initialize", wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_WorkerhubContract *WorkerhubContractSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Initialize(&_WorkerhubContract.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Initialize(&_WorkerhubContract.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerhubContract *WorkerhubContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerhubContract *WorkerhubContractSession) Pause() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Pause(&_WorkerhubContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Pause(&_WorkerhubContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerhubContract *WorkerhubContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerhubContract *WorkerhubContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.RenounceOwnership(&_WorkerhubContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.RenounceOwnership(&_WorkerhubContract.TransactOpts)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_WorkerhubContract *WorkerhubContractTransactor) SetSubmitDuration(opts *bind.TransactOpts, submitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "setSubmitDuration", submitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_WorkerhubContract *WorkerhubContractSession) SetSubmitDuration(submitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SetSubmitDuration(&_WorkerhubContract.TransactOpts, submitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) SetSubmitDuration(submitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SetSubmitDuration(&_WorkerhubContract.TransactOpts, submitDuration)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_WorkerhubContract *WorkerhubContractTransactor) SetWEAIAddress(opts *bind.TransactOpts, wEAIToken common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "setWEAIAddress", wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_WorkerhubContract *WorkerhubContractSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SetWEAIAddress(&_WorkerhubContract.TransactOpts, wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SetWEAIAddress(&_WorkerhubContract.TransactOpts, wEAIToken)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_WorkerhubContract *WorkerhubContractTransactor) SubmitSolution(opts *bind.TransactOpts, inferId uint64, solution []byte) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "submitSolution", inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_WorkerhubContract *WorkerhubContractSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SubmitSolution(&_WorkerhubContract.TransactOpts, inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.SubmitSolution(&_WorkerhubContract.TransactOpts, inferId, solution)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerhubContract *WorkerhubContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerhubContract *WorkerhubContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.TransferOwnership(&_WorkerhubContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerhubContract.Contract.TransferOwnership(&_WorkerhubContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerhubContract *WorkerhubContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerhubContract *WorkerhubContractSession) Unpause() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Unpause(&_WorkerhubContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Unpause(&_WorkerhubContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerhubContract *WorkerhubContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerhubContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerhubContract *WorkerhubContractSession) Receive() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Receive(&_WorkerhubContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerhubContract *WorkerhubContractTransactorSession) Receive() (*types.Transaction, error) {
	return _WorkerhubContract.Contract.Receive(&_WorkerhubContract.TransactOpts)
}

// WorkerhubContractAppendToBatchIterator is returned from FilterAppendToBatch and is used to iterate over the raw logs and unpacked data for AppendToBatch events raised by the WorkerhubContract contract.
type WorkerhubContractAppendToBatchIterator struct {
	Event *WorkerhubContractAppendToBatch // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractAppendToBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractAppendToBatch)
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
		it.Event = new(WorkerhubContractAppendToBatch)
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
func (it *WorkerhubContractAppendToBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractAppendToBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractAppendToBatch represents a AppendToBatch event raised by the WorkerhubContract contract.
type WorkerhubContractAppendToBatch struct {
	BatchId uint64
	ModelId uint32
	InferId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAppendToBatch is a free log retrieval operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterAppendToBatch(opts *bind.FilterOpts, batchId []uint64, modelId []uint32, inferId []uint64) (*WorkerhubContractAppendToBatchIterator, error) {

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

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractAppendToBatchIterator{contract: _WorkerhubContract.contract, event: "AppendToBatch", logs: logs, sub: sub}, nil
}

// WatchAppendToBatch is a free log subscription operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchAppendToBatch(opts *bind.WatchOpts, sink chan<- *WorkerhubContractAppendToBatch, batchId []uint64, modelId []uint32, inferId []uint64) (event.Subscription, error) {

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

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractAppendToBatch)
				if err := _WorkerhubContract.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseAppendToBatch(log types.Log) (*WorkerhubContractAppendToBatch, error) {
	event := new(WorkerhubContractAppendToBatch)
	if err := _WorkerhubContract.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the WorkerhubContract contract.
type WorkerhubContractInferenceStatusUpdateIterator struct {
	Event *WorkerhubContractInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractInferenceStatusUpdate)
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
		it.Event = new(WorkerhubContractInferenceStatusUpdate)
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
func (it *WorkerhubContractInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the WorkerhubContract contract.
type WorkerhubContractInferenceStatusUpdate struct {
	InferenceId uint64
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []uint64) (*WorkerhubContractInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractInferenceStatusUpdateIterator{contract: _WorkerhubContract.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *WorkerhubContractInferenceStatusUpdate, inferenceId []uint64) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractInferenceStatusUpdate)
				if err := _WorkerhubContract.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseInferenceStatusUpdate(log types.Log) (*WorkerhubContractInferenceStatusUpdate, error) {
	event := new(WorkerhubContractInferenceStatusUpdate)
	if err := _WorkerhubContract.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerhubContract contract.
type WorkerhubContractInitializedIterator struct {
	Event *WorkerhubContractInitialized // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractInitialized)
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
		it.Event = new(WorkerhubContractInitialized)
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
func (it *WorkerhubContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractInitialized represents a Initialized event raised by the WorkerhubContract contract.
type WorkerhubContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerhubContractInitializedIterator, error) {

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractInitializedIterator{contract: _WorkerhubContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerhubContractInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractInitialized)
				if err := _WorkerhubContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseInitialized(log types.Log) (*WorkerhubContractInitialized, error) {
	event := new(WorkerhubContractInitialized)
	if err := _WorkerhubContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the WorkerhubContract contract.
type WorkerhubContractNewAssignmentIterator struct {
	Event *WorkerhubContractNewAssignment // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractNewAssignment)
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
		it.Event = new(WorkerhubContractNewAssignment)
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
func (it *WorkerhubContractNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractNewAssignment represents a NewAssignment event raised by the WorkerhubContract contract.
type WorkerhubContractNewAssignment struct {
	InferenceId uint64
	Miner       common.Address
	ExpiredAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterNewAssignment(opts *bind.FilterOpts, inferenceId []uint64, miner []common.Address) (*WorkerhubContractNewAssignmentIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractNewAssignmentIterator{contract: _WorkerhubContract.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *WorkerhubContractNewAssignment, inferenceId []uint64, miner []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractNewAssignment)
				if err := _WorkerhubContract.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseNewAssignment(log types.Log) (*WorkerhubContractNewAssignment, error) {
	event := new(WorkerhubContractNewAssignment)
	if err := _WorkerhubContract.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the WorkerhubContract contract.
type WorkerhubContractNewInferenceIterator struct {
	Event *WorkerhubContractNewInference // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractNewInference)
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
		it.Event = new(WorkerhubContractNewInference)
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
func (it *WorkerhubContractNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractNewInference represents a NewInference event raised by the WorkerhubContract contract.
type WorkerhubContractNewInference struct {
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
func (_WorkerhubContract *WorkerhubContractFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []uint64, creator []common.Address, modelId []uint32) (*WorkerhubContractNewInferenceIterator, error) {

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

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractNewInferenceIterator{contract: _WorkerhubContract.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerhubContractNewInference, inferenceId []uint64, creator []common.Address, modelId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractNewInference)
				if err := _WorkerhubContract.contract.UnpackLog(event, "NewInference", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseNewInference(log types.Log) (*WorkerhubContractNewInference, error) {
	event := new(WorkerhubContractNewInference)
	if err := _WorkerhubContract.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerhubContract contract.
type WorkerhubContractOwnershipTransferredIterator struct {
	Event *WorkerhubContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractOwnershipTransferred)
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
		it.Event = new(WorkerhubContractOwnershipTransferred)
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
func (it *WorkerhubContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerhubContract contract.
type WorkerhubContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerhubContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractOwnershipTransferredIterator{contract: _WorkerhubContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerhubContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractOwnershipTransferred)
				if err := _WorkerhubContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerhubContractOwnershipTransferred, error) {
	event := new(WorkerhubContractOwnershipTransferred)
	if err := _WorkerhubContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerhubContract contract.
type WorkerhubContractPausedIterator struct {
	Event *WorkerhubContractPaused // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractPaused)
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
		it.Event = new(WorkerhubContractPaused)
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
func (it *WorkerhubContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractPaused represents a Paused event raised by the WorkerhubContract contract.
type WorkerhubContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerhubContractPausedIterator, error) {

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractPausedIterator{contract: _WorkerhubContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerhubContractPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractPaused)
				if err := _WorkerhubContract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParsePaused(log types.Log) (*WorkerhubContractPaused, error) {
	event := new(WorkerhubContractPaused)
	if err := _WorkerhubContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the WorkerhubContract contract.
type WorkerhubContractSolutionSubmissionIterator struct {
	Event *WorkerhubContractSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractSolutionSubmission)
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
		it.Event = new(WorkerhubContractSolutionSubmission)
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
func (it *WorkerhubContractSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractSolutionSubmission represents a SolutionSubmission event raised by the WorkerhubContract contract.
type WorkerhubContractSolutionSubmission struct {
	Miner   common.Address
	InferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, inferId []*big.Int) (*WorkerhubContractSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractSolutionSubmissionIterator{contract: _WorkerhubContract.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *WorkerhubContractSolutionSubmission, miner []common.Address, inferId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractSolutionSubmission)
				if err := _WorkerhubContract.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseSolutionSubmission(log types.Log) (*WorkerhubContractSolutionSubmission, error) {
	event := new(WorkerhubContractSolutionSubmission)
	if err := _WorkerhubContract.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the WorkerhubContract contract.
type WorkerhubContractStreamedDataIterator struct {
	Event *WorkerhubContractStreamedData // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractStreamedData)
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
		it.Event = new(WorkerhubContractStreamedData)
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
func (it *WorkerhubContractStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractStreamedData represents a StreamedData event raised by the WorkerhubContract contract.
type WorkerhubContractStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*WorkerhubContractStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractStreamedDataIterator{contract: _WorkerhubContract.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *WorkerhubContractStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractStreamedData)
				if err := _WorkerhubContract.contract.UnpackLog(event, "StreamedData", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseStreamedData(log types.Log) (*WorkerhubContractStreamedData, error) {
	event := new(WorkerhubContractStreamedData)
	if err := _WorkerhubContract.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerhubContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerhubContract contract.
type WorkerhubContractUnpausedIterator struct {
	Event *WorkerhubContractUnpaused // Event containing the contract specifics and raw log

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
func (it *WorkerhubContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerhubContractUnpaused)
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
		it.Event = new(WorkerhubContractUnpaused)
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
func (it *WorkerhubContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerhubContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerhubContractUnpaused represents a Unpaused event raised by the WorkerhubContract contract.
type WorkerhubContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerhubContract *WorkerhubContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerhubContractUnpausedIterator, error) {

	logs, sub, err := _WorkerhubContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerhubContractUnpausedIterator{contract: _WorkerhubContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerhubContract *WorkerhubContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerhubContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerhubContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerhubContractUnpaused)
				if err := _WorkerhubContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WorkerhubContract *WorkerhubContractFilterer) ParseUnpaused(log types.Log) (*WorkerhubContractUnpaused, error) {
	event := new(WorkerhubContractUnpaused)
	if err := _WorkerhubContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
