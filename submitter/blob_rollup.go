// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BlobRollupBatchData is an auto generated low-level Go binding around an user-defined struct.
type BlobRollupBatchData struct {
	BlockNumber    uint64
	BlockWitness   []byte
	PreStateRoot   [32]byte
	PostStateRoot  [32]byte
	WithdrawalRoot [32]byte
	Signature      BlobRollupBatchSignature
}

// BlobRollupBatchSignature is an auto generated low-level Go binding around an user-defined struct.
type BlobRollupBatchSignature struct {
	Signers   [][]byte
	Signature []byte
}

// BlobRollupMetaData contains all meta data concerning the BlobRollup contract.
var BlobRollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"Challenge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchToVersionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"evaluation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txs\",\"type\":\"bytes\"}],\"name\":\"getPointValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashGetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blockWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"preStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBlobRollup.BatchSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structBlobRollup.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"storageBatchs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blockWitness\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"preStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBlobRollup.BatchSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structBlobRollup.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRoots\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlobRollupABI is the input ABI used to generate the binding from.
// Deprecated: Use BlobRollupMetaData.ABI instead.
var BlobRollupABI = BlobRollupMetaData.ABI

// BlobRollup is an auto generated Go binding around an Ethereum contract.
type BlobRollup struct {
	BlobRollupCaller     // Read-only binding to the contract
	BlobRollupTransactor // Write-only binding to the contract
	BlobRollupFilterer   // Log filterer for contract events
}

// BlobRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlobRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlobRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlobRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlobRollupSession struct {
	Contract     *BlobRollup       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlobRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlobRollupCallerSession struct {
	Contract *BlobRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlobRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlobRollupTransactorSession struct {
	Contract     *BlobRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlobRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlobRollupRaw struct {
	Contract *BlobRollup // Generic contract binding to access the raw methods on
}

// BlobRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlobRollupCallerRaw struct {
	Contract *BlobRollupCaller // Generic read-only contract binding to access the raw methods on
}

// BlobRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlobRollupTransactorRaw struct {
	Contract *BlobRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlobRollup creates a new instance of BlobRollup, bound to a specific deployed contract.
func NewBlobRollup(address common.Address, backend bind.ContractBackend) (*BlobRollup, error) {
	contract, err := bindBlobRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlobRollup{BlobRollupCaller: BlobRollupCaller{contract: contract}, BlobRollupTransactor: BlobRollupTransactor{contract: contract}, BlobRollupFilterer: BlobRollupFilterer{contract: contract}}, nil
}

// NewBlobRollupCaller creates a new read-only instance of BlobRollup, bound to a specific deployed contract.
func NewBlobRollupCaller(address common.Address, caller bind.ContractCaller) (*BlobRollupCaller, error) {
	contract, err := bindBlobRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobRollupCaller{contract: contract}, nil
}

// NewBlobRollupTransactor creates a new write-only instance of BlobRollup, bound to a specific deployed contract.
func NewBlobRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobRollupTransactor, error) {
	contract, err := bindBlobRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobRollupTransactor{contract: contract}, nil
}

// NewBlobRollupFilterer creates a new log filterer instance of BlobRollup, bound to a specific deployed contract.
func NewBlobRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobRollupFilterer, error) {
	contract, err := bindBlobRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobRollupFilterer{contract: contract}, nil
}

// bindBlobRollup binds a generic wrapper to an already deployed contract.
func bindBlobRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlobRollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlobRollup *BlobRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobRollup.Contract.BlobRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlobRollup *BlobRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobRollup.Contract.BlobRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlobRollup *BlobRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobRollup.Contract.BlobRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlobRollup *BlobRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlobRollup *BlobRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlobRollup *BlobRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobRollup.Contract.contract.Transact(opts, method, params...)
}

