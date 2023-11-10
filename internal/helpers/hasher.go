package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// ShortenString shortens a string to a specified length using SHA256 hash.
func HashString(s ...string) string {
	// Create a new SHA256 hash object.
	h := sha256.New()

	// Write the input string to the hash object with salt.
	h.Write([]byte(strings.Join(s, "")))

	// Return the first `length` characters of the hash string.
	return hex.EncodeToString(h.Sum(nil))
}
