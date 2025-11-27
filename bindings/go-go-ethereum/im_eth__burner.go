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

// ImETHBurnerMetaData contains all meta data concerning the ImETHBurner contract.
var ImETHBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// ImETHBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use ImETHBurnerMetaData.ABI instead.
var ImETHBurnerABI = ImETHBurnerMetaData.ABI

// ImETHBurner is an auto generated Go binding around an Ethereum contract.
type ImETHBurner struct {
	ImETHBurnerCaller     // Read-only binding to the contract
	ImETHBurnerTransactor // Write-only binding to the contract
	ImETHBurnerFilterer   // Log filterer for contract events
}

// ImETHBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImETHBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImETHBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImETHBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImETHBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImETHBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImETHBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImETHBurnerSession struct {
	Contract     *ImETHBurner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImETHBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImETHBurnerCallerSession struct {
	Contract *ImETHBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ImETHBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImETHBurnerTransactorSession struct {
	Contract     *ImETHBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ImETHBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImETHBurnerRaw struct {
	Contract *ImETHBurner // Generic contract binding to access the raw methods on
}

// ImETHBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImETHBurnerCallerRaw struct {
	Contract *ImETHBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// ImETHBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImETHBurnerTransactorRaw struct {
	Contract *ImETHBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImETHBurner creates a new instance of ImETHBurner, bound to a specific deployed contract.
func NewImETHBurner(address common.Address, backend bind.ContractBackend) (*ImETHBurner, error) {
	contract, err := bindImETHBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImETHBurner{ImETHBurnerCaller: ImETHBurnerCaller{contract: contract}, ImETHBurnerTransactor: ImETHBurnerTransactor{contract: contract}, ImETHBurnerFilterer: ImETHBurnerFilterer{contract: contract}}, nil
}

// NewImETHBurnerCaller creates a new read-only instance of ImETHBurner, bound to a specific deployed contract.
func NewImETHBurnerCaller(address common.Address, caller bind.ContractCaller) (*ImETHBurnerCaller, error) {
	contract, err := bindImETHBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImETHBurnerCaller{contract: contract}, nil
}

// NewImETHBurnerTransactor creates a new write-only instance of ImETHBurner, bound to a specific deployed contract.
func NewImETHBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*ImETHBurnerTransactor, error) {
	contract, err := bindImETHBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImETHBurnerTransactor{contract: contract}, nil
}

// NewImETHBurnerFilterer creates a new log filterer instance of ImETHBurner, bound to a specific deployed contract.
func NewImETHBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*ImETHBurnerFilterer, error) {
	contract, err := bindImETHBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImETHBurnerFilterer{contract: contract}, nil
}

