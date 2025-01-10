// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package orderpayment

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

// OrderpaymentMetaData contains all meta data concerning the Orderpayment contract.
var OrderpaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OrderPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OrderpaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderpaymentMetaData.ABI instead.
var OrderpaymentABI = OrderpaymentMetaData.ABI

// Orderpayment is an auto generated Go binding around an Ethereum contract.
type Orderpayment struct {
	OrderpaymentCaller     // Read-only binding to the contract
	OrderpaymentTransactor // Write-only binding to the contract
	OrderpaymentFilterer   // Log filterer for contract events
}

// OrderpaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderpaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderpaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderpaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderpaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderpaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderpaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderpaymentSession struct {
	Contract     *Orderpayment     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderpaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderpaymentCallerSession struct {
	Contract *OrderpaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OrderpaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderpaymentTransactorSession struct {
	Contract     *OrderpaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OrderpaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderpaymentRaw struct {
	Contract *Orderpayment // Generic contract binding to access the raw methods on
}

// OrderpaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderpaymentCallerRaw struct {
	Contract *OrderpaymentCaller // Generic read-only contract binding to access the raw methods on
}

// OrderpaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderpaymentTransactorRaw struct {
	Contract *OrderpaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderpayment creates a new instance of Orderpayment, bound to a specific deployed contract.
func NewOrderpayment(address common.Address, backend bind.ContractBackend) (*Orderpayment, error) {
	contract, err := bindOrderpayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderpayment{OrderpaymentCaller: OrderpaymentCaller{contract: contract}, OrderpaymentTransactor: OrderpaymentTransactor{contract: contract}, OrderpaymentFilterer: OrderpaymentFilterer{contract: contract}}, nil
}

// NewOrderpaymentCaller creates a new read-only instance of Orderpayment, bound to a specific deployed contract.
func NewOrderpaymentCaller(address common.Address, caller bind.ContractCaller) (*OrderpaymentCaller, error) {
	contract, err := bindOrderpayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderpaymentCaller{contract: contract}, nil
}

// NewOrderpaymentTransactor creates a new write-only instance of Orderpayment, bound to a specific deployed contract.
func NewOrderpaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderpaymentTransactor, error) {
	contract, err := bindOrderpayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderpaymentTransactor{contract: contract}, nil
}

// NewOrderpaymentFilterer creates a new log filterer instance of Orderpayment, bound to a specific deployed contract.
func NewOrderpaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderpaymentFilterer, error) {
	contract, err := bindOrderpayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderpaymentFilterer{contract: contract}, nil
}

