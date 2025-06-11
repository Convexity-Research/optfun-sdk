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

// TakerQ is an auto generated low-level Go binding around an user-defined struct.
type TakerQ struct {
	Trader common.Address
	Size   *big.Int
}

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONTRACT_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeCycle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"badDebt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cursor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cycles\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSettled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"strikePrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"settlementPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"denominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"detCA\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"detCB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"detPA\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"detPB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fixPrice\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOraclePrice\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"levels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"vol\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"head\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tail\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"makerIds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"long\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"long\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"makerFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"makerQ\",\"inputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"size\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"next\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"key\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"prev\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"midCA\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"midCB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"midPA\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"midPB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paidOut\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeOrder\",\"inputs\":[{\"name\":\"option\",\"type\":\"uint8\",\"internalType\":\"enumOptionType\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"placeOrder\",\"inputs\":[{\"name\":\"option\",\"type\":\"uint8\",\"internalType\":\"enumOptionType\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"longCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"shortCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"longPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"shortPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingLongCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingShortCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingLongPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingShortPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedForwarder\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleChunk\",\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"done\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"short\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"short\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startCycle\",\"inputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"summaries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"takerFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tqHead\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"traders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"viewTakerQueue\",\"inputs\":[{\"name\":\"isPut\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isBid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTakerQ[]\",\"components\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"size\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralWithdrawn\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CycleSettled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CycleStarted\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"strike\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Liquidated\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCancelled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderFilled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isPut\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"btcPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPlaced\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isBuy\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isPut\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFixed\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTSIZE is a free data retrieval call binding the contract method 0xb3413bda.
//
// Solidity: function CONTRACT_SIZE() view returns(uint256)
func (_Market *MarketCaller) CONTRACTSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "CONTRACT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONTRACTSIZE is a free data retrieval call binding the contract method 0xb3413bda.
//
// Solidity: function CONTRACT_SIZE() view returns(uint256)
func (_Market *MarketSession) CONTRACTSIZE() (*big.Int, error) {
	return _Market.Contract.CONTRACTSIZE(&_Market.CallOpts)
}

// CONTRACTSIZE is a free data retrieval call binding the contract method 0xb3413bda.
//
// Solidity: function CONTRACT_SIZE() view returns(uint256)
func (_Market *MarketCallerSession) CONTRACTSIZE() (*big.Int, error) {
	return _Market.Contract.CONTRACTSIZE(&_Market.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Market *MarketCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Market *MarketSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Market.Contract.UPGRADEINTERFACEVERSION(&_Market.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Market *MarketCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Market.Contract.UPGRADEINTERFACEVERSION(&_Market.CallOpts)
}

// ActiveCycle is a free data retrieval call binding the contract method 0x8706fb16.
//
// Solidity: function activeCycle() view returns(uint256)
func (_Market *MarketCaller) ActiveCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "activeCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveCycle is a free data retrieval call binding the contract method 0x8706fb16.
//
// Solidity: function activeCycle() view returns(uint256)
func (_Market *MarketSession) ActiveCycle() (*big.Int, error) {
	return _Market.Contract.ActiveCycle(&_Market.CallOpts)
}

// ActiveCycle is a free data retrieval call binding the contract method 0x8706fb16.
//
// Solidity: function activeCycle() view returns(uint256)
func (_Market *MarketCallerSession) ActiveCycle() (*big.Int, error) {
	return _Market.Contract.ActiveCycle(&_Market.CallOpts)
}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Market *MarketCaller) BadDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "badDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Market *MarketSession) BadDebt() (*big.Int, error) {
	return _Market.Contract.BadDebt(&_Market.CallOpts)
}

