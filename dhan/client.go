package dhan

import (
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