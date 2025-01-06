// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iv3swaprouter

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

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterMetaData contains all meta data concerning the IV3SwapRouter contract.
var IV3SwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IV3SwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IV3SwapRouterMetaData.ABI instead.
var IV3SwapRouterABI = IV3SwapRouterMetaData.ABI

// IV3SwapRouter is an auto generated Go binding around an Ethereum contract.
type IV3SwapRouter struct {
	IV3SwapRouterCaller     // Read-only binding to the contract
	IV3SwapRouterTransactor // Write-only binding to the contract
	IV3SwapRouterFilterer   // Log filterer for contract events
}

// IV3SwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IV3SwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3SwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IV3SwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3SwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IV3SwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IV3SwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IV3SwapRouterSession struct {
	Contract     *IV3SwapRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IV3SwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IV3SwapRouterCallerSession struct {
	Contract *IV3SwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IV3SwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IV3SwapRouterTransactorSession struct {
	Contract     *IV3SwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IV3SwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IV3SwapRouterRaw struct {
	Contract *IV3SwapRouter // Generic contract binding to access the raw methods on
}

// IV3SwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IV3SwapRouterCallerRaw struct {
	Contract *IV3SwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IV3SwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IV3SwapRouterTransactorRaw struct {
	Contract *IV3SwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIV3SwapRouter creates a new instance of IV3SwapRouter, bound to a specific deployed contract.
func NewIV3SwapRouter(address common.Address, backend bind.ContractBackend) (*IV3SwapRouter, error) {
	contract, err := bindIV3SwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IV3SwapRouter{IV3SwapRouterCaller: IV3SwapRouterCaller{contract: contract}, IV3SwapRouterTransactor: IV3SwapRouterTransactor{contract: contract}, IV3SwapRouterFilterer: IV3SwapRouterFilterer{contract: contract}}, nil
}

// NewIV3SwapRouterCaller creates a new read-only instance of IV3SwapRouter, bound to a specific deployed contract.
func NewIV3SwapRouterCaller(address common.Address, caller bind.ContractCaller) (*IV3SwapRouterCaller, error) {
	contract, err := bindIV3SwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IV3SwapRouterCaller{contract: contract}, nil
}

// NewIV3SwapRouterTransactor creates a new write-only instance of IV3SwapRouter, bound to a specific deployed contract.
func NewIV3SwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IV3SwapRouterTransactor, error) {
	contract, err := bindIV3SwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IV3SwapRouterTransactor{contract: contract}, nil
}

// NewIV3SwapRouterFilterer creates a new log filterer instance of IV3SwapRouter, bound to a specific deployed contract.
func NewIV3SwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IV3SwapRouterFilterer, error) {
	contract, err := bindIV3SwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IV3SwapRouterFilterer{contract: contract}, nil
}

// bindIV3SwapRouter binds a generic wrapper to an already deployed contract.
func bindIV3SwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IV3SwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IV3SwapRouter *IV3SwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IV3SwapRouter.Contract.IV3SwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IV3SwapRouter *IV3SwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.IV3SwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IV3SwapRouter *IV3SwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.IV3SwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IV3SwapRouter *IV3SwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IV3SwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IV3SwapRouter *IV3SwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IV3SwapRouter *IV3SwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IV3SwapRouter *IV3SwapRouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IV3SwapRouter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IV3SwapRouter *IV3SwapRouterSession) WETH9() (common.Address, error) {
	return _IV3SwapRouter.Contract.WETH9(&_IV3SwapRouter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_IV3SwapRouter *IV3SwapRouterCallerSession) WETH9() (common.Address, error) {
	return _IV3SwapRouter.Contract.WETH9(&_IV3SwapRouter.CallOpts)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IV3SwapRouter *IV3SwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IV3SwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IV3SwapRouter *IV3SwapRouterSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.ExactInputSingle(&_IV3SwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_IV3SwapRouter *IV3SwapRouterTransactorSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.ExactInputSingle(&_IV3SwapRouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3SwapRouter *IV3SwapRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _IV3SwapRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3SwapRouter *IV3SwapRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.Multicall(&_IV3SwapRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IV3SwapRouter *IV3SwapRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.Multicall(&_IV3SwapRouter.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IV3SwapRouter.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IV3SwapRouter *IV3SwapRouterSession) RefundETH() (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.RefundETH(&_IV3SwapRouter.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.RefundETH(&_IV3SwapRouter.TransactOpts)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.SweepToken(&_IV3SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.SweepToken(&_IV3SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.UnwrapWETH9(&_IV3SwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_IV3SwapRouter *IV3SwapRouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IV3SwapRouter.Contract.UnwrapWETH9(&_IV3SwapRouter.TransactOpts, amountMinimum, recipient)
}
