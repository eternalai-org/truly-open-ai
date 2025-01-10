// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sunpumplaunchpad

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

// SunpumplaunchpadMetaData contains all meta data concerning the Sunpumplaunchpad contract.
var SunpumplaunchpadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LaunchPending\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"oldLauncher\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newLauncher\",\"type\":\"address\"}],\"name\":\"LauncherChanged\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"MinTxFeeSet\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"MintFeeSet\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"oldOperator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"PurchaseFeeSet\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SaleFeeSet\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"tokenIndex\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"TokenCreate\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenLaunched\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"name\":\"tokenReserve\",\"type\":\"uint256\"}],\"name\":\"TokenPurchased\",\"type\":\"event\"},{\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSold\",\"type\":\"event\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"LAUNCH_FEE\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"LAUNCH_THRESHOLD\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"LAUNCH_TRX_RESERVE\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"TOKEN_SUPPLY\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"TOTAL_SALE\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"VIRTUAL_TOKEN_RESERVE_AMOUNT\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"VIRTUAL_TRX_RESERVE_AMOUNT\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"_becomeNewImplementation\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"name\":\"acceptOwner\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"createAndInitPurchase\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"deadAddress\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"getExactTokenAmountForPurchase\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"trxAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"getExactTokenAmountForPurchaseWithFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"name\":\"getExactTrxAmountForSale\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"name\":\"getExactTrxAmountForSaleWithFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenAmountByPurchase\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenAmountByPurchaseWithFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenState\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"trxAmount\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTrxAmountBySale\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"trxAmount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTrxAmountBySaleWithFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"implementation\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\"},{\"name\":\"_v2Router\",\"type\":\"address\"},{\"name\":\"_salefee\",\"type\":\"uint256\"},{\"name\":\"_purchasefee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"launchFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"launchToDEX\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"launcher\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"maxPurachaseAmount\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"minTxFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"mintFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"operator\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"owner\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"bool\"}],\"name\":\"pause\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"name\":\"pausePad\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"pendingImplementation\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"pendingOwner\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"purchaseFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"AmountMin\",\"type\":\"uint256\"}],\"name\":\"purchaseToken\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"renounceTokenOwnership\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"name\":\"rerunPad\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"saleFee\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"name\":\"AmountMin\",\"type\":\"uint256\"}],\"name\":\"saleToken\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newLauncher\",\"type\":\"address\"}],\"name\":\"setLauncher\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setMinTxFee\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_newMintFee\",\"type\":\"uint256\"},{\"name\":\"_newMinTxFee\",\"type\":\"uint256\"}],\"name\":\"setMintAndMinTxFee\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setMintFee\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newOp\",\"type\":\"address\"}],\"name\":\"setOperator\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setPurchaseFee\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setSaleFee\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setVault\",\"stateMutability\":\"Nonpayable\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"inputs\":[{\"type\":\"uint256\"}],\"name\":\"tokenAddress\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"uint256\"}],\"name\":\"tokenCount\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"inputs\":[{\"type\":\"address\"}],\"name\":\"tokenCreator\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"v2Router\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"type\":\"address\"}],\"name\":\"vault\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"outputs\":[{\"name\":\"TRXReserve\",\"type\":\"uint256\"},{\"name\":\"TokenReserve\",\"type\":\"uint256\"},{\"name\":\"launched\",\"type\":\"bool\"}],\"inputs\":[{\"type\":\"address\"}],\"name\":\"virtualPools\",\"stateMutability\":\"View\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SunpumplaunchpadABI is the input ABI used to generate the binding from.
// Deprecated: Use SunpumplaunchpadMetaData.ABI instead.
var SunpumplaunchpadABI = SunpumplaunchpadMetaData.ABI

// Sunpumplaunchpad is an auto generated Go binding around an Ethereum contract.
type Sunpumplaunchpad struct {
	SunpumplaunchpadCaller     // Read-only binding to the contract
	SunpumplaunchpadTransactor // Write-only binding to the contract
	SunpumplaunchpadFilterer   // Log filterer for contract events
}

// SunpumplaunchpadCaller is an auto generated read-only Go binding around an Ethereum contract.
type SunpumplaunchpadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SunpumplaunchpadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SunpumplaunchpadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SunpumplaunchpadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SunpumplaunchpadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SunpumplaunchpadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SunpumplaunchpadSession struct {
	Contract     *Sunpumplaunchpad // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SunpumplaunchpadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SunpumplaunchpadCallerSession struct {
	Contract *SunpumplaunchpadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SunpumplaunchpadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SunpumplaunchpadTransactorSession struct {
	Contract     *SunpumplaunchpadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SunpumplaunchpadRaw is an auto generated low-level Go binding around an Ethereum contract.
type SunpumplaunchpadRaw struct {
	Contract *Sunpumplaunchpad // Generic contract binding to access the raw methods on
}

// SunpumplaunchpadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SunpumplaunchpadCallerRaw struct {
	Contract *SunpumplaunchpadCaller // Generic read-only contract binding to access the raw methods on
}

// SunpumplaunchpadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SunpumplaunchpadTransactorRaw struct {
	Contract *SunpumplaunchpadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSunpumplaunchpad creates a new instance of Sunpumplaunchpad, bound to a specific deployed contract.
func NewSunpumplaunchpad(address common.Address, backend bind.ContractBackend) (*Sunpumplaunchpad, error) {
	contract, err := bindSunpumplaunchpad(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sunpumplaunchpad{SunpumplaunchpadCaller: SunpumplaunchpadCaller{contract: contract}, SunpumplaunchpadTransactor: SunpumplaunchpadTransactor{contract: contract}, SunpumplaunchpadFilterer: SunpumplaunchpadFilterer{contract: contract}}, nil
}

// NewSunpumplaunchpadCaller creates a new read-only instance of Sunpumplaunchpad, bound to a specific deployed contract.
func NewSunpumplaunchpadCaller(address common.Address, caller bind.ContractCaller) (*SunpumplaunchpadCaller, error) {
	contract, err := bindSunpumplaunchpad(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadCaller{contract: contract}, nil
}

// NewSunpumplaunchpadTransactor creates a new write-only instance of Sunpumplaunchpad, bound to a specific deployed contract.
func NewSunpumplaunchpadTransactor(address common.Address, transactor bind.ContractTransactor) (*SunpumplaunchpadTransactor, error) {
	contract, err := bindSunpumplaunchpad(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadTransactor{contract: contract}, nil
}

// NewSunpumplaunchpadFilterer creates a new log filterer instance of Sunpumplaunchpad, bound to a specific deployed contract.
func NewSunpumplaunchpadFilterer(address common.Address, filterer bind.ContractFilterer) (*SunpumplaunchpadFilterer, error) {
	contract, err := bindSunpumplaunchpad(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadFilterer{contract: contract}, nil
}

// bindSunpumplaunchpad binds a generic wrapper to an already deployed contract.
func bindSunpumplaunchpad(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SunpumplaunchpadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sunpumplaunchpad *SunpumplaunchpadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sunpumplaunchpad.Contract.SunpumplaunchpadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sunpumplaunchpad *SunpumplaunchpadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SunpumplaunchpadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sunpumplaunchpad *SunpumplaunchpadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SunpumplaunchpadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sunpumplaunchpad *SunpumplaunchpadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sunpumplaunchpad.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.contract.Transact(opts, method, params...)
}

// LAUNCHFEE is a paid mutator transaction binding the contract method 0x09197a81.
//
// Solidity: function LAUNCH_FEE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) LAUNCHFEE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "LAUNCH_FEE")
}

// LAUNCHFEE is a paid mutator transaction binding the contract method 0x09197a81.
//
// Solidity: function LAUNCH_FEE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) LAUNCHFEE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHFEE(&_Sunpumplaunchpad.TransactOpts)
}

// LAUNCHFEE is a paid mutator transaction binding the contract method 0x09197a81.
//
// Solidity: function LAUNCH_FEE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) LAUNCHFEE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHFEE(&_Sunpumplaunchpad.TransactOpts)
}

// LAUNCHTHRESHOLD is a paid mutator transaction binding the contract method 0x1d32c2da.
//
// Solidity: function LAUNCH_THRESHOLD() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) LAUNCHTHRESHOLD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "LAUNCH_THRESHOLD")
}

// LAUNCHTHRESHOLD is a paid mutator transaction binding the contract method 0x1d32c2da.
//
// Solidity: function LAUNCH_THRESHOLD() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) LAUNCHTHRESHOLD() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHTHRESHOLD(&_Sunpumplaunchpad.TransactOpts)
}

