// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisend

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

// MultiSendSend is an auto generated low-level Go binding around an user-defined struct.
type MultiSendSend struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
}

// MultiSendMetaData contains all meta data concerning the MultiSend contract.
var MultiSendMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiERC20Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structMultiSend.Send[]\",\"name\":\"sends\",\"type\":\"tuple[]\"}],\"name\":\"multiSend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461005b5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36108c590816100618239f35b600080fdfe60406080815260049081361015610020575b5050361561001e57600080fd5b005b600091823560e01c806335a21728146103e757806369328dec14610331578063715018a6146102d75780638da5cb5b146102ab578063ee15882b146101345763f2fde38b1461006f5750610011565b3461013057602036600319011261013057610088610490565b906100916104dc565b6001600160a01b039182169283156100de57505082546001600160a01b0319811683178455167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b8280fd5b50602092836003193601126102a85767ffffffffffffffff8235818111610130573660238201121561013057808401359182116101305760249283820191843691606086020101116102a85791908280925b8084106101d7575050505034036101a05750505160018152f35b60649360149293519362461bcd60e51b8552840152820152731b5a5cdb585d18da081cd95b9908185b5bdd5b9d60621b6044820152fd5b9293919290916001600160a01b03806101f96101f4888689610559565b61057f565b16610266575061022b6102178a610211888689610559565b0161057f565b89610223888689610559565b01359061083b565b87610237868487610559565b013581018091116102545761024c9094610534565b929190610186565b634e487b7160e01b8352601187528583fd5b946102a361024c929661027d6101f484878a610559565b1661028d8c61021185888b610559565b8b61029985888b610559565b0135913390610593565b610534565b80fd5b5050346102d357816003193601126102d357905490516001600160a01b039091168152602090f35b5080fd5b83346102a857806003193601126102a8576102f06104dc565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b509190346102d35760603660031901126102d35761034d610490565b60243591604435916001600160a01b0380841692918385036103e3576103716104dc565b1690816103895750505061038692935061083b565b80f35b809294919350519363a9059cbb60e01b60208601526024850152604484015260448352608083019083821067ffffffffffffffff8311176103d05761038694955052610604565b634e487b7160e01b855260418652602485fd5b8680fd5b8382346102d35760603660031901126102d357610402610490565b9067ffffffffffffffff60243581811161048c5761042390369084016104ab565b9190926044359182116104885761043f919493943691016104ab565b6001600160a01b03909316929091855b81811061045a578680f35b806102a361046f6101f461048394868b61082b565b61047a83878961082b565b35903389610593565b61044f565b8580fd5b8480fd5b600435906001600160a01b03821682036104a657565b600080fd5b9181601f840112156104a65782359167ffffffffffffffff83116104a6576020808501948460051b0101116104a657565b6000546001600160a01b031633036104f057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60001981146105435760010190565b634e487b7160e01b600052601160045260246000fd5b9190811015610569576060020190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b03811681036104a65790565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815260a081019181831067ffffffffffffffff8411176105ee576105ec92604052610604565b565b634e487b7160e01b600052604160045260246000fd5b60018060a01b031690604051604081019080821067ffffffffffffffff8311176105ee57610676916040526020938482527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858301526000808587829751910182855af161067061070c565b91610757565b8051918215918483156106e8575b5050509050156106915750565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b9193818094500103126102d3578201519081151582036102a8575080388084610684565b3d156107525767ffffffffffffffff903d8281116105ee5760405192601f8201601f19908116603f01168401908111848210176105ee5760405282523d6000602084013e565b606090565b919290156107b9575081511561076b575090565b3b156107745790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156107cc5750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b828510610812575050604492506000838284010152601f80199101168101030190fd5b84810182015186860160440152938101938593506107ef565b91908110156105695760051b0190565b600080809381935af161084c61070c565b501561085457565b60405162461bcd60e51b81526020600482015260136024820152721d1c985b9cd9995c88195d1a0819985a5b1959606a1b6044820152606490fdfea26469706673582212204e76d9ede77380337a919bdbe4566fcfd7e3dd9d0dd6b7f5258b4cdf2a86958964736f6c63430008130033",
}

// MultiSendABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSendMetaData.ABI instead.
var MultiSendABI = MultiSendMetaData.ABI

// MultiSendBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiSendMetaData.Bin instead.
var MultiSendBin = MultiSendMetaData.Bin

// DeployMultiSend deploys a new Ethereum contract, binding an instance of MultiSend to it.
func DeployMultiSend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiSend, error) {
	parsed, err := MultiSendMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiSendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSend{MultiSendCaller: MultiSendCaller{contract: contract}, MultiSendTransactor: MultiSendTransactor{contract: contract}, MultiSendFilterer: MultiSendFilterer{contract: contract}}, nil
}

// MultiSend is an auto generated Go binding around an Ethereum contract.
type MultiSend struct {
	MultiSendCaller     // Read-only binding to the contract
	MultiSendTransactor // Write-only binding to the contract
	MultiSendFilterer   // Log filterer for contract events
}

// MultiSendCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSendSession struct {
	Contract     *MultiSend        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSendCallerSession struct {
	Contract *MultiSendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultiSendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSendTransactorSession struct {
	Contract     *MultiSendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultiSendRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSendRaw struct {
	Contract *MultiSend // Generic contract binding to access the raw methods on
}

// MultiSendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSendCallerRaw struct {
	Contract *MultiSendCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSendTransactorRaw struct {
	Contract *MultiSendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSend creates a new instance of MultiSend, bound to a specific deployed contract.
func NewMultiSend(address common.Address, backend bind.ContractBackend) (*MultiSend, error) {
	contract, err := bindMultiSend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSend{MultiSendCaller: MultiSendCaller{contract: contract}, MultiSendTransactor: MultiSendTransactor{contract: contract}, MultiSendFilterer: MultiSendFilterer{contract: contract}}, nil
}

// NewMultiSendCaller creates a new read-only instance of MultiSend, bound to a specific deployed contract.
func NewMultiSendCaller(address common.Address, caller bind.ContractCaller) (*MultiSendCaller, error) {
	contract, err := bindMultiSend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSendCaller{contract: contract}, nil
}

// NewMultiSendTransactor creates a new write-only instance of MultiSend, bound to a specific deployed contract.
func NewMultiSendTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSendTransactor, error) {
	contract, err := bindMultiSend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSendTransactor{contract: contract}, nil
}

// NewMultiSendFilterer creates a new log filterer instance of MultiSend, bound to a specific deployed contract.
func NewMultiSendFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSendFilterer, error) {
	contract, err := bindMultiSend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSendFilterer{contract: contract}, nil
}

// bindMultiSend binds a generic wrapper to an already deployed contract.
func bindMultiSend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiSendMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSend *MultiSendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSend.Contract.MultiSendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSend *MultiSendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiSendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSend *MultiSendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiSendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSend *MultiSendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSend *MultiSendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSend *MultiSendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSend.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiSend *MultiSendCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiSend.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiSend *MultiSendSession) Owner() (common.Address, error) {
	return _MultiSend.Contract.Owner(&_MultiSend.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MultiSend *MultiSendCallerSession) Owner() (common.Address, error) {
	return _MultiSend.Contract.Owner(&_MultiSend.CallOpts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_MultiSend *MultiSendTransactor) MultiERC20Transfer(opts *bind.TransactOpts, _token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MultiSend.contract.Transact(opts, "multiERC20Transfer", _token, _addresses, _amounts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_MultiSend *MultiSendSession) MultiERC20Transfer(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiERC20Transfer(&_MultiSend.TransactOpts, _token, _addresses, _amounts)
}

// MultiERC20Transfer is a paid mutator transaction binding the contract method 0x35a21728.
//
// Solidity: function multiERC20Transfer(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_MultiSend *MultiSendTransactorSession) MultiERC20Transfer(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiERC20Transfer(&_MultiSend.TransactOpts, _token, _addresses, _amounts)
}

// MultiSend is a paid mutator transaction binding the contract method 0xee15882b.
//
// Solidity: function multiSend((address,address,uint256)[] sends) payable returns(bool)
func (_MultiSend *MultiSendTransactor) MultiSend(opts *bind.TransactOpts, sends []MultiSendSend) (*types.Transaction, error) {
	return _MultiSend.contract.Transact(opts, "multiSend", sends)
}

// MultiSend is a paid mutator transaction binding the contract method 0xee15882b.
//
// Solidity: function multiSend((address,address,uint256)[] sends) payable returns(bool)
func (_MultiSend *MultiSendSession) MultiSend(sends []MultiSendSend) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiSend(&_MultiSend.TransactOpts, sends)
}

// MultiSend is a paid mutator transaction binding the contract method 0xee15882b.
//
// Solidity: function multiSend((address,address,uint256)[] sends) payable returns(bool)
func (_MultiSend *MultiSendTransactorSession) MultiSend(sends []MultiSendSend) (*types.Transaction, error) {
	return _MultiSend.Contract.MultiSend(&_MultiSend.TransactOpts, sends)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiSend *MultiSendTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSend.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiSend *MultiSendSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultiSend.Contract.RenounceOwnership(&_MultiSend.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MultiSend *MultiSendTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MultiSend.Contract.RenounceOwnership(&_MultiSend.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiSend *MultiSendTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSend.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiSend *MultiSendSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiSend.Contract.TransferOwnership(&_MultiSend.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MultiSend *MultiSendTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MultiSend.Contract.TransferOwnership(&_MultiSend.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address to) returns()
func (_MultiSend *MultiSendTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MultiSend.contract.Transact(opts, "withdraw", token, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address to) returns()
func (_MultiSend *MultiSendSession) Withdraw(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MultiSend.Contract.Withdraw(&_MultiSend.TransactOpts, token, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address to) returns()
func (_MultiSend *MultiSendTransactorSession) Withdraw(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MultiSend.Contract.Withdraw(&_MultiSend.TransactOpts, token, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSend *MultiSendTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSend.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSend *MultiSendSession) Receive() (*types.Transaction, error) {
	return _MultiSend.Contract.Receive(&_MultiSend.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSend *MultiSendTransactorSession) Receive() (*types.Transaction, error) {
	return _MultiSend.Contract.Receive(&_MultiSend.TransactOpts)
}

// MultiSendOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MultiSend contract.
type MultiSendOwnershipTransferredIterator struct {
	Event *MultiSendOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MultiSendOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSendOwnershipTransferred)
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
		it.Event = new(MultiSendOwnershipTransferred)
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
func (it *MultiSendOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSendOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSendOwnershipTransferred represents a OwnershipTransferred event raised by the MultiSend contract.
type MultiSendOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultiSend *MultiSendFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MultiSendOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiSend.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSendOwnershipTransferredIterator{contract: _MultiSend.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MultiSend *MultiSendFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MultiSendOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MultiSend.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSendOwnershipTransferred)
				if err := _MultiSend.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MultiSend *MultiSendFilterer) ParseOwnershipTransferred(log types.Log) (*MultiSendOwnershipTransferred, error) {
	event := new(MultiSendOwnershipTransferred)
	if err := _MultiSend.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
