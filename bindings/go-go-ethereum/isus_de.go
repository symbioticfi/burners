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

// ISUSDeMetaData contains all meta data concerning the ISUSDe contract.
var ISUSDeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cooldownDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cooldownShares\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCooldownDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// ISUSDeABI is the input ABI used to generate the binding from.
// Deprecated: Use ISUSDeMetaData.ABI instead.
var ISUSDeABI = ISUSDeMetaData.ABI

// ISUSDe is an auto generated Go binding around an Ethereum contract.
type ISUSDe struct {
	ISUSDeCaller     // Read-only binding to the contract
	ISUSDeTransactor // Write-only binding to the contract
	ISUSDeFilterer   // Log filterer for contract events
}

// ISUSDeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISUSDeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISUSDeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISUSDeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISUSDeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISUSDeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISUSDeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISUSDeSession struct {
	Contract     *ISUSDe           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISUSDeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISUSDeCallerSession struct {
	Contract *ISUSDeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISUSDeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISUSDeTransactorSession struct {
	Contract     *ISUSDeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISUSDeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISUSDeRaw struct {
	Contract *ISUSDe // Generic contract binding to access the raw methods on
}

// ISUSDeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISUSDeCallerRaw struct {
	Contract *ISUSDeCaller // Generic read-only contract binding to access the raw methods on
}

// ISUSDeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISUSDeTransactorRaw struct {
	Contract *ISUSDeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISUSDe creates a new instance of ISUSDe, bound to a specific deployed contract.
func NewISUSDe(address common.Address, backend bind.ContractBackend) (*ISUSDe, error) {
	contract, err := bindISUSDe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISUSDe{ISUSDeCaller: ISUSDeCaller{contract: contract}, ISUSDeTransactor: ISUSDeTransactor{contract: contract}, ISUSDeFilterer: ISUSDeFilterer{contract: contract}}, nil
}

// NewISUSDeCaller creates a new read-only instance of ISUSDe, bound to a specific deployed contract.
func NewISUSDeCaller(address common.Address, caller bind.ContractCaller) (*ISUSDeCaller, error) {
	contract, err := bindISUSDe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISUSDeCaller{contract: contract}, nil
}

// NewISUSDeTransactor creates a new write-only instance of ISUSDe, bound to a specific deployed contract.
func NewISUSDeTransactor(address common.Address, transactor bind.ContractTransactor) (*ISUSDeTransactor, error) {
	contract, err := bindISUSDe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISUSDeTransactor{contract: contract}, nil
}

// NewISUSDeFilterer creates a new log filterer instance of ISUSDe, bound to a specific deployed contract.
func NewISUSDeFilterer(address common.Address, filterer bind.ContractFilterer) (*ISUSDeFilterer, error) {
	contract, err := bindISUSDe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISUSDeFilterer{contract: contract}, nil
}

// bindISUSDe binds a generic wrapper to an already deployed contract.
func bindISUSDe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISUSDeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISUSDe *ISUSDeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISUSDe.Contract.ISUSDeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISUSDe *ISUSDeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISUSDe.Contract.ISUSDeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISUSDe *ISUSDeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISUSDe.Contract.ISUSDeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISUSDe *ISUSDeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISUSDe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISUSDe *ISUSDeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISUSDe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISUSDe *ISUSDeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISUSDe.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ISUSDe *ISUSDeCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISUSDe.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ISUSDe *ISUSDeSession) Asset() (common.Address, error) {
	return _ISUSDe.Contract.Asset(&_ISUSDe.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ISUSDe *ISUSDeCallerSession) Asset() (common.Address, error) {
	return _ISUSDe.Contract.Asset(&_ISUSDe.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint24)
func (_ISUSDe *ISUSDeCaller) CooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISUSDe.contract.Call(opts, &out, "cooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint24)
func (_ISUSDe *ISUSDeSession) CooldownDuration() (*big.Int, error) {
	return _ISUSDe.Contract.CooldownDuration(&_ISUSDe.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint24)
func (_ISUSDe *ISUSDeCallerSession) CooldownDuration() (*big.Int, error) {
	return _ISUSDe.Contract.CooldownDuration(&_ISUSDe.CallOpts)
}

// CooldownShares is a paid mutator transaction binding the contract method 0x9343d9e1.
//
// Solidity: function cooldownShares(uint256 shares) returns(uint256 assets)
func (_ISUSDe *ISUSDeTransactor) CooldownShares(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "cooldownShares", shares)
}

// CooldownShares is a paid mutator transaction binding the contract method 0x9343d9e1.
//
// Solidity: function cooldownShares(uint256 shares) returns(uint256 assets)
func (_ISUSDe *ISUSDeSession) CooldownShares(shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.CooldownShares(&_ISUSDe.TransactOpts, shares)
}

// CooldownShares is a paid mutator transaction binding the contract method 0x9343d9e1.
//
// Solidity: function cooldownShares(uint256 shares) returns(uint256 assets)
func (_ISUSDe *ISUSDeTransactorSession) CooldownShares(shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.CooldownShares(&_ISUSDe.TransactOpts, shares)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ISUSDe *ISUSDeTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ISUSDe *ISUSDeSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Deposit(&_ISUSDe.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ISUSDe *ISUSDeTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Deposit(&_ISUSDe.TransactOpts, assets, receiver)
}

// PreviewRedeem is a paid mutator transaction binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) returns(uint256)
func (_ISUSDe *ISUSDeTransactor) PreviewRedeem(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "previewRedeem", shares)
}

// PreviewRedeem is a paid mutator transaction binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) returns(uint256)
func (_ISUSDe *ISUSDeSession) PreviewRedeem(shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.PreviewRedeem(&_ISUSDe.TransactOpts, shares)
}

// PreviewRedeem is a paid mutator transaction binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) returns(uint256)
func (_ISUSDe *ISUSDeTransactorSession) PreviewRedeem(shares *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.PreviewRedeem(&_ISUSDe.TransactOpts, shares)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_ISUSDe *ISUSDeTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "redeem", shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_ISUSDe *ISUSDeSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Redeem(&_ISUSDe.TransactOpts, shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_ISUSDe *ISUSDeTransactorSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Redeem(&_ISUSDe.TransactOpts, shares, receiver, _owner)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0xce23eb3c.
//
// Solidity: function setCooldownDuration(uint24 duration) returns()
func (_ISUSDe *ISUSDeTransactor) SetCooldownDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "setCooldownDuration", duration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0xce23eb3c.
//
// Solidity: function setCooldownDuration(uint24 duration) returns()
func (_ISUSDe *ISUSDeSession) SetCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.SetCooldownDuration(&_ISUSDe.TransactOpts, duration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0xce23eb3c.
//
// Solidity: function setCooldownDuration(uint24 duration) returns()
func (_ISUSDe *ISUSDeTransactorSession) SetCooldownDuration(duration *big.Int) (*types.Transaction, error) {
	return _ISUSDe.Contract.SetCooldownDuration(&_ISUSDe.TransactOpts, duration)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address receiver) returns()
func (_ISUSDe *ISUSDeTransactor) Unstake(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.contract.Transact(opts, "unstake", receiver)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address receiver) returns()
func (_ISUSDe *ISUSDeSession) Unstake(receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Unstake(&_ISUSDe.TransactOpts, receiver)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address receiver) returns()
func (_ISUSDe *ISUSDeTransactorSession) Unstake(receiver common.Address) (*types.Transaction, error) {
	return _ISUSDe.Contract.Unstake(&_ISUSDe.TransactOpts, receiver)
}
