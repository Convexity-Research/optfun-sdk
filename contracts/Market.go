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

// Cycle is an auto generated low-level Go binding around an user-defined struct.
type Cycle struct {
	IsSettled       bool
	StrikePrice     uint64
	SettlementPrice uint64
}

// Level is an auto generated low-level Go binding around an user-defined struct.
type Level struct {
	Vol  *big.Int
	Head uint32
	Tail uint32
}

// SharedStorageUserAccount is an auto generated low-level Go binding around an user-defined struct.
type SharedStorageUserAccount struct {
	ActiveInCycle      bool
	LiquidationQueued  bool
	Balance            uint64
	LiquidationFeeOwed uint64
	ScratchPnL         uint64
	Gap                *big.Int
	LongCalls          uint32
	ShortCalls         uint32
	LongPuts           uint32
	ShortPuts          uint32
	PendingLongCalls   uint32
	PendingShortCalls  uint32
	PendingLongPuts    uint32
	PendingShortPuts   uint32
}

// TakerQ is an auto generated low-level Go binding around an user-defined struct.
type TakerQ struct {
	Trader       common.Address
	Size         uint64
	TakerOrderId int32
}

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelAndClose\",\"inputs\":[{\"name\":\"buyCallPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sellCallPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyPutPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sellPutPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveCycle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCycles\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCycle\",\"components\":[{\"name\":\"isSettled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"strikePrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"settlementPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLevels\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structLevel\",\"components\":[{\"name\":\"vol\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"head\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tail\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMmBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumTraders\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTakerQ\",\"inputs\":[{\"name\":\"side\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structTakerQ[]\",\"components\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"takerOrderId\",\"type\":\"int32\",\"internalType\":\"int32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserAccounts\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSharedStorage.UserAccount\",\"components\":[{\"name\":\"activeInCycle\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"liquidationQueued\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"balance\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liquidationFeeOwed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"scratchPnL\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_gap\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"longCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"shortCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"longPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"shortPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingLongCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingShortCalls\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingLongPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pendingShortPuts\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserOrders\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWhitelist\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isLiquidatable\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidatable\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"long\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPriceBuy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPriceSell\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cycleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeMultiOrder\",\"inputs\":[{\"name\":\"sides\",\"type\":\"uint8[]\",\"internalType\":\"enumMarketSide[]\"},{\"name\":\"sizes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"limitPrices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"cycleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"placeOrder\",\"inputs\":[{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumMarketSide\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cycleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExtension\",\"inputs\":[{\"name\":\"ext\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleChunk\",\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pauseNextCycle\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"short\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPriceBuy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitPriceSell\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cycleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startCycle\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralWithdrawn\",\"inputs\":[{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CycleSettled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CycleStarted\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"strike\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LimitOrderCancelled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"makerOrderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LimitOrderFilled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"makerOrderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"takerOrderId\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketSide\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cashTaker\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"cashMaker\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"btcPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LimitOrderPlaced\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"makerOrderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"limitPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketSide\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Liquidated\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"liqFeeOwed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFixed\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settled\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"trader\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TakerOrderPlaced\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"takerOrderId\",\"type\":\"int32\",\"indexed\":false,\"internalType\":\"int32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketSide\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TakerOrderRemaining\",\"inputs\":[{\"name\":\"cycleId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"takerOrderId\",\"type\":\"int32\",\"indexed\":false,\"internalType\":\"int32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMarketSide\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountInLiquidation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CycleActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CycleAlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CycleNotStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InTraderList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCycle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWhitelistSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketNotLive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOrderOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSecurityCouncil\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OraclePriceCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StillSolvent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TraderLiquidatable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
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

// GetActiveCycle is a free data retrieval call binding the contract method 0xce294cf8.
//
// Solidity: function getActiveCycle() view returns(uint256)
func (_Market *MarketCaller) GetActiveCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getActiveCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveCycle is a free data retrieval call binding the contract method 0xce294cf8.
//
// Solidity: function getActiveCycle() view returns(uint256)
func (_Market *MarketSession) GetActiveCycle() (*big.Int, error) {
	return _Market.Contract.GetActiveCycle(&_Market.CallOpts)
}

// GetActiveCycle is a free data retrieval call binding the contract method 0xce294cf8.
//
// Solidity: function getActiveCycle() view returns(uint256)
func (_Market *MarketCallerSession) GetActiveCycle() (*big.Int, error) {
	return _Market.Contract.GetActiveCycle(&_Market.CallOpts)
}

// GetCollateralToken is a free data retrieval call binding the contract method 0x96423663.
//
// Solidity: function getCollateralToken() view returns(address)
func (_Market *MarketCaller) GetCollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getCollateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateralToken is a free data retrieval call binding the contract method 0x96423663.
//
// Solidity: function getCollateralToken() view returns(address)
func (_Market *MarketSession) GetCollateralToken() (common.Address, error) {
	return _Market.Contract.GetCollateralToken(&_Market.CallOpts)
}

