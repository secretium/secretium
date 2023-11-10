package helpers

import (
	"os"
)

// Getenv returns the value of the environment variable associated with the given key.
func Getenv(key, fallback string) string {
	// Check if the environment variable exists for the given key
	value, ok := os.LookupEnv(key)
	if ok {
		// If the environment variable exists, return its value
		return value
	}

	// If the environment variable does not exist, return the fallback value
	return fallback
}
