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

// IETHxBurnerMetaData contains all meta data concerning the IETHxBurner contract.
var IETHxBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STADER_CONFIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKE_POOLS_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USER_WITHDRAW_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[{\"name\":\"maxWithdrawalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidETHxMaximumWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// IETHxBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IETHxBurnerMetaData.ABI instead.
var IETHxBurnerABI = IETHxBurnerMetaData.ABI

// IETHxBurner is an auto generated Go binding around an Ethereum contract.
type IETHxBurner struct {
	IETHxBurnerCaller     // Read-only binding to the contract
	IETHxBurnerTransactor // Write-only binding to the contract
	IETHxBurnerFilterer   // Log filterer for contract events
}

// IETHxBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IETHxBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHxBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IETHxBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHxBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IETHxBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHxBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IETHxBurnerSession struct {
	Contract     *IETHxBurner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IETHxBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IETHxBurnerCallerSession struct {
	Contract *IETHxBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IETHxBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IETHxBurnerTransactorSession struct {
	Contract     *IETHxBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IETHxBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IETHxBurnerRaw struct {
	Contract *IETHxBurner // Generic contract binding to access the raw methods on
}

// IETHxBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IETHxBurnerCallerRaw struct {
	Contract *IETHxBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IETHxBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IETHxBurnerTransactorRaw struct {
	Contract *IETHxBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIETHxBurner creates a new instance of IETHxBurner, bound to a specific deployed contract.
func NewIETHxBurner(address common.Address, backend bind.ContractBackend) (*IETHxBurner, error) {
	contract, err := bindIETHxBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IETHxBurner{IETHxBurnerCaller: IETHxBurnerCaller{contract: contract}, IETHxBurnerTransactor: IETHxBurnerTransactor{contract: contract}, IETHxBurnerFilterer: IETHxBurnerFilterer{contract: contract}}, nil
}

// NewIETHxBurnerCaller creates a new read-only instance of IETHxBurner, bound to a specific deployed contract.
func NewIETHxBurnerCaller(address common.Address, caller bind.ContractCaller) (*IETHxBurnerCaller, error) {
	contract, err := bindIETHxBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IETHxBurnerCaller{contract: contract}, nil
}

// NewIETHxBurnerTransactor creates a new write-only instance of IETHxBurner, bound to a specific deployed contract.
func NewIETHxBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IETHxBurnerTransactor, error) {
	contract, err := bindIETHxBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IETHxBurnerTransactor{contract: contract}, nil
}

// NewIETHxBurnerFilterer creates a new log filterer instance of IETHxBurner, bound to a specific deployed contract.
func NewIETHxBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IETHxBurnerFilterer, error) {
	contract, err := bindIETHxBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IETHxBurnerFilterer{contract: contract}, nil
}

// bindIETHxBurner binds a generic wrapper to an already deployed contract.
func bindIETHxBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IETHxBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHxBurner *IETHxBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHxBurner.Contract.IETHxBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHxBurner *IETHxBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHxBurner.Contract.IETHxBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHxBurner *IETHxBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHxBurner.Contract.IETHxBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHxBurner *IETHxBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHxBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHxBurner *IETHxBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHxBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHxBurner *IETHxBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHxBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IETHxBurner *IETHxBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IETHxBurner *IETHxBurnerSession) COLLATERAL() (common.Address, error) {
	return _IETHxBurner.Contract.COLLATERAL(&_IETHxBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IETHxBurner *IETHxBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IETHxBurner.Contract.COLLATERAL(&_IETHxBurner.CallOpts)
}

// STADERCONFIG is a free data retrieval call binding the contract method 0x6c50e8b7.
//
// Solidity: function STADER_CONFIG() view returns(address)
func (_IETHxBurner *IETHxBurnerCaller) STADERCONFIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "STADER_CONFIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STADERCONFIG is a free data retrieval call binding the contract method 0x6c50e8b7.
//
// Solidity: function STADER_CONFIG() view returns(address)
func (_IETHxBurner *IETHxBurnerSession) STADERCONFIG() (common.Address, error) {
	return _IETHxBurner.Contract.STADERCONFIG(&_IETHxBurner.CallOpts)
}

// STADERCONFIG is a free data retrieval call binding the contract method 0x6c50e8b7.
//
// Solidity: function STADER_CONFIG() view returns(address)
func (_IETHxBurner *IETHxBurnerCallerSession) STADERCONFIG() (common.Address, error) {
	return _IETHxBurner.Contract.STADERCONFIG(&_IETHxBurner.CallOpts)
}

// STAKEPOOLSMANAGER is a free data retrieval call binding the contract method 0x05f4d08a.
//
// Solidity: function STAKE_POOLS_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerCaller) STAKEPOOLSMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "STAKE_POOLS_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEPOOLSMANAGER is a free data retrieval call binding the contract method 0x05f4d08a.
//
// Solidity: function STAKE_POOLS_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerSession) STAKEPOOLSMANAGER() (common.Address, error) {
	return _IETHxBurner.Contract.STAKEPOOLSMANAGER(&_IETHxBurner.CallOpts)
}

// STAKEPOOLSMANAGER is a free data retrieval call binding the contract method 0x05f4d08a.
//
// Solidity: function STAKE_POOLS_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerCallerSession) STAKEPOOLSMANAGER() (common.Address, error) {
	return _IETHxBurner.Contract.STAKEPOOLSMANAGER(&_IETHxBurner.CallOpts)
}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerCaller) USERWITHDRAWMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "USER_WITHDRAW_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerSession) USERWITHDRAWMANAGER() (common.Address, error) {
	return _IETHxBurner.Contract.USERWITHDRAWMANAGER(&_IETHxBurner.CallOpts)
}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(address)
func (_IETHxBurner *IETHxBurnerCallerSession) USERWITHDRAWMANAGER() (common.Address, error) {
	return _IETHxBurner.Contract.USERWITHDRAWMANAGER(&_IETHxBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IETHxBurner *IETHxBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IETHxBurner *IETHxBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IETHxBurner.Contract.RequestIds(&_IETHxBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IETHxBurner *IETHxBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IETHxBurner.Contract.RequestIds(&_IETHxBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IETHxBurner *IETHxBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IETHxBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IETHxBurner *IETHxBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _IETHxBurner.Contract.RequestIdsLength(&_IETHxBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IETHxBurner *IETHxBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IETHxBurner.Contract.RequestIdsLength(&_IETHxBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IETHxBurner *IETHxBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.contract.Transact(opts, "triggerBurn", requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IETHxBurner *IETHxBurnerSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.Contract.TriggerBurn(&_IETHxBurner.TransactOpts, requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IETHxBurner *IETHxBurnerTransactorSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.Contract.TriggerBurn(&_IETHxBurner.TransactOpts, requestId)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxWithdrawalAmount) returns(uint256 requestId)
func (_IETHxBurner *IETHxBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts, maxWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.contract.Transact(opts, "triggerWithdrawal", maxWithdrawalAmount)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxWithdrawalAmount) returns(uint256 requestId)
func (_IETHxBurner *IETHxBurnerSession) TriggerWithdrawal(maxWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.Contract.TriggerWithdrawal(&_IETHxBurner.TransactOpts, maxWithdrawalAmount)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxWithdrawalAmount) returns(uint256 requestId)
func (_IETHxBurner *IETHxBurnerTransactorSession) TriggerWithdrawal(maxWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _IETHxBurner.Contract.TriggerWithdrawal(&_IETHxBurner.TransactOpts, maxWithdrawalAmount)
}

// IETHxBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IETHxBurner contract.
type IETHxBurnerTriggerBurnIterator struct {
	Event *IETHxBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *IETHxBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHxBurnerTriggerBurn)
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
		it.Event = new(IETHxBurnerTriggerBurn)
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
func (it *IETHxBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHxBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHxBurnerTriggerBurn represents a TriggerBurn event raised by the IETHxBurner contract.
type IETHxBurnerTriggerBurn struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IETHxBurner *IETHxBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*IETHxBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IETHxBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &IETHxBurnerTriggerBurnIterator{contract: _IETHxBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IETHxBurner *IETHxBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IETHxBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IETHxBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHxBurnerTriggerBurn)
				if err := _IETHxBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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
func (_IETHxBurner *IETHxBurnerFilterer) ParseTriggerBurn(log types.Log) (*IETHxBurnerTriggerBurn, error) {
	event := new(IETHxBurnerTriggerBurn)
	if err := _IETHxBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IETHxBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the IETHxBurner contract.
type IETHxBurnerTriggerWithdrawalIterator struct {
	Event *IETHxBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

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
func (it *IETHxBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHxBurnerTriggerWithdrawal)
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
		it.Event = new(IETHxBurnerTriggerWithdrawal)
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
func (it *IETHxBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHxBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHxBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the IETHxBurner contract.
type IETHxBurnerTriggerWithdrawal struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_IETHxBurner *IETHxBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*IETHxBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IETHxBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &IETHxBurnerTriggerWithdrawalIterator{contract: _IETHxBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0x262b8826d3cce07380cc79eea8c1390c2680f3cc26276b7aecaed4868164cb45.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256 requestId)
func (_IETHxBurner *IETHxBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *IETHxBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IETHxBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHxBurnerTriggerWithdrawal)
				if err := _IETHxBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
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
func (_IETHxBurner *IETHxBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*IETHxBurnerTriggerWithdrawal, error) {
	event := new(IETHxBurnerTriggerWithdrawal)
	if err := _IETHxBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