// LAUNCHTHRESHOLD is a paid mutator transaction binding the contract method 0x1d32c2da.
//
// Solidity: function LAUNCH_THRESHOLD() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) LAUNCHTHRESHOLD() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHTHRESHOLD(&_Sunpumplaunchpad.TransactOpts)
}

// LAUNCHTRXRESERVE is a paid mutator transaction binding the contract method 0x5da6454e.
//
// Solidity: function LAUNCH_TRX_RESERVE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) LAUNCHTRXRESERVE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "LAUNCH_TRX_RESERVE")
}

// LAUNCHTRXRESERVE is a paid mutator transaction binding the contract method 0x5da6454e.
//
// Solidity: function LAUNCH_TRX_RESERVE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) LAUNCHTRXRESERVE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHTRXRESERVE(&_Sunpumplaunchpad.TransactOpts)
}

// LAUNCHTRXRESERVE is a paid mutator transaction binding the contract method 0x5da6454e.
//
// Solidity: function LAUNCH_TRX_RESERVE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) LAUNCHTRXRESERVE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LAUNCHTRXRESERVE(&_Sunpumplaunchpad.TransactOpts)
}

// TOKENSUPPLY is a paid mutator transaction binding the contract method 0xb152f6cf.
//
// Solidity: function TOKEN_SUPPLY() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) TOKENSUPPLY(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "TOKEN_SUPPLY")
}

// TOKENSUPPLY is a paid mutator transaction binding the contract method 0xb152f6cf.
//
// Solidity: function TOKEN_SUPPLY() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) TOKENSUPPLY() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TOKENSUPPLY(&_Sunpumplaunchpad.TransactOpts)
}

// TOKENSUPPLY is a paid mutator transaction binding the contract method 0xb152f6cf.
//
// Solidity: function TOKEN_SUPPLY() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) TOKENSUPPLY() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TOKENSUPPLY(&_Sunpumplaunchpad.TransactOpts)
}

// TOTALSALE is a paid mutator transaction binding the contract method 0x910cffe0.
//
// Solidity: function TOTAL_SALE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) TOTALSALE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "TOTAL_SALE")
}

// TOTALSALE is a paid mutator transaction binding the contract method 0x910cffe0.
//
// Solidity: function TOTAL_SALE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) TOTALSALE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TOTALSALE(&_Sunpumplaunchpad.TransactOpts)
}

// TOTALSALE is a paid mutator transaction binding the contract method 0x910cffe0.
//
// Solidity: function TOTAL_SALE() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) TOTALSALE() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TOTALSALE(&_Sunpumplaunchpad.TransactOpts)
}

// VIRTUALTOKENRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x516cf387.
//
// Solidity: function VIRTUAL_TOKEN_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) VIRTUALTOKENRESERVEAMOUNT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "VIRTUAL_TOKEN_RESERVE_AMOUNT")
}

// VIRTUALTOKENRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x516cf387.
//
// Solidity: function VIRTUAL_TOKEN_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) VIRTUALTOKENRESERVEAMOUNT() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VIRTUALTOKENRESERVEAMOUNT(&_Sunpumplaunchpad.TransactOpts)
}

// VIRTUALTOKENRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x516cf387.
//
// Solidity: function VIRTUAL_TOKEN_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) VIRTUALTOKENRESERVEAMOUNT() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VIRTUALTOKENRESERVEAMOUNT(&_Sunpumplaunchpad.TransactOpts)
}

// VIRTUALTRXRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x3013355f.
//
// Solidity: function VIRTUAL_TRX_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) VIRTUALTRXRESERVEAMOUNT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "VIRTUAL_TRX_RESERVE_AMOUNT")
}

// VIRTUALTRXRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x3013355f.
//
// Solidity: function VIRTUAL_TRX_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) VIRTUALTRXRESERVEAMOUNT() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VIRTUALTRXRESERVEAMOUNT(&_Sunpumplaunchpad.TransactOpts)
}

// VIRTUALTRXRESERVEAMOUNT is a paid mutator transaction binding the contract method 0x3013355f.
//
// Solidity: function VIRTUAL_TRX_RESERVE_AMOUNT() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) VIRTUALTRXRESERVEAMOUNT() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VIRTUALTRXRESERVEAMOUNT(&_Sunpumplaunchpad.TransactOpts)
}

// BecomeNewImplementation is a paid mutator transaction binding the contract method 0xb7f7e90d.
//
// Solidity: function _becomeNewImplementation(address proxy) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) BecomeNewImplementation(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "_becomeNewImplementation", proxy)
}

// BecomeNewImplementation is a paid mutator transaction binding the contract method 0xb7f7e90d.
//
// Solidity: function _becomeNewImplementation(address proxy) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) BecomeNewImplementation(proxy common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.BecomeNewImplementation(&_Sunpumplaunchpad.TransactOpts, proxy)
}

// BecomeNewImplementation is a paid mutator transaction binding the contract method 0xb7f7e90d.
//
// Solidity: function _becomeNewImplementation(address proxy) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) BecomeNewImplementation(proxy common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.BecomeNewImplementation(&_Sunpumplaunchpad.TransactOpts, proxy)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) AcceptOwner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.AcceptOwner(&_Sunpumplaunchpad.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xebbc4965.
//
// Solidity: function acceptOwner() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.AcceptOwner(&_Sunpumplaunchpad.TransactOpts)
}

// CreateAndInitPurchase is a paid mutator transaction binding the contract method 0x2f70d762.
//
// Solidity: function createAndInitPurchase(string name, string symbol) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) CreateAndInitPurchase(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "createAndInitPurchase", name, symbol)
}

// CreateAndInitPurchase is a paid mutator transaction binding the contract method 0x2f70d762.
//
// Solidity: function createAndInitPurchase(string name, string symbol) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) CreateAndInitPurchase(name string, symbol string) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.CreateAndInitPurchase(&_Sunpumplaunchpad.TransactOpts, name, symbol)
}

// CreateAndInitPurchase is a paid mutator transaction binding the contract method 0x2f70d762.
//
// Solidity: function createAndInitPurchase(string name, string symbol) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) CreateAndInitPurchase(name string, symbol string) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.CreateAndInitPurchase(&_Sunpumplaunchpad.TransactOpts, name, symbol)
}

// DeadAddress is a paid mutator transaction binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) DeadAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "deadAddress")
}

// DeadAddress is a paid mutator transaction binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) DeadAddress() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.DeadAddress(&_Sunpumplaunchpad.TransactOpts)
}

// DeadAddress is a paid mutator transaction binding the contract method 0x27c8f835.
//
// Solidity: function deadAddress() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) DeadAddress() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.DeadAddress(&_Sunpumplaunchpad.TransactOpts)
}

// GetExactTokenAmountForPurchase is a paid mutator transaction binding the contract method 0x3ee09daa.
//
// Solidity: function getExactTokenAmountForPurchase(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetExactTokenAmountForPurchase(opts *bind.TransactOpts, token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getExactTokenAmountForPurchase", token, tokenAmount)
}

// GetExactTokenAmountForPurchase is a paid mutator transaction binding the contract method 0x3ee09daa.
//
// Solidity: function getExactTokenAmountForPurchase(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetExactTokenAmountForPurchase(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTokenAmountForPurchase(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetExactTokenAmountForPurchase is a paid mutator transaction binding the contract method 0x3ee09daa.
//
// Solidity: function getExactTokenAmountForPurchase(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetExactTokenAmountForPurchase(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTokenAmountForPurchase(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetExactTokenAmountForPurchaseWithFee is a paid mutator transaction binding the contract method 0xec1f743b.
//
// Solidity: function getExactTokenAmountForPurchaseWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetExactTokenAmountForPurchaseWithFee(opts *bind.TransactOpts, token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getExactTokenAmountForPurchaseWithFee", token, tokenAmount)
}