// GetCollateralToken is a free data retrieval call binding the contract method 0x96423663.
//
// Solidity: function getCollateralToken() view returns(address)
func (_Market *MarketCallerSession) GetCollateralToken() (common.Address, error) {
	return _Market.Contract.GetCollateralToken(&_Market.CallOpts)
}

// GetCycles is a free data retrieval call binding the contract method 0x774d922b.
//
// Solidity: function getCycles(uint256 cycleId) view returns((bool,uint64,uint64))
func (_Market *MarketCaller) GetCycles(opts *bind.CallOpts, cycleId *big.Int) (Cycle, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getCycles", cycleId)

	if err != nil {
		return *new(Cycle), err
	}

	out0 := *abi.ConvertType(out[0], new(Cycle)).(*Cycle)

	return out0, err

}

// GetCycles is a free data retrieval call binding the contract method 0x774d922b.
//
// Solidity: function getCycles(uint256 cycleId) view returns((bool,uint64,uint64))
func (_Market *MarketSession) GetCycles(cycleId *big.Int) (Cycle, error) {
	return _Market.Contract.GetCycles(&_Market.CallOpts, cycleId)
}

// GetCycles is a free data retrieval call binding the contract method 0x774d922b.
//
// Solidity: function getCycles(uint256 cycleId) view returns((bool,uint64,uint64))
func (_Market *MarketCallerSession) GetCycles(cycleId *big.Int) (Cycle, error) {
	return _Market.Contract.GetCycles(&_Market.CallOpts, cycleId)
}

// GetLevels is a free data retrieval call binding the contract method 0xd084ba7e.
//
// Solidity: function getLevels(uint32 key) view returns((uint128,uint32,uint32))
func (_Market *MarketCaller) GetLevels(opts *bind.CallOpts, key uint32) (Level, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getLevels", key)

	if err != nil {
		return *new(Level), err
	}

	out0 := *abi.ConvertType(out[0], new(Level)).(*Level)

	return out0, err

}

// GetLevels is a free data retrieval call binding the contract method 0xd084ba7e.
//
// Solidity: function getLevels(uint32 key) view returns((uint128,uint32,uint32))
func (_Market *MarketSession) GetLevels(key uint32) (Level, error) {
	return _Market.Contract.GetLevels(&_Market.CallOpts, key)
}

// GetLevels is a free data retrieval call binding the contract method 0xd084ba7e.
//
// Solidity: function getLevels(uint32 key) view returns((uint128,uint32,uint32))
func (_Market *MarketCallerSession) GetLevels(key uint32) (Level, error) {
	return _Market.Contract.GetLevels(&_Market.CallOpts, key)
}

// GetMmBps is a free data retrieval call binding the contract method 0xd0b7a062.
//
// Solidity: function getMmBps() pure returns(uint256)
func (_Market *MarketCaller) GetMmBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getMmBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMmBps is a free data retrieval call binding the contract method 0xd0b7a062.
//
// Solidity: function getMmBps() pure returns(uint256)
func (_Market *MarketSession) GetMmBps() (*big.Int, error) {
	return _Market.Contract.GetMmBps(&_Market.CallOpts)
}

