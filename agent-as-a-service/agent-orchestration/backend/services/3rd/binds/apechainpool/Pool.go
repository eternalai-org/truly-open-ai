// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package apechainpool

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

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"communityFee1New\",\"type\":\"uint8\"}],\"name\":\"CommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeZto\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeOtz\",\"type\":\"uint16\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"Incentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"liquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"LiquidityCooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"newTickSpacing\",\"type\":\"int24\"}],\"name\":\"TickSpacing\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeIncentive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorageOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"}],\"name\":\"getInnerCumulatives\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"innerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"innerSecondsSpentPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"innerSecondsSpent\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"getTimepoints\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulatives\",\"type\":\"uint160[]\"},{\"internalType\":\"uint112[]\",\"name\":\"volatilityCumulatives\",\"type\":\"uint112[]\"},{\"internalType\":\"uint256[]\",\"name\":\"volumePerAvgLiquiditys\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalState\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"feeZto\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeOtz\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timepointIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"communityFeeToken1\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"initialPrice\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidityDesired\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityActual\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastLiquidityAddTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"fees0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fees1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"communityFee0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"communityFee1\",\"type\":\"uint8\"}],\"name\":\"setCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"setIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newLiquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"setLiquidityCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newTickSpacing\",\"type\":\"int24\"}],\"name\":\"setTickSpacing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapSupportingFeeOnInputTokens\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"tickTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityTotal\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"outerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"outerSecondsPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"outerSecondsSpent\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"timepoints\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulative\",\"type\":\"uint160\"},{\"internalType\":\"uint88\",\"name\":\"volatilityCumulative\",\"type\":\"uint88\"},{\"internalType\":\"int24\",\"name\":\"averageTick\",\"type\":\"int24\"},{\"internalType\":\"uint144\",\"name\":\"volumePerLiquidityCumulative\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth0Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth1Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Pool *PoolCaller) ActiveIncentive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "activeIncentive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Pool *PoolSession) ActiveIncentive() (common.Address, error) {
	return _Pool.Contract.ActiveIncentive(&_Pool.CallOpts)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Pool *PoolCallerSession) ActiveIncentive() (common.Address, error) {
	return _Pool.Contract.ActiveIncentive(&_Pool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Pool *PoolCaller) DataStorageOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "dataStorageOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Pool *PoolSession) DataStorageOperator() (common.Address, error) {
	return _Pool.Contract.DataStorageOperator(&_Pool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Pool *PoolCallerSession) DataStorageOperator() (common.Address, error) {
	return _Pool.Contract.DataStorageOperator(&_Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCallerSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Pool *PoolCaller) GetInnerCumulatives(opts *bind.CallOpts, bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getInnerCumulatives", bottomTick, topTick)

	outstruct := new(struct {
		InnerTickCumulative           *big.Int
		InnerSecondsSpentPerLiquidity *big.Int
		InnerSecondsSpent             uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnerTickCumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpentPerLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpent = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Pool *PoolSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _Pool.Contract.GetInnerCumulatives(&_Pool.CallOpts, bottomTick, topTick)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Pool *PoolCallerSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _Pool.Contract.GetInnerCumulatives(&_Pool.CallOpts, bottomTick, topTick)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Pool *PoolCaller) GetTimepoints(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTimepoints", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                []*big.Int
		SecondsPerLiquidityCumulatives []*big.Int
		VolatilityCumulatives          []*big.Int
		VolumePerAvgLiquiditys         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulatives = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolatilityCumulatives = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolumePerAvgLiquiditys = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Pool *PoolSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _Pool.Contract.GetTimepoints(&_Pool.CallOpts, secondsAgos)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Pool *PoolCallerSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _Pool.Contract.GetTimepoints(&_Pool.CallOpts, secondsAgos)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_Pool *PoolCaller) GlobalState(opts *bind.CallOpts) (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "globalState")

	outstruct := new(struct {
		Price              *big.Int
		Tick               *big.Int
		FeeZto             uint16
		FeeOtz             uint16
		TimepointIndex     uint16
		CommunityFeeToken0 uint8
		CommunityFeeToken1 uint8
		Unlocked           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeZto = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.FeeOtz = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.TimepointIndex = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken0 = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.CommunityFeeToken1 = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_Pool *PoolSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _Pool.Contract.GlobalState(&_Pool.CallOpts)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 feeZto, uint16 feeOtz, uint16 timepointIndex, uint8 communityFeeToken0, uint8 communityFeeToken1, bool unlocked)
func (_Pool *PoolCallerSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	return _Pool.Contract.GlobalState(&_Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Pool *PoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Pool *PoolSession) Liquidity() (*big.Int, error) {
	return _Pool.Contract.Liquidity(&_Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Pool *PoolCallerSession) Liquidity() (*big.Int, error) {
	return _Pool.Contract.Liquidity(&_Pool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Pool *PoolCaller) LiquidityCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "liquidityCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Pool *PoolSession) LiquidityCooldown() (uint32, error) {
	return _Pool.Contract.LiquidityCooldown(&_Pool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Pool *PoolCallerSession) LiquidityCooldown() (uint32, error) {
	return _Pool.Contract.LiquidityCooldown(&_Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Pool *PoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Pool *PoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _Pool.Contract.MaxLiquidityPerTick(&_Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Pool *PoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _Pool.Contract.MaxLiquidityPerTick(&_Pool.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Pool *PoolCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Liquidity                 *big.Int
		LastLiquidityAddTimestamp uint32
		InnerFeeGrowth0Token      *big.Int
		InnerFeeGrowth1Token      *big.Int
		Fees0                     *big.Int
		Fees1                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLiquidityAddTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.InnerFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InnerFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Pool *PoolSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _Pool.Contract.Positions(&_Pool.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Pool *PoolCallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _Pool.Contract.Positions(&_Pool.CallOpts, arg0)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Pool *PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Pool *PoolSession) TickSpacing() (*big.Int, error) {
	return _Pool.Contract.TickSpacing(&_Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Pool *PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _Pool.Contract.TickSpacing(&_Pool.CallOpts)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Pool *PoolCaller) TickTable(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tickTable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Pool *PoolSession) TickTable(arg0 int16) (*big.Int, error) {
	return _Pool.Contract.TickTable(&_Pool.CallOpts, arg0)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Pool *PoolCallerSession) TickTable(arg0 int16) (*big.Int, error) {
	return _Pool.Contract.TickTable(&_Pool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Pool *PoolCaller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "ticks", arg0)

	outstruct := new(struct {
		LiquidityTotal           *big.Int
		LiquidityDelta           *big.Int
		OuterFeeGrowth0Token     *big.Int
		OuterFeeGrowth1Token     *big.Int
		OuterTickCumulative      *big.Int
		OuterSecondsPerLiquidity *big.Int
		OuterSecondsSpent        uint32
		Initialized              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OuterTickCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsPerLiquidity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsSpent = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Pool *PoolSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _Pool.Contract.Ticks(&_Pool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Pool *PoolCallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _Pool.Contract.Ticks(&_Pool.CallOpts, arg0)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Pool *PoolCaller) Timepoints(opts *bind.CallOpts, index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "timepoints", index)

	outstruct := new(struct {
		Initialized                   bool
		BlockTimestamp                uint32
		TickCumulative                *big.Int
		SecondsPerLiquidityCumulative *big.Int
		VolatilityCumulative          *big.Int
		AverageTick                   *big.Int
		VolumePerLiquidityCumulative  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulative = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VolatilityCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AverageTick = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VolumePerLiquidityCumulative = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Pool *PoolSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _Pool.Contract.Timepoints(&_Pool.CallOpts, index)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Pool *PoolCallerSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _Pool.Contract.Timepoints(&_Pool.CallOpts, index)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolSession) Token0() (common.Address, error) {
	return _Pool.Contract.Token0(&_Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolCallerSession) Token0() (common.Address, error) {
	return _Pool.Contract.Token0(&_Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolSession) Token1() (common.Address, error) {
	return _Pool.Contract.Token1(&_Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolCallerSession) Token1() (common.Address, error) {
	return _Pool.Contract.Token1(&_Pool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Pool *PoolCaller) TotalFeeGrowth0Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalFeeGrowth0Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Pool *PoolSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _Pool.Contract.TotalFeeGrowth0Token(&_Pool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Pool *PoolCallerSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _Pool.Contract.TotalFeeGrowth0Token(&_Pool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Pool *PoolCaller) TotalFeeGrowth1Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalFeeGrowth1Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Pool *PoolSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _Pool.Contract.TotalFeeGrowth1Token(&_Pool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Pool *PoolCallerSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _Pool.Contract.TotalFeeGrowth1Token(&_Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Pool *PoolTransactor) Burn(opts *bind.TransactOpts, bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burn", bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Pool *PoolSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Pool *PoolTransactorSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, bottomTick, topTick, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Pool *PoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "collect", recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Pool *PoolSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Collect(&_Pool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Pool *PoolTransactorSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Collect(&_Pool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Pool *PoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Pool *PoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Flash(&_Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Pool *PoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Flash(&_Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, initialPrice *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Pool *PoolSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Pool *PoolTransactorSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, initialPrice)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Pool *PoolSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Pool *PoolTransactorSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_Pool *PoolTransactor) SetCommunityFee(opts *bind.TransactOpts, communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setCommunityFee", communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_Pool *PoolSession) SetCommunityFee(communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _Pool.Contract.SetCommunityFee(&_Pool.TransactOpts, communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x7c0112b7.
//
// Solidity: function setCommunityFee(uint8 communityFee0, uint8 communityFee1) returns()
func (_Pool *PoolTransactorSession) SetCommunityFee(communityFee0 uint8, communityFee1 uint8) (*types.Transaction, error) {
	return _Pool.Contract.SetCommunityFee(&_Pool.TransactOpts, communityFee0, communityFee1)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Pool *PoolTransactor) SetIncentive(opts *bind.TransactOpts, virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setIncentive", virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Pool *PoolSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetIncentive(&_Pool.TransactOpts, virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Pool *PoolTransactorSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetIncentive(&_Pool.TransactOpts, virtualPoolAddress)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Pool *PoolTransactor) SetLiquidityCooldown(opts *bind.TransactOpts, newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setLiquidityCooldown", newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Pool *PoolSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Pool.Contract.SetLiquidityCooldown(&_Pool.TransactOpts, newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Pool *PoolTransactorSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Pool.Contract.SetLiquidityCooldown(&_Pool.TransactOpts, newLiquidityCooldown)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Pool *PoolTransactor) SetTickSpacing(opts *bind.TransactOpts, newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setTickSpacing", newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Pool *PoolSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetTickSpacing(&_Pool.TransactOpts, newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Pool *PoolTransactorSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetTickSpacing(&_Pool.TransactOpts, newTickSpacing)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swap", recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolTransactorSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolTransactor) SwapSupportingFeeOnInputTokens(opts *bind.TransactOpts, sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swapSupportingFeeOnInputTokens", sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.SwapSupportingFeeOnInputTokens(&_Pool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Pool *PoolTransactorSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.SwapSupportingFeeOnInputTokens(&_Pool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pool contract.
type PoolBurnIterator struct {
	Event *PoolBurn // Event containing the contract specifics and raw log

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
func (it *PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBurn)
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
		it.Event = new(PoolBurn)
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
func (it *PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBurn represents a Burn event raised by the Pool contract.
type PoolBurn struct {
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*PoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &PoolBurnIterator{contract: _Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PoolBurn, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBurn)
				if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) ParseBurn(log types.Log) (*PoolBurn, error) {
	event := new(PoolBurn)
	if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the Pool contract.
type PoolCollectIterator struct {
	Event *PoolCollect // Event containing the contract specifics and raw log

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
func (it *PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCollect)
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
		it.Event = new(PoolCollect)
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
func (it *PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCollect represents a Collect event raised by the Pool contract.
type PoolCollect struct {
	Owner      common.Address
	Recipient  common.Address
	BottomTick *big.Int
	TopTick    *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Pool *PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*PoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &PoolCollectIterator{contract: _Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Pool *PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *PoolCollect, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCollect)
				if err := _Pool.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Pool *PoolFilterer) ParseCollect(log types.Log) (*PoolCollect, error) {
	event := new(PoolCollect)
	if err := _Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolCommunityFeeIterator is returned from FilterCommunityFee and is used to iterate over the raw logs and unpacked data for CommunityFee events raised by the Pool contract.
type PoolCommunityFeeIterator struct {
	Event *PoolCommunityFee // Event containing the contract specifics and raw log

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
func (it *PoolCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCommunityFee)
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
		it.Event = new(PoolCommunityFee)
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
func (it *PoolCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCommunityFee represents a CommunityFee event raised by the Pool contract.
type PoolCommunityFee struct {
	CommunityFee0New uint8
	CommunityFee1New uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommunityFee is a free log retrieval operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_Pool *PoolFilterer) FilterCommunityFee(opts *bind.FilterOpts) (*PoolCommunityFeeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return &PoolCommunityFeeIterator{contract: _Pool.contract, event: "CommunityFee", logs: logs, sub: sub}, nil
}

// WatchCommunityFee is a free log subscription operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_Pool *PoolFilterer) WatchCommunityFee(opts *bind.WatchOpts, sink chan<- *PoolCommunityFee) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCommunityFee)
				if err := _Pool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
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

// ParseCommunityFee is a log parse operation binding the contract event 0x9e22b964b08e25c3aaa72102bb0071c089258fb82d51271a8ddf5c24921356ee.
//
// Solidity: event CommunityFee(uint8 communityFee0New, uint8 communityFee1New)
func (_Pool *PoolFilterer) ParseCommunityFee(log types.Log) (*PoolCommunityFee, error) {
	event := new(PoolCommunityFee)
	if err := _Pool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Pool contract.
type PoolFeeIterator struct {
	Event *PoolFee // Event containing the contract specifics and raw log

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
func (it *PoolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFee)
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
		it.Event = new(PoolFee)
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
func (it *PoolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFee represents a Fee event raised by the Pool contract.
type PoolFee struct {
	FeeZto uint16
	FeeOtz uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_Pool *PoolFilterer) FilterFee(opts *bind.FilterOpts) (*PoolFeeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &PoolFeeIterator{contract: _Pool.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_Pool *PoolFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *PoolFee) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFee)
				if err := _Pool.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x8a89de70856bccec096661388f305b9a75f5f65cb0d8a0e1e803c39dabedb57f.
//
// Solidity: event Fee(uint16 feeZto, uint16 feeOtz)
func (_Pool *PoolFilterer) ParseFee(log types.Log) (*PoolFee, error) {
	event := new(PoolFee)
	if err := _Pool.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the Pool contract.
type PoolFlashIterator struct {
	Event *PoolFlash // Event containing the contract specifics and raw log

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
func (it *PoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFlash)
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
		it.Event = new(PoolFlash)
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
func (it *PoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFlash represents a Flash event raised by the Pool contract.
type PoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Pool *PoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*PoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PoolFlashIterator{contract: _Pool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Pool *PoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *PoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFlash)
				if err := _Pool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Pool *PoolFilterer) ParseFlash(log types.Log) (*PoolFlash, error) {
	event := new(PoolFlash)
	if err := _Pool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolIncentiveIterator is returned from FilterIncentive and is used to iterate over the raw logs and unpacked data for Incentive events raised by the Pool contract.
type PoolIncentiveIterator struct {
	Event *PoolIncentive // Event containing the contract specifics and raw log

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
func (it *PoolIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolIncentive)
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
		it.Event = new(PoolIncentive)
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
func (it *PoolIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolIncentive represents a Incentive event raised by the Pool contract.
type PoolIncentive struct {
	VirtualPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentive is a free log retrieval operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Pool *PoolFilterer) FilterIncentive(opts *bind.FilterOpts, virtualPoolAddress []common.Address) (*PoolIncentiveIterator, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolIncentiveIterator{contract: _Pool.contract, event: "Incentive", logs: logs, sub: sub}, nil
}

// WatchIncentive is a free log subscription operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Pool *PoolFilterer) WatchIncentive(opts *bind.WatchOpts, sink chan<- *PoolIncentive, virtualPoolAddress []common.Address) (event.Subscription, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolIncentive)
				if err := _Pool.contract.UnpackLog(event, "Incentive", log); err != nil {
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

// ParseIncentive is a log parse operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Pool *PoolFilterer) ParseIncentive(log types.Log) (*PoolIncentive, error) {
	event := new(PoolIncentive)
	if err := _Pool.contract.UnpackLog(event, "Incentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Pool contract.
type PoolInitializeIterator struct {
	Event *PoolInitialize // Event containing the contract specifics and raw log

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
func (it *PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInitialize)
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
		it.Event = new(PoolInitialize)
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
func (it *PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInitialize represents a Initialize event raised by the Pool contract.
type PoolInitialize struct {
	Price *big.Int
	Tick  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Pool *PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*PoolInitializeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &PoolInitializeIterator{contract: _Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Pool *PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInitialize)
				if err := _Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Pool *PoolFilterer) ParseInitialize(log types.Log) (*PoolInitialize, error) {
	event := new(PoolInitialize)
	if err := _Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolLiquidityCooldownIterator is returned from FilterLiquidityCooldown and is used to iterate over the raw logs and unpacked data for LiquidityCooldown events raised by the Pool contract.
type PoolLiquidityCooldownIterator struct {
	Event *PoolLiquidityCooldown // Event containing the contract specifics and raw log

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
func (it *PoolLiquidityCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolLiquidityCooldown)
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
		it.Event = new(PoolLiquidityCooldown)
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
func (it *PoolLiquidityCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolLiquidityCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolLiquidityCooldown represents a LiquidityCooldown event raised by the Pool contract.
type PoolLiquidityCooldown struct {
	LiquidityCooldown uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCooldown is a free log retrieval operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Pool *PoolFilterer) FilterLiquidityCooldown(opts *bind.FilterOpts) (*PoolLiquidityCooldownIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return &PoolLiquidityCooldownIterator{contract: _Pool.contract, event: "LiquidityCooldown", logs: logs, sub: sub}, nil
}

// WatchLiquidityCooldown is a free log subscription operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Pool *PoolFilterer) WatchLiquidityCooldown(opts *bind.WatchOpts, sink chan<- *PoolLiquidityCooldown) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolLiquidityCooldown)
				if err := _Pool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
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

// ParseLiquidityCooldown is a log parse operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Pool *PoolFilterer) ParseLiquidityCooldown(log types.Log) (*PoolLiquidityCooldown, error) {
	event := new(PoolLiquidityCooldown)
	if err := _Pool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pool contract.
type PoolMintIterator struct {
	Event *PoolMint // Event containing the contract specifics and raw log

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
func (it *PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMint)
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
		it.Event = new(PoolMint)
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
func (it *PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMint represents a Mint event raised by the Pool contract.
type PoolMint struct {
	Sender          common.Address
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*PoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &PoolMintIterator{contract: _Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PoolMint, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMint)
				if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Pool *PoolFilterer) ParseMint(log types.Log) (*PoolMint, error) {
	event := new(PoolMint)
	if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pool contract.
type PoolSwapIterator struct {
	Event *PoolSwap // Event containing the contract specifics and raw log

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
func (it *PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwap)
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
		it.Event = new(PoolSwap)
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
func (it *PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwap represents a Swap event raised by the Pool contract.
type PoolSwap struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Price     *big.Int
	Liquidity *big.Int
	Tick      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Pool *PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PoolSwapIterator{contract: _Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Pool *PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwap)
				if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Pool *PoolFilterer) ParseSwap(log types.Log) (*PoolSwap, error) {
	event := new(PoolSwap)
	if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTickSpacingIterator is returned from FilterTickSpacing and is used to iterate over the raw logs and unpacked data for TickSpacing events raised by the Pool contract.
type PoolTickSpacingIterator struct {
	Event *PoolTickSpacing // Event containing the contract specifics and raw log

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
func (it *PoolTickSpacingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTickSpacing)
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
		it.Event = new(PoolTickSpacing)
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
func (it *PoolTickSpacingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTickSpacingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTickSpacing represents a TickSpacing event raised by the Pool contract.
type PoolTickSpacing struct {
	NewTickSpacing *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTickSpacing is a free log retrieval operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_Pool *PoolFilterer) FilterTickSpacing(opts *bind.FilterOpts) (*PoolTickSpacingIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return &PoolTickSpacingIterator{contract: _Pool.contract, event: "TickSpacing", logs: logs, sub: sub}, nil
}

// WatchTickSpacing is a free log subscription operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_Pool *PoolFilterer) WatchTickSpacing(opts *bind.WatchOpts, sink chan<- *PoolTickSpacing) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TickSpacing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTickSpacing)
				if err := _Pool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
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

// ParseTickSpacing is a log parse operation binding the contract event 0x01413b1d5d4c359e9a0daa7909ecda165f6e8c51fe2ff529d74b22a5a7c02645.
//
// Solidity: event TickSpacing(int24 newTickSpacing)
func (_Pool *PoolFilterer) ParseTickSpacing(log types.Log) (*PoolTickSpacing, error) {
	event := new(PoolTickSpacing)
	if err := _Pool.contract.UnpackLog(event, "TickSpacing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
