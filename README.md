# OptFun SDK Usage Guide

The OptFun SDK provides functionality to interact with OptFun options markets and perform operations like trading options, managing collateral, placing orders, and accessing market information.

## Installation

```go
go get github.com/Convexity-Research/optfun-sdk
```

## Quick Start

```go
import (
    optfunsdk "github.com/Convexity-Research/optfun-sdk/sdk"
    contracts "github.com/Convexity-Research/optfun-sdk/contracts"
)

// Initialize the SDK
client, fromAddress, txOpts, err := optfunsdk.LoadSDK()
if err != nil {
    log.Fatalf("Failed to load SDK: %v", err)
}
```

## Constants

The SDK provides constants for market operations:

```go
// Market sides for order placement  
const (
    CALL_BUY  = uint8(0) // Buy call options
    CALL_SELL = uint8(1) // Sell call options  
    PUT_BUY   = uint8(2) // Buy put options
    PUT_SELL  = uint8(3) // Sell put options
)

// Pre-defined market address
var BtcMarket = common.HexToAddress("0xB7C609cFfa0e47DB2467ea03fF3e598bF59361A5")
```

## Core Features

### 1. Market Information

#### Get Active Cycle
```go
func getActiveCycle(client *ethclient.Client) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    activeCycle, err := market.ActiveCycle(nil)
    if err != nil {
        log.Fatalf("Failed to get active cycle: %v", err)
    }

    fmt.Printf("Active Cycle: %v\n", activeCycle)
}
```

#### Get Cycle Information
```go
func getCycleInfo(client *ethclient.Client, cycleId *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    cycle, err := market.Cycles(nil, cycleId)
    if err != nil {
        log.Fatalf("Failed to get cycle info: %v", err)
    }

    fmt.Printf("Cycle Settled: %v, Strike Price: %v, Settlement Price: %v\n", 
        cycle.IsSettled, cycle.StrikePrice, cycle.SettlementPrice)
}
```

#### Check User Account Information
```go
func getUserAccount(client *ethclient.Client, userAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    account, err := market.UserAccounts(nil, userAddress)
    if err != nil {
        log.Fatalf("Failed to get user account: %v", err)
    }

    fmt.Printf("Balance: %v, Active in Cycle: %v\n", account.Balance, account.ActiveInCycle)
    fmt.Printf("Long Calls: %v, Short Calls: %v, Long Puts: %v, Short Puts: %v\n",
        account.LongCalls, account.ShortCalls, account.LongPuts, account.ShortPuts)
    fmt.Printf("Pending - Long Calls: %v, Short Calls: %v, Long Puts: %v, Short Puts: %v\n",
        account.PendingLongCalls, account.PendingShortCalls, account.PendingLongPuts, account.PendingShortPuts)
}
```

### 2. Collateral Management

