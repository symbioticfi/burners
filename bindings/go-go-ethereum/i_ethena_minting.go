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

// IEthenaMintingBlockTotals is an auto generated low-level Go binding around an user-defined struct.
type IEthenaMintingBlockTotals struct {
	MintedPerBlock   *big.Int
	RedeemedPerBlock *big.Int
}

// IEthenaMintingGlobalConfig is an auto generated low-level Go binding around an user-defined struct.
type IEthenaMintingGlobalConfig struct {
	GlobalMaxMintPerBlock   *big.Int
	GlobalMaxRedeemPerBlock *big.Int
}

// IEthenaMintingOrder is an auto generated low-level Go binding around an user-defined struct.
type IEthenaMintingOrder struct {
	OrderId          string
	OrderType        uint8
	Expiry           *big.Int
	Nonce            *big.Int
	Benefactor       common.Address
	Beneficiary      common.Address
	CollateralAsset  common.Address
	CollateralAmount *big.Int
	UsdeAmount       *big.Int
}

// IEthenaMintingSignature is an auto generated low-level Go binding around an user-defined struct.
type IEthenaMintingSignature struct {
	SignatureType  uint8
	SignatureBytes []byte
}

// IEthenaMintingTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type IEthenaMintingTokenConfig struct {
	TokenType         uint8
	IsActive          bool
	MaxMintPerBlock   *big.Int
	MaxRedeemPerBlock *big.Int
}

// IEthenaMintingMetaData contains all meta data concerning the IEthenaMinting contract.
var IEthenaMintingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addWhitelistedBenefactor\",\"inputs\":[{\"name\":\"benefactor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"globalConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.GlobalConfig\",\"components\":[{\"name\":\"globalMaxMintPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"globalMaxRedeemPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.Order\",\"components\":[{\"name\":\"order_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"order_type\",\"type\":\"uint8\",\"internalType\":\"enumIEthenaMinting.OrderType\"},{\"name\":\"expiry\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"nonce\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"benefactor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral_asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral_amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"usde_amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.Order\",\"components\":[{\"name\":\"order_id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"order_type\",\"type\":\"uint8\",\"internalType\":\"enumIEthenaMinting.OrderType\"},{\"name\":\"expiry\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"nonce\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"benefactor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral_asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral_amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"usde_amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.Signature\",\"components\":[{\"name\":\"signature_type\",\"type\":\"uint8\",\"internalType\":\"enumIEthenaMinting.SignatureType\"},{\"name\":\"signature_bytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenConfig\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.TokenConfig\",\"components\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumIEthenaMinting.TokenType\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"maxMintPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"maxRedeemPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPerBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.BlockTotals\",\"components\":[{\"name\":\"mintedPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"redeemedPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalPerBlockPerAsset\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEthenaMinting.BlockTotals\",\"components\":[{\"name\":\"mintedPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"redeemedPerBlock\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"}]",
}

// IEthenaMintingABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthenaMintingMetaData.ABI instead.
var IEthenaMintingABI = IEthenaMintingMetaData.ABI

// IEthenaMinting is an auto generated Go binding around an Ethereum contract.
type IEthenaMinting struct {
	IEthenaMintingCaller     // Read-only binding to the contract
	IEthenaMintingTransactor // Write-only binding to the contract
	IEthenaMintingFilterer   // Log filterer for contract events
}

// IEthenaMintingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthenaMintingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthenaMintingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthenaMintingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthenaMintingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthenaMintingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthenaMintingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthenaMintingSession struct {
	Contract     *IEthenaMinting   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthenaMintingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthenaMintingCallerSession struct {
	Contract *IEthenaMintingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IEthenaMintingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthenaMintingTransactorSession struct {
	Contract     *IEthenaMintingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IEthenaMintingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthenaMintingRaw struct {
	Contract *IEthenaMinting // Generic contract binding to access the raw methods on
}

// IEthenaMintingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthenaMintingCallerRaw struct {
	Contract *IEthenaMintingCaller // Generic read-only contract binding to access the raw methods on
}

// IEthenaMintingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthenaMintingTransactorRaw struct {
	Contract *IEthenaMintingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthenaMinting creates a new instance of IEthenaMinting, bound to a specific deployed contract.
func NewIEthenaMinting(address common.Address, backend bind.ContractBackend) (*IEthenaMinting, error) {
	contract, err := bindIEthenaMinting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthenaMinting{IEthenaMintingCaller: IEthenaMintingCaller{contract: contract}, IEthenaMintingTransactor: IEthenaMintingTransactor{contract: contract}, IEthenaMintingFilterer: IEthenaMintingFilterer{contract: contract}}, nil
}

// NewIEthenaMintingCaller creates a new read-only instance of IEthenaMinting, bound to a specific deployed contract.
func NewIEthenaMintingCaller(address common.Address, caller bind.ContractCaller) (*IEthenaMintingCaller, error) {
	contract, err := bindIEthenaMinting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthenaMintingCaller{contract: contract}, nil
}

// NewIEthenaMintingTransactor creates a new write-only instance of IEthenaMinting, bound to a specific deployed contract.
func NewIEthenaMintingTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthenaMintingTransactor, error) {
	contract, err := bindIEthenaMinting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthenaMintingTransactor{contract: contract}, nil
}

