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

// IWstETHMetaData contains all meta data concerning the IWstETH contract.
var IWstETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getStETHByWstETH\",\"inputs\":[{\"name\":\"_wstETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwrap\",\"inputs\":[{\"name\":\"_wstETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IWstETHABI is the input ABI used to generate the binding from.
// Deprecated: Use IWstETHMetaData.ABI instead.
var IWstETHABI = IWstETHMetaData.ABI

// IWstETH is an auto generated Go binding around an Ethereum contract.
type IWstETH struct {
	IWstETHCaller     // Read-only binding to the contract
	IWstETHTransactor // Write-only binding to the contract
	IWstETHFilterer   // Log filterer for contract events
}

// IWstETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWstETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWstETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWstETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWstETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWstETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWstETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWstETHSession struct {
	Contract     *IWstETH          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWstETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWstETHCallerSession struct {
	Contract *IWstETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IWstETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWstETHTransactorSession struct {
	Contract     *IWstETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IWstETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWstETHRaw struct {
	Contract *IWstETH // Generic contract binding to access the raw methods on
}

// IWstETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWstETHCallerRaw struct {
	Contract *IWstETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWstETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWstETHTransactorRaw struct {
	Contract *IWstETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWstETH creates a new instance of IWstETH, bound to a specific deployed contract.
func NewIWstETH(address common.Address, backend bind.ContractBackend) (*IWstETH, error) {
	contract, err := bindIWstETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWstETH{IWstETHCaller: IWstETHCaller{contract: contract}, IWstETHTransactor: IWstETHTransactor{contract: contract}, IWstETHFilterer: IWstETHFilterer{contract: contract}}, nil
}

// NewIWstETHCaller creates a new read-only instance of IWstETH, bound to a specific deployed contract.
func NewIWstETHCaller(address common.Address, caller bind.ContractCaller) (*IWstETHCaller, error) {
	contract, err := bindIWstETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWstETHCaller{contract: contract}, nil
}

// NewIWstETHTransactor creates a new write-only instance of IWstETH, bound to a specific deployed contract.
func NewIWstETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWstETHTransactor, error) {
	contract, err := bindIWstETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWstETHTransactor{contract: contract}, nil
}

// NewIWstETHFilterer creates a new log filterer instance of IWstETH, bound to a specific deployed contract.
func NewIWstETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWstETHFilterer, error) {
	contract, err := bindIWstETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWstETHFilterer{contract: contract}, nil
}

// bindIWstETH binds a generic wrapper to an already deployed contract.
func bindIWstETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWstETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWstETH *IWstETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWstETH.Contract.IWstETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWstETH *IWstETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWstETH.Contract.IWstETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWstETH *IWstETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWstETH.Contract.IWstETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWstETH *IWstETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWstETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWstETH *IWstETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWstETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWstETH *IWstETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWstETH.Contract.contract.Transact(opts, method, params...)
}

// GetStETHByWstETH is a paid mutator transaction binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHTransactor) GetStETHByWstETH(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.contract.Transact(opts, "getStETHByWstETH", _wstETHAmount)
}

// GetStETHByWstETH is a paid mutator transaction binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.Contract.GetStETHByWstETH(&_IWstETH.TransactOpts, _wstETHAmount)
}

// GetStETHByWstETH is a paid mutator transaction binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHTransactorSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.Contract.GetStETHByWstETH(&_IWstETH.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHTransactor) Unwrap(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.contract.Transact(opts, "unwrap", _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.Contract.Unwrap(&_IWstETH.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_IWstETH *IWstETHTransactorSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _IWstETH.Contract.Unwrap(&_IWstETH.TransactOpts, _wstETHAmount)
}
