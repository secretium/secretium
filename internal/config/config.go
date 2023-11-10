package config

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
	_ "time/tzdata"

	"github.com/secretium/secretium/internal/constants"
	"github.com/secretium/secretium/internal/helpers"
	"github.com/secretium/secretium/internal/messages"
)

// Config contains secret key, master password, domain and server configuration.
type Config struct {
	SecretKey, MasterUsername, MasterPassword, Domain, DomainSchema string
	Server                                                          *server
}

// Server contains port, read and write timeout.
type server struct {
	Port, ReadTimeout, WriteTimeout int
}

// New returns a new instance of Config after validation.
func New() (*Config, error) {
	// Validate secret key.
	if os.Getenv("SECRET_KEY") == "" {
		return nil, errors.New(messages.ErrConfigSecretKeyEmpty)
	}

	// Validate length of the secret key.
	if len(os.Getenv("SECRET_KEY")) < constants.ConstConfigSecretKeyMinLength {
		return nil, fmt.Errorf(messages.ErrConfigSecretKeyLengthNotValid, constants.ConstConfigSecretKeyMinLength)
	}

	// Validate master username.
	if os.Getenv("MASTER_USERNAME") == "" {
		return nil, errors.New(messages.ErrConfigMasterUsernameEmpty)
	}

	// Validate length of the master username.
	if len(os.Getenv("MASTER_USERNAME")) < constants.ConstConfigMasterUsernameMinLength ||
		len(os.Getenv("MASTER_USERNAME")) > constants.ConstConfigMasterUsernameMaxLength {
		return nil, fmt.Errorf(
			messages.ErrConfigMasterUsernameLengthNotValid,
			constants.ConstConfigMasterUsernameMinLength,
			constants.ConstConfigMasterUsernameMaxLength,
		)
	}

	// Validate master password.
	if os.Getenv("MASTER_PASSWORD") == "" {
		return nil, errors.New(messages.ErrConfigMasterPasswordEmpty)
	}

	// Validate length of the master password.
	if len(os.Getenv("MASTER_PASSWORD")) < constants.ConstConfigMasterPasswordMinLength ||
		len(os.Getenv("MASTER_PASSWORD")) > constants.ConstConfigMasterPasswordMaxLength {
		return nil, fmt.Errorf(
			messages.ErrConfigMasterPasswordLengthNotValid,
			constants.ConstConfigMasterPasswordMinLength,
			constants.ConstConfigMasterPasswordMaxLength,
		)
	}

	// Validate domain URL.
	if err := helpers.IsValidURL(os.Getenv("DOMAIN")); err != nil {
		return nil, err
	}

	// Check, if the domain HTTP schema is present.
	if len(os.Getenv("DOMAIN_SCHEMA")) > 0 {
		// Validate domain HTTP schema.
		if !slices.Contains([]string{"https", "http"}, os.Getenv("DOMAIN_SCHEMA")) {
			return nil, errors.New(messages.ErrConfigDomainSchemaNotValid)
		}
	}

	// Validate server port.
	port, err := strconv.Atoi(helpers.Getenv("SERVER_PORT", constants.ConstConfigServerPort))
	if err != nil {
		return nil, errors.New(messages.ErrConfigServerPortNotValid)
	}

	// Validate server read timeout.
	readTimeout, err := strconv.Atoi(helpers.Getenv("SERVER_READ_TIMEOUT", constants.ConstConfigServerReadTimeout))
	if err != nil {
		return nil, errors.New(messages.ErrConfigServerReadTimeoutNotValid)
	}

	// Validate server write timeout.
	writeTimeout, err := strconv.Atoi(helpers.Getenv("SERVER_WRITE_TIMEOUT", constants.ConstConfigServerWriteTimeout))
	if err != nil {
		return nil, errors.New(messages.ErrConfigServerWriteTimeoutNotValid)
	}

	// Validate server timezone.
	if _, err := time.LoadLocation(helpers.Getenv("SERVER_TIMEZONE", constants.ConstConfigServerTimezone)); err != nil {
		return nil, errors.New(messages.ErrConfigServerTimezoneNotValid)
	}

	return &Config{
		SecretKey:      os.Getenv("SECRET_KEY"),
		MasterUsername: os.Getenv("MASTER_USERNAME"),
		MasterPassword: os.Getenv("MASTER_PASSWORD"),
		Domain:         os.Getenv("DOMAIN"),
		DomainSchema:   helpers.Getenv("DOMAIN_SCHEMA", constants.ConstConfigDomainSchema),
		Server: &server{
			Port:         port,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}, nil
}
