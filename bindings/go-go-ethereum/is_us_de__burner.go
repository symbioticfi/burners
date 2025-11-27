// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardsv2contracts

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

// IsUSDeBurnerMetaData contains all meta data concerning the IsUSDeBurner contract.
var IsUSDeBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveUSDeMinter\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerClaim\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerInstantClaim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[],\"outputs\":[{\"name\":\"requestId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerClaim\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerInstantClaim\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"HasCooldown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAsset\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoCooldown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SufficientApproval\",\"inputs\":[]}]",
}

// IsUSDeBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IsUSDeBurnerMetaData.ABI instead.
var IsUSDeBurnerABI = IsUSDeBurnerMetaData.ABI

// IsUSDeBurner is an auto generated Go binding around an Ethereum contract.
type IsUSDeBurner struct {
	IsUSDeBurnerCaller     // Read-only binding to the contract
	IsUSDeBurnerTransactor // Write-only binding to the contract
	IsUSDeBurnerFilterer   // Log filterer for contract events
}

// IsUSDeBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IsUSDeBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsUSDeBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IsUSDeBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsUSDeBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IsUSDeBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsUSDeBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IsUSDeBurnerSession struct {
	Contract     *IsUSDeBurner     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IsUSDeBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IsUSDeBurnerCallerSession struct {
	Contract *IsUSDeBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IsUSDeBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IsUSDeBurnerTransactorSession struct {
	Contract     *IsUSDeBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IsUSDeBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IsUSDeBurnerRaw struct {
	Contract *IsUSDeBurner // Generic contract binding to access the raw methods on
}

// IsUSDeBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IsUSDeBurnerCallerRaw struct {
	Contract *IsUSDeBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IsUSDeBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IsUSDeBurnerTransactorRaw struct {
	Contract *IsUSDeBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIsUSDeBurner creates a new instance of IsUSDeBurner, bound to a specific deployed contract.
func NewIsUSDeBurner(address common.Address, backend bind.ContractBackend) (*IsUSDeBurner, error) {
	contract, err := bindIsUSDeBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurner{IsUSDeBurnerCaller: IsUSDeBurnerCaller{contract: contract}, IsUSDeBurnerTransactor: IsUSDeBurnerTransactor{contract: contract}, IsUSDeBurnerFilterer: IsUSDeBurnerFilterer{contract: contract}}, nil
}

// NewIsUSDeBurnerCaller creates a new read-only instance of IsUSDeBurner, bound to a specific deployed contract.
func NewIsUSDeBurnerCaller(address common.Address, caller bind.ContractCaller) (*IsUSDeBurnerCaller, error) {
	contract, err := bindIsUSDeBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerCaller{contract: contract}, nil
}

// NewIsUSDeBurnerTransactor creates a new write-only instance of IsUSDeBurner, bound to a specific deployed contract.
func NewIsUSDeBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IsUSDeBurnerTransactor, error) {
	contract, err := bindIsUSDeBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerTransactor{contract: contract}, nil
}

// NewIsUSDeBurnerFilterer creates a new log filterer instance of IsUSDeBurner, bound to a specific deployed contract.
func NewIsUSDeBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IsUSDeBurnerFilterer, error) {
	contract, err := bindIsUSDeBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerFilterer{contract: contract}, nil
}

// bindIsUSDeBurner binds a generic wrapper to an already deployed contract.
func bindIsUSDeBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IsUSDeBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsUSDeBurner *IsUSDeBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IsUSDeBurner.Contract.IsUSDeBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsUSDeBurner *IsUSDeBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.IsUSDeBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsUSDeBurner *IsUSDeBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.IsUSDeBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsUSDeBurner *IsUSDeBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IsUSDeBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsUSDeBurner *IsUSDeBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsUSDeBurner *IsUSDeBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IsUSDeBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerSession) COLLATERAL() (common.Address, error) {
	return _IsUSDeBurner.Contract.COLLATERAL(&_IsUSDeBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IsUSDeBurner.Contract.COLLATERAL(&_IsUSDeBurner.CallOpts)
}

// USDE is a free data retrieval call binding the contract method 0x42de081b.
//
// Solidity: function USDE() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerCaller) USDE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IsUSDeBurner.contract.Call(opts, &out, "USDE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDE is a free data retrieval call binding the contract method 0x42de081b.
//
// Solidity: function USDE() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerSession) USDE() (common.Address, error) {
	return _IsUSDeBurner.Contract.USDE(&_IsUSDeBurner.CallOpts)
}

