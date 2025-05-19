package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateID generates a random ID of specified byte length
func GenerateID(byteLength int) (string, error) {
	// Generate random bytes
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	
	// Encode to base64 (URL-safe, no padding) for compactness
	id := base64.RawURLEncoding.EncodeToString(bytes)

	// Truncate to approximate length (base64 expands by ~4/3)
	return id[:byteLength*4/3], nil
}
