// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ihybridmodel

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

// IHybridModelMetaData contains all meta data concerning the IHybridModel contract.
var IHybridModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"referenceId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IHybridModelABI is the input ABI used to generate the binding from.
// Deprecated: Use IHybridModelMetaData.ABI instead.
var IHybridModelABI = IHybridModelMetaData.ABI

// IHybridModel is an auto generated Go binding around an Ethereum contract.
type IHybridModel struct {
	IHybridModelCaller     // Read-only binding to the contract
	IHybridModelTransactor // Write-only binding to the contract
	IHybridModelFilterer   // Log filterer for contract events
}

// IHybridModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHybridModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHybridModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHybridModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHybridModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHybridModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHybridModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHybridModelSession struct {
	Contract     *IHybridModel     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHybridModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHybridModelCallerSession struct {
	Contract *IHybridModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IHybridModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHybridModelTransactorSession struct {
	Contract     *IHybridModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IHybridModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHybridModelRaw struct {
	Contract *IHybridModel // Generic contract binding to access the raw methods on
}

// IHybridModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHybridModelCallerRaw struct {
	Contract *IHybridModelCaller // Generic read-only contract binding to access the raw methods on
}

// IHybridModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHybridModelTransactorRaw struct {
	Contract *IHybridModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHybridModel creates a new instance of IHybridModel, bound to a specific deployed contract.
func NewIHybridModel(address common.Address, backend bind.ContractBackend) (*IHybridModel, error) {
	contract, err := bindIHybridModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHybridModel{IHybridModelCaller: IHybridModelCaller{contract: contract}, IHybridModelTransactor: IHybridModelTransactor{contract: contract}, IHybridModelFilterer: IHybridModelFilterer{contract: contract}}, nil
}

// NewIHybridModelCaller creates a new read-only instance of IHybridModel, bound to a specific deployed contract.
func NewIHybridModelCaller(address common.Address, caller bind.ContractCaller) (*IHybridModelCaller, error) {
	contract, err := bindIHybridModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHybridModelCaller{contract: contract}, nil
}

// NewIHybridModelTransactor creates a new write-only instance of IHybridModel, bound to a specific deployed contract.
func NewIHybridModelTransactor(address common.Address, transactor bind.ContractTransactor) (*IHybridModelTransactor, error) {
	contract, err := bindIHybridModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHybridModelTransactor{contract: contract}, nil
}

// NewIHybridModelFilterer creates a new log filterer instance of IHybridModel, bound to a specific deployed contract.
func NewIHybridModelFilterer(address common.Address, filterer bind.ContractFilterer) (*IHybridModelFilterer, error) {
	contract, err := bindIHybridModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHybridModelFilterer{contract: contract}, nil
}

// bindIHybridModel binds a generic wrapper to an already deployed contract.
func bindIHybridModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHybridModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHybridModel *IHybridModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHybridModel.Contract.IHybridModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHybridModel *IHybridModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHybridModel.Contract.IHybridModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHybridModel *IHybridModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHybridModel.Contract.IHybridModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHybridModel *IHybridModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHybridModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHybridModel *IHybridModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHybridModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHybridModel *IHybridModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHybridModel.Contract.contract.Transact(opts, method, params...)
}

// Infer is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _data) payable returns(uint256 referenceId)
func (_IHybridModel *IHybridModelTransactor) Infer(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _IHybridModel.contract.Transact(opts, "infer", _data)
}

// Infer is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _data) payable returns(uint256 referenceId)
func (_IHybridModel *IHybridModelSession) Infer(_data []byte) (*types.Transaction, error) {
	return _IHybridModel.Contract.Infer(&_IHybridModel.TransactOpts, _data)
}

// Infer is a paid mutator transaction binding the contract method 0x67e950a8.
//
// Solidity: function infer(bytes _data) payable returns(uint256 referenceId)
func (_IHybridModel *IHybridModelTransactorSession) Infer(_data []byte) (*types.Transaction, error) {
	return _IHybridModel.Contract.Infer(&_IHybridModel.TransactOpts, _data)
}
