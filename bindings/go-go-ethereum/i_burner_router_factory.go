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

// IBurnerRouterInitParams is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterInitParams struct {
	Owner                    common.Address
	Collateral               common.Address
	Delay                    *big.Int
	GlobalReceiver           common.Address
	NetworkReceivers         []IBurnerRouterNetworkReceiver
	OperatorNetworkReceivers []IBurnerRouterOperatorNetworkReceiver
}

// IBurnerRouterNetworkReceiver is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterNetworkReceiver struct {
	Network  common.Address
	Receiver common.Address
}

// IBurnerRouterOperatorNetworkReceiver is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Receiver common.Address
}

// IBurnerRouterFactoryMetaData contains all meta data concerning the IBurnerRouterFactory contract.
var IBurnerRouterFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBurnerRouter.InitParams\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"globalReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"networkReceivers\",\"type\":\"tuple[]\",\"internalType\":\"structIBurnerRouter.NetworkReceiver[]\",\"components\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"operatorNetworkReceivers\",\"type\":\"tuple[]\",\"internalType\":\"structIBurnerRouter.OperatorNetworkReceiver[]\",\"components\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"entity\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEntity\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEntities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddEntity\",\"inputs\":[{\"name\":\"entity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EntityNotExist\",\"inputs\":[]}]",
}

// IBurnerRouterFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IBurnerRouterFactoryMetaData.ABI instead.
var IBurnerRouterFactoryABI = IBurnerRouterFactoryMetaData.ABI

// IBurnerRouterFactory is an auto generated Go binding around an Ethereum contract.
type IBurnerRouterFactory struct {
	IBurnerRouterFactoryCaller     // Read-only binding to the contract
	IBurnerRouterFactoryTransactor // Write-only binding to the contract
	IBurnerRouterFactoryFilterer   // Log filterer for contract events
}

// IBurnerRouterFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBurnerRouterFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBurnerRouterFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBurnerRouterFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBurnerRouterFactorySession struct {
	Contract     *IBurnerRouterFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBurnerRouterFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBurnerRouterFactoryCallerSession struct {
	Contract *IBurnerRouterFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IBurnerRouterFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBurnerRouterFactoryTransactorSession struct {
	Contract     *IBurnerRouterFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IBurnerRouterFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBurnerRouterFactoryRaw struct {
	Contract *IBurnerRouterFactory // Generic contract binding to access the raw methods on
}

// IBurnerRouterFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBurnerRouterFactoryCallerRaw struct {
	Contract *IBurnerRouterFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IBurnerRouterFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBurnerRouterFactoryTransactorRaw struct {
	Contract *IBurnerRouterFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBurnerRouterFactory creates a new instance of IBurnerRouterFactory, bound to a specific deployed contract.
func NewIBurnerRouterFactory(address common.Address, backend bind.ContractBackend) (*IBurnerRouterFactory, error) {
	contract, err := bindIBurnerRouterFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFactory{IBurnerRouterFactoryCaller: IBurnerRouterFactoryCaller{contract: contract}, IBurnerRouterFactoryTransactor: IBurnerRouterFactoryTransactor{contract: contract}, IBurnerRouterFactoryFilterer: IBurnerRouterFactoryFilterer{contract: contract}}, nil
}

// NewIBurnerRouterFactoryCaller creates a new read-only instance of IBurnerRouterFactory, bound to a specific deployed contract.
func NewIBurnerRouterFactoryCaller(address common.Address, caller bind.ContractCaller) (*IBurnerRouterFactoryCaller, error) {
	contract, err := bindIBurnerRouterFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFactoryCaller{contract: contract}, nil
}

// NewIBurnerRouterFactoryTransactor creates a new write-only instance of IBurnerRouterFactory, bound to a specific deployed contract.
func NewIBurnerRouterFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IBurnerRouterFactoryTransactor, error) {
	contract, err := bindIBurnerRouterFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFactoryTransactor{contract: contract}, nil
}

// NewIBurnerRouterFactoryFilterer creates a new log filterer instance of IBurnerRouterFactory, bound to a specific deployed contract.
func NewIBurnerRouterFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IBurnerRouterFactoryFilterer, error) {
	contract, err := bindIBurnerRouterFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFactoryFilterer{contract: contract}, nil
}

// bindIBurnerRouterFactory binds a generic wrapper to an already deployed contract.
func bindIBurnerRouterFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBurnerRouterFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnerRouterFactory *IBurnerRouterFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnerRouterFactory.Contract.IBurnerRouterFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnerRouterFactory *IBurnerRouterFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.IBurnerRouterFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnerRouterFactory *IBurnerRouterFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.IBurnerRouterFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnerRouterFactory *IBurnerRouterFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnerRouterFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnerRouterFactory *IBurnerRouterFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnerRouterFactory *IBurnerRouterFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.contract.Transact(opts, method, params...)
}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCaller) Entity(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IBurnerRouterFactory.contract.Call(opts, &out, "entity", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactorySession) Entity(index *big.Int) (common.Address, error) {
	return _IBurnerRouterFactory.Contract.Entity(&_IBurnerRouterFactory.CallOpts, index)
}

// Entity is a free data retrieval call binding the contract method 0xb42ba2a2.
//
// Solidity: function entity(uint256 index) view returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCallerSession) Entity(index *big.Int) (common.Address, error) {
	return _IBurnerRouterFactory.Contract.Entity(&_IBurnerRouterFactory.CallOpts, index)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address account) view returns(bool)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCaller) IsEntity(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IBurnerRouterFactory.contract.Call(opts, &out, "isEntity", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address account) view returns(bool)
func (_IBurnerRouterFactory *IBurnerRouterFactorySession) IsEntity(account common.Address) (bool, error) {
	return _IBurnerRouterFactory.Contract.IsEntity(&_IBurnerRouterFactory.CallOpts, account)
}

// IsEntity is a free data retrieval call binding the contract method 0x14887c58.
//
// Solidity: function isEntity(address account) view returns(bool)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCallerSession) IsEntity(account common.Address) (bool, error) {
	return _IBurnerRouterFactory.Contract.IsEntity(&_IBurnerRouterFactory.CallOpts, account)
}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCaller) TotalEntities(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBurnerRouterFactory.contract.Call(opts, &out, "totalEntities")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_IBurnerRouterFactory *IBurnerRouterFactorySession) TotalEntities() (*big.Int, error) {
	return _IBurnerRouterFactory.Contract.TotalEntities(&_IBurnerRouterFactory.CallOpts)
}

