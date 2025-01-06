// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package brigde

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
)

// BrigdeMetaData contains all meta data concerning the Brigde contract.
var BrigdeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extddr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"BridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"externalAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"bridgeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"externalAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"bridgeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"safeMultisigContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainIdEth_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWrappedToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BrigdeABI is the input ABI used to generate the binding from.
// Deprecated: Use BrigdeMetaData.ABI instead.
var BrigdeABI = BrigdeMetaData.ABI

// Brigde is an auto generated Go binding around an Ethereum contract.
type Brigde struct {
	BrigdeCaller     // Read-only binding to the contract
	BrigdeTransactor // Write-only binding to the contract
	BrigdeFilterer   // Log filterer for contract events
}

// BrigdeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrigdeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrigdeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrigdeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrigdeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrigdeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrigdeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrigdeSession struct {
	Contract     *Brigde           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrigdeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrigdeCallerSession struct {
	Contract *BrigdeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrigdeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrigdeTransactorSession struct {
	Contract     *BrigdeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrigdeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrigdeRaw struct {
	Contract *Brigde // Generic contract binding to access the raw methods on
}

// BrigdeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrigdeCallerRaw struct {
	Contract *BrigdeCaller // Generic read-only contract binding to access the raw methods on
}

// BrigdeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrigdeTransactorRaw struct {
	Contract *BrigdeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrigde creates a new instance of Brigde, bound to a specific deployed contract.
func NewBrigde(address common.Address, backend bind.ContractBackend) (*Brigde, error) {
	contract, err := bindBrigde(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brigde{BrigdeCaller: BrigdeCaller{contract: contract}, BrigdeTransactor: BrigdeTransactor{contract: contract}, BrigdeFilterer: BrigdeFilterer{contract: contract}}, nil
}

// NewBrigdeCaller creates a new read-only instance of Brigde, bound to a specific deployed contract.
func NewBrigdeCaller(address common.Address, caller bind.ContractCaller) (*BrigdeCaller, error) {
	contract, err := bindBrigde(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrigdeCaller{contract: contract}, nil
}

// NewBrigdeTransactor creates a new write-only instance of Brigde, bound to a specific deployed contract.
func NewBrigdeTransactor(address common.Address, transactor bind.ContractTransactor) (*BrigdeTransactor, error) {
	contract, err := bindBrigde(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrigdeTransactor{contract: contract}, nil
}

// NewBrigdeFilterer creates a new log filterer instance of Brigde, bound to a specific deployed contract.
func NewBrigdeFilterer(address common.Address, filterer bind.ContractFilterer) (*BrigdeFilterer, error) {
	contract, err := bindBrigde(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrigdeFilterer{contract: contract}, nil
}

// bindBrigde binds a generic wrapper to an already deployed contract.
func bindBrigde(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrigdeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brigde *BrigdeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brigde.Contract.BrigdeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brigde *BrigdeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brigde.Contract.BrigdeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brigde *BrigdeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brigde.Contract.BrigdeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brigde *BrigdeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brigde.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brigde *BrigdeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brigde.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brigde *BrigdeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brigde.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Brigde *BrigdeCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brigde.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Brigde *BrigdeSession) ETHTOKEN() (common.Address, error) {
	return _Brigde.Contract.ETHTOKEN(&_Brigde.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Brigde *BrigdeCallerSession) ETHTOKEN() (common.Address, error) {
	return _Brigde.Contract.ETHTOKEN(&_Brigde.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brigde *BrigdeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brigde.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brigde *BrigdeSession) Owner() (common.Address, error) {
	return _Brigde.Contract.Owner(&_Brigde.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brigde *BrigdeCallerSession) Owner() (common.Address, error) {
	return _Brigde.Contract.Owner(&_Brigde.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Brigde *BrigdeTransactor) BridgeToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "bridgeToken", token, amount, externalAddr, destChainId)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Brigde *BrigdeSession) BridgeToken(token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.BridgeToken(&_Brigde.TransactOpts, token, amount, externalAddr, destChainId)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x0e93b35c.
//
// Solidity: function bridgeToken(address token, uint256 amount, string externalAddr, uint256 destChainId) returns()
func (_Brigde *BrigdeTransactorSession) BridgeToken(token common.Address, amount *big.Int, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.BridgeToken(&_Brigde.TransactOpts, token, amount, externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Brigde *BrigdeTransactor) BridgeToken0(opts *bind.TransactOpts, externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "bridgeToken0", externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Brigde *BrigdeSession) BridgeToken0(externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.BridgeToken0(&_Brigde.TransactOpts, externalAddr, destChainId)
}

// BridgeToken0 is a paid mutator transaction binding the contract method 0xd4546d23.
//
// Solidity: function bridgeToken(string externalAddr, uint256 destChainId) payable returns()
func (_Brigde *BrigdeTransactorSession) BridgeToken0(externalAddr string, destChainId *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.BridgeToken0(&_Brigde.TransactOpts, externalAddr, destChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address safeMultisigContractAddress, uint256 chainIdEth_) returns()
func (_Brigde *BrigdeTransactor) Initialize(opts *bind.TransactOpts, safeMultisigContractAddress common.Address, chainIdEth_ *big.Int) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "initialize", safeMultisigContractAddress, chainIdEth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address safeMultisigContractAddress, uint256 chainIdEth_) returns()
func (_Brigde *BrigdeSession) Initialize(safeMultisigContractAddress common.Address, chainIdEth_ *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Initialize(&_Brigde.TransactOpts, safeMultisigContractAddress, chainIdEth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address safeMultisigContractAddress, uint256 chainIdEth_) returns()
func (_Brigde *BrigdeTransactorSession) Initialize(safeMultisigContractAddress common.Address, chainIdEth_ *big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Initialize(&_Brigde.TransactOpts, safeMultisigContractAddress, chainIdEth_)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeTransactor) Mint(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "mint", tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Mint(&_Brigde.TransactOpts, tokens, recipients, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x5530f4a5.
//
// Solidity: function mint(address[] tokens, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeTransactorSession) Mint(tokens []common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Mint(&_Brigde.TransactOpts, tokens, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeTransactor) Mint0(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "mint0", token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Mint0(&_Brigde.TransactOpts, token, recipients, amounts)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa3bf277e.
//
// Solidity: function mint(address token, address[] recipients, uint256[] amounts) returns()
func (_Brigde *BrigdeTransactorSession) Mint0(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Brigde.Contract.Mint0(&_Brigde.TransactOpts, token, recipients, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brigde *BrigdeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brigde *BrigdeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brigde.Contract.RenounceOwnership(&_Brigde.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brigde *BrigdeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brigde.Contract.RenounceOwnership(&_Brigde.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brigde *BrigdeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Brigde.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brigde *BrigdeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brigde.Contract.TransferOwnership(&_Brigde.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brigde *BrigdeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brigde.Contract.TransferOwnership(&_Brigde.TransactOpts, newOwner)
}

// BrigdeBridgeTokenIterator is returned from FilterBridgeToken and is used to iterate over the raw logs and unpacked data for BridgeToken events raised by the Brigde contract.
type BrigdeBridgeTokenIterator struct {
	Event *BrigdeBridgeToken // Event containing the contract specifics and raw log

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
func (it *BrigdeBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrigdeBridgeToken)
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
		it.Event = new(BrigdeBridgeToken)
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
func (it *BrigdeBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrigdeBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrigdeBridgeToken represents a BridgeToken event raised by the Brigde contract.
type BrigdeBridgeToken struct {
	Token       common.Address
	Burner      common.Address
	Amount      *big.Int
	Extddr      string
	DestChainId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeToken is a free log retrieval operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Brigde *BrigdeFilterer) FilterBridgeToken(opts *bind.FilterOpts) (*BrigdeBridgeTokenIterator, error) {

	logs, sub, err := _Brigde.contract.FilterLogs(opts, "BridgeToken")
	if err != nil {
		return nil, err
	}
	return &BrigdeBridgeTokenIterator{contract: _Brigde.contract, event: "BridgeToken", logs: logs, sub: sub}, nil
}

// WatchBridgeToken is a free log subscription operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Brigde *BrigdeFilterer) WatchBridgeToken(opts *bind.WatchOpts, sink chan<- *BrigdeBridgeToken) (event.Subscription, error) {

	logs, sub, err := _Brigde.contract.WatchLogs(opts, "BridgeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrigdeBridgeToken)
				if err := _Brigde.contract.UnpackLog(event, "BridgeToken", log); err != nil {
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

// ParseBridgeToken is a log parse operation binding the contract event 0xc28e54186544d7357308b86c8319edd275e0db552d62381cf49f827791845c61.
//
// Solidity: event BridgeToken(address token, address burner, uint256 amount, string extddr, uint256 destChainId)
func (_Brigde *BrigdeFilterer) ParseBridgeToken(log types.Log) (*BrigdeBridgeToken, error) {
	event := new(BrigdeBridgeToken)
	if err := _Brigde.contract.UnpackLog(event, "BridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrigdeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Brigde contract.
type BrigdeInitializedIterator struct {
	Event *BrigdeInitialized // Event containing the contract specifics and raw log

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
func (it *BrigdeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrigdeInitialized)
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
		it.Event = new(BrigdeInitialized)
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
func (it *BrigdeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrigdeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrigdeInitialized represents a Initialized event raised by the Brigde contract.
type BrigdeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Brigde *BrigdeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BrigdeInitializedIterator, error) {

	logs, sub, err := _Brigde.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BrigdeInitializedIterator{contract: _Brigde.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Brigde *BrigdeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BrigdeInitialized) (event.Subscription, error) {

	logs, sub, err := _Brigde.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrigdeInitialized)
				if err := _Brigde.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Brigde *BrigdeFilterer) ParseInitialized(log types.Log) (*BrigdeInitialized, error) {
	event := new(BrigdeInitialized)
	if err := _Brigde.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrigdeMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Brigde contract.
type BrigdeMintIterator struct {
	Event *BrigdeMint // Event containing the contract specifics and raw log

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
func (it *BrigdeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrigdeMint)
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
		it.Event = new(BrigdeMint)
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
func (it *BrigdeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrigdeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrigdeMint represents a Mint event raised by the Brigde contract.
type BrigdeMint struct {
	Tokens     []common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) FilterMint(opts *bind.FilterOpts) (*BrigdeMintIterator, error) {

	logs, sub, err := _Brigde.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &BrigdeMintIterator{contract: _Brigde.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BrigdeMint) (event.Subscription, error) {

	logs, sub, err := _Brigde.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrigdeMint)
				if err := _Brigde.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xe9914506df53b6ba40090fea5ed4edb71623a51062de3125c2dc65b23de6d05e.
//
// Solidity: event Mint(address[] tokens, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) ParseMint(log types.Log) (*BrigdeMint, error) {
	event := new(BrigdeMint)
	if err := _Brigde.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrigdeMint0Iterator is returned from FilterMint0 and is used to iterate over the raw logs and unpacked data for Mint0 events raised by the Brigde contract.
type BrigdeMint0Iterator struct {
	Event *BrigdeMint0 // Event containing the contract specifics and raw log

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
func (it *BrigdeMint0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrigdeMint0)
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
		it.Event = new(BrigdeMint0)
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
func (it *BrigdeMint0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrigdeMint0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrigdeMint0 represents a Mint0 event raised by the Brigde contract.
type BrigdeMint0 struct {
	Token      common.Address
	Recipients []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint0 is a free log retrieval operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) FilterMint0(opts *bind.FilterOpts) (*BrigdeMint0Iterator, error) {

	logs, sub, err := _Brigde.contract.FilterLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return &BrigdeMint0Iterator{contract: _Brigde.contract, event: "Mint0", logs: logs, sub: sub}, nil
}

// WatchMint0 is a free log subscription operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) WatchMint0(opts *bind.WatchOpts, sink chan<- *BrigdeMint0) (event.Subscription, error) {

	logs, sub, err := _Brigde.contract.WatchLogs(opts, "Mint0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrigdeMint0)
				if err := _Brigde.contract.UnpackLog(event, "Mint0", log); err != nil {
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

// ParseMint0 is a log parse operation binding the contract event 0xa20ca4d8d83b89ff090c0ea7b3c3c600625d46681874e0c0d1e35a1d1d4964dd.
//
// Solidity: event Mint(address token, address[] recipients, uint256[] amounts)
func (_Brigde *BrigdeFilterer) ParseMint0(log types.Log) (*BrigdeMint0, error) {
	event := new(BrigdeMint0)
	if err := _Brigde.contract.UnpackLog(event, "Mint0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrigdeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Brigde contract.
type BrigdeOwnershipTransferredIterator struct {
	Event *BrigdeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrigdeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrigdeOwnershipTransferred)
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
		it.Event = new(BrigdeOwnershipTransferred)
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
func (it *BrigdeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrigdeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrigdeOwnershipTransferred represents a OwnershipTransferred event raised by the Brigde contract.
type BrigdeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brigde *BrigdeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrigdeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brigde.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrigdeOwnershipTransferredIterator{contract: _Brigde.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brigde *BrigdeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrigdeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brigde.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrigdeOwnershipTransferred)
				if err := _Brigde.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Brigde *BrigdeFilterer) ParseOwnershipTransferred(log types.Log) (*BrigdeOwnershipTransferred, error) {
	event := new(BrigdeOwnershipTransferred)
	if err := _Brigde.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