// NewIEthenaMintingFilterer creates a new log filterer instance of IEthenaMinting, bound to a specific deployed contract.
func NewIEthenaMintingFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthenaMintingFilterer, error) {
	contract, err := bindIEthenaMinting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthenaMintingFilterer{contract: contract}, nil
}

// bindIEthenaMinting binds a generic wrapper to an already deployed contract.
func bindIEthenaMinting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthenaMintingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthenaMinting *IEthenaMintingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthenaMinting.Contract.IEthenaMintingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthenaMinting *IEthenaMintingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.IEthenaMintingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthenaMinting *IEthenaMintingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.IEthenaMintingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthenaMinting *IEthenaMintingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthenaMinting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthenaMinting *IEthenaMintingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthenaMinting *IEthenaMintingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.contract.Transact(opts, method, params...)
}

// GlobalConfig is a free data retrieval call binding the contract method 0xa7c1abe0.
//
// Solidity: function globalConfig() view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCaller) GlobalConfig(opts *bind.CallOpts) (IEthenaMintingGlobalConfig, error) {
	var out []interface{}
	err := _IEthenaMinting.contract.Call(opts, &out, "globalConfig")

	if err != nil {
		return *new(IEthenaMintingGlobalConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEthenaMintingGlobalConfig)).(*IEthenaMintingGlobalConfig)

	return out0, err

}

// GlobalConfig is a free data retrieval call binding the contract method 0xa7c1abe0.
//
// Solidity: function globalConfig() view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingSession) GlobalConfig() (IEthenaMintingGlobalConfig, error) {
	return _IEthenaMinting.Contract.GlobalConfig(&_IEthenaMinting.CallOpts)
}

// GlobalConfig is a free data retrieval call binding the contract method 0xa7c1abe0.
//
// Solidity: function globalConfig() view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCallerSession) GlobalConfig() (IEthenaMintingGlobalConfig, error) {
	return _IEthenaMinting.Contract.GlobalConfig(&_IEthenaMinting.CallOpts)
}

// HashOrder is a free data retrieval call binding the contract method 0x7cef7e91.
//
// Solidity: function hashOrder((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order) view returns(bytes32)
func (_IEthenaMinting *IEthenaMintingCaller) HashOrder(opts *bind.CallOpts, order IEthenaMintingOrder) ([32]byte, error) {
	var out []interface{}
	err := _IEthenaMinting.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x7cef7e91.
//
// Solidity: function hashOrder((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order) view returns(bytes32)
func (_IEthenaMinting *IEthenaMintingSession) HashOrder(order IEthenaMintingOrder) ([32]byte, error) {
	return _IEthenaMinting.Contract.HashOrder(&_IEthenaMinting.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x7cef7e91.
//
// Solidity: function hashOrder((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order) view returns(bytes32)
func (_IEthenaMinting *IEthenaMintingCallerSession) HashOrder(order IEthenaMintingOrder) ([32]byte, error) {
	return _IEthenaMinting.Contract.HashOrder(&_IEthenaMinting.CallOpts, order)
}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address asset) view returns((uint8,bool,uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCaller) TokenConfig(opts *bind.CallOpts, asset common.Address) (IEthenaMintingTokenConfig, error) {
	var out []interface{}
	err := _IEthenaMinting.contract.Call(opts, &out, "tokenConfig", asset)

	if err != nil {
		return *new(IEthenaMintingTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEthenaMintingTokenConfig)).(*IEthenaMintingTokenConfig)

	return out0, err

}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address asset) view returns((uint8,bool,uint128,uint128))
func (_IEthenaMinting *IEthenaMintingSession) TokenConfig(asset common.Address) (IEthenaMintingTokenConfig, error) {
	return _IEthenaMinting.Contract.TokenConfig(&_IEthenaMinting.CallOpts, asset)
}

// TokenConfig is a free data retrieval call binding the contract method 0xfe136c4e.
//
// Solidity: function tokenConfig(address asset) view returns((uint8,bool,uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCallerSession) TokenConfig(asset common.Address) (IEthenaMintingTokenConfig, error) {
	return _IEthenaMinting.Contract.TokenConfig(&_IEthenaMinting.CallOpts, asset)
}

// TotalPerBlock is a free data retrieval call binding the contract method 0x92408dc7.
//
// Solidity: function totalPerBlock(uint256 blockNumber) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCaller) TotalPerBlock(opts *bind.CallOpts, blockNumber *big.Int) (IEthenaMintingBlockTotals, error) {
	var out []interface{}
	err := _IEthenaMinting.contract.Call(opts, &out, "totalPerBlock", blockNumber)

	if err != nil {
		return *new(IEthenaMintingBlockTotals), err
	}

	out0 := *abi.ConvertType(out[0], new(IEthenaMintingBlockTotals)).(*IEthenaMintingBlockTotals)

	return out0, err

}

