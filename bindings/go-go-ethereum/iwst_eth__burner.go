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

// IwstETHBurnerMetaData contains all meta data concerning the IwstETHBurner contract.
var IwstETHBurnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"COLLATERAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIDO_WITHDRAWAL_QUEUE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_STETH_WITHDRAWAL_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_STETH_WITHDRAWAL_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRequestIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestIdsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerBurn\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerBurnBatch\",\"inputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"hints\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerWithdrawal\",\"inputs\":[{\"name\":\"maxRequests\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TriggerBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerBurnBatch\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TriggerWithdrawal\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[]}]",
}

// IwstETHBurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IwstETHBurnerMetaData.ABI instead.
var IwstETHBurnerABI = IwstETHBurnerMetaData.ABI

// IwstETHBurner is an auto generated Go binding around an Ethereum contract.
type IwstETHBurner struct {
	IwstETHBurnerCaller     // Read-only binding to the contract
	IwstETHBurnerTransactor // Write-only binding to the contract
	IwstETHBurnerFilterer   // Log filterer for contract events
}

// IwstETHBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IwstETHBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwstETHBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IwstETHBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwstETHBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IwstETHBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwstETHBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IwstETHBurnerSession struct {
	Contract     *IwstETHBurner    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IwstETHBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IwstETHBurnerCallerSession struct {
	Contract *IwstETHBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IwstETHBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IwstETHBurnerTransactorSession struct {
	Contract     *IwstETHBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IwstETHBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IwstETHBurnerRaw struct {
	Contract *IwstETHBurner // Generic contract binding to access the raw methods on
}

// IwstETHBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IwstETHBurnerCallerRaw struct {
	Contract *IwstETHBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// IwstETHBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IwstETHBurnerTransactorRaw struct {
	Contract *IwstETHBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIwstETHBurner creates a new instance of IwstETHBurner, bound to a specific deployed contract.
func NewIwstETHBurner(address common.Address, backend bind.ContractBackend) (*IwstETHBurner, error) {
	contract, err := bindIwstETHBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurner{IwstETHBurnerCaller: IwstETHBurnerCaller{contract: contract}, IwstETHBurnerTransactor: IwstETHBurnerTransactor{contract: contract}, IwstETHBurnerFilterer: IwstETHBurnerFilterer{contract: contract}}, nil
}

// NewIwstETHBurnerCaller creates a new read-only instance of IwstETHBurner, bound to a specific deployed contract.
func NewIwstETHBurnerCaller(address common.Address, caller bind.ContractCaller) (*IwstETHBurnerCaller, error) {
	contract, err := bindIwstETHBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerCaller{contract: contract}, nil
}

// NewIwstETHBurnerTransactor creates a new write-only instance of IwstETHBurner, bound to a specific deployed contract.
func NewIwstETHBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IwstETHBurnerTransactor, error) {
	contract, err := bindIwstETHBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerTransactor{contract: contract}, nil
}

// NewIwstETHBurnerFilterer creates a new log filterer instance of IwstETHBurner, bound to a specific deployed contract.
func NewIwstETHBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IwstETHBurnerFilterer, error) {
	contract, err := bindIwstETHBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerFilterer{contract: contract}, nil
}

// bindIwstETHBurner binds a generic wrapper to an already deployed contract.
func bindIwstETHBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IwstETHBurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IwstETHBurner *IwstETHBurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IwstETHBurner.Contract.IwstETHBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IwstETHBurner *IwstETHBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.IwstETHBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IwstETHBurner *IwstETHBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.IwstETHBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IwstETHBurner *IwstETHBurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IwstETHBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IwstETHBurner *IwstETHBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IwstETHBurner *IwstETHBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.contract.Transact(opts, method, params...)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCaller) COLLATERAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "COLLATERAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IwstETHBurner *IwstETHBurnerSession) COLLATERAL() (common.Address, error) {
	return _IwstETHBurner.Contract.COLLATERAL(&_IwstETHBurner.CallOpts)
}

