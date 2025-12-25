// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AuctionContractMetaData contains all meta data concerning the AuctionContract contract.
var AuctionContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdcToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdcFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ethFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"calculateTotalPrice\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"_nft\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_durationMinutes\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentBuyer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"durationMinutes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethMap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPriceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestETHPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUSDCPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"seller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumNftAuctionV2.AuctionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdcMap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcPriceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdcToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withDraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionEnd\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidUSD\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateAuction\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nft\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"durationMinutes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// AuctionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionContractMetaData.ABI instead.
var AuctionContractABI = AuctionContractMetaData.ABI

// AuctionContract is an auto generated Go binding around an Ethereum contract.
type AuctionContract struct {
	AuctionContractCaller     // Read-only binding to the contract
	AuctionContractTransactor // Write-only binding to the contract
	AuctionContractFilterer   // Log filterer for contract events
}

// AuctionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionContractSession struct {
	Contract     *AuctionContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionContractCallerSession struct {
	Contract *AuctionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AuctionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionContractTransactorSession struct {
	Contract     *AuctionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AuctionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionContractRaw struct {
	Contract *AuctionContract // Generic contract binding to access the raw methods on
}

// AuctionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionContractCallerRaw struct {
	Contract *AuctionContractCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionContractTransactorRaw struct {
	Contract *AuctionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionContract creates a new instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContract(address common.Address, backend bind.ContractBackend) (*AuctionContract, error) {
	contract, err := bindAuctionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionContract{AuctionContractCaller: AuctionContractCaller{contract: contract}, AuctionContractTransactor: AuctionContractTransactor{contract: contract}, AuctionContractFilterer: AuctionContractFilterer{contract: contract}}, nil
}

// NewAuctionContractCaller creates a new read-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractCaller(address common.Address, caller bind.ContractCaller) (*AuctionContractCaller, error) {
	contract, err := bindAuctionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractCaller{contract: contract}, nil
}

// NewAuctionContractTransactor creates a new write-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionContractTransactor, error) {
	contract, err := bindAuctionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractTransactor{contract: contract}, nil
}

// NewAuctionContractFilterer creates a new log filterer instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionContractFilterer, error) {
	contract, err := bindAuctionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionContractFilterer{contract: contract}, nil
}

// bindAuctionContract binds a generic wrapper to an already deployed contract.
func bindAuctionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.AuctionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transact(opts, method, params...)
}

// CalculateTotalPrice is a free data retrieval call binding the contract method 0x0c8788bb.
//
// Solidity: function calculateTotalPrice(address _user) view returns(uint256)
func (_AuctionContract *AuctionContractCaller) CalculateTotalPrice(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "calculateTotalPrice", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTotalPrice is a free data retrieval call binding the contract method 0x0c8788bb.
//
// Solidity: function calculateTotalPrice(address _user) view returns(uint256)
func (_AuctionContract *AuctionContractSession) CalculateTotalPrice(_user common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.CalculateTotalPrice(&_AuctionContract.CallOpts, _user)
}

// CalculateTotalPrice is a free data retrieval call binding the contract method 0x0c8788bb.
//
// Solidity: function calculateTotalPrice(address _user) view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) CalculateTotalPrice(_user common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.CalculateTotalPrice(&_AuctionContract.CallOpts, _user)
}

// CurrentBuyer is a free data retrieval call binding the contract method 0x5bf3c518.
//
// Solidity: function currentBuyer() view returns(address)
func (_AuctionContract *AuctionContractCaller) CurrentBuyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "currentBuyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentBuyer is a free data retrieval call binding the contract method 0x5bf3c518.
//
// Solidity: function currentBuyer() view returns(address)
func (_AuctionContract *AuctionContractSession) CurrentBuyer() (common.Address, error) {
	return _AuctionContract.Contract.CurrentBuyer(&_AuctionContract.CallOpts)
}

