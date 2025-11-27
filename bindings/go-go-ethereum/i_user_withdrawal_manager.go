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

// IUserWithdrawalManagerMetaData contains all meta data concerning the IUserWithdrawalManager contract.
var IUserWithdrawalManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeUserWithdrawalRequest\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextRequestIdToFinalize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestWithdraw\",\"inputs\":[{\"name\":\"_ethXAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IUserWithdrawalManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IUserWithdrawalManagerMetaData.ABI instead.
var IUserWithdrawalManagerABI = IUserWithdrawalManagerMetaData.ABI

// IUserWithdrawalManager is an auto generated Go binding around an Ethereum contract.
type IUserWithdrawalManager struct {
	IUserWithdrawalManagerCaller     // Read-only binding to the contract
	IUserWithdrawalManagerTransactor // Write-only binding to the contract
	IUserWithdrawalManagerFilterer   // Log filterer for contract events
}

// IUserWithdrawalManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUserWithdrawalManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUserWithdrawalManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUserWithdrawalManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUserWithdrawalManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUserWithdrawalManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUserWithdrawalManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUserWithdrawalManagerSession struct {
	Contract     *IUserWithdrawalManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IUserWithdrawalManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUserWithdrawalManagerCallerSession struct {
	Contract *IUserWithdrawalManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IUserWithdrawalManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUserWithdrawalManagerTransactorSession struct {
	Contract     *IUserWithdrawalManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IUserWithdrawalManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUserWithdrawalManagerRaw struct {
	Contract *IUserWithdrawalManager // Generic contract binding to access the raw methods on
}

// IUserWithdrawalManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUserWithdrawalManagerCallerRaw struct {
	Contract *IUserWithdrawalManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IUserWithdrawalManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUserWithdrawalManagerTransactorRaw struct {
	Contract *IUserWithdrawalManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUserWithdrawalManager creates a new instance of IUserWithdrawalManager, bound to a specific deployed contract.
func NewIUserWithdrawalManager(address common.Address, backend bind.ContractBackend) (*IUserWithdrawalManager, error) {
	contract, err := bindIUserWithdrawalManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUserWithdrawalManager{IUserWithdrawalManagerCaller: IUserWithdrawalManagerCaller{contract: contract}, IUserWithdrawalManagerTransactor: IUserWithdrawalManagerTransactor{contract: contract}, IUserWithdrawalManagerFilterer: IUserWithdrawalManagerFilterer{contract: contract}}, nil
}

// NewIUserWithdrawalManagerCaller creates a new read-only instance of IUserWithdrawalManager, bound to a specific deployed contract.
func NewIUserWithdrawalManagerCaller(address common.Address, caller bind.ContractCaller) (*IUserWithdrawalManagerCaller, error) {
	contract, err := bindIUserWithdrawalManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUserWithdrawalManagerCaller{contract: contract}, nil
}

// NewIUserWithdrawalManagerTransactor creates a new write-only instance of IUserWithdrawalManager, bound to a specific deployed contract.
func NewIUserWithdrawalManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IUserWithdrawalManagerTransactor, error) {
	contract, err := bindIUserWithdrawalManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUserWithdrawalManagerTransactor{contract: contract}, nil
}

// NewIUserWithdrawalManagerFilterer creates a new log filterer instance of IUserWithdrawalManager, bound to a specific deployed contract.
func NewIUserWithdrawalManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IUserWithdrawalManagerFilterer, error) {
	contract, err := bindIUserWithdrawalManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUserWithdrawalManagerFilterer{contract: contract}, nil
}

// bindIUserWithdrawalManager binds a generic wrapper to an already deployed contract.
func bindIUserWithdrawalManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUserWithdrawalManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUserWithdrawalManager *IUserWithdrawalManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUserWithdrawalManager.Contract.IUserWithdrawalManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUserWithdrawalManager *IUserWithdrawalManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.IUserWithdrawalManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUserWithdrawalManager *IUserWithdrawalManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.IUserWithdrawalManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUserWithdrawalManager *IUserWithdrawalManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUserWithdrawalManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.contract.Transact(opts, method, params...)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUserWithdrawalManager.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerSession) NextRequestId() (*big.Int, error) {
	return _IUserWithdrawalManager.Contract.NextRequestId(&_IUserWithdrawalManager.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerCallerSession) NextRequestId() (*big.Int, error) {
	return _IUserWithdrawalManager.Contract.NextRequestId(&_IUserWithdrawalManager.CallOpts)
}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerCaller) NextRequestIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUserWithdrawalManager.contract.Call(opts, &out, "nextRequestIdToFinalize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerSession) NextRequestIdToFinalize() (*big.Int, error) {
	return _IUserWithdrawalManager.Contract.NextRequestIdToFinalize(&_IUserWithdrawalManager.CallOpts)
}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerCallerSession) NextRequestIdToFinalize() (*big.Int, error) {
	return _IUserWithdrawalManager.Contract.NextRequestIdToFinalize(&_IUserWithdrawalManager.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _requestId) returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactor) Claim(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _IUserWithdrawalManager.contract.Transact(opts, "claim", _requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _requestId) returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerSession) Claim(_requestId *big.Int) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.Claim(&_IUserWithdrawalManager.TransactOpts, _requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _requestId) returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactorSession) Claim(_requestId *big.Int) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.Claim(&_IUserWithdrawalManager.TransactOpts, _requestId)
}

// FinalizeUserWithdrawalRequest is a paid mutator transaction binding the contract method 0xad8a16dc.
//
// Solidity: function finalizeUserWithdrawalRequest() returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactor) FinalizeUserWithdrawalRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUserWithdrawalManager.contract.Transact(opts, "finalizeUserWithdrawalRequest")
}

// FinalizeUserWithdrawalRequest is a paid mutator transaction binding the contract method 0xad8a16dc.
//
// Solidity: function finalizeUserWithdrawalRequest() returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerSession) FinalizeUserWithdrawalRequest() (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.FinalizeUserWithdrawalRequest(&_IUserWithdrawalManager.TransactOpts)
}

// FinalizeUserWithdrawalRequest is a paid mutator transaction binding the contract method 0xad8a16dc.
//
// Solidity: function finalizeUserWithdrawalRequest() returns()
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactorSession) FinalizeUserWithdrawalRequest() (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.FinalizeUserWithdrawalRequest(&_IUserWithdrawalManager.TransactOpts)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _ethXAmount, address _owner) returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactor) RequestWithdraw(opts *bind.TransactOpts, _ethXAmount *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IUserWithdrawalManager.contract.Transact(opts, "requestWithdraw", _ethXAmount, _owner)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _ethXAmount, address _owner) returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerSession) RequestWithdraw(_ethXAmount *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.RequestWithdraw(&_IUserWithdrawalManager.TransactOpts, _ethXAmount, _owner)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _ethXAmount, address _owner) returns(uint256)
func (_IUserWithdrawalManager *IUserWithdrawalManagerTransactorSession) RequestWithdraw(_ethXAmount *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IUserWithdrawalManager.Contract.RequestWithdraw(&_IUserWithdrawalManager.TransactOpts, _ethXAmount, _owner)
}
