package main

import (
	"log"
	"math/big"

	"github.com/Convexity-Research/optfun-sdk/contracts"
	"github.com/Convexity-Research/optfun-sdk/sdk"
)

func main() {
	log.Println("Testing OptFun SDK connection...")

	// Load SDK
	client, fromAddress, auth, err := sdk.LoadSDK()
	if err != nil {
		log.Fatalf("Failed to load SDK: %v", err)
	}

	if auth == nil {
		log.Println("No private key provided - running in read-only mode")
	} else {
		log.Printf("Private key loaded - running in full mode")
		log.Printf("Trading address: %s", fromAddress.Hex())
	}

	// Test connection to market
	market, err := contracts.NewMarket(sdk.BtcMarket, client)
	if err != nil {
		log.Fatalf("Failed to connect to market: %v", err)
	}

	log.Printf("Connected to market at: %s", sdk.BtcMarket.Hex())

	// Get market information
	name, err := market.GetName(nil)
	if err != nil {
		log.Printf("Warning: failed to get market name: %v", err)
	} else {
		log.Printf("Market name: %s", name)
	}

	// Get current cycle
	activeCycle, err := market.GetActiveCycle(nil)
	if err != nil {
		log.Fatalf("Failed to get active cycle: %v", err)
	}
	log.Printf("Current active cycle: %s", activeCycle.String())

	// Get oracle price
	oraclePrice, err := market.GetOraclePrice(nil)
	if err != nil {
		log.Printf("Warning: failed to get oracle price: %v", err)
	} else {
		log.Printf("Current oracle price: $%d", oraclePrice)
	}

	// Check if market is paused
	paused, err := market.Paused(nil)
	if err != nil {
		log.Printf("Warning: failed to check if market is paused: %v", err)
	} else {
		log.Printf("Market paused: %t", paused)
	}

	// Get user account information if private key is provided
	if auth != nil {
		account, err := market.GetUserAccounts(nil, fromAddress)
		if err != nil {
			log.Printf("Warning: failed to get user account: %v", err)
		} else {
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

			// Check if user is whitelisted
			whitelisted, err := market.GetWhitelist(nil, fromAddress)
			if err != nil {
				log.Printf("Warning: failed to check whitelist status: %v", err)
			} else {
				log.Printf("Whitelisted: %t", whitelisted)
			}
		}
	}

	// Test price formatting
	log.Println("\nPrice Configuration Test:")
	log.Printf("Buy call price ($0.01): %s", big.NewInt(1).String())
	log.Printf("Sell call price ($1000): %s", big.NewInt(100000).String())
	log.Printf("Buy put price ($0.01): %s", big.NewInt(1).String())

	log.Println("\n✅ SDK connection test completed successfully!")
	log.Println("You can now run the automated strategy with: go run automated_strategy.go")
}