// USDE is a free data retrieval call binding the contract method 0x42de081b.
//
// Solidity: function USDE() view returns(address)
func (_IsUSDeBurner *IsUSDeBurnerCallerSession) USDE() (common.Address, error) {
	return _IsUSDeBurner.Contract.USDE(&_IsUSDeBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(address[] requestIds)
func (_IsUSDeBurner *IsUSDeBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _IsUSDeBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(address[] requestIds)
func (_IsUSDeBurner *IsUSDeBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]common.Address, error) {
	return _IsUSDeBurner.Contract.RequestIds(&_IsUSDeBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(address[] requestIds)
func (_IsUSDeBurner *IsUSDeBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]common.Address, error) {
	return _IsUSDeBurner.Contract.RequestIds(&_IsUSDeBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsUSDeBurner *IsUSDeBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IsUSDeBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsUSDeBurner *IsUSDeBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _IsUSDeBurner.Contract.RequestIdsLength(&_IsUSDeBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsUSDeBurner *IsUSDeBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IsUSDeBurner.Contract.RequestIdsLength(&_IsUSDeBurner.CallOpts)
}

// ApproveUSDeMinter is a paid mutator transaction binding the contract method 0xe854535e.
//
// Solidity: function approveUSDeMinter() returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactor) ApproveUSDeMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsUSDeBurner.contract.Transact(opts, "approveUSDeMinter")
}

// ApproveUSDeMinter is a paid mutator transaction binding the contract method 0xe854535e.
//
// Solidity: function approveUSDeMinter() returns()
func (_IsUSDeBurner *IsUSDeBurnerSession) ApproveUSDeMinter() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.ApproveUSDeMinter(&_IsUSDeBurner.TransactOpts)
}

// ApproveUSDeMinter is a paid mutator transaction binding the contract method 0xe854535e.
//
// Solidity: function approveUSDeMinter() returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactorSession) ApproveUSDeMinter() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.ApproveUSDeMinter(&_IsUSDeBurner.TransactOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0xe9e17920.
//
// Solidity: function triggerBurn(address asset) returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.contract.Transact(opts, "triggerBurn", asset)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0xe9e17920.
//
// Solidity: function triggerBurn(address asset) returns()
func (_IsUSDeBurner *IsUSDeBurnerSession) TriggerBurn(asset common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerBurn(&_IsUSDeBurner.TransactOpts, asset)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0xe9e17920.
//
// Solidity: function triggerBurn(address asset) returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactorSession) TriggerBurn(asset common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerBurn(&_IsUSDeBurner.TransactOpts, asset)
}

// TriggerClaim is a paid mutator transaction binding the contract method 0x4bf20e20.
//
// Solidity: function triggerClaim(address requestId) returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactor) TriggerClaim(opts *bind.TransactOpts, requestId common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.contract.Transact(opts, "triggerClaim", requestId)
}

// TriggerClaim is a paid mutator transaction binding the contract method 0x4bf20e20.
//
// Solidity: function triggerClaim(address requestId) returns()
func (_IsUSDeBurner *IsUSDeBurnerSession) TriggerClaim(requestId common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerClaim(&_IsUSDeBurner.TransactOpts, requestId)
}

// TriggerClaim is a paid mutator transaction binding the contract method 0x4bf20e20.
//
// Solidity: function triggerClaim(address requestId) returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactorSession) TriggerClaim(requestId common.Address) (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerClaim(&_IsUSDeBurner.TransactOpts, requestId)
}

// TriggerInstantClaim is a paid mutator transaction binding the contract method 0x71f16aad.
//
// Solidity: function triggerInstantClaim() returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactor) TriggerInstantClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsUSDeBurner.contract.Transact(opts, "triggerInstantClaim")
}

// TriggerInstantClaim is a paid mutator transaction binding the contract method 0x71f16aad.
//
// Solidity: function triggerInstantClaim() returns()
func (_IsUSDeBurner *IsUSDeBurnerSession) TriggerInstantClaim() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerInstantClaim(&_IsUSDeBurner.TransactOpts)
}

