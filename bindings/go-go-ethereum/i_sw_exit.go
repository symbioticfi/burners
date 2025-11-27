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

// ISwEXITMetaData contains all meta data concerning the ISwEXIT contract.
var ISwEXITMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createWithdrawRequest\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeWithdrawal\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastTokenIdCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processWithdrawals\",\"inputs\":[{\"name\":\"_lastTokenIdToProcess\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRequestMaximum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawRequestMinimum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// ISwEXITABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwEXITMetaData.ABI instead.
var ISwEXITABI = ISwEXITMetaData.ABI

// ISwEXIT is an auto generated Go binding around an Ethereum contract.
type ISwEXIT struct {
	ISwEXITCaller     // Read-only binding to the contract
	ISwEXITTransactor // Write-only binding to the contract
	ISwEXITFilterer   // Log filterer for contract events
}

// ISwEXITCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwEXITCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwEXITTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwEXITTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwEXITFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwEXITFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwEXITSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwEXITSession struct {
	Contract     *ISwEXIT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwEXITCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwEXITCallerSession struct {
	Contract *ISwEXITCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ISwEXITTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwEXITTransactorSession struct {
	Contract     *ISwEXITTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ISwEXITRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwEXITRaw struct {
	Contract *ISwEXIT // Generic contract binding to access the raw methods on
}

// ISwEXITCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwEXITCallerRaw struct {
	Contract *ISwEXITCaller // Generic read-only contract binding to access the raw methods on
}

// ISwEXITTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwEXITTransactorRaw struct {
	Contract *ISwEXITTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwEXIT creates a new instance of ISwEXIT, bound to a specific deployed contract.
func NewISwEXIT(address common.Address, backend bind.ContractBackend) (*ISwEXIT, error) {
	contract, err := bindISwEXIT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwEXIT{ISwEXITCaller: ISwEXITCaller{contract: contract}, ISwEXITTransactor: ISwEXITTransactor{contract: contract}, ISwEXITFilterer: ISwEXITFilterer{contract: contract}}, nil
}

// NewISwEXITCaller creates a new read-only instance of ISwEXIT, bound to a specific deployed contract.
func NewISwEXITCaller(address common.Address, caller bind.ContractCaller) (*ISwEXITCaller, error) {
	contract, err := bindISwEXIT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwEXITCaller{contract: contract}, nil
}

// NewISwEXITTransactor creates a new write-only instance of ISwEXIT, bound to a specific deployed contract.
func NewISwEXITTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwEXITTransactor, error) {
	contract, err := bindISwEXIT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwEXITTransactor{contract: contract}, nil
}

// NewISwEXITFilterer creates a new log filterer instance of ISwEXIT, bound to a specific deployed contract.
func NewISwEXITFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwEXITFilterer, error) {
	contract, err := bindISwEXIT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwEXITFilterer{contract: contract}, nil
}

// bindISwEXIT binds a generic wrapper to an already deployed contract.
func bindISwEXIT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwEXITMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwEXIT *ISwEXITRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwEXIT.Contract.ISwEXITCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwEXIT *ISwEXITRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwEXIT.Contract.ISwEXITTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwEXIT *ISwEXITRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwEXIT.Contract.ISwEXITTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwEXIT *ISwEXITCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwEXIT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwEXIT *ISwEXITTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwEXIT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwEXIT *ISwEXITTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwEXIT.Contract.contract.Transact(opts, method, params...)
}

