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

// IsfrxETHBurnerMetaData contains all meta data concerning the IsfrxETHBurner contract.
var IsfrxETHBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FRAX_ETHER_REDEMPTION_QUEUE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// IsfrxETHBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IsfrxETHBurnerMetaData.ABI instead.
var IsfrxETHBurnerABI = IsfrxETHBurnerMetaData.ABI

// IsfrxETHBurner is an auto generated Go binding around an Ethereum contract.
type IsfrxETHBurner struct {
	IsfrxETHBurnerCaller     // Read-only binding to the contract
	IsfrxETHBurnerTransactor // Write-only binding to the contract
	IsfrxETHBurnerFilterer   // Log filterer for contract events
}

// IsfrxETHBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IsfrxETHBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsfrxETHBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IsfrxETHBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsfrxETHBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IsfrxETHBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IsfrxETHBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IsfrxETHBurnerSession struct {
	Contract     *IsfrxETHBurner   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IsfrxETHBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IsfrxETHBurnerCallerSession struct {
	Contract *IsfrxETHBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IsfrxETHBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IsfrxETHBurnerTransactorSession struct {
	Contract     *IsfrxETHBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IsfrxETHBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IsfrxETHBurnerRaw struct {
	Contract *IsfrxETHBurner // Generic contract binding to access the raw methods on
}

// IsfrxETHBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IsfrxETHBurnerCallerRaw struct {
	Contract *IsfrxETHBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IsfrxETHBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IsfrxETHBurnerTransactorRaw struct {
	Contract *IsfrxETHBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIsfrxETHBurner creates a new instance of IsfrxETHBurner, bound to a specific deployed contract.
func NewIsfrxETHBurner(address common.Address, backend bind.ContractBackend) (*IsfrxETHBurner, error) {
	contract, err := bindIsfrxETHBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurner{IsfrxETHBurnerCaller: IsfrxETHBurnerCaller{contract: contract}, IsfrxETHBurnerTransactor: IsfrxETHBurnerTransactor{contract: contract}, IsfrxETHBurnerFilterer: IsfrxETHBurnerFilterer{contract: contract}}, nil
}

// NewIsfrxETHBurnerCaller creates a new read-only instance of IsfrxETHBurner, bound to a specific deployed contract.
func NewIsfrxETHBurnerCaller(address common.Address, caller bind.ContractCaller) (*IsfrxETHBurnerCaller, error) {
	contract, err := bindIsfrxETHBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurnerCaller{contract: contract}, nil
}

// NewIsfrxETHBurnerTransactor creates a new write-only instance of IsfrxETHBurner, bound to a specific deployed contract.
func NewIsfrxETHBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IsfrxETHBurnerTransactor, error) {
	contract, err := bindIsfrxETHBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurnerTransactor{contract: contract}, nil
}

// NewIsfrxETHBurnerFilterer creates a new log filterer instance of IsfrxETHBurner, bound to a specific deployed contract.
func NewIsfrxETHBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IsfrxETHBurnerFilterer, error) {
	contract, err := bindIsfrxETHBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurnerFilterer{contract: contract}, nil
}

// bindIsfrxETHBurner binds a generic wrapper to an already deployed contract.
func bindIsfrxETHBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IsfrxETHBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsfrxETHBurner *IsfrxETHBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IsfrxETHBurner.Contract.IsfrxETHBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsfrxETHBurner *IsfrxETHBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.IsfrxETHBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsfrxETHBurner *IsfrxETHBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.IsfrxETHBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IsfrxETHBurner *IsfrxETHBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IsfrxETHBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IsfrxETHBurner *IsfrxETHBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IsfrxETHBurner *IsfrxETHBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IsfrxETHBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerSession) COLLATERAL() (common.Address, error) {
	return _IsfrxETHBurner.Contract.COLLATERAL(&_IsfrxETHBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IsfrxETHBurner.Contract.COLLATERAL(&_IsfrxETHBurner.CallOpts)
}

// FRAXETHERREDEMPTIONQUEUE is a free data retrieval call binding the contract method 0xea98e28c.
//
// Solidity: function FRAX_ETHER_REDEMPTION_QUEUE() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerCaller) FRAXETHERREDEMPTIONQUEUE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IsfrxETHBurner.contract.Call(opts, &out, "FRAX_ETHER_REDEMPTION_QUEUE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FRAXETHERREDEMPTIONQUEUE is a free data retrieval call binding the contract method 0xea98e28c.
//
// Solidity: function FRAX_ETHER_REDEMPTION_QUEUE() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerSession) FRAXETHERREDEMPTIONQUEUE() (common.Address, error) {
	return _IsfrxETHBurner.Contract.FRAXETHERREDEMPTIONQUEUE(&_IsfrxETHBurner.CallOpts)
}

