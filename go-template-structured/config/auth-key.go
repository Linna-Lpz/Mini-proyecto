package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateRandomKey() string {
	bytes := make([]byte, 32) // 256 bits
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Error generating random key:", err)
	}
	
	return base64.URLEncoding.EncodeToString(bytes)
}