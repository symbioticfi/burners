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

// IWithdrawalQueueMetaData contains all meta data concerning the IWithdrawalQueue contract.
var IWithdrawalQueueMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_STETH_WITHDRAWAL_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_STETH_WITHDRAWAL_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimWithdrawal\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimWithdrawals\",\"inputs\":[{\"name\":\"_requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_hints\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[{\"name\":\"_lastRequestIdToBeFinalized\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxShareRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"findCheckpointHints\",\"inputs\":[{\"name\":\"_requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_firstIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lastIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"hintIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestWithdrawals\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IWithdrawalQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use IWithdrawalQueueMetaData.ABI instead.
var IWithdrawalQueueABI = IWithdrawalQueueMetaData.ABI

// IWithdrawalQueue is an auto generated Go binding around an Ethereum contract.
type IWithdrawalQueue struct {
	IWithdrawalQueueCaller     // Read-only binding to the contract
	IWithdrawalQueueTransactor // Write-only binding to the contract
	IWithdrawalQueueFilterer   // Log filterer for contract events
}

// IWithdrawalQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWithdrawalQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawalQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWithdrawalQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawalQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWithdrawalQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWithdrawalQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWithdrawalQueueSession struct {
	Contract     *IWithdrawalQueue // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWithdrawalQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWithdrawalQueueCallerSession struct {
	Contract *IWithdrawalQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IWithdrawalQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWithdrawalQueueTransactorSession struct {
	Contract     *IWithdrawalQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IWithdrawalQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWithdrawalQueueRaw struct {
	Contract *IWithdrawalQueue // Generic contract binding to access the raw methods on
}

// IWithdrawalQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWithdrawalQueueCallerRaw struct {
	Contract *IWithdrawalQueueCaller // Generic read-only contract binding to access the raw methods on
}

// IWithdrawalQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWithdrawalQueueTransactorRaw struct {
	Contract *IWithdrawalQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWithdrawalQueue creates a new instance of IWithdrawalQueue, bound to a specific deployed contract.
func NewIWithdrawalQueue(address common.Address, backend bind.ContractBackend) (*IWithdrawalQueue, error) {
	contract, err := bindIWithdrawalQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWithdrawalQueue{IWithdrawalQueueCaller: IWithdrawalQueueCaller{contract: contract}, IWithdrawalQueueTransactor: IWithdrawalQueueTransactor{contract: contract}, IWithdrawalQueueFilterer: IWithdrawalQueueFilterer{contract: contract}}, nil
}

// NewIWithdrawalQueueCaller creates a new read-only instance of IWithdrawalQueue, bound to a specific deployed contract.
func NewIWithdrawalQueueCaller(address common.Address, caller bind.ContractCaller) (*IWithdrawalQueueCaller, error) {
	contract, err := bindIWithdrawalQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWithdrawalQueueCaller{contract: contract}, nil
}

// NewIWithdrawalQueueTransactor creates a new write-only instance of IWithdrawalQueue, bound to a specific deployed contract.
func NewIWithdrawalQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*IWithdrawalQueueTransactor, error) {
	contract, err := bindIWithdrawalQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWithdrawalQueueTransactor{contract: contract}, nil
}

// NewIWithdrawalQueueFilterer creates a new log filterer instance of IWithdrawalQueue, bound to a specific deployed contract.
func NewIWithdrawalQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*IWithdrawalQueueFilterer, error) {
	contract, err := bindIWithdrawalQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWithdrawalQueueFilterer{contract: contract}, nil
}

