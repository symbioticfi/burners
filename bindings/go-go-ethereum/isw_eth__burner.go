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

// IswETHBurnerMetaData contains all meta data concerning the IswETHBurner contract.
var IswETHBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SWEXIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[{\"name\":\"maxRequests\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"firstRequestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastRequestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"firstRequestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lastRequestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// IswETHBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IswETHBurnerMetaData.ABI instead.
var IswETHBurnerABI = IswETHBurnerMetaData.ABI

// IswETHBurner is an auto generated Go binding around an Ethereum contract.
type IswETHBurner struct {
	IswETHBurnerCaller     // Read-only binding to the contract
	IswETHBurnerTransactor // Write-only binding to the contract
	IswETHBurnerFilterer   // Log filterer for contract events
}

// IswETHBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IswETHBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IswETHBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IswETHBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IswETHBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IswETHBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IswETHBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IswETHBurnerSession struct {
	Contract     *IswETHBurner     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IswETHBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IswETHBurnerCallerSession struct {
	Contract *IswETHBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IswETHBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IswETHBurnerTransactorSession struct {
	Contract     *IswETHBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IswETHBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IswETHBurnerRaw struct {
	Contract *IswETHBurner // Generic contract binding to access the raw methods on
}

// IswETHBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IswETHBurnerCallerRaw struct {
	Contract *IswETHBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IswETHBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IswETHBurnerTransactorRaw struct {
	Contract *IswETHBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIswETHBurner creates a new instance of IswETHBurner, bound to a specific deployed contract.
func NewIswETHBurner(address common.Address, backend bind.ContractBackend) (*IswETHBurner, error) {
	contract, err := bindIswETHBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IswETHBurner{IswETHBurnerCaller: IswETHBurnerCaller{contract: contract}, IswETHBurnerTransactor: IswETHBurnerTransactor{contract: contract}, IswETHBurnerFilterer: IswETHBurnerFilterer{contract: contract}}, nil
}

// NewIswETHBurnerCaller creates a new read-only instance of IswETHBurner, bound to a specific deployed contract.
func NewIswETHBurnerCaller(address common.Address, caller bind.ContractCaller) (*IswETHBurnerCaller, error) {
	contract, err := bindIswETHBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IswETHBurnerCaller{contract: contract}, nil
}

// NewIswETHBurnerTransactor creates a new write-only instance of IswETHBurner, bound to a specific deployed contract.
func NewIswETHBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IswETHBurnerTransactor, error) {
	contract, err := bindIswETHBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IswETHBurnerTransactor{contract: contract}, nil
}

// NewIswETHBurnerFilterer creates a new log filterer instance of IswETHBurner, bound to a specific deployed contract.
func NewIswETHBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IswETHBurnerFilterer, error) {
	contract, err := bindIswETHBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IswETHBurnerFilterer{contract: contract}, nil
}

// bindIswETHBurner binds a generic wrapper to an already deployed contract.
func bindIswETHBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IswETHBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IswETHBurner *IswETHBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IswETHBurner.Contract.IswETHBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IswETHBurner *IswETHBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IswETHBurner.Contract.IswETHBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IswETHBurner *IswETHBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IswETHBurner.Contract.IswETHBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IswETHBurner *IswETHBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IswETHBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IswETHBurner *IswETHBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IswETHBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IswETHBurner *IswETHBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IswETHBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IswETHBurner *IswETHBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IswETHBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IswETHBurner *IswETHBurnerSession) COLLATERAL() (common.Address, error) {
	return _IswETHBurner.Contract.COLLATERAL(&_IswETHBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IswETHBurner *IswETHBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IswETHBurner.Contract.COLLATERAL(&_IswETHBurner.CallOpts)
}

// SWEXIT is a free data retrieval call binding the contract method 0x127ed559.
//
// Solidity: function SWEXIT() view returns(address)
func (_IswETHBurner *IswETHBurnerCaller) SWEXIT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IswETHBurner.contract.Call(opts, &out, "SWEXIT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWEXIT is a free data retrieval call binding the contract method 0x127ed559.
//
// Solidity: function SWEXIT() view returns(address)
func (_IswETHBurner *IswETHBurnerSession) SWEXIT() (common.Address, error) {
	return _IswETHBurner.Contract.SWEXIT(&_IswETHBurner.CallOpts)
}

// SWEXIT is a free data retrieval call binding the contract method 0x127ed559.
//
// Solidity: function SWEXIT() view returns(address)
func (_IswETHBurner *IswETHBurnerCallerSession) SWEXIT() (common.Address, error) {
	return _IswETHBurner.Contract.SWEXIT(&_IswETHBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IswETHBurner *IswETHBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IswETHBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IswETHBurner *IswETHBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IswETHBurner.Contract.RequestIds(&_IswETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IswETHBurner *IswETHBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IswETHBurner.Contract.RequestIds(&_IswETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IswETHBurner *IswETHBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IswETHBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IswETHBurner *IswETHBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _IswETHBurner.Contract.RequestIdsLength(&_IswETHBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IswETHBurner *IswETHBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IswETHBurner.Contract.RequestIdsLength(&_IswETHBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IswETHBurner *IswETHBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.contract.Transact(opts, "triggerBurn", requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IswETHBurner *IswETHBurnerSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.Contract.TriggerBurn(&_IswETHBurner.TransactOpts, requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IswETHBurner *IswETHBurnerTransactorSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.Contract.TriggerBurn(&_IswETHBurner.TransactOpts, requestId)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts, maxRequests *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.contract.Transact(opts, "triggerWithdrawal", maxRequests)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerSession) TriggerWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.Contract.TriggerWithdrawal(&_IswETHBurner.TransactOpts, maxRequests)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerTransactorSession) TriggerWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _IswETHBurner.Contract.TriggerWithdrawal(&_IswETHBurner.TransactOpts, maxRequests)
}

// IswETHBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IswETHBurner contract.
type IswETHBurnerTriggerBurnIterator struct {
	Event *IswETHBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *IswETHBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IswETHBurnerTriggerBurn)
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
		it.Event = new(IswETHBurnerTriggerBurn)
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
func (it *IswETHBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IswETHBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IswETHBurnerTriggerBurn represents a TriggerBurn event raised by the IswETHBurner contract.
type IswETHBurnerTriggerBurn struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IswETHBurner *IswETHBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*IswETHBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IswETHBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &IswETHBurnerTriggerBurnIterator{contract: _IswETHBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IswETHBurner *IswETHBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IswETHBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IswETHBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IswETHBurnerTriggerBurn)
				if err := _IswETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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
func (_IswETHBurner *IswETHBurnerFilterer) ParseTriggerBurn(log types.Log) (*IswETHBurnerTriggerBurn, error) {
	event := new(IswETHBurnerTriggerBurn)
	if err := _IswETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IswETHBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the IswETHBurner contract.
type IswETHBurnerTriggerWithdrawalIterator struct {
	Event *IswETHBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

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
func (it *IswETHBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IswETHBurnerTriggerWithdrawal)
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
		it.Event = new(IswETHBurnerTriggerWithdrawal)
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
func (it *IswETHBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IswETHBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IswETHBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the IswETHBurner contract.
type IswETHBurnerTriggerWithdrawal struct {
	Caller         common.Address
	FirstRequestId *big.Int
	LastRequestId  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0x2af9b173527ef5f4bc3130bac428abdaf7646958117405d702f2d49774b79c12.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*IswETHBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IswETHBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &IswETHBurnerTriggerWithdrawalIterator{contract: _IswETHBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0x2af9b173527ef5f4bc3130bac428abdaf7646958117405d702f2d49774b79c12.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *IswETHBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IswETHBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IswETHBurnerTriggerWithdrawal)
				if err := _IswETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
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

// ParseTriggerWithdrawal is a log parse operation binding the contract event 0x2af9b173527ef5f4bc3130bac428abdaf7646958117405d702f2d49774b79c12.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 firstRequestId, uint256 lastRequestId)
func (_IswETHBurner *IswETHBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*IswETHBurnerTriggerWithdrawal, error) {
	event := new(IswETHBurnerTriggerWithdrawal)
	if err := _IswETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
