// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagenttokendeployer

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

// IAGENTTokenDeployerMetaData contains all meta data concerning the IAGENTTokenDeployer contract.
var IAGENTTokenDeployerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AGENTTokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"callWithValue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IAGENTTokenDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAGENTTokenDeployerMetaData.ABI instead.
var IAGENTTokenDeployerABI = IAGENTTokenDeployerMetaData.ABI

// IAGENTTokenDeployer is an auto generated Go binding around an Ethereum contract.
type IAGENTTokenDeployer struct {
	IAGENTTokenDeployerCaller     // Read-only binding to the contract
	IAGENTTokenDeployerTransactor // Write-only binding to the contract
	IAGENTTokenDeployerFilterer   // Log filterer for contract events
}

// IAGENTTokenDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAGENTTokenDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAGENTTokenDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAGENTTokenDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAGENTTokenDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAGENTTokenDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAGENTTokenDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAGENTTokenDeployerSession struct {
	Contract     *IAGENTTokenDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IAGENTTokenDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAGENTTokenDeployerCallerSession struct {
	Contract *IAGENTTokenDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IAGENTTokenDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAGENTTokenDeployerTransactorSession struct {
	Contract     *IAGENTTokenDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IAGENTTokenDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAGENTTokenDeployerRaw struct {
	Contract *IAGENTTokenDeployer // Generic contract binding to access the raw methods on
}

// IAGENTTokenDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAGENTTokenDeployerCallerRaw struct {
	Contract *IAGENTTokenDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// IAGENTTokenDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAGENTTokenDeployerTransactorRaw struct {
	Contract *IAGENTTokenDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAGENTTokenDeployer creates a new instance of IAGENTTokenDeployer, bound to a specific deployed contract.
func NewIAGENTTokenDeployer(address common.Address, backend bind.ContractBackend) (*IAGENTTokenDeployer, error) {
	contract, err := bindIAGENTTokenDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAGENTTokenDeployer{IAGENTTokenDeployerCaller: IAGENTTokenDeployerCaller{contract: contract}, IAGENTTokenDeployerTransactor: IAGENTTokenDeployerTransactor{contract: contract}, IAGENTTokenDeployerFilterer: IAGENTTokenDeployerFilterer{contract: contract}}, nil
}

// NewIAGENTTokenDeployerCaller creates a new read-only instance of IAGENTTokenDeployer, bound to a specific deployed contract.
func NewIAGENTTokenDeployerCaller(address common.Address, caller bind.ContractCaller) (*IAGENTTokenDeployerCaller, error) {
	contract, err := bindIAGENTTokenDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAGENTTokenDeployerCaller{contract: contract}, nil
}

// NewIAGENTTokenDeployerTransactor creates a new write-only instance of IAGENTTokenDeployer, bound to a specific deployed contract.
func NewIAGENTTokenDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAGENTTokenDeployerTransactor, error) {
	contract, err := bindIAGENTTokenDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAGENTTokenDeployerTransactor{contract: contract}, nil
}

// NewIAGENTTokenDeployerFilterer creates a new log filterer instance of IAGENTTokenDeployer, bound to a specific deployed contract.
func NewIAGENTTokenDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAGENTTokenDeployerFilterer, error) {
	contract, err := bindIAGENTTokenDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAGENTTokenDeployerFilterer{contract: contract}, nil
}

// bindIAGENTTokenDeployer binds a generic wrapper to an already deployed contract.
func bindIAGENTTokenDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAGENTTokenDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAGENTTokenDeployer.Contract.IAGENTTokenDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.IAGENTTokenDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.IAGENTTokenDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAGENTTokenDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.contract.Transact(opts, method, params...)
}

// GetToken is a free data retrieval call binding the contract method 0x155bf4e2.
//
// Solidity: function getToken(bytes32 salt) view returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerCaller) GetToken(opts *bind.CallOpts, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IAGENTTokenDeployer.contract.Call(opts, &out, "getToken", salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x155bf4e2.
//
// Solidity: function getToken(bytes32 salt) view returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerSession) GetToken(salt [32]byte) (common.Address, error) {
	return _IAGENTTokenDeployer.Contract.GetToken(&_IAGENTTokenDeployer.CallOpts, salt)
}

// GetToken is a free data retrieval call binding the contract method 0x155bf4e2.
//
// Solidity: function getToken(bytes32 salt) view returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerCallerSession) GetToken(salt [32]byte) (common.Address, error) {
	return _IAGENTTokenDeployer.Contract.GetToken(&_IAGENTTokenDeployer.CallOpts, salt)
}

// CallWithValue is a paid mutator transaction binding the contract method 0x04ac018a.
//
// Solidity: function callWithValue(address target, bytes data, uint256 value) payable returns(bytes)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactor) CallWithValue(opts *bind.TransactOpts, target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.contract.Transact(opts, "callWithValue", target, data, value)
}

