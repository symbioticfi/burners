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

// IrETHBurnerMetaData contains all meta data concerning the IrETHBurner contract.
var IrETHBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assetAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ethAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IrETHBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IrETHBurnerMetaData.ABI instead.
var IrETHBurnerABI = IrETHBurnerMetaData.ABI

// IrETHBurner is an auto generated Go binding around an Ethereum contract.
type IrETHBurner struct {
	IrETHBurnerCaller     // Read-only binding to the contract
	IrETHBurnerTransactor // Write-only binding to the contract
	IrETHBurnerFilterer   // Log filterer for contract events
}

// IrETHBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IrETHBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrETHBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IrETHBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrETHBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IrETHBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrETHBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IrETHBurnerSession struct {
	Contract     *IrETHBurner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IrETHBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IrETHBurnerCallerSession struct {
	Contract *IrETHBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IrETHBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IrETHBurnerTransactorSession struct {
	Contract     *IrETHBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IrETHBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IrETHBurnerRaw struct {
	Contract *IrETHBurner // Generic contract binding to access the raw methods on
}

// IrETHBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IrETHBurnerCallerRaw struct {
	Contract *IrETHBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IrETHBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IrETHBurnerTransactorRaw struct {
	Contract *IrETHBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIrETHBurner creates a new instance of IrETHBurner, bound to a specific deployed contract.
func NewIrETHBurner(address common.Address, backend bind.ContractBackend) (*IrETHBurner, error) {
	contract, err := bindIrETHBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IrETHBurner{IrETHBurnerCaller: IrETHBurnerCaller{contract: contract}, IrETHBurnerTransactor: IrETHBurnerTransactor{contract: contract}, IrETHBurnerFilterer: IrETHBurnerFilterer{contract: contract}}, nil
}

// NewIrETHBurnerCaller creates a new read-only instance of IrETHBurner, bound to a specific deployed contract.
func NewIrETHBurnerCaller(address common.Address, caller bind.ContractCaller) (*IrETHBurnerCaller, error) {
	contract, err := bindIrETHBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IrETHBurnerCaller{contract: contract}, nil
}

// NewIrETHBurnerTransactor creates a new write-only instance of IrETHBurner, bound to a specific deployed contract.
func NewIrETHBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IrETHBurnerTransactor, error) {
	contract, err := bindIrETHBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IrETHBurnerTransactor{contract: contract}, nil
}

// NewIrETHBurnerFilterer creates a new log filterer instance of IrETHBurner, bound to a specific deployed contract.
func NewIrETHBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IrETHBurnerFilterer, error) {
	contract, err := bindIrETHBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IrETHBurnerFilterer{contract: contract}, nil
}

// bindIrETHBurner binds a generic wrapper to an already deployed contract.
func bindIrETHBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IrETHBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IrETHBurner *IrETHBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IrETHBurner.Contract.IrETHBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IrETHBurner *IrETHBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IrETHBurner.Contract.IrETHBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IrETHBurner *IrETHBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IrETHBurner.Contract.IrETHBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IrETHBurner *IrETHBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IrETHBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IrETHBurner *IrETHBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IrETHBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IrETHBurner *IrETHBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IrETHBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IrETHBurner *IrETHBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IrETHBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IrETHBurner *IrETHBurnerSession) COLLATERAL() (common.Address, error) {
	return _IrETHBurner.Contract.COLLATERAL(&_IrETHBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IrETHBurner *IrETHBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IrETHBurner.Contract.COLLATERAL(&_IrETHBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 amount) returns()
func (_IrETHBurner *IrETHBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IrETHBurner.contract.Transact(opts, "triggerBurn", amount)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 amount) returns()
func (_IrETHBurner *IrETHBurnerSession) TriggerBurn(amount *big.Int) (*types.Transaction, error) {
	return _IrETHBurner.Contract.TriggerBurn(&_IrETHBurner.TransactOpts, amount)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 amount) returns()
func (_IrETHBurner *IrETHBurnerTransactorSession) TriggerBurn(amount *big.Int) (*types.Transaction, error) {
	return _IrETHBurner.Contract.TriggerBurn(&_IrETHBurner.TransactOpts, amount)
}

// IrETHBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IrETHBurner contract.
type IrETHBurnerTriggerBurnIterator struct {
	Event *IrETHBurnerTriggerBurn // Event containing the contract specifics and raw log

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
func (it *IrETHBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IrETHBurnerTriggerBurn)
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
		it.Event = new(IrETHBurnerTriggerBurn)
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
func (it *IrETHBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IrETHBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IrETHBurnerTriggerBurn represents a TriggerBurn event raised by the IrETHBurner contract.
type IrETHBurnerTriggerBurn struct {
	Caller      common.Address
	AssetAmount *big.Int
	EthAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0xf04af4aa419dbccb13024c7e3c652ec214ad43e79b4e23f36741bbe275808681.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 assetAmount, uint256 ethAmount)
func (_IrETHBurner *IrETHBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*IrETHBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IrETHBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &IrETHBurnerTriggerBurnIterator{contract: _IrETHBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0xf04af4aa419dbccb13024c7e3c652ec214ad43e79b4e23f36741bbe275808681.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 assetAmount, uint256 ethAmount)
func (_IrETHBurner *IrETHBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IrETHBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IrETHBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IrETHBurnerTriggerBurn)
				if err := _IrETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
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

// ParseTriggerBurn is a log parse operation binding the contract event 0xf04af4aa419dbccb13024c7e3c652ec214ad43e79b4e23f36741bbe275808681.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 assetAmount, uint256 ethAmount)
func (_IrETHBurner *IrETHBurnerFilterer) ParseTriggerBurn(log types.Log) (*IrETHBurnerTriggerBurn, error) {
	event := new(IrETHBurnerTriggerBurn)
	if err := _IrETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
