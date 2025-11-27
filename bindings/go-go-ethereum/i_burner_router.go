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

// IBurnerRouterMetaData contains all meta data concerning the IBurnerRouter contract.
var IBurnerRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptGlobalReceiver\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOperatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"networkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onSlash\",\"inputs\":[{\"name\":\"subnetwork\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingGlobalReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOperatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerTransfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AcceptDelay\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcceptGlobalReceiver\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcceptNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcceptOperatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetGlobalReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetOperatorNetworkReceiver\",\"inputs\":[{\"name\":\"network\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerTransfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateNetworkReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOperatorNetworkReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReceiverSetEpochsDelay\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotReady\",\"inputs\":[]}]",
}

// IBurnerRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IBurnerRouterMetaData.ABI instead.
var IBurnerRouterABI = IBurnerRouterMetaData.ABI

// IBurnerRouter is an auto generated Go binding around an Ethereum contract.
type IBurnerRouter struct {
	IBurnerRouterCaller     // Read-only binding to the contract
	IBurnerRouterTransactor // Write-only binding to the contract
	IBurnerRouterFilterer   // Log filterer for contract events
}

// IBurnerRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBurnerRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBurnerRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBurnerRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnerRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBurnerRouterSession struct {
	Contract     *IBurnerRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBurnerRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBurnerRouterCallerSession struct {
	Contract *IBurnerRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBurnerRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBurnerRouterTransactorSession struct {
	Contract     *IBurnerRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBurnerRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBurnerRouterRaw struct {
	Contract *IBurnerRouter // Generic contract binding to access the raw methods on
}

// IBurnerRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBurnerRouterCallerRaw struct {
	Contract *IBurnerRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IBurnerRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBurnerRouterTransactorRaw struct {
	Contract *IBurnerRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBurnerRouter creates a new instance of IBurnerRouter, bound to a specific deployed contract.
func NewIBurnerRouter(address common.Address, backend bind.ContractBackend) (*IBurnerRouter, error) {
	contract, err := bindIBurnerRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouter{IBurnerRouterCaller: IBurnerRouterCaller{contract: contract}, IBurnerRouterTransactor: IBurnerRouterTransactor{contract: contract}, IBurnerRouterFilterer: IBurnerRouterFilterer{contract: contract}}, nil
}

// NewIBurnerRouterCaller creates a new read-only instance of IBurnerRouter, bound to a specific deployed contract.
func NewIBurnerRouterCaller(address common.Address, caller bind.ContractCaller) (*IBurnerRouterCaller, error) {
	contract, err := bindIBurnerRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterCaller{contract: contract}, nil
}

// NewIBurnerRouterTransactor creates a new write-only instance of IBurnerRouter, bound to a specific deployed contract.
func NewIBurnerRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IBurnerRouterTransactor, error) {
	contract, err := bindIBurnerRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterTransactor{contract: contract}, nil
}

// NewIBurnerRouterFilterer creates a new log filterer instance of IBurnerRouter, bound to a specific deployed contract.
func NewIBurnerRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IBurnerRouterFilterer, error) {
	contract, err := bindIBurnerRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterFilterer{contract: contract}, nil
}

// bindIBurnerRouter binds a generic wrapper to an already deployed contract.
func bindIBurnerRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBurnerRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnerRouter *IBurnerRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnerRouter.Contract.IBurnerRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnerRouter *IBurnerRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.IBurnerRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnerRouter *IBurnerRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.IBurnerRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnerRouter *IBurnerRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnerRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnerRouter *IBurnerRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnerRouter *IBurnerRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256)
func (_IBurnerRouter *IBurnerRouterCaller) BalanceOf(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "balanceOf", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256)
func (_IBurnerRouter *IBurnerRouterSession) BalanceOf(receiver common.Address) (*big.Int, error) {
	return _IBurnerRouter.Contract.BalanceOf(&_IBurnerRouter.CallOpts, receiver)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256)