// TriggerInstantClaim is a paid mutator transaction binding the contract method 0x71f16aad.
//
// Solidity: function triggerInstantClaim() returns()
func (_IsUSDeBurner *IsUSDeBurnerTransactorSession) TriggerInstantClaim() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerInstantClaim(&_IsUSDeBurner.TransactOpts)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(address requestId)
func (_IsUSDeBurner *IsUSDeBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsUSDeBurner.contract.Transact(opts, "triggerWithdrawal")
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(address requestId)
func (_IsUSDeBurner *IsUSDeBurnerSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerWithdrawal(&_IsUSDeBurner.TransactOpts)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(address requestId)
func (_IsUSDeBurner *IsUSDeBurnerTransactorSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _IsUSDeBurner.Contract.TriggerWithdrawal(&_IsUSDeBurner.TransactOpts)
}

// IsUSDeBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerBurnIterator struct {
	Event *IsUSDeBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *IsUSDeBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsUSDeBurnerTriggerBurn)
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
		it.Event = new(IsUSDeBurnerTriggerBurn)
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
func (it *IsUSDeBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsUSDeBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsUSDeBurnerTriggerBurn represents a TriggerBurn event raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerBurn struct {
	Caller common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x10b34285c16e79ef46ec8143f038d4bdad522c314d150d22c43f7ac9c742a327.
//
// Solidity: event TriggerBurn(address indexed caller, address indexed asset, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address, asset []common.Address) (*IsUSDeBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerTriggerBurnIterator{contract: _IsUSDeBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x10b34285c16e79ef46ec8143f038d4bdad522c314d150d22c43f7ac9c742a327.
//
// Solidity: event TriggerBurn(address indexed caller, address indexed asset, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IsUSDeBurnerTriggerBurn, caller []common.Address, asset []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsUSDeBurnerTriggerBurn)
				if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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

// ParseTriggerBurn is a log parse operation binding the contract event 0x10b34285c16e79ef46ec8143f038d4bdad522c314d150d22c43f7ac9c742a327.
//
// Solidity: event TriggerBurn(address indexed caller, address indexed asset, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) ParseTriggerBurn(log types.Log) (*IsUSDeBurnerTriggerBurn, error) {
	event := new(IsUSDeBurnerTriggerBurn)
	if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IsUSDeBurnerTriggerClaimIterator is returned from FilterTriggerClaim and is used to iterate over the raw logs and unpacked data for TriggerClaim events raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerClaimIterator struct {
	Event *IsUSDeBurnerTriggerClaim // Event containing the contract specifics and raw log

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
func (it *IsUSDeBurnerTriggerClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsUSDeBurnerTriggerClaim)
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
		it.Event = new(IsUSDeBurnerTriggerClaim)
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
func (it *IsUSDeBurnerTriggerClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsUSDeBurnerTriggerClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsUSDeBurnerTriggerClaim represents a TriggerClaim event raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerClaim struct {
	Caller    common.Address
	RequestId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerClaim is a free log retrieval operation binding the contract event 0x57e41aed337843ff9b9328b03a65f1ef655c06d5dd77f62f3ee089ce69ddd511.
//
// Solidity: event TriggerClaim(address indexed caller, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) FilterTriggerClaim(opts *bind.FilterOpts, caller []common.Address) (*IsUSDeBurnerTriggerClaimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.FilterLogs(opts, "TriggerClaim", callerRule)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerTriggerClaimIterator{contract: _IsUSDeBurner.contract, event: "TriggerClaim", logs: logs, sub: sub}, nil
}

// WatchTriggerClaim is a free log subscription operation binding the contract event 0x57e41aed337843ff9b9328b03a65f1ef655c06d5dd77f62f3ee089ce69ddd511.
//
// Solidity: event TriggerClaim(address indexed caller, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) WatchTriggerClaim(opts *bind.WatchOpts, sink chan<- *IsUSDeBurnerTriggerClaim, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.WatchLogs(opts, "TriggerClaim", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsUSDeBurnerTriggerClaim)
				if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerClaim", log); err != nil {
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

// ParseTriggerClaim is a log parse operation binding the contract event 0x57e41aed337843ff9b9328b03a65f1ef655c06d5dd77f62f3ee089ce69ddd511.
//
// Solidity: event TriggerClaim(address indexed caller, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) ParseTriggerClaim(log types.Log) (*IsUSDeBurnerTriggerClaim, error) {
	event := new(IsUSDeBurnerTriggerClaim)
	if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IsUSDeBurnerTriggerInstantClaimIterator is returned from FilterTriggerInstantClaim and is used to iterate over the raw logs and unpacked data for TriggerInstantClaim events raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerInstantClaimIterator struct {
	Event *IsUSDeBurnerTriggerInstantClaim // Event containing the contract specifics and raw log

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
func (it *IsUSDeBurnerTriggerInstantClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsUSDeBurnerTriggerInstantClaim)
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
		it.Event = new(IsUSDeBurnerTriggerInstantClaim)
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
func (it *IsUSDeBurnerTriggerInstantClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsUSDeBurnerTriggerInstantClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsUSDeBurnerTriggerInstantClaim represents a TriggerInstantClaim event raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerInstantClaim struct {
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTriggerInstantClaim is a free log retrieval operation binding the contract event 0x44ce5b54555a9c96767133da974ebafd206eafbbdc523f635e09f75b7f4171d6.
//
// Solidity: event TriggerInstantClaim(address indexed caller, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) FilterTriggerInstantClaim(opts *bind.FilterOpts, caller []common.Address) (*IsUSDeBurnerTriggerInstantClaimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.FilterLogs(opts, "TriggerInstantClaim", callerRule)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerTriggerInstantClaimIterator{contract: _IsUSDeBurner.contract, event: "TriggerInstantClaim", logs: logs, sub: sub}, nil
}

// WatchTriggerInstantClaim is a free log subscription operation binding the contract event 0x44ce5b54555a9c96767133da974ebafd206eafbbdc523f635e09f75b7f4171d6.
//
// Solidity: event TriggerInstantClaim(address indexed caller, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) WatchTriggerInstantClaim(opts *bind.WatchOpts, sink chan<- *IsUSDeBurnerTriggerInstantClaim, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.WatchLogs(opts, "TriggerInstantClaim", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsUSDeBurnerTriggerInstantClaim)
				if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerInstantClaim", log); err != nil {
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

// ParseTriggerInstantClaim is a log parse operation binding the contract event 0x44ce5b54555a9c96767133da974ebafd206eafbbdc523f635e09f75b7f4171d6.
//
// Solidity: event TriggerInstantClaim(address indexed caller, uint256 amount)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) ParseTriggerInstantClaim(log types.Log) (*IsUSDeBurnerTriggerInstantClaim, error) {
	event := new(IsUSDeBurnerTriggerInstantClaim)
	if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerInstantClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IsUSDeBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerWithdrawalIterator struct {
	Event *IsUSDeBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

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
func (it *IsUSDeBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsUSDeBurnerTriggerWithdrawal)
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
		it.Event = new(IsUSDeBurnerTriggerWithdrawal)
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
func (it *IsUSDeBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsUSDeBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsUSDeBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the IsUSDeBurner contract.
type IsUSDeBurnerTriggerWithdrawal struct {
	Caller    common.Address
	Amount    *big.Int
	RequestId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0xa3a2f7ea0a86c4c3a1be020951dce1acdc9e7b12aac6fe2104f3c90d7e9ee693.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 amount, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*IsUSDeBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &IsUSDeBurnerTriggerWithdrawalIterator{contract: _IsUSDeBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0xa3a2f7ea0a86c4c3a1be020951dce1acdc9e7b12aac6fe2104f3c90d7e9ee693.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 amount, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *IsUSDeBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsUSDeBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsUSDeBurnerTriggerWithdrawal)
				if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
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

// ParseTriggerWithdrawal is a log parse operation binding the contract event 0xa3a2f7ea0a86c4c3a1be020951dce1acdc9e7b12aac6fe2104f3c90d7e9ee693.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 amount, address requestId)
func (_IsUSDeBurner *IsUSDeBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*IsUSDeBurnerTriggerWithdrawal, error) {
	event := new(IsUSDeBurnerTriggerWithdrawal)
	if err := _IsUSDeBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