// GetExactTokenAmountForPurchaseWithFee is a paid mutator transaction binding the contract method 0xec1f743b.
//
// Solidity: function getExactTokenAmountForPurchaseWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetExactTokenAmountForPurchaseWithFee(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTokenAmountForPurchaseWithFee(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetExactTokenAmountForPurchaseWithFee is a paid mutator transaction binding the contract method 0xec1f743b.
//
// Solidity: function getExactTokenAmountForPurchaseWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetExactTokenAmountForPurchaseWithFee(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTokenAmountForPurchaseWithFee(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetExactTrxAmountForSale is a paid mutator transaction binding the contract method 0x2d324eac.
//
// Solidity: function getExactTrxAmountForSale(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetExactTrxAmountForSale(opts *bind.TransactOpts, token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getExactTrxAmountForSale", token, trxAmount)
}

// GetExactTrxAmountForSale is a paid mutator transaction binding the contract method 0x2d324eac.
//
// Solidity: function getExactTrxAmountForSale(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetExactTrxAmountForSale(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTrxAmountForSale(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetExactTrxAmountForSale is a paid mutator transaction binding the contract method 0x2d324eac.
//
// Solidity: function getExactTrxAmountForSale(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetExactTrxAmountForSale(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTrxAmountForSale(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetExactTrxAmountForSaleWithFee is a paid mutator transaction binding the contract method 0x44388b14.
//
// Solidity: function getExactTrxAmountForSaleWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetExactTrxAmountForSaleWithFee(opts *bind.TransactOpts, token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getExactTrxAmountForSaleWithFee", token, trxAmount)
}

// GetExactTrxAmountForSaleWithFee is a paid mutator transaction binding the contract method 0x44388b14.
//
// Solidity: function getExactTrxAmountForSaleWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetExactTrxAmountForSaleWithFee(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTrxAmountForSaleWithFee(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetExactTrxAmountForSaleWithFee is a paid mutator transaction binding the contract method 0x44388b14.
//
// Solidity: function getExactTrxAmountForSaleWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetExactTrxAmountForSaleWithFee(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetExactTrxAmountForSaleWithFee(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetPrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getPrice", token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetPrice(&_Sunpumplaunchpad.TransactOpts, token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetPrice(&_Sunpumplaunchpad.TransactOpts, token)
}

// GetTokenAmountByPurchase is a paid mutator transaction binding the contract method 0x1e105689.
//
// Solidity: function getTokenAmountByPurchase(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetTokenAmountByPurchase(opts *bind.TransactOpts, token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getTokenAmountByPurchase", token, trxAmount)
}

// GetTokenAmountByPurchase is a paid mutator transaction binding the contract method 0x1e105689.
//
// Solidity: function getTokenAmountByPurchase(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetTokenAmountByPurchase(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenAmountByPurchase(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetTokenAmountByPurchase is a paid mutator transaction binding the contract method 0x1e105689.
//
// Solidity: function getTokenAmountByPurchase(address token, uint256 trxAmount) View returns(uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetTokenAmountByPurchase(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenAmountByPurchase(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetTokenAmountByPurchaseWithFee is a paid mutator transaction binding the contract method 0x0bce3861.
//
// Solidity: function getTokenAmountByPurchaseWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetTokenAmountByPurchaseWithFee(opts *bind.TransactOpts, token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getTokenAmountByPurchaseWithFee", token, trxAmount)
}

// GetTokenAmountByPurchaseWithFee is a paid mutator transaction binding the contract method 0x0bce3861.
//
// Solidity: function getTokenAmountByPurchaseWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetTokenAmountByPurchaseWithFee(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenAmountByPurchaseWithFee(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetTokenAmountByPurchaseWithFee is a paid mutator transaction binding the contract method 0x0bce3861.
//
// Solidity: function getTokenAmountByPurchaseWithFee(address token, uint256 trxAmount) View returns(uint256 tokenAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetTokenAmountByPurchaseWithFee(token common.Address, trxAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenAmountByPurchaseWithFee(&_Sunpumplaunchpad.TransactOpts, token, trxAmount)
}

// GetTokenState is a paid mutator transaction binding the contract method 0x0b3eb970.
//
// Solidity: function getTokenState(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetTokenState(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getTokenState", token)
}

// GetTokenState is a paid mutator transaction binding the contract method 0x0b3eb970.
//
// Solidity: function getTokenState(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetTokenState(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenState(&_Sunpumplaunchpad.TransactOpts, token)
}

// GetTokenState is a paid mutator transaction binding the contract method 0x0b3eb970.
//
// Solidity: function getTokenState(address token) View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetTokenState(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTokenState(&_Sunpumplaunchpad.TransactOpts, token)
}

// GetTrxAmountBySale is a paid mutator transaction binding the contract method 0x2a5c0b6e.
//
// Solidity: function getTrxAmountBySale(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetTrxAmountBySale(opts *bind.TransactOpts, token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getTrxAmountBySale", token, tokenAmount)
}

// GetTrxAmountBySale is a paid mutator transaction binding the contract method 0x2a5c0b6e.
//
// Solidity: function getTrxAmountBySale(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetTrxAmountBySale(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTrxAmountBySale(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetTrxAmountBySale is a paid mutator transaction binding the contract method 0x2a5c0b6e.
//
// Solidity: function getTrxAmountBySale(address token, uint256 tokenAmount) View returns(uint256 trxAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetTrxAmountBySale(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTrxAmountBySale(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetTrxAmountBySaleWithFee is a paid mutator transaction binding the contract method 0xa85e75a8.
//
// Solidity: function getTrxAmountBySaleWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) GetTrxAmountBySaleWithFee(opts *bind.TransactOpts, token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "getTrxAmountBySaleWithFee", token, tokenAmount)
}

// GetTrxAmountBySaleWithFee is a paid mutator transaction binding the contract method 0xa85e75a8.
//
// Solidity: function getTrxAmountBySaleWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) GetTrxAmountBySaleWithFee(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTrxAmountBySaleWithFee(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// GetTrxAmountBySaleWithFee is a paid mutator transaction binding the contract method 0xa85e75a8.
//
// Solidity: function getTrxAmountBySaleWithFee(address token, uint256 tokenAmount) View returns(uint256 trxAmount, uint256 fee)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) GetTrxAmountBySaleWithFee(token common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.GetTrxAmountBySaleWithFee(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Implementation() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Implementation(&_Sunpumplaunchpad.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Implementation() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Implementation(&_Sunpumplaunchpad.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _vault, address _v2Router, uint256 _salefee, uint256 _purchasefee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Initialize(opts *bind.TransactOpts, _vault common.Address, _v2Router common.Address, _salefee *big.Int, _purchasefee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "initialize", _vault, _v2Router, _salefee, _purchasefee)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _vault, address _v2Router, uint256 _salefee, uint256 _purchasefee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Initialize(_vault common.Address, _v2Router common.Address, _salefee *big.Int, _purchasefee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Initialize(&_Sunpumplaunchpad.TransactOpts, _vault, _v2Router, _salefee, _purchasefee)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address _vault, address _v2Router, uint256 _salefee, uint256 _purchasefee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Initialize(_vault common.Address, _v2Router common.Address, _salefee *big.Int, _purchasefee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Initialize(&_Sunpumplaunchpad.TransactOpts, _vault, _v2Router, _salefee, _purchasefee)
}

// LaunchFee is a paid mutator transaction binding the contract method 0xcf3cf573.
//
// Solidity: function launchFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) LaunchFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "launchFee")
}

// LaunchFee is a paid mutator transaction binding the contract method 0xcf3cf573.
//
// Solidity: function launchFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) LaunchFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LaunchFee(&_Sunpumplaunchpad.TransactOpts)
}

// LaunchFee is a paid mutator transaction binding the contract method 0xcf3cf573.
//
// Solidity: function launchFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) LaunchFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LaunchFee(&_Sunpumplaunchpad.TransactOpts)
}

// LaunchToDEX is a paid mutator transaction binding the contract method 0x5a47db1d.
//
// Solidity: function launchToDEX(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) LaunchToDEX(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "launchToDEX", token)
}

