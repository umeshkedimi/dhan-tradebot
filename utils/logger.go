package utils

import (
	"io"
	"log"
	"os"
)

// Create and returns a logger that writes to both console and file
func NewLogger(logFile string) *log.Logger {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("❌ Failed to open log file: %v", err)
	}

	// MultiWriter sends output to both file and terminal
    multiWriter := io.MultiWriter(os.Stdout, file)

	// Create and return the logger
    return log.New(multiWriter, "LOG: ", log.LstdFlags|log.Lshortfile)
}