package sdk

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type OrderEngine struct {
	privateKey *ecdsa.PrivateKey
	baseURL    string
	chainID    *big.Int
	httpClient *http.Client
}

type Order struct {
	MarketID   string `json:"marketId"`
	CycleID    string `json:"cycleId"`
	CycleNonce int    `json:"cycleNonce"`
	Side       string `json:"side"`       // "buy" or "sell"
	OptionType string `json:"optionType"` // "call" or "put"
	Size       int    `json:"size"`
	LimitPrice int    `json:"limitPrice"`
}

type OrdersRequest struct {
	Signature string  `json:"signature"`
	Orders    []Order `json:"orders"`
}

type CancelAndCloseRequest struct {
	Signature string `json:"signature"`
	MarketID  int    `json:"marketId"`
	Nonce     int    `json:"nonce"`
}

type OrderParams struct {
	MarketID       string
	CycleID        string
	CycleNonce     int
	Size           int
	BuyLimitPrice  int
	SellLimitPrice int
}

type CancelAndCloseParams struct {
	MarketID int
	Nonce    int
}

// EIP-712 domain and types
var (
	OrderEngineTypedDataDomain = apitypes.TypedDataDomain{
		Name:              "OptServer",
		Version:           "2",
		ChainId:           math.NewHexOrDecimal256(999), // HyperEVM chain ID
		VerifyingContract: "0x0000000000000000000000000000000000000000",
	}

	OrdersTypedDataTypes = apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"Order": {
			{Name: "marketId", Type: "uint256"},
			{Name: "cycleId", Type: "uint256"},
			{Name: "cycleNonce", Type: "uint256"},
			{Name: "optionType", Type: "string"},
			{Name: "side", Type: "string"},
			{Name: "size", Type: "uint256"},
			{Name: "limitPrice", Type: "uint256"},
		},
		"Orders": {
			{Name: "orders", Type: "Order[]"},
		},
	}

	CancelAndCloseTypedDataTypes = apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"CancelAndClose": {
			{Name: "marketId", Type: "uint256"},
			{Name: "nonce", Type: "uint256"},
		},
	}
)

