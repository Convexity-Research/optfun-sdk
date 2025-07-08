package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Convexity-Research/optfun-sdk/contracts"
	"github.com/Convexity-Research/optfun-sdk/sdk"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// StrategyConfig holds the configuration for the automated trading strategy
type StrategyConfig struct {
	MarketAddress      common.Address
	ContractSize       *big.Int
	BuyCallPrice       *big.Int // $0.01 per contract
	SellCallPrice      *big.Int // $1000 per contract
	BuyPutPrice        *big.Int // $0.01 per contract
	MaxRetries         int
	RetryDelay         time.Duration
	CycleCheckInterval time.Duration
}

// TradingStrategy represents the automated trading strategy
type TradingStrategy struct {
	client    *ethclient.Client
	market    *contracts.Market
	config    StrategyConfig
	lastCycle *big.Int
}

// NewTradingStrategy creates a new instance of the trading strategy
func NewTradingStrategy(config StrategyConfig) (*TradingStrategy, error) {
	// Load SDK
	client, fromAddress, auth, err := sdk.LoadSDK()
	if err != nil {
		return nil, fmt.Errorf("failed to load SDK: %w", err)
	}

	if auth == nil {
		return nil, fmt.Errorf("no private key provided - cannot execute transactions")
	}

	// Create market contract instance
	market, err := contracts.NewMarket(config.MarketAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create market contract instance: %w", err)
	}

	// Get current active cycle
	activeCycle, err := market.GetActiveCycle(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get active cycle: %w", err)
	}

	log.Printf("Connected to market at %s", config.MarketAddress.Hex())
	log.Printf("Current active cycle: %s", activeCycle.String())
	log.Printf("Trading from address: %s", fromAddress.Hex())

	return &TradingStrategy{
		client:    client,
		market:    market,
		config:    config,
		lastCycle: activeCycle,
	}, nil
}

// Run starts the automated trading strategy
func (ts *TradingStrategy) Run(ctx context.Context) error {
	log.Println("Starting automated trading strategy...")
	log.Printf("Strategy configuration:")
	log.Printf("  Contract size: %s", ts.config.ContractSize.String())
	log.Printf("  Buy call price: $%s", ts.config.BuyCallPrice.String())
	log.Printf("  Sell call price: $%s", ts.config.SellCallPrice.String())
	log.Printf("  Buy put price: $%s", ts.config.BuyPutPrice.String())

	ticker := time.NewTicker(ts.config.CycleCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Strategy stopped by context cancellation")
			return ctx.Err()
		case <-ticker.C:
			if err := ts.checkAndExecuteStrategy(); err != nil {
				log.Printf("Error in strategy execution: %v", err)
			}
		}
	}
}

// checkAndExecuteStrategy checks if a new cycle has started and executes the strategy
func (ts *TradingStrategy) checkAndExecuteStrategy() error {
	// Get current active cycle
	activeCycle, err := ts.market.GetActiveCycle(nil)
	if err != nil {
		return fmt.Errorf("failed to get active cycle: %w", err)
	}

	// Check if this is a new cycle
	if activeCycle.Cmp(ts.lastCycle) > 0 {
		log.Printf("New cycle detected: %s (previous: %s)", activeCycle.String(), ts.lastCycle.String())

		// Execute the strategy for the new cycle
		if err := ts.executeStrategy(activeCycle); err != nil {
			return fmt.Errorf("failed to execute strategy for cycle %s: %w", activeCycle.String(), err)
		}

		ts.lastCycle = activeCycle
	}

	return nil
}

