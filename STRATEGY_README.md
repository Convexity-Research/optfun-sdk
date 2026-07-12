# Automated Trading Strategy

This script implements an automated trading strategy that leverages the OptFun SDK to place limit orders at the start of market cycles.

## Strategy Overview

The strategy automatically places the following limit orders at the beginning of each market cycle:

1. **Buy Call Options**: 1 contract at $0.01 per contract
2. **Sell Call Options**: 1 contract at $1000 per contract  
3. **Buy Put Options**: 1 contract at $0.01 per contract

## Features

- **Automated Cycle Detection**: Monitors for new market cycles and automatically executes orders
- **Retry Logic**: Implements retry mechanism for failed transactions
- **Market State Validation**: Checks if market is paused before placing orders
- **Graceful Shutdown**: Handles interrupt signals for clean shutdown
- **Account Monitoring**: Displays user account information and position details
- **Transaction Confirmation**: Waits for transaction confirmations and validates success

## Prerequisites

1. **Environment Setup**: Create a `.env` file with the following variables:
   ```
   RPC_URL=<your_ethereum_rpc_url>
   PRIVATE_KEY=<your_private_key_without_0x_prefix>
   CHAIN_ID=999  # HyperEVM chain ID
   ```

2. **Account Requirements**:
   - Sufficient collateral deposited in the market
   - Account must be whitelisted for trading
   - Sufficient balance for transaction fees

## Installation

1. Ensure you have Go 1.24.2 or later installed
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Testing Connection

Before running the automated strategy, test your SDK connection:

```bash
go run test_sdk_connection.go
```

This will verify:
- SDK connection to the blockchain
- Market contract accessibility
- Current cycle and oracle price
- Account information (if private key provided)
- Whitelist status

### Basic Usage

Run the automated strategy:

```bash
go run automated_strategy.go
```

### Build and Run

Build the executables:

```bash
# Build test script
go build -o test_sdk_connection test_sdk_connection.go

# Build automated strategy
go build -o automated_strategy automated_strategy.go
```

Run the executables:

```bash
# Test connection first
./test_sdk_connection

# Run automated strategy
./automated_strategy
```

## Configuration

The strategy configuration can be modified in the `main()` function:

```go
config := StrategyConfig{
    MarketAddress:      sdk.BtcMarket,      // Market contract address
    ContractSize:       big.NewInt(1),      // Number of contracts per order
    BuyCallPrice:       big.NewInt(1),      // Buy call price ($0.01)
    SellCallPrice:      big.NewInt(100000), // Sell call price ($1000)
    BuyPutPrice:        big.NewInt(1),      // Buy put price ($0.01)
    MaxRetries:         3,                  // Maximum retry attempts
    RetryDelay:         5 * time.Second,    // Delay between retries
    CycleCheckInterval: 30 * time.Second,   // How often to check for new cycles
}
```

## Strategy Logic

### Cycle Detection
- The script polls the market every 30 seconds (configurable) to check for new cycles
- When a new cycle is detected, it automatically executes the trading strategy

### Order Placement
- Uses `PlaceMultiOrder` for efficiency to place all orders in a single transaction
- Orders are placed with the following sides:
  - `CALL_BUY`: Buy call options at $0.01
  - `CALL_SELL`: Sell call options at $1000
  - `PUT_BUY`: Buy put options at $0.01

### Risk Management
- Checks if market is paused before placing orders
- Validates transaction success and waits for confirmations
- Implements retry logic for failed transactions
- Monitors account status and position information

## Monitoring

The script provides comprehensive logging:

- Connection status and current cycle information
- Strategy execution details
- Transaction hashes and confirmation status
- Account balance and position information
- Error messages and retry attempts

## Safety Features

1. **Market Pause Check**: Won't place orders if market is paused
2. **Transaction Validation**: Confirms transaction success before proceeding
3. **Graceful Shutdown**: Handles interrupt signals (Ctrl+C) for clean exit
4. **Error Handling**: Comprehensive error handling with retry logic
5. **Account Monitoring**: Displays account status and position information

## Important Notes

⚠️ **Risk Warning**: This is an automated trading strategy. Please ensure you understand the risks involved and have sufficient collateral before running.

⚠️ **Testing**: Test the strategy on a testnet first before using on mainnet.

⚠️ **Monitoring**: Always monitor the strategy while it's running and be prepared to stop it if needed.

## Troubleshooting

### Common Issues

1. **"no private key provided"**: Ensure your `.env` file contains a valid `PRIVATE_KEY`
2. **"failed to connect"**: Check your `RPC_URL` in the `.env` file
3. **"insufficient balance"**: Ensure you have sufficient collateral deposited
4. **"not whitelisted"**: Contact the market operators to get whitelisted

### Logs

The script provides detailed logs to help diagnose issues:
- Connection status
- Cycle detection
- Order placement attempts
- Transaction confirmations
- Account information

## Example Output

```
2024/01/15 10:30:00 Connected to market at 0x137b3dab3366854a1f68a7b893d3f4fad1429896
2024/01/15 10:30:00 Current active cycle: 123
2024/01/15 10:30:00 Trading from address: 0x1234567890abcdef...
2024/01/15 10:30:00 Starting automated trading strategy...
2024/01/15 10:30:00 Strategy configuration:
2024/01/15 10:30:00   Contract size: 1
2024/01/15 10:30:00   Buy call price: $1
2024/01/15 10:30:00   Sell call price: $100000
2024/01/15 10:30:00   Buy put price: $1
2024/01/15 10:30:30 New cycle detected: 124 (previous: 123)
2024/01/15 10:30:30 Executing strategy for cycle 124
2024/01/15 10:30:30 Current oracle price: $45000
2024/01/15 10:30:30 Placing multi-order:
2024/01/15 10:30:30   Buy 1 calls at $1
2024/01/15 10:30:30   Sell 1 calls at $100000
2024/01/15 10:30:30   Buy 1 puts at $1
2024/01/15 10:30:35 Multi-order transaction submitted: 0xabcdef123456...
2024/01/15 10:30:40 Waiting for transaction confirmation...
2024/01/15 10:30:45 Transaction confirmed in block 1234567
2024/01/15 10:30:45 Strategy execution completed for cycle 124
``` 