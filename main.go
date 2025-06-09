package main

import (
	"github.com/umeshkedimi/dhan-tradebot/dhan"
	"github.com/umeshkedimi/dhan-tradebot/utils"
	"github.com/umeshkedimi/dhan-tradebot/telegram"
	"os"
	"time"
	"log"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	_ = godotenv.Load(".env")
	logger := utils.NewLogger("algo.log")
	logger.Println("🚀 Algo Starting...")

	dhanClient := dhan.InitDhanClient()
	logger.Printf("✅ Dhan client initialized: %s", dhanClient.ClientID)

	telegram.StartTelegramListener(dhanClient)
	logger.Println("📡 Telegram bot listener activated.")

	// Get chat ID for exit alert (optional)
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	// Real-time PnL monitoring loop
	for {
		pnl, err := dhanClient.GetPnL()
		if err != nil {
			logger.Printf("❌ Error fetching PnL: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Printf("📊 Current PnL: ₹%.2f", pnl)

		exit, reason := dhanClient.ShouldExit(pnl)
		if exit {
			logger.Printf("🚨 Exit Triggered: %s", reason)

			// Convert chatID to int64 before sending Telegram message
			var cid int64
			_, err := fmt.Sscanf(chatID, "%d", &cid)
			if err == nil {
				telegram.SendMessage(cid, fmt.Sprintf("🚨 Exit: %s", reason))
			} else {
				log.Printf("⚠️ Could not parse chat ID: %v", err)
			}

			break
		}

		time.Sleep(1 * time.Second) // Check PnL every 5 seconds
	}
}