// LaunchToDEX is a paid mutator transaction binding the contract method 0x5a47db1d.
//
// Solidity: function launchToDEX(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) LaunchToDEX(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LaunchToDEX(&_Sunpumplaunchpad.TransactOpts, token)
}

// LaunchToDEX is a paid mutator transaction binding the contract method 0x5a47db1d.
//
// Solidity: function launchToDEX(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) LaunchToDEX(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.LaunchToDEX(&_Sunpumplaunchpad.TransactOpts, token)
}

// Launcher is a paid mutator transaction binding the contract method 0x16eebd1e.
//
// Solidity: function launcher() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Launcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "launcher")
}

// Launcher is a paid mutator transaction binding the contract method 0x16eebd1e.
//
// Solidity: function launcher() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Launcher() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Launcher(&_Sunpumplaunchpad.TransactOpts)
}

// Launcher is a paid mutator transaction binding the contract method 0x16eebd1e.
//
// Solidity: function launcher() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Launcher() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Launcher(&_Sunpumplaunchpad.TransactOpts)
}

// MaxPurachaseAmount is a paid mutator transaction binding the contract method 0xceee6c1b.
//
// Solidity: function maxPurachaseAmount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) MaxPurachaseAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "maxPurachaseAmount")
}

// MaxPurachaseAmount is a paid mutator transaction binding the contract method 0xceee6c1b.
//
// Solidity: function maxPurachaseAmount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) MaxPurachaseAmount() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MaxPurachaseAmount(&_Sunpumplaunchpad.TransactOpts)
}

// MaxPurachaseAmount is a paid mutator transaction binding the contract method 0xceee6c1b.
//
// Solidity: function maxPurachaseAmount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) MaxPurachaseAmount() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MaxPurachaseAmount(&_Sunpumplaunchpad.TransactOpts)
}

// MinTxFee is a paid mutator transaction binding the contract method 0xfe29b4e8.
//
// Solidity: function minTxFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) MinTxFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "minTxFee")
}

// MinTxFee is a paid mutator transaction binding the contract method 0xfe29b4e8.
//
// Solidity: function minTxFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) MinTxFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MinTxFee(&_Sunpumplaunchpad.TransactOpts)
}

// MinTxFee is a paid mutator transaction binding the contract method 0xfe29b4e8.
//
// Solidity: function minTxFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) MinTxFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MinTxFee(&_Sunpumplaunchpad.TransactOpts)
}

// MintFee is a paid mutator transaction binding the contract method 0x13966db5.
//
// Solidity: function mintFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) MintFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "mintFee")
}

// MintFee is a paid mutator transaction binding the contract method 0x13966db5.
//
// Solidity: function mintFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) MintFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MintFee(&_Sunpumplaunchpad.TransactOpts)
}

// MintFee is a paid mutator transaction binding the contract method 0x13966db5.
//
// Solidity: function mintFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) MintFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.MintFee(&_Sunpumplaunchpad.TransactOpts)
}

// Operator is a paid mutator transaction binding the contract method 0x570ca735.
//
// Solidity: function operator() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Operator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "operator")
}

// Operator is a paid mutator transaction binding the contract method 0x570ca735.
//
// Solidity: function operator() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Operator() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Operator(&_Sunpumplaunchpad.TransactOpts)
}

// Operator is a paid mutator transaction binding the contract method 0x570ca735.
//
// Solidity: function operator() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Operator() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Operator(&_Sunpumplaunchpad.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Owner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Owner(&_Sunpumplaunchpad.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Owner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Owner(&_Sunpumplaunchpad.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() View returns(bool)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() View returns(bool)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Pause() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Pause(&_Sunpumplaunchpad.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() View returns(bool)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Pause() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Pause(&_Sunpumplaunchpad.TransactOpts)
}

// PausePad is a paid mutator transaction binding the contract method 0x32adb109.
//
// Solidity: function pausePad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) PausePad(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "pausePad")
}

// PausePad is a paid mutator transaction binding the contract method 0x32adb109.
//
// Solidity: function pausePad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) PausePad() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PausePad(&_Sunpumplaunchpad.TransactOpts)
}

// PausePad is a paid mutator transaction binding the contract method 0x32adb109.
//
// Solidity: function pausePad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) PausePad() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PausePad(&_Sunpumplaunchpad.TransactOpts)
}

// PendingImplementation is a paid mutator transaction binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) PendingImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "pendingImplementation")
}

// PendingImplementation is a paid mutator transaction binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) PendingImplementation() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PendingImplementation(&_Sunpumplaunchpad.TransactOpts)
}

// PendingImplementation is a paid mutator transaction binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) PendingImplementation() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PendingImplementation(&_Sunpumplaunchpad.TransactOpts)
}

// PendingOwner is a paid mutator transaction binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) PendingOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "pendingOwner")
}

// PendingOwner is a paid mutator transaction binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) PendingOwner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PendingOwner(&_Sunpumplaunchpad.TransactOpts)
}

// PendingOwner is a paid mutator transaction binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) PendingOwner() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PendingOwner(&_Sunpumplaunchpad.TransactOpts)
}

// PurchaseFee is a paid mutator transaction binding the contract method 0x14b5e981.
//
// Solidity: function purchaseFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) PurchaseFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "purchaseFee")
}

// PurchaseFee is a paid mutator transaction binding the contract method 0x14b5e981.
//
// Solidity: function purchaseFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) PurchaseFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PurchaseFee(&_Sunpumplaunchpad.TransactOpts)
}

// PurchaseFee is a paid mutator transaction binding the contract method 0x14b5e981.
//
// Solidity: function purchaseFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) PurchaseFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PurchaseFee(&_Sunpumplaunchpad.TransactOpts)
}

// PurchaseToken is a paid mutator transaction binding the contract method 0x1cc2c911.
//
// Solidity: function purchaseToken(address token, uint256 AmountMin) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) PurchaseToken(opts *bind.TransactOpts, token common.Address, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "purchaseToken", token, AmountMin)
}

// PurchaseToken is a paid mutator transaction binding the contract method 0x1cc2c911.
//
// Solidity: function purchaseToken(address token, uint256 AmountMin) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) PurchaseToken(token common.Address, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PurchaseToken(&_Sunpumplaunchpad.TransactOpts, token, AmountMin)
}

// PurchaseToken is a paid mutator transaction binding the contract method 0x1cc2c911.
//
// Solidity: function purchaseToken(address token, uint256 AmountMin) payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) PurchaseToken(token common.Address, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.PurchaseToken(&_Sunpumplaunchpad.TransactOpts, token, AmountMin)
}

// RenounceTokenOwnership is a paid mutator transaction binding the contract method 0x03f76476.
//
// Solidity: function renounceTokenOwnership(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) RenounceTokenOwnership(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "renounceTokenOwnership", token)
}

// RenounceTokenOwnership is a paid mutator transaction binding the contract method 0x03f76476.
//
// Solidity: function renounceTokenOwnership(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) RenounceTokenOwnership(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.RenounceTokenOwnership(&_Sunpumplaunchpad.TransactOpts, token)
}

// RenounceTokenOwnership is a paid mutator transaction binding the contract method 0x03f76476.
//
// Solidity: function renounceTokenOwnership(address token) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) RenounceTokenOwnership(token common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.RenounceTokenOwnership(&_Sunpumplaunchpad.TransactOpts, token)
}

// RerunPad is a paid mutator transaction binding the contract method 0xd92748ea.
//
// Solidity: function rerunPad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) RerunPad(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "rerunPad")
}

// RerunPad is a paid mutator transaction binding the contract method 0xd92748ea.
//
// Solidity: function rerunPad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) RerunPad() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.RerunPad(&_Sunpumplaunchpad.TransactOpts)
}

// RerunPad is a paid mutator transaction binding the contract method 0xd92748ea.
//
// Solidity: function rerunPad() Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) RerunPad() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.RerunPad(&_Sunpumplaunchpad.TransactOpts)
}

// SaleFee is a paid mutator transaction binding the contract method 0x178021e3.
//
// Solidity: function saleFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SaleFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "saleFee")
}