// GetLastTokenIdCreated is a free data retrieval call binding the contract method 0x061a499f.
//
// Solidity: function getLastTokenIdCreated() view returns(uint256)
func (_ISwEXIT *ISwEXITCaller) GetLastTokenIdCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwEXIT.contract.Call(opts, &out, "getLastTokenIdCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastTokenIdCreated is a free data retrieval call binding the contract method 0x061a499f.
//
// Solidity: function getLastTokenIdCreated() view returns(uint256)
func (_ISwEXIT *ISwEXITSession) GetLastTokenIdCreated() (*big.Int, error) {
	return _ISwEXIT.Contract.GetLastTokenIdCreated(&_ISwEXIT.CallOpts)
}

// GetLastTokenIdCreated is a free data retrieval call binding the contract method 0x061a499f.
//
// Solidity: function getLastTokenIdCreated() view returns(uint256)
func (_ISwEXIT *ISwEXITCallerSession) GetLastTokenIdCreated() (*big.Int, error) {
	return _ISwEXIT.Contract.GetLastTokenIdCreated(&_ISwEXIT.CallOpts)
}

// WithdrawRequestMaximum is a free data retrieval call binding the contract method 0xef8526f3.
//
// Solidity: function withdrawRequestMaximum() view returns(uint256)
func (_ISwEXIT *ISwEXITCaller) WithdrawRequestMaximum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwEXIT.contract.Call(opts, &out, "withdrawRequestMaximum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawRequestMaximum is a free data retrieval call binding the contract method 0xef8526f3.
//
// Solidity: function withdrawRequestMaximum() view returns(uint256)
func (_ISwEXIT *ISwEXITSession) WithdrawRequestMaximum() (*big.Int, error) {
	return _ISwEXIT.Contract.WithdrawRequestMaximum(&_ISwEXIT.CallOpts)
}

// WithdrawRequestMaximum is a free data retrieval call binding the contract method 0xef8526f3.
//
// Solidity: function withdrawRequestMaximum() view returns(uint256)
func (_ISwEXIT *ISwEXITCallerSession) WithdrawRequestMaximum() (*big.Int, error) {
	return _ISwEXIT.Contract.WithdrawRequestMaximum(&_ISwEXIT.CallOpts)
}

// WithdrawRequestMinimum is a free data retrieval call binding the contract method 0xf049db24.
//
// Solidity: function withdrawRequestMinimum() view returns(uint256)
func (_ISwEXIT *ISwEXITCaller) WithdrawRequestMinimum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwEXIT.contract.Call(opts, &out, "withdrawRequestMinimum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawRequestMinimum is a free data retrieval call binding the contract method 0xf049db24.
//
// Solidity: function withdrawRequestMinimum() view returns(uint256)
func (_ISwEXIT *ISwEXITSession) WithdrawRequestMinimum() (*big.Int, error) {
	return _ISwEXIT.Contract.WithdrawRequestMinimum(&_ISwEXIT.CallOpts)
}

// WithdrawRequestMinimum is a free data retrieval call binding the contract method 0xf049db24.
//
// Solidity: function withdrawRequestMinimum() view returns(uint256)
func (_ISwEXIT *ISwEXITCallerSession) WithdrawRequestMinimum() (*big.Int, error) {
	return _ISwEXIT.Contract.WithdrawRequestMinimum(&_ISwEXIT.CallOpts)
}

// CreateWithdrawRequest is a paid mutator transaction binding the contract method 0x74dc9d1a.
//
// Solidity: function createWithdrawRequest(uint256 amount) returns()
func (_ISwEXIT *ISwEXITTransactor) CreateWithdrawRequest(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.contract.Transact(opts, "createWithdrawRequest", amount)
}

// CreateWithdrawRequest is a paid mutator transaction binding the contract method 0x74dc9d1a.
//
// Solidity: function createWithdrawRequest(uint256 amount) returns()
func (_ISwEXIT *ISwEXITSession) CreateWithdrawRequest(amount *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.CreateWithdrawRequest(&_ISwEXIT.TransactOpts, amount)
}

// CreateWithdrawRequest is a paid mutator transaction binding the contract method 0x74dc9d1a.
//
// Solidity: function createWithdrawRequest(uint256 amount) returns()
func (_ISwEXIT *ISwEXITTransactorSession) CreateWithdrawRequest(amount *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.CreateWithdrawRequest(&_ISwEXIT.TransactOpts, amount)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x5e15c749.
//
// Solidity: function finalizeWithdrawal(uint256 tokenId) returns()
func (_ISwEXIT *ISwEXITTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.contract.Transact(opts, "finalizeWithdrawal", tokenId)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x5e15c749.
//
// Solidity: function finalizeWithdrawal(uint256 tokenId) returns()
func (_ISwEXIT *ISwEXITSession) FinalizeWithdrawal(tokenId *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.FinalizeWithdrawal(&_ISwEXIT.TransactOpts, tokenId)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x5e15c749.
//
// Solidity: function finalizeWithdrawal(uint256 tokenId) returns()
func (_ISwEXIT *ISwEXITTransactorSession) FinalizeWithdrawal(tokenId *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.FinalizeWithdrawal(&_ISwEXIT.TransactOpts, tokenId)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0x152fcb0c.
//
// Solidity: function processWithdrawals(uint256 _lastTokenIdToProcess) returns()
func (_ISwEXIT *ISwEXITTransactor) ProcessWithdrawals(opts *bind.TransactOpts, _lastTokenIdToProcess *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.contract.Transact(opts, "processWithdrawals", _lastTokenIdToProcess)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0x152fcb0c.
//
// Solidity: function processWithdrawals(uint256 _lastTokenIdToProcess) returns()
func (_ISwEXIT *ISwEXITSession) ProcessWithdrawals(_lastTokenIdToProcess *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.ProcessWithdrawals(&_ISwEXIT.TransactOpts, _lastTokenIdToProcess)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0x152fcb0c.
//
// Solidity: function processWithdrawals(uint256 _lastTokenIdToProcess) returns()
func (_ISwEXIT *ISwEXITTransactorSession) ProcessWithdrawals(_lastTokenIdToProcess *big.Int) (*types.Transaction, error) {
	return _ISwEXIT.Contract.ProcessWithdrawals(&_ISwEXIT.TransactOpts, _lastTokenIdToProcess)
}