// CurrentBuyer is a free data retrieval call binding the contract method 0x5bf3c518.
//
// Solidity: function currentBuyer() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) CurrentBuyer() (common.Address, error) {
	return _AuctionContract.Contract.CurrentBuyer(&_AuctionContract.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) CurrentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_AuctionContract *AuctionContractSession) CurrentPrice() (*big.Int, error) {
	return _AuctionContract.Contract.CurrentPrice(&_AuctionContract.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) CurrentPrice() (*big.Int, error) {
	return _AuctionContract.Contract.CurrentPrice(&_AuctionContract.CallOpts)
}

// DurationMinutes is a free data retrieval call binding the contract method 0x25f1f3dd.
//
// Solidity: function durationMinutes() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) DurationMinutes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "durationMinutes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DurationMinutes is a free data retrieval call binding the contract method 0x25f1f3dd.
//
// Solidity: function durationMinutes() view returns(uint256)
func (_AuctionContract *AuctionContractSession) DurationMinutes() (*big.Int, error) {
	return _AuctionContract.Contract.DurationMinutes(&_AuctionContract.CallOpts)
}

// DurationMinutes is a free data retrieval call binding the contract method 0x25f1f3dd.
//
// Solidity: function durationMinutes() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) DurationMinutes() (*big.Int, error) {
	return _AuctionContract.Contract.DurationMinutes(&_AuctionContract.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_AuctionContract *AuctionContractSession) EndTime() (*big.Int, error) {
	return _AuctionContract.Contract.EndTime(&_AuctionContract.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) EndTime() (*big.Int, error) {
	return _AuctionContract.Contract.EndTime(&_AuctionContract.CallOpts)
}

// EthMap is a free data retrieval call binding the contract method 0xe6ebc5ee.
//
// Solidity: function ethMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractCaller) EthMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "ethMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthMap is a free data retrieval call binding the contract method 0xe6ebc5ee.
//
// Solidity: function ethMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractSession) EthMap(arg0 common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.EthMap(&_AuctionContract.CallOpts, arg0)
}

// EthMap is a free data retrieval call binding the contract method 0xe6ebc5ee.
//
// Solidity: function ethMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) EthMap(arg0 common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.EthMap(&_AuctionContract.CallOpts, arg0)
}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractCaller) EthPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "ethPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractSession) EthPriceFeed() (common.Address, error) {
	return _AuctionContract.Contract.EthPriceFeed(&_AuctionContract.CallOpts)
}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) EthPriceFeed() (common.Address, error) {
	return _AuctionContract.Contract.EthPriceFeed(&_AuctionContract.CallOpts)
}

// GetLatestETHPrice is a free data retrieval call binding the contract method 0x777e0d86.
//
// Solidity: function getLatestETHPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) GetLatestETHPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "getLatestETHPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestETHPrice is a free data retrieval call binding the contract method 0x777e0d86.
//
// Solidity: function getLatestETHPrice() view returns(uint256)
func (_AuctionContract *AuctionContractSession) GetLatestETHPrice() (*big.Int, error) {
	return _AuctionContract.Contract.GetLatestETHPrice(&_AuctionContract.CallOpts)
}

// GetLatestETHPrice is a free data retrieval call binding the contract method 0x777e0d86.
//
// Solidity: function getLatestETHPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) GetLatestETHPrice() (*big.Int, error) {
	return _AuctionContract.Contract.GetLatestETHPrice(&_AuctionContract.CallOpts)
}

// GetLatestUSDCPrice is a free data retrieval call binding the contract method 0x7f281a64.
//
// Solidity: function getLatestUSDCPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) GetLatestUSDCPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "getLatestUSDCPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestUSDCPrice is a free data retrieval call binding the contract method 0x7f281a64.
//
// Solidity: function getLatestUSDCPrice() view returns(uint256)
func (_AuctionContract *AuctionContractSession) GetLatestUSDCPrice() (*big.Int, error) {
	return _AuctionContract.Contract.GetLatestUSDCPrice(&_AuctionContract.CallOpts)
}