// SaleFee is a paid mutator transaction binding the contract method 0x178021e3.
//
// Solidity: function saleFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SaleFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SaleFee(&_Sunpumplaunchpad.TransactOpts)
}

// SaleFee is a paid mutator transaction binding the contract method 0x178021e3.
//
// Solidity: function saleFee() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SaleFee() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SaleFee(&_Sunpumplaunchpad.TransactOpts)
}

// SaleToken is a paid mutator transaction binding the contract method 0xd19aa2b9.
//
// Solidity: function saleToken(address token, uint256 tokenAmount, uint256 AmountMin) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SaleToken(opts *bind.TransactOpts, token common.Address, tokenAmount *big.Int, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "saleToken", token, tokenAmount, AmountMin)
}

// SaleToken is a paid mutator transaction binding the contract method 0xd19aa2b9.
//
// Solidity: function saleToken(address token, uint256 tokenAmount, uint256 AmountMin) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SaleToken(token common.Address, tokenAmount *big.Int, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SaleToken(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount, AmountMin)
}

// SaleToken is a paid mutator transaction binding the contract method 0xd19aa2b9.
//
// Solidity: function saleToken(address token, uint256 tokenAmount, uint256 AmountMin) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SaleToken(token common.Address, tokenAmount *big.Int, AmountMin *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SaleToken(&_Sunpumplaunchpad.TransactOpts, token, tokenAmount, AmountMin)
}

// SetLauncher is a paid mutator transaction binding the contract method 0xf4c094c8.
//
// Solidity: function setLauncher(address newLauncher) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetLauncher(opts *bind.TransactOpts, newLauncher common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setLauncher", newLauncher)
}

// SetLauncher is a paid mutator transaction binding the contract method 0xf4c094c8.
//
// Solidity: function setLauncher(address newLauncher) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetLauncher(newLauncher common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetLauncher(&_Sunpumplaunchpad.TransactOpts, newLauncher)
}

// SetLauncher is a paid mutator transaction binding the contract method 0xf4c094c8.
//
// Solidity: function setLauncher(address newLauncher) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetLauncher(newLauncher common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetLauncher(&_Sunpumplaunchpad.TransactOpts, newLauncher)
}

// SetMinTxFee is a paid mutator transaction binding the contract method 0x20160b07.
//
// Solidity: function setMinTxFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetMinTxFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setMinTxFee", newFee)
}

// SetMinTxFee is a paid mutator transaction binding the contract method 0x20160b07.
//
// Solidity: function setMinTxFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetMinTxFee(newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMinTxFee(&_Sunpumplaunchpad.TransactOpts, newFee)
}

// SetMinTxFee is a paid mutator transaction binding the contract method 0x20160b07.
//
// Solidity: function setMinTxFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetMinTxFee(newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMinTxFee(&_Sunpumplaunchpad.TransactOpts, newFee)
}

// SetMintAndMinTxFee is a paid mutator transaction binding the contract method 0x12e6c2c3.
//
// Solidity: function setMintAndMinTxFee(uint256 _newMintFee, uint256 _newMinTxFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetMintAndMinTxFee(opts *bind.TransactOpts, _newMintFee *big.Int, _newMinTxFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setMintAndMinTxFee", _newMintFee, _newMinTxFee)
}

// SetMintAndMinTxFee is a paid mutator transaction binding the contract method 0x12e6c2c3.
//
// Solidity: function setMintAndMinTxFee(uint256 _newMintFee, uint256 _newMinTxFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetMintAndMinTxFee(_newMintFee *big.Int, _newMinTxFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMintAndMinTxFee(&_Sunpumplaunchpad.TransactOpts, _newMintFee, _newMinTxFee)
}

// SetMintAndMinTxFee is a paid mutator transaction binding the contract method 0x12e6c2c3.
//
// Solidity: function setMintAndMinTxFee(uint256 _newMintFee, uint256 _newMinTxFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetMintAndMinTxFee(_newMintFee *big.Int, _newMinTxFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMintAndMinTxFee(&_Sunpumplaunchpad.TransactOpts, _newMintFee, _newMinTxFee)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetMintFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setMintFee", newFee)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetMintFee(newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMintFee(&_Sunpumplaunchpad.TransactOpts, newFee)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 newFee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetMintFee(newFee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetMintFee(&_Sunpumplaunchpad.TransactOpts, newFee)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOp) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetOperator(opts *bind.TransactOpts, newOp common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setOperator", newOp)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOp) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetOperator(newOp common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetOperator(&_Sunpumplaunchpad.TransactOpts, newOp)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOp) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetOperator(newOp common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetOperator(&_Sunpumplaunchpad.TransactOpts, newOp)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetPendingOwner(&_Sunpumplaunchpad.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetPendingOwner(&_Sunpumplaunchpad.TransactOpts, newPendingOwner)
}

// SetPurchaseFee is a paid mutator transaction binding the contract method 0x6402cdc3.
//
// Solidity: function setPurchaseFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetPurchaseFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setPurchaseFee", _fee)
}

// SetPurchaseFee is a paid mutator transaction binding the contract method 0x6402cdc3.
//
// Solidity: function setPurchaseFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetPurchaseFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetPurchaseFee(&_Sunpumplaunchpad.TransactOpts, _fee)
}

// SetPurchaseFee is a paid mutator transaction binding the contract method 0x6402cdc3.
//
// Solidity: function setPurchaseFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetPurchaseFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetPurchaseFee(&_Sunpumplaunchpad.TransactOpts, _fee)
}

// SetSaleFee is a paid mutator transaction binding the contract method 0xbdcafc55.
//
// Solidity: function setSaleFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetSaleFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setSaleFee", _fee)
}

// SetSaleFee is a paid mutator transaction binding the contract method 0xbdcafc55.
//
// Solidity: function setSaleFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetSaleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetSaleFee(&_Sunpumplaunchpad.TransactOpts, _fee)
}

// SetSaleFee is a paid mutator transaction binding the contract method 0xbdcafc55.
//
// Solidity: function setSaleFee(uint256 _fee) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetSaleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetSaleFee(&_Sunpumplaunchpad.TransactOpts, _fee)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _addr) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) SetVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "setVault", _addr)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _addr) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) SetVault(_addr common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetVault(&_Sunpumplaunchpad.TransactOpts, _addr)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _addr) Nonpayable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) SetVault(_addr common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.SetVault(&_Sunpumplaunchpad.TransactOpts, _addr)
}

// TokenAddress is a paid mutator transaction binding the contract method 0x9e6b26ba.
//
// Solidity: function tokenAddress(uint256 ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) TokenAddress(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "tokenAddress", arg0)
}

// TokenAddress is a paid mutator transaction binding the contract method 0x9e6b26ba.
//
// Solidity: function tokenAddress(uint256 ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) TokenAddress(arg0 *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenAddress(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// TokenAddress is a paid mutator transaction binding the contract method 0x9e6b26ba.
//
// Solidity: function tokenAddress(uint256 ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) TokenAddress(arg0 *big.Int) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenAddress(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// TokenCount is a paid mutator transaction binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) TokenCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "tokenCount")
}

// TokenCount is a paid mutator transaction binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) TokenCount() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenCount(&_Sunpumplaunchpad.TransactOpts)
}

// TokenCount is a paid mutator transaction binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() View returns(uint256)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) TokenCount() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenCount(&_Sunpumplaunchpad.TransactOpts)
}

// TokenCreator is a paid mutator transaction binding the contract method 0x23774af2.
//
// Solidity: function tokenCreator(address ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) TokenCreator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "tokenCreator", arg0)
}

// TokenCreator is a paid mutator transaction binding the contract method 0x23774af2.
//
// Solidity: function tokenCreator(address ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) TokenCreator(arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenCreator(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// TokenCreator is a paid mutator transaction binding the contract method 0x23774af2.
//
// Solidity: function tokenCreator(address ) View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) TokenCreator(arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.TokenCreator(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// V2Router is a paid mutator transaction binding the contract method 0xdeadbc14.
//
// Solidity: function v2Router() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) V2Router(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "v2Router")
}