#### Deposit Collateral
```go
func depositCollateral(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    // Get collateral token address and approve spending
    collateralTokenAddress, err := market.CollateralToken(nil)
    if err != nil {
        log.Fatalf("Failed to get collateral token: %v", err)
    }

    collateralToken, err := contracts.NewErc20(collateralTokenAddress, client)
    if err != nil {
        log.Fatalf("Failed to load collateral token: %v", err)
    }

    depositAmount := big.NewInt(1000000) // 1 USDC (6 decimals)
    
    // Approve token spending
    _, err = collateralToken.Approve(txOpts, marketAddress, depositAmount)
    if err != nil {
        log.Fatalf("Failed to approve token: %v", err)
    }

    // Deposit collateral
    tx, err := market.DepositCollateral(txOpts, depositAmount)
    if err != nil {
        log.Fatalf("Failed to deposit collateral: %v", err)
    }

    fmt.Printf("Collateral deposited. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Withdraw Collateral
```go
func withdrawCollateral(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    withdrawAmount := big.NewInt(500000) // 0.5 USDC

    tx, err := market.WithdrawCollateral(txOpts, withdrawAmount)
    if err != nil {
        log.Fatalf("Failed to withdraw collateral: %v", err)
    }

    fmt.Printf("Collateral withdrawn. Transaction hash: %s\n", tx.Hash().Hex())
}
```

### 3. Trading Operations

#### Place Limit Order
```go
func placeLimitOrder(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    // Get active cycle
    activeCycle, err := market.ActiveCycle(nil)
    if err != nil {
        log.Fatalf("Failed to get active cycle: %v", err)
    }

    // Order parameters
    side := uint8(0)              // 0 = Buy, 1 = Sell
    size := big.NewInt(10)        // 10 contracts
    limitPrice := big.NewInt(100) // Price per contract
    cycleId := activeCycle        // Use active cycle

    tx, err := market.PlaceOrder(txOpts, side, size, limitPrice, cycleId)
    if err != nil {
        log.Fatalf("Failed to place order: %v", err)
    }

    fmt.Printf("Order placed. Transaction hash: %s\n", tx.Hash().Hex())
    // Note: PlaceOrder returns an orderId, but in Go bindings this is available in the transaction receipt
}
```

#### Cancel Order
```go
func cancelOrder(client *ethclient.Client, txOpts *bind.TransactOpts, orderId *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.CancelOrder(txOpts, orderId)
    if err != nil {
        log.Fatalf("Failed to cancel order: %v", err)
    }

    fmt.Printf("Order cancelled. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Long Position
```go
func longPosition(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    // Get active cycle
    activeCycle, err := market.ActiveCycle(nil)
    if err != nil {
        log.Fatalf("Failed to get active cycle: %v", err)
    }

    size := big.NewInt(5)             // Buy 5 options
    limitPriceBuy := big.NewInt(100)  // Maximum price willing to pay
    limitPriceSell := big.NewInt(90)  // Minimum price willing to sell at  
    cycleId := activeCycle            // Use active cycle

    tx, err := market.Long(txOpts, size, limitPriceBuy, limitPriceSell, cycleId)
    if err != nil {
        log.Fatalf("Failed to buy options: %v", err)
    }

    fmt.Printf("Options purchased. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Short Position
```go
func shortPosition(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    // Get active cycle
    activeCycle, err := market.ActiveCycle(nil)
    if err != nil {
        log.Fatalf("Failed to get active cycle: %v", err)
    }

    size := big.NewInt(3)             // Sell 3 options
    limitPriceBuy := big.NewInt(80)   // Maximum price willing to pay to buy back
    limitPriceSell := big.NewInt(110) // Minimum price willing to sell at
    cycleId := activeCycle            // Use active cycle

    tx, err := market.Short(txOpts, size, limitPriceBuy, limitPriceSell, cycleId)
    if err != nil {
        log.Fatalf("Failed to sell options: %v", err)
    }

    fmt.Printf("Options sold. Transaction hash: %s\n", tx.Hash().Hex())
}
```

### 4. Position Management

#### Get User Positions
```go
func getUserPositions(client *ethclient.Client, userAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    account, err := market.UserAccounts(nil, userAddress)
    if err != nil {
        log.Fatalf("Failed to get user account: %v", err)
    }

    fmt.Printf("Current Positions:\n")
    fmt.Printf("  Long Calls: %v, Short Calls: %v\n", account.LongCalls, account.ShortCalls)
    fmt.Printf("  Long Puts: %v, Short Puts: %v\n", account.LongPuts, account.ShortPuts)
    fmt.Printf("Pending Positions:\n")
    fmt.Printf("  Pending Long Calls: %v, Pending Short Calls: %v\n", account.PendingLongCalls, account.PendingShortCalls)
    fmt.Printf("  Pending Long Puts: %v, Pending Short Puts: %v\n", account.PendingLongPuts, account.PendingShortPuts)
}
```

#### Check if User is Liquidatable
```go
func checkLiquidatable(client *ethclient.Client, userAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    isLiquidatable, err := market.IsLiquidatable(nil, userAddress)
    if err != nil {
        log.Fatalf("Failed to check liquidation status: %v", err)
    }

    fmt.Printf("User %s is liquidatable: %v\n", userAddress.Hex(), isLiquidatable)
}
```

#### Check if User is Liquidatable at Specific Price
```go
func checkLiquidatableAtPrice(client *ethclient.Client, userAddress common.Address, price uint64) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    isLiquidatable, err := market.IsLiquidatable0(nil, userAddress, price)
    if err != nil {
        log.Fatalf("Failed to check liquidation status at price: %v", err)
    }

    fmt.Printf("User %s is liquidatable at price %d: %v\n", userAddress.Hex(), price, isLiquidatable)
}
```

#### Liquidate User
```go
func liquidateUser(client *ethclient.Client, txOpts *bind.TransactOpts, userAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.Liquidate(txOpts, userAddress)
    if err != nil {
        log.Fatalf("Failed to liquidate user: %v", err)
    }

    fmt.Printf("User liquidated. Transaction hash: %s\n", tx.Hash().Hex())
}
```

### 5. Oracle Price Information

#### Get Current Oracle Price
```go
func getOraclePrice(client *ethclient.Client) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    price, err := market.GetOraclePrice(nil)
    if err != nil {
        log.Fatalf("Failed to get oracle price: %v", err)
    }

    fmt.Printf("Current oracle price: %d\n", price)
}
```

#### Check if Address is Whitelisted
```go
func checkWhitelist(client *ethclient.Client, address common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    isWhitelisted, err := market.Whitelist(nil, address)
    if err != nil {
        log.Fatalf("Failed to check whitelist status: %v", err)
    }

    fmt.Printf("Address %s is whitelisted: %v\n", address.Hex(), isWhitelisted)
}
```

### 6. Market Administration

#### Start New Cycle
```go
func startCycle(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.StartCycle(txOpts)
    if err != nil {
        log.Fatalf("Failed to start cycle: %v", err)
    }

    fmt.Printf("New cycle started. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Settle Chunk
```go
func settleChunk(client *ethclient.Client, txOpts *bind.TransactOpts, maxSettlements *big.Int, pauseNextCycle bool) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.SettleChunk(txOpts, maxSettlements, pauseNextCycle)
    if err != nil {
        log.Fatalf("Failed to settle chunk: %v", err)
    }

    fmt.Printf("Chunk settled. Transaction hash: %s\n", tx.Hash().Hex())
}
```

## SDK Utilities

### Get Current Gas Price
```go
func updateGasPrice(client *ethclient.Client, txOpts *bind.TransactOpts) {
    gasPrice, err := optfunsdk.GetGasPrice(client)
    if err != nil {
        log.Fatalf("Failed to get gas price: %v", err)
    }
    
    txOpts.GasPrice = gasPrice
    fmt.Printf("Updated gas price: %v\n", gasPrice)
}
```

### Get Next Nonce
```go
func updateNonce(client *ethclient.Client, txOpts *bind.TransactOpts, address common.Address) {
    nonce, err := optfunsdk.GetNextNonce(client, address)
    if err != nil {
        log.Fatalf("Failed to get nonce: %v", err)
    }
    
    txOpts.Nonce = big.NewInt(int64(nonce))
    fmt.Printf("Updated nonce: %v\n", nonce)
}
```

## Environment Configuration

The SDK requires certain environment variables to be set up. Create a `.env` file in your project root with the following variables:

```env
# Required
RPC_URL=https://your-hyperevm-rpc-url
PRIVATE_KEY=your-private-key

# Optional
CHAIN_ID=999
```

### Environment Variables Details:
- `RPC_URL`: Required for all operations. This is your connection to the blockchain.
- `PRIVATE_KEY`: Required only if you need to send transactions (trading, collateral management, etc.).
- `CHAIN_ID`: Optional. Defaults to a HyperEVM. Set this when working with testnet or other chain deployments.

## Best Practices

1. **Transaction Management**
   - Always update nonce and gas price before sending transactions
   - Wait for transaction receipts to confirm operations
   - Handle transaction failures gracefully

2. **Collateral Management**
   - Ensure sufficient collateral before opening positions
   - Monitor collateral ratios to avoid liquidation
   - Approve adequate token allowances

3. **Risk Management**
   - Understand option expiry times and settlement procedures
   - Monitor market conditions and oracle prices
   - Set appropriate position sizes based on risk tolerance

4. **Error Handling**
   - Always check for errors in all SDK operations
   - Implement retry logic for network-related failures
   - Validate addresses and parameters before transactions

5. **Security**
   - Keep private keys secure and never commit them to version control
   - Use environment variables for sensitive configuration
   - Verify contract addresses before interactions

## Example: Complete Trading Workflow

```go
func completeWorkflow() {
    // Initialize SDK
    client, fromAddress, txOpts, err := optfunsdk.LoadSDK()
    if err != nil {
        log.Fatalf("Failed to load SDK: %v", err)
    }

    // 1. Deposit collateral
    depositCollateral(client, txOpts)

    // 2. Check account info
    getUserAccount(client, fromAddress)

    // 3. Place a buy order
    placeLimitOrder(client, txOpts)

    // 4. Check positions
    getUserPositions(client, fromAddress)

    // 5. Withdraw some collateral
    withdrawCollateral(client, txOpts)
}
```

This SDK provides a comprehensive interface for interacting with OptFun options markets, enabling sophisticated trading strategies and risk management.
