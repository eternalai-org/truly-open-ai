// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrumfactory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"DefaultCommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFarmingAddress\",\"type\":\"address\"}],\"name\":\"FarmingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"FeeConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Owner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseFeeConfiguration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCommunityFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"setBaseFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"setDefaultCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmingAddress\",\"type\":\"address\"}],\"name\":\"setFarmingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactoryCaller) BaseFeeConfiguration(opts *bind.CallOpts) (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "baseFeeConfiguration")

	outstruct := new(struct {
		Alpha1      uint16
		Alpha2      uint16
		Beta1       uint32
		Beta2       uint32
		Gamma1      uint16
		Gamma2      uint16
		VolumeBeta  uint32
		VolumeGamma uint16
		BaseFee     uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Alpha1 = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Alpha2 = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Beta1 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Beta2 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Gamma1 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Gamma2 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.VolumeBeta = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.VolumeGamma = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.BaseFee = *abi.ConvertType(out[8], new(uint16)).(*uint16)

	return *outstruct, err

}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactorySession) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _Factory.Contract.BaseFeeConfiguration(&_Factory.CallOpts)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactoryCallerSession) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _Factory.Contract.BaseFeeConfiguration(&_Factory.CallOpts)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Factory *FactoryCaller) DefaultCommunityFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "defaultCommunityFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Factory *FactorySession) DefaultCommunityFee() (uint8, error) {
	return _Factory.Contract.DefaultCommunityFee(&_Factory.CallOpts)
}