// GetMmBps is a free data retrieval call binding the contract method 0xd0b7a062.
//
// Solidity: function getMmBps() pure returns(uint256)
func (_Market *MarketCallerSession) GetMmBps() (*big.Int, error) {
	return _Market.Contract.GetMmBps(&_Market.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Market *MarketCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Market *MarketSession) GetName() (string, error) {
	return _Market.Contract.GetName(&_Market.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Market *MarketCallerSession) GetName() (string, error) {
	return _Market.Contract.GetName(&_Market.CallOpts)
}

// GetNumTraders is a free data retrieval call binding the contract method 0xe6707e16.
//
// Solidity: function getNumTraders() view returns(uint256)
func (_Market *MarketCaller) GetNumTraders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getNumTraders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTraders is a free data retrieval call binding the contract method 0xe6707e16.
//
// Solidity: function getNumTraders() view returns(uint256)
func (_Market *MarketSession) GetNumTraders() (*big.Int, error) {
	return _Market.Contract.GetNumTraders(&_Market.CallOpts)
}

// GetNumTraders is a free data retrieval call binding the contract method 0xe6707e16.
//
// Solidity: function getNumTraders() view returns(uint256)
func (_Market *MarketCallerSession) GetNumTraders() (*big.Int, error) {
	return _Market.Contract.GetNumTraders(&_Market.CallOpts)
}

// GetTakerQ is a free data retrieval call binding the contract method 0x1b5fe22d.
//
// Solidity: function getTakerQ(uint256 side) view returns((address,uint64,int32)[])
func (_Market *MarketCaller) GetTakerQ(opts *bind.CallOpts, side *big.Int) ([]TakerQ, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getTakerQ", side)

	if err != nil {
		return *new([]TakerQ), err
	}

	out0 := *abi.ConvertType(out[0], new([]TakerQ)).(*[]TakerQ)

	return out0, err

}

// GetTakerQ is a free data retrieval call binding the contract method 0x1b5fe22d.
//
// Solidity: function getTakerQ(uint256 side) view returns((address,uint64,int32)[])
func (_Market *MarketSession) GetTakerQ(side *big.Int) ([]TakerQ, error) {
	return _Market.Contract.GetTakerQ(&_Market.CallOpts, side)
}

// GetTakerQ is a free data retrieval call binding the contract method 0x1b5fe22d.
//
// Solidity: function getTakerQ(uint256 side) view returns((address,uint64,int32)[])
func (_Market *MarketCallerSession) GetTakerQ(side *big.Int) ([]TakerQ, error) {
	return _Market.Contract.GetTakerQ(&_Market.CallOpts, side)
}

// GetUserAccounts is a free data retrieval call binding the contract method 0xad872a86.
//
// Solidity: function getUserAccounts(address trader) view returns((bool,bool,uint64,uint64,uint64,uint48,uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint32))
func (_Market *MarketCaller) GetUserAccounts(opts *bind.CallOpts, trader common.Address) (SharedStorageUserAccount, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getUserAccounts", trader)

	if err != nil {
		return *new(SharedStorageUserAccount), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStorageUserAccount)).(*SharedStorageUserAccount)

	return out0, err

}

// GetUserAccounts is a free data retrieval call binding the contract method 0xad872a86.
//
// Solidity: function getUserAccounts(address trader) view returns((bool,bool,uint64,uint64,uint64,uint48,uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint32))
func (_Market *MarketSession) GetUserAccounts(trader common.Address) (SharedStorageUserAccount, error) {
	return _Market.Contract.GetUserAccounts(&_Market.CallOpts, trader)
}

// GetUserAccounts is a free data retrieval call binding the contract method 0xad872a86.
//
// Solidity: function getUserAccounts(address trader) view returns((bool,bool,uint64,uint64,uint64,uint48,uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint32))
func (_Market *MarketCallerSession) GetUserAccounts(trader common.Address) (SharedStorageUserAccount, error) {
	return _Market.Contract.GetUserAccounts(&_Market.CallOpts, trader)
}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address trader) view returns(uint32[])
func (_Market *MarketCaller) GetUserOrders(opts *bind.CallOpts, trader common.Address) ([]uint32, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getUserOrders", trader)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address trader) view returns(uint32[])
func (_Market *MarketSession) GetUserOrders(trader common.Address) ([]uint32, error) {
	return _Market.Contract.GetUserOrders(&_Market.CallOpts, trader)
}

// GetUserOrders is a free data retrieval call binding the contract method 0x63c69f08.
//
// Solidity: function getUserOrders(address trader) view returns(uint32[])
func (_Market *MarketCallerSession) GetUserOrders(trader common.Address) ([]uint32, error) {
	return _Market.Contract.GetUserOrders(&_Market.CallOpts, trader)
}

// GetWhitelist is a free data retrieval call binding the contract method 0x30edc0f5.
//
// Solidity: function getWhitelist(address account) view returns(bool)
func (_Market *MarketCaller) GetWhitelist(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getWhitelist", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWhitelist is a free data retrieval call binding the contract method 0x30edc0f5.
//
// Solidity: function getWhitelist(address account) view returns(bool)
func (_Market *MarketSession) GetWhitelist(account common.Address) (bool, error) {
	return _Market.Contract.GetWhitelist(&_Market.CallOpts, account)
}

// GetWhitelist is a free data retrieval call binding the contract method 0x30edc0f5.
//
// Solidity: function getWhitelist(address account) view returns(bool)
func (_Market *MarketCallerSession) GetWhitelist(account common.Address) (bool, error) {
	return _Market.Contract.GetWhitelist(&_Market.CallOpts, account)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address trader) view returns(bool)
func (_Market *MarketCaller) IsLiquidatable(opts *bind.CallOpts, trader common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isLiquidatable", trader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address trader) view returns(bool)
func (_Market *MarketSession) IsLiquidatable(trader common.Address) (bool, error) {
	return _Market.Contract.IsLiquidatable(&_Market.CallOpts, trader)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address trader) view returns(bool)
func (_Market *MarketCallerSession) IsLiquidatable(trader common.Address) (bool, error) {
	return _Market.Contract.IsLiquidatable(&_Market.CallOpts, trader)
}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0xab68e733.
//
// Solidity: function isLiquidatable(address trader, uint64 price) view returns(bool)
func (_Market *MarketCaller) IsLiquidatable0(opts *bind.CallOpts, trader common.Address, price uint64) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isLiquidatable0", trader, price)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0xab68e733.
//
// Solidity: function isLiquidatable(address trader, uint64 price) view returns(bool)
func (_Market *MarketSession) IsLiquidatable0(trader common.Address, price uint64) (bool, error) {
	return _Market.Contract.IsLiquidatable0(&_Market.CallOpts, trader, price)
}

// IsLiquidatable0 is a free data retrieval call binding the contract method 0xab68e733.
//
// Solidity: function isLiquidatable(address trader, uint64 price) view returns(bool)
func (_Market *MarketCallerSession) IsLiquidatable0(trader common.Address, price uint64) (bool, error) {
	return _Market.Contract.IsLiquidatable0(&_Market.CallOpts, trader, price)
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

// CancelAndClose is a paid mutator transaction binding the contract method 0x8e17eec6.
//
// Solidity: function cancelAndClose(uint256 buyCallPrice, uint256 sellCallPrice, uint256 buyPutPrice, uint256 sellPutPrice) returns()
func (_Market *MarketTransactor) CancelAndClose(opts *bind.TransactOpts, buyCallPrice *big.Int, sellCallPrice *big.Int, buyPutPrice *big.Int, sellPutPrice *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "cancelAndClose", buyCallPrice, sellCallPrice, buyPutPrice, sellPutPrice)
}

// CancelAndClose is a paid mutator transaction binding the contract method 0x8e17eec6.
//
// Solidity: function cancelAndClose(uint256 buyCallPrice, uint256 sellCallPrice, uint256 buyPutPrice, uint256 sellPutPrice) returns()
func (_Market *MarketSession) CancelAndClose(buyCallPrice *big.Int, sellCallPrice *big.Int, buyPutPrice *big.Int, sellPutPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelAndClose(&_Market.TransactOpts, buyCallPrice, sellCallPrice, buyPutPrice, sellPutPrice)
}

// CancelAndClose is a paid mutator transaction binding the contract method 0x8e17eec6.
//
// Solidity: function cancelAndClose(uint256 buyCallPrice, uint256 sellCallPrice, uint256 buyPutPrice, uint256 sellPutPrice) returns()
func (_Market *MarketTransactorSession) CancelAndClose(buyCallPrice *big.Int, sellCallPrice *big.Int, buyPutPrice *big.Int, sellPutPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.CancelAndClose(&_Market.TransactOpts, buyCallPrice, sellCallPrice, buyPutPrice, sellPutPrice)
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

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xf5207793.
//
// Solidity: function depositCollateral(uint256 amount, bytes signature) returns()
func (_Market *MarketTransactor) DepositCollateral0(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "depositCollateral0", amount, signature)
}

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xf5207793.
//
// Solidity: function depositCollateral(uint256 amount, bytes signature) returns()
func (_Market *MarketSession) DepositCollateral0(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.DepositCollateral0(&_Market.TransactOpts, amount, signature)
}

// DepositCollateral0 is a paid mutator transaction binding the contract method 0xf5207793.
//
// Solidity: function depositCollateral(uint256 amount, bytes signature) returns()
func (_Market *MarketTransactorSession) DepositCollateral0(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Market.Contract.DepositCollateral0(&_Market.TransactOpts, amount, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b885ac3.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder, address _governance) returns()
func (_Market *MarketTransactor) Initialize(opts *bind.TransactOpts, _name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "initialize", _name, _feeRecipient, _collateralToken, _forwarder, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b885ac3.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder, address _governance) returns()
func (_Market *MarketSession) Initialize(_name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _name, _feeRecipient, _collateralToken, _forwarder, _governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x0b885ac3.
//
// Solidity: function initialize(string _name, address _feeRecipient, address _collateralToken, address _forwarder, address _governance) returns()
func (_Market *MarketTransactorSession) Initialize(_name string, _feeRecipient common.Address, _collateralToken common.Address, _forwarder common.Address, _governance common.Address) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _name, _feeRecipient, _collateralToken, _forwarder, _governance)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address trader) returns()
func (_Market *MarketTransactor) Liquidate(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "liquidate", trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address trader) returns()
func (_Market *MarketSession) Liquidate(trader common.Address) (*types.Transaction, error) {
	return _Market.Contract.Liquidate(&_Market.TransactOpts, trader)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address trader) returns()
func (_Market *MarketTransactorSession) Liquidate(trader common.Address) (*types.Transaction, error) {
	return _Market.Contract.Liquidate(&_Market.TransactOpts, trader)
}

// Long is a paid mutator transaction binding the contract method 0x7d3525b2.
//
// Solidity: function long(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketTransactor) Long(opts *bind.TransactOpts, size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "long", size, limitPriceBuy, limitPriceSell, cycleId)
}

// Long is a paid mutator transaction binding the contract method 0x7d3525b2.
//
// Solidity: function long(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketSession) Long(size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Long(&_Market.TransactOpts, size, limitPriceBuy, limitPriceSell, cycleId)
}

// Long is a paid mutator transaction binding the contract method 0x7d3525b2.
//
// Solidity: function long(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketTransactorSession) Long(size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Long(&_Market.TransactOpts, size, limitPriceBuy, limitPriceSell, cycleId)
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

// PlaceMultiOrder is a paid mutator transaction binding the contract method 0x1fd3eeda.
//
// Solidity: function placeMultiOrder(uint8[] sides, uint256[] sizes, uint256[] limitPrices, uint256 cycleId) returns()
func (_Market *MarketTransactor) PlaceMultiOrder(opts *bind.TransactOpts, sides []uint8, sizes []*big.Int, limitPrices []*big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "placeMultiOrder", sides, sizes, limitPrices, cycleId)
}

