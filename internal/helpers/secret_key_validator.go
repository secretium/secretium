package helpers

import (
	"fmt"

	"github.com/secretium/secretium/internal/messages"
)

// IsSecretKeyValid returns nil if the given secret key is valid.
func IsSecretKeyValid(key string, length int) error {
	// Check, if the given secret key is empty.
	if key == "" {
		return fmt.Errorf(messages.ErrSecretKeyEmptyOrNotFound)
	}

	// Check, if the given secret key is of the correct length.
	if len(key) != length {
		return fmt.Errorf(messages.ErrSecretKeyLengthNotValid, length)
	}

	return nil
}
