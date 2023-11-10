package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// ShortenString shortens a string to a specified length using SHA256 hash.
func ShortenString(s string, length int) string {
	// Create a new SHA256 hash object.
	h := sha256.New()

	// Write the input string to the hash object.
	h.Write([]byte(s))

	// Convert the hash to a hexadecimal string.
	hashString := hex.EncodeToString(h.Sum(nil))

	// Return the first `length` characters of the hash string.
	return hashString[:length]
}
