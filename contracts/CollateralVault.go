// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CollateralVaultMetaData contains all meta data concerning the CollateralVault contract.
var CollateralVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWAL_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAuthorizedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelWithdrawalRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserActiveRequests\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalUnlockTime\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestWithdrawal\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOrderEngine\",\"inputs\":[{\"name\":\"_orderEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSettlementEngine\",\"inputs\":[{\"name\":\"_settlementEngine\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settlementEngine\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"usedSignatures\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userRequestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawCollateralSigned\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CollateralBelowZero\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralWithdrawn\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenAuthorized\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCancelled\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequested\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpiredSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWhitelistSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSecurityCouncil\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UsedSignature\",\"inputs\":[]}]",
}

// CollateralVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralVaultMetaData.ABI instead.
var CollateralVaultABI = CollateralVaultMetaData.ABI

// CollateralVault is an auto generated Go binding around an Ethereum contract.
type CollateralVault struct {
	CollateralVaultCaller     // Read-only binding to the contract
	CollateralVaultTransactor // Write-only binding to the contract
	CollateralVaultFilterer   // Log filterer for contract events
}

// CollateralVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralVaultSession struct {
	Contract     *CollateralVault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralVaultCallerSession struct {
	Contract *CollateralVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CollateralVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralVaultTransactorSession struct {
	Contract     *CollateralVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CollateralVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralVaultRaw struct {
	Contract *CollateralVault // Generic contract binding to access the raw methods on
}

// CollateralVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralVaultCallerRaw struct {
	Contract *CollateralVaultCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralVaultTransactorRaw struct {
	Contract *CollateralVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateralVault creates a new instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVault(address common.Address, backend bind.ContractBackend) (*CollateralVault, error) {
	contract, err := bindCollateralVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollateralVault{CollateralVaultCaller: CollateralVaultCaller{contract: contract}, CollateralVaultTransactor: CollateralVaultTransactor{contract: contract}, CollateralVaultFilterer: CollateralVaultFilterer{contract: contract}}, nil
}

// NewCollateralVaultCaller creates a new read-only instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultCaller(address common.Address, caller bind.ContractCaller) (*CollateralVaultCaller, error) {
	contract, err := bindCollateralVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCaller{contract: contract}, nil
}

// NewCollateralVaultTransactor creates a new write-only instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralVaultTransactor, error) {
	contract, err := bindCollateralVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultTransactor{contract: contract}, nil
}

// NewCollateralVaultFilterer creates a new log filterer instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralVaultFilterer, error) {
	contract, err := bindCollateralVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultFilterer{contract: contract}, nil
}

// bindCollateralVault binds a generic wrapper to an already deployed contract.
func bindCollateralVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollateralVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralVault *CollateralVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralVault.Contract.CollateralVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralVault *CollateralVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.Contract.CollateralVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralVault *CollateralVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralVault.Contract.CollateralVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralVault *CollateralVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralVault *CollateralVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralVault *CollateralVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralVault.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CollateralVault *CollateralVaultCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CollateralVault *CollateralVaultSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CollateralVault.Contract.UPGRADEINTERFACEVERSION(&_CollateralVault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CollateralVault *CollateralVaultCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CollateralVault.Contract.UPGRADEINTERFACEVERSION(&_CollateralVault.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) WITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_CollateralVault *CollateralVaultSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _CollateralVault.Contract.WITHDRAWALDELAY(&_CollateralVault.CallOpts)
}

// WITHDRAWALDELAY is a free data retrieval call binding the contract method 0x0ebb172a.
//
// Solidity: function WITHDRAWAL_DELAY() view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) WITHDRAWALDELAY() (*big.Int, error) {
	return _CollateralVault.Contract.WITHDRAWALDELAY(&_CollateralVault.CallOpts)
}

// AuthorizedTokens is a free data retrieval call binding the contract method 0x9bc7c586.
//
// Solidity: function authorizedTokens(address ) view returns(bool)
func (_CollateralVault *CollateralVaultCaller) AuthorizedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "authorizedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedTokens is a free data retrieval call binding the contract method 0x9bc7c586.
//
// Solidity: function authorizedTokens(address ) view returns(bool)
func (_CollateralVault *CollateralVaultSession) AuthorizedTokens(arg0 common.Address) (bool, error) {
	return _CollateralVault.Contract.AuthorizedTokens(&_CollateralVault.CallOpts, arg0)
}

