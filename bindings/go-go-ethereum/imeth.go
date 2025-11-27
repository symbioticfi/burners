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

// IMETHMetaData contains all meta data concerning the IMETH contract.
var IMETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// IMETHABI is the input ABI used to generate the binding from.
// Deprecated: Use IMETHMetaData.ABI instead.
var IMETHABI = IMETHMetaData.ABI

// IMETH is an auto generated Go binding around an Ethereum contract.
type IMETH struct {
	IMETHCaller     // Read-only binding to the contract
	IMETHTransactor // Write-only binding to the contract
	IMETHFilterer   // Log filterer for contract events
}

// IMETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMETHSession struct {
	Contract     *IMETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMETHCallerSession struct {
	Contract *IMETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IMETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMETHTransactorSession struct {
	Contract     *IMETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMETHRaw struct {
	Contract *IMETH // Generic contract binding to access the raw methods on
}

// IMETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMETHCallerRaw struct {
	Contract *IMETHCaller // Generic read-only contract binding to access the raw methods on
}

// IMETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMETHTransactorRaw struct {
	Contract *IMETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMETH creates a new instance of IMETH, bound to a specific deployed contract.
func NewIMETH(address common.Address, backend bind.ContractBackend) (*IMETH, error) {
	contract, err := bindIMETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMETH{IMETHCaller: IMETHCaller{contract: contract}, IMETHTransactor: IMETHTransactor{contract: contract}, IMETHFilterer: IMETHFilterer{contract: contract}}, nil
}

// NewIMETHCaller creates a new read-only instance of IMETH, bound to a specific deployed contract.
func NewIMETHCaller(address common.Address, caller bind.ContractCaller) (*IMETHCaller, error) {
	contract, err := bindIMETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMETHCaller{contract: contract}, nil
}

// NewIMETHTransactor creates a new write-only instance of IMETH, bound to a specific deployed contract.
func NewIMETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IMETHTransactor, error) {
	contract, err := bindIMETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMETHTransactor{contract: contract}, nil
}

// NewIMETHFilterer creates a new log filterer instance of IMETH, bound to a specific deployed contract.
func NewIMETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IMETHFilterer, error) {
	contract, err := bindIMETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMETHFilterer{contract: contract}, nil
}

// bindIMETH binds a generic wrapper to an already deployed contract.
func bindIMETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMETH *IMETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMETH.Contract.IMETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMETH *IMETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMETH.Contract.IMETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMETH *IMETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMETH.Contract.IMETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMETH *IMETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMETH *IMETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMETH *IMETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMETH.Contract.contract.Transact(opts, method, params...)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_IMETH *IMETHCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMETH.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_IMETH *IMETHSession) StakingContract() (common.Address, error) {
	return _IMETH.Contract.StakingContract(&_IMETH.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_IMETH *IMETHCallerSession) StakingContract() (common.Address, error) {
	return _IMETH.Contract.StakingContract(&_IMETH.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_IMETH *IMETHTransactor) Mint(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMETH.contract.Transact(opts, "mint", staker, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_IMETH *IMETHSession) Mint(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMETH.Contract.Mint(&_IMETH.TransactOpts, staker, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_IMETH *IMETHTransactorSession) Mint(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMETH.Contract.Mint(&_IMETH.TransactOpts, staker, amount)
}
