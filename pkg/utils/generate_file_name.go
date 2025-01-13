package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateFileName(bytes int) string {
	randomBytes := make([]byte, bytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("failed to generate file name: %v", err)
	}
	return hex.EncodeToString(randomBytes)
}
