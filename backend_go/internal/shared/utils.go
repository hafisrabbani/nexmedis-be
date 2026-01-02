package shared

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func GenerateApiKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func HashApiKey(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}

func EncryptEmail(email string) []byte {
	return []byte(email + ":" + os.Getenv("APP_SECRET"))
}