// COLLATERAL is a free data retrieval call binding the contract method 0x24bbab8b.
//
// Solidity: function COLLATERAL() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCallerSession) COLLATERAL() (common.Address, error) {
	return _IwstETHBurner.Contract.COLLATERAL(&_IwstETHBurner.CallOpts)
}

// LIDOWITHDRAWALQUEUE is a free data retrieval call binding the contract method 0xb8c77774.
//
// Solidity: function LIDO_WITHDRAWAL_QUEUE() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCaller) LIDOWITHDRAWALQUEUE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "LIDO_WITHDRAWAL_QUEUE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOWITHDRAWALQUEUE is a free data retrieval call binding the contract method 0xb8c77774.
//
// Solidity: function LIDO_WITHDRAWAL_QUEUE() view returns(address)
func (_IwstETHBurner *IwstETHBurnerSession) LIDOWITHDRAWALQUEUE() (common.Address, error) {
	return _IwstETHBurner.Contract.LIDOWITHDRAWALQUEUE(&_IwstETHBurner.CallOpts)
}

// LIDOWITHDRAWALQUEUE is a free data retrieval call binding the contract method 0xb8c77774.
//
// Solidity: function LIDO_WITHDRAWAL_QUEUE() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCallerSession) LIDOWITHDRAWALQUEUE() (common.Address, error) {
	return _IwstETHBurner.Contract.LIDOWITHDRAWALQUEUE(&_IwstETHBurner.CallOpts)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCaller) MAXSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "MAX_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IwstETHBurner.Contract.MAXSTETHWITHDRAWALAMOUNT(&_IwstETHBurner.CallOpts)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCallerSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IwstETHBurner.Contract.MAXSTETHWITHDRAWALAMOUNT(&_IwstETHBurner.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCaller) MINSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "MIN_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IwstETHBurner.Contract.MINSTETHWITHDRAWALAMOUNT(&_IwstETHBurner.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCallerSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IwstETHBurner.Contract.MINSTETHWITHDRAWALAMOUNT(&_IwstETHBurner.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IwstETHBurner *IwstETHBurnerSession) STETH() (common.Address, error) {
	return _IwstETHBurner.Contract.STETH(&_IwstETHBurner.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IwstETHBurner *IwstETHBurnerCallerSession) STETH() (common.Address, error) {
	return _IwstETHBurner.Contract.STETH(&_IwstETHBurner.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerCaller) RequestIds(opts *bind.CallOpts, index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "requestIds", index, maxRequestIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IwstETHBurner.Contract.RequestIds(&_IwstETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIds is a free data retrieval call binding the contract method 0x4383ee3d.
//
// Solidity: function requestIds(uint256 index, uint256 maxRequestIds) view returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerCallerSession) RequestIds(index *big.Int, maxRequestIds *big.Int) ([]*big.Int, error) {
	return _IwstETHBurner.Contract.RequestIds(&_IwstETHBurner.CallOpts, index, maxRequestIds)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCaller) RequestIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IwstETHBurner.contract.Call(opts, &out, "requestIdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerSession) RequestIdsLength() (*big.Int, error) {
	return _IwstETHBurner.Contract.RequestIdsLength(&_IwstETHBurner.CallOpts)
}

// RequestIdsLength is a free data retrieval call binding the contract method 0x45a67f51.
//
// Solidity: function requestIdsLength() view returns(uint256)
func (_IwstETHBurner *IwstETHBurnerCallerSession) RequestIdsLength() (*big.Int, error) {
	return _IwstETHBurner.Contract.RequestIdsLength(&_IwstETHBurner.CallOpts)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IwstETHBurner *IwstETHBurnerTransactor) TriggerBurn(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.contract.Transact(opts, "triggerBurn", requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IwstETHBurner *IwstETHBurnerSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerBurn(&_IwstETHBurner.TransactOpts, requestId)
}

// TriggerBurn is a paid mutator transaction binding the contract method 0x0bc8cbcf.
//
// Solidity: function triggerBurn(uint256 requestId) returns()
func (_IwstETHBurner *IwstETHBurnerTransactorSession) TriggerBurn(requestId *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerBurn(&_IwstETHBurner.TransactOpts, requestId)
}

// TriggerBurnBatch is a paid mutator transaction binding the contract method 0x5faeff4c.
//
// Solidity: function triggerBurnBatch(uint256[] requestIds, uint256[] hints) returns()
func (_IwstETHBurner *IwstETHBurnerTransactor) TriggerBurnBatch(opts *bind.TransactOpts, requestIds []*big.Int, hints []*big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.contract.Transact(opts, "triggerBurnBatch", requestIds, hints)
}

// TriggerBurnBatch is a paid mutator transaction binding the contract method 0x5faeff4c.
//
// Solidity: function triggerBurnBatch(uint256[] requestIds, uint256[] hints) returns()
func (_IwstETHBurner *IwstETHBurnerSession) TriggerBurnBatch(requestIds []*big.Int, hints []*big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerBurnBatch(&_IwstETHBurner.TransactOpts, requestIds, hints)
}

// TriggerBurnBatch is a paid mutator transaction binding the contract method 0x5faeff4c.
//
// Solidity: function triggerBurnBatch(uint256[] requestIds, uint256[] hints) returns()
func (_IwstETHBurner *IwstETHBurnerTransactorSession) TriggerBurnBatch(requestIds []*big.Int, hints []*big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerBurnBatch(&_IwstETHBurner.TransactOpts, requestIds, hints)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerTransactor) TriggerWithdrawal(opts *bind.TransactOpts, maxRequests *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.contract.Transact(opts, "triggerWithdrawal", maxRequests)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerSession) TriggerWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerWithdrawal(&_IwstETHBurner.TransactOpts, maxRequests)
}

// TriggerWithdrawal is a paid mutator transaction binding the contract method 0x92284cb6.
//
// Solidity: function triggerWithdrawal(uint256 maxRequests) returns(uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerTransactorSession) TriggerWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _IwstETHBurner.Contract.TriggerWithdrawal(&_IwstETHBurner.TransactOpts, maxRequests)
}

// IwstETHBurnerTriggerBurnIterator is returned from FilterTriggerBurn and is used to iterate over the raw logs and unpacked data for TriggerBurn events raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerBurnIterator struct {
	Event *IwstETHBurnerTriggerBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IwstETHBurnerTriggerBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IwstETHBurnerTriggerBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IwstETHBurnerTriggerBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IwstETHBurnerTriggerBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IwstETHBurnerTriggerBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IwstETHBurnerTriggerBurn represents a TriggerBurn event raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerBurn struct {
	Caller    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurn is a free log retrieval operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IwstETHBurner *IwstETHBurnerFilterer) FilterTriggerBurn(opts *bind.FilterOpts, caller []common.Address) (*IwstETHBurnerTriggerBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.FilterLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerTriggerBurnIterator{contract: _IwstETHBurner.contract, event: "TriggerBurn", logs: logs, sub: sub}, nil
}

// WatchTriggerBurn is a free log subscription operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IwstETHBurner *IwstETHBurnerFilterer) WatchTriggerBurn(opts *bind.WatchOpts, sink chan<- *IwstETHBurnerTriggerBurn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.WatchLogs(opts, "TriggerBurn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IwstETHBurnerTriggerBurn)
				if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerBurn is a log parse operation binding the contract event 0x4ee9ebe815db7e3e7645cd8bd4d663bf6b60b8505a4dd52fcf38aaef8b6d9884.
//
// Solidity: event TriggerBurn(address indexed caller, uint256 requestId)
func (_IwstETHBurner *IwstETHBurnerFilterer) ParseTriggerBurn(log types.Log) (*IwstETHBurnerTriggerBurn, error) {
	event := new(IwstETHBurnerTriggerBurn)
	if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IwstETHBurnerTriggerBurnBatchIterator is returned from FilterTriggerBurnBatch and is used to iterate over the raw logs and unpacked data for TriggerBurnBatch events raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerBurnBatchIterator struct {
	Event *IwstETHBurnerTriggerBurnBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IwstETHBurnerTriggerBurnBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IwstETHBurnerTriggerBurnBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IwstETHBurnerTriggerBurnBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IwstETHBurnerTriggerBurnBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IwstETHBurnerTriggerBurnBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IwstETHBurnerTriggerBurnBatch represents a TriggerBurnBatch event raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerBurnBatch struct {
	Caller     common.Address
	RequestIds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTriggerBurnBatch is a free log retrieval operation binding the contract event 0x30d1c85d591b30123c542e62c0c42de83290455affaab177afca24b1c3aecd1e.
//
// Solidity: event TriggerBurnBatch(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) FilterTriggerBurnBatch(opts *bind.FilterOpts, caller []common.Address) (*IwstETHBurnerTriggerBurnBatchIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.FilterLogs(opts, "TriggerBurnBatch", callerRule)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerTriggerBurnBatchIterator{contract: _IwstETHBurner.contract, event: "TriggerBurnBatch", logs: logs, sub: sub}, nil
}

// WatchTriggerBurnBatch is a free log subscription operation binding the contract event 0x30d1c85d591b30123c542e62c0c42de83290455affaab177afca24b1c3aecd1e.
//
// Solidity: event TriggerBurnBatch(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) WatchTriggerBurnBatch(opts *bind.WatchOpts, sink chan<- *IwstETHBurnerTriggerBurnBatch, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.WatchLogs(opts, "TriggerBurnBatch", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IwstETHBurnerTriggerBurnBatch)
				if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerBurnBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerBurnBatch is a log parse operation binding the contract event 0x30d1c85d591b30123c542e62c0c42de83290455affaab177afca24b1c3aecd1e.
//
// Solidity: event TriggerBurnBatch(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) ParseTriggerBurnBatch(log types.Log) (*IwstETHBurnerTriggerBurnBatch, error) {
	event := new(IwstETHBurnerTriggerBurnBatch)
	if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerBurnBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IwstETHBurnerTriggerWithdrawalIterator is returned from FilterTriggerWithdrawal and is used to iterate over the raw logs and unpacked data for TriggerWithdrawal events raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerWithdrawalIterator struct {
	Event *IwstETHBurnerTriggerWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IwstETHBurnerTriggerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IwstETHBurnerTriggerWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IwstETHBurnerTriggerWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IwstETHBurnerTriggerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IwstETHBurnerTriggerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IwstETHBurnerTriggerWithdrawal represents a TriggerWithdrawal event raised by the IwstETHBurner contract.
type IwstETHBurnerTriggerWithdrawal struct {
	Caller     common.Address
	RequestIds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTriggerWithdrawal is a free log retrieval operation binding the contract event 0xe60bc3f6bd772e2234b4831b9c71ac461b4afc653329fb877fba5853b724ae75.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) FilterTriggerWithdrawal(opts *bind.FilterOpts, caller []common.Address) (*IwstETHBurnerTriggerWithdrawalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.FilterLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return &IwstETHBurnerTriggerWithdrawalIterator{contract: _IwstETHBurner.contract, event: "TriggerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchTriggerWithdrawal is a free log subscription operation binding the contract event 0xe60bc3f6bd772e2234b4831b9c71ac461b4afc653329fb877fba5853b724ae75.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) WatchTriggerWithdrawal(opts *bind.WatchOpts, sink chan<- *IwstETHBurnerTriggerWithdrawal, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _IwstETHBurner.contract.WatchLogs(opts, "TriggerWithdrawal", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IwstETHBurnerTriggerWithdrawal)
				if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerWithdrawal is a log parse operation binding the contract event 0xe60bc3f6bd772e2234b4831b9c71ac461b4afc653329fb877fba5853b724ae75.
//
// Solidity: event TriggerWithdrawal(address indexed caller, uint256[] requestIds)
func (_IwstETHBurner *IwstETHBurnerFilterer) ParseTriggerWithdrawal(log types.Log) (*IwstETHBurnerTriggerWithdrawal, error) {
	event := new(IwstETHBurnerTriggerWithdrawal)
	if err := _IwstETHBurner.contract.UnpackLog(event, "TriggerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
