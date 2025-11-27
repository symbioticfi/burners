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

// IRocketTokenRETHMetaData contains all meta data concerning the IRocketTokenRETH contract.
var IRocketTokenRETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_rethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEthValue\",\"inputs\":[{\"name\":\"_rethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRethValue\",\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IRocketTokenRETHABI is the input ABI used to generate the binding from.
// Deprecated: Use IRocketTokenRETHMetaData.ABI instead.
var IRocketTokenRETHABI = IRocketTokenRETHMetaData.ABI

// IRocketTokenRETH is an auto generated Go binding around an Ethereum contract.
type IRocketTokenRETH struct {
	IRocketTokenRETHCaller     // Read-only binding to the contract
	IRocketTokenRETHTransactor // Write-only binding to the contract
	IRocketTokenRETHFilterer   // Log filterer for contract events
}

// IRocketTokenRETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRocketTokenRETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRocketTokenRETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRocketTokenRETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRocketTokenRETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRocketTokenRETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRocketTokenRETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRocketTokenRETHSession struct {
	Contract     *IRocketTokenRETH // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRocketTokenRETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRocketTokenRETHCallerSession struct {
	Contract *IRocketTokenRETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IRocketTokenRETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRocketTokenRETHTransactorSession struct {
	Contract     *IRocketTokenRETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IRocketTokenRETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRocketTokenRETHRaw struct {
	Contract *IRocketTokenRETH // Generic contract binding to access the raw methods on
}

// IRocketTokenRETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRocketTokenRETHCallerRaw struct {
	Contract *IRocketTokenRETHCaller // Generic read-only contract binding to access the raw methods on
}

// IRocketTokenRETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRocketTokenRETHTransactorRaw struct {
	Contract *IRocketTokenRETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRocketTokenRETH creates a new instance of IRocketTokenRETH, bound to a specific deployed contract.
func NewIRocketTokenRETH(address common.Address, backend bind.ContractBackend) (*IRocketTokenRETH, error) {
	contract, err := bindIRocketTokenRETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRocketTokenRETH{IRocketTokenRETHCaller: IRocketTokenRETHCaller{contract: contract}, IRocketTokenRETHTransactor: IRocketTokenRETHTransactor{contract: contract}, IRocketTokenRETHFilterer: IRocketTokenRETHFilterer{contract: contract}}, nil
}

// NewIRocketTokenRETHCaller creates a new read-only instance of IRocketTokenRETH, bound to a specific deployed contract.
func NewIRocketTokenRETHCaller(address common.Address, caller bind.ContractCaller) (*IRocketTokenRETHCaller, error) {
	contract, err := bindIRocketTokenRETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRocketTokenRETHCaller{contract: contract}, nil
}

// NewIRocketTokenRETHTransactor creates a new write-only instance of IRocketTokenRETH, bound to a specific deployed contract.
func NewIRocketTokenRETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IRocketTokenRETHTransactor, error) {
	contract, err := bindIRocketTokenRETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRocketTokenRETHTransactor{contract: contract}, nil
}

// NewIRocketTokenRETHFilterer creates a new log filterer instance of IRocketTokenRETH, bound to a specific deployed contract.
func NewIRocketTokenRETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IRocketTokenRETHFilterer, error) {
	contract, err := bindIRocketTokenRETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRocketTokenRETHFilterer{contract: contract}, nil
}

// bindIRocketTokenRETH binds a generic wrapper to an already deployed contract.
func bindIRocketTokenRETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRocketTokenRETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRocketTokenRETH *IRocketTokenRETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRocketTokenRETH.Contract.IRocketTokenRETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRocketTokenRETH *IRocketTokenRETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.IRocketTokenRETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRocketTokenRETH *IRocketTokenRETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.IRocketTokenRETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRocketTokenRETH *IRocketTokenRETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRocketTokenRETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRocketTokenRETH *IRocketTokenRETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRocketTokenRETH *IRocketTokenRETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.contract.Transact(opts, method, params...)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCaller) GetEthValue(opts *bind.CallOpts, _rethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IRocketTokenRETH.contract.Call(opts, &out, "getEthValue", _rethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHSession) GetEthValue(_rethAmount *big.Int) (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetEthValue(&_IRocketTokenRETH.CallOpts, _rethAmount)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCallerSession) GetEthValue(_rethAmount *big.Int) (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetEthValue(&_IRocketTokenRETH.CallOpts, _rethAmount)
}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCaller) GetRethValue(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IRocketTokenRETH.contract.Call(opts, &out, "getRethValue", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHSession) GetRethValue(_ethAmount *big.Int) (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetRethValue(&_IRocketTokenRETH.CallOpts, _ethAmount)
}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCallerSession) GetRethValue(_ethAmount *big.Int) (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetRethValue(&_IRocketTokenRETH.CallOpts, _ethAmount)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCaller) GetTotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRocketTokenRETH.contract.Call(opts, &out, "getTotalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHSession) GetTotalCollateral() (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetTotalCollateral(&_IRocketTokenRETH.CallOpts)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_IRocketTokenRETH *IRocketTokenRETHCallerSession) GetTotalCollateral() (*big.Int, error) {
	return _IRocketTokenRETH.Contract.GetTotalCollateral(&_IRocketTokenRETH.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_IRocketTokenRETH *IRocketTokenRETHTransactor) Burn(opts *bind.TransactOpts, _rethAmount *big.Int) (*types.Transaction, error) {
	return _IRocketTokenRETH.contract.Transact(opts, "burn", _rethAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_IRocketTokenRETH *IRocketTokenRETHSession) Burn(_rethAmount *big.Int) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.Burn(&_IRocketTokenRETH.TransactOpts, _rethAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_IRocketTokenRETH *IRocketTokenRETHTransactorSession) Burn(_rethAmount *big.Int) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.Burn(&_IRocketTokenRETH.TransactOpts, _rethAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_IRocketTokenRETH *IRocketTokenRETHTransactor) Mint(opts *bind.TransactOpts, _ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IRocketTokenRETH.contract.Transact(opts, "mint", _ethAmount, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_IRocketTokenRETH *IRocketTokenRETHSession) Mint(_ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.Mint(&_IRocketTokenRETH.TransactOpts, _ethAmount, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_IRocketTokenRETH *IRocketTokenRETHTransactorSession) Mint(_ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IRocketTokenRETH.Contract.Mint(&_IRocketTokenRETH.TransactOpts, _ethAmount, _to)
}
