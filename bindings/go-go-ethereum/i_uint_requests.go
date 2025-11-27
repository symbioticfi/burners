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

// IUintRequestsMetaData contains all meta data concerning the IUintRequests contract.
var IUintRequestsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// IUintRequestsABI is the input ABI used to generate the binding from.
// Deprecated: Use IUintRequestsMetaData.ABI instead.
var IUintRequestsABI = IUintRequestsMetaData.ABI

// IUintRequests is an auto generated Go binding around an Ethereum contract.
type IUintRequests struct {
	IUintRequestsCaller     // Read-only binding to the contract
	IUintRequestsTransactor // Write-only binding to the contract
	IUintRequestsFilterer   // Log filterer for contract events
}

// IUintRequestsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUintRequestsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUintRequestsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUintRequestsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUintRequestsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUintRequestsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUintRequestsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUintRequestsSession struct {
	Contract     *IUintRequests    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUintRequestsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUintRequestsCallerSession struct {
	Contract *IUintRequestsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IUintRequestsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUintRequestsTransactorSession struct {
	Contract     *IUintRequestsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IUintRequestsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUintRequestsRaw struct {
	Contract *IUintRequests // Generic contract binding to access the raw methods on
}

// IUintRequestsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUintRequestsCallerRaw struct {
	Contract *IUintRequestsCaller // Generic read-only contract binding to access the raw methods on
}

// IUintRequestsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUintRequestsTransactorRaw struct {
	Contract *IUintRequestsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUintRequests creates a new instance of IUintRequests, bound to a specific deployed contract.
func NewIUintRequests(address common.Address, backend bind.ContractBackend) (*IUintRequests, error) {
	contract, err := bindIUintRequests(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUintRequests{IUintRequestsCaller: IUintRequestsCaller{contract: contract}, IUintRequestsTransactor: IUintRequestsTransactor{contract: contract}, IUintRequestsFilterer: IUintRequestsFilterer{contract: contract}}, nil
}

// NewIUintRequestsCaller creates a new read-only instance of IUintRequests, bound to a specific deployed contract.
func NewIUintRequestsCaller(address common.Address, caller bind.ContractCaller) (*IUintRequestsCaller, error) {
	contract, err := bindIUintRequests(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUintRequestsCaller{contract: contract}, nil
}

// NewIUintRequestsTransactor creates a new write-only instance of IUintRequests, bound to a specific deployed contract.
func NewIUintRequestsTransactor(address common.Address, transactor bind.ContractTransactor) (*IUintRequestsTransactor, error) {
	contract, err := bindIUintRequests(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUintRequestsTransactor{contract: contract}, nil
}

// NewIUintRequestsFilterer creates a new log filterer instance of IUintRequests, bound to a specific deployed contract.
func NewIUintRequestsFilterer(address common.Address, filterer bind.ContractFilterer) (*IUintRequestsFilterer, error) {
	contract, err := bindIUintRequests(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUintRequestsFilterer{contract: contract}, nil
}

// bindIUintRequests binds a generic wrapper to an already deployed contract.
func bindIUintRequests(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUintRequestsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUintRequests *IUintRequestsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUintRequests.Contract.IUintRequestsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUintRequests *IUintRequestsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUintRequests.Contract.IUintRequestsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUintRequests *IUintRequestsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUintRequests.Contract.IUintRequestsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUintRequests *IUintRequestsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUintRequests.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUintRequests *IUintRequestsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUintRequests.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUintRequests *IUintRequestsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUintRequests.Contract.contract.Transact(opts, method, params...)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IUintRequests *IUintRequestsCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IUintRequests.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IUintRequests *IUintRequestsSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IUintRequests.Contract.RequestIds(&_IUintRequests.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IUintRequests *IUintRequestsCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IUintRequests.Contract.RequestIds(&_IUintRequests.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IUintRequests *IUintRequestsCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUintRequests.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IUintRequests *IUintRequestsSession) RequestIdsLength() (*big.Int, error) {
	return _IUintRequests.Contract.RequestIdsLength(&_IUintRequests.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IUintRequests *IUintRequestsCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IUintRequests.Contract.RequestIdsLength(&_IUintRequests.CallOpts)
}
