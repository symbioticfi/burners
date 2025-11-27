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

// IStaderStakePoolsManagerMetaData contains all meta data concerning the IStaderStakePoolsManager contract.
var IStaderStakePoolsManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"_assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"_shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// IStaderStakePoolsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IStaderStakePoolsManagerMetaData.ABI instead.
var IStaderStakePoolsManagerABI = IStaderStakePoolsManagerMetaData.ABI

// IStaderStakePoolsManager is an auto generated Go binding around an Ethereum contract.
type IStaderStakePoolsManager struct {
	IStaderStakePoolsManagerCaller     // Read-only binding to the contract
	IStaderStakePoolsManagerTransactor // Write-only binding to the contract
	IStaderStakePoolsManagerFilterer   // Log filterer for contract events
}

// IStaderStakePoolsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStaderStakePoolsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStaderStakePoolsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStaderStakePoolsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStaderStakePoolsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStaderStakePoolsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStaderStakePoolsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStaderStakePoolsManagerSession struct {
	Contract     *IStaderStakePoolsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IStaderStakePoolsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStaderStakePoolsManagerCallerSession struct {
	Contract *IStaderStakePoolsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IStaderStakePoolsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStaderStakePoolsManagerTransactorSession struct {
	Contract     *IStaderStakePoolsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IStaderStakePoolsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStaderStakePoolsManagerRaw struct {
	Contract *IStaderStakePoolsManager // Generic contract binding to access the raw methods on
}

// IStaderStakePoolsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStaderStakePoolsManagerCallerRaw struct {
	Contract *IStaderStakePoolsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStaderStakePoolsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStaderStakePoolsManagerTransactorRaw struct {
	Contract *IStaderStakePoolsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaderStakePoolsManager creates a new instance of IStaderStakePoolsManager, bound to a specific deployed contract.
func NewIStaderStakePoolsManager(address common.Address, backend bind.ContractBackend) (*IStaderStakePoolsManager, error) {
	contract, err := bindIStaderStakePoolsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaderStakePoolsManager{IStaderStakePoolsManagerCaller: IStaderStakePoolsManagerCaller{contract: contract}, IStaderStakePoolsManagerTransactor: IStaderStakePoolsManagerTransactor{contract: contract}, IStaderStakePoolsManagerFilterer: IStaderStakePoolsManagerFilterer{contract: contract}}, nil
}

// NewIStaderStakePoolsManagerCaller creates a new read-only instance of IStaderStakePoolsManager, bound to a specific deployed contract.
func NewIStaderStakePoolsManagerCaller(address common.Address, caller bind.ContractCaller) (*IStaderStakePoolsManagerCaller, error) {
	contract, err := bindIStaderStakePoolsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStaderStakePoolsManagerCaller{contract: contract}, nil
}

// NewIStaderStakePoolsManagerTransactor creates a new write-only instance of IStaderStakePoolsManager, bound to a specific deployed contract.
func NewIStaderStakePoolsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStaderStakePoolsManagerTransactor, error) {
	contract, err := bindIStaderStakePoolsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStaderStakePoolsManagerTransactor{contract: contract}, nil
}

// NewIStaderStakePoolsManagerFilterer creates a new log filterer instance of IStaderStakePoolsManager, bound to a specific deployed contract.
func NewIStaderStakePoolsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStaderStakePoolsManagerFilterer, error) {
	contract, err := bindIStaderStakePoolsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStaderStakePoolsManagerFilterer{contract: contract}, nil
}

// bindIStaderStakePoolsManager binds a generic wrapper to an already deployed contract.
func bindIStaderStakePoolsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStaderStakePoolsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaderStakePoolsManager.Contract.IStaderStakePoolsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaderStakePoolsManager.Contract.IStaderStakePoolsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaderStakePoolsManager.Contract.IStaderStakePoolsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaderStakePoolsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaderStakePoolsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaderStakePoolsManager.Contract.contract.Transact(opts, method, params...)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStaderStakePoolsManager.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _IStaderStakePoolsManager.Contract.PreviewDeposit(&_IStaderStakePoolsManager.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _IStaderStakePoolsManager.Contract.PreviewDeposit(&_IStaderStakePoolsManager.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerCaller) PreviewWithdraw(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStaderStakePoolsManager.contract.Call(opts, &out, "previewWithdraw", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerSession) PreviewWithdraw(_shares *big.Int) (*big.Int, error) {
	return _IStaderStakePoolsManager.Contract.PreviewWithdraw(&_IStaderStakePoolsManager.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_IStaderStakePoolsManager *IStaderStakePoolsManagerCallerSession) PreviewWithdraw(_shares *big.Int) (*big.Int, error) {
	return _IStaderStakePoolsManager.Contract.PreviewWithdraw(&_IStaderStakePoolsManager.CallOpts, _shares)
}