// TotalPerBlock is a free data retrieval call binding the contract method 0x92408dc7.
//
// Solidity: function totalPerBlock(uint256 blockNumber) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingSession) TotalPerBlock(blockNumber *big.Int) (IEthenaMintingBlockTotals, error) {
	return _IEthenaMinting.Contract.TotalPerBlock(&_IEthenaMinting.CallOpts, blockNumber)
}

// TotalPerBlock is a free data retrieval call binding the contract method 0x92408dc7.
//
// Solidity: function totalPerBlock(uint256 blockNumber) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCallerSession) TotalPerBlock(blockNumber *big.Int) (IEthenaMintingBlockTotals, error) {
	return _IEthenaMinting.Contract.TotalPerBlock(&_IEthenaMinting.CallOpts, blockNumber)
}

// TotalPerBlockPerAsset is a free data retrieval call binding the contract method 0xa693635e.
//
// Solidity: function totalPerBlockPerAsset(uint256 blockNumber, address asset) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCaller) TotalPerBlockPerAsset(opts *bind.CallOpts, blockNumber *big.Int, asset common.Address) (IEthenaMintingBlockTotals, error) {
	var out []interface{}
	err := _IEthenaMinting.contract.Call(opts, &out, "totalPerBlockPerAsset", blockNumber, asset)

	if err != nil {
		return *new(IEthenaMintingBlockTotals), err
	}

	out0 := *abi.ConvertType(out[0], new(IEthenaMintingBlockTotals)).(*IEthenaMintingBlockTotals)

	return out0, err

}

// TotalPerBlockPerAsset is a free data retrieval call binding the contract method 0xa693635e.
//
// Solidity: function totalPerBlockPerAsset(uint256 blockNumber, address asset) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingSession) TotalPerBlockPerAsset(blockNumber *big.Int, asset common.Address) (IEthenaMintingBlockTotals, error) {
	return _IEthenaMinting.Contract.TotalPerBlockPerAsset(&_IEthenaMinting.CallOpts, blockNumber, asset)
}

// TotalPerBlockPerAsset is a free data retrieval call binding the contract method 0xa693635e.
//
// Solidity: function totalPerBlockPerAsset(uint256 blockNumber, address asset) view returns((uint128,uint128))
func (_IEthenaMinting *IEthenaMintingCallerSession) TotalPerBlockPerAsset(blockNumber *big.Int, asset common.Address) (IEthenaMintingBlockTotals, error) {
	return _IEthenaMinting.Contract.TotalPerBlockPerAsset(&_IEthenaMinting.CallOpts, blockNumber, asset)
}

// AddWhitelistedBenefactor is a paid mutator transaction binding the contract method 0x16255c43.
//
// Solidity: function addWhitelistedBenefactor(address benefactor) returns()
func (_IEthenaMinting *IEthenaMintingTransactor) AddWhitelistedBenefactor(opts *bind.TransactOpts, benefactor common.Address) (*types.Transaction, error) {
	return _IEthenaMinting.contract.Transact(opts, "addWhitelistedBenefactor", benefactor)
}

// AddWhitelistedBenefactor is a paid mutator transaction binding the contract method 0x16255c43.
//
// Solidity: function addWhitelistedBenefactor(address benefactor) returns()
func (_IEthenaMinting *IEthenaMintingSession) AddWhitelistedBenefactor(benefactor common.Address) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.AddWhitelistedBenefactor(&_IEthenaMinting.TransactOpts, benefactor)
}

// AddWhitelistedBenefactor is a paid mutator transaction binding the contract method 0x16255c43.
//
// Solidity: function addWhitelistedBenefactor(address benefactor) returns()
func (_IEthenaMinting *IEthenaMintingTransactorSession) AddWhitelistedBenefactor(benefactor common.Address) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.AddWhitelistedBenefactor(&_IEthenaMinting.TransactOpts, benefactor)
}

// Redeem is a paid mutator transaction binding the contract method 0x75c890dc.
//
// Solidity: function redeem((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order, (uint8,bytes) signature) returns()
func (_IEthenaMinting *IEthenaMintingTransactor) Redeem(opts *bind.TransactOpts, order IEthenaMintingOrder, signature IEthenaMintingSignature) (*types.Transaction, error) {
	return _IEthenaMinting.contract.Transact(opts, "redeem", order, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0x75c890dc.
//
// Solidity: function redeem((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order, (uint8,bytes) signature) returns()
func (_IEthenaMinting *IEthenaMintingSession) Redeem(order IEthenaMintingOrder, signature IEthenaMintingSignature) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.Redeem(&_IEthenaMinting.TransactOpts, order, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0x75c890dc.
//
// Solidity: function redeem((string,uint8,uint120,uint128,address,address,address,uint128,uint128) order, (uint8,bytes) signature) returns()
func (_IEthenaMinting *IEthenaMintingTransactorSession) Redeem(order IEthenaMintingOrder, signature IEthenaMintingSignature) (*types.Transaction, error) {
	return _IEthenaMinting.Contract.Redeem(&_IEthenaMinting.TransactOpts, order, signature)
}