// PlaceMultiOrder is a paid mutator transaction binding the contract method 0x1fd3eeda.
//
// Solidity: function placeMultiOrder(uint8[] sides, uint256[] sizes, uint256[] limitPrices, uint256 cycleId) returns()
func (_Market *MarketSession) PlaceMultiOrder(sides []uint8, sizes []*big.Int, limitPrices []*big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceMultiOrder(&_Market.TransactOpts, sides, sizes, limitPrices, cycleId)
}

// PlaceMultiOrder is a paid mutator transaction binding the contract method 0x1fd3eeda.
//
// Solidity: function placeMultiOrder(uint8[] sides, uint256[] sizes, uint256[] limitPrices, uint256 cycleId) returns()
func (_Market *MarketTransactorSession) PlaceMultiOrder(sides []uint8, sizes []*big.Int, limitPrices []*big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceMultiOrder(&_Market.TransactOpts, sides, sizes, limitPrices, cycleId)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4622c218.
//
// Solidity: function placeOrder(uint8 side, uint256 size, uint256 limitPrice, uint256 cycleId) returns(uint256 orderId)
func (_Market *MarketTransactor) PlaceOrder(opts *bind.TransactOpts, side uint8, size *big.Int, limitPrice *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "placeOrder", side, size, limitPrice, cycleId)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4622c218.
//
// Solidity: function placeOrder(uint8 side, uint256 size, uint256 limitPrice, uint256 cycleId) returns(uint256 orderId)
func (_Market *MarketSession) PlaceOrder(side uint8, size *big.Int, limitPrice *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder(&_Market.TransactOpts, side, size, limitPrice, cycleId)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x4622c218.
//
// Solidity: function placeOrder(uint8 side, uint256 size, uint256 limitPrice, uint256 cycleId) returns(uint256 orderId)
func (_Market *MarketTransactorSession) PlaceOrder(side uint8, size *big.Int, limitPrice *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PlaceOrder(&_Market.TransactOpts, side, size, limitPrice, cycleId)
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

// SetExtension is a paid mutator transaction binding the contract method 0xab622d9a.
//
// Solidity: function setExtension(address ext) returns()
func (_Market *MarketTransactor) SetExtension(opts *bind.TransactOpts, ext common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setExtension", ext)
}

// SetExtension is a paid mutator transaction binding the contract method 0xab622d9a.
//
// Solidity: function setExtension(address ext) returns()
func (_Market *MarketSession) SetExtension(ext common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetExtension(&_Market.TransactOpts, ext)
}

// SetExtension is a paid mutator transaction binding the contract method 0xab622d9a.
//
// Solidity: function setExtension(address ext) returns()
func (_Market *MarketTransactorSession) SetExtension(ext common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetExtension(&_Market.TransactOpts, ext)
}

// SettleChunk is a paid mutator transaction binding the contract method 0xde2bb589.
//
// Solidity: function settleChunk(uint256 max, bool pauseNextCycle) returns()
func (_Market *MarketTransactor) SettleChunk(opts *bind.TransactOpts, max *big.Int, pauseNextCycle bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "settleChunk", max, pauseNextCycle)
}

// SettleChunk is a paid mutator transaction binding the contract method 0xde2bb589.
//
// Solidity: function settleChunk(uint256 max, bool pauseNextCycle) returns()
func (_Market *MarketSession) SettleChunk(max *big.Int, pauseNextCycle bool) (*types.Transaction, error) {
	return _Market.Contract.SettleChunk(&_Market.TransactOpts, max, pauseNextCycle)
}

// SettleChunk is a paid mutator transaction binding the contract method 0xde2bb589.
//
// Solidity: function settleChunk(uint256 max, bool pauseNextCycle) returns()
func (_Market *MarketTransactorSession) SettleChunk(max *big.Int, pauseNextCycle bool) (*types.Transaction, error) {
	return _Market.Contract.SettleChunk(&_Market.TransactOpts, max, pauseNextCycle)
}

// Short is a paid mutator transaction binding the contract method 0x83850047.
//
// Solidity: function short(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketTransactor) Short(opts *bind.TransactOpts, size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "short", size, limitPriceBuy, limitPriceSell, cycleId)
}

// Short is a paid mutator transaction binding the contract method 0x83850047.
//
// Solidity: function short(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketSession) Short(size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Short(&_Market.TransactOpts, size, limitPriceBuy, limitPriceSell, cycleId)
}

