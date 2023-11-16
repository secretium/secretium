package helpers

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/secretium/secretium/internal/constants"
	"github.com/secretium/secretium/internal/messages"
)

// ConfigValidation validates the configuration settings.
//
// This function checks the validity of various configuration settings such as the secret key,
// master username, master password, domain URL, domain HTTP schema, and server timezone.
// It returns an error if any of the configuration settings are invalid.
//
// Returns:
// - error: An error indicating the invalid configuration setting, or nil if all settings are valid.
func ConfigValidation() error {
	// Check SECRET_KEY.
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return errors.New(messages.ErrConfigSecretKeyEmpty)
	}
	if len(secretKey) < constants.ConstConfigSecretKeyMinLength {
		return fmt.Errorf(messages.ErrConfigSecretKeyLengthNotValid, constants.ConstConfigSecretKeyMinLength)
	}

	// Check MASTER_USERNAME.
	masterUsername := os.Getenv("MASTER_USERNAME")
	if masterUsername == "" {
		return errors.New(messages.ErrConfigMasterUsernameEmpty)
	}
	if len(masterUsername) < constants.ConstConfigMasterUsernameMinLength ||
		len(masterUsername) > constants.ConstConfigMasterUsernameMaxLength {
		return fmt.Errorf(
			messages.ErrConfigMasterUsernameLengthNotValid,
			constants.ConstConfigMasterUsernameMinLength,
			constants.ConstConfigMasterUsernameMaxLength,
		)
	}

	// Check MASTER_PASSWORD.
	masterPassword := os.Getenv("MASTER_PASSWORD")
	if masterPassword == "" {
		return errors.New(messages.ErrConfigMasterPasswordEmpty)
	}
	if len(masterPassword) < constants.ConstConfigMasterPasswordMinLength ||
		len(masterPassword) > constants.ConstConfigMasterPasswordMaxLength {
		return fmt.Errorf(
			messages.ErrConfigMasterPasswordLengthNotValid,
			constants.ConstConfigMasterPasswordMinLength,
			constants.ConstConfigMasterPasswordMaxLength,
		)
	}

	// Check DOMAIN.
	domain := os.Getenv("DOMAIN")
	if domain != "" {
		if err := IsValidURL(domain); err != nil {
			return err
		}
	}

	// Check DOMAIN_SCHEMA.
	domainSchema := os.Getenv("DOMAIN_SCHEMA")
	if domainSchema != "" {
		if !slices.Contains([]string{"https", "http"}, domainSchema) {
			return errors.New(messages.ErrConfigDomainSchemaNotValid)
		}
	}

	// Check SERVER_TIMEZONE.
	serverTimezone := Getenv("SERVER_TIMEZONE", constants.ConstConfigServerTimezone)
	_, err := time.LoadLocation(serverTimezone)
	if err != nil {
		return errors.New(messages.ErrConfigServerTimezoneNotValid)
	}

	return nil
}
