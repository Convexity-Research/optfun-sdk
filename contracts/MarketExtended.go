package contracts

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"unsafe"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MARK_PX_PRECOMPILE is the address of the mark price precompile contract
var MARK_PX_PRECOMPILE = common.HexToAddress("0x0000000000000000000000000000000000000806")

// ErrOraclePriceCallFailed represents an error when the oracle price call fails
var ErrOraclePriceCallFailed = errors.New("oracle price call failed")

// MarketSide represents the different order sides in the market
const (
	CALL_BUY  = uint8(0) // Buy call options
	CALL_SELL = uint8(1) // Sell call options
	PUT_BUY   = uint8(2) // Buy put options
	PUT_SELL  = uint8(3) // Sell put options
)

// getCallerFromContract extracts the ContractCaller from a BoundContract using reflection
func getCallerFromContract(contract *bind.BoundContract) bind.ContractCaller {
	contractValue := reflect.ValueOf(contract).Elem()
	callerField := contractValue.FieldByName("caller")
	if !callerField.IsValid() {
		return nil
	}

	// Use unsafe to access unexported field
	callerPtr := unsafe.Pointer(callerField.UnsafeAddr())
	caller := *(*bind.ContractCaller)(callerPtr)
	return caller
}

// GetOraclePrice mimics the Solidity _getOraclePrice() internal function
// It calls the mark price precompile and adjusts the price based on collateral decimals
func (_Market *MarketCaller) GetOraclePrice(opts *bind.CallOpts) (uint64, error) {
	// Extract the client from the existing contract binding
	client := getCallerFromContract(_Market.contract)
	if client == nil {
		return 0, errors.New("failed to extract client from contract")
	}

	// Hardcode collateral decimals to 6
	collateralDecimals := uint8(6)

	// Prepare call data: encode uint256(0)
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return 0, err
	}

	arguments := abi.Arguments{
		{Type: uint256Type},
	}

	callData, err := arguments.Pack(big.NewInt(0))
	if err != nil {
		return 0, err
	}

	// Make static call to precompile
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &MARK_PX_PRECOMPILE,
		Data: callData,
	}, nil)

	if err != nil {
		return 0, ErrOraclePriceCallFailed
	}

	if len(result) == 0 {
		return 0, ErrOraclePriceCallFailed
	}

	// Decode result as uint64
	uint64Type, err := abi.NewType("uint64", "", nil)
	if err != nil {
		return 0, err
	}

	returnArgs := abi.Arguments{
		{Type: uint64Type},
	}

	decoded, err := returnArgs.Unpack(result)
	if err != nil {
		return 0, err
	}

	if len(decoded) == 0 {
		return 0, ErrOraclePriceCallFailed
	}

	rawPrice, ok := decoded[0].(uint64)
	if !ok {
		return 0, ErrOraclePriceCallFailed
	}

	// Price always returned with 1 extra decimal, so subtract by 1 from collateral decimals
	// Apply adjustment: price * 10^(collateralDecimals - 1)
	if collateralDecimals == 0 {
		return rawPrice, nil
	}

	adjustment := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(collateralDecimals-1)), nil)
	adjustedPrice := new(big.Int).Mul(big.NewInt(int64(rawPrice)), adjustment)

	// Check if result fits in uint64
	if !adjustedPrice.IsUint64() {
		return 0, errors.New("adjusted price overflow")
	}

	return adjustedPrice.Uint64(), nil
}

// GetOraclePrice mimics the Solidity _getOraclePrice() internal function
// It calls the mark price precompile and adjusts the price based on collateral decimals
func (_Market *Market) GetOraclePrice(opts *bind.CallOpts) (uint64, error) {
	return _Market.MarketCaller.GetOraclePrice(opts)
}

// GetOraclePrice mimics the Solidity _getOraclePrice() internal function
// It calls the mark price precompile and adjusts the price based on collateral decimals
func (_Market *MarketSession) GetOraclePrice() (uint64, error) {
	return _Market.Contract.GetOraclePrice(&_Market.CallOpts)
}

// GetOraclePrice mimics the Solidity _getOraclePrice() internal function
// It calls the mark price precompile and adjusts the price based on collateral decimals
func (_Market *MarketCallerSession) GetOraclePrice() (uint64, error) {
	return _Market.Contract.GetOraclePrice(&_Market.CallOpts)
}