// TotalEntities is a free data retrieval call binding the contract method 0x5cd8b15e.
//
// Solidity: function totalEntities() view returns(uint256)
func (_IBurnerRouterFactory *IBurnerRouterFactoryCallerSession) TotalEntities() (*big.Int, error) {
	return _IBurnerRouterFactory.Contract.TotalEntities(&_IBurnerRouterFactory.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x3c4a80c8.
//
// Solidity: function create((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactoryTransactor) Create(opts *bind.TransactOpts, params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _IBurnerRouterFactory.contract.Transact(opts, "create", params)
}

// Create is a paid mutator transaction binding the contract method 0x3c4a80c8.
//
// Solidity: function create((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactorySession) Create(params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.Create(&_IBurnerRouterFactory.TransactOpts, params)
}

// Create is a paid mutator transaction binding the contract method 0x3c4a80c8.
//
// Solidity: function create((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns(address)
func (_IBurnerRouterFactory *IBurnerRouterFactoryTransactorSession) Create(params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _IBurnerRouterFactory.Contract.Create(&_IBurnerRouterFactory.TransactOpts, params)
}

// IBurnerRouterFactoryAddEntityIterator is returned from FilterAddEntity and is used to iterate over the raw logs and unpacked data for AddEntity events raised by the IBurnerRouterFactory contract.
type IBurnerRouterFactoryAddEntityIterator struct {
	Event *IBurnerRouterFactoryAddEntity // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterFactoryAddEntityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterFactoryAddEntity)
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
		it.Event = new(IBurnerRouterFactoryAddEntity)
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
func (it *IBurnerRouterFactoryAddEntityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterFactoryAddEntityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterFactoryAddEntity represents a AddEntity event raised by the IBurnerRouterFactory contract.
type IBurnerRouterFactoryAddEntity struct {
	Entity common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddEntity is a free log retrieval operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_IBurnerRouterFactory *IBurnerRouterFactoryFilterer) FilterAddEntity(opts *bind.FilterOpts, entity []common.Address) (*IBurnerRouterFactoryAddEntityIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _IBurnerRouterFactory.contract.FilterLogs(opts, "AddEntity", entityRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFactoryAddEntityIterator{contract: _IBurnerRouterFactory.contract, event: "AddEntity", logs: logs, sub: sub}, nil
}

// WatchAddEntity is a free log subscription operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_IBurnerRouterFactory *IBurnerRouterFactoryFilterer) WatchAddEntity(opts *bind.WatchOpts, sink chan<- *IBurnerRouterFactoryAddEntity, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _IBurnerRouterFactory.contract.WatchLogs(opts, "AddEntity", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterFactoryAddEntity)
				if err := _IBurnerRouterFactory.contract.UnpackLog(event, "AddEntity", log); err != nil {
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

// ParseAddEntity is a log parse operation binding the contract event 0xb919910dcefbf753bfd926ab3b1d3f85d877190c3d01ba1bd585047b99b99f0b.
//
// Solidity: event AddEntity(address indexed entity)
func (_IBurnerRouterFactory *IBurnerRouterFactoryFilterer) ParseAddEntity(log types.Log) (*IBurnerRouterFactoryAddEntity, error) {
	event := new(IBurnerRouterFactoryAddEntity)
	if err := _IBurnerRouterFactory.contract.UnpackLog(event, "AddEntity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