// V2Router is a paid mutator transaction binding the contract method 0xdeadbc14.
//
// Solidity: function v2Router() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) V2Router() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.V2Router(&_Sunpumplaunchpad.TransactOpts)
}

// V2Router is a paid mutator transaction binding the contract method 0xdeadbc14.
//
// Solidity: function v2Router() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) V2Router() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.V2Router(&_Sunpumplaunchpad.TransactOpts)
}

// Vault is a paid mutator transaction binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Vault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "vault")
}

// Vault is a paid mutator transaction binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Vault() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Vault(&_Sunpumplaunchpad.TransactOpts)
}

// Vault is a paid mutator transaction binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() View returns(address)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Vault() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Vault(&_Sunpumplaunchpad.TransactOpts)
}

// VirtualPools is a paid mutator transaction binding the contract method 0x1e228192.
//
// Solidity: function virtualPools(address ) View returns(uint256 TRXReserve, uint256 TokenReserve, bool launched)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) VirtualPools(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.Transact(opts, "virtualPools", arg0)
}

// VirtualPools is a paid mutator transaction binding the contract method 0x1e228192.
//
// Solidity: function virtualPools(address ) View returns(uint256 TRXReserve, uint256 TokenReserve, bool launched)
func (_Sunpumplaunchpad *SunpumplaunchpadSession) VirtualPools(arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VirtualPools(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// VirtualPools is a paid mutator transaction binding the contract method 0x1e228192.
//
// Solidity: function virtualPools(address ) View returns(uint256 TRXReserve, uint256 TokenReserve, bool launched)
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) VirtualPools(arg0 common.Address) (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.VirtualPools(&_Sunpumplaunchpad.TransactOpts, arg0)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sunpumplaunchpad.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadSession) Receive() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Receive(&_Sunpumplaunchpad.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sunpumplaunchpad *SunpumplaunchpadTransactorSession) Receive() (*types.Transaction, error) {
	return _Sunpumplaunchpad.Contract.Receive(&_Sunpumplaunchpad.TransactOpts)
}

// SunpumplaunchpadLaunchPendingIterator is returned from FilterLaunchPending and is used to iterate over the raw logs and unpacked data for LaunchPending events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadLaunchPendingIterator struct {
	Event *SunpumplaunchpadLaunchPending // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadLaunchPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadLaunchPending)
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
		it.Event = new(SunpumplaunchpadLaunchPending)
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
func (it *SunpumplaunchpadLaunchPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadLaunchPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadLaunchPending represents a LaunchPending event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadLaunchPending struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLaunchPending is a free log retrieval operation binding the contract event 0xff274cd97aba8af276149429fbc7ea387e14da22dcd51779c691af908f4feb64.
//
// Solidity: event LaunchPending(address token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterLaunchPending(opts *bind.FilterOpts) (*SunpumplaunchpadLaunchPendingIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "LaunchPending")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadLaunchPendingIterator{contract: _Sunpumplaunchpad.contract, event: "LaunchPending", logs: logs, sub: sub}, nil
}

// WatchLaunchPending is a free log subscription operation binding the contract event 0xff274cd97aba8af276149429fbc7ea387e14da22dcd51779c691af908f4feb64.
//
// Solidity: event LaunchPending(address token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchLaunchPending(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadLaunchPending) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "LaunchPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadLaunchPending)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "LaunchPending", log); err != nil {
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

// ParseLaunchPending is a log parse operation binding the contract event 0xff274cd97aba8af276149429fbc7ea387e14da22dcd51779c691af908f4feb64.
//
// Solidity: event LaunchPending(address token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseLaunchPending(log types.Log) (*SunpumplaunchpadLaunchPending, error) {
	event := new(SunpumplaunchpadLaunchPending)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "LaunchPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadLauncherChangedIterator is returned from FilterLauncherChanged and is used to iterate over the raw logs and unpacked data for LauncherChanged events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadLauncherChangedIterator struct {
	Event *SunpumplaunchpadLauncherChanged // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadLauncherChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadLauncherChanged)
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
		it.Event = new(SunpumplaunchpadLauncherChanged)
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
func (it *SunpumplaunchpadLauncherChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadLauncherChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadLauncherChanged represents a LauncherChanged event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadLauncherChanged struct {
	OldLauncher common.Address
	NewLauncher common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLauncherChanged is a free log retrieval operation binding the contract event 0x349b3ca858f1d049aa7b4e826494f79354dbe5f1125ec5442e2d309d21646ec2.
//
// Solidity: event LauncherChanged(address indexed oldLauncher, address indexed newLauncher)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterLauncherChanged(opts *bind.FilterOpts, oldLauncher []common.Address, newLauncher []common.Address) (*SunpumplaunchpadLauncherChangedIterator, error) {

	var oldLauncherRule []interface{}
	for _, oldLauncherItem := range oldLauncher {
		oldLauncherRule = append(oldLauncherRule, oldLauncherItem)
	}
	var newLauncherRule []interface{}
	for _, newLauncherItem := range newLauncher {
		newLauncherRule = append(newLauncherRule, newLauncherItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "LauncherChanged", oldLauncherRule, newLauncherRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadLauncherChangedIterator{contract: _Sunpumplaunchpad.contract, event: "LauncherChanged", logs: logs, sub: sub}, nil
}

// WatchLauncherChanged is a free log subscription operation binding the contract event 0x349b3ca858f1d049aa7b4e826494f79354dbe5f1125ec5442e2d309d21646ec2.
//
// Solidity: event LauncherChanged(address indexed oldLauncher, address indexed newLauncher)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchLauncherChanged(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadLauncherChanged, oldLauncher []common.Address, newLauncher []common.Address) (event.Subscription, error) {

	var oldLauncherRule []interface{}
	for _, oldLauncherItem := range oldLauncher {
		oldLauncherRule = append(oldLauncherRule, oldLauncherItem)
	}
	var newLauncherRule []interface{}
	for _, newLauncherItem := range newLauncher {
		newLauncherRule = append(newLauncherRule, newLauncherItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "LauncherChanged", oldLauncherRule, newLauncherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadLauncherChanged)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "LauncherChanged", log); err != nil {
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

// ParseLauncherChanged is a log parse operation binding the contract event 0x349b3ca858f1d049aa7b4e826494f79354dbe5f1125ec5442e2d309d21646ec2.
//
// Solidity: event LauncherChanged(address indexed oldLauncher, address indexed newLauncher)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseLauncherChanged(log types.Log) (*SunpumplaunchpadLauncherChanged, error) {
	event := new(SunpumplaunchpadLauncherChanged)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "LauncherChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadMinTxFeeSetIterator is returned from FilterMinTxFeeSet and is used to iterate over the raw logs and unpacked data for MinTxFeeSet events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadMinTxFeeSetIterator struct {
	Event *SunpumplaunchpadMinTxFeeSet // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadMinTxFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadMinTxFeeSet)
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
		it.Event = new(SunpumplaunchpadMinTxFeeSet)
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
func (it *SunpumplaunchpadMinTxFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadMinTxFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadMinTxFeeSet represents a MinTxFeeSet event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadMinTxFeeSet struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinTxFeeSet is a free log retrieval operation binding the contract event 0x091f43688ad4f42d7f02fb81c7e28b693ca4d5d825c27572437bc308cc46aba2.
//
// Solidity: event MinTxFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterMinTxFeeSet(opts *bind.FilterOpts) (*SunpumplaunchpadMinTxFeeSetIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "MinTxFeeSet")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadMinTxFeeSetIterator{contract: _Sunpumplaunchpad.contract, event: "MinTxFeeSet", logs: logs, sub: sub}, nil
}

// WatchMinTxFeeSet is a free log subscription operation binding the contract event 0x091f43688ad4f42d7f02fb81c7e28b693ca4d5d825c27572437bc308cc46aba2.
//
// Solidity: event MinTxFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchMinTxFeeSet(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadMinTxFeeSet) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "MinTxFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadMinTxFeeSet)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "MinTxFeeSet", log); err != nil {
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

// ParseMinTxFeeSet is a log parse operation binding the contract event 0x091f43688ad4f42d7f02fb81c7e28b693ca4d5d825c27572437bc308cc46aba2.
//
// Solidity: event MinTxFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseMinTxFeeSet(log types.Log) (*SunpumplaunchpadMinTxFeeSet, error) {
	event := new(SunpumplaunchpadMinTxFeeSet)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "MinTxFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadMintFeeSetIterator is returned from FilterMintFeeSet and is used to iterate over the raw logs and unpacked data for MintFeeSet events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadMintFeeSetIterator struct {
	Event *SunpumplaunchpadMintFeeSet // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadMintFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadMintFeeSet)
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
		it.Event = new(SunpumplaunchpadMintFeeSet)
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
func (it *SunpumplaunchpadMintFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadMintFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadMintFeeSet represents a MintFeeSet event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadMintFeeSet struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintFeeSet is a free log retrieval operation binding the contract event 0x387269377ae17304805d5f88cea4252e5ca47346783c279aeb9e8627335a49ac.
//
// Solidity: event MintFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterMintFeeSet(opts *bind.FilterOpts) (*SunpumplaunchpadMintFeeSetIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "MintFeeSet")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadMintFeeSetIterator{contract: _Sunpumplaunchpad.contract, event: "MintFeeSet", logs: logs, sub: sub}, nil
}

// WatchMintFeeSet is a free log subscription operation binding the contract event 0x387269377ae17304805d5f88cea4252e5ca47346783c279aeb9e8627335a49ac.
//
// Solidity: event MintFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchMintFeeSet(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadMintFeeSet) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "MintFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadMintFeeSet)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "MintFeeSet", log); err != nil {
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

// ParseMintFeeSet is a log parse operation binding the contract event 0x387269377ae17304805d5f88cea4252e5ca47346783c279aeb9e8627335a49ac.
//
// Solidity: event MintFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseMintFeeSet(log types.Log) (*SunpumplaunchpadMintFeeSet, error) {
	event := new(SunpumplaunchpadMintFeeSet)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "MintFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadOperatorChangedIterator struct {
	Event *SunpumplaunchpadOperatorChanged // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadOperatorChanged)
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
		it.Event = new(SunpumplaunchpadOperatorChanged)
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
func (it *SunpumplaunchpadOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadOperatorChanged represents a OperatorChanged event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadOperatorChanged struct {
	OldOperator common.Address
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterOperatorChanged(opts *bind.FilterOpts, oldOperator []common.Address, newOperator []common.Address) (*SunpumplaunchpadOperatorChangedIterator, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "OperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadOperatorChangedIterator{contract: _Sunpumplaunchpad.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadOperatorChanged, oldOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "OperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadOperatorChanged)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseOperatorChanged(log types.Log) (*SunpumplaunchpadOperatorChanged, error) {
	event := new(SunpumplaunchpadOperatorChanged)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadOwnerChangedIterator struct {
	Event *SunpumplaunchpadOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadOwnerChanged)
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
		it.Event = new(SunpumplaunchpadOwnerChanged)
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
func (it *SunpumplaunchpadOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadOwnerChanged represents a OwnerChanged event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SunpumplaunchpadOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadOwnerChangedIterator{contract: _Sunpumplaunchpad.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadOwnerChanged)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseOwnerChanged(log types.Log) (*SunpumplaunchpadOwnerChanged, error) {
	event := new(SunpumplaunchpadOwnerChanged)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadPendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadPendingOwnerSetIterator struct {
	Event *SunpumplaunchpadPendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadPendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadPendingOwnerSet)
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
		it.Event = new(SunpumplaunchpadPendingOwnerSet)
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
func (it *SunpumplaunchpadPendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadPendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadPendingOwnerSet represents a PendingOwnerSet event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadPendingOwnerSet struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*SunpumplaunchpadPendingOwnerSetIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "PendingOwnerSet", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadPendingOwnerSetIterator{contract: _Sunpumplaunchpad.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadPendingOwnerSet, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "PendingOwnerSet", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadPendingOwnerSet)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParsePendingOwnerSet(log types.Log) (*SunpumplaunchpadPendingOwnerSet, error) {
	event := new(SunpumplaunchpadPendingOwnerSet)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadPurchaseFeeSetIterator is returned from FilterPurchaseFeeSet and is used to iterate over the raw logs and unpacked data for PurchaseFeeSet events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadPurchaseFeeSetIterator struct {
	Event *SunpumplaunchpadPurchaseFeeSet // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadPurchaseFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadPurchaseFeeSet)
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
		it.Event = new(SunpumplaunchpadPurchaseFeeSet)
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
func (it *SunpumplaunchpadPurchaseFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadPurchaseFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadPurchaseFeeSet represents a PurchaseFeeSet event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadPurchaseFeeSet struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPurchaseFeeSet is a free log retrieval operation binding the contract event 0x525ad74c8a8eb66a2372d62022c6d2813136ef0f41574506841342146cf694db.
//
// Solidity: event PurchaseFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterPurchaseFeeSet(opts *bind.FilterOpts) (*SunpumplaunchpadPurchaseFeeSetIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "PurchaseFeeSet")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadPurchaseFeeSetIterator{contract: _Sunpumplaunchpad.contract, event: "PurchaseFeeSet", logs: logs, sub: sub}, nil
}

// WatchPurchaseFeeSet is a free log subscription operation binding the contract event 0x525ad74c8a8eb66a2372d62022c6d2813136ef0f41574506841342146cf694db.
//
// Solidity: event PurchaseFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchPurchaseFeeSet(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadPurchaseFeeSet) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "PurchaseFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadPurchaseFeeSet)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "PurchaseFeeSet", log); err != nil {
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

// ParsePurchaseFeeSet is a log parse operation binding the contract event 0x525ad74c8a8eb66a2372d62022c6d2813136ef0f41574506841342146cf694db.
//
// Solidity: event PurchaseFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParsePurchaseFeeSet(log types.Log) (*SunpumplaunchpadPurchaseFeeSet, error) {
	event := new(SunpumplaunchpadPurchaseFeeSet)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "PurchaseFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadSaleFeeSetIterator is returned from FilterSaleFeeSet and is used to iterate over the raw logs and unpacked data for SaleFeeSet events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadSaleFeeSetIterator struct {
	Event *SunpumplaunchpadSaleFeeSet // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadSaleFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadSaleFeeSet)
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
		it.Event = new(SunpumplaunchpadSaleFeeSet)
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
func (it *SunpumplaunchpadSaleFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadSaleFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadSaleFeeSet represents a SaleFeeSet event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadSaleFeeSet struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSaleFeeSet is a free log retrieval operation binding the contract event 0x3063d3516a6a2d04f3fcbbb8096b055b0976989088fe108ad8081c526a594abd.
//
// Solidity: event SaleFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterSaleFeeSet(opts *bind.FilterOpts) (*SunpumplaunchpadSaleFeeSetIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "SaleFeeSet")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadSaleFeeSetIterator{contract: _Sunpumplaunchpad.contract, event: "SaleFeeSet", logs: logs, sub: sub}, nil
}

// WatchSaleFeeSet is a free log subscription operation binding the contract event 0x3063d3516a6a2d04f3fcbbb8096b055b0976989088fe108ad8081c526a594abd.
//
// Solidity: event SaleFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchSaleFeeSet(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadSaleFeeSet) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "SaleFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadSaleFeeSet)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "SaleFeeSet", log); err != nil {
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

// ParseSaleFeeSet is a log parse operation binding the contract event 0x3063d3516a6a2d04f3fcbbb8096b055b0976989088fe108ad8081c526a594abd.
//
// Solidity: event SaleFeeSet(uint256 oldFee, uint256 newFee)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseSaleFeeSet(log types.Log) (*SunpumplaunchpadSaleFeeSet, error) {
	event := new(SunpumplaunchpadSaleFeeSet)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "SaleFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadTokenCreateIterator is returned from FilterTokenCreate and is used to iterate over the raw logs and unpacked data for TokenCreate events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenCreateIterator struct {
	Event *SunpumplaunchpadTokenCreate // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadTokenCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadTokenCreate)
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
		it.Event = new(SunpumplaunchpadTokenCreate)
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
func (it *SunpumplaunchpadTokenCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadTokenCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadTokenCreate represents a TokenCreate event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenCreate struct {
	TokenAddress common.Address
	TokenIndex   *big.Int
	Creator      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenCreate is a free log retrieval operation binding the contract event 0x1ff0a01c8968e3551472812164f233abb579247de887db8cbb18281c149bee7a.
//
// Solidity: event TokenCreate(address tokenAddress, uint256 tokenIndex, address creator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterTokenCreate(opts *bind.FilterOpts) (*SunpumplaunchpadTokenCreateIterator, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "TokenCreate")
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadTokenCreateIterator{contract: _Sunpumplaunchpad.contract, event: "TokenCreate", logs: logs, sub: sub}, nil
}

// WatchTokenCreate is a free log subscription operation binding the contract event 0x1ff0a01c8968e3551472812164f233abb579247de887db8cbb18281c149bee7a.
//
// Solidity: event TokenCreate(address tokenAddress, uint256 tokenIndex, address creator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchTokenCreate(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadTokenCreate) (event.Subscription, error) {

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "TokenCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadTokenCreate)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenCreate", log); err != nil {
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

// ParseTokenCreate is a log parse operation binding the contract event 0x1ff0a01c8968e3551472812164f233abb579247de887db8cbb18281c149bee7a.
//
// Solidity: event TokenCreate(address tokenAddress, uint256 tokenIndex, address creator)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseTokenCreate(log types.Log) (*SunpumplaunchpadTokenCreate, error) {
	event := new(SunpumplaunchpadTokenCreate)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadTokenLaunchedIterator is returned from FilterTokenLaunched and is used to iterate over the raw logs and unpacked data for TokenLaunched events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenLaunchedIterator struct {
	Event *SunpumplaunchpadTokenLaunched // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadTokenLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadTokenLaunched)
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
		it.Event = new(SunpumplaunchpadTokenLaunched)
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
func (it *SunpumplaunchpadTokenLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadTokenLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadTokenLaunched represents a TokenLaunched event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenLaunched struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenLaunched is a free log retrieval operation binding the contract event 0x2ab676eef3f76f1bd4e765a352c6cd81e62702f7ad3d363291c8b60582a45250.
//
// Solidity: event TokenLaunched(address indexed token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterTokenLaunched(opts *bind.FilterOpts, token []common.Address) (*SunpumplaunchpadTokenLaunchedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "TokenLaunched", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadTokenLaunchedIterator{contract: _Sunpumplaunchpad.contract, event: "TokenLaunched", logs: logs, sub: sub}, nil
}

// WatchTokenLaunched is a free log subscription operation binding the contract event 0x2ab676eef3f76f1bd4e765a352c6cd81e62702f7ad3d363291c8b60582a45250.
//
// Solidity: event TokenLaunched(address indexed token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchTokenLaunched(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadTokenLaunched, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "TokenLaunched", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadTokenLaunched)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenLaunched", log); err != nil {
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

// ParseTokenLaunched is a log parse operation binding the contract event 0x2ab676eef3f76f1bd4e765a352c6cd81e62702f7ad3d363291c8b60582a45250.
//
// Solidity: event TokenLaunched(address indexed token)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseTokenLaunched(log types.Log) (*SunpumplaunchpadTokenLaunched, error) {
	event := new(SunpumplaunchpadTokenLaunched)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenLaunched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadTokenPurchasedIterator is returned from FilterTokenPurchased and is used to iterate over the raw logs and unpacked data for TokenPurchased events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenPurchasedIterator struct {
	Event *SunpumplaunchpadTokenPurchased // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadTokenPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadTokenPurchased)
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
		it.Event = new(SunpumplaunchpadTokenPurchased)
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
func (it *SunpumplaunchpadTokenPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadTokenPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadTokenPurchased represents a TokenPurchased event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenPurchased struct {
	Token        common.Address
	Buyer        common.Address
	TrxAmount    *big.Int
	Fee          *big.Int
	TokenAmount  *big.Int
	TokenReserve *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchased is a free log retrieval operation binding the contract event 0x63abb62535c21a5d221cf9c15994097b8880cc986d82faf80f57382b998dbae5.
//
// Solidity: event TokenPurchased(address indexed token, address indexed buyer, uint256 trxAmount, uint256 fee, uint256 tokenAmount, uint256 tokenReserve)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterTokenPurchased(opts *bind.FilterOpts, token []common.Address, buyer []common.Address) (*SunpumplaunchpadTokenPurchasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "TokenPurchased", tokenRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadTokenPurchasedIterator{contract: _Sunpumplaunchpad.contract, event: "TokenPurchased", logs: logs, sub: sub}, nil
}

// WatchTokenPurchased is a free log subscription operation binding the contract event 0x63abb62535c21a5d221cf9c15994097b8880cc986d82faf80f57382b998dbae5.
//
// Solidity: event TokenPurchased(address indexed token, address indexed buyer, uint256 trxAmount, uint256 fee, uint256 tokenAmount, uint256 tokenReserve)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchTokenPurchased(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadTokenPurchased, token []common.Address, buyer []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "TokenPurchased", tokenRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadTokenPurchased)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenPurchased", log); err != nil {
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

// ParseTokenPurchased is a log parse operation binding the contract event 0x63abb62535c21a5d221cf9c15994097b8880cc986d82faf80f57382b998dbae5.
//
// Solidity: event TokenPurchased(address indexed token, address indexed buyer, uint256 trxAmount, uint256 fee, uint256 tokenAmount, uint256 tokenReserve)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseTokenPurchased(log types.Log) (*SunpumplaunchpadTokenPurchased, error) {
	event := new(SunpumplaunchpadTokenPurchased)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SunpumplaunchpadTokenSoldIterator is returned from FilterTokenSold and is used to iterate over the raw logs and unpacked data for TokenSold events raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenSoldIterator struct {
	Event *SunpumplaunchpadTokenSold // Event containing the contract specifics and raw log

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
func (it *SunpumplaunchpadTokenSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SunpumplaunchpadTokenSold)
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
		it.Event = new(SunpumplaunchpadTokenSold)
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
func (it *SunpumplaunchpadTokenSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SunpumplaunchpadTokenSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SunpumplaunchpadTokenSold represents a TokenSold event raised by the Sunpumplaunchpad contract.
type SunpumplaunchpadTokenSold struct {
	Token       common.Address
	Seller      common.Address
	TrxAmount   *big.Int
	Fee         *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenSold is a free log retrieval operation binding the contract event 0x9387a595ac4be9038bbb9751abad8baa3dcf219dd9e19abb81552bd521fe3546.
//
// Solidity: event TokenSold(address indexed token, address indexed seller, uint256 trxAmount, uint256 fee, uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) FilterTokenSold(opts *bind.FilterOpts, token []common.Address, seller []common.Address) (*SunpumplaunchpadTokenSoldIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.FilterLogs(opts, "TokenSold", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SunpumplaunchpadTokenSoldIterator{contract: _Sunpumplaunchpad.contract, event: "TokenSold", logs: logs, sub: sub}, nil
}

// WatchTokenSold is a free log subscription operation binding the contract event 0x9387a595ac4be9038bbb9751abad8baa3dcf219dd9e19abb81552bd521fe3546.
//
// Solidity: event TokenSold(address indexed token, address indexed seller, uint256 trxAmount, uint256 fee, uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) WatchTokenSold(opts *bind.WatchOpts, sink chan<- *SunpumplaunchpadTokenSold, token []common.Address, seller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sunpumplaunchpad.contract.WatchLogs(opts, "TokenSold", tokenRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SunpumplaunchpadTokenSold)
				if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenSold", log); err != nil {
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

// ParseTokenSold is a log parse operation binding the contract event 0x9387a595ac4be9038bbb9751abad8baa3dcf219dd9e19abb81552bd521fe3546.
//
// Solidity: event TokenSold(address indexed token, address indexed seller, uint256 trxAmount, uint256 fee, uint256 tokenAmount)
func (_Sunpumplaunchpad *SunpumplaunchpadFilterer) ParseTokenSold(log types.Log) (*SunpumplaunchpadTokenSold, error) {
	event := new(SunpumplaunchpadTokenSold)
	if err := _Sunpumplaunchpad.contract.UnpackLog(event, "TokenSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