// BadDebt is a free data retrieval call binding the contract method 0xbbcac557.
//
// Solidity: function badDebt() view returns(uint256)
func (_Market *MarketCallerSession) BadDebt() (*big.Int, error) {
	return _Market.Contract.BadDebt(&_Market.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Market *MarketCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Market *MarketSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Balances(&_Market.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Market *MarketCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Balances(&_Market.CallOpts, arg0)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint64)
func (_Market *MarketCaller) CollateralDecimals(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "collateralDecimals")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint64)
func (_Market *MarketSession) CollateralDecimals() (uint64, error) {
	return _Market.Contract.CollateralDecimals(&_Market.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint64)
func (_Market *MarketCallerSession) CollateralDecimals() (uint64, error) {
	return _Market.Contract.CollateralDecimals(&_Market.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Market *MarketCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Market *MarketSession) CollateralToken() (common.Address, error) {
	return _Market.Contract.CollateralToken(&_Market.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_Market *MarketCallerSession) CollateralToken() (common.Address, error) {
	return _Market.Contract.CollateralToken(&_Market.CallOpts)
}

// Cursor is a free data retrieval call binding the contract method 0x36a0b9ef.
//
// Solidity: function cursor() view returns(uint256)
func (_Market *MarketCaller) Cursor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "cursor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cursor is a free data retrieval call binding the contract method 0x36a0b9ef.
//
// Solidity: function cursor() view returns(uint256)
func (_Market *MarketSession) Cursor() (*big.Int, error) {
	return _Market.Contract.Cursor(&_Market.CallOpts)
}

// Cursor is a free data retrieval call binding the contract method 0x36a0b9ef.
//
// Solidity: function cursor() view returns(uint256)
func (_Market *MarketCallerSession) Cursor() (*big.Int, error) {
	return _Market.Contract.Cursor(&_Market.CallOpts)
}

// Cycles is a free data retrieval call binding the contract method 0xafbce3b9.
//
// Solidity: function cycles(uint256 ) view returns(bool active, bool isSettled, uint64 strikePrice, uint64 settlementPrice)
func (_Market *MarketCaller) Cycles(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active          bool
	IsSettled       bool
	StrikePrice     uint64
	SettlementPrice uint64
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "cycles", arg0)

	outstruct := new(struct {
		Active          bool
		IsSettled       bool
		StrikePrice     uint64
		SettlementPrice uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsSettled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.StrikePrice = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.SettlementPrice = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// Cycles is a free data retrieval call binding the contract method 0xafbce3b9.
//
// Solidity: function cycles(uint256 ) view returns(bool active, bool isSettled, uint64 strikePrice, uint64 settlementPrice)
func (_Market *MarketSession) Cycles(arg0 *big.Int) (struct {
	Active          bool
	IsSettled       bool
	StrikePrice     uint64
	SettlementPrice uint64
}, error) {
	return _Market.Contract.Cycles(&_Market.CallOpts, arg0)
}

// Cycles is a free data retrieval call binding the contract method 0xafbce3b9.
//
// Solidity: function cycles(uint256 ) view returns(bool active, bool isSettled, uint64 strikePrice, uint64 settlementPrice)
func (_Market *MarketCallerSession) Cycles(arg0 *big.Int) (struct {
	Active          bool
	IsSettled       bool
	StrikePrice     uint64
	SettlementPrice uint64
}, error) {
	return _Market.Contract.Cycles(&_Market.CallOpts, arg0)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Market *MarketCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Market *MarketSession) Denominator() (*big.Int, error) {
	return _Market.Contract.Denominator(&_Market.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Market *MarketCallerSession) Denominator() (*big.Int, error) {
	return _Market.Contract.Denominator(&_Market.CallOpts)
}

// DetCA is a free data retrieval call binding the contract method 0xd8a19b0c.
//
// Solidity: function detCA(uint16 ) view returns(uint256)
func (_Market *MarketCaller) DetCA(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "detCA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetCA is a free data retrieval call binding the contract method 0xd8a19b0c.
//
// Solidity: function detCA(uint16 ) view returns(uint256)
func (_Market *MarketSession) DetCA(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetCA(&_Market.CallOpts, arg0)
}

// DetCA is a free data retrieval call binding the contract method 0xd8a19b0c.
//
// Solidity: function detCA(uint16 ) view returns(uint256)
func (_Market *MarketCallerSession) DetCA(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetCA(&_Market.CallOpts, arg0)
}

// DetCB is a free data retrieval call binding the contract method 0xd8569551.
//
// Solidity: function detCB(uint16 ) view returns(uint256)
func (_Market *MarketCaller) DetCB(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "detCB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetCB is a free data retrieval call binding the contract method 0xd8569551.
//
// Solidity: function detCB(uint16 ) view returns(uint256)
func (_Market *MarketSession) DetCB(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetCB(&_Market.CallOpts, arg0)
}

// DetCB is a free data retrieval call binding the contract method 0xd8569551.
//
// Solidity: function detCB(uint16 ) view returns(uint256)
func (_Market *MarketCallerSession) DetCB(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetCB(&_Market.CallOpts, arg0)
}

// DetPA is a free data retrieval call binding the contract method 0xa109fb5b.
//
// Solidity: function detPA(uint16 ) view returns(uint256)
func (_Market *MarketCaller) DetPA(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "detPA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetPA is a free data retrieval call binding the contract method 0xa109fb5b.
//
// Solidity: function detPA(uint16 ) view returns(uint256)
func (_Market *MarketSession) DetPA(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetPA(&_Market.CallOpts, arg0)
}

// DetPA is a free data retrieval call binding the contract method 0xa109fb5b.
//
// Solidity: function detPA(uint16 ) view returns(uint256)
func (_Market *MarketCallerSession) DetPA(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetPA(&_Market.CallOpts, arg0)
}

// DetPB is a free data retrieval call binding the contract method 0xda55afad.
//
// Solidity: function detPB(uint16 ) view returns(uint256)
func (_Market *MarketCaller) DetPB(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "detPB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetPB is a free data retrieval call binding the contract method 0xda55afad.
//
// Solidity: function detPB(uint16 ) view returns(uint256)
func (_Market *MarketSession) DetPB(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetPB(&_Market.CallOpts, arg0)
}

// DetPB is a free data retrieval call binding the contract method 0xda55afad.
//
// Solidity: function detPB(uint16 ) view returns(uint256)
func (_Market *MarketCallerSession) DetPB(arg0 uint16) (*big.Int, error) {
	return _Market.Contract.DetPB(&_Market.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Market *MarketCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Market *MarketSession) FeeRecipient() (common.Address, error) {
	return _Market.Contract.FeeRecipient(&_Market.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Market *MarketCallerSession) FeeRecipient() (common.Address, error) {
	return _Market.Contract.FeeRecipient(&_Market.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x2af78bd1.
//
// Solidity: function getOraclePrice(uint32 index) view returns(uint64)
func (_Market *MarketCaller) GetOraclePrice(opts *bind.CallOpts, index uint32) (uint64, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getOraclePrice", index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x2af78bd1.
//
// Solidity: function getOraclePrice(uint32 index) view returns(uint64)
func (_Market *MarketSession) GetOraclePrice(index uint32) (uint64, error) {
	return _Market.Contract.GetOraclePrice(&_Market.CallOpts, index)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x2af78bd1.
//
// Solidity: function getOraclePrice(uint32 index) view returns(uint64)
func (_Market *MarketCallerSession) GetOraclePrice(index uint32) (uint64, error) {
	return _Market.Contract.GetOraclePrice(&_Market.CallOpts, index)
}

// InList is a free data retrieval call binding the contract method 0x25b198eb.
//
// Solidity: function inList(address ) view returns(bool)
func (_Market *MarketCaller) InList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "inList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InList is a free data retrieval call binding the contract method 0x25b198eb.
//
// Solidity: function inList(address ) view returns(bool)
func (_Market *MarketSession) InList(arg0 common.Address) (bool, error) {
	return _Market.Contract.InList(&_Market.CallOpts, arg0)
}

// InList is a free data retrieval call binding the contract method 0x25b198eb.
//
// Solidity: function inList(address ) view returns(bool)
func (_Market *MarketCallerSession) InList(arg0 common.Address) (bool, error) {
	return _Market.Contract.InList(&_Market.CallOpts, arg0)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Market.Contract.IsTrustedForwarder(&_Market.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Market.Contract.IsTrustedForwarder(&_Market.CallOpts, forwarder)
}

// Levels is a free data retrieval call binding the contract method 0x45805991.
//
// Solidity: function levels(uint32 ) view returns(uint128 vol, uint16 head, uint16 tail)
func (_Market *MarketCaller) Levels(opts *bind.CallOpts, arg0 uint32) (struct {
	Vol  *big.Int
	Head uint16
	Tail uint16
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "levels", arg0)

	outstruct := new(struct {
		Vol  *big.Int
		Head uint16
		Tail uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Vol = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Head = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Tail = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// Levels is a free data retrieval call binding the contract method 0x45805991.
//
// Solidity: function levels(uint32 ) view returns(uint128 vol, uint16 head, uint16 tail)
func (_Market *MarketSession) Levels(arg0 uint32) (struct {
	Vol  *big.Int
	Head uint16
	Tail uint16
}, error) {
	return _Market.Contract.Levels(&_Market.CallOpts, arg0)
}

// Levels is a free data retrieval call binding the contract method 0x45805991.
//
// Solidity: function levels(uint32 ) view returns(uint128 vol, uint16 head, uint16 tail)
func (_Market *MarketCallerSession) Levels(arg0 uint32) (struct {
	Vol  *big.Int
	Head uint16
	Tail uint16
}, error) {
	return _Market.Contract.Levels(&_Market.CallOpts, arg0)
}

// MakerFeeBps is a free data retrieval call binding the contract method 0xc8d594db.
//
// Solidity: function makerFeeBps() view returns(int256)
func (_Market *MarketCaller) MakerFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "makerFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerFeeBps is a free data retrieval call binding the contract method 0xc8d594db.
//
// Solidity: function makerFeeBps() view returns(int256)
func (_Market *MarketSession) MakerFeeBps() (*big.Int, error) {
	return _Market.Contract.MakerFeeBps(&_Market.CallOpts)
}

// MakerFeeBps is a free data retrieval call binding the contract method 0xc8d594db.
//
// Solidity: function makerFeeBps() view returns(int256)
func (_Market *MarketCallerSession) MakerFeeBps() (*big.Int, error) {
	return _Market.Contract.MakerFeeBps(&_Market.CallOpts)
}

// MakerQ is a free data retrieval call binding the contract method 0x93067139.
//
// Solidity: function makerQ(uint16 ) view returns(address trader, uint128 size, uint16 next, uint32 key, uint16 prev)
func (_Market *MarketCaller) MakerQ(opts *bind.CallOpts, arg0 uint16) (struct {
	Trader common.Address
	Size   *big.Int
	Next   uint16
	Key    uint32
	Prev   uint16
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "makerQ", arg0)

	outstruct := new(struct {
		Trader common.Address
		Size   *big.Int
		Next   uint16
		Key    uint32
		Prev   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Size = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Next = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Key = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Prev = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// MakerQ is a free data retrieval call binding the contract method 0x93067139.
//
// Solidity: function makerQ(uint16 ) view returns(address trader, uint128 size, uint16 next, uint32 key, uint16 prev)
func (_Market *MarketSession) MakerQ(arg0 uint16) (struct {
	Trader common.Address
	Size   *big.Int
	Next   uint16
	Key    uint32
	Prev   uint16
}, error) {
	return _Market.Contract.MakerQ(&_Market.CallOpts, arg0)
}

// MakerQ is a free data retrieval call binding the contract method 0x93067139.
//
// Solidity: function makerQ(uint16 ) view returns(address trader, uint128 size, uint16 next, uint32 key, uint16 prev)
func (_Market *MarketCallerSession) MakerQ(arg0 uint16) (struct {
	Trader common.Address
	Size   *big.Int
	Next   uint16
	Key    uint32
	Prev   uint16
}, error) {
	return _Market.Contract.MakerQ(&_Market.CallOpts, arg0)
}

// MidCA is a free data retrieval call binding the contract method 0xaedb34cb.
//
// Solidity: function midCA(uint8 ) view returns(uint256)
func (_Market *MarketCaller) MidCA(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "midCA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidCA is a free data retrieval call binding the contract method 0xaedb34cb.
//
// Solidity: function midCA(uint8 ) view returns(uint256)
func (_Market *MarketSession) MidCA(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidCA(&_Market.CallOpts, arg0)
}

// MidCA is a free data retrieval call binding the contract method 0xaedb34cb.
//
// Solidity: function midCA(uint8 ) view returns(uint256)
func (_Market *MarketCallerSession) MidCA(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidCA(&_Market.CallOpts, arg0)
}

// MidCB is a free data retrieval call binding the contract method 0x7e593e3d.
//
// Solidity: function midCB(uint8 ) view returns(uint256)
func (_Market *MarketCaller) MidCB(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "midCB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidCB is a free data retrieval call binding the contract method 0x7e593e3d.
//
// Solidity: function midCB(uint8 ) view returns(uint256)
func (_Market *MarketSession) MidCB(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidCB(&_Market.CallOpts, arg0)
}

// MidCB is a free data retrieval call binding the contract method 0x7e593e3d.
//
// Solidity: function midCB(uint8 ) view returns(uint256)
func (_Market *MarketCallerSession) MidCB(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidCB(&_Market.CallOpts, arg0)
}

// MidPA is a free data retrieval call binding the contract method 0xcee7c9cf.
//
// Solidity: function midPA(uint8 ) view returns(uint256)
func (_Market *MarketCaller) MidPA(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "midPA", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidPA is a free data retrieval call binding the contract method 0xcee7c9cf.
//
// Solidity: function midPA(uint8 ) view returns(uint256)
func (_Market *MarketSession) MidPA(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidPA(&_Market.CallOpts, arg0)
}

// MidPA is a free data retrieval call binding the contract method 0xcee7c9cf.
//
// Solidity: function midPA(uint8 ) view returns(uint256)
func (_Market *MarketCallerSession) MidPA(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidPA(&_Market.CallOpts, arg0)
}

// MidPB is a free data retrieval call binding the contract method 0xe04ebfda.
//
// Solidity: function midPB(uint8 ) view returns(uint256)
func (_Market *MarketCaller) MidPB(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "midPB", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidPB is a free data retrieval call binding the contract method 0xe04ebfda.
//
// Solidity: function midPB(uint8 ) view returns(uint256)
func (_Market *MarketSession) MidPB(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidPB(&_Market.CallOpts, arg0)
}

// MidPB is a free data retrieval call binding the contract method 0xe04ebfda.
//
// Solidity: function midPB(uint8 ) view returns(uint256)
func (_Market *MarketCallerSession) MidPB(arg0 uint8) (*big.Int, error) {
	return _Market.Contract.MidPB(&_Market.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Market *MarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Market *MarketSession) Name() (string, error) {
	return _Market.Contract.Name(&_Market.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Market *MarketCallerSession) Name() (string, error) {
	return _Market.Contract.Name(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// PaidOut is a free data retrieval call binding the contract method 0x5c76ca2d.
//
// Solidity: function paidOut() view returns(uint256)
func (_Market *MarketCaller) PaidOut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "paidOut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaidOut is a free data retrieval call binding the contract method 0x5c76ca2d.
//
// Solidity: function paidOut() view returns(uint256)
func (_Market *MarketSession) PaidOut() (*big.Int, error) {
	return _Market.Contract.PaidOut(&_Market.CallOpts)
}

// PaidOut is a free data retrieval call binding the contract method 0x5c76ca2d.
//
// Solidity: function paidOut() view returns(uint256)
func (_Market *MarketCallerSession) PaidOut() (*big.Int, error) {
	return _Market.Contract.PaidOut(&_Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCallerSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(uint32 longCalls, uint32 shortCalls, uint32 longPuts, uint32 shortPuts, uint32 pendingLongCalls, uint32 pendingShortCalls, uint32 pendingLongPuts, uint32 pendingShortPuts)
func (_Market *MarketCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LongCalls         uint32
	ShortCalls        uint32
	LongPuts          uint32
	ShortPuts         uint32
	PendingLongCalls  uint32
	PendingShortCalls uint32
	PendingLongPuts   uint32
	PendingShortPuts  uint32
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		LongCalls         uint32
		ShortCalls        uint32
		LongPuts          uint32
		ShortPuts         uint32
		PendingLongCalls  uint32
		PendingShortCalls uint32
		PendingLongPuts   uint32
		PendingShortPuts  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongCalls = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ShortCalls = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LongPuts = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.ShortPuts = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.PendingLongCalls = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.PendingShortCalls = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.PendingLongPuts = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.PendingShortPuts = *abi.ConvertType(out[7], new(uint32)).(*uint32)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(uint32 longCalls, uint32 shortCalls, uint32 longPuts, uint32 shortPuts, uint32 pendingLongCalls, uint32 pendingShortCalls, uint32 pendingLongPuts, uint32 pendingShortPuts)
func (_Market *MarketSession) Positions(arg0 *big.Int) (struct {
	LongCalls         uint32
	ShortCalls        uint32
	LongPuts          uint32
	ShortPuts         uint32
	PendingLongCalls  uint32
	PendingShortCalls uint32
	PendingLongPuts   uint32
	PendingShortPuts  uint32
}, error) {
	return _Market.Contract.Positions(&_Market.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) view returns(uint32 longCalls, uint32 shortCalls, uint32 longPuts, uint32 shortPuts, uint32 pendingLongCalls, uint32 pendingShortCalls, uint32 pendingLongPuts, uint32 pendingShortPuts)
func (_Market *MarketCallerSession) Positions(arg0 *big.Int) (struct {
	LongCalls         uint32
	ShortCalls        uint32
	LongPuts          uint32
	ShortPuts         uint32
	PendingLongCalls  uint32
	PendingShortCalls uint32
	PendingLongPuts   uint32
	PendingShortPuts  uint32
}, error) {
	return _Market.Contract.Positions(&_Market.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Market *MarketCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Market *MarketSession) ProxiableUUID() ([32]byte, error) {
	return _Market.Contract.ProxiableUUID(&_Market.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Market *MarketCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Market.Contract.ProxiableUUID(&_Market.CallOpts)
}

// Summaries is a free data retrieval call binding the contract method 0x9696d83d.
//
// Solidity: function summaries(uint256 ) view returns(uint256)
func (_Market *MarketCaller) Summaries(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "summaries", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Summaries is a free data retrieval call binding the contract method 0x9696d83d.
//
// Solidity: function summaries(uint256 ) view returns(uint256)
func (_Market *MarketSession) Summaries(arg0 *big.Int) (*big.Int, error) {
	return _Market.Contract.Summaries(&_Market.CallOpts, arg0)
}

// Summaries is a free data retrieval call binding the contract method 0x9696d83d.
//
// Solidity: function summaries(uint256 ) view returns(uint256)
func (_Market *MarketCallerSession) Summaries(arg0 *big.Int) (*big.Int, error) {
	return _Market.Contract.Summaries(&_Market.CallOpts, arg0)
}

// TakerFeeBps is a free data retrieval call binding the contract method 0xd448eebc.
//
// Solidity: function takerFeeBps() view returns(int256)
func (_Market *MarketCaller) TakerFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "takerFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TakerFeeBps is a free data retrieval call binding the contract method 0xd448eebc.
//
// Solidity: function takerFeeBps() view returns(int256)
func (_Market *MarketSession) TakerFeeBps() (*big.Int, error) {
	return _Market.Contract.TakerFeeBps(&_Market.CallOpts)
}

// TakerFeeBps is a free data retrieval call binding the contract method 0xd448eebc.
//
// Solidity: function takerFeeBps() view returns(int256)
func (_Market *MarketCallerSession) TakerFeeBps() (*big.Int, error) {
	return _Market.Contract.TakerFeeBps(&_Market.CallOpts)
}

// TqHead is a free data retrieval call binding the contract method 0x337a7a3e.
//
// Solidity: function tqHead(uint256 , uint256 ) view returns(uint256)
func (_Market *MarketCaller) TqHead(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "tqHead", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TqHead is a free data retrieval call binding the contract method 0x337a7a3e.
//
// Solidity: function tqHead(uint256 , uint256 ) view returns(uint256)
func (_Market *MarketSession) TqHead(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.TqHead(&_Market.CallOpts, arg0, arg1)
}

// TqHead is a free data retrieval call binding the contract method 0x337a7a3e.
//
// Solidity: function tqHead(uint256 , uint256 ) view returns(uint256)
func (_Market *MarketCallerSession) TqHead(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.TqHead(&_Market.CallOpts, arg0, arg1)
}

// Traders is a free data retrieval call binding the contract method 0x135e563d.
//
// Solidity: function traders(uint256 ) view returns(address)
func (_Market *MarketCaller) Traders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "traders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Traders is a free data retrieval call binding the contract method 0x135e563d.
//
// Solidity: function traders(uint256 ) view returns(address)
func (_Market *MarketSession) Traders(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.Traders(&_Market.CallOpts, arg0)
}

// Traders is a free data retrieval call binding the contract method 0x135e563d.
//
// Solidity: function traders(uint256 ) view returns(address)
func (_Market *MarketCallerSession) Traders(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.Traders(&_Market.CallOpts, arg0)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Market *MarketCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Market *MarketSession) TrustedForwarder() (common.Address, error) {
	return _Market.Contract.TrustedForwarder(&_Market.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Market *MarketCallerSession) TrustedForwarder() (common.Address, error) {
	return _Market.Contract.TrustedForwarder(&_Market.CallOpts)
}

// ViewTakerQueue is a free data retrieval call binding the contract method 0xd27929a1.
//
// Solidity: function viewTakerQueue(bool isPut, bool isBid) view returns((address,uint96)[])
func (_Market *MarketCaller) ViewTakerQueue(opts *bind.CallOpts, isPut bool, isBid bool) ([]TakerQ, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "viewTakerQueue", isPut, isBid)

	if err != nil {
		return *new([]TakerQ), err
	}

	out0 := *abi.ConvertType(out[0], new([]TakerQ)).(*[]TakerQ)

	return out0, err

}

// ViewTakerQueue is a free data retrieval call binding the contract method 0xd27929a1.
//
// Solidity: function viewTakerQueue(bool isPut, bool isBid) view returns((address,uint96)[])
func (_Market *MarketSession) ViewTakerQueue(isPut bool, isBid bool) ([]TakerQ, error) {
	return _Market.Contract.ViewTakerQueue(&_Market.CallOpts, isPut, isBid)
}

// ViewTakerQueue is a free data retrieval call binding the contract method 0xd27929a1.
//
// Solidity: function viewTakerQueue(bool isPut, bool isBid) view returns((address,uint96)[])
func (_Market *MarketCallerSession) ViewTakerQueue(isPut bool, isBid bool) ([]TakerQ, error) {
	return _Market.Contract.ViewTakerQueue(&_Market.CallOpts, isPut, isBid)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Market *MarketCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Market *MarketSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Market.Contract.Whitelist(&_Market.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Market *MarketCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Market.Contract.Whitelist(&_Market.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Market *MarketTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Market *MarketSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelOrder(&_Market.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Market *MarketTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelOrder(&_Market.TransactOpts, orderId)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_Market *MarketTransactor) DepositCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "depositCollateral", amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_Market *MarketSession) DepositCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositCollateral(&_Market.TransactOpts, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_Market *MarketTransactorSession) DepositCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositCollateral(&_Market.TransactOpts, amount)
}

// FixPrice is a paid mutator transaction binding the contract method 0x75593a39.
//
// Solidity: function fixPrice() returns()
func (_Market *MarketTransactor) FixPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "fixPrice")
}

// FixPrice is a paid mutator transaction binding the contract method 0x75593a39.
//
// Solidity: function fixPrice() returns()
func (_Market *MarketSession) FixPrice() (*types.Transaction, error) {
	return _Market.Contract.FixPrice(&_Market.TransactOpts)
}

// FixPrice is a paid mutator transaction binding the contract method 0x75593a39.
//
// Solidity: function fixPrice() returns()
func (_Market *MarketTransactorSession) FixPrice() (*types.Transaction, error) {
	return _Market.Contract.FixPrice(&_Market.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder) returns()
func (_Market *MarketTransactor) Initialize(opts *bind.TransactOpts, _name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "initialize", _name, _feeRecipient, _collateralToken, _forwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder) returns()
func (_Market *MarketSession) Initialize(_name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _name, _feeRecipient, _collateralToken, _forwarder)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder) returns()
func (_Market *MarketTransactorSession) Initialize(_name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _name, _feeRecipient, _collateralToken, _forwarder)
}

// Liquidate is a paid mutator transaction binding the contract method 0x820b68cc.
//
// Solidity: function liquidate(uint16[] makerIds, address trader) returns()
func (_Market *MarketTransactor) Liquidate(opts *bind.TransactOpts, makerIds []uint16, trader common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "liquidate", makerIds, trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x820b68cc.
//
// Solidity: function liquidate(uint16[] makerIds, address trader) returns()
func (_Market *MarketSession) Liquidate(makerIds []uint16, trader common.Address) (*types.Transaction, error) {
	return _Market.Contract.Liquidate(&_Market.TransactOpts, makerIds, trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x820b68cc.
//
// Solidity: function liquidate(uint16[] makerIds, address trader) returns()
func (_Market *MarketTransactorSession) Liquidate(makerIds []uint16, trader common.Address) (*types.Transaction, error) {
	return _Market.Contract.Liquidate(&_Market.TransactOpts, makerIds, trader)
}

// Long is a paid mutator transaction binding the contract method 0x488de1c0.
//
// Solidity: function long(uint256 size, bytes signature) returns()
func (_Market *MarketTransactor) Long(opts *bind.TransactOpts, size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "long", size, signature)
}

// Long is a paid mutator transaction binding the contract method 0x488de1c0.
//
// Solidity: function long(uint256 size, bytes signature) returns()
func (_Market *MarketSession) Long(size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.Long(&_Market.TransactOpts, size, signature)
}

// Long is a paid mutator transaction binding the contract method 0x488de1c0.
//
// Solidity: function long(uint256 size, bytes signature) returns()
func (_Market *MarketTransactorSession) Long(size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.Long(&_Market.TransactOpts, size, signature)
}

// Long0 is a paid mutator transaction binding the contract method 0x9db4bcd6.
//
// Solidity: function long(uint256 size) returns()
func (_Market *MarketTransactor) Long0(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "long0", size)
}

// Long0 is a paid mutator transaction binding the contract method 0x9db4bcd6.
//
// Solidity: function long(uint256 size) returns()
func (_Market *MarketSession) Long0(size *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Long0(&_Market.TransactOpts, size)
}

// Long0 is a paid mutator transaction binding the contract method 0x9db4bcd6.
//
// Solidity: function long(uint256 size) returns()
func (_Market *MarketTransactorSession) Long0(size *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Long0(&_Market.TransactOpts, size)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Market *MarketTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Market *MarketSession) Pause() (*types.Transaction, error) {
	return _Market.Contract.Pause(&_Market.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Market *MarketTransactorSession) Pause() (*types.Transaction, error) {
	return _Market.Contract.Pause(&_Market.TransactOpts)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x758a55d4.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice, bytes signature) returns(uint256 orderId)
func (_Market *MarketTransactor) PlaceOrder(opts *bind.TransactOpts, option uint8, side uint8, size *big.Int, limitPrice *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "placeOrder", option, side, size, limitPrice, signature)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x758a55d4.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice, bytes signature) returns(uint256 orderId)
func (_Market *MarketSession) PlaceOrder(option uint8, side uint8, size *big.Int, limitPrice *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder(&_Market.TransactOpts, option, side, size, limitPrice, signature)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x758a55d4.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice, bytes signature) returns(uint256 orderId)
func (_Market *MarketTransactorSession) PlaceOrder(option uint8, side uint8, size *big.Int, limitPrice *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder(&_Market.TransactOpts, option, side, size, limitPrice, signature)
}

// PlaceOrder0 is a paid mutator transaction binding the contract method 0xeca15d9f.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice) returns(uint256 orderId)
func (_Market *MarketTransactor) PlaceOrder0(opts *bind.TransactOpts, option uint8, side uint8, size *big.Int, limitPrice *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "placeOrder0", option, side, size, limitPrice)
}

// PlaceOrder0 is a paid mutator transaction binding the contract method 0xeca15d9f.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice) returns(uint256 orderId)
func (_Market *MarketSession) PlaceOrder0(option uint8, side uint8, size *big.Int, limitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder0(&_Market.TransactOpts, option, side, size, limitPrice)
}

// PlaceOrder0 is a paid mutator transaction binding the contract method 0xeca15d9f.
//
// Solidity: function placeOrder(uint8 option, uint8 side, uint256 size, uint256 limitPrice) returns(uint256 orderId)
func (_Market *MarketTransactorSession) PlaceOrder0(option uint8, side uint8, size *big.Int, limitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder0(&_Market.TransactOpts, option, side, size, limitPrice)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Market *MarketTransactor) SetTrustedForwarder(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setTrustedForwarder", _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Market *MarketSession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetTrustedForwarder(&_Market.TransactOpts, _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Market *MarketTransactorSession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetTrustedForwarder(&_Market.TransactOpts, _forwarder)
}

// SettleChunk is a paid mutator transaction binding the contract method 0x6ee6e29d.
//
// Solidity: function settleChunk(uint256 max) returns(bool done)
func (_Market *MarketTransactor) SettleChunk(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "settleChunk", max)
}

// SettleChunk is a paid mutator transaction binding the contract method 0x6ee6e29d.
//
// Solidity: function settleChunk(uint256 max) returns(bool done)
func (_Market *MarketSession) SettleChunk(max *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SettleChunk(&_Market.TransactOpts, max)
}

// SettleChunk is a paid mutator transaction binding the contract method 0x6ee6e29d.
//
// Solidity: function settleChunk(uint256 max) returns(bool done)
func (_Market *MarketTransactorSession) SettleChunk(max *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SettleChunk(&_Market.TransactOpts, max)
}

// Short is a paid mutator transaction binding the contract method 0x97303b62.
//
// Solidity: function short(uint256 size, bytes signature) returns()
func (_Market *MarketTransactor) Short(opts *bind.TransactOpts, size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "short", size, signature)
}

// Short is a paid mutator transaction binding the contract method 0x97303b62.
//
// Solidity: function short(uint256 size, bytes signature) returns()
func (_Market *MarketSession) Short(size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.Short(&_Market.TransactOpts, size, signature)
}

// Short is a paid mutator transaction binding the contract method 0x97303b62.
//
// Solidity: function short(uint256 size, bytes signature) returns()
func (_Market *MarketTransactorSession) Short(size *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.Short(&_Market.TransactOpts, size, signature)
}

// Short0 is a paid mutator transaction binding the contract method 0xc4a2a1e8.
//
// Solidity: function short(uint256 size) returns()
func (_Market *MarketTransactor) Short0(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "short0", size)
}

// Short0 is a paid mutator transaction binding the contract method 0xc4a2a1e8.
//
// Solidity: function short(uint256 size) returns()
func (_Market *MarketSession) Short0(size *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Short0(&_Market.TransactOpts, size)
}

// Short0 is a paid mutator transaction binding the contract method 0xc4a2a1e8.
//
// Solidity: function short(uint256 size) returns()
func (_Market *MarketTransactorSession) Short0(size *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Short0(&_Market.TransactOpts, size)
}

// StartCycle is a paid mutator transaction binding the contract method 0xe50b9ecd.
//
// Solidity: function startCycle(uint256 expiry) returns()
func (_Market *MarketTransactor) StartCycle(opts *bind.TransactOpts, expiry *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "startCycle", expiry)
}

// StartCycle is a paid mutator transaction binding the contract method 0xe50b9ecd.
//
// Solidity: function startCycle(uint256 expiry) returns()
func (_Market *MarketSession) StartCycle(expiry *big.Int) (*types.Transaction, error) {
	return _Market.Contract.StartCycle(&_Market.TransactOpts, expiry)
}

// StartCycle is a paid mutator transaction binding the contract method 0xe50b9ecd.
//
// Solidity: function startCycle(uint256 expiry) returns()
func (_Market *MarketTransactorSession) StartCycle(expiry *big.Int) (*types.Transaction, error) {
	return _Market.Contract.StartCycle(&_Market.TransactOpts, expiry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Market *MarketTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Market *MarketSession) Unpause() (*types.Transaction, error) {
	return _Market.Contract.Unpause(&_Market.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Market *MarketTransactorSession) Unpause() (*types.Transaction, error) {
	return _Market.Contract.Unpause(&_Market.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Market *MarketTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Market *MarketSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Market.Contract.UpgradeToAndCall(&_Market.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Market *MarketTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Market.Contract.UpgradeToAndCall(&_Market.TransactOpts, newImplementation, data)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_Market *MarketTransactor) WithdrawCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "withdrawCollateral", amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_Market *MarketSession) WithdrawCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.WithdrawCollateral(&_Market.TransactOpts, amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_Market *MarketTransactorSession) WithdrawCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.WithdrawCollateral(&_Market.TransactOpts, amount)
}

// MarketCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the Market contract.
type MarketCollateralDepositedIterator struct {
	Event *MarketCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *MarketCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCollateralDeposited)
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
		it.Event = new(MarketCollateralDeposited)
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
func (it *MarketCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCollateralDeposited represents a CollateralDeposited event raised by the Market contract.
type MarketCollateralDeposited struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xd7243f6f8212d5188fd054141cf6ea89cfc0d91facb8c3afe2f88a1358480142.
//
// Solidity: event CollateralDeposited(address indexed trader, uint256 amount)
func (_Market *MarketFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, trader []common.Address) (*MarketCollateralDepositedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CollateralDeposited", traderRule)
	if err != nil {
		return nil, err
	}
	return &MarketCollateralDepositedIterator{contract: _Market.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xd7243f6f8212d5188fd054141cf6ea89cfc0d91facb8c3afe2f88a1358480142.
//
// Solidity: event CollateralDeposited(address indexed trader, uint256 amount)
func (_Market *MarketFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *MarketCollateralDeposited, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CollateralDeposited", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCollateralDeposited)
				if err := _Market.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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
func (_Market *MarketFilterer) ParseCollateralDeposited(log types.Log) (*MarketCollateralDeposited, error) {
	event := new(MarketCollateralDeposited)
	if err := _Market.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCollateralWithdrawnIterator is returned from FilterCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for CollateralWithdrawn events raised by the Market contract.
type MarketCollateralWithdrawnIterator struct {
	Event *MarketCollateralWithdrawn // Event containing the contract specifics and raw log

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
func (it *MarketCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCollateralWithdrawn)
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
		it.Event = new(MarketCollateralWithdrawn)
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
func (it *MarketCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCollateralWithdrawn represents a CollateralWithdrawn event raised by the Market contract.
type MarketCollateralWithdrawn struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralWithdrawn is a free log retrieval operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount)
func (_Market *MarketFilterer) FilterCollateralWithdrawn(opts *bind.FilterOpts, trader []common.Address) (*MarketCollateralWithdrawnIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CollateralWithdrawn", traderRule)
	if err != nil {
		return nil, err
	}
	return &MarketCollateralWithdrawnIterator{contract: _Market.contract, event: "CollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCollateralWithdrawn is a free log subscription operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount)
func (_Market *MarketFilterer) WatchCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketCollateralWithdrawn, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CollateralWithdrawn", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCollateralWithdrawn)
				if err := _Market.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
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

// ParseCollateralWithdrawn is a log parse operation binding the contract event 0xc30fcfbcaac9e0deffa719714eaa82396ff506a0d0d0eebe170830177288715d.
//
// Solidity: event CollateralWithdrawn(address indexed trader, uint256 amount)
func (_Market *MarketFilterer) ParseCollateralWithdrawn(log types.Log) (*MarketCollateralWithdrawn, error) {
	event := new(MarketCollateralWithdrawn)
	if err := _Market.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCycleSettledIterator is returned from FilterCycleSettled and is used to iterate over the raw logs and unpacked data for CycleSettled events raised by the Market contract.
type MarketCycleSettledIterator struct {
	Event *MarketCycleSettled // Event containing the contract specifics and raw log

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
func (it *MarketCycleSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCycleSettled)
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
		it.Event = new(MarketCycleSettled)
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
func (it *MarketCycleSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCycleSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCycleSettled represents a CycleSettled event raised by the Market contract.
type MarketCycleSettled struct {
	CycleId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCycleSettled is a free log retrieval operation binding the contract event 0xfcb4ce7a9eef0b06a841c256c81216fb6035f604416175854e6c9f05ff4fdf8d.
//
// Solidity: event CycleSettled(uint256 indexed cycleId)
func (_Market *MarketFilterer) FilterCycleSettled(opts *bind.FilterOpts, cycleId []*big.Int) (*MarketCycleSettledIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CycleSettled", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketCycleSettledIterator{contract: _Market.contract, event: "CycleSettled", logs: logs, sub: sub}, nil
}

// WatchCycleSettled is a free log subscription operation binding the contract event 0xfcb4ce7a9eef0b06a841c256c81216fb6035f604416175854e6c9f05ff4fdf8d.
//
// Solidity: event CycleSettled(uint256 indexed cycleId)
func (_Market *MarketFilterer) WatchCycleSettled(opts *bind.WatchOpts, sink chan<- *MarketCycleSettled, cycleId []*big.Int) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CycleSettled", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCycleSettled)
				if err := _Market.contract.UnpackLog(event, "CycleSettled", log); err != nil {
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

// ParseCycleSettled is a log parse operation binding the contract event 0xfcb4ce7a9eef0b06a841c256c81216fb6035f604416175854e6c9f05ff4fdf8d.
//
// Solidity: event CycleSettled(uint256 indexed cycleId)
func (_Market *MarketFilterer) ParseCycleSettled(log types.Log) (*MarketCycleSettled, error) {
	event := new(MarketCycleSettled)
	if err := _Market.contract.UnpackLog(event, "CycleSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCycleStartedIterator is returned from FilterCycleStarted and is used to iterate over the raw logs and unpacked data for CycleStarted events raised by the Market contract.
type MarketCycleStartedIterator struct {
	Event *MarketCycleStarted // Event containing the contract specifics and raw log

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
func (it *MarketCycleStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCycleStarted)
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
		it.Event = new(MarketCycleStarted)
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
func (it *MarketCycleStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCycleStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCycleStarted represents a CycleStarted event raised by the Market contract.
type MarketCycleStarted struct {
	CycleId *big.Int
	Strike  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCycleStarted is a free log retrieval operation binding the contract event 0x62eeb403dc233a7676f50cdb0fd72dae595aeb7c63039adce5caacc33254f9c0.
//
// Solidity: event CycleStarted(uint256 indexed cycleId, uint256 strike)
func (_Market *MarketFilterer) FilterCycleStarted(opts *bind.FilterOpts, cycleId []*big.Int) (*MarketCycleStartedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CycleStarted", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketCycleStartedIterator{contract: _Market.contract, event: "CycleStarted", logs: logs, sub: sub}, nil
}

// WatchCycleStarted is a free log subscription operation binding the contract event 0x62eeb403dc233a7676f50cdb0fd72dae595aeb7c63039adce5caacc33254f9c0.
//
// Solidity: event CycleStarted(uint256 indexed cycleId, uint256 strike)
func (_Market *MarketFilterer) WatchCycleStarted(opts *bind.WatchOpts, sink chan<- *MarketCycleStarted, cycleId []*big.Int) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CycleStarted", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCycleStarted)
				if err := _Market.contract.UnpackLog(event, "CycleStarted", log); err != nil {
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

// ParseCycleStarted is a log parse operation binding the contract event 0x62eeb403dc233a7676f50cdb0fd72dae595aeb7c63039adce5caacc33254f9c0.
//
// Solidity: event CycleStarted(uint256 indexed cycleId, uint256 strike)
func (_Market *MarketFilterer) ParseCycleStarted(log types.Log) (*MarketCycleStarted, error) {
	event := new(MarketCycleStarted)
	if err := _Market.contract.UnpackLog(event, "CycleStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Market contract.
type MarketInitializedIterator struct {
	Event *MarketInitialized // Event containing the contract specifics and raw log

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
func (it *MarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketInitialized)
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
		it.Event = new(MarketInitialized)
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
func (it *MarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketInitialized represents a Initialized event raised by the Market contract.
type MarketInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Market *MarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketInitializedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketInitializedIterator{contract: _Market.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Market *MarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketInitialized) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketInitialized)
				if err := _Market.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Market *MarketFilterer) ParseInitialized(log types.Log) (*MarketInitialized, error) {
	event := new(MarketInitialized)
	if err := _Market.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketLiquidatedIterator is returned from FilterLiquidated and is used to iterate over the raw logs and unpacked data for Liquidated events raised by the Market contract.
type MarketLiquidatedIterator struct {
	Event *MarketLiquidated // Event containing the contract specifics and raw log

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
func (it *MarketLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketLiquidated)
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
		it.Event = new(MarketLiquidated)
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
func (it *MarketLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketLiquidated represents a Liquidated event raised by the Market contract.
type MarketLiquidated struct {
	CycleId *big.Int
	Trader  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0x03b910a1639200ab2b0061788a7e4d39f45a305c900ac7f92a77b9a8447df9dc.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader)
func (_Market *MarketFilterer) FilterLiquidated(opts *bind.FilterOpts, cycleId []*big.Int, trader []common.Address) (*MarketLiquidatedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Liquidated", cycleIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &MarketLiquidatedIterator{contract: _Market.contract, event: "Liquidated", logs: logs, sub: sub}, nil
}

// WatchLiquidated is a free log subscription operation binding the contract event 0x03b910a1639200ab2b0061788a7e4d39f45a305c900ac7f92a77b9a8447df9dc.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader)
func (_Market *MarketFilterer) WatchLiquidated(opts *bind.WatchOpts, sink chan<- *MarketLiquidated, cycleId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Liquidated", cycleIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketLiquidated)
				if err := _Market.contract.UnpackLog(event, "Liquidated", log); err != nil {
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

// ParseLiquidated is a log parse operation binding the contract event 0x03b910a1639200ab2b0061788a7e4d39f45a305c900ac7f92a77b9a8447df9dc.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader)
func (_Market *MarketFilterer) ParseLiquidated(log types.Log) (*MarketLiquidated, error) {
	event := new(MarketLiquidated)
	if err := _Market.contract.UnpackLog(event, "Liquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Market contract.
type MarketOrderCancelledIterator struct {
	Event *MarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *MarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOrderCancelled)
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
		it.Event = new(MarketOrderCancelled)
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
func (it *MarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOrderCancelled represents a OrderCancelled event raised by the Market contract.
type MarketOrderCancelled struct {
	CycleId    *big.Int
	OrderId    *big.Int
	Size       *big.Int
	LimitPrice *big.Int
	Maker      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xe382001f9763d7d4e4edcf8785744bcb729a4bb962bb1b83b0f3ba0451145f0b.
//
// Solidity: event OrderCancelled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) FilterOrderCancelled(opts *bind.FilterOpts, cycleId []*big.Int, maker []common.Address) (*MarketOrderCancelledIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OrderCancelled", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOrderCancelledIterator{contract: _Market.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xe382001f9763d7d4e4edcf8785744bcb729a4bb962bb1b83b0f3ba0451145f0b.
//
// Solidity: event OrderCancelled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *MarketOrderCancelled, cycleId []*big.Int, maker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OrderCancelled", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOrderCancelled)
				if err := _Market.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xe382001f9763d7d4e4edcf8785744bcb729a4bb962bb1b83b0f3ba0451145f0b.
//
// Solidity: event OrderCancelled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) ParseOrderCancelled(log types.Log) (*MarketOrderCancelled, error) {
	event := new(MarketOrderCancelled)
	if err := _Market.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Market contract.
type MarketOrderFilledIterator struct {
	Event *MarketOrderFilled // Event containing the contract specifics and raw log

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
func (it *MarketOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOrderFilled)
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
		it.Event = new(MarketOrderFilled)
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
func (it *MarketOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOrderFilled represents a OrderFilled event raised by the Market contract.
type MarketOrderFilled struct {
	CycleId    *big.Int
	OrderId    *big.Int
	Size       *big.Int
	LimitPrice *big.Int
	IsBuy      bool
	IsPut      bool
	Taker      common.Address
	Maker      common.Address
	BtcPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0x102c95915c5aabd72f079785777cd19491895ca2766e057fc571462d5d9d9198.
//
// Solidity: event OrderFilled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed taker, address indexed maker, uint256 btcPrice)
func (_Market *MarketFilterer) FilterOrderFilled(opts *bind.FilterOpts, cycleId []*big.Int, taker []common.Address, maker []common.Address) (*MarketOrderFilledIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OrderFilled", cycleIdRule, takerRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOrderFilledIterator{contract: _Market.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0x102c95915c5aabd72f079785777cd19491895ca2766e057fc571462d5d9d9198.
//
// Solidity: event OrderFilled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed taker, address indexed maker, uint256 btcPrice)
func (_Market *MarketFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *MarketOrderFilled, cycleId []*big.Int, taker []common.Address, maker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OrderFilled", cycleIdRule, takerRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOrderFilled)
				if err := _Market.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0x102c95915c5aabd72f079785777cd19491895ca2766e057fc571462d5d9d9198.
//
// Solidity: event OrderFilled(uint256 indexed cycleId, uint256 orderId, uint128 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed taker, address indexed maker, uint256 btcPrice)
func (_Market *MarketFilterer) ParseOrderFilled(log types.Log) (*MarketOrderFilled, error) {
	event := new(MarketOrderFilled)
	if err := _Market.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the Market contract.
type MarketOrderPlacedIterator struct {
	Event *MarketOrderPlaced // Event containing the contract specifics and raw log

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
func (it *MarketOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOrderPlaced)
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
		it.Event = new(MarketOrderPlaced)
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
func (it *MarketOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOrderPlaced represents a OrderPlaced event raised by the Market contract.
type MarketOrderPlaced struct {
	CycleId    *big.Int
	OrderId    *big.Int
	Size       *big.Int
	LimitPrice *big.Int
	IsBuy      bool
	IsPut      bool
	Maker      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0x1b489ca75ba8d0a2bf1afadb14a7b3f5d645b513987a3affed1e58559ce562d6.
//
// Solidity: event OrderPlaced(uint256 indexed cycleId, uint256 orderId, uint256 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed maker)
func (_Market *MarketFilterer) FilterOrderPlaced(opts *bind.FilterOpts, cycleId []*big.Int, maker []common.Address) (*MarketOrderPlacedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OrderPlaced", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOrderPlacedIterator{contract: _Market.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0x1b489ca75ba8d0a2bf1afadb14a7b3f5d645b513987a3affed1e58559ce562d6.
//
// Solidity: event OrderPlaced(uint256 indexed cycleId, uint256 orderId, uint256 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed maker)
func (_Market *MarketFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *MarketOrderPlaced, cycleId []*big.Int, maker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OrderPlaced", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOrderPlaced)
				if err := _Market.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0x1b489ca75ba8d0a2bf1afadb14a7b3f5d645b513987a3affed1e58559ce562d6.
//
// Solidity: event OrderPlaced(uint256 indexed cycleId, uint256 orderId, uint256 size, uint256 limitPrice, bool isBuy, bool isPut, address indexed maker)
func (_Market *MarketFilterer) ParseOrderPlaced(log types.Log) (*MarketOrderPlaced, error) {
	event := new(MarketOrderPlaced)
	if err := _Market.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Market contract.
type MarketOwnershipTransferredIterator struct {
	Event *MarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOwnershipTransferred)
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
		it.Event = new(MarketOwnershipTransferred)
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
func (it *MarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOwnershipTransferred represents a OwnershipTransferred event raised by the Market contract.
type MarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOwnershipTransferredIterator{contract: _Market.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOwnershipTransferred)
				if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Market *MarketFilterer) ParseOwnershipTransferred(log types.Log) (*MarketOwnershipTransferred, error) {
	event := new(MarketOwnershipTransferred)
	if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Market contract.
type MarketPausedIterator struct {
	Event *MarketPaused // Event containing the contract specifics and raw log

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
func (it *MarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPaused)
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
		it.Event = new(MarketPaused)
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
func (it *MarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPaused represents a Paused event raised by the Market contract.
type MarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketPausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketPausedIterator{contract: _Market.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketPaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPaused)
				if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Market *MarketFilterer) ParsePaused(log types.Log) (*MarketPaused, error) {
	event := new(MarketPaused)
	if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPriceFixedIterator is returned from FilterPriceFixed and is used to iterate over the raw logs and unpacked data for PriceFixed events raised by the Market contract.
type MarketPriceFixedIterator struct {
	Event *MarketPriceFixed // Event containing the contract specifics and raw log

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
func (it *MarketPriceFixedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPriceFixed)
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
		it.Event = new(MarketPriceFixed)
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
func (it *MarketPriceFixedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPriceFixedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPriceFixed represents a PriceFixed event raised by the Market contract.
type MarketPriceFixed struct {
	CycleId *big.Int
	Price   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPriceFixed is a free log retrieval operation binding the contract event 0x54bde8919daf9fd4d1dc5758bff2ebda5cb1e6faf9fae1b9e3d5cf3683be0d35.
//
// Solidity: event PriceFixed(uint256 indexed cycleId, uint64 price)
func (_Market *MarketFilterer) FilterPriceFixed(opts *bind.FilterOpts, cycleId []*big.Int) (*MarketPriceFixedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "PriceFixed", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketPriceFixedIterator{contract: _Market.contract, event: "PriceFixed", logs: logs, sub: sub}, nil
}

// WatchPriceFixed is a free log subscription operation binding the contract event 0x54bde8919daf9fd4d1dc5758bff2ebda5cb1e6faf9fae1b9e3d5cf3683be0d35.
//
// Solidity: event PriceFixed(uint256 indexed cycleId, uint64 price)
func (_Market *MarketFilterer) WatchPriceFixed(opts *bind.WatchOpts, sink chan<- *MarketPriceFixed, cycleId []*big.Int) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "PriceFixed", cycleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPriceFixed)
				if err := _Market.contract.UnpackLog(event, "PriceFixed", log); err != nil {
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

// ParsePriceFixed is a log parse operation binding the contract event 0x54bde8919daf9fd4d1dc5758bff2ebda5cb1e6faf9fae1b9e3d5cf3683be0d35.
//
// Solidity: event PriceFixed(uint256 indexed cycleId, uint64 price)
func (_Market *MarketFilterer) ParsePriceFixed(log types.Log) (*MarketPriceFixed, error) {
	event := new(MarketPriceFixed)
	if err := _Market.contract.UnpackLog(event, "PriceFixed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketSettledIterator is returned from FilterSettled and is used to iterate over the raw logs and unpacked data for Settled events raised by the Market contract.
type MarketSettledIterator struct {
	Event *MarketSettled // Event containing the contract specifics and raw log

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
func (it *MarketSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketSettled)
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
		it.Event = new(MarketSettled)
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
func (it *MarketSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketSettled represents a Settled event raised by the Market contract.
type MarketSettled struct {
	CycleId *big.Int
	Trader  common.Address
	Pnl     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettled is a free log retrieval operation binding the contract event 0x7d603afd81152e037b1b2ed7303fe825424ec4b2600470781ad901d963b740dd.
//
// Solidity: event Settled(uint256 indexed cycleId, address indexed trader, int256 pnl)
func (_Market *MarketFilterer) FilterSettled(opts *bind.FilterOpts, cycleId []*big.Int, trader []common.Address) (*MarketSettledIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Settled", cycleIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &MarketSettledIterator{contract: _Market.contract, event: "Settled", logs: logs, sub: sub}, nil
}

// WatchSettled is a free log subscription operation binding the contract event 0x7d603afd81152e037b1b2ed7303fe825424ec4b2600470781ad901d963b740dd.
//
// Solidity: event Settled(uint256 indexed cycleId, address indexed trader, int256 pnl)
func (_Market *MarketFilterer) WatchSettled(opts *bind.WatchOpts, sink chan<- *MarketSettled, cycleId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Settled", cycleIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketSettled)
				if err := _Market.contract.UnpackLog(event, "Settled", log); err != nil {
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

// ParseSettled is a log parse operation binding the contract event 0x7d603afd81152e037b1b2ed7303fe825424ec4b2600470781ad901d963b740dd.
//
// Solidity: event Settled(uint256 indexed cycleId, address indexed trader, int256 pnl)
func (_Market *MarketFilterer) ParseSettled(log types.Log) (*MarketSettled, error) {
	event := new(MarketSettled)
	if err := _Market.contract.UnpackLog(event, "Settled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Market contract.
type MarketUnpausedIterator struct {
	Event *MarketUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUnpaused)
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
		it.Event = new(MarketUnpaused)
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
func (it *MarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUnpaused represents a Unpaused event raised by the Market contract.
type MarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketUnpausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketUnpausedIterator{contract: _Market.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUnpaused)
				if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Market *MarketFilterer) ParseUnpaused(log types.Log) (*MarketUnpaused, error) {
	event := new(MarketUnpaused)
	if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Market contract.
type MarketUpgradedIterator struct {
	Event *MarketUpgraded // Event containing the contract specifics and raw log

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
func (it *MarketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUpgraded)
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
		it.Event = new(MarketUpgraded)
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
func (it *MarketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUpgraded represents a Upgraded event raised by the Market contract.
type MarketUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Market *MarketFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketUpgradedIterator{contract: _Market.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Market *MarketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUpgraded)
				if err := _Market.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Market *MarketFilterer) ParseUpgraded(log types.Log) (*MarketUpgraded, error) {
	event := new(MarketUpgraded)
	if err := _Market.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