// Short is a paid mutator transaction binding the contract method 0x83850047.
//
// Solidity: function short(uint256 size, uint256 limitPriceBuy, uint256 limitPriceSell, uint256 cycleId) returns()
func (_Market *MarketTransactorSession) Short(size *big.Int, limitPriceBuy *big.Int, limitPriceSell *big.Int, cycleId *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Short(&_Market.TransactOpts, size, limitPriceBuy, limitPriceSell, cycleId)
}

// StartCycle is a paid mutator transaction binding the contract method 0x562cad23.
//
// Solidity: function startCycle() returns()
func (_Market *MarketTransactor) StartCycle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "startCycle")
}

// StartCycle is a paid mutator transaction binding the contract method 0x562cad23.
//
// Solidity: function startCycle() returns()
func (_Market *MarketSession) StartCycle() (*types.Transaction, error) {
	return _Market.Contract.StartCycle(&_Market.TransactOpts)
}

// StartCycle is a paid mutator transaction binding the contract method 0x562cad23.
//
// Solidity: function startCycle() returns()
func (_Market *MarketTransactorSession) StartCycle() (*types.Transaction, error) {
	return _Market.Contract.StartCycle(&_Market.TransactOpts)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Market *MarketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Market.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Market *MarketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Market.Contract.Fallback(&_Market.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Market *MarketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Market.Contract.Fallback(&_Market.TransactOpts, calldata)
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

// MarketLimitOrderCancelledIterator is returned from FilterLimitOrderCancelled and is used to iterate over the raw logs and unpacked data for LimitOrderCancelled events raised by the Market contract.
type MarketLimitOrderCancelledIterator struct {
	Event *MarketLimitOrderCancelled // Event containing the contract specifics and raw log

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
func (it *MarketLimitOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketLimitOrderCancelled)
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
		it.Event = new(MarketLimitOrderCancelled)
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
func (it *MarketLimitOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketLimitOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketLimitOrderCancelled represents a LimitOrderCancelled event raised by the Market contract.
type MarketLimitOrderCancelled struct {
	CycleId      *big.Int
	MakerOrderId *big.Int
	Size         *big.Int
	LimitPrice   *big.Int
	Maker        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLimitOrderCancelled is a free log retrieval operation binding the contract event 0x1dd5cf68bf69ecb3e22036f61b054535bd40347e9710a236ed3c6d5a99ec371a.
//
// Solidity: event LimitOrderCancelled(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) FilterLimitOrderCancelled(opts *bind.FilterOpts, cycleId []*big.Int, maker []common.Address) (*MarketLimitOrderCancelledIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "LimitOrderCancelled", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketLimitOrderCancelledIterator{contract: _Market.contract, event: "LimitOrderCancelled", logs: logs, sub: sub}, nil
}

// WatchLimitOrderCancelled is a free log subscription operation binding the contract event 0x1dd5cf68bf69ecb3e22036f61b054535bd40347e9710a236ed3c6d5a99ec371a.
//
// Solidity: event LimitOrderCancelled(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) WatchLimitOrderCancelled(opts *bind.WatchOpts, sink chan<- *MarketLimitOrderCancelled, cycleId []*big.Int, maker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "LimitOrderCancelled", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketLimitOrderCancelled)
				if err := _Market.contract.UnpackLog(event, "LimitOrderCancelled", log); err != nil {
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

// ParseLimitOrderCancelled is a log parse operation binding the contract event 0x1dd5cf68bf69ecb3e22036f61b054535bd40347e9710a236ed3c6d5a99ec371a.
//
// Solidity: event LimitOrderCancelled(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, address indexed maker)
func (_Market *MarketFilterer) ParseLimitOrderCancelled(log types.Log) (*MarketLimitOrderCancelled, error) {
	event := new(MarketLimitOrderCancelled)
	if err := _Market.contract.UnpackLog(event, "LimitOrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketLimitOrderFilledIterator is returned from FilterLimitOrderFilled and is used to iterate over the raw logs and unpacked data for LimitOrderFilled events raised by the Market contract.
type MarketLimitOrderFilledIterator struct {
	Event *MarketLimitOrderFilled // Event containing the contract specifics and raw log

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
func (it *MarketLimitOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketLimitOrderFilled)
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
		it.Event = new(MarketLimitOrderFilled)
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
func (it *MarketLimitOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketLimitOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketLimitOrderFilled represents a LimitOrderFilled event raised by the Market contract.
type MarketLimitOrderFilled struct {
	CycleId      *big.Int
	MakerOrderId *big.Int
	TakerOrderId *big.Int
	Size         *big.Int
	LimitPrice   *big.Int
	Side         uint8
	Taker        common.Address
	Maker        common.Address
	CashTaker    *big.Int
	CashMaker    *big.Int
	BtcPrice     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLimitOrderFilled is a free log retrieval operation binding the contract event 0x49dac6294dcbe9dfa9fe99064a7f2702118e9ec00e8a067c69b2b9ce2d08eccd.
//
// Solidity: event LimitOrderFilled(uint256 indexed cycleId, uint256 makerOrderId, int256 takerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed taker, address indexed maker, int256 cashTaker, int256 cashMaker, uint256 btcPrice)
func (_Market *MarketFilterer) FilterLimitOrderFilled(opts *bind.FilterOpts, cycleId []*big.Int, taker []common.Address, maker []common.Address) (*MarketLimitOrderFilledIterator, error) {

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

	logs, sub, err := _Market.contract.FilterLogs(opts, "LimitOrderFilled", cycleIdRule, takerRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketLimitOrderFilledIterator{contract: _Market.contract, event: "LimitOrderFilled", logs: logs, sub: sub}, nil
}

// WatchLimitOrderFilled is a free log subscription operation binding the contract event 0x49dac6294dcbe9dfa9fe99064a7f2702118e9ec00e8a067c69b2b9ce2d08eccd.
//
// Solidity: event LimitOrderFilled(uint256 indexed cycleId, uint256 makerOrderId, int256 takerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed taker, address indexed maker, int256 cashTaker, int256 cashMaker, uint256 btcPrice)
func (_Market *MarketFilterer) WatchLimitOrderFilled(opts *bind.WatchOpts, sink chan<- *MarketLimitOrderFilled, cycleId []*big.Int, taker []common.Address, maker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Market.contract.WatchLogs(opts, "LimitOrderFilled", cycleIdRule, takerRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketLimitOrderFilled)
				if err := _Market.contract.UnpackLog(event, "LimitOrderFilled", log); err != nil {
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

// ParseLimitOrderFilled is a log parse operation binding the contract event 0x49dac6294dcbe9dfa9fe99064a7f2702118e9ec00e8a067c69b2b9ce2d08eccd.
//
// Solidity: event LimitOrderFilled(uint256 indexed cycleId, uint256 makerOrderId, int256 takerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed taker, address indexed maker, int256 cashTaker, int256 cashMaker, uint256 btcPrice)
func (_Market *MarketFilterer) ParseLimitOrderFilled(log types.Log) (*MarketLimitOrderFilled, error) {
	event := new(MarketLimitOrderFilled)
	if err := _Market.contract.UnpackLog(event, "LimitOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketLimitOrderPlacedIterator is returned from FilterLimitOrderPlaced and is used to iterate over the raw logs and unpacked data for LimitOrderPlaced events raised by the Market contract.
type MarketLimitOrderPlacedIterator struct {
	Event *MarketLimitOrderPlaced // Event containing the contract specifics and raw log

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
func (it *MarketLimitOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketLimitOrderPlaced)
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
		it.Event = new(MarketLimitOrderPlaced)
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
func (it *MarketLimitOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketLimitOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketLimitOrderPlaced represents a LimitOrderPlaced event raised by the Market contract.
type MarketLimitOrderPlaced struct {
	CycleId      *big.Int
	MakerOrderId *big.Int
	Size         *big.Int
	LimitPrice   *big.Int
	Side         uint8
	Maker        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLimitOrderPlaced is a free log retrieval operation binding the contract event 0x0be99a1df32ac89abadc9bae2496b27c104dae5ae43969be6052ef280a0ed2c6.
//
// Solidity: event LimitOrderPlaced(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed maker)
func (_Market *MarketFilterer) FilterLimitOrderPlaced(opts *bind.FilterOpts, cycleId []*big.Int, maker []common.Address) (*MarketLimitOrderPlacedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "LimitOrderPlaced", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &MarketLimitOrderPlacedIterator{contract: _Market.contract, event: "LimitOrderPlaced", logs: logs, sub: sub}, nil
}

// WatchLimitOrderPlaced is a free log subscription operation binding the contract event 0x0be99a1df32ac89abadc9bae2496b27c104dae5ae43969be6052ef280a0ed2c6.
//
// Solidity: event LimitOrderPlaced(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed maker)
func (_Market *MarketFilterer) WatchLimitOrderPlaced(opts *bind.WatchOpts, sink chan<- *MarketLimitOrderPlaced, cycleId []*big.Int, maker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "LimitOrderPlaced", cycleIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketLimitOrderPlaced)
				if err := _Market.contract.UnpackLog(event, "LimitOrderPlaced", log); err != nil {
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

// ParseLimitOrderPlaced is a log parse operation binding the contract event 0x0be99a1df32ac89abadc9bae2496b27c104dae5ae43969be6052ef280a0ed2c6.
//
// Solidity: event LimitOrderPlaced(uint256 indexed cycleId, uint256 makerOrderId, uint256 size, uint256 limitPrice, uint8 side, address indexed maker)
func (_Market *MarketFilterer) ParseLimitOrderPlaced(log types.Log) (*MarketLimitOrderPlaced, error) {
	event := new(MarketLimitOrderPlaced)
	if err := _Market.contract.UnpackLog(event, "LimitOrderPlaced", log); err != nil {
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
	CycleId    *big.Int
	Trader     common.Address
	LiqFeeOwed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0xf4c6cbfcc96248be8ecbaf76de0fee34f71f2fadd9af537dd38c2657621930d6.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader, uint256 liqFeeOwed)
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

// WatchLiquidated is a free log subscription operation binding the contract event 0xf4c6cbfcc96248be8ecbaf76de0fee34f71f2fadd9af537dd38c2657621930d6.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader, uint256 liqFeeOwed)
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

// ParseLiquidated is a log parse operation binding the contract event 0xf4c6cbfcc96248be8ecbaf76de0fee34f71f2fadd9af537dd38c2657621930d6.
//
// Solidity: event Liquidated(uint256 indexed cycleId, address indexed trader, uint256 liqFeeOwed)
func (_Market *MarketFilterer) ParseLiquidated(log types.Log) (*MarketLiquidated, error) {
	event := new(MarketLiquidated)
	if err := _Market.contract.UnpackLog(event, "Liquidated", log); err != nil {
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

// MarketTakerOrderPlacedIterator is returned from FilterTakerOrderPlaced and is used to iterate over the raw logs and unpacked data for TakerOrderPlaced events raised by the Market contract.
type MarketTakerOrderPlacedIterator struct {
	Event *MarketTakerOrderPlaced // Event containing the contract specifics and raw log

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
func (it *MarketTakerOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTakerOrderPlaced)
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
		it.Event = new(MarketTakerOrderPlaced)
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
func (it *MarketTakerOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTakerOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTakerOrderPlaced represents a TakerOrderPlaced event raised by the Market contract.
type MarketTakerOrderPlaced struct {
	CycleId      *big.Int
	TakerOrderId int32
	Size         *big.Int
	Side         uint8
	Taker        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTakerOrderPlaced is a free log retrieval operation binding the contract event 0xdd67da5d2d18e4e450567dcad68b213da4e86181939bb6c057b280d8c0810d31.
//
// Solidity: event TakerOrderPlaced(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) FilterTakerOrderPlaced(opts *bind.FilterOpts, cycleId []*big.Int, taker []common.Address) (*MarketTakerOrderPlacedIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "TakerOrderPlaced", cycleIdRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &MarketTakerOrderPlacedIterator{contract: _Market.contract, event: "TakerOrderPlaced", logs: logs, sub: sub}, nil
}

// WatchTakerOrderPlaced is a free log subscription operation binding the contract event 0xdd67da5d2d18e4e450567dcad68b213da4e86181939bb6c057b280d8c0810d31.
//
// Solidity: event TakerOrderPlaced(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) WatchTakerOrderPlaced(opts *bind.WatchOpts, sink chan<- *MarketTakerOrderPlaced, cycleId []*big.Int, taker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "TakerOrderPlaced", cycleIdRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTakerOrderPlaced)
				if err := _Market.contract.UnpackLog(event, "TakerOrderPlaced", log); err != nil {
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

// ParseTakerOrderPlaced is a log parse operation binding the contract event 0xdd67da5d2d18e4e450567dcad68b213da4e86181939bb6c057b280d8c0810d31.
//
// Solidity: event TakerOrderPlaced(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) ParseTakerOrderPlaced(log types.Log) (*MarketTakerOrderPlaced, error) {
	event := new(MarketTakerOrderPlaced)
	if err := _Market.contract.UnpackLog(event, "TakerOrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketTakerOrderRemainingIterator is returned from FilterTakerOrderRemaining and is used to iterate over the raw logs and unpacked data for TakerOrderRemaining events raised by the Market contract.
type MarketTakerOrderRemainingIterator struct {
	Event *MarketTakerOrderRemaining // Event containing the contract specifics and raw log

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
func (it *MarketTakerOrderRemainingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTakerOrderRemaining)
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
		it.Event = new(MarketTakerOrderRemaining)
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
func (it *MarketTakerOrderRemainingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTakerOrderRemainingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTakerOrderRemaining represents a TakerOrderRemaining event raised by the Market contract.
type MarketTakerOrderRemaining struct {
	CycleId      *big.Int
	TakerOrderId int32
	Size         *big.Int
	Side         uint8
	Taker        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTakerOrderRemaining is a free log retrieval operation binding the contract event 0x1f1f20e1bedacc81242c958ebe95e161d407d3533075fad36d452ca697f1eed2.
//
// Solidity: event TakerOrderRemaining(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) FilterTakerOrderRemaining(opts *bind.FilterOpts, cycleId []*big.Int, taker []common.Address) (*MarketTakerOrderRemainingIterator, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "TakerOrderRemaining", cycleIdRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &MarketTakerOrderRemainingIterator{contract: _Market.contract, event: "TakerOrderRemaining", logs: logs, sub: sub}, nil
}

// WatchTakerOrderRemaining is a free log subscription operation binding the contract event 0x1f1f20e1bedacc81242c958ebe95e161d407d3533075fad36d452ca697f1eed2.
//
// Solidity: event TakerOrderRemaining(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) WatchTakerOrderRemaining(opts *bind.WatchOpts, sink chan<- *MarketTakerOrderRemaining, cycleId []*big.Int, taker []common.Address) (event.Subscription, error) {

	var cycleIdRule []interface{}
	for _, cycleIdItem := range cycleId {
		cycleIdRule = append(cycleIdRule, cycleIdItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "TakerOrderRemaining", cycleIdRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTakerOrderRemaining)
				if err := _Market.contract.UnpackLog(event, "TakerOrderRemaining", log); err != nil {
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

// ParseTakerOrderRemaining is a log parse operation binding the contract event 0x1f1f20e1bedacc81242c958ebe95e161d407d3533075fad36d452ca697f1eed2.
//
// Solidity: event TakerOrderRemaining(uint256 indexed cycleId, int32 takerOrderId, uint256 size, uint8 side, address indexed taker)
func (_Market *MarketFilterer) ParseTakerOrderRemaining(log types.Log) (*MarketTakerOrderRemaining, error) {
	event := new(MarketTakerOrderRemaining)
	if err := _Market.contract.UnpackLog(event, "TakerOrderRemaining", log); err != nil {
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
