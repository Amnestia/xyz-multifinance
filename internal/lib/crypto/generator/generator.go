package generator

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
)

func GenerateAPIKey() string {
	b := make([]byte, 64)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func GenerateClientID() string {
	return uuid.New().String()
}
