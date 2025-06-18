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

#### Deposit Collateral with Signature (Gasless)
```go
func depositCollateralWithSignature(client *ethclient.Client, txOpts *bind.TransactOpts, signature []byte) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    depositAmount := big.NewInt(1000000) // 1 USDC (6 decimals)

    tx, err := market.DepositCollateral0(txOpts, depositAmount, signature)
    if err != nil {
        log.Fatalf("Failed to deposit collateral with signature: %v", err)
    }

    fmt.Printf("Collateral deposited with signature. Transaction hash: %s\n", tx.Hash().Hex())
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
func buyOptions(client *ethclient.Client, txOpts *bind.TransactOpts) {
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

    size := big.NewInt(5) // Buy 5 options

    tx, err := market.Long(txOpts, size, activeCycle)
    if err != nil {
        log.Fatalf("Failed to buy options: %v", err)
    }

    fmt.Printf("Options purchased. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Market Sell (Short Position)
```go
func sellOptions(client *ethclient.Client, txOpts *bind.TransactOpts) {
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

    size := big.NewInt(3) // Sell 3 options

    tx, err := market.Short(txOpts, size, activeCycle)
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

### 5. Orderbook Data

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

### 6. Market Information and Utilities

#### Get Number of Traders
```go
func getNumTraders(client *ethclient.Client) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    numTraders, err := market.GetNumTraders(nil)
    if err != nil {
        log.Fatalf("Failed to get number of traders: %v", err)
    }

    fmt.Printf("Number of traders: %v\n", numTraders)
}
```

#### Get Trader Address by Index
```go
func getTraderAddress(client *ethclient.Client, traderIndex *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    traderAddress, err := market.Traders(nil, traderIndex)
    if err != nil {
        log.Fatalf("Failed to get trader address: %v", err)
    }

    fmt.Printf("Trader at index %v: %s\n", traderIndex, traderAddress.Hex())
}
```

#### Check User Orders
```go
func getUserOrders(client *ethclient.Client, userAddress common.Address, orderIndex *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    orderInfo, err := market.UserOrders(nil, userAddress, orderIndex)
    if err != nil {
        log.Fatalf("Failed to get user orders: %v", err)
    }

    fmt.Printf("User order at index %v: %v\n", orderIndex, orderInfo)
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

### 7. Market Administration (Owner Only)

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
func settleChunk(client *ethclient.Client, txOpts *bind.TransactOpts, maxSettlements *big.Int) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.SettleChunk(txOpts, maxSettlements)
    if err != nil {
        log.Fatalf("Failed to settle chunk: %v", err)
    }

    fmt.Printf("Chunk settled. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Set Trusted Forwarder
```go
func setTrustedForwarder(client *ethclient.Client, txOpts *bind.TransactOpts, forwarderAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    tx, err := market.SetTrustedForwarder(txOpts, forwarderAddress)
    if err != nil {
        log.Fatalf("Failed to set trusted forwarder: %v", err)
    }

    fmt.Printf("Trusted forwarder set. Transaction hash: %s\n", tx.Hash().Hex())
}
```

#### Check if Forwarder is Trusted
```go
func checkTrustedForwarder(client *ethclient.Client, forwarderAddress common.Address) {
    marketAddress := common.HexToAddress("YOUR_MARKET_ADDRESS")
    market, err := contracts.NewMarket(marketAddress, client)
    if err != nil {
        log.Fatalf("Failed to load market: %v", err)
    }

    isTrusted, err := market.IsTrustedForwarder(nil, forwarderAddress)
    if err != nil {
        log.Fatalf("Failed to check trusted forwarder: %v", err)
    }

    fmt.Printf("Forwarder %s is trusted: %v\n", forwarderAddress.Hex(), isTrusted)
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

- `CollateralDeposited`: When collateral is deposited
- `CollateralWithdrawn`: When collateral is withdrawn
- `LimitOrderPlaced`: When a new limit order is placed
- `LimitOrderFilled`: When an order is matched and filled
- `LimitOrderCancelled`: When an order is cancelled
- `TakerOrderPlaced`: When a market order is placed
- `TakerOrderRemaining`: When a market order has remaining size
- `CycleStarted`: When a new options cycle begins
- `CycleSettled`: When a cycle is settled
- `PriceFixed`: When the settlement price is fixed
- `Settled`: When a position is settled
- `Liquidated`: When a position is liquidated

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
