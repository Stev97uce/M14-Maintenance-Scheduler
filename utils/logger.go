package utils

import (
	"log"
	"time"
)

func LogInfo(msg string) {
	log.Printf("ℹ️ [%s] %s", time.Now().Format(time.RFC3339), msg)
}

func LogError(msg string, err error) {
	log.Printf("❌ [%s] %s: %v", time.Now().Format(time.RFC3339), msg, err)
}
