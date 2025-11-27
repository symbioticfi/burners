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

// IStakingMetaData contains all meta data concerning the IStaking contract.
var IStakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"claimUnstakeRequest\",\"inputs\":[{\"name\":\"unstakeRequestID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mETHToETH\",\"inputs\":[{\"name\":\"mETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumUnstakeBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeRequest\",\"inputs\":[{\"name\":\"methAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"minETHAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingMetaData.ABI instead.
var IStakingABI = IStakingMetaData.ABI

// IStaking is an auto generated Go binding around an Ethereum contract.
type IStaking struct {
	IStakingCaller     // Read-only binding to the contract
	IStakingTransactor // Write-only binding to the contract
	IStakingFilterer   // Log filterer for contract events
}

// IStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingSession struct {
	Contract     *IStaking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingCallerSession struct {
	Contract *IStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingTransactorSession struct {
	Contract     *IStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingRaw struct {
	Contract *IStaking // Generic contract binding to access the raw methods on
}

// IStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingCallerRaw struct {
	Contract *IStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingTransactorRaw struct {
	Contract *IStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaking creates a new instance of IStaking, bound to a specific deployed contract.
func NewIStaking(address common.Address, backend bind.ContractBackend) (*IStaking, error) {
	contract, err := bindIStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaking{IStakingCaller: IStakingCaller{contract: contract}, IStakingTransactor: IStakingTransactor{contract: contract}, IStakingFilterer: IStakingFilterer{contract: contract}}, nil
}

// NewIStakingCaller creates a new read-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingCaller(address common.Address, caller bind.ContractCaller) (*IStakingCaller, error) {
	contract, err := bindIStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingCaller{contract: contract}, nil
}

// NewIStakingTransactor creates a new write-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingTransactor, error) {
	contract, err := bindIStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingTransactor{contract: contract}, nil
}

// NewIStakingFilterer creates a new log filterer instance of IStaking, bound to a specific deployed contract.
func NewIStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingFilterer, error) {
	contract, err := bindIStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingFilterer{contract: contract}, nil
}

// bindIStaking binds a generic wrapper to an already deployed contract.
func bindIStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.IStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transact(opts, method, params...)
}

// METHToETH is a free data retrieval call binding the contract method 0x5890c11c.
//
// Solidity: function mETHToETH(uint256 mETHAmount) view returns(uint256)
func (_IStaking *IStakingCaller) METHToETH(opts *bind.CallOpts, mETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "mETHToETH", mETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// METHToETH is a free data retrieval call binding the contract method 0x5890c11c.
//
// Solidity: function mETHToETH(uint256 mETHAmount) view returns(uint256)
func (_IStaking *IStakingSession) METHToETH(mETHAmount *big.Int) (*big.Int, error) {
	return _IStaking.Contract.METHToETH(&_IStaking.CallOpts, mETHAmount)
}

// METHToETH is a free data retrieval call binding the contract method 0x5890c11c.
//
// Solidity: function mETHToETH(uint256 mETHAmount) view returns(uint256)
func (_IStaking *IStakingCallerSession) METHToETH(mETHAmount *big.Int) (*big.Int, error) {
	return _IStaking.Contract.METHToETH(&_IStaking.CallOpts, mETHAmount)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_IStaking *IStakingCaller) MinimumUnstakeBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "minimumUnstakeBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_IStaking *IStakingSession) MinimumUnstakeBound() (*big.Int, error) {
	return _IStaking.Contract.MinimumUnstakeBound(&_IStaking.CallOpts)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_IStaking *IStakingCallerSession) MinimumUnstakeBound() (*big.Int, error) {
	return _IStaking.Contract.MinimumUnstakeBound(&_IStaking.CallOpts)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x2bf67650.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID) returns()
func (_IStaking *IStakingTransactor) ClaimUnstakeRequest(opts *bind.TransactOpts, unstakeRequestID *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "claimUnstakeRequest", unstakeRequestID)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x2bf67650.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID) returns()
func (_IStaking *IStakingSession) ClaimUnstakeRequest(unstakeRequestID *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.ClaimUnstakeRequest(&_IStaking.TransactOpts, unstakeRequestID)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x2bf67650.
//
// Solidity: function claimUnstakeRequest(uint256 unstakeRequestID) returns()
func (_IStaking *IStakingTransactorSession) ClaimUnstakeRequest(unstakeRequestID *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.ClaimUnstakeRequest(&_IStaking.TransactOpts, unstakeRequestID)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x891ef43e.
//
// Solidity: function unstakeRequest(uint128 methAmount, uint128 minETHAmount) returns(uint256)
func (_IStaking *IStakingTransactor) UnstakeRequest(opts *bind.TransactOpts, methAmount *big.Int, minETHAmount *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "unstakeRequest", methAmount, minETHAmount)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x891ef43e.
//
// Solidity: function unstakeRequest(uint128 methAmount, uint128 minETHAmount) returns(uint256)
func (_IStaking *IStakingSession) UnstakeRequest(methAmount *big.Int, minETHAmount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.UnstakeRequest(&_IStaking.TransactOpts, methAmount, minETHAmount)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x891ef43e.
//
// Solidity: function unstakeRequest(uint128 methAmount, uint128 minETHAmount) returns(uint256)
func (_IStaking *IStakingTransactorSession) UnstakeRequest(methAmount *big.Int, minETHAmount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.UnstakeRequest(&_IStaking.TransactOpts, methAmount, minETHAmount)
}
