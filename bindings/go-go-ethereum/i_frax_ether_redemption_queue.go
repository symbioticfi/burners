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

// IFraxEtherRedemptionQueueMetaData contains all meta data concerning the IFraxEtherRedemptionQueue contract.
var IFraxEtherRedemptionQueueMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"burnRedemptionTicketNft\",\"inputs\":[{\"name\":\"_nftId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enterRedemptionQueueViaSfrxEth\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sfrxEthAmount\",\"type\":\"uint120\",\"internalType\":\"uint120\"}],\"outputs\":[{\"name\":\"_nftId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redemptionQueueState\",\"inputs\":[],\"outputs\":[{\"name\":\"nextNftId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"queueLengthSecs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"redemptionFee\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"earlyExitFee\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"}]",
}

// IFraxEtherRedemptionQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use IFraxEtherRedemptionQueueMetaData.ABI instead.
var IFraxEtherRedemptionQueueABI = IFraxEtherRedemptionQueueMetaData.ABI

// IFraxEtherRedemptionQueue is an auto generated Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueue struct {
	IFraxEtherRedemptionQueueCaller     // Read-only binding to the contract
	IFraxEtherRedemptionQueueTransactor // Write-only binding to the contract
	IFraxEtherRedemptionQueueFilterer   // Log filterer for contract events
}

// IFraxEtherRedemptionQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFraxEtherRedemptionQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFraxEtherRedemptionQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFraxEtherRedemptionQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFraxEtherRedemptionQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFraxEtherRedemptionQueueSession struct {
	Contract     *IFraxEtherRedemptionQueue // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IFraxEtherRedemptionQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFraxEtherRedemptionQueueCallerSession struct {
	Contract *IFraxEtherRedemptionQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IFraxEtherRedemptionQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFraxEtherRedemptionQueueTransactorSession struct {
	Contract     *IFraxEtherRedemptionQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IFraxEtherRedemptionQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueueRaw struct {
	Contract *IFraxEtherRedemptionQueue // Generic contract binding to access the raw methods on
}

// IFraxEtherRedemptionQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueueCallerRaw struct {
	Contract *IFraxEtherRedemptionQueueCaller // Generic read-only contract binding to access the raw methods on
}

// IFraxEtherRedemptionQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFraxEtherRedemptionQueueTransactorRaw struct {
	Contract *IFraxEtherRedemptionQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFraxEtherRedemptionQueue creates a new instance of IFraxEtherRedemptionQueue, bound to a specific deployed contract.
func NewIFraxEtherRedemptionQueue(address common.Address, backend bind.ContractBackend) (*IFraxEtherRedemptionQueue, error) {
	contract, err := bindIFraxEtherRedemptionQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFraxEtherRedemptionQueue{IFraxEtherRedemptionQueueCaller: IFraxEtherRedemptionQueueCaller{contract: contract}, IFraxEtherRedemptionQueueTransactor: IFraxEtherRedemptionQueueTransactor{contract: contract}, IFraxEtherRedemptionQueueFilterer: IFraxEtherRedemptionQueueFilterer{contract: contract}}, nil
}

// NewIFraxEtherRedemptionQueueCaller creates a new read-only instance of IFraxEtherRedemptionQueue, bound to a specific deployed contract.
func NewIFraxEtherRedemptionQueueCaller(address common.Address, caller bind.ContractCaller) (*IFraxEtherRedemptionQueueCaller, error) {
	contract, err := bindIFraxEtherRedemptionQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFraxEtherRedemptionQueueCaller{contract: contract}, nil
}

// NewIFraxEtherRedemptionQueueTransactor creates a new write-only instance of IFraxEtherRedemptionQueue, bound to a specific deployed contract.
func NewIFraxEtherRedemptionQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*IFraxEtherRedemptionQueueTransactor, error) {
	contract, err := bindIFraxEtherRedemptionQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFraxEtherRedemptionQueueTransactor{contract: contract}, nil
}

// NewIFraxEtherRedemptionQueueFilterer creates a new log filterer instance of IFraxEtherRedemptionQueue, bound to a specific deployed contract.
func NewIFraxEtherRedemptionQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*IFraxEtherRedemptionQueueFilterer, error) {
	contract, err := bindIFraxEtherRedemptionQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFraxEtherRedemptionQueueFilterer{contract: contract}, nil
}

// bindIFraxEtherRedemptionQueue binds a generic wrapper to an already deployed contract.
func bindIFraxEtherRedemptionQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFraxEtherRedemptionQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFraxEtherRedemptionQueue.Contract.IFraxEtherRedemptionQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.IFraxEtherRedemptionQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.IFraxEtherRedemptionQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFraxEtherRedemptionQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.contract.Transact(opts, method, params...)
}

