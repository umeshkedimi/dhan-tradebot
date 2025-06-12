package utils

import (
	"os"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tmp := "test.log"
	defer os.Remove(tmp)
	logger := NewLogger(tmp)
	if logger == nil {
		t.Fatal("expected logger")
	}
	msg := "hello"
	logger.Println(msg)
	data, err := os.ReadFile(tmp)
	if err != nil {
		t.Fatalf("read log file error: %v", err)
	}
	if !strings.Contains(string(data), msg) {
		t.Fatalf("log file does not contain message")
	}
}