// executeStrategy places the limit orders for the given cycle
func (ts *TradingStrategy) executeStrategy(cycleId *big.Int) error {
	log.Printf("Executing strategy for cycle %s", cycleId.String())

	// Check if market is paused
	paused, err := ts.market.Paused(nil)
	if err != nil {
		return fmt.Errorf("failed to check if market is paused: %w", err)
	}

	if paused {
		log.Println("Market is paused, skipping strategy execution")
		return nil
	}

	// Get current oracle price for reference
	oraclePrice, err := ts.market.GetOraclePrice(nil)
	if err != nil {
		log.Printf("Warning: failed to get oracle price: %v", err)
	} else {
		log.Printf("Current oracle price: $%d", oraclePrice)
	}

	// Place orders using PlaceMultiOrder for efficiency
	// Order sides: CALL_BUY, CALL_SELL, PUT_BUY
	sides := []uint8{
		contracts.CALL_BUY,  // Buy calls for $0.01
		contracts.CALL_SELL, // Sell calls for $1000
		contracts.PUT_BUY,   // Buy puts for $0.01
	}

	sizes := []*big.Int{
		ts.config.ContractSize, // Buy calls
		ts.config.ContractSize, // Sell calls
		ts.config.ContractSize, // Buy puts
	}

	limitPrices := []*big.Int{
		ts.config.BuyCallPrice,  // $0.01 for buying calls
		ts.config.SellCallPrice, // $1000 for selling calls
		ts.config.BuyPutPrice,   // $0.01 for buying puts
	}

	log.Printf("Placing multi-order:")
	log.Printf("  Buy %s calls at $%s", ts.config.ContractSize.String(), ts.config.BuyCallPrice.String())
	log.Printf("  Sell %s calls at $%s", ts.config.ContractSize.String(), ts.config.SellCallPrice.String())
	log.Printf("  Buy %s puts at $%s", ts.config.ContractSize.String(), ts.config.BuyPutPrice.String())

	// Execute the multi-order with retries
	var tx *types.Transaction

	for attempt := 0; attempt <= ts.config.MaxRetries; attempt++ {
		if attempt > 0 {
			log.Printf("Retry attempt %d/%d", attempt, ts.config.MaxRetries)
			time.Sleep(ts.config.RetryDelay)
		}

		// Get fresh auth for each attempt
		_, _, auth, authErr := sdk.LoadSDK()
		if authErr != nil {
			log.Printf("Failed to get fresh auth: %v", authErr)
			continue
		}

		tx, err = ts.market.PlaceMultiOrder(auth, sides, sizes, limitPrices, cycleId)
		if err != nil {
			log.Printf("Failed to place multi-order (attempt %d): %v", attempt+1, err)
			if attempt == ts.config.MaxRetries {
				return fmt.Errorf("failed to place multi-order after %d attempts: %w", ts.config.MaxRetries+1, err)
			}
			continue
		}

		log.Printf("Multi-order transaction submitted: %s", tx.Hash().Hex())
		break
	}

	// Wait for transaction confirmation
	if tx != nil {
		log.Printf("Waiting for transaction confirmation...")
		receipt, err := bind.WaitMined(context.Background(), ts.client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for transaction confirmation: %w", err)
		}

		if receipt.Status == 0 {
			return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
		}

		log.Printf("Transaction confirmed in block %d", receipt.BlockNumber)
	}

	log.Printf("Strategy execution completed for cycle %s", cycleId.String())
	return nil
}

// GetUserAccountInfo retrieves and logs user account information
func (ts *TradingStrategy) GetUserAccountInfo() error {
	_, fromAddress, _, err := sdk.LoadSDK()
	if err != nil {
		return fmt.Errorf("failed to load SDK: %w", err)
	}

	account, err := ts.market.GetUserAccounts(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get user account: %w", err)
	}

	log.Printf("User Account Information:")
	log.Printf("  Active in cycle: %t", account.ActiveInCycle)
	log.Printf("  Liquidation queued: %t", account.LiquidationQueued)
	log.Printf("  Balance: %d", account.Balance)
	log.Printf("  Liquidation fee owed: %d", account.LiquidationFeeOwed)
	log.Printf("  Scratch PnL: %d", account.ScratchPnL)
	log.Printf("  Long calls: %d", account.LongCalls)
	log.Printf("  Short calls: %d", account.ShortCalls)
	log.Printf("  Long puts: %d", account.LongPuts)
	log.Printf("  Short puts: %d", account.ShortPuts)
	log.Printf("  Pending long calls: %d", account.PendingLongCalls)
	log.Printf("  Pending short calls: %d", account.PendingShortCalls)
	log.Printf("  Pending long puts: %d", account.PendingLongPuts)
	log.Printf("  Pending short puts: %d", account.PendingShortPuts)

	return nil
}

func main() {
	// Configuration for the automated trading strategy
	config := StrategyConfig{
		MarketAddress:      sdk.BtcMarket,      // Use the BTC market from SDK constants
		ContractSize:       big.NewInt(1),      // 1 contract per order
		BuyCallPrice:       big.NewInt(1),      // $0.01 (in smallest unit)
		SellCallPrice:      big.NewInt(100000), // $1000 (in smallest unit)
		BuyPutPrice:        big.NewInt(1),      // $0.01 (in smallest unit)
		MaxRetries:         3,
		RetryDelay:         5 * time.Second,
		CycleCheckInterval: 30 * time.Second, // Check for new cycles every 30 seconds
	}

	// Create and run the trading strategy
	strategy, err := NewTradingStrategy(config)
	if err != nil {
		log.Fatalf("Failed to create trading strategy: %v", err)
	}

	// Print initial account information
	if err := strategy.GetUserAccountInfo(); err != nil {
		log.Printf("Warning: failed to get account info: %v", err)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown on interrupt
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Received shutdown signal, stopping strategy...")
		cancel()
	}()

	// Run the strategy
	if err := strategy.Run(ctx); err != nil {
		log.Fatalf("Strategy failed: %v", err)
	}
}
