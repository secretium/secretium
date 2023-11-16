package config

import (
	"errors"
	"os"
	"strconv"
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
	// Validate config.
	if err := helpers.ConfigValidation(); err != nil {
		return nil, err
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

	return &Config{
		SecretKey:      os.Getenv("SECRET_KEY"),
		MasterUsername: os.Getenv("MASTER_USERNAME"),
		MasterPassword: os.Getenv("MASTER_PASSWORD"),
		Domain:         helpers.Getenv("DOMAIN", constants.ConstConfigDomain),
		DomainSchema:   helpers.Getenv("DOMAIN_SCHEMA", constants.ConstConfigDomainSchema),
		Server: &server{
			Port:         port,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}, nil
}