// bindOrderpayment binds a generic wrapper to an already deployed contract.
func bindOrderpayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderpaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderpayment *OrderpaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orderpayment.Contract.OrderpaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderpayment *OrderpaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderpayment.Contract.OrderpaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderpayment *OrderpaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderpayment.Contract.OrderpaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderpayment *OrderpaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orderpayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderpayment *OrderpaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderpayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderpayment *OrderpaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderpayment.Contract.contract.Transact(opts, method, params...)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Orderpayment *OrderpaymentCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Orderpayment.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Orderpayment *OrderpaymentSession) BaseToken() (common.Address, error) {
	return _Orderpayment.Contract.BaseToken(&_Orderpayment.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Orderpayment *OrderpaymentCallerSession) BaseToken() (common.Address, error) {
	return _Orderpayment.Contract.BaseToken(&_Orderpayment.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Orderpayment *OrderpaymentCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Orderpayment.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Orderpayment *OrderpaymentSession) GetChainId() (*big.Int, error) {
	return _Orderpayment.Contract.GetChainId(&_Orderpayment.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_Orderpayment *OrderpaymentCallerSession) GetChainId() (*big.Int, error) {
	return _Orderpayment.Contract.GetChainId(&_Orderpayment.CallOpts)
}

// OrderIds is a free data retrieval call binding the contract method 0xf13886ec.
//
// Solidity: function orderIds(bytes32 ) view returns(bool)
func (_Orderpayment *OrderpaymentCaller) OrderIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Orderpayment.contract.Call(opts, &out, "orderIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrderIds is a free data retrieval call binding the contract method 0xf13886ec.
//
// Solidity: function orderIds(bytes32 ) view returns(bool)
func (_Orderpayment *OrderpaymentSession) OrderIds(arg0 [32]byte) (bool, error) {
	return _Orderpayment.Contract.OrderIds(&_Orderpayment.CallOpts, arg0)
}

// OrderIds is a free data retrieval call binding the contract method 0xf13886ec.
//
// Solidity: function orderIds(bytes32 ) view returns(bool)
func (_Orderpayment *OrderpaymentCallerSession) OrderIds(arg0 [32]byte) (bool, error) {
	return _Orderpayment.Contract.OrderIds(&_Orderpayment.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Orderpayment *OrderpaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Orderpayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Orderpayment *OrderpaymentSession) Owner() (common.Address, error) {
	return _Orderpayment.Contract.Owner(&_Orderpayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Orderpayment *OrderpaymentCallerSession) Owner() (common.Address, error) {
	return _Orderpayment.Contract.Owner(&_Orderpayment.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_Orderpayment *OrderpaymentTransactor) Initialize(opts *bind.TransactOpts, _baseToken common.Address) (*types.Transaction, error) {
	return _Orderpayment.contract.Transact(opts, "initialize", _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_Orderpayment *OrderpaymentSession) Initialize(_baseToken common.Address) (*types.Transaction, error) {
	return _Orderpayment.Contract.Initialize(&_Orderpayment.TransactOpts, _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _baseToken) returns()
func (_Orderpayment *OrderpaymentTransactorSession) Initialize(_baseToken common.Address) (*types.Transaction, error) {
	return _Orderpayment.Contract.Initialize(&_Orderpayment.TransactOpts, _baseToken)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Orderpayment *OrderpaymentTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Orderpayment.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Orderpayment *OrderpaymentSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Orderpayment.Contract.Multicall(&_Orderpayment.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Orderpayment *OrderpaymentTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Orderpayment.Contract.Multicall(&_Orderpayment.TransactOpts, data)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7473df4c.
//
// Solidity: function payOrder(bytes32 orderId, address seller, uint256 amount) returns()
func (_Orderpayment *OrderpaymentTransactor) PayOrder(opts *bind.TransactOpts, orderId [32]byte, seller common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Orderpayment.contract.Transact(opts, "payOrder", orderId, seller, amount)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7473df4c.
//
// Solidity: function payOrder(bytes32 orderId, address seller, uint256 amount) returns()
func (_Orderpayment *OrderpaymentSession) PayOrder(orderId [32]byte, seller common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Orderpayment.Contract.PayOrder(&_Orderpayment.TransactOpts, orderId, seller, amount)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7473df4c.
//
// Solidity: function payOrder(bytes32 orderId, address seller, uint256 amount) returns()
func (_Orderpayment *OrderpaymentTransactorSession) PayOrder(orderId [32]byte, seller common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Orderpayment.Contract.PayOrder(&_Orderpayment.TransactOpts, orderId, seller, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderpayment *OrderpaymentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderpayment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderpayment *OrderpaymentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderpayment.Contract.RenounceOwnership(&_Orderpayment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderpayment *OrderpaymentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderpayment.Contract.RenounceOwnership(&_Orderpayment.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Orderpayment *OrderpaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Orderpayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Orderpayment *OrderpaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Orderpayment.Contract.TransferOwnership(&_Orderpayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Orderpayment *OrderpaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Orderpayment.Contract.TransferOwnership(&_Orderpayment.TransactOpts, newOwner)
}

// OrderpaymentInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Orderpayment contract.
type OrderpaymentInitializedIterator struct {
	Event *OrderpaymentInitialized // Event containing the contract specifics and raw log

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
func (it *OrderpaymentInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderpaymentInitialized)
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
		it.Event = new(OrderpaymentInitialized)
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
func (it *OrderpaymentInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderpaymentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderpaymentInitialized represents a Initialized event raised by the Orderpayment contract.
type OrderpaymentInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Orderpayment *OrderpaymentFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrderpaymentInitializedIterator, error) {

	logs, sub, err := _Orderpayment.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrderpaymentInitializedIterator{contract: _Orderpayment.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Orderpayment *OrderpaymentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrderpaymentInitialized) (event.Subscription, error) {

	logs, sub, err := _Orderpayment.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderpaymentInitialized)
				if err := _Orderpayment.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Orderpayment *OrderpaymentFilterer) ParseInitialized(log types.Log) (*OrderpaymentInitialized, error) {
	event := new(OrderpaymentInitialized)
	if err := _Orderpayment.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderpaymentOrderPaidIterator is returned from FilterOrderPaid and is used to iterate over the raw logs and unpacked data for OrderPaid events raised by the Orderpayment contract.
type OrderpaymentOrderPaidIterator struct {
	Event *OrderpaymentOrderPaid // Event containing the contract specifics and raw log

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
func (it *OrderpaymentOrderPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderpaymentOrderPaid)
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
		it.Event = new(OrderpaymentOrderPaid)
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
func (it *OrderpaymentOrderPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderpaymentOrderPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderpaymentOrderPaid represents a OrderPaid event raised by the Orderpayment contract.
type OrderpaymentOrderPaid struct {
	OrderId [32]byte
	Buyer   common.Address
	Seller  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderPaid is a free log retrieval operation binding the contract event 0xc2522570932e6dff27df2e5c31cfd70be3653d564375e29575d4360aafca4eb5.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed buyer, address indexed seller, uint256 amount)
func (_Orderpayment *OrderpaymentFilterer) FilterOrderPaid(opts *bind.FilterOpts, orderId [][32]byte, buyer []common.Address, seller []common.Address) (*OrderpaymentOrderPaidIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Orderpayment.contract.FilterLogs(opts, "OrderPaid", orderIdRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &OrderpaymentOrderPaidIterator{contract: _Orderpayment.contract, event: "OrderPaid", logs: logs, sub: sub}, nil
}

// WatchOrderPaid is a free log subscription operation binding the contract event 0xc2522570932e6dff27df2e5c31cfd70be3653d564375e29575d4360aafca4eb5.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed buyer, address indexed seller, uint256 amount)
func (_Orderpayment *OrderpaymentFilterer) WatchOrderPaid(opts *bind.WatchOpts, sink chan<- *OrderpaymentOrderPaid, orderId [][32]byte, buyer []common.Address, seller []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Orderpayment.contract.WatchLogs(opts, "OrderPaid", orderIdRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderpaymentOrderPaid)
				if err := _Orderpayment.contract.UnpackLog(event, "OrderPaid", log); err != nil {
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

// ParseOrderPaid is a log parse operation binding the contract event 0xc2522570932e6dff27df2e5c31cfd70be3653d564375e29575d4360aafca4eb5.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed buyer, address indexed seller, uint256 amount)
func (_Orderpayment *OrderpaymentFilterer) ParseOrderPaid(log types.Log) (*OrderpaymentOrderPaid, error) {
	event := new(OrderpaymentOrderPaid)
	if err := _Orderpayment.contract.UnpackLog(event, "OrderPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderpaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Orderpayment contract.
type OrderpaymentOwnershipTransferredIterator struct {
	Event *OrderpaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrderpaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderpaymentOwnershipTransferred)
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
		it.Event = new(OrderpaymentOwnershipTransferred)
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
func (it *OrderpaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderpaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderpaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Orderpayment contract.
type OrderpaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Orderpayment *OrderpaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrderpaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderpayment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderpaymentOwnershipTransferredIterator{contract: _Orderpayment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Orderpayment *OrderpaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrderpaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderpayment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderpaymentOwnershipTransferred)
				if err := _Orderpayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Orderpayment *OrderpaymentFilterer) ParseOwnershipTransferred(log types.Log) (*OrderpaymentOwnershipTransferred, error) {
	event := new(OrderpaymentOwnershipTransferred)
	if err := _Orderpayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