// FRAXETHERREDEMPTIONQUEUE is a free data retrieval call binding the contract method 0xea98e28c.
//
// Solidity: function FRAX_ETHER_REDEMPTION_QUEUE() view returns(address)
func (_IsfrxETHBurner *IsfrxETHBurnerCallerSession) FRAXETHERREDEMPTIONQUEUE() (common.Address, error) {
	return _IsfrxETHBurner.Contract.FRAXETHERREDEMPTIONQUEUE(&_IsfrxETHBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IsfrxETHBurner *IsfrxETHBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IsfrxETHBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IsfrxETHBurner *IsfrxETHBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IsfrxETHBurner.Contract.RequestIds(&_IsfrxETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IsfrxETHBurner *IsfrxETHBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IsfrxETHBurner.Contract.RequestIds(&_IsfrxETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsfrxETHBurner *IsfrxETHBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IsfrxETHBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsfrxETHBurner *IsfrxETHBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _IsfrxETHBurner.Contract.RequestIdsLength(&_IsfrxETHBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IsfrxETHBurner *IsfrxETHBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IsfrxETHBurner.Contract.RequestIdsLength(&_IsfrxETHBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IsfrxETHBurner *IsfrxETHBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IsfrxETHBurner.contract.Transact(opts, "triggerBurn", requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IsfrxETHBurner *IsfrxETHBurnerSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.TriggerBurn(&_IsfrxETHBurner.TransactOpts, requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IsfrxETHBurner *IsfrxETHBurnerTransactorSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.TriggerBurn(&_IsfrxETHBurner.TransactOpts, requestId)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IsfrxETHBurner.contract.Transact(opts, "triggerWithdrawal")
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.TriggerWithdrawal(&_IsfrxETHBurner.TransactOpts)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerTransactorSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _IsfrxETHBurner.Contract.TriggerWithdrawal(&_IsfrxETHBurner.TransactOpts)
}

// IsfrxETHBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IsfrxETHBurner contract.
type IsfrxETHBurnerTriggerBurnIterator struct {
	Event *IsfrxETHBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *IsfrxETHBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsfrxETHBurnerTriggerBurn)
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
		it.Event = new(IsfrxETHBurnerTriggerBurn)
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
func (it *IsfrxETHBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsfrxETHBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsfrxETHBurnerTriggerBurn represents a TriggerBurn event raised by the IsfrxETHBurner contract.
type IsfrxETHBurnerTriggerBurn struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*IsfrxETHBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsfrxETHBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurnerTriggerBurnIterator{contract: _IsfrxETHBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IsfrxETHBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsfrxETHBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsfrxETHBurnerTriggerBurn)
				if err := _IsfrxETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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

// ParseTriggerBurn is a log parse operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) ParseTriggerBurn(log types.Log) (*IsfrxETHBurnerTriggerBurn, error) {
	event := new(IsfrxETHBurnerTriggerBurn)
	if err := _IsfrxETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IsfrxETHBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the IsfrxETHBurner contract.
type IsfrxETHBurnerTriggerWithdrawalIterator struct {
	Event *IsfrxETHBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

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
func (it *IsfrxETHBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IsfrxETHBurnerTriggerWithdrawal)
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
		it.Event = new(IsfrxETHBurnerTriggerWithdrawal)
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
func (it *IsfrxETHBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IsfrxETHBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IsfrxETHBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the IsfrxETHBurner contract.
type IsfrxETHBurnerTriggerWithdrawal struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*IsfrxETHBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsfrxETHBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &IsfrxETHBurnerTriggerWithdrawalIterator{contract: _IsfrxETHBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *IsfrxETHBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IsfrxETHBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IsfrxETHBurnerTriggerWithdrawal)
				if err := _IsfrxETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
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

// ParseTriggerWithdrawal is a log parse operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_IsfrxETHBurner *IsfrxETHBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*IsfrxETHBurnerTriggerWithdrawal, error) {
	event := new(IsfrxETHBurnerTriggerWithdrawal)
	if err := _IsfrxETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
