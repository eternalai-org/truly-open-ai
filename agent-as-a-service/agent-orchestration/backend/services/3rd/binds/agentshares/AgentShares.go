// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agentshares

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

// AgentSharesMetaData contains all meta data concerning the AgentShares contract.
var AgentSharesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyMax\",\"type\":\"uint256\"}],\"name\":\"buyShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"deployToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBuyPriceAfterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSellPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getSellPriceAfterFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyMin\",\"type\":\"uint256\"}],\"name\":\"sellShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeDestination\",\"type\":\"address\"}],\"name\":\"setFeeDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolAdmin\",\"type\":\"address\"}],\"name\":\"setProtocolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercent\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDeployed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AgentSharesABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentSharesMetaData.ABI instead.
var AgentSharesABI = AgentSharesMetaData.ABI

// AgentShares is an auto generated Go binding around an Ethereum contract.
type AgentShares struct {
	AgentSharesCaller     // Read-only binding to the contract
	AgentSharesTransactor // Write-only binding to the contract
	AgentSharesFilterer   // Log filterer for contract events
}

// AgentSharesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentSharesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSharesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentSharesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSharesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentSharesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSharesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSharesSession struct {
	Contract     *AgentShares      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentSharesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentSharesCallerSession struct {
	Contract *AgentSharesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AgentSharesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentSharesTransactorSession struct {
	Contract     *AgentSharesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AgentSharesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentSharesRaw struct {
	Contract *AgentShares // Generic contract binding to access the raw methods on
}

// AgentSharesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentSharesCallerRaw struct {
	Contract *AgentSharesCaller // Generic read-only contract binding to access the raw methods on
}

// AgentSharesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentSharesTransactorRaw struct {
	Contract *AgentSharesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentShares creates a new instance of AgentShares, bound to a specific deployed contract.
func NewAgentShares(address common.Address, backend bind.ContractBackend) (*AgentShares, error) {
	contract, err := bindAgentShares(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentShares{AgentSharesCaller: AgentSharesCaller{contract: contract}, AgentSharesTransactor: AgentSharesTransactor{contract: contract}, AgentSharesFilterer: AgentSharesFilterer{contract: contract}}, nil
}

// NewAgentSharesCaller creates a new read-only instance of AgentShares, bound to a specific deployed contract.
func NewAgentSharesCaller(address common.Address, caller bind.ContractCaller) (*AgentSharesCaller, error) {
	contract, err := bindAgentShares(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSharesCaller{contract: contract}, nil
}

// NewAgentSharesTransactor creates a new write-only instance of AgentShares, bound to a specific deployed contract.
func NewAgentSharesTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentSharesTransactor, error) {
	contract, err := bindAgentShares(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSharesTransactor{contract: contract}, nil
}

// NewAgentSharesFilterer creates a new log filterer instance of AgentShares, bound to a specific deployed contract.
func NewAgentSharesFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentSharesFilterer, error) {
	contract, err := bindAgentShares(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentSharesFilterer{contract: contract}, nil
}

// bindAgentShares binds a generic wrapper to an already deployed contract.
func bindAgentShares(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentSharesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentShares *AgentSharesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentShares.Contract.AgentSharesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentShares *AgentSharesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentShares.Contract.AgentSharesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentShares *AgentSharesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentShares.Contract.AgentSharesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentShares *AgentSharesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentShares.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentShares *AgentSharesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentShares.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentShares *AgentSharesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentShares.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AgentShares *AgentSharesCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AgentShares *AgentSharesSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.BalanceOf(&_AgentShares.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.BalanceOf(&_AgentShares.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AgentShares *AgentSharesCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AgentShares *AgentSharesSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _AgentShares.Contract.BalanceOfBatch(&_AgentShares.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AgentShares *AgentSharesCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _AgentShares.Contract.BalanceOfBatch(&_AgentShares.CallOpts, accounts, ids)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_AgentShares *AgentSharesCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_AgentShares *AgentSharesSession) BaseToken() (common.Address, error) {
	return _AgentShares.Contract.BaseToken(&_AgentShares.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_AgentShares *AgentSharesCallerSession) BaseToken() (common.Address, error) {
	return _AgentShares.Contract.BaseToken(&_AgentShares.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xc157253d.
//
// Solidity: function getBuyPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCaller) GetBuyPrice(opts *bind.CallOpts, tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getBuyPrice", tokenId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0xc157253d.
//
// Solidity: function getBuyPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesSession) GetBuyPrice(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetBuyPrice(&_AgentShares.CallOpts, tokenId, amount)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xc157253d.
//
// Solidity: function getBuyPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) GetBuyPrice(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetBuyPrice(&_AgentShares.CallOpts, tokenId, amount)
}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x063a741f.
//
// Solidity: function getBuyPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCaller) GetBuyPriceAfterFee(opts *bind.CallOpts, tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getBuyPriceAfterFee", tokenId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x063a741f.
//
// Solidity: function getBuyPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesSession) GetBuyPriceAfterFee(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetBuyPriceAfterFee(&_AgentShares.CallOpts, tokenId, amount)
}

// GetBuyPriceAfterFee is a free data retrieval call binding the contract method 0x063a741f.
//
// Solidity: function getBuyPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) GetBuyPriceAfterFee(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetBuyPriceAfterFee(&_AgentShares.CallOpts, tokenId, amount)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_AgentShares *AgentSharesCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_AgentShares *AgentSharesSession) GetChainId() (*big.Int, error) {
	return _AgentShares.Contract.GetChainId(&_AgentShares.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_AgentShares *AgentSharesCallerSession) GetChainId() (*big.Int, error) {
	return _AgentShares.Contract.GetChainId(&_AgentShares.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_AgentShares *AgentSharesCaller) GetPrice(opts *bind.CallOpts, supply *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getPrice", supply, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_AgentShares *AgentSharesSession) GetPrice(supply *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetPrice(&_AgentShares.CallOpts, supply, amount)
}

// GetPrice is a free data retrieval call binding the contract method 0x5cf4ee91.
//
// Solidity: function getPrice(uint256 supply, uint256 amount) pure returns(uint256)
func (_AgentShares *AgentSharesCallerSession) GetPrice(supply *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetPrice(&_AgentShares.CallOpts, supply, amount)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x9477d85d.
//
// Solidity: function getSellPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCaller) GetSellPrice(opts *bind.CallOpts, tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getSellPrice", tokenId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPrice is a free data retrieval call binding the contract method 0x9477d85d.
//
// Solidity: function getSellPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesSession) GetSellPrice(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetSellPrice(&_AgentShares.CallOpts, tokenId, amount)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x9477d85d.
//
// Solidity: function getSellPrice(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) GetSellPrice(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetSellPrice(&_AgentShares.CallOpts, tokenId, amount)
}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0xcd9c7121.
//
// Solidity: function getSellPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCaller) GetSellPriceAfterFee(opts *bind.CallOpts, tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "getSellPriceAfterFee", tokenId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0xcd9c7121.
//
// Solidity: function getSellPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesSession) GetSellPriceAfterFee(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetSellPriceAfterFee(&_AgentShares.CallOpts, tokenId, amount)
}

// GetSellPriceAfterFee is a free data retrieval call binding the contract method 0xcd9c7121.
//
// Solidity: function getSellPriceAfterFee(uint256 tokenId, uint256 amount) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) GetSellPriceAfterFee(tokenId *big.Int, amount *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.GetSellPriceAfterFee(&_AgentShares.CallOpts, tokenId, amount)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AgentShares *AgentSharesCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AgentShares *AgentSharesSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _AgentShares.Contract.IsApprovedForAll(&_AgentShares.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AgentShares *AgentSharesCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _AgentShares.Contract.IsApprovedForAll(&_AgentShares.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentShares *AgentSharesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentShares *AgentSharesSession) Owner() (common.Address, error) {
	return _AgentShares.Contract.Owner(&_AgentShares.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentShares *AgentSharesCallerSession) Owner() (common.Address, error) {
	return _AgentShares.Contract.Owner(&_AgentShares.CallOpts)
}

// ProtocolAdmin is a free data retrieval call binding the contract method 0x420f6861.
//
// Solidity: function protocolAdmin() view returns(address)
func (_AgentShares *AgentSharesCaller) ProtocolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "protocolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolAdmin is a free data retrieval call binding the contract method 0x420f6861.
//
// Solidity: function protocolAdmin() view returns(address)
func (_AgentShares *AgentSharesSession) ProtocolAdmin() (common.Address, error) {
	return _AgentShares.Contract.ProtocolAdmin(&_AgentShares.CallOpts)
}

// ProtocolAdmin is a free data retrieval call binding the contract method 0x420f6861.
//
// Solidity: function protocolAdmin() view returns(address)
func (_AgentShares *AgentSharesCallerSession) ProtocolAdmin() (common.Address, error) {
	return _AgentShares.Contract.ProtocolAdmin(&_AgentShares.CallOpts)
}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_AgentShares *AgentSharesCaller) ProtocolFeeDestination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "protocolFeeDestination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_AgentShares *AgentSharesSession) ProtocolFeeDestination() (common.Address, error) {
	return _AgentShares.Contract.ProtocolFeeDestination(&_AgentShares.CallOpts)
}

// ProtocolFeeDestination is a free data retrieval call binding the contract method 0x4ce7957c.
//
// Solidity: function protocolFeeDestination() view returns(address)
func (_AgentShares *AgentSharesCallerSession) ProtocolFeeDestination() (common.Address, error) {
	return _AgentShares.Contract.ProtocolFeeDestination(&_AgentShares.CallOpts)
}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_AgentShares *AgentSharesCaller) ProtocolFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "protocolFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_AgentShares *AgentSharesSession) ProtocolFeePercent() (*big.Int, error) {
	return _AgentShares.Contract.ProtocolFeePercent(&_AgentShares.CallOpts)
}

// ProtocolFeePercent is a free data retrieval call binding the contract method 0xd6e6eb9f.
//
// Solidity: function protocolFeePercent() view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) ProtocolFeePercent() (*big.Int, error) {
	return _AgentShares.Contract.ProtocolFeePercent(&_AgentShares.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentShares *AgentSharesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentShares *AgentSharesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentShares.Contract.SupportsInterface(&_AgentShares.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AgentShares *AgentSharesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AgentShares.Contract.SupportsInterface(&_AgentShares.CallOpts, interfaceId)
}

// TokenDeployed is a free data retrieval call binding the contract method 0x1039a12d.
//
// Solidity: function tokenDeployed() view returns(address)
func (_AgentShares *AgentSharesCaller) TokenDeployed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "tokenDeployed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDeployed is a free data retrieval call binding the contract method 0x1039a12d.
//
// Solidity: function tokenDeployed() view returns(address)
func (_AgentShares *AgentSharesSession) TokenDeployed() (common.Address, error) {
	return _AgentShares.Contract.TokenDeployed(&_AgentShares.CallOpts)
}

// TokenDeployed is a free data retrieval call binding the contract method 0x1039a12d.
//
// Solidity: function tokenDeployed() view returns(address)
func (_AgentShares *AgentSharesCallerSession) TokenDeployed() (common.Address, error) {
	return _AgentShares.Contract.TokenDeployed(&_AgentShares.CallOpts)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_AgentShares *AgentSharesCaller) TokenSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "tokenSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_AgentShares *AgentSharesSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.TokenSupply(&_AgentShares.CallOpts, arg0)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_AgentShares *AgentSharesCallerSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _AgentShares.Contract.TokenSupply(&_AgentShares.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_AgentShares *AgentSharesCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_AgentShares *AgentSharesSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _AgentShares.Contract.Tokens(&_AgentShares.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_AgentShares *AgentSharesCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _AgentShares.Contract.Tokens(&_AgentShares.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AgentShares *AgentSharesCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AgentShares.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AgentShares *AgentSharesSession) Uri(arg0 *big.Int) (string, error) {
	return _AgentShares.Contract.Uri(&_AgentShares.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AgentShares *AgentSharesCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _AgentShares.Contract.Uri(&_AgentShares.CallOpts, arg0)
}

// BuyShares is a paid mutator transaction binding the contract method 0x7aabf58a.
//
// Solidity: function buyShares(uint256 tokenId, uint256 amount, uint256 supplyMax) returns()
func (_AgentShares *AgentSharesTransactor) BuyShares(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, supplyMax *big.Int) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "buyShares", tokenId, amount, supplyMax)
}

// BuyShares is a paid mutator transaction binding the contract method 0x7aabf58a.
//
// Solidity: function buyShares(uint256 tokenId, uint256 amount, uint256 supplyMax) returns()
func (_AgentShares *AgentSharesSession) BuyShares(tokenId *big.Int, amount *big.Int, supplyMax *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.BuyShares(&_AgentShares.TransactOpts, tokenId, amount, supplyMax)
}

// BuyShares is a paid mutator transaction binding the contract method 0x7aabf58a.
//
// Solidity: function buyShares(uint256 tokenId, uint256 amount, uint256 supplyMax) returns()
func (_AgentShares *AgentSharesTransactorSession) BuyShares(tokenId *big.Int, amount *big.Int, supplyMax *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.BuyShares(&_AgentShares.TransactOpts, tokenId, amount, supplyMax)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xc49662c5.
//
// Solidity: function claimToken(uint256 tokenId, address user) returns()
func (_AgentShares *AgentSharesTransactor) ClaimToken(opts *bind.TransactOpts, tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "claimToken", tokenId, user)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xc49662c5.
//
// Solidity: function claimToken(uint256 tokenId, address user) returns()
func (_AgentShares *AgentSharesSession) ClaimToken(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.ClaimToken(&_AgentShares.TransactOpts, tokenId, user)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xc49662c5.
//
// Solidity: function claimToken(uint256 tokenId, address user) returns()
func (_AgentShares *AgentSharesTransactorSession) ClaimToken(tokenId *big.Int, user common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.ClaimToken(&_AgentShares.TransactOpts, tokenId, user)
}

// DeployToken is a paid mutator transaction binding the contract method 0xda68ed12.
//
// Solidity: function deployToken(uint256 tokenId, string name, string symbol) returns(address)
func (_AgentShares *AgentSharesTransactor) DeployToken(opts *bind.TransactOpts, tokenId *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "deployToken", tokenId, name, symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0xda68ed12.
//
// Solidity: function deployToken(uint256 tokenId, string name, string symbol) returns(address)
func (_AgentShares *AgentSharesSession) DeployToken(tokenId *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _AgentShares.Contract.DeployToken(&_AgentShares.TransactOpts, tokenId, name, symbol)
}

// DeployToken is a paid mutator transaction binding the contract method 0xda68ed12.
//
// Solidity: function deployToken(uint256 tokenId, string name, string symbol) returns(address)
func (_AgentShares *AgentSharesTransactorSession) DeployToken(tokenId *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _AgentShares.Contract.DeployToken(&_AgentShares.TransactOpts, tokenId, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_AgentShares *AgentSharesTransactor) Initialize(opts *bind.TransactOpts, _baseToken common.Address) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "initialize", _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_AgentShares *AgentSharesSession) Initialize(_baseToken common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.Initialize(&_AgentShares.TransactOpts, _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_AgentShares *AgentSharesTransactorSession) Initialize(_baseToken common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.Initialize(&_AgentShares.TransactOpts, _baseToken)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AgentShares *AgentSharesTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AgentShares *AgentSharesSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AgentShares.Contract.Multicall(&_AgentShares.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_AgentShares *AgentSharesTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AgentShares.Contract.Multicall(&_AgentShares.TransactOpts, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentShares *AgentSharesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentShares *AgentSharesSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentShares.Contract.RenounceOwnership(&_AgentShares.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentShares *AgentSharesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentShares.Contract.RenounceOwnership(&_AgentShares.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AgentShares *AgentSharesTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AgentShares *AgentSharesSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.Contract.SafeBatchTransferFrom(&_AgentShares.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AgentShares *AgentSharesTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.Contract.SafeBatchTransferFrom(&_AgentShares.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AgentShares *AgentSharesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AgentShares *AgentSharesSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.Contract.SafeTransferFrom(&_AgentShares.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AgentShares *AgentSharesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AgentShares.Contract.SafeTransferFrom(&_AgentShares.TransactOpts, from, to, id, amount, data)
}

// SellShares is a paid mutator transaction binding the contract method 0xd67c6872.
//
// Solidity: function sellShares(uint256 tokenId, uint256 amount, uint256 supplyMin) returns()
func (_AgentShares *AgentSharesTransactor) SellShares(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, supplyMin *big.Int) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "sellShares", tokenId, amount, supplyMin)
}

// SellShares is a paid mutator transaction binding the contract method 0xd67c6872.
//
// Solidity: function sellShares(uint256 tokenId, uint256 amount, uint256 supplyMin) returns()
func (_AgentShares *AgentSharesSession) SellShares(tokenId *big.Int, amount *big.Int, supplyMin *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.SellShares(&_AgentShares.TransactOpts, tokenId, amount, supplyMin)
}

// SellShares is a paid mutator transaction binding the contract method 0xd67c6872.
//
// Solidity: function sellShares(uint256 tokenId, uint256 amount, uint256 supplyMin) returns()
func (_AgentShares *AgentSharesTransactorSession) SellShares(tokenId *big.Int, amount *big.Int, supplyMin *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.SellShares(&_AgentShares.TransactOpts, tokenId, amount, supplyMin)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentShares *AgentSharesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentShares *AgentSharesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentShares.Contract.SetApprovalForAll(&_AgentShares.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AgentShares *AgentSharesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AgentShares.Contract.SetApprovalForAll(&_AgentShares.TransactOpts, operator, approved)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_AgentShares *AgentSharesTransactor) SetFeeDestination(opts *bind.TransactOpts, _feeDestination common.Address) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "setFeeDestination", _feeDestination)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_AgentShares *AgentSharesSession) SetFeeDestination(_feeDestination common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.SetFeeDestination(&_AgentShares.TransactOpts, _feeDestination)
}

// SetFeeDestination is a paid mutator transaction binding the contract method 0xfbe53234.
//
// Solidity: function setFeeDestination(address _feeDestination) returns()
func (_AgentShares *AgentSharesTransactorSession) SetFeeDestination(_feeDestination common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.SetFeeDestination(&_AgentShares.TransactOpts, _feeDestination)
}

// SetProtocolAdmin is a paid mutator transaction binding the contract method 0x9a09b285.
//
// Solidity: function setProtocolAdmin(address _protocolAdmin) returns()
func (_AgentShares *AgentSharesTransactor) SetProtocolAdmin(opts *bind.TransactOpts, _protocolAdmin common.Address) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "setProtocolAdmin", _protocolAdmin)
}

// SetProtocolAdmin is a paid mutator transaction binding the contract method 0x9a09b285.
//
// Solidity: function setProtocolAdmin(address _protocolAdmin) returns()
func (_AgentShares *AgentSharesSession) SetProtocolAdmin(_protocolAdmin common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.SetProtocolAdmin(&_AgentShares.TransactOpts, _protocolAdmin)
}

// SetProtocolAdmin is a paid mutator transaction binding the contract method 0x9a09b285.
//
// Solidity: function setProtocolAdmin(address _protocolAdmin) returns()
func (_AgentShares *AgentSharesTransactorSession) SetProtocolAdmin(_protocolAdmin common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.SetProtocolAdmin(&_AgentShares.TransactOpts, _protocolAdmin)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_AgentShares *AgentSharesTransactor) SetProtocolFeePercent(opts *bind.TransactOpts, _feePercent *big.Int) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "setProtocolFeePercent", _feePercent)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_AgentShares *AgentSharesSession) SetProtocolFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.SetProtocolFeePercent(&_AgentShares.TransactOpts, _feePercent)
}

// SetProtocolFeePercent is a paid mutator transaction binding the contract method 0xa4983421.
//
// Solidity: function setProtocolFeePercent(uint256 _feePercent) returns()
func (_AgentShares *AgentSharesTransactorSession) SetProtocolFeePercent(_feePercent *big.Int) (*types.Transaction, error) {
	return _AgentShares.Contract.SetProtocolFeePercent(&_AgentShares.TransactOpts, _feePercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentShares *AgentSharesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentShares.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentShares *AgentSharesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.TransferOwnership(&_AgentShares.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentShares *AgentSharesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentShares.Contract.TransferOwnership(&_AgentShares.TransactOpts, newOwner)
}

// AgentSharesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AgentShares contract.
type AgentSharesApprovalForAllIterator struct {
	Event *AgentSharesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AgentSharesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesApprovalForAll)
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
		it.Event = new(AgentSharesApprovalForAll)
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
func (it *AgentSharesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesApprovalForAll represents a ApprovalForAll event raised by the AgentShares contract.
type AgentSharesApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_AgentShares *AgentSharesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*AgentSharesApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesApprovalForAllIterator{contract: _AgentShares.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_AgentShares *AgentSharesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AgentSharesApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesApprovalForAll)
				if err := _AgentShares.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_AgentShares *AgentSharesFilterer) ParseApprovalForAll(log types.Log) (*AgentSharesApprovalForAll, error) {
	event := new(AgentSharesApprovalForAll)
	if err := _AgentShares.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentShares contract.
type AgentSharesInitializedIterator struct {
	Event *AgentSharesInitialized // Event containing the contract specifics and raw log

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
func (it *AgentSharesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesInitialized)
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
		it.Event = new(AgentSharesInitialized)
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
func (it *AgentSharesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesInitialized represents a Initialized event raised by the AgentShares contract.
type AgentSharesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentShares *AgentSharesFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentSharesInitializedIterator, error) {

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentSharesInitializedIterator{contract: _AgentShares.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentShares *AgentSharesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentSharesInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesInitialized)
				if err := _AgentShares.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AgentShares *AgentSharesFilterer) ParseInitialized(log types.Log) (*AgentSharesInitialized, error) {
	event := new(AgentSharesInitialized)
	if err := _AgentShares.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentShares contract.
type AgentSharesOwnershipTransferredIterator struct {
	Event *AgentSharesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentSharesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesOwnershipTransferred)
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
		it.Event = new(AgentSharesOwnershipTransferred)
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
func (it *AgentSharesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesOwnershipTransferred represents a OwnershipTransferred event raised by the AgentShares contract.
type AgentSharesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentShares *AgentSharesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentSharesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesOwnershipTransferredIterator{contract: _AgentShares.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentShares *AgentSharesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentSharesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesOwnershipTransferred)
				if err := _AgentShares.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AgentShares *AgentSharesFilterer) ParseOwnershipTransferred(log types.Log) (*AgentSharesOwnershipTransferred, error) {
	event := new(AgentSharesOwnershipTransferred)
	if err := _AgentShares.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesTokenClaimedIterator is returned from FilterTokenClaimed and is used to iterate over the raw logs and unpacked data for TokenClaimed events raised by the AgentShares contract.
type AgentSharesTokenClaimedIterator struct {
	Event *AgentSharesTokenClaimed // Event containing the contract specifics and raw log

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
func (it *AgentSharesTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesTokenClaimed)
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
		it.Event = new(AgentSharesTokenClaimed)
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
func (it *AgentSharesTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesTokenClaimed represents a TokenClaimed event raised by the AgentShares contract.
type AgentSharesTokenClaimed struct {
	TokenId *big.Int
	Token   common.Address
	User    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenClaimed is a free log retrieval operation binding the contract event 0xa0126cb27d0e7a0ae1b6240b529bbdaaa427fd657cc72b15acf3ddbac115b66e.
//
// Solidity: event TokenClaimed(uint256 indexed tokenId, address indexed token, address indexed user, uint256 amount)
func (_AgentShares *AgentSharesFilterer) FilterTokenClaimed(opts *bind.FilterOpts, tokenId []*big.Int, token []common.Address, user []common.Address) (*AgentSharesTokenClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "TokenClaimed", tokenIdRule, tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesTokenClaimedIterator{contract: _AgentShares.contract, event: "TokenClaimed", logs: logs, sub: sub}, nil
}

// WatchTokenClaimed is a free log subscription operation binding the contract event 0xa0126cb27d0e7a0ae1b6240b529bbdaaa427fd657cc72b15acf3ddbac115b66e.
//
// Solidity: event TokenClaimed(uint256 indexed tokenId, address indexed token, address indexed user, uint256 amount)
func (_AgentShares *AgentSharesFilterer) WatchTokenClaimed(opts *bind.WatchOpts, sink chan<- *AgentSharesTokenClaimed, tokenId []*big.Int, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "TokenClaimed", tokenIdRule, tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesTokenClaimed)
				if err := _AgentShares.contract.UnpackLog(event, "TokenClaimed", log); err != nil {
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

// ParseTokenClaimed is a log parse operation binding the contract event 0xa0126cb27d0e7a0ae1b6240b529bbdaaa427fd657cc72b15acf3ddbac115b66e.
//
// Solidity: event TokenClaimed(uint256 indexed tokenId, address indexed token, address indexed user, uint256 amount)
func (_AgentShares *AgentSharesFilterer) ParseTokenClaimed(log types.Log) (*AgentSharesTokenClaimed, error) {
	event := new(AgentSharesTokenClaimed)
	if err := _AgentShares.contract.UnpackLog(event, "TokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesTokenDeployedIterator is returned from FilterTokenDeployed and is used to iterate over the raw logs and unpacked data for TokenDeployed events raised by the AgentShares contract.
type AgentSharesTokenDeployedIterator struct {
	Event *AgentSharesTokenDeployed // Event containing the contract specifics and raw log

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
func (it *AgentSharesTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesTokenDeployed)
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
		it.Event = new(AgentSharesTokenDeployed)
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
func (it *AgentSharesTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesTokenDeployed represents a TokenDeployed event raised by the AgentShares contract.
type AgentSharesTokenDeployed struct {
	TokenId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeployed is a free log retrieval operation binding the contract event 0xb0886d5b89c3540effacb4587a5c495172ace26a507fef96fb8ccb8376ae9df5.
//
// Solidity: event TokenDeployed(uint256 indexed tokenId, address indexed token)
func (_AgentShares *AgentSharesFilterer) FilterTokenDeployed(opts *bind.FilterOpts, tokenId []*big.Int, token []common.Address) (*AgentSharesTokenDeployedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "TokenDeployed", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesTokenDeployedIterator{contract: _AgentShares.contract, event: "TokenDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenDeployed is a free log subscription operation binding the contract event 0xb0886d5b89c3540effacb4587a5c495172ace26a507fef96fb8ccb8376ae9df5.
//
// Solidity: event TokenDeployed(uint256 indexed tokenId, address indexed token)
func (_AgentShares *AgentSharesFilterer) WatchTokenDeployed(opts *bind.WatchOpts, sink chan<- *AgentSharesTokenDeployed, tokenId []*big.Int, token []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "TokenDeployed", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesTokenDeployed)
				if err := _AgentShares.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
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

// ParseTokenDeployed is a log parse operation binding the contract event 0xb0886d5b89c3540effacb4587a5c495172ace26a507fef96fb8ccb8376ae9df5.
//
// Solidity: event TokenDeployed(uint256 indexed tokenId, address indexed token)
func (_AgentShares *AgentSharesFilterer) ParseTokenDeployed(log types.Log) (*AgentSharesTokenDeployed, error) {
	event := new(AgentSharesTokenDeployed)
	if err := _AgentShares.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the AgentShares contract.
type AgentSharesTradeIterator struct {
	Event *AgentSharesTrade // Event containing the contract specifics and raw log

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
func (it *AgentSharesTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesTrade)
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
		it.Event = new(AgentSharesTrade)
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
func (it *AgentSharesTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesTrade represents a Trade event raised by the AgentShares contract.
type AgentSharesTrade struct {
	Trader         common.Address
	TokenId        *big.Int
	IsBuy          bool
	ShareAmount    *big.Int
	EthAmount      *big.Int
	ProtocolAmount *big.Int
	Supply         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x12d0646903287d48eb117ac55a8bcc90d4357c4180221d5b33e83e73860440ec.
//
// Solidity: event Trade(address trader, uint256 tokenId, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolAmount, uint256 supply)
func (_AgentShares *AgentSharesFilterer) FilterTrade(opts *bind.FilterOpts) (*AgentSharesTradeIterator, error) {

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &AgentSharesTradeIterator{contract: _AgentShares.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x12d0646903287d48eb117ac55a8bcc90d4357c4180221d5b33e83e73860440ec.
//
// Solidity: event Trade(address trader, uint256 tokenId, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolAmount, uint256 supply)
func (_AgentShares *AgentSharesFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *AgentSharesTrade) (event.Subscription, error) {

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesTrade)
				if err := _AgentShares.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x12d0646903287d48eb117ac55a8bcc90d4357c4180221d5b33e83e73860440ec.
//
// Solidity: event Trade(address trader, uint256 tokenId, bool isBuy, uint256 shareAmount, uint256 ethAmount, uint256 protocolAmount, uint256 supply)
func (_AgentShares *AgentSharesFilterer) ParseTrade(log types.Log) (*AgentSharesTrade, error) {
	event := new(AgentSharesTrade)
	if err := _AgentShares.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the AgentShares contract.
type AgentSharesTransferBatchIterator struct {
	Event *AgentSharesTransferBatch // Event containing the contract specifics and raw log

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
func (it *AgentSharesTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesTransferBatch)
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
		it.Event = new(AgentSharesTransferBatch)
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
func (it *AgentSharesTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesTransferBatch represents a TransferBatch event raised by the AgentShares contract.
type AgentSharesTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_AgentShares *AgentSharesFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AgentSharesTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesTransferBatchIterator{contract: _AgentShares.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_AgentShares *AgentSharesFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *AgentSharesTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesTransferBatch)
				if err := _AgentShares.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_AgentShares *AgentSharesFilterer) ParseTransferBatch(log types.Log) (*AgentSharesTransferBatch, error) {
	event := new(AgentSharesTransferBatch)
	if err := _AgentShares.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the AgentShares contract.
type AgentSharesTransferSingleIterator struct {
	Event *AgentSharesTransferSingle // Event containing the contract specifics and raw log

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
func (it *AgentSharesTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesTransferSingle)
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
		it.Event = new(AgentSharesTransferSingle)
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
func (it *AgentSharesTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesTransferSingle represents a TransferSingle event raised by the AgentShares contract.
type AgentSharesTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AgentShares *AgentSharesFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AgentSharesTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesTransferSingleIterator{contract: _AgentShares.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AgentShares *AgentSharesFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *AgentSharesTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesTransferSingle)
				if err := _AgentShares.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AgentShares *AgentSharesFilterer) ParseTransferSingle(log types.Log) (*AgentSharesTransferSingle, error) {
	event := new(AgentSharesTransferSingle)
	if err := _AgentShares.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSharesURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the AgentShares contract.
type AgentSharesURIIterator struct {
	Event *AgentSharesURI // Event containing the contract specifics and raw log

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
func (it *AgentSharesURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSharesURI)
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
		it.Event = new(AgentSharesURI)
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
func (it *AgentSharesURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSharesURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSharesURI represents a URI event raised by the AgentShares contract.
type AgentSharesURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AgentShares *AgentSharesFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*AgentSharesURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AgentShares.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &AgentSharesURIIterator{contract: _AgentShares.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AgentShares *AgentSharesFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *AgentSharesURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AgentShares.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSharesURI)
				if err := _AgentShares.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AgentShares *AgentSharesFilterer) ParseURI(log types.Log) (*AgentSharesURI, error) {
	event := new(AgentSharesURI)
	if err := _AgentShares.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