// RedemptionQueueState is a free data retrieval call binding the contract method 0x1494ef63.
//
// Solidity: function redemptionQueueState() view returns(uint64 nextNftId, uint64 queueLengthSecs, uint64 redemptionFee, uint64 earlyExitFee)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueCaller) RedemptionQueueState(opts *bind.CallOpts) (struct {
	NextNftId       uint64
	QueueLengthSecs uint64
	RedemptionFee   uint64
	EarlyExitFee    uint64
}, error) {
	var out []interface{}
	err := _IFraxEtherRedemptionQueue.contract.Call(opts, &out, "redemptionQueueState")

	outstruct := new(struct {
		NextNftId       uint64
		QueueLengthSecs uint64
		RedemptionFee   uint64
		EarlyExitFee    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NextNftId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.QueueLengthSecs = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RedemptionFee = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EarlyExitFee = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// RedemptionQueueState is a free data retrieval call binding the contract method 0x1494ef63.
//
// Solidity: function redemptionQueueState() view returns(uint64 nextNftId, uint64 queueLengthSecs, uint64 redemptionFee, uint64 earlyExitFee)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueSession) RedemptionQueueState() (struct {
	NextNftId       uint64
	QueueLengthSecs uint64
	RedemptionFee   uint64
	EarlyExitFee    uint64
}, error) {
	return _IFraxEtherRedemptionQueue.Contract.RedemptionQueueState(&_IFraxEtherRedemptionQueue.CallOpts)
}

// RedemptionQueueState is a free data retrieval call binding the contract method 0x1494ef63.
//
// Solidity: function redemptionQueueState() view returns(uint64 nextNftId, uint64 queueLengthSecs, uint64 redemptionFee, uint64 earlyExitFee)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueCallerSession) RedemptionQueueState() (struct {
	NextNftId       uint64
	QueueLengthSecs uint64
	RedemptionFee   uint64
	EarlyExitFee    uint64
}, error) {
	return _IFraxEtherRedemptionQueue.Contract.RedemptionQueueState(&_IFraxEtherRedemptionQueue.CallOpts)
}

// BurnRedemptionTicketNft is a paid mutator transaction binding the contract method 0x0a5334e5.
//
// Solidity: function burnRedemptionTicketNft(uint256 _nftId, address _recipient) returns()
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactor) BurnRedemptionTicketNft(opts *bind.TransactOpts, _nftId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.contract.Transact(opts, "burnRedemptionTicketNft", _nftId, _recipient)
}

// BurnRedemptionTicketNft is a paid mutator transaction binding the contract method 0x0a5334e5.
//
// Solidity: function burnRedemptionTicketNft(uint256 _nftId, address _recipient) returns()
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueSession) BurnRedemptionTicketNft(_nftId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.BurnRedemptionTicketNft(&_IFraxEtherRedemptionQueue.TransactOpts, _nftId, _recipient)
}

// BurnRedemptionTicketNft is a paid mutator transaction binding the contract method 0x0a5334e5.
//
// Solidity: function burnRedemptionTicketNft(uint256 _nftId, address _recipient) returns()
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactorSession) BurnRedemptionTicketNft(_nftId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.BurnRedemptionTicketNft(&_IFraxEtherRedemptionQueue.TransactOpts, _nftId, _recipient)
}

// EnterRedemptionQueueViaSfrxEth is a paid mutator transaction binding the contract method 0xe14cbb8a.
//
// Solidity: function enterRedemptionQueueViaSfrxEth(address _recipient, uint120 _sfrxEthAmount) returns(uint256 _nftId)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactor) EnterRedemptionQueueViaSfrxEth(opts *bind.TransactOpts, _recipient common.Address, _sfrxEthAmount *big.Int) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.contract.Transact(opts, "enterRedemptionQueueViaSfrxEth", _recipient, _sfrxEthAmount)
}

// EnterRedemptionQueueViaSfrxEth is a paid mutator transaction binding the contract method 0xe14cbb8a.
//
// Solidity: function enterRedemptionQueueViaSfrxEth(address _recipient, uint120 _sfrxEthAmount) returns(uint256 _nftId)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueSession) EnterRedemptionQueueViaSfrxEth(_recipient common.Address, _sfrxEthAmount *big.Int) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.EnterRedemptionQueueViaSfrxEth(&_IFraxEtherRedemptionQueue.TransactOpts, _recipient, _sfrxEthAmount)
}

// EnterRedemptionQueueViaSfrxEth is a paid mutator transaction binding the contract method 0xe14cbb8a.
//
// Solidity: function enterRedemptionQueueViaSfrxEth(address _recipient, uint120 _sfrxEthAmount) returns(uint256 _nftId)
func (_IFraxEtherRedemptionQueue *IFraxEtherRedemptionQueueTransactorSession) EnterRedemptionQueueViaSfrxEth(_recipient common.Address, _sfrxEthAmount *big.Int) (*types.Transaction, error) {
	return _IFraxEtherRedemptionQueue.Contract.EnterRedemptionQueueViaSfrxEth(&_IFraxEtherRedemptionQueue.TransactOpts, _recipient, _sfrxEthAmount)
}