// bindImETHBurner binds a generic wrapper to an already deployed contract.
func bindImETHBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ImETHBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImETHBurner *ImETHBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImETHBurner.Contract.ImETHBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImETHBurner *ImETHBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImETHBurner.Contract.ImETHBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImETHBurner *ImETHBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImETHBurner.Contract.ImETHBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImETHBurner *ImETHBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImETHBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImETHBurner *ImETHBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImETHBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImETHBurner *ImETHBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImETHBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_ImETHBurner *ImETHBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ImETHBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_ImETHBurner *ImETHBurnerSession) COLLATERAL() (common.Address, error) {
	return _ImETHBurner.Contract.COLLATERAL(&_ImETHBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_ImETHBurner *ImETHBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _ImETHBurner.Contract.COLLATERAL(&_ImETHBurner.CallOpts)
}

// STAKING is a free data retrieval call binding the contract method 0x97610f30.
//
// Solidity: function STAKING() view returns(address)
func (_ImETHBurner *ImETHBurnerCaller) STAKING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ImETHBurner.contract.Call(opts, &out, "STAKING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKING is a free data retrieval call binding the contract method 0x97610f30.
//
// Solidity: function STAKING() view returns(address)
func (_ImETHBurner *ImETHBurnerSession) STAKING() (common.Address, error) {
	return _ImETHBurner.Contract.STAKING(&_ImETHBurner.CallOpts)
}

// STAKING is a free data retrieval call binding the contract method 0x97610f30.
//
// Solidity: function STAKING() view returns(address)
func (_ImETHBurner *ImETHBurnerCallerSession) STAKING() (common.Address, error) {
	return _ImETHBurner.Contract.STAKING(&_ImETHBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_ImETHBurner *ImETHBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ImETHBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_ImETHBurner *ImETHBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _ImETHBurner.Contract.RequestIds(&_ImETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_ImETHBurner *ImETHBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _ImETHBurner.Contract.RequestIds(&_ImETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_ImETHBurner *ImETHBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ImETHBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_ImETHBurner *ImETHBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _ImETHBurner.Contract.RequestIdsLength(&_ImETHBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_ImETHBurner *ImETHBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _ImETHBurner.Contract.RequestIdsLength(&_ImETHBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_ImETHBurner *ImETHBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _ImETHBurner.contract.Transact(opts, "triggerBurn", requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_ImETHBurner *ImETHBurnerSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _ImETHBurner.Contract.TriggerBurn(&_ImETHBurner.TransactOpts, requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_ImETHBurner *ImETHBurnerTransactorSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _ImETHBurner.Contract.TriggerBurn(&_ImETHBurner.TransactOpts, requestId)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_ImETHBurner *ImETHBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImETHBurner.contract.Transact(opts, "triggerWithdrawal")
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_ImETHBurner *ImETHBurnerSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _ImETHBurner.Contract.TriggerWithdrawal(&_ImETHBurner.TransactOpts)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x041e0185.
//
// Solidity: function triggerWithdrawal() returns(uint256 requestId)
func (_ImETHBurner *ImETHBurnerTransactorSession) TriggerWithdrawal() (*types.Transaction, error) {
	return _ImETHBurner.Contract.TriggerWithdrawal(&_ImETHBurner.TransactOpts)
}

// ImETHBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the ImETHBurner contract.
type ImETHBurnerTriggerBurnIterator struct {
	Event *ImETHBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *ImETHBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImETHBurnerTriggerBurn)
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
		it.Event = new(ImETHBurnerTriggerBurn)
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
func (it *ImETHBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImETHBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImETHBurnerTriggerBurn represents a TriggerBurn event raised by the ImETHBurner contract.
type ImETHBurnerTriggerBurn struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_ImETHBurner *ImETHBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*ImETHBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ImETHBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &ImETHBurnerTriggerBurnIterator{contract: _ImETHBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_ImETHBurner *ImETHBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *ImETHBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ImETHBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImETHBurnerTriggerBurn)
				if err := _ImETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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
func (_ImETHBurner *ImETHBurnerFilterer) ParseTriggerBurn(log types.Log) (*ImETHBurnerTriggerBurn, error) {
	event := new(ImETHBurnerTriggerBurn)
	if err := _ImETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImETHBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the ImETHBurner contract.
type ImETHBurnerTriggerWithdrawalIterator struct {
	Event *ImETHBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

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
func (it *ImETHBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImETHBurnerTriggerWithdrawal)
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
		it.Event = new(ImETHBurnerTriggerWithdrawal)
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
func (it *ImETHBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImETHBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImETHBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the ImETHBurner contract.
type ImETHBurnerTriggerWithdrawal struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_ImETHBurner *ImETHBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*ImETHBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ImETHBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &ImETHBurnerTriggerWithdrawalIterator{contract: _ImETHBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_ImETHBurner *ImETHBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *ImETHBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ImETHBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImETHBurnerTriggerWithdrawal)
				if err := _ImETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
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
func (_ImETHBurner *ImETHBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*ImETHBurnerTriggerWithdrawal, error) {
	event := new(ImETHBurnerTriggerWithdrawal)
	if err := _ImETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
