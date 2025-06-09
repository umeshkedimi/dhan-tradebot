package dhan

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
    "strconv"

    "github.com/joho/godotenv"
    "github.com/go-resty/resty/v2"
)

type DhanClient struct {
	ClientID string
	AccessToken string
	Target        float64
    Stoploss      float64
    TrailStart    float64
    TrailStep     float64
    CurrentTrail  float64
}

func InitDhanClient() *DhanClient {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	clientID := os.Getenv("DHAN_CLIENT_ID")
	accessToken := os.Getenv("DHAN_ACCESS_TOKEN")
	target, _ := strconv.ParseFloat(os.Getenv("PNL_TARGET"), 64)
	stoploss, _ := strconv.ParseFloat(os.Getenv("PNL_STOPLOSS"), 64)
	trailStart, _ := strconv.ParseFloat(os.Getenv("TRAIL_START"), 64)
	trailStep, _ := strconv.ParseFloat(os.Getenv("TRAIL_STEP"), 64)

	if clientID == "" || accessToken == "" {
		log.Fatal("❌ Dhan credentials not found in environment variables")
	}

	return &DhanClient{
		ClientID:    clientID,
		AccessToken: accessToken,
		Target:       target,
    	Stoploss:     stoploss,
    	TrailStart:   trailStart,
    	TrailStep:    trailStep,
    	CurrentTrail: stoploss, // initialize with base SL
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

func (dc *DhanClient) ShouldExit(pnl float64) (bool, string) {
    // Check normal target
    if pnl >= dc.Target {
        return true, "🎯 Target hit"
    }

    // Check trailing SL only if enabled
    if dc.TrailStart > 0 {
        if pnl >= dc.TrailStart && pnl-dc.CurrentTrail >= dc.TrailStep {
            dc.CurrentTrail += dc.TrailStep
            log.Printf("🔁 Trailing SL moved up to ₹%.2f", dc.CurrentTrail)
        }
        if pnl <= dc.CurrentTrail {
            return true, fmt.Sprintf("🛑 Trailing SL hit (₹%.2f)", dc.CurrentTrail)
        }
    } else {
        // Fallback to normal SL
        if pnl <= dc.Stoploss {
            return true, "🛑 Stop-loss hit"
        }
    }

    return false, ""
}