func (_IBurnerRouter *IBurnerRouterCallerSession) BalanceOf(receiver common.Address) (*big.Int, error) {
	return _IBurnerRouter.Contract.BalanceOf(&_IBurnerRouter.CallOpts, receiver)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_IBurnerRouter *IBurnerRouterCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_IBurnerRouter *IBurnerRouterSession) Collateral() (common.Address, error) {
	return _IBurnerRouter.Contract.Collateral(&_IBurnerRouter.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_IBurnerRouter *IBurnerRouterCallerSession) Collateral() (common.Address, error) {
	return _IBurnerRouter.Contract.Collateral(&_IBurnerRouter.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48)
func (_IBurnerRouter *IBurnerRouterCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48)
func (_IBurnerRouter *IBurnerRouterSession) Delay() (*big.Int, error) {
	return _IBurnerRouter.Contract.Delay(&_IBurnerRouter.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48)
func (_IBurnerRouter *IBurnerRouterCallerSession) Delay() (*big.Int, error) {
	return _IBurnerRouter.Contract.Delay(&_IBurnerRouter.CallOpts)
}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address)
func (_IBurnerRouter *IBurnerRouterCaller) GlobalReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "globalReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address)
func (_IBurnerRouter *IBurnerRouterSession) GlobalReceiver() (common.Address, error) {
	return _IBurnerRouter.Contract.GlobalReceiver(&_IBurnerRouter.CallOpts)
}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address)
func (_IBurnerRouter *IBurnerRouterCallerSession) GlobalReceiver() (common.Address, error) {
	return _IBurnerRouter.Contract.GlobalReceiver(&_IBurnerRouter.CallOpts)
}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_IBurnerRouter *IBurnerRouterCaller) LastBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "lastBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_IBurnerRouter *IBurnerRouterSession) LastBalance() (*big.Int, error) {
	return _IBurnerRouter.Contract.LastBalance(&_IBurnerRouter.CallOpts)
}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_IBurnerRouter *IBurnerRouterCallerSession) LastBalance() (*big.Int, error) {
	return _IBurnerRouter.Contract.LastBalance(&_IBurnerRouter.CallOpts)
}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address)
func (_IBurnerRouter *IBurnerRouterCaller) NetworkReceiver(opts *bind.CallOpts, network common.Address) (common.Address, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "networkReceiver", network)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address)
func (_IBurnerRouter *IBurnerRouterSession) NetworkReceiver(network common.Address) (common.Address, error) {
	return _IBurnerRouter.Contract.NetworkReceiver(&_IBurnerRouter.CallOpts, network)
}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address)
func (_IBurnerRouter *IBurnerRouterCallerSession) NetworkReceiver(network common.Address) (common.Address, error) {
	return _IBurnerRouter.Contract.NetworkReceiver(&_IBurnerRouter.CallOpts, network)
}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address)
func (_IBurnerRouter *IBurnerRouterCaller) OperatorNetworkReceiver(opts *bind.CallOpts, network common.Address, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "operatorNetworkReceiver", network, operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address)
func (_IBurnerRouter *IBurnerRouterSession) OperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, error) {
	return _IBurnerRouter.Contract.OperatorNetworkReceiver(&_IBurnerRouter.CallOpts, network, operator)
}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address)
func (_IBurnerRouter *IBurnerRouterCallerSession) OperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, error) {
	return _IBurnerRouter.Contract.OperatorNetworkReceiver(&_IBurnerRouter.CallOpts, network, operator)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48, uint48)
