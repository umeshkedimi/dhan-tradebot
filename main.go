package main

import (
	"github.com/umeshkedimi/dhan-tradebot/dhan"
	"github.com/umeshkedimi/dhan-tradebot/utils"
	"github.com/umeshkedimi/dhan-tradebot/telegram"
    "time"
)

func main() {
	logger := utils.NewLogger("algo.log")
	logger.Println("🚀 Algo Starting...")
    
	dhanClient := dhan.InitDhanClient()
	logger.Printf("✅ Dhan client initialized: %s", dhanClient.ClientID)

    telegram.StartTelegramListener()
    logger.Println("📡 Telegram bot listener activated.")

    // Keep main.go running
    for {
        time.Sleep(1 * time.Second)
    }
}