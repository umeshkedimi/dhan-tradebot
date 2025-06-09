package dhan

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DhanClient struct {
	ClientID string
	AccessToken string
}

func InitDhanClient() *DhanClient {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	clientID := os.Getenv("DHAN_CLIENT_ID")
	accessToken := os.Getenv("DHAN_ACCESS_TOKEN")

	if clientID == "" || accessToken == "" {
		log.Fatal("❌ Dhan credentials not found in environment variables")
	}

	return &DhanClient{
		ClientID:    clientID,
		AccessToken: accessToken,
}
}

// PlaceOrder simulates placing an order with Dhan (for now)
func (dc *DhanClient) PlaceOrder(tradeType string) string {
    var action string
    if tradeType == "buy" {
        action = "🟢 Placing BUY order..."
    } else if tradeType == "sell" {
        action = "🔴 Placing SELL order..."
    } else {
        return "❌ Invalid order type!"
    }

    // Simulate order logic — real Dhan API will go here
    fmt.Printf("%s [ClientID: %s, AccessToken: %s]\n", action, dc.ClientID, dc.AccessToken)

    return fmt.Sprintf("✅ Order placed: %s (simulated)", tradeType)
}
