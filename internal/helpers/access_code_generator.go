package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateAccessCode generates a random access code of the given size.
func GenerateAccessCode(size int) (string, error) {
	// Create a byte slice with the buffer size.
	buffer := make([]byte, size/2)

	// Generate random bytes and store them in the byte slice.
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a hexadecimal string and return the first 'size' characters.
	return hex.EncodeToString(buffer)[:size], nil
}
