# Racks.Win Contract Documentation

## Overview

racks.win is a decentralized options trading platform on HyperEVM that enables traders to buy and sell 1-minute expiry at-the-money (ATM) options. The platform provides a streamlined trading experience for short-term options strategies with automated settlement and collateral management.

### Key Features
- **High leverage**: Up to 1000x leverage on deposited collateral
- **1-minute expiry cycles**: Each trading cycle lasts exactly 60 seconds from initiation to settlement
- **At-the-money strikes**: All options are automatically struck at the current oracle price when the cycle begins
- **Binary settlement**: Options expire either in-the-money (ITM) or out-of-the-money (OTM)
- **Gasless trading**: Sponsored by the racks.win team using ERC2771 meta-transactions
- **Automated liquidations**: Protect system solvency through automated position liquidation

### Use Cases
- **Speculating**: Capitalize on small price movements with minimal capital via high leverage
- **Hedging**: Quickly hedge spot positions over a 1-minute period
- **Yield Farming**: Provide liquidity by selling options and collecting premiums
- **Market Making**: Simultaneous buying and selling to capture bid-ask spreads


## Public/External Functions

### Collateral Management

#### `depositCollateral(uint256 amount, bytes memory signature)`
Deposits USDT collateral with whitelist signature verification.
- **Parameters**: 
  - `amount`: Amount of USDT to deposit (6 decimals)
  - `signature`: Signature from whitelist signer to authorize the address
- **Effects**: Adds user to whitelist and credits their balance
- **Events**: `CollateralDeposited`

#### `depositCollateral(uint256 amount)`
Deposits USDT collateral for already whitelisted users.
- **Parameters**: `amount`: Amount of USDT to deposit
- **Requirements**: User must be whitelisted
- **Events**: `CollateralDeposited`

#### `withdrawCollateral(uint256 amount)`
Withdraws USDT collateral from the platform.
- **Parameters**: `amount`: Amount to withdraw
- **Requirements**: 
  - User must have no open positions or orders
  - Sufficient balance available
- **Events**: `CollateralWithdrawn`

## SDK - OrderEngine

The OrderEngine provides a Go SDK for interacting with the Racks.Win trading API using EIP-712 signed requests.

### Setup
```go
engine, err := NewOrderEngine()
```

**Required Environment Variables:**
- `PRIVATE_KEY`: Your private key for signing requests
- `API_URL`: Backend API URL (defaults to `http://localhost:8080`)

### Functions

#### `Long(OrderParams)`
Creates a long position (buy calls + sell puts).

#### `Short(OrderParams)` 
Creates a short position (buy puts + sell calls).

#### `CancelAndClose(CancelAndCloseParams)`
Cancels all orders and closes positions for a market.

All methods use EIP-712 signing and submit requests directly to the trading API.
