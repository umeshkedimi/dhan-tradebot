package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"log"
	"time"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var lastUpdateID int

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	Message   struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// Load .env and fetch token
func getBotToken() string {
	_ = godotenv.Load(".env")
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}

func initLastUpdateID() {
    botToken := getBotToken()
    url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", botToken)

    resp, err := http.Get(url)
    if err != nil {
        log.Printf("❌ Error initializing update ID: %v", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    var result struct {
        OK     bool              `json:"ok"`
        Result []TelegramUpdate  `json:"result"`
    }

    if err := json.Unmarshal(body, &result); err != nil {
        log.Printf("❌ JSON error initializing update ID: %v", err)
        return
    }

    if len(result.Result) > 0 {
        lastUpdateID = result.Result[len(result.Result)-1].UpdateID
        log.Printf("🔁 Skipping old messages. Starting from update_id: %d", lastUpdateID+1)
    }
}


func getTelegramUpdates() ([]TelegramUpdate, error) {
	botToken := getBotToken()
	offset := lastUpdateID + 1
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", botToken, offset)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result struct {
		Ok     bool              `json:"ok"`
		Result []TelegramUpdate `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

func SendMessage(chatID int64, text string) {
	botToken := getBotToken()
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	payload := fmt.Sprintf("chat_id=%d&text=%s", chatID, text)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(payload))
	if err != nil {
		log.Printf("❌ Error sending message: %v", err)
		return
	}

	defer resp.Body.Close()
}

// This will run in a goroutine
func StartTelegramListener() {
	initLastUpdateID() // Skips old messages on startup
    go func() {
        log.Println("📡 Telegram listener started...")
        for {
            updates, err := getTelegramUpdates()
            if err != nil {
                log.Printf("❌ Error getting updates: %v", err)
                time.Sleep(5 * time.Second)
                continue
            }

            for _, update := range updates {
                lastUpdateID = update.UpdateID

                text := strings.ToLower(update.Message.Text)
                chatID := update.Message.Chat.ID

                log.Printf("📨 Received Telegram message: %s", text)

                switch text {
                case "buy":
                    SendMessage(chatID, "🟢 Buy command received.")
                case "sell":
                    SendMessage(chatID, "🔴 Sell command received.")
                case "pnls":
                    SendMessage(chatID, "💰 PnL: ₹1234.56 (dummy)")
                default:
                    SendMessage(chatID, "🤖 Unknown command. Try: buy, sell, pnls")
                }
            }

            time.Sleep(2 * time.Second)
        }
    }()
}