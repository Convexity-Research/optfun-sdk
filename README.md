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

    fmt.Printf("Cycle Active: %v, Strike Price: %v, Settlement Price: %v\n", 
        cycle.Active, cycle.StrikePrice, cycle.SettlementPrice)
}
```

#### Check User Balance
```go
func getUserBalance(client *ethclient.Client, userAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    balance, err := market.Balances(nil, userAddress)
    if err != nil {
        log.Fatalf("Failed to get user balance: %v", err)
    }

    fmt.Printf("User Balance: %v\n", balance)
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

    // Order parameters
    optionType := uint8(0)        // 0 = Call, 1 = Put
    side := uint8(0)              // 0 = Buy, 1 = Sell
    size := big.NewInt(10)        // 10 contracts
    limitPrice := big.NewInt(100) // Price per contract

    tx, err := market.PlaceOrder0(txOpts, optionType, side, size, limitPrice)
    if err != nil {
        log.Fatalf("Failed to place order: %v", err)
    }

    fmt.Printf("Order placed. Transaction hash: %s\n", tx.Hash().Hex())
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

#### Market Buy (Long Position)
```go
func buyCallOptions(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    size := big.NewInt(5) // Buy 5 call options

    tx, err := market.Long0(txOpts, size)
    if err != nil {
        log.Fatalf("Failed to buy options: %v", err)
    }

    fmt.Printf("Call options purchased. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Market Sell (Short Position)
```go
func sellCallOptions(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    size := big.NewInt(3) // Sell 3 call options

    tx, err := market.Short0(txOpts, size)
    if err != nil {
        log.Fatalf("Failed to sell options: %v", err)
    }

    fmt.Printf("Call options sold. Transaction hash: %s\n", tx.Hash().Hex())
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

    // Get user's trader ID (you'll need to track this or get it from events)
    traderId := big.NewInt(1) // Example trader ID

    positions, err := market.Positions(nil, traderId)
    if err != nil {
        log.Fatalf("Failed to get positions: %v", err)
    }

    fmt.Printf("Long Calls: %v, Short Calls: %v, Long Puts: %v, Short Puts: %v\n",
        positions.LongCalls, positions.ShortCalls, positions.LongPuts, positions.ShortPuts)
}
```

#### View Taker Queue
```go
func viewTakerQueue(client *ethclient.Client, isPut bool, isBid bool) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    queue, err := market.ViewTakerQueue(nil, isPut, isBid)
    if err != nil {
        log.Fatalf("Failed to get taker queue: %v", err)
    }

    fmt.Printf("Taker Queue Length: %d\n", len(queue))
    for i, entry := range queue {
        fmt.Printf("Entry %d: Trader %s, Size %v\n", i, entry.Trader.Hex(), entry.Size)
    }
}

#### Get Orderbook Data

**Note**: This function replicates the internal `dumpBook` functionality from the Solidity contract by reading the bitmap data structures using only external view methods and implementing the bit-scanning logic in Go.

```go
func getOrderbook(client *ethclient.Client, isBid bool, isPut bool) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    
    levels, err := contracts.GetOrderbook(client, marketAddress, isBid, isPut)
    if err != nil {
        log.Fatalf("Failed to get orderbook: %v", err)
    }

    fmt.Printf("Orderbook Levels Count: %d\n", len(levels))
    for i, level := range levels {
        fmt.Printf("Level %d: Tick %d, Volume %v\n", i, level.Tick, level.Vol)
    }
}

// Example: Get bid side for calls
func getBidCallOrderbook(client *ethclient.Client) {
    getOrderbook(client, true, false) // isBid=true, isPut=false
}

// Example: Get ask side for puts  
func getAskPutOrderbook(client *ethclient.Client) {
    getOrderbook(client, false, true) // isBid=false, isPut=true
}
```

### 5. Market Administration (Owner Only)

#### Start New Cycle
```go
func startCycle(client *ethclient.Client, txOpts *bind.TransactOpts, expiry *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.StartCycle(txOpts, expiry)
    if err != nil {
        log.Fatalf("Failed to start cycle: %v", err)
    }

    fmt.Printf("New cycle started. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Fix Settlement Price
```go
func fixPrice(client *ethclient.Client, txOpts *bind.TransactOpts) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.FixPrice(txOpts)
    if err != nil {
        log.Fatalf("Failed to fix price: %v", err)
    }

    fmt.Printf("Price fixed. Transaction hash: %s\n", tx.Hash().Hex())
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

## Contract Events

The Market contract emits several events that you can listen to:

- `OrderPlaced`: When a new order is placed
- `OrderFilled`: When an order is matched and filled
- `OrderCancelled`: When an order is cancelled
- `CollateralDeposited`: When collateral is deposited
- `CollateralWithdrawn`: When collateral is withdrawn
- `CycleStarted`: When a new options cycle begins
- `PriceFixed`: When the settlement price is fixed
- `Settled`: When a position is settled

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

    // 2. Check balance
    getUserBalance(client, fromAddress)

    // 3. Place a buy order
    placeLimitOrder(client, txOpts)

    // 4. Check positions
    getUserPositions(client, fromAddress)

    // 5. Withdraw some collateral
    withdrawCollateral(client, txOpts)
}
```

This SDK provides a comprehensive interface for interacting with OptFun options markets, enabling sophisticated trading strategies and risk management.