// DefaultCommunityFee is a free data retrieval call binding the contract method 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (_Factory *FactoryCallerSession) DefaultCommunityFee() (uint8, error) {
	return _Factory.Contract.DefaultCommunityFee(&_Factory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Factory *FactoryCaller) FarmingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "farmingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Factory *FactorySession) FarmingAddress() (common.Address, error) {
	return _Factory.Contract.FarmingAddress(&_Factory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_Factory *FactoryCallerSession) FarmingAddress() (common.Address, error) {
	return _Factory.Contract.FarmingAddress(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactorySession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCallerSession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Factory *FactoryCaller) PoolByPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "poolByPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Factory *FactorySession) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Factory.Contract.PoolByPair(&_Factory.CallOpts, arg0, arg1)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_Factory *FactoryCallerSession) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Factory.Contract.PoolByPair(&_Factory.CallOpts, arg0, arg1)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Factory *FactoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Factory *FactorySession) PoolDeployer() (common.Address, error) {
	return _Factory.Contract.PoolDeployer(&_Factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_Factory *FactoryCallerSession) PoolDeployer() (common.Address, error) {
	return _Factory.Contract.PoolDeployer(&_Factory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Factory *FactoryCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Factory *FactorySession) VaultAddress() (common.Address, error) {
	return _Factory.Contract.VaultAddress(&_Factory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Factory *FactoryCallerSession) VaultAddress() (common.Address, error) {
	return _Factory.Contract.VaultAddress(&_Factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Factory *FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPool", tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Factory *FactorySession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreatePool(&_Factory.TransactOpts, tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_Factory *FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreatePool(&_Factory.TransactOpts, tokenA, tokenB)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Factory *FactoryTransactor) SetBaseFeeConfiguration(opts *bind.TransactOpts, alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Factory *FactorySession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Factory.Contract.SetBaseFeeConfiguration(&_Factory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_Factory *FactoryTransactorSession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _Factory.Contract.SetBaseFeeConfiguration(&_Factory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Factory *FactoryTransactor) SetDefaultCommunityFee(opts *bind.TransactOpts, newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setDefaultCommunityFee", newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Factory *FactorySession) SetDefaultCommunityFee(newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Factory.Contract.SetDefaultCommunityFee(&_Factory.TransactOpts, newDefaultCommunityFee)
}

// SetDefaultCommunityFee is a paid mutator transaction binding the contract method 0x371e3521.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (_Factory *FactoryTransactorSession) SetDefaultCommunityFee(newDefaultCommunityFee uint8) (*types.Transaction, error) {
	return _Factory.Contract.SetDefaultCommunityFee(&_Factory.TransactOpts, newDefaultCommunityFee)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Factory *FactoryTransactor) SetFarmingAddress(opts *bind.TransactOpts, _farmingAddress common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFarmingAddress", _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Factory *FactorySession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFarmingAddress(&_Factory.TransactOpts, _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_Factory *FactoryTransactorSession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFarmingAddress(&_Factory.TransactOpts, _farmingAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Factory *FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Factory *FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetOwner(&_Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Factory *FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetOwner(&_Factory.TransactOpts, _owner)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Factory *FactoryTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Factory *FactorySession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetVaultAddress(&_Factory.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_Factory *FactoryTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetVaultAddress(&_Factory.TransactOpts, _vaultAddress)
}

// FactoryDefaultCommunityFeeIterator is returned from FilterDefaultCommunityFee and is used to iterate over the raw logs and unpacked data for DefaultCommunityFee events raised by the Factory contract.
type FactoryDefaultCommunityFeeIterator struct {
	Event *FactoryDefaultCommunityFee // Event containing the contract specifics and raw log

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
func (it *FactoryDefaultCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryDefaultCommunityFee)
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
		it.Event = new(FactoryDefaultCommunityFee)
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
func (it *FactoryDefaultCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryDefaultCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryDefaultCommunityFee represents a DefaultCommunityFee event raised by the Factory contract.
type FactoryDefaultCommunityFee struct {
	NewDefaultCommunityFee uint8
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDefaultCommunityFee is a free log retrieval operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Factory *FactoryFilterer) FilterDefaultCommunityFee(opts *bind.FilterOpts) (*FactoryDefaultCommunityFeeIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return &FactoryDefaultCommunityFeeIterator{contract: _Factory.contract, event: "DefaultCommunityFee", logs: logs, sub: sub}, nil
}

// WatchDefaultCommunityFee is a free log subscription operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Factory *FactoryFilterer) WatchDefaultCommunityFee(opts *bind.WatchOpts, sink chan<- *FactoryDefaultCommunityFee) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "DefaultCommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryDefaultCommunityFee)
				if err := _Factory.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
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

// ParseDefaultCommunityFee is a log parse operation binding the contract event 0x88cb5103fd9d88d417e72dc496030c71c65d1500548a9e9530e7d812b6a35558.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (_Factory *FactoryFilterer) ParseDefaultCommunityFee(log types.Log) (*FactoryDefaultCommunityFee, error) {
	event := new(FactoryDefaultCommunityFee)
	if err := _Factory.contract.UnpackLog(event, "DefaultCommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryFarmingAddressIterator is returned from FilterFarmingAddress and is used to iterate over the raw logs and unpacked data for FarmingAddress events raised by the Factory contract.
type FactoryFarmingAddressIterator struct {
	Event *FactoryFarmingAddress // Event containing the contract specifics and raw log

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
func (it *FactoryFarmingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryFarmingAddress)
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
		it.Event = new(FactoryFarmingAddress)
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
func (it *FactoryFarmingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryFarmingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryFarmingAddress represents a FarmingAddress event raised by the Factory contract.
type FactoryFarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFarmingAddress is a free log retrieval operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Factory *FactoryFilterer) FilterFarmingAddress(opts *bind.FilterOpts, newFarmingAddress []common.Address) (*FactoryFarmingAddressIterator, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return &FactoryFarmingAddressIterator{contract: _Factory.contract, event: "FarmingAddress", logs: logs, sub: sub}, nil
}

// WatchFarmingAddress is a free log subscription operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Factory *FactoryFilterer) WatchFarmingAddress(opts *bind.WatchOpts, sink chan<- *FactoryFarmingAddress, newFarmingAddress []common.Address) (event.Subscription, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryFarmingAddress)
				if err := _Factory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
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

// ParseFarmingAddress is a log parse operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_Factory *FactoryFilterer) ParseFarmingAddress(log types.Log) (*FactoryFarmingAddress, error) {
	event := new(FactoryFarmingAddress)
	if err := _Factory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryFeeConfigurationIterator is returned from FilterFeeConfiguration and is used to iterate over the raw logs and unpacked data for FeeConfiguration events raised by the Factory contract.
type FactoryFeeConfigurationIterator struct {
	Event *FactoryFeeConfiguration // Event containing the contract specifics and raw log

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
func (it *FactoryFeeConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryFeeConfiguration)
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
		it.Event = new(FactoryFeeConfiguration)
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
func (it *FactoryFeeConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryFeeConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryFeeConfiguration represents a FeeConfiguration event raised by the Factory contract.
type FactoryFeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeConfiguration is a free log retrieval operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactoryFilterer) FilterFeeConfiguration(opts *bind.FilterOpts) (*FactoryFeeConfigurationIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return &FactoryFeeConfigurationIterator{contract: _Factory.contract, event: "FeeConfiguration", logs: logs, sub: sub}, nil
}

// WatchFeeConfiguration is a free log subscription operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactoryFilterer) WatchFeeConfiguration(opts *bind.WatchOpts, sink chan<- *FactoryFeeConfiguration) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryFeeConfiguration)
				if err := _Factory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
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

// ParseFeeConfiguration is a log parse operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_Factory *FactoryFilterer) ParseFeeConfiguration(log types.Log) (*FactoryFeeConfiguration, error) {
	event := new(FactoryFeeConfiguration)
	if err := _Factory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryOwnerIterator is returned from FilterOwner and is used to iterate over the raw logs and unpacked data for Owner events raised by the Factory contract.
type FactoryOwnerIterator struct {
	Event *FactoryOwner // Event containing the contract specifics and raw log

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
func (it *FactoryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryOwner)
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
		it.Event = new(FactoryOwner)
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
func (it *FactoryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryOwner represents a Owner event raised by the Factory contract.
type FactoryOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwner is a free log retrieval operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Factory *FactoryFilterer) FilterOwner(opts *bind.FilterOpts, newOwner []common.Address) (*FactoryOwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryOwnerIterator{contract: _Factory.contract, event: "Owner", logs: logs, sub: sub}, nil
}

// WatchOwner is a free log subscription operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Factory *FactoryFilterer) WatchOwner(opts *bind.WatchOpts, sink chan<- *FactoryOwner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryOwner)
				if err := _Factory.contract.UnpackLog(event, "Owner", log); err != nil {
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

// ParseOwner is a log parse operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_Factory *FactoryFilterer) ParseOwner(log types.Log) (*FactoryOwner, error) {
	event := new(FactoryOwner)
	if err := _Factory.contract.UnpackLog(event, "Owner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryPoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the Factory contract.
type FactoryPoolIterator struct {
	Event *FactoryPool // Event containing the contract specifics and raw log

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
func (it *FactoryPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPool)
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
		it.Event = new(FactoryPool)
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
func (it *FactoryPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPool represents a Pool event raised by the Factory contract.
type FactoryPool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Factory *FactoryFilterer) FilterPool(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*FactoryPoolIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &FactoryPoolIterator{contract: _Factory.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Factory *FactoryFilterer) WatchPool(opts *bind.WatchOpts, sink chan<- *FactoryPool, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPool)
				if err := _Factory.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_Factory *FactoryFilterer) ParsePool(log types.Log) (*FactoryPool, error) {
	event := new(FactoryPool)
	if err := _Factory.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryVaultAddressIterator is returned from FilterVaultAddress and is used to iterate over the raw logs and unpacked data for VaultAddress events raised by the Factory contract.
type FactoryVaultAddressIterator struct {
	Event *FactoryVaultAddress // Event containing the contract specifics and raw log

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
func (it *FactoryVaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryVaultAddress)
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
		it.Event = new(FactoryVaultAddress)
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
func (it *FactoryVaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryVaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryVaultAddress represents a VaultAddress event raised by the Factory contract.
type FactoryVaultAddress struct {
	NewVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultAddress is a free log retrieval operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Factory *FactoryFilterer) FilterVaultAddress(opts *bind.FilterOpts, newVaultAddress []common.Address) (*FactoryVaultAddressIterator, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &FactoryVaultAddressIterator{contract: _Factory.contract, event: "VaultAddress", logs: logs, sub: sub}, nil
}

// WatchVaultAddress is a free log subscription operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Factory *FactoryFilterer) WatchVaultAddress(opts *bind.WatchOpts, sink chan<- *FactoryVaultAddress, newVaultAddress []common.Address) (event.Subscription, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryVaultAddress)
				if err := _Factory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
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

// ParseVaultAddress is a log parse operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_Factory *FactoryFilterer) ParseVaultAddress(log types.Log) (*FactoryVaultAddress, error) {
	event := new(FactoryVaultAddress)
	if err := _Factory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