// GetLatestUSDCPrice is a free data retrieval call binding the contract method 0x7f281a64.
//
// Solidity: function getLatestUSDCPrice() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) GetLatestUSDCPrice() (*big.Int, error) {
	return _AuctionContract.Contract.GetLatestUSDCPrice(&_AuctionContract.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_AuctionContract *AuctionContractCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_AuctionContract *AuctionContractSession) Nft() (common.Address, error) {
	return _AuctionContract.Contract.Nft(&_AuctionContract.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) Nft() (common.Address, error) {
	return _AuctionContract.Contract.Nft(&_AuctionContract.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) pure returns(bytes4)
func (_AuctionContract *AuctionContractCaller) OnERC721Received(opts *bind.CallOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) ([4]byte, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "onERC721Received", _operator, _from, _tokenId, _data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) pure returns(bytes4)
func (_AuctionContract *AuctionContractSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) ([4]byte, error) {
	return _AuctionContract.Contract.OnERC721Received(&_AuctionContract.CallOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) pure returns(bytes4)
func (_AuctionContract *AuctionContractCallerSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) ([4]byte, error) {
	return _AuctionContract.Contract.OnERC721Received(&_AuctionContract.CallOpts, _operator, _from, _tokenId, _data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionContract *AuctionContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionContract *AuctionContractSession) Owner() (common.Address, error) {
	return _AuctionContract.Contract.Owner(&_AuctionContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) Owner() (common.Address, error) {
	return _AuctionContract.Contract.Owner(&_AuctionContract.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_AuctionContract *AuctionContractCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_AuctionContract *AuctionContractSession) Seller() (common.Address, error) {
	return _AuctionContract.Contract.Seller(&_AuctionContract.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) Seller() (common.Address, error) {
	return _AuctionContract.Contract.Seller(&_AuctionContract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_AuctionContract *AuctionContractCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_AuctionContract *AuctionContractSession) State() (uint8, error) {
	return _AuctionContract.Contract.State(&_AuctionContract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_AuctionContract *AuctionContractCallerSession) State() (uint8, error) {
	return _AuctionContract.Contract.State(&_AuctionContract.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_AuctionContract *AuctionContractCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_AuctionContract *AuctionContractSession) TokenId() (*big.Int, error) {
	return _AuctionContract.Contract.TokenId(&_AuctionContract.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) TokenId() (*big.Int, error) {
	return _AuctionContract.Contract.TokenId(&_AuctionContract.CallOpts)
}

// UsdcMap is a free data retrieval call binding the contract method 0xb3ef84a2.
//
// Solidity: function usdcMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractCaller) UsdcMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "usdcMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcMap is a free data retrieval call binding the contract method 0xb3ef84a2.
//
// Solidity: function usdcMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractSession) UsdcMap(arg0 common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.UsdcMap(&_AuctionContract.CallOpts, arg0)
}

// UsdcMap is a free data retrieval call binding the contract method 0xb3ef84a2.
//
// Solidity: function usdcMap(address ) view returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) UsdcMap(arg0 common.Address) (*big.Int, error) {
	return _AuctionContract.Contract.UsdcMap(&_AuctionContract.CallOpts, arg0)
}

// UsdcPriceFeed is a free data retrieval call binding the contract method 0x58f40399.
//
// Solidity: function usdcPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractCaller) UsdcPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "usdcPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcPriceFeed is a free data retrieval call binding the contract method 0x58f40399.
//
// Solidity: function usdcPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractSession) UsdcPriceFeed() (common.Address, error) {
	return _AuctionContract.Contract.UsdcPriceFeed(&_AuctionContract.CallOpts)
}

// UsdcPriceFeed is a free data retrieval call binding the contract method 0x58f40399.
//
// Solidity: function usdcPriceFeed() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) UsdcPriceFeed() (common.Address, error) {
	return _AuctionContract.Contract.UsdcPriceFeed(&_AuctionContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_AuctionContract *AuctionContractCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionContract.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_AuctionContract *AuctionContractSession) UsdcToken() (common.Address, error) {
	return _AuctionContract.Contract.UsdcToken(&_AuctionContract.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_AuctionContract *AuctionContractCallerSession) UsdcToken() (common.Address, error) {
	return _AuctionContract.Contract.UsdcToken(&_AuctionContract.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 usdcAmount) payable returns()
func (_AuctionContract *AuctionContractTransactor) Bid(opts *bind.TransactOpts, usdcAmount *big.Int) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "bid", usdcAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 usdcAmount) payable returns()
func (_AuctionContract *AuctionContractSession) Bid(usdcAmount *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Bid(&_AuctionContract.TransactOpts, usdcAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 usdcAmount) payable returns()
func (_AuctionContract *AuctionContractTransactorSession) Bid(usdcAmount *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Bid(&_AuctionContract.TransactOpts, usdcAmount)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address _nft, uint256 _tokenId, uint256 _durationMinutes) returns()
func (_AuctionContract *AuctionContractTransactor) CreateAuction(opts *bind.TransactOpts, _nft common.Address, _tokenId *big.Int, _durationMinutes *big.Int) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "createAuction", _nft, _tokenId, _durationMinutes)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address _nft, uint256 _tokenId, uint256 _durationMinutes) returns()
func (_AuctionContract *AuctionContractSession) CreateAuction(_nft common.Address, _tokenId *big.Int, _durationMinutes *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.CreateAuction(&_AuctionContract.TransactOpts, _nft, _tokenId, _durationMinutes)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address _nft, uint256 _tokenId, uint256 _durationMinutes) returns()
func (_AuctionContract *AuctionContractTransactorSession) CreateAuction(_nft common.Address, _tokenId *big.Int, _durationMinutes *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.CreateAuction(&_AuctionContract.TransactOpts, _nft, _tokenId, _durationMinutes)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AuctionContract *AuctionContractTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AuctionContract *AuctionContractSession) Finalize() (*types.Transaction, error) {
	return _AuctionContract.Contract.Finalize(&_AuctionContract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_AuctionContract *AuctionContractTransactorSession) Finalize() (*types.Transaction, error) {
	return _AuctionContract.Contract.Finalize(&_AuctionContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionContract *AuctionContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionContract *AuctionContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionContract.Contract.RenounceOwnership(&_AuctionContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionContract *AuctionContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionContract.Contract.RenounceOwnership(&_AuctionContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionContract *AuctionContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionContract *AuctionContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionContract.Contract.TransferOwnership(&_AuctionContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionContract *AuctionContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionContract.Contract.TransferOwnership(&_AuctionContract.TransactOpts, newOwner)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_AuctionContract *AuctionContractTransactor) WithDraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "withDraw")
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_AuctionContract *AuctionContractSession) WithDraw() (*types.Transaction, error) {
	return _AuctionContract.Contract.WithDraw(&_AuctionContract.TransactOpts)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_AuctionContract *AuctionContractTransactorSession) WithDraw() (*types.Transaction, error) {
	return _AuctionContract.Contract.WithDraw(&_AuctionContract.TransactOpts)
}

// AuctionContractAuctionEndIterator is returned from FilterAuctionEnd and is used to iterate over the raw logs and unpacked data for AuctionEnd events raised by the AuctionContract contract.
type AuctionContractAuctionEndIterator struct {
	Event *AuctionContractAuctionEnd // Event containing the contract specifics and raw log

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
func (it *AuctionContractAuctionEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractAuctionEnd)
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
		it.Event = new(AuctionContractAuctionEnd)
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
func (it *AuctionContractAuctionEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractAuctionEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractAuctionEnd represents a AuctionEnd event raised by the AuctionContract contract.
type AuctionContractAuctionEnd struct {
	Buyer common.Address
	Bid   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnd is a free log retrieval operation binding the contract event 0x6ae96154c7c63a88543f3fdfca38bc8f87f4e2b034479c9429858b4b56fa0bec.
//
// Solidity: event AuctionEnd(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) FilterAuctionEnd(opts *bind.FilterOpts, buyer []common.Address) (*AuctionContractAuctionEndIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "AuctionEnd", buyerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractAuctionEndIterator{contract: _AuctionContract.contract, event: "AuctionEnd", logs: logs, sub: sub}, nil
}

// WatchAuctionEnd is a free log subscription operation binding the contract event 0x6ae96154c7c63a88543f3fdfca38bc8f87f4e2b034479c9429858b4b56fa0bec.
//
// Solidity: event AuctionEnd(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) WatchAuctionEnd(opts *bind.WatchOpts, sink chan<- *AuctionContractAuctionEnd, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "AuctionEnd", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractAuctionEnd)
				if err := _AuctionContract.contract.UnpackLog(event, "AuctionEnd", log); err != nil {
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

// ParseAuctionEnd is a log parse operation binding the contract event 0x6ae96154c7c63a88543f3fdfca38bc8f87f4e2b034479c9429858b4b56fa0bec.
//
// Solidity: event AuctionEnd(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) ParseAuctionEnd(log types.Log) (*AuctionContractAuctionEnd, error) {
	event := new(AuctionContractAuctionEnd)
	if err := _AuctionContract.contract.UnpackLog(event, "AuctionEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionContractBidUSDIterator is returned from FilterBidUSD and is used to iterate over the raw logs and unpacked data for BidUSD events raised by the AuctionContract contract.
type AuctionContractBidUSDIterator struct {
	Event *AuctionContractBidUSD // Event containing the contract specifics and raw log

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
func (it *AuctionContractBidUSDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractBidUSD)
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
		it.Event = new(AuctionContractBidUSD)
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
func (it *AuctionContractBidUSDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractBidUSDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractBidUSD represents a BidUSD event raised by the AuctionContract contract.
type AuctionContractBidUSD struct {
	Buyer common.Address
	Bid   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBidUSD is a free log retrieval operation binding the contract event 0x0cc687afcc9ee18a07769e118ca5bca8d94af38992e1b8de14ca8d6c205c1277.
//
// Solidity: event BidUSD(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) FilterBidUSD(opts *bind.FilterOpts, buyer []common.Address) (*AuctionContractBidUSDIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "BidUSD", buyerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractBidUSDIterator{contract: _AuctionContract.contract, event: "BidUSD", logs: logs, sub: sub}, nil
}

// WatchBidUSD is a free log subscription operation binding the contract event 0x0cc687afcc9ee18a07769e118ca5bca8d94af38992e1b8de14ca8d6c205c1277.
//
// Solidity: event BidUSD(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) WatchBidUSD(opts *bind.WatchOpts, sink chan<- *AuctionContractBidUSD, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "BidUSD", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractBidUSD)
				if err := _AuctionContract.contract.UnpackLog(event, "BidUSD", log); err != nil {
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

// ParseBidUSD is a log parse operation binding the contract event 0x0cc687afcc9ee18a07769e118ca5bca8d94af38992e1b8de14ca8d6c205c1277.
//
// Solidity: event BidUSD(address indexed buyer, uint256 bid)
func (_AuctionContract *AuctionContractFilterer) ParseBidUSD(log types.Log) (*AuctionContractBidUSD, error) {
	event := new(AuctionContractBidUSD)
	if err := _AuctionContract.contract.UnpackLog(event, "BidUSD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionContractCreateAuctionIterator is returned from FilterCreateAuction and is used to iterate over the raw logs and unpacked data for CreateAuction events raised by the AuctionContract contract.
type AuctionContractCreateAuctionIterator struct {
	Event *AuctionContractCreateAuction // Event containing the contract specifics and raw log

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
func (it *AuctionContractCreateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractCreateAuction)
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
		it.Event = new(AuctionContractCreateAuction)
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
func (it *AuctionContractCreateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractCreateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractCreateAuction represents a CreateAuction event raised by the AuctionContract contract.
type AuctionContractCreateAuction struct {
	Seller          common.Address
	Nft             common.Address
	TokenId         *big.Int
	DurationMinutes *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateAuction is a free log retrieval operation binding the contract event 0xfb02b3662d6236b2e42a0fcea31ad824bcd600646b7da6cd6c61c24cebdbda82.
//
// Solidity: event CreateAuction(address indexed seller, address indexed nft, uint256 tokenId, uint256 durationMinutes)
func (_AuctionContract *AuctionContractFilterer) FilterCreateAuction(opts *bind.FilterOpts, seller []common.Address, nft []common.Address) (*AuctionContractCreateAuctionIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "CreateAuction", sellerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractCreateAuctionIterator{contract: _AuctionContract.contract, event: "CreateAuction", logs: logs, sub: sub}, nil
}

// WatchCreateAuction is a free log subscription operation binding the contract event 0xfb02b3662d6236b2e42a0fcea31ad824bcd600646b7da6cd6c61c24cebdbda82.
//
// Solidity: event CreateAuction(address indexed seller, address indexed nft, uint256 tokenId, uint256 durationMinutes)
func (_AuctionContract *AuctionContractFilterer) WatchCreateAuction(opts *bind.WatchOpts, sink chan<- *AuctionContractCreateAuction, seller []common.Address, nft []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "CreateAuction", sellerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractCreateAuction)
				if err := _AuctionContract.contract.UnpackLog(event, "CreateAuction", log); err != nil {
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

// ParseCreateAuction is a log parse operation binding the contract event 0xfb02b3662d6236b2e42a0fcea31ad824bcd600646b7da6cd6c61c24cebdbda82.
//
// Solidity: event CreateAuction(address indexed seller, address indexed nft, uint256 tokenId, uint256 durationMinutes)
func (_AuctionContract *AuctionContractFilterer) ParseCreateAuction(log types.Log) (*AuctionContractCreateAuction, error) {
	event := new(AuctionContractCreateAuction)
	if err := _AuctionContract.contract.UnpackLog(event, "CreateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionContract contract.
type AuctionContractOwnershipTransferredIterator struct {
	Event *AuctionContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractOwnershipTransferred)
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
		it.Event = new(AuctionContractOwnershipTransferred)
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
func (it *AuctionContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionContract contract.
type AuctionContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionContract *AuctionContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractOwnershipTransferredIterator{contract: _AuctionContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionContract *AuctionContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractOwnershipTransferred)
				if err := _AuctionContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionContract *AuctionContractFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionContractOwnershipTransferred, error) {
	event := new(AuctionContractOwnershipTransferred)
	if err := _AuctionContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
