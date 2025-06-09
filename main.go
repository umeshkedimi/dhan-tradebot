package main

import (
	"github.com/umeshkedimi/dhan-tradebot/dhan"
	"github.com/umeshkedimi/dhan-tradebot/utils"
)

func main() {
	logger := utils.NewLogger("algo.log")
	logger.Println("🚀 Algo Starting...")
    
	dhanClient := dhan.InitDhanClient()
	logger.Println("✅ Dhan client initialized.")
    logger.Printf("Client ID: %s\n", dhanClient.ClientID)
}