// BatchToVersionHash is a free data retrieval call binding the contract method 0xf4b10abc.
//
// Solidity: function batchToVersionHash(uint64 ) view returns(bytes32)
func (_BlobRollup *BlobRollupCaller) BatchToVersionHash(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "batchToVersionHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchToVersionHash is a free data retrieval call binding the contract method 0xf4b10abc.
//
// Solidity: function batchToVersionHash(uint64 ) view returns(bytes32)
func (_BlobRollup *BlobRollupSession) BatchToVersionHash(arg0 uint64) ([32]byte, error) {
	return _BlobRollup.Contract.BatchToVersionHash(&_BlobRollup.CallOpts, arg0)
}

// BatchToVersionHash is a free data retrieval call binding the contract method 0xf4b10abc.
//
// Solidity: function batchToVersionHash(uint64 ) view returns(bytes32)
func (_BlobRollup *BlobRollupCallerSession) BatchToVersionHash(arg0 uint64) ([32]byte, error) {
	return _BlobRollup.Contract.BatchToVersionHash(&_BlobRollup.CallOpts, arg0)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupCaller) DataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "dataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupSession) DataHash() ([32]byte, error) {
	return _BlobRollup.Contract.DataHash(&_BlobRollup.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupCallerSession) DataHash() ([32]byte, error) {
	return _BlobRollup.Contract.DataHash(&_BlobRollup.CallOpts)
}

// Evaluation is a free data retrieval call binding the contract method 0xe53c14c0.
//
// Solidity: function evaluation(bytes input) view returns(bytes _output)
func (_BlobRollup *BlobRollupCaller) Evaluation(opts *bind.CallOpts, input []byte) ([]byte, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "evaluation", input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Evaluation is a free data retrieval call binding the contract method 0xe53c14c0.
//
// Solidity: function evaluation(bytes input) view returns(bytes _output)
func (_BlobRollup *BlobRollupSession) Evaluation(input []byte) ([]byte, error) {
	return _BlobRollup.Contract.Evaluation(&_BlobRollup.CallOpts, input)
}

// Evaluation is a free data retrieval call binding the contract method 0xe53c14c0.
//
// Solidity: function evaluation(bytes input) view returns(bytes _output)
func (_BlobRollup *BlobRollupCallerSession) Evaluation(input []byte) ([]byte, error) {
	return _BlobRollup.Contract.Evaluation(&_BlobRollup.CallOpts, input)
}

// GetDataHash is a free data retrieval call binding the contract method 0xa4da2290.
//
// Solidity: function getDataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupCaller) GetDataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "getDataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDataHash is a free data retrieval call binding the contract method 0xa4da2290.
//
// Solidity: function getDataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupSession) GetDataHash() ([32]byte, error) {
	return _BlobRollup.Contract.GetDataHash(&_BlobRollup.CallOpts)
}

// GetDataHash is a free data retrieval call binding the contract method 0xa4da2290.
//
// Solidity: function getDataHash() view returns(bytes32)
func (_BlobRollup *BlobRollupCallerSession) GetDataHash() ([32]byte, error) {
	return _BlobRollup.Contract.GetDataHash(&_BlobRollup.CallOpts)
}

// GetPointValue is a free data retrieval call binding the contract method 0xc5923e43.
//
// Solidity: function getPointValue(bytes txs) pure returns(uint256)
func (_BlobRollup *BlobRollupCaller) GetPointValue(opts *bind.CallOpts, txs []byte) (*big.Int, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "getPointValue", txs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPointValue is a free data retrieval call binding the contract method 0xc5923e43.
//
// Solidity: function getPointValue(bytes txs) pure returns(uint256)
func (_BlobRollup *BlobRollupSession) GetPointValue(txs []byte) (*big.Int, error) {
	return _BlobRollup.Contract.GetPointValue(&_BlobRollup.CallOpts, txs)
}

// GetPointValue is a free data retrieval call binding the contract method 0xc5923e43.
//
// Solidity: function getPointValue(bytes txs) pure returns(uint256)
func (_BlobRollup *BlobRollupCallerSession) GetPointValue(txs []byte) (*big.Int, error) {
	return _BlobRollup.Contract.GetPointValue(&_BlobRollup.CallOpts, txs)
}

// HashGetter is a free data retrieval call binding the contract method 0xe5ec2a65.
//
// Solidity: function hashGetter() view returns(address)
func (_BlobRollup *BlobRollupCaller) HashGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "hashGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashGetter is a free data retrieval call binding the contract method 0xe5ec2a65.
//
// Solidity: function hashGetter() view returns(address)
func (_BlobRollup *BlobRollupSession) HashGetter() (common.Address, error) {
	return _BlobRollup.Contract.HashGetter(&_BlobRollup.CallOpts)
}

// HashGetter is a free data retrieval call binding the contract method 0xe5ec2a65.
//
// Solidity: function hashGetter() view returns(address)
func (_BlobRollup *BlobRollupCallerSession) HashGetter() (common.Address, error) {
	return _BlobRollup.Contract.HashGetter(&_BlobRollup.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_BlobRollup *BlobRollupCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_BlobRollup *BlobRollupSession) LastBatchSequenced() (uint64, error) {
	return _BlobRollup.Contract.LastBatchSequenced(&_BlobRollup.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_BlobRollup *BlobRollupCallerSession) LastBatchSequenced() (uint64, error) {
	return _BlobRollup.Contract.LastBatchSequenced(&_BlobRollup.CallOpts)
}

// LastL2BlockNumber is a free data retrieval call binding the contract method 0xf02deda1.
//
// Solidity: function lastL2BlockNumber() view returns(uint64)
func (_BlobRollup *BlobRollupCaller) LastL2BlockNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "lastL2BlockNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastL2BlockNumber is a free data retrieval call binding the contract method 0xf02deda1.
//
// Solidity: function lastL2BlockNumber() view returns(uint64)
func (_BlobRollup *BlobRollupSession) LastL2BlockNumber() (uint64, error) {
	return _BlobRollup.Contract.LastL2BlockNumber(&_BlobRollup.CallOpts)
}

// LastL2BlockNumber is a free data retrieval call binding the contract method 0xf02deda1.
//
// Solidity: function lastL2BlockNumber() view returns(uint64)
func (_BlobRollup *BlobRollupCallerSession) LastL2BlockNumber() (uint64, error) {
	return _BlobRollup.Contract.LastL2BlockNumber(&_BlobRollup.CallOpts)
}

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(uint64 blockNumber, bytes32 withdrawalRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_BlobRollup *BlobRollupCaller) StorageBatchs(opts *bind.CallOpts, arg0 uint64) (struct {
	BlockNumber     uint64
	WithdrawalRoot  [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "storageBatchs", arg0)

	outstruct := new(struct {
		BlockNumber     uint64
		WithdrawalRoot  [32]byte
		Commitment      [32]byte
		StateRoot       [32]byte
		OriginTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.WithdrawalRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Commitment = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.OriginTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(uint64 blockNumber, bytes32 withdrawalRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_BlobRollup *BlobRollupSession) StorageBatchs(arg0 uint64) (struct {
	BlockNumber     uint64
	WithdrawalRoot  [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	return _BlobRollup.Contract.StorageBatchs(&_BlobRollup.CallOpts, arg0)
}

// StorageBatchs is a free data retrieval call binding the contract method 0xe291c2f5.
//
// Solidity: function storageBatchs(uint64 ) view returns(uint64 blockNumber, bytes32 withdrawalRoot, bytes32 commitment, bytes32 stateRoot, uint256 originTimestamp)
func (_BlobRollup *BlobRollupCallerSession) StorageBatchs(arg0 uint64) (struct {
	BlockNumber     uint64
	WithdrawalRoot  [32]byte
	Commitment      [32]byte
	StateRoot       [32]byte
	OriginTimestamp *big.Int
}, error) {
	return _BlobRollup.Contract.StorageBatchs(&_BlobRollup.CallOpts, arg0)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(uint64 batchIndex, uint256 timestamp)
func (_BlobRollup *BlobRollupCaller) WithdrawalRoots(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BatchIndex uint64
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _BlobRollup.contract.Call(opts, &out, "withdrawalRoots", arg0)

	outstruct := new(struct {
		BatchIndex uint64
		Timestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(uint64 batchIndex, uint256 timestamp)
func (_BlobRollup *BlobRollupSession) WithdrawalRoots(arg0 [32]byte) (struct {
	BatchIndex uint64
	Timestamp  *big.Int
}, error) {
	return _BlobRollup.Contract.WithdrawalRoots(&_BlobRollup.CallOpts, arg0)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(uint64 batchIndex, uint256 timestamp)
func (_BlobRollup *BlobRollupCallerSession) WithdrawalRoots(arg0 [32]byte) (struct {
	BatchIndex uint64
	Timestamp  *big.Int
}, error) {
	return _BlobRollup.Contract.WithdrawalRoots(&_BlobRollup.CallOpts, arg0)
}

// Challenge is a paid mutator transaction binding the contract method 0xeff6fc12.
//
// Solidity: function challenge(uint64 batchNum) returns()
func (_BlobRollup *BlobRollupTransactor) Challenge(opts *bind.TransactOpts, batchNum uint64) (*types.Transaction, error) {
	return _BlobRollup.contract.Transact(opts, "challenge", batchNum)
}

// Challenge is a paid mutator transaction binding the contract method 0xeff6fc12.
//
// Solidity: function challenge(uint64 batchNum) returns()
func (_BlobRollup *BlobRollupSession) Challenge(batchNum uint64) (*types.Transaction, error) {
	return _BlobRollup.Contract.Challenge(&_BlobRollup.TransactOpts, batchNum)
}

// Challenge is a paid mutator transaction binding the contract method 0xeff6fc12.
//
// Solidity: function challenge(uint64 batchNum) returns()
func (_BlobRollup *BlobRollupTransactorSession) Challenge(batchNum uint64) (*types.Transaction, error) {
	return _BlobRollup.Contract.Challenge(&_BlobRollup.TransactOpts, batchNum)
}

// Prove is a paid mutator transaction binding the contract method 0x1a00cab1.
//
// Solidity: function prove((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupTransactor) Prove(opts *bind.TransactOpts, batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.contract.Transact(opts, "prove", batches)
}

// Prove is a paid mutator transaction binding the contract method 0x1a00cab1.
//
// Solidity: function prove((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupSession) Prove(batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.Contract.Prove(&_BlobRollup.TransactOpts, batches)
}

// Prove is a paid mutator transaction binding the contract method 0x1a00cab1.
//
// Solidity: function prove((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupTransactorSession) Prove(batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.Contract.Prove(&_BlobRollup.TransactOpts, batches)
}

// Submit is a paid mutator transaction binding the contract method 0x0f30b953.
//
// Solidity: function submit((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupTransactor) Submit(opts *bind.TransactOpts, batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.contract.Transact(opts, "submit", batches)
}

// Submit is a paid mutator transaction binding the contract method 0x0f30b953.
//
// Solidity: function submit((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupSession) Submit(batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.Contract.Submit(&_BlobRollup.TransactOpts, batches)
}

// Submit is a paid mutator transaction binding the contract method 0x0f30b953.
//
// Solidity: function submit((uint64,bytes,bytes32,bytes32,bytes32,(bytes[],bytes))[] batches) returns()
func (_BlobRollup *BlobRollupTransactorSession) Submit(batches []BlobRollupBatchData) (*types.Transaction, error) {
	return _BlobRollup.Contract.Submit(&_BlobRollup.TransactOpts, batches)
}

// BlobRollupChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the BlobRollup contract.
type BlobRollupChallengeIterator struct {
	Event *BlobRollupChallenge // Event containing the contract specifics and raw log

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
func (it *BlobRollupChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobRollupChallenge)
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
		it.Event = new(BlobRollupChallenge)
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
func (it *BlobRollupChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlobRollupChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlobRollupChallenge represents a Challenge event raised by the BlobRollup contract.
type BlobRollupChallenge struct {
	BatchNum uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x7da0610d1319f5b4847cddae888b1c8022030c5b06b912f67d7c8cd6ca8d39cc.
//
// Solidity: event Challenge(uint64 batchNum)
func (_BlobRollup *BlobRollupFilterer) FilterChallenge(opts *bind.FilterOpts) (*BlobRollupChallengeIterator, error) {

	logs, sub, err := _BlobRollup.contract.FilterLogs(opts, "Challenge")
	if err != nil {
		return nil, err
	}
	return &BlobRollupChallengeIterator{contract: _BlobRollup.contract, event: "Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x7da0610d1319f5b4847cddae888b1c8022030c5b06b912f67d7c8cd6ca8d39cc.
//
// Solidity: event Challenge(uint64 batchNum)
func (_BlobRollup *BlobRollupFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *BlobRollupChallenge) (event.Subscription, error) {

	logs, sub, err := _BlobRollup.contract.WatchLogs(opts, "Challenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlobRollupChallenge)
				if err := _BlobRollup.contract.UnpackLog(event, "Challenge", log); err != nil {
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

// ParseChallenge is a log parse operation binding the contract event 0x7da0610d1319f5b4847cddae888b1c8022030c5b06b912f67d7c8cd6ca8d39cc.
//
// Solidity: event Challenge(uint64 batchNum)
func (_BlobRollup *BlobRollupFilterer) ParseChallenge(log types.Log) (*BlobRollupChallenge, error) {
	event := new(BlobRollupChallenge)
	if err := _BlobRollup.contract.UnpackLog(event, "Challenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
