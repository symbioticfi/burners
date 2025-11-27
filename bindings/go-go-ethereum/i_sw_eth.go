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

// ISwETHMetaData contains all meta data concerning the ISwETH contract.
var ISwETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// ISwETHABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwETHMetaData.ABI instead.
var ISwETHABI = ISwETHMetaData.ABI

// ISwETH is an auto generated Go binding around an Ethereum contract.
type ISwETH struct {
	ISwETHCaller     // Read-only binding to the contract
	ISwETHTransactor // Write-only binding to the contract
	ISwETHFilterer   // Log filterer for contract events
}

// ISwETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwETHSession struct {
	Contract     *ISwETH           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwETHCallerSession struct {
	Contract *ISwETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISwETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwETHTransactorSession struct {
	Contract     *ISwETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwETHRaw struct {
	Contract *ISwETH // Generic contract binding to access the raw methods on
}

// ISwETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwETHCallerRaw struct {
	Contract *ISwETHCaller // Generic read-only contract binding to access the raw methods on
}

// ISwETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwETHTransactorRaw struct {
	Contract *ISwETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwETH creates a new instance of ISwETH, bound to a specific deployed contract.
func NewISwETH(address common.Address, backend bind.ContractBackend) (*ISwETH, error) {
	contract, err := bindISwETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwETH{ISwETHCaller: ISwETHCaller{contract: contract}, ISwETHTransactor: ISwETHTransactor{contract: contract}, ISwETHFilterer: ISwETHFilterer{contract: contract}}, nil
}

// NewISwETHCaller creates a new read-only instance of ISwETH, bound to a specific deployed contract.
func NewISwETHCaller(address common.Address, caller bind.ContractCaller) (*ISwETHCaller, error) {
	contract, err := bindISwETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwETHCaller{contract: contract}, nil
}

// NewISwETHTransactor creates a new write-only instance of ISwETH, bound to a specific deployed contract.
func NewISwETHTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwETHTransactor, error) {
	contract, err := bindISwETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwETHTransactor{contract: contract}, nil
}

// NewISwETHFilterer creates a new log filterer instance of ISwETH, bound to a specific deployed contract.
func NewISwETHFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwETHFilterer, error) {
	contract, err := bindISwETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwETHFilterer{contract: contract}, nil
}

// bindISwETH binds a generic wrapper to an already deployed contract.
func bindISwETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwETH *ISwETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwETH.Contract.ISwETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwETH *ISwETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwETH.Contract.ISwETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwETH *ISwETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwETH.Contract.ISwETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwETH *ISwETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwETH *ISwETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwETH *ISwETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwETH.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ISwETH *ISwETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ISwETH *ISwETHSession) Deposit() (*types.Transaction, error) {
	return _ISwETH.Contract.Deposit(&_ISwETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ISwETH *ISwETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _ISwETH.Contract.Deposit(&_ISwETH.TransactOpts)
}
