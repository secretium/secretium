package helpers

import (
	"errors"
	"net/url"

	"github.com/secretium/secretium/internal/messages"
)

// IsValidURL returns nil if the given URL is valid.
func IsValidURL(uri string) error {
	// Check if the domain has a valid URL.
	if _, err := url.Parse(uri); err != nil {
		return errors.New(messages.ErrConfigDomainNotValid)
	}

	return nil
}