func (_IBurnerRouter *IBurnerRouterCaller) PendingDelay(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "pendingDelay")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48, uint48)
func (_IBurnerRouter *IBurnerRouterSession) PendingDelay() (*big.Int, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingDelay(&_IBurnerRouter.CallOpts)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48, uint48)
func (_IBurnerRouter *IBurnerRouterCallerSession) PendingDelay() (*big.Int, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingDelay(&_IBurnerRouter.CallOpts)
}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCaller) PendingGlobalReceiver(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "pendingGlobalReceiver")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterSession) PendingGlobalReceiver() (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingGlobalReceiver(&_IBurnerRouter.CallOpts)
}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCallerSession) PendingGlobalReceiver() (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingGlobalReceiver(&_IBurnerRouter.CallOpts)
}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCaller) PendingNetworkReceiver(opts *bind.CallOpts, network common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "pendingNetworkReceiver", network)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterSession) PendingNetworkReceiver(network common.Address) (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingNetworkReceiver(&_IBurnerRouter.CallOpts, network)
}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCallerSession) PendingNetworkReceiver(network common.Address) (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingNetworkReceiver(&_IBurnerRouter.CallOpts, network)
}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCaller) PendingOperatorNetworkReceiver(opts *bind.CallOpts, network common.Address, operator common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IBurnerRouter.contract.Call(opts, &out, "pendingOperatorNetworkReceiver", network, operator)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterSession) PendingOperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingOperatorNetworkReceiver(&_IBurnerRouter.CallOpts, network, operator)
}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address, uint48)
func (_IBurnerRouter *IBurnerRouterCallerSession) PendingOperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, *big.Int, error) {
	return _IBurnerRouter.Contract.PendingOperatorNetworkReceiver(&_IBurnerRouter.CallOpts, network, operator)
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_IBurnerRouter *IBurnerRouterTransactor) AcceptDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "acceptDelay")
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_IBurnerRouter *IBurnerRouterSession) AcceptDelay() (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptDelay(&_IBurnerRouter.TransactOpts)
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) AcceptDelay() (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptDelay(&_IBurnerRouter.TransactOpts)
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_IBurnerRouter *IBurnerRouterTransactor) AcceptGlobalReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "acceptGlobalReceiver")
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_IBurnerRouter *IBurnerRouterSession) AcceptGlobalReceiver() (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptGlobalReceiver(&_IBurnerRouter.TransactOpts)
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) AcceptGlobalReceiver() (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptGlobalReceiver(&_IBurnerRouter.TransactOpts)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) AcceptNetworkReceiver(opts *bind.TransactOpts, network common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "acceptNetworkReceiver", network)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_IBurnerRouter *IBurnerRouterSession) AcceptNetworkReceiver(network common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptNetworkReceiver(&_IBurnerRouter.TransactOpts, network)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) AcceptNetworkReceiver(network common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptNetworkReceiver(&_IBurnerRouter.TransactOpts, network)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) AcceptOperatorNetworkReceiver(opts *bind.TransactOpts, network common.Address, operator common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "acceptOperatorNetworkReceiver", network, operator)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_IBurnerRouter *IBurnerRouterSession) AcceptOperatorNetworkReceiver(network common.Address, operator common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptOperatorNetworkReceiver(&_IBurnerRouter.TransactOpts, network, operator)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) AcceptOperatorNetworkReceiver(network common.Address, operator common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.AcceptOperatorNetworkReceiver(&_IBurnerRouter.TransactOpts, network, operator)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) OnSlash(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "onSlash", subnetwork, operator, amount, captureTimestamp)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns()
func (_IBurnerRouter *IBurnerRouterSession) OnSlash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.OnSlash(&_IBurnerRouter.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) OnSlash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.OnSlash(&_IBurnerRouter.TransactOpts, subnetwork, operator, amount, captureTimestamp)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) SetDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "setDelay", newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_IBurnerRouter *IBurnerRouterSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetDelay(&_IBurnerRouter.TransactOpts, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetDelay(&_IBurnerRouter.TransactOpts, newDelay)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) SetGlobalReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "setGlobalReceiver", receiver)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_IBurnerRouter *IBurnerRouterSession) SetGlobalReceiver(receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetGlobalReceiver(&_IBurnerRouter.TransactOpts, receiver)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) SetGlobalReceiver(receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetGlobalReceiver(&_IBurnerRouter.TransactOpts, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) SetNetworkReceiver(opts *bind.TransactOpts, network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "setNetworkReceiver", network, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterSession) SetNetworkReceiver(network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetNetworkReceiver(&_IBurnerRouter.TransactOpts, network, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) SetNetworkReceiver(network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetNetworkReceiver(&_IBurnerRouter.TransactOpts, network, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactor) SetOperatorNetworkReceiver(opts *bind.TransactOpts, network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "setOperatorNetworkReceiver", network, operator, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterSession) SetOperatorNetworkReceiver(network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetOperatorNetworkReceiver(&_IBurnerRouter.TransactOpts, network, operator, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_IBurnerRouter *IBurnerRouterTransactorSession) SetOperatorNetworkReceiver(network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.SetOperatorNetworkReceiver(&_IBurnerRouter.TransactOpts, network, operator, receiver)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_IBurnerRouter *IBurnerRouterTransactor) TriggerTransfer(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.contract.Transact(opts, "triggerTransfer", receiver)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_IBurnerRouter *IBurnerRouterSession) TriggerTransfer(receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.TriggerTransfer(&_IBurnerRouter.TransactOpts, receiver)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_IBurnerRouter *IBurnerRouterTransactorSession) TriggerTransfer(receiver common.Address) (*types.Transaction, error) {
	return _IBurnerRouter.Contract.TriggerTransfer(&_IBurnerRouter.TransactOpts, receiver)
}

// IBurnerRouterAcceptDelayIterator is returned from FilterAcceptDelay and is used to iterate over the raw logs and unpacked data for AcceptDelay events raised by the IBurnerRouter contract.
type IBurnerRouterAcceptDelayIterator struct {
	Event *IBurnerRouterAcceptDelay // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterAcceptDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterAcceptDelay)
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
		it.Event = new(IBurnerRouterAcceptDelay)
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
func (it *IBurnerRouterAcceptDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterAcceptDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterAcceptDelay represents a AcceptDelay event raised by the IBurnerRouter contract.
type IBurnerRouterAcceptDelay struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAcceptDelay is a free log retrieval operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_IBurnerRouter *IBurnerRouterFilterer) FilterAcceptDelay(opts *bind.FilterOpts) (*IBurnerRouterAcceptDelayIterator, error) {

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "AcceptDelay")
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterAcceptDelayIterator{contract: _IBurnerRouter.contract, event: "AcceptDelay", logs: logs, sub: sub}, nil
}

// WatchAcceptDelay is a free log subscription operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_IBurnerRouter *IBurnerRouterFilterer) WatchAcceptDelay(opts *bind.WatchOpts, sink chan<- *IBurnerRouterAcceptDelay) (event.Subscription, error) {

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "AcceptDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterAcceptDelay)
				if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptDelay", log); err != nil {
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

// ParseAcceptDelay is a log parse operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_IBurnerRouter *IBurnerRouterFilterer) ParseAcceptDelay(log types.Log) (*IBurnerRouterAcceptDelay, error) {
	event := new(IBurnerRouterAcceptDelay)
	if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterAcceptGlobalReceiverIterator is returned from FilterAcceptGlobalReceiver and is used to iterate over the raw logs and unpacked data for AcceptGlobalReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterAcceptGlobalReceiverIterator struct {
	Event *IBurnerRouterAcceptGlobalReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterAcceptGlobalReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterAcceptGlobalReceiver)
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
		it.Event = new(IBurnerRouterAcceptGlobalReceiver)
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
func (it *IBurnerRouterAcceptGlobalReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterAcceptGlobalReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterAcceptGlobalReceiver represents a AcceptGlobalReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterAcceptGlobalReceiver struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAcceptGlobalReceiver is a free log retrieval operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_IBurnerRouter *IBurnerRouterFilterer) FilterAcceptGlobalReceiver(opts *bind.FilterOpts) (*IBurnerRouterAcceptGlobalReceiverIterator, error) {

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "AcceptGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterAcceptGlobalReceiverIterator{contract: _IBurnerRouter.contract, event: "AcceptGlobalReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptGlobalReceiver is a free log subscription operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_IBurnerRouter *IBurnerRouterFilterer) WatchAcceptGlobalReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterAcceptGlobalReceiver) (event.Subscription, error) {

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "AcceptGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterAcceptGlobalReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptGlobalReceiver", log); err != nil {
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

// ParseAcceptGlobalReceiver is a log parse operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_IBurnerRouter *IBurnerRouterFilterer) ParseAcceptGlobalReceiver(log types.Log) (*IBurnerRouterAcceptGlobalReceiver, error) {
	event := new(IBurnerRouterAcceptGlobalReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptGlobalReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterAcceptNetworkReceiverIterator is returned from FilterAcceptNetworkReceiver and is used to iterate over the raw logs and unpacked data for AcceptNetworkReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterAcceptNetworkReceiverIterator struct {
	Event *IBurnerRouterAcceptNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterAcceptNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterAcceptNetworkReceiver)
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
		it.Event = new(IBurnerRouterAcceptNetworkReceiver)
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
func (it *IBurnerRouterAcceptNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterAcceptNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterAcceptNetworkReceiver represents a AcceptNetworkReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterAcceptNetworkReceiver struct {
	Network common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAcceptNetworkReceiver is a free log retrieval operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterAcceptNetworkReceiver(opts *bind.FilterOpts, network []common.Address) (*IBurnerRouterAcceptNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "AcceptNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterAcceptNetworkReceiverIterator{contract: _IBurnerRouter.contract, event: "AcceptNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptNetworkReceiver is a free log subscription operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchAcceptNetworkReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterAcceptNetworkReceiver, network []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "AcceptNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterAcceptNetworkReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptNetworkReceiver", log); err != nil {
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

// ParseAcceptNetworkReceiver is a log parse operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseAcceptNetworkReceiver(log types.Log) (*IBurnerRouterAcceptNetworkReceiver, error) {
	event := new(IBurnerRouterAcceptNetworkReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterAcceptOperatorNetworkReceiverIterator is returned from FilterAcceptOperatorNetworkReceiver and is used to iterate over the raw logs and unpacked data for AcceptOperatorNetworkReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterAcceptOperatorNetworkReceiverIterator struct {
	Event *IBurnerRouterAcceptOperatorNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterAcceptOperatorNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterAcceptOperatorNetworkReceiver)
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
		it.Event = new(IBurnerRouterAcceptOperatorNetworkReceiver)
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
func (it *IBurnerRouterAcceptOperatorNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterAcceptOperatorNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterAcceptOperatorNetworkReceiver represents a AcceptOperatorNetworkReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterAcceptOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptOperatorNetworkReceiver is a free log retrieval operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterAcceptOperatorNetworkReceiver(opts *bind.FilterOpts, network []common.Address, operator []common.Address) (*IBurnerRouterAcceptOperatorNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "AcceptOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterAcceptOperatorNetworkReceiverIterator{contract: _IBurnerRouter.contract, event: "AcceptOperatorNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptOperatorNetworkReceiver is a free log subscription operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchAcceptOperatorNetworkReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterAcceptOperatorNetworkReceiver, network []common.Address, operator []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "AcceptOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterAcceptOperatorNetworkReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptOperatorNetworkReceiver", log); err != nil {
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

// ParseAcceptOperatorNetworkReceiver is a log parse operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseAcceptOperatorNetworkReceiver(log types.Log) (*IBurnerRouterAcceptOperatorNetworkReceiver, error) {
	event := new(IBurnerRouterAcceptOperatorNetworkReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "AcceptOperatorNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterSetDelayIterator is returned from FilterSetDelay and is used to iterate over the raw logs and unpacked data for SetDelay events raised by the IBurnerRouter contract.
type IBurnerRouterSetDelayIterator struct {
	Event *IBurnerRouterSetDelay // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterSetDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterSetDelay)
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
		it.Event = new(IBurnerRouterSetDelay)
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
func (it *IBurnerRouterSetDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterSetDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterSetDelay represents a SetDelay event raised by the IBurnerRouter contract.
type IBurnerRouterSetDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDelay is a free log retrieval operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterSetDelay(opts *bind.FilterOpts) (*IBurnerRouterSetDelayIterator, error) {

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "SetDelay")
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterSetDelayIterator{contract: _IBurnerRouter.contract, event: "SetDelay", logs: logs, sub: sub}, nil
}

// WatchSetDelay is a free log subscription operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchSetDelay(opts *bind.WatchOpts, sink chan<- *IBurnerRouterSetDelay) (event.Subscription, error) {

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "SetDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterSetDelay)
				if err := _IBurnerRouter.contract.UnpackLog(event, "SetDelay", log); err != nil {
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

// ParseSetDelay is a log parse operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseSetDelay(log types.Log) (*IBurnerRouterSetDelay, error) {
	event := new(IBurnerRouterSetDelay)
	if err := _IBurnerRouter.contract.UnpackLog(event, "SetDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterSetGlobalReceiverIterator is returned from FilterSetGlobalReceiver and is used to iterate over the raw logs and unpacked data for SetGlobalReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterSetGlobalReceiverIterator struct {
	Event *IBurnerRouterSetGlobalReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterSetGlobalReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterSetGlobalReceiver)
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
		it.Event = new(IBurnerRouterSetGlobalReceiver)
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
func (it *IBurnerRouterSetGlobalReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterSetGlobalReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterSetGlobalReceiver represents a SetGlobalReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterSetGlobalReceiver struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalReceiver is a free log retrieval operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterSetGlobalReceiver(opts *bind.FilterOpts) (*IBurnerRouterSetGlobalReceiverIterator, error) {

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "SetGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterSetGlobalReceiverIterator{contract: _IBurnerRouter.contract, event: "SetGlobalReceiver", logs: logs, sub: sub}, nil
}

// WatchSetGlobalReceiver is a free log subscription operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchSetGlobalReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterSetGlobalReceiver) (event.Subscription, error) {

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "SetGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterSetGlobalReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "SetGlobalReceiver", log); err != nil {
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

// ParseSetGlobalReceiver is a log parse operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseSetGlobalReceiver(log types.Log) (*IBurnerRouterSetGlobalReceiver, error) {
	event := new(IBurnerRouterSetGlobalReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "SetGlobalReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterSetNetworkReceiverIterator is returned from FilterSetNetworkReceiver and is used to iterate over the raw logs and unpacked data for SetNetworkReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterSetNetworkReceiverIterator struct {
	Event *IBurnerRouterSetNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterSetNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterSetNetworkReceiver)
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
		it.Event = new(IBurnerRouterSetNetworkReceiver)
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
func (it *IBurnerRouterSetNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterSetNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterSetNetworkReceiver represents a SetNetworkReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterSetNetworkReceiver struct {
	Network  common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetNetworkReceiver is a free log retrieval operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterSetNetworkReceiver(opts *bind.FilterOpts, network []common.Address) (*IBurnerRouterSetNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "SetNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterSetNetworkReceiverIterator{contract: _IBurnerRouter.contract, event: "SetNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchSetNetworkReceiver is a free log subscription operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchSetNetworkReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterSetNetworkReceiver, network []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "SetNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterSetNetworkReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "SetNetworkReceiver", log); err != nil {
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

// ParseSetNetworkReceiver is a log parse operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseSetNetworkReceiver(log types.Log) (*IBurnerRouterSetNetworkReceiver, error) {
	event := new(IBurnerRouterSetNetworkReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "SetNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterSetOperatorNetworkReceiverIterator is returned from FilterSetOperatorNetworkReceiver and is used to iterate over the raw logs and unpacked data for SetOperatorNetworkReceiver events raised by the IBurnerRouter contract.
type IBurnerRouterSetOperatorNetworkReceiverIterator struct {
	Event *IBurnerRouterSetOperatorNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterSetOperatorNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterSetOperatorNetworkReceiver)
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
		it.Event = new(IBurnerRouterSetOperatorNetworkReceiver)
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
func (it *IBurnerRouterSetOperatorNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterSetOperatorNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterSetOperatorNetworkReceiver represents a SetOperatorNetworkReceiver event raised by the IBurnerRouter contract.
type IBurnerRouterSetOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperatorNetworkReceiver is a free log retrieval operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterSetOperatorNetworkReceiver(opts *bind.FilterOpts, network []common.Address, operator []common.Address) (*IBurnerRouterSetOperatorNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "SetOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterSetOperatorNetworkReceiverIterator{contract: _IBurnerRouter.contract, event: "SetOperatorNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchSetOperatorNetworkReceiver is a free log subscription operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchSetOperatorNetworkReceiver(opts *bind.WatchOpts, sink chan<- *IBurnerRouterSetOperatorNetworkReceiver, network []common.Address, operator []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "SetOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterSetOperatorNetworkReceiver)
				if err := _IBurnerRouter.contract.UnpackLog(event, "SetOperatorNetworkReceiver", log); err != nil {
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

// ParseSetOperatorNetworkReceiver is a log parse operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseSetOperatorNetworkReceiver(log types.Log) (*IBurnerRouterSetOperatorNetworkReceiver, error) {
	event := new(IBurnerRouterSetOperatorNetworkReceiver)
	if err := _IBurnerRouter.contract.UnpackLog(event, "SetOperatorNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnerRouterTriggerTransferIterator is returned from FilterTriggerTransfer and is used to iterate over the raw logs and unpacked data for TriggerTransfer events raised by the IBurnerRouter contract.
type IBurnerRouterTriggerTransferIterator struct {
	Event *IBurnerRouterTriggerTransfer // Event containing the contract specifics and raw log

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
func (it *IBurnerRouterTriggerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnerRouterTriggerTransfer)
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
		it.Event = new(IBurnerRouterTriggerTransfer)
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
func (it *IBurnerRouterTriggerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnerRouterTriggerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnerRouterTriggerTransfer represents a TriggerTransfer event raised by the IBurnerRouter contract.
type IBurnerRouterTriggerTransfer struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTriggerTransfer is a free log retrieval operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_IBurnerRouter *IBurnerRouterFilterer) FilterTriggerTransfer(opts *bind.FilterOpts, receiver []common.Address) (*IBurnerRouterTriggerTransferIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IBurnerRouter.contract.FilterLogs(opts, "TriggerTransfer", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IBurnerRouterTriggerTransferIterator{contract: _IBurnerRouter.contract, event: "TriggerTransfer", logs: logs, sub: sub}, nil
}

// WatchTriggerTransfer is a free log subscription operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_IBurnerRouter *IBurnerRouterFilterer) WatchTriggerTransfer(opts *bind.WatchOpts, sink chan<- *IBurnerRouterTriggerTransfer, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IBurnerRouter.contract.WatchLogs(opts, "TriggerTransfer", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnerRouterTriggerTransfer)
				if err := _IBurnerRouter.contract.UnpackLog(event, "TriggerTransfer", log); err != nil {
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

// ParseTriggerTransfer is a log parse operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_IBurnerRouter *IBurnerRouterFilterer) ParseTriggerTransfer(log types.Log) (*IBurnerRouterTriggerTransfer, error) {
	event := new(IBurnerRouterTriggerTransfer)
	if err := _IBurnerRouter.contract.UnpackLog(event, "TriggerTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