func NewOrderEngine() (*OrderEngine, error) {
	baseURL := os.Getenv("API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // Default API URL
	}

	privateKey, err := GetPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID := big.NewInt(ChainId)
	if os.Getenv("CHAIN_ID") != "" {
		chainIdStr := os.Getenv("CHAIN_ID")
		chainIdInt, err := strconv.ParseInt(chainIdStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid CHAIN_ID: %w", err)
		}
		chainID = big.NewInt(chainIdInt)
	}

	return &OrderEngine{
		privateKey: privateKey,
		address:    address,
		baseURL:    baseURL,
		chainID:    chainID,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

func (oe *OrderEngine) Long(params OrderParams) error {
	orders := []Order{}

	// For long: buy calls if buyLimitPrice > 0
	if params.BuyLimitPrice > 0 {
		orders = append(orders, Order{
			MarketID:   params.MarketID,
			CycleID:    params.CycleID,
			CycleNonce: params.CycleNonce,
			Side:       "buy",
			OptionType: "call",
			Size:       params.Size,
			LimitPrice: params.BuyLimitPrice,
		})
		params.CycleNonce++
	}

	// For long: sell puts if sellLimitPrice > 0
	if params.SellLimitPrice > 0 {
		orders = append(orders, Order{
			MarketID:   params.MarketID,
			CycleID:    params.CycleID,
			CycleNonce: params.CycleNonce,
			Side:       "sell",
			OptionType: "put",
			Size:       params.Size,
			LimitPrice: params.SellLimitPrice,
		})
	}

	return oe.submitOrders(orders)
}

func (oe *OrderEngine) Short(params OrderParams) error {
	orders := []Order{}

	// For short: buy puts if buyLimitPrice > 0
	if params.BuyLimitPrice > 0 {
		orders = append(orders, Order{
			MarketID:   params.MarketID,
			CycleID:    params.CycleID,
			CycleNonce: params.CycleNonce,
			Side:       "buy",
			OptionType: "put",
			Size:       params.Size,
			LimitPrice: params.BuyLimitPrice,
		})
		params.CycleNonce++
	}

	// For short: sell calls if sellLimitPrice > 0
	if params.SellLimitPrice > 0 {
		orders = append(orders, Order{
			MarketID:   params.MarketID,
			CycleID:    params.CycleID,
			CycleNonce: params.CycleNonce,
			Side:       "sell",
			OptionType: "call",
			Size:       params.Size,
			LimitPrice: params.SellLimitPrice,
		})
	}

	return oe.submitOrders(orders)
}

func (oe *OrderEngine) CancelAndClose(params CancelAndCloseParams) error {
	// Create typed data for signing
	typedData := apitypes.TypedData{
		Types:       CancelAndCloseTypedDataTypes,
		PrimaryType: "CancelAndClose",
		Domain:      OrderEngineTypedDataDomain,
		Message: apitypes.TypedDataMessage{
			"marketId": params.MarketID,
			"nonce":    params.Nonce,
		},
	}

	// Sign the typed data
	signature, err := oe.signTypedData(typedData)
	if err != nil {
		return fmt.Errorf("failed to sign cancel and close request: %w", err)
	}

	// Create request
	request := CancelAndCloseRequest{
		Signature: signature,
		MarketID:  params.MarketID,
		Nonce:     params.Nonce,
	}

	// Submit to API
	return oe.submitCancelAndClose(request)
}

func (oe *OrderEngine) submitOrders(orders []Order) error {
	if len(orders) == 0 {
		return fmt.Errorf("no orders to submit")
	}

	// Convert orders to typed data message format
	ordersMessage := make([]interface{}, len(orders))
	for i, order := range orders {
		marketID, err := strconv.ParseInt(order.MarketID, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid marketId: %w", err)
		}
		cycleID, err := strconv.ParseInt(order.CycleID, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid cycleId: %w", err)
		}

		ordersMessage[i] = map[string]interface{}{
			"marketId":   marketID,
			"cycleId":    cycleID,
			"cycleNonce": order.CycleNonce,
			"optionType": order.OptionType,
			"side":       order.Side,
			"size":       order.Size,
			"limitPrice": order.LimitPrice,
		}
	}

	// Create typed data for signing
	typedData := apitypes.TypedData{
		Types:       OrdersTypedDataTypes,
		PrimaryType: "Orders",
		Domain:      OrderEngineTypedDataDomain,
		Message: apitypes.TypedDataMessage{
			"orders": ordersMessage,
		},
	}

	// Sign the typed data
	signature, err := oe.signTypedData(typedData)
	if err != nil {
		return fmt.Errorf("failed to sign orders: %w", err)
	}

	// Create request
	request := OrdersRequest{
		Signature: signature,
		Orders:    orders,
	}

	// Submit to API
	return oe.submitOrdersToAPI(request)
}

// signTypedData signs typed data using EIP-712
func (oe *OrderEngine) signTypedData(typedData apitypes.TypedData) (string, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return "", fmt.Errorf("failed to hash domain: %w", err)
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return "", fmt.Errorf("failed to hash message: %w", err)
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)

	signature, err := crypto.Sign(hash.Bytes(), oe.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}

	// Adjust recovery ID
	signature[64] += 27

	return fmt.Sprintf("0x%x", signature), nil
}

func (oe *OrderEngine) submitOrdersToAPI(request OrdersRequest) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", oe.baseURL+"/order-book/orders", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := oe.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit orders: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %d %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	return nil
}

func (oe *OrderEngine) submitCancelAndClose(request CancelAndCloseRequest) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", oe.baseURL+"/order-book/cancel-and-close", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := oe.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit cancel and close: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %d %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	return nil
}