// bindIWithdrawalQueue binds a generic wrapper to an already deployed contract.
func bindIWithdrawalQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWithdrawalQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWithdrawalQueue *IWithdrawalQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWithdrawalQueue.Contract.IWithdrawalQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWithdrawalQueue *IWithdrawalQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.IWithdrawalQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWithdrawalQueue *IWithdrawalQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.IWithdrawalQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWithdrawalQueue *IWithdrawalQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWithdrawalQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWithdrawalQueue *IWithdrawalQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWithdrawalQueue *IWithdrawalQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.contract.Transact(opts, method, params...)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) MAXSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "MAX_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.MAXSTETHWITHDRAWALAMOUNT(&_IWithdrawalQueue.CallOpts)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.MAXSTETHWITHDRAWALAMOUNT(&_IWithdrawalQueue.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) MINSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "MIN_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.MINSTETHWITHDRAWALAMOUNT(&_IWithdrawalQueue.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.MINSTETHWITHDRAWALAMOUNT(&_IWithdrawalQueue.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IWithdrawalQueue *IWithdrawalQueueSession) STETH() (common.Address, error) {
	return _IWithdrawalQueue.Contract.STETH(&_IWithdrawalQueue.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) STETH() (common.Address, error) {
	return _IWithdrawalQueue.Contract.STETH(&_IWithdrawalQueue.CallOpts)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) FindCheckpointHints(opts *bind.CallOpts, _requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "findCheckpointHints", _requestIds, _firstIndex, _lastIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_IWithdrawalQueue *IWithdrawalQueueSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _IWithdrawalQueue.Contract.FindCheckpointHints(&_IWithdrawalQueue.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _IWithdrawalQueue.Contract.FindCheckpointHints(&_IWithdrawalQueue.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) GetLastCheckpointIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "getLastCheckpointIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueSession) GetLastCheckpointIndex() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.GetLastCheckpointIndex(&_IWithdrawalQueue.CallOpts)
}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) GetLastCheckpointIndex() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.GetLastCheckpointIndex(&_IWithdrawalQueue.CallOpts)
}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCaller) GetLastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWithdrawalQueue.contract.Call(opts, &out, "getLastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueSession) GetLastRequestId() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.GetLastRequestId(&_IWithdrawalQueue.CallOpts)
}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_IWithdrawalQueue *IWithdrawalQueueCallerSession) GetLastRequestId() (*big.Int, error) {
	return _IWithdrawalQueue.Contract.GetLastRequestId(&_IWithdrawalQueue.CallOpts)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactor) ClaimWithdrawal(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.contract.Transact(opts, "claimWithdrawal", _requestId)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_IWithdrawalQueue *IWithdrawalQueueSession) ClaimWithdrawal(_requestId *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.ClaimWithdrawal(&_IWithdrawalQueue.TransactOpts, _requestId)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactorSession) ClaimWithdrawal(_requestId *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.ClaimWithdrawal(&_IWithdrawalQueue.TransactOpts, _requestId)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactor) ClaimWithdrawals(opts *bind.TransactOpts, _requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.contract.Transact(opts, "claimWithdrawals", _requestIds, _hints)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_IWithdrawalQueue *IWithdrawalQueueSession) ClaimWithdrawals(_requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.ClaimWithdrawals(&_IWithdrawalQueue.TransactOpts, _requestIds, _hints)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactorSession) ClaimWithdrawals(_requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.ClaimWithdrawals(&_IWithdrawalQueue.TransactOpts, _requestIds, _hints)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactor) Finalize(opts *bind.TransactOpts, _lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.contract.Transact(opts, "finalize", _lastRequestIdToBeFinalized, _maxShareRate)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_IWithdrawalQueue *IWithdrawalQueueSession) Finalize(_lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.Finalize(&_IWithdrawalQueue.TransactOpts, _lastRequestIdToBeFinalized, _maxShareRate)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_IWithdrawalQueue *IWithdrawalQueueTransactorSession) Finalize(_lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.Finalize(&_IWithdrawalQueue.TransactOpts, _lastRequestIdToBeFinalized, _maxShareRate)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_IWithdrawalQueue *IWithdrawalQueueTransactor) RequestWithdrawals(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IWithdrawalQueue.contract.Transact(opts, "requestWithdrawals", _amounts, _owner)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_IWithdrawalQueue *IWithdrawalQueueSession) RequestWithdrawals(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.RequestWithdrawals(&_IWithdrawalQueue.TransactOpts, _amounts, _owner)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_IWithdrawalQueue *IWithdrawalQueueTransactorSession) RequestWithdrawals(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _IWithdrawalQueue.Contract.RequestWithdrawals(&_IWithdrawalQueue.TransactOpts, _amounts, _owner)
}