// CallWithValue is a paid mutator transaction binding the contract method 0x04ac018a.
//
// Solidity: function callWithValue(address target, bytes data, uint256 value) payable returns(bytes)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerSession) CallWithValue(target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.CallWithValue(&_IAGENTTokenDeployer.TransactOpts, target, data, value)
}

// CallWithValue is a paid mutator transaction binding the contract method 0x04ac018a.
//
// Solidity: function callWithValue(address target, bytes data, uint256 value) payable returns(bytes)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactorSession) CallWithValue(target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.CallWithValue(&_IAGENTTokenDeployer.TransactOpts, target, data, value)
}

// CreateToken is a paid mutator transaction binding the contract method 0x0f99f157.
//
// Solidity: function createToken(bytes32 salt, string name, string symbol, uint256 amount) payable returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactor) CreateToken(opts *bind.TransactOpts, salt [32]byte, name string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.contract.Transact(opts, "createToken", salt, name, symbol, amount)
}

// CreateToken is a paid mutator transaction binding the contract method 0x0f99f157.
//
// Solidity: function createToken(bytes32 salt, string name, string symbol, uint256 amount) payable returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerSession) CreateToken(salt [32]byte, name string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.CreateToken(&_IAGENTTokenDeployer.TransactOpts, salt, name, symbol, amount)
}

// CreateToken is a paid mutator transaction binding the contract method 0x0f99f157.
//
// Solidity: function createToken(bytes32 salt, string name, string symbol, uint256 amount) payable returns(address)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactorSession) CreateToken(salt [32]byte, name string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.CreateToken(&_IAGENTTokenDeployer.TransactOpts, salt, name, symbol, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.Multicall(&_IAGENTTokenDeployer.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _IAGENTTokenDeployer.Contract.Multicall(&_IAGENTTokenDeployer.TransactOpts, data)
}

// IAGENTTokenDeployerAGENTTokenCreatedIterator is returned from FilterAGENTTokenCreated and is used to iterate over the raw logs and unpacked data for AGENTTokenCreated events raised by the IAGENTTokenDeployer contract.
type IAGENTTokenDeployerAGENTTokenCreatedIterator struct {
	Event *IAGENTTokenDeployerAGENTTokenCreated // Event containing the contract specifics and raw log

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
func (it *IAGENTTokenDeployerAGENTTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAGENTTokenDeployerAGENTTokenCreated)
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
		it.Event = new(IAGENTTokenDeployerAGENTTokenCreated)
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
func (it *IAGENTTokenDeployerAGENTTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAGENTTokenDeployerAGENTTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAGENTTokenDeployerAGENTTokenCreated represents a AGENTTokenCreated event raised by the IAGENTTokenDeployer contract.
type IAGENTTokenDeployerAGENTTokenCreated struct {
	Salt  [32]byte
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAGENTTokenCreated is a free log retrieval operation binding the contract event 0xb1ed6fee395fd3653fcafe0c34960115183da0ce5e99275b123893359ab04dc2.
//
// Solidity: event AGENTTokenCreated(bytes32 indexed salt, address indexed token)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerFilterer) FilterAGENTTokenCreated(opts *bind.FilterOpts, salt [][32]byte, token []common.Address) (*IAGENTTokenDeployerAGENTTokenCreatedIterator, error) {

	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IAGENTTokenDeployer.contract.FilterLogs(opts, "AGENTTokenCreated", saltRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IAGENTTokenDeployerAGENTTokenCreatedIterator{contract: _IAGENTTokenDeployer.contract, event: "AGENTTokenCreated", logs: logs, sub: sub}, nil
}

// WatchAGENTTokenCreated is a free log subscription operation binding the contract event 0xb1ed6fee395fd3653fcafe0c34960115183da0ce5e99275b123893359ab04dc2.
//
// Solidity: event AGENTTokenCreated(bytes32 indexed salt, address indexed token)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerFilterer) WatchAGENTTokenCreated(opts *bind.WatchOpts, sink chan<- *IAGENTTokenDeployerAGENTTokenCreated, salt [][32]byte, token []common.Address) (event.Subscription, error) {

	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IAGENTTokenDeployer.contract.WatchLogs(opts, "AGENTTokenCreated", saltRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAGENTTokenDeployerAGENTTokenCreated)
				if err := _IAGENTTokenDeployer.contract.UnpackLog(event, "AGENTTokenCreated", log); err != nil {
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

// ParseAGENTTokenCreated is a log parse operation binding the contract event 0xb1ed6fee395fd3653fcafe0c34960115183da0ce5e99275b123893359ab04dc2.
//
// Solidity: event AGENTTokenCreated(bytes32 indexed salt, address indexed token)
func (_IAGENTTokenDeployer *IAGENTTokenDeployerFilterer) ParseAGENTTokenCreated(log types.Log) (*IAGENTTokenDeployerAGENTTokenCreated, error) {
	event := new(IAGENTTokenDeployerAGENTTokenCreated)
	if err := _IAGENTTokenDeployer.contract.UnpackLog(event, "AGENTTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