// AuthorizedTokens is a free data retrieval call binding the contract method 0x9bc7c586.
//
// Solidity: function authorizedTokens(address ) view returns(bool)
func (_CollateralVault *CollateralVaultCallerSession) AuthorizedTokens(arg0 common.Address) (bool, error) {
	return _CollateralVault.Contract.AuthorizedTokens(&_CollateralVault.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(int256)
func (_CollateralVault *CollateralVaultCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(int256)
func (_CollateralVault *CollateralVaultSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CollateralVault.Contract.Balances(&_CollateralVault.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(int256)
func (_CollateralVault *CollateralVaultCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CollateralVault.Contract.Balances(&_CollateralVault.CallOpts, arg0, arg1)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_CollateralVault *CollateralVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_CollateralVault *CollateralVaultSession) FeeRecipient() (common.Address, error) {
	return _CollateralVault.Contract.FeeRecipient(&_CollateralVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _CollateralVault.Contract.FeeRecipient(&_CollateralVault.CallOpts)
}

// GetUserActiveRequests is a free data retrieval call binding the contract method 0x815bed7f.
//
// Solidity: function getUserActiveRequests(address user, address token) view returns(uint256[])
func (_CollateralVault *CollateralVaultCaller) GetUserActiveRequests(opts *bind.CallOpts, user common.Address, token common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "getUserActiveRequests", user, token)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserActiveRequests is a free data retrieval call binding the contract method 0x815bed7f.
//
// Solidity: function getUserActiveRequests(address user, address token) view returns(uint256[])
func (_CollateralVault *CollateralVaultSession) GetUserActiveRequests(user common.Address, token common.Address) ([]*big.Int, error) {
	return _CollateralVault.Contract.GetUserActiveRequests(&_CollateralVault.CallOpts, user, token)
}

// GetUserActiveRequests is a free data retrieval call binding the contract method 0x815bed7f.
//
// Solidity: function getUserActiveRequests(address user, address token) view returns(uint256[])
func (_CollateralVault *CollateralVaultCallerSession) GetUserActiveRequests(user common.Address, token common.Address) ([]*big.Int, error) {
	return _CollateralVault.Contract.GetUserActiveRequests(&_CollateralVault.CallOpts, user, token)
}

// GetWithdrawalUnlockTime is a free data retrieval call binding the contract method 0x7254f681.
//
// Solidity: function getWithdrawalUnlockTime(uint256 requestId) view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) GetWithdrawalUnlockTime(opts *bind.CallOpts, requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "getWithdrawalUnlockTime", requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalUnlockTime is a free data retrieval call binding the contract method 0x7254f681.
//
// Solidity: function getWithdrawalUnlockTime(uint256 requestId) view returns(uint256)
func (_CollateralVault *CollateralVaultSession) GetWithdrawalUnlockTime(requestId *big.Int) (*big.Int, error) {
	return _CollateralVault.Contract.GetWithdrawalUnlockTime(&_CollateralVault.CallOpts, requestId)
}

// GetWithdrawalUnlockTime is a free data retrieval call binding the contract method 0x7254f681.
//
// Solidity: function getWithdrawalUnlockTime(uint256 requestId) view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) GetWithdrawalUnlockTime(requestId *big.Int) (*big.Int, error) {
	return _CollateralVault.Contract.GetWithdrawalUnlockTime(&_CollateralVault.CallOpts, requestId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_CollateralVault *CollateralVaultCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_CollateralVault *CollateralVaultSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _CollateralVault.Contract.IsTrustedForwarder(&_CollateralVault.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_CollateralVault *CollateralVaultCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _CollateralVault.Contract.IsTrustedForwarder(&_CollateralVault.CallOpts, forwarder)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_CollateralVault *CollateralVaultSession) NextRequestId() (*big.Int, error) {
	return _CollateralVault.Contract.NextRequestId(&_CollateralVault.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) NextRequestId() (*big.Int, error) {
	return _CollateralVault.Contract.NextRequestId(&_CollateralVault.CallOpts)
}

// OrderEngine is a free data retrieval call binding the contract method 0x5132519b.
//
// Solidity: function orderEngine() view returns(address)
func (_CollateralVault *CollateralVaultCaller) OrderEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "orderEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderEngine is a free data retrieval call binding the contract method 0x5132519b.
//
// Solidity: function orderEngine() view returns(address)
func (_CollateralVault *CollateralVaultSession) OrderEngine() (common.Address, error) {
	return _CollateralVault.Contract.OrderEngine(&_CollateralVault.CallOpts)
}

// OrderEngine is a free data retrieval call binding the contract method 0x5132519b.
//
// Solidity: function orderEngine() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) OrderEngine() (common.Address, error) {
	return _CollateralVault.Contract.OrderEngine(&_CollateralVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultSession) Owner() (common.Address, error) {
	return _CollateralVault.Contract.Owner(&_CollateralVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) Owner() (common.Address, error) {
	return _CollateralVault.Contract.Owner(&_CollateralVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CollateralVault *CollateralVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CollateralVault *CollateralVaultSession) Paused() (bool, error) {
	return _CollateralVault.Contract.Paused(&_CollateralVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CollateralVault *CollateralVaultCallerSession) Paused() (bool, error) {
	return _CollateralVault.Contract.Paused(&_CollateralVault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralVault *CollateralVaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralVault *CollateralVaultSession) ProxiableUUID() ([32]byte, error) {
	return _CollateralVault.Contract.ProxiableUUID(&_CollateralVault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CollateralVault *CollateralVaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CollateralVault.Contract.ProxiableUUID(&_CollateralVault.CallOpts)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_CollateralVault *CollateralVaultCaller) SettlementEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "settlementEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_CollateralVault *CollateralVaultSession) SettlementEngine() (common.Address, error) {
	return _CollateralVault.Contract.SettlementEngine(&_CollateralVault.CallOpts)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) SettlementEngine() (common.Address, error) {
	return _CollateralVault.Contract.SettlementEngine(&_CollateralVault.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_CollateralVault *CollateralVaultCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_CollateralVault *CollateralVaultSession) TrustedForwarder() (common.Address, error) {
	return _CollateralVault.Contract.TrustedForwarder(&_CollateralVault.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) TrustedForwarder() (common.Address, error) {
	return _CollateralVault.Contract.TrustedForwarder(&_CollateralVault.CallOpts)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_CollateralVault *CollateralVaultCaller) UsedSignatures(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "usedSignatures", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_CollateralVault *CollateralVaultSession) UsedSignatures(arg0 []byte) (bool, error) {
	return _CollateralVault.Contract.UsedSignatures(&_CollateralVault.CallOpts, arg0)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xe949580e.
//
// Solidity: function usedSignatures(bytes ) view returns(bool)
func (_CollateralVault *CollateralVaultCallerSession) UsedSignatures(arg0 []byte) (bool, error) {
	return _CollateralVault.Contract.UsedSignatures(&_CollateralVault.CallOpts, arg0)
}

// UserRequestIds is a free data retrieval call binding the contract method 0xdd4f9a3b.
//
// Solidity: function userRequestIds(address , address , uint256 ) view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) UserRequestIds(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "userRequestIds", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRequestIds is a free data retrieval call binding the contract method 0xdd4f9a3b.
//
// Solidity: function userRequestIds(address , address , uint256 ) view returns(uint256)
func (_CollateralVault *CollateralVaultSession) UserRequestIds(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _CollateralVault.Contract.UserRequestIds(&_CollateralVault.CallOpts, arg0, arg1, arg2)
}

// UserRequestIds is a free data retrieval call binding the contract method 0xdd4f9a3b.
//
// Solidity: function userRequestIds(address , address , uint256 ) view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) UserRequestIds(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _CollateralVault.Contract.UserRequestIds(&_CollateralVault.CallOpts, arg0, arg1, arg2)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_CollateralVault *CollateralVaultCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_CollateralVault *CollateralVaultSession) Whitelist(arg0 common.Address) (bool, error) {
	return _CollateralVault.Contract.Whitelist(&_CollateralVault.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_CollateralVault *CollateralVaultCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _CollateralVault.Contract.Whitelist(&_CollateralVault.CallOpts, arg0)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x937b2581.
//
// Solidity: function withdrawalRequests(uint256 ) view returns(address user, address token, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultCaller) WithdrawalRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "withdrawalRequests", arg0)

	outstruct := new(struct {
		User       common.Address
		Token      common.Address
		Amount     *big.Int
		UnlockTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x937b2581.
//
// Solidity: function withdrawalRequests(uint256 ) view returns(address user, address token, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultSession) WithdrawalRequests(arg0 *big.Int) (struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _CollateralVault.Contract.WithdrawalRequests(&_CollateralVault.CallOpts, arg0)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x937b2581.
//
// Solidity: function withdrawalRequests(uint256 ) view returns(address user, address token, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultCallerSession) WithdrawalRequests(arg0 *big.Int) (struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _CollateralVault.Contract.WithdrawalRequests(&_CollateralVault.CallOpts, arg0)
}

// AddAuthorizedToken is a paid mutator transaction binding the contract method 0xd9d23d64.
//
// Solidity: function addAuthorizedToken(address token, bool authorized) returns()
func (_CollateralVault *CollateralVaultTransactor) AddAuthorizedToken(opts *bind.TransactOpts, token common.Address, authorized bool) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "addAuthorizedToken", token, authorized)
}

// AddAuthorizedToken is a paid mutator transaction binding the contract method 0xd9d23d64.
//
// Solidity: function addAuthorizedToken(address token, bool authorized) returns()
func (_CollateralVault *CollateralVaultSession) AddAuthorizedToken(token common.Address, authorized bool) (*types.Transaction, error) {
	return _CollateralVault.Contract.AddAuthorizedToken(&_CollateralVault.TransactOpts, token, authorized)
}

// AddAuthorizedToken is a paid mutator transaction binding the contract method 0xd9d23d64.
//
// Solidity: function addAuthorizedToken(address token, bool authorized) returns()
func (_CollateralVault *CollateralVaultTransactorSession) AddAuthorizedToken(token common.Address, authorized bool) (*types.Transaction, error) {
	return _CollateralVault.Contract.AddAuthorizedToken(&_CollateralVault.TransactOpts, token, authorized)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x45b05d09.
//
// Solidity: function cancelWithdrawalRequest(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultTransactor) CancelWithdrawalRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "cancelWithdrawalRequest", requestId)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x45b05d09.
//
// Solidity: function cancelWithdrawalRequest(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultSession) CancelWithdrawalRequest(requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.CancelWithdrawalRequest(&_CollateralVault.TransactOpts, requestId)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0x45b05d09.
//
// Solidity: function cancelWithdrawalRequest(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultTransactorSession) CancelWithdrawalRequest(requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.CancelWithdrawalRequest(&_CollateralVault.TransactOpts, requestId)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x2a3a0141.
//
// Solidity: function depositCollateral(address token, uint256 amount, bytes signature) returns()
func (_CollateralVault *CollateralVaultTransactor) DepositCollateral(opts *bind.TransactOpts, token common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "depositCollateral", token, amount, signature)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x2a3a0141.
//
// Solidity: function depositCollateral(address token, uint256 amount, bytes signature) returns()
func (_CollateralVault *CollateralVaultSession) DepositCollateral(token common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral(&_CollateralVault.TransactOpts, token, amount, signature)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0x2a3a0141.
//
// Solidity: function depositCollateral(address token, uint256 amount, bytes signature) returns()
func (_CollateralVault *CollateralVaultTransactorSession) DepositCollateral(token common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral(&_CollateralVault.TransactOpts, token, amount, signature)
}

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactor) DepositCollateral0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "depositCollateral0", token, amount)
}

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_CollateralVault *CollateralVaultSession) DepositCollateral0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral0(&_CollateralVault.TransactOpts, token, amount)
}

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address token, uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactorSession) DepositCollateral0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral0(&_CollateralVault.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _feeRecipient, address _forwarder, address _governance) returns()
func (_CollateralVault *CollateralVaultTransactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "initialize", _feeRecipient, _forwarder, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _feeRecipient, address _forwarder, address _governance) returns()
func (_CollateralVault *CollateralVaultSession) Initialize(_feeRecipient common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.Initialize(&_CollateralVault.TransactOpts, _feeRecipient, _forwarder, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _feeRecipient, address _forwarder, address _governance) returns()
func (_CollateralVault *CollateralVaultTransactorSession) Initialize(_feeRecipient common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.Initialize(&_CollateralVault.TransactOpts, _feeRecipient, _forwarder, _governance)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CollateralVault *CollateralVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CollateralVault *CollateralVaultSession) Pause() (*types.Transaction, error) {
	return _CollateralVault.Contract.Pause(&_CollateralVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CollateralVault *CollateralVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _CollateralVault.Contract.Pause(&_CollateralVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralVault.Contract.RenounceOwnership(&_CollateralVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralVault.Contract.RenounceOwnership(&_CollateralVault.TransactOpts)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address token, uint256 amount) returns(uint256 requestId)
func (_CollateralVault *CollateralVaultTransactor) RequestWithdrawal(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "requestWithdrawal", token, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address token, uint256 amount) returns(uint256 requestId)
func (_CollateralVault *CollateralVaultSession) RequestWithdrawal(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.RequestWithdrawal(&_CollateralVault.TransactOpts, token, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xda95ebf7.
//
// Solidity: function requestWithdrawal(address token, uint256 amount) returns(uint256 requestId)
func (_CollateralVault *CollateralVaultTransactorSession) RequestWithdrawal(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.RequestWithdrawal(&_CollateralVault.TransactOpts, token, amount)
}

// SetOrderEngine is a paid mutator transaction binding the contract method 0x5acc17b3.
//
// Solidity: function setOrderEngine(address _orderEngine) returns()
func (_CollateralVault *CollateralVaultTransactor) SetOrderEngine(opts *bind.TransactOpts, _orderEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "setOrderEngine", _orderEngine)
}

// SetOrderEngine is a paid mutator transaction binding the contract method 0x5acc17b3.
//
// Solidity: function setOrderEngine(address _orderEngine) returns()
func (_CollateralVault *CollateralVaultSession) SetOrderEngine(_orderEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.SetOrderEngine(&_CollateralVault.TransactOpts, _orderEngine)
}

// SetOrderEngine is a paid mutator transaction binding the contract method 0x5acc17b3.
//
// Solidity: function setOrderEngine(address _orderEngine) returns()
func (_CollateralVault *CollateralVaultTransactorSession) SetOrderEngine(_orderEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.SetOrderEngine(&_CollateralVault.TransactOpts, _orderEngine)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _settlementEngine) returns()
func (_CollateralVault *CollateralVaultTransactor) SetSettlementEngine(opts *bind.TransactOpts, _settlementEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "setSettlementEngine", _settlementEngine)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _settlementEngine) returns()
func (_CollateralVault *CollateralVaultSession) SetSettlementEngine(_settlementEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.SetSettlementEngine(&_CollateralVault.TransactOpts, _settlementEngine)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _settlementEngine) returns()
func (_CollateralVault *CollateralVaultTransactorSession) SetSettlementEngine(_settlementEngine common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.SetSettlementEngine(&_CollateralVault.TransactOpts, _settlementEngine)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.TransferOwnership(&_CollateralVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.TransferOwnership(&_CollateralVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CollateralVault *CollateralVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CollateralVault *CollateralVaultSession) Unpause() (*types.Transaction, error) {
	return _CollateralVault.Contract.Unpause(&_CollateralVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CollateralVault *CollateralVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _CollateralVault.Contract.Unpause(&_CollateralVault.TransactOpts)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x33f2f2e7.
//
// Solidity: function updateBalance(address token, address trader, int256 amount) returns()
func (_CollateralVault *CollateralVaultTransactor) UpdateBalance(opts *bind.TransactOpts, token common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "updateBalance", token, trader, amount)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x33f2f2e7.
//
// Solidity: function updateBalance(address token, address trader, int256 amount) returns()
func (_CollateralVault *CollateralVaultSession) UpdateBalance(token common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.UpdateBalance(&_CollateralVault.TransactOpts, token, trader, amount)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x33f2f2e7.
//
// Solidity: function updateBalance(address token, address trader, int256 amount) returns()
func (_CollateralVault *CollateralVaultTransactorSession) UpdateBalance(token common.Address, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.UpdateBalance(&_CollateralVault.TransactOpts, token, trader, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralVault *CollateralVaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralVault *CollateralVaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.UpgradeToAndCall(&_CollateralVault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CollateralVault *CollateralVaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.UpgradeToAndCall(&_CollateralVault.TransactOpts, newImplementation, data)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultTransactor) WithdrawCollateral(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "withdrawCollateral", requestId)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultSession) WithdrawCollateral(requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateral(&_CollateralVault.TransactOpts, requestId)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 requestId) returns()
func (_CollateralVault *CollateralVaultTransactorSession) WithdrawCollateral(requestId *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateral(&_CollateralVault.TransactOpts, requestId)
}

// WithdrawCollateralSigned is a paid mutator transaction binding the contract method 0xf6467852.
//
// Solidity: function withdrawCollateralSigned(address token, uint256 amount, uint256 expiry, bytes signature) returns()
func (_CollateralVault *CollateralVaultTransactor) WithdrawCollateralSigned(opts *bind.TransactOpts, token common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "withdrawCollateralSigned", token, amount, expiry, signature)
}

// WithdrawCollateralSigned is a paid mutator transaction binding the contract method 0xf6467852.
//
// Solidity: function withdrawCollateralSigned(address token, uint256 amount, uint256 expiry, bytes signature) returns()
func (_CollateralVault *CollateralVaultSession) WithdrawCollateralSigned(token common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateralSigned(&_CollateralVault.TransactOpts, token, amount, expiry, signature)
}

// WithdrawCollateralSigned is a paid mutator transaction binding the contract method 0xf6467852.
//
// Solidity: function withdrawCollateralSigned(address token, uint256 amount, uint256 expiry, bytes signature) returns()
func (_CollateralVault *CollateralVaultTransactorSession) WithdrawCollateralSigned(token common.Address, amount *big.Int, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateralSigned(&_CollateralVault.TransactOpts, token, amount, expiry, signature)
}

// CollateralVaultCollateralBelowZeroIterator is returned from FilterCollateralBelowZero and is used to iterate over the raw logs and unpacked data for CollateralBelowZero events raised by the CollateralVault contract.
type CollateralVaultCollateralBelowZeroIterator struct {
	Event *CollateralVaultCollateralBelowZero // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralBelowZeroIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralBelowZero)
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
		it.Event = new(CollateralVaultCollateralBelowZero)
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
func (it *CollateralVaultCollateralBelowZeroIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralBelowZeroIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralBelowZero represents a CollateralBelowZero event raised by the CollateralVault contract.
type CollateralVaultCollateralBelowZero struct {
	Trader common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralBelowZero is a free log retrieval operation binding the contract event 0xe146f09bc7254ace77d0523f7a2456f6c20bfefd8453954379d834962709fce4.
//
// Solidity: event CollateralBelowZero(address indexed trader, address indexed token, int256 amount)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralBelowZero(opts *bind.FilterOpts, trader []common.Address, token []common.Address) (*CollateralVaultCollateralBelowZeroIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralBelowZero", traderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralBelowZeroIterator{contract: _CollateralVault.contract, event: "CollateralBelowZero", logs: logs, sub: sub}, nil
}

// WatchCollateralBelowZero is a free log subscription operation binding the contract event 0xe146f09bc7254ace77d0523f7a2456f6c20bfefd8453954379d834962709fce4.
//
// Solidity: event CollateralBelowZero(address indexed trader, address indexed token, int256 amount)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralBelowZero(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralBelowZero, trader []common.Address, token []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralBelowZero", traderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralBelowZero)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralBelowZero", log); err != nil {
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

// ParseCollateralBelowZero is a log parse operation binding the contract event 0xe146f09bc7254ace77d0523f7a2456f6c20bfefd8453954379d834962709fce4.
//
// Solidity: event CollateralBelowZero(address indexed trader, address indexed token, int256 amount)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralBelowZero(log types.Log) (*CollateralVaultCollateralBelowZero, error) {
	event := new(CollateralVaultCollateralBelowZero)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralBelowZero", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the CollateralVault contract.
type CollateralVaultCollateralDepositedIterator struct {
	Event *CollateralVaultCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralDeposited)
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
		it.Event = new(CollateralVaultCollateralDeposited)
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
func (it *CollateralVaultCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralDeposited represents a CollateralDeposited event raised by the CollateralVault contract.
type CollateralVaultCollateralDeposited struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xd7243f6f8212d5188fd054141cf6ea89cfc0d91facb8c3afe2f88a1358480142.
//
// Solidity: event CollateralDeposited(address indexed trader, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, trader []common.Address) (*CollateralVaultCollateralDepositedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralDeposited", traderRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralDepositedIterator{contract: _CollateralVault.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xd7243f6f8212d5188fd054141cf6ea89cfc0d91facb8c3afe2f88a1358480142.
//
// Solidity: event CollateralDeposited(address indexed trader, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralDeposited, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralDeposited", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralDeposited)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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

// ParseCollateralDeposited is a log parse operation binding the contract event 0xd7243f6f8212d5188fd054141cf6ea89cfc0d91facb8c3afe2f88a1358480142.
//
// Solidity: event CollateralDeposited(address indexed trader, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralDeposited(log types.Log) (*CollateralVaultCollateralDeposited, error) {
	event := new(CollateralVaultCollateralDeposited)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultCollateralWithdrawnIterator is returned from FilterCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for CollateralWithdrawn events raised by the CollateralVault contract.
type CollateralVaultCollateralWithdrawnIterator struct {
	Event *CollateralVaultCollateralWithdrawn // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralWithdrawn)
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
		it.Event = new(CollateralVaultCollateralWithdrawn)
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
func (it *CollateralVaultCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralWithdrawn represents a CollateralWithdrawn event raised by the CollateralVault contract.
type CollateralVaultCollateralWithdrawn struct {
	Trader common.Address
	Amount *big.Int
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralWithdrawn is a free log retrieval operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount, uint256 expiry)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralWithdrawn(opts *bind.FilterOpts, trader []common.Address) (*CollateralVaultCollateralWithdrawnIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralWithdrawn", traderRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralWithdrawnIterator{contract: _CollateralVault.contract, event: "CollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCollateralWithdrawn is a free log subscription operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount, uint256 expiry)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralWithdrawn, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralWithdrawn", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralWithdrawn)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
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

// ParseCollateralWithdrawn is a log parse operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount, uint256 expiry)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralWithdrawn(log types.Log) (*CollateralVaultCollateralWithdrawn, error) {
	event := new(CollateralVaultCollateralWithdrawn)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CollateralVault contract.
type CollateralVaultInitializedIterator struct {
	Event *CollateralVaultInitialized // Event containing the contract specifics and raw log

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
func (it *CollateralVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultInitialized)
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
		it.Event = new(CollateralVaultInitialized)
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
func (it *CollateralVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultInitialized represents a Initialized event raised by the CollateralVault contract.
type CollateralVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralVault *CollateralVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*CollateralVaultInitializedIterator, error) {

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CollateralVaultInitializedIterator{contract: _CollateralVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralVault *CollateralVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CollateralVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultInitialized)
				if err := _CollateralVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CollateralVault *CollateralVaultFilterer) ParseInitialized(log types.Log) (*CollateralVaultInitialized, error) {
	event := new(CollateralVaultInitialized)
	if err := _CollateralVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CollateralVault contract.
type CollateralVaultOwnershipTransferredIterator struct {
	Event *CollateralVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollateralVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultOwnershipTransferred)
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
		it.Event = new(CollateralVaultOwnershipTransferred)
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
func (it *CollateralVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultOwnershipTransferred represents a OwnershipTransferred event raised by the CollateralVault contract.
type CollateralVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollateralVault *CollateralVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollateralVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultOwnershipTransferredIterator{contract: _CollateralVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollateralVault *CollateralVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultOwnershipTransferred)
				if err := _CollateralVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CollateralVault *CollateralVaultFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralVaultOwnershipTransferred, error) {
	event := new(CollateralVaultOwnershipTransferred)
	if err := _CollateralVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CollateralVault contract.
type CollateralVaultPausedIterator struct {
	Event *CollateralVaultPaused // Event containing the contract specifics and raw log

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
func (it *CollateralVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultPaused)
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
		it.Event = new(CollateralVaultPaused)
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
func (it *CollateralVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultPaused represents a Paused event raised by the CollateralVault contract.
type CollateralVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CollateralVault *CollateralVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*CollateralVaultPausedIterator, error) {

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CollateralVaultPausedIterator{contract: _CollateralVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CollateralVault *CollateralVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CollateralVaultPaused) (event.Subscription, error) {

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultPaused)
				if err := _CollateralVault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CollateralVault *CollateralVaultFilterer) ParsePaused(log types.Log) (*CollateralVaultPaused, error) {
	event := new(CollateralVaultPaused)
	if err := _CollateralVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultTokenAuthorizedIterator is returned from FilterTokenAuthorized and is used to iterate over the raw logs and unpacked data for TokenAuthorized events raised by the CollateralVault contract.
type CollateralVaultTokenAuthorizedIterator struct {
	Event *CollateralVaultTokenAuthorized // Event containing the contract specifics and raw log

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
func (it *CollateralVaultTokenAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultTokenAuthorized)
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
		it.Event = new(CollateralVaultTokenAuthorized)
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
func (it *CollateralVaultTokenAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultTokenAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultTokenAuthorized represents a TokenAuthorized event raised by the CollateralVault contract.
type CollateralVaultTokenAuthorized struct {
	Token      common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenAuthorized is a free log retrieval operation binding the contract event 0xdcef1c68f04eb69ace296de06030dc89d1110f7d56bca60edf1b83d9f402078b.
//
// Solidity: event TokenAuthorized(address indexed token, bool authorized)
func (_CollateralVault *CollateralVaultFilterer) FilterTokenAuthorized(opts *bind.FilterOpts, token []common.Address) (*CollateralVaultTokenAuthorizedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "TokenAuthorized", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultTokenAuthorizedIterator{contract: _CollateralVault.contract, event: "TokenAuthorized", logs: logs, sub: sub}, nil
}

// WatchTokenAuthorized is a free log subscription operation binding the contract event 0xdcef1c68f04eb69ace296de06030dc89d1110f7d56bca60edf1b83d9f402078b.
//
// Solidity: event TokenAuthorized(address indexed token, bool authorized)
func (_CollateralVault *CollateralVaultFilterer) WatchTokenAuthorized(opts *bind.WatchOpts, sink chan<- *CollateralVaultTokenAuthorized, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "TokenAuthorized", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultTokenAuthorized)
				if err := _CollateralVault.contract.UnpackLog(event, "TokenAuthorized", log); err != nil {
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

// ParseTokenAuthorized is a log parse operation binding the contract event 0xdcef1c68f04eb69ace296de06030dc89d1110f7d56bca60edf1b83d9f402078b.
//
// Solidity: event TokenAuthorized(address indexed token, bool authorized)
func (_CollateralVault *CollateralVaultFilterer) ParseTokenAuthorized(log types.Log) (*CollateralVaultTokenAuthorized, error) {
	event := new(CollateralVaultTokenAuthorized)
	if err := _CollateralVault.contract.UnpackLog(event, "TokenAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CollateralVault contract.
type CollateralVaultUnpausedIterator struct {
	Event *CollateralVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *CollateralVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultUnpaused)
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
		it.Event = new(CollateralVaultUnpaused)
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
func (it *CollateralVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultUnpaused represents a Unpaused event raised by the CollateralVault contract.
type CollateralVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CollateralVault *CollateralVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CollateralVaultUnpausedIterator, error) {

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CollateralVaultUnpausedIterator{contract: _CollateralVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CollateralVault *CollateralVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CollateralVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultUnpaused)
				if err := _CollateralVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CollateralVault *CollateralVaultFilterer) ParseUnpaused(log types.Log) (*CollateralVaultUnpaused, error) {
	event := new(CollateralVaultUnpaused)
	if err := _CollateralVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CollateralVault contract.
type CollateralVaultUpgradedIterator struct {
	Event *CollateralVaultUpgraded // Event containing the contract specifics and raw log

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
func (it *CollateralVaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultUpgraded)
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
		it.Event = new(CollateralVaultUpgraded)
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
func (it *CollateralVaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultUpgraded represents a Upgraded event raised by the CollateralVault contract.
type CollateralVaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralVault *CollateralVaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CollateralVaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultUpgradedIterator{contract: _CollateralVault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralVault *CollateralVaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CollateralVaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultUpgraded)
				if err := _CollateralVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CollateralVault *CollateralVaultFilterer) ParseUpgraded(log types.Log) (*CollateralVaultUpgraded, error) {
	event := new(CollateralVaultUpgraded)
	if err := _CollateralVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultWithdrawalCancelledIterator is returned from FilterWithdrawalCancelled and is used to iterate over the raw logs and unpacked data for WithdrawalCancelled events raised by the CollateralVault contract.
type CollateralVaultWithdrawalCancelledIterator struct {
	Event *CollateralVaultWithdrawalCancelled // Event containing the contract specifics and raw log

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
func (it *CollateralVaultWithdrawalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultWithdrawalCancelled)
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
		it.Event = new(CollateralVaultWithdrawalCancelled)
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
func (it *CollateralVaultWithdrawalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultWithdrawalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultWithdrawalCancelled represents a WithdrawalCancelled event raised by the CollateralVault contract.
type CollateralVaultWithdrawalCancelled struct {
	Trader    common.Address
	Token     common.Address
	RequestId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCancelled is a free log retrieval operation binding the contract event 0xc0b6b4b152b639767dd4ffd28f2a0f05b4ffd672264a6c4d34447185a03cc517.
//
// Solidity: event WithdrawalCancelled(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) FilterWithdrawalCancelled(opts *bind.FilterOpts, trader []common.Address, token []common.Address, requestId []*big.Int) (*CollateralVaultWithdrawalCancelledIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "WithdrawalCancelled", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultWithdrawalCancelledIterator{contract: _CollateralVault.contract, event: "WithdrawalCancelled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCancelled is a free log subscription operation binding the contract event 0xc0b6b4b152b639767dd4ffd28f2a0f05b4ffd672264a6c4d34447185a03cc517.
//
// Solidity: event WithdrawalCancelled(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) WatchWithdrawalCancelled(opts *bind.WatchOpts, sink chan<- *CollateralVaultWithdrawalCancelled, trader []common.Address, token []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "WithdrawalCancelled", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultWithdrawalCancelled)
				if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
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

// ParseWithdrawalCancelled is a log parse operation binding the contract event 0xc0b6b4b152b639767dd4ffd28f2a0f05b4ffd672264a6c4d34447185a03cc517.
//
// Solidity: event WithdrawalCancelled(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) ParseWithdrawalCancelled(log types.Log) (*CollateralVaultWithdrawalCancelled, error) {
	event := new(CollateralVaultWithdrawalCancelled)
	if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the CollateralVault contract.
type CollateralVaultWithdrawalCompletedIterator struct {
	Event *CollateralVaultWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *CollateralVaultWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultWithdrawalCompleted)
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
		it.Event = new(CollateralVaultWithdrawalCompleted)
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
func (it *CollateralVaultWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultWithdrawalCompleted represents a WithdrawalCompleted event raised by the CollateralVault contract.
type CollateralVaultWithdrawalCompleted struct {
	Trader    common.Address
	Token     common.Address
	RequestId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xe01911a56be554fae2e8c04ebcc1c56c757ce4583c883b878180a6ce99fc8a09.
//
// Solidity: event WithdrawalCompleted(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts, trader []common.Address, token []common.Address, requestId []*big.Int) (*CollateralVaultWithdrawalCompletedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "WithdrawalCompleted", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultWithdrawalCompletedIterator{contract: _CollateralVault.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xe01911a56be554fae2e8c04ebcc1c56c757ce4583c883b878180a6ce99fc8a09.
//
// Solidity: event WithdrawalCompleted(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *CollateralVaultWithdrawalCompleted, trader []common.Address, token []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "WithdrawalCompleted", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultWithdrawalCompleted)
				if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xe01911a56be554fae2e8c04ebcc1c56c757ce4583c883b878180a6ce99fc8a09.
//
// Solidity: event WithdrawalCompleted(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount)
func (_CollateralVault *CollateralVaultFilterer) ParseWithdrawalCompleted(log types.Log) (*CollateralVaultWithdrawalCompleted, error) {
	event := new(CollateralVaultWithdrawalCompleted)
	if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the CollateralVault contract.
type CollateralVaultWithdrawalRequestedIterator struct {
	Event *CollateralVaultWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *CollateralVaultWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultWithdrawalRequested)
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
		it.Event = new(CollateralVaultWithdrawalRequested)
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
func (it *CollateralVaultWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultWithdrawalRequested represents a WithdrawalRequested event raised by the CollateralVault contract.
type CollateralVaultWithdrawalRequested struct {
	Trader     common.Address
	Token      common.Address
	RequestId  *big.Int
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: event WithdrawalRequested(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, trader []common.Address, token []common.Address, requestId []*big.Int) (*CollateralVaultWithdrawalRequestedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "WithdrawalRequested", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultWithdrawalRequestedIterator{contract: _CollateralVault.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: event WithdrawalRequested(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *CollateralVaultWithdrawalRequested, trader []common.Address, token []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "WithdrawalRequested", traderRule, tokenRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultWithdrawalRequested)
				if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x3aeb15af61588a39bcfafb19ed853140d195c2a924537afbf9a6d04348e76a69.
//
// Solidity: event WithdrawalRequested(address indexed trader, address indexed token, uint256 indexed requestId, uint256 amount, uint256 unlockTime)
func (_CollateralVault *CollateralVaultFilterer) ParseWithdrawalRequested(log types.Log) (*CollateralVaultWithdrawalRequested, error) {
	event := new(CollateralVaultWithdrawalRequested)
	if err := _CollateralVault.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
