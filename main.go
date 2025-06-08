package main

import (
	"github.com/umeshkedimi/dhan-tradebot/utils"
)

func main() {
	logger := utils.NewLogger("algo.log")
    logger.Println("🚀 Hello Trader! Algo with custom logger started.")
    logger.Println("✅ Logger working fine. Time to trade!")
}