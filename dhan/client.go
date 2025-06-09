package dhan

import (
    "fmt"
    "log"
    "os"
    "encoding/json"

    "github.com/joho/godotenv"
    "github.com/go-resty/resty/v2"
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

func (dc *DhanClient) GetPnL() (float64, error) {
    client := resty.New()

    resp, err := client.R().
        SetHeader("access-token", dc.AccessToken).
        SetHeader("client-id", dc.ClientID).
        Get("https://api.dhan.co/positions")

    if err != nil {
        return 0, fmt.Errorf("❌ Error fetching positions: %v", err)
    }

    var positions []struct {
        RealisedProfit   float64 `json:"realisedProfit"`
        UnrealisedProfit float64 `json:"unrealisedProfit"`
    }

    if err := json.Unmarshal(resp.Body(), &positions); err != nil {
        return 0, fmt.Errorf("❌ JSON parse error: %v", err)
    }

    totalPnL := 0.0
    for _, pos := range positions {
        totalPnL += pos.RealisedProfit + pos.UnrealisedProfit
    }

    return totalPnL, nil
}
