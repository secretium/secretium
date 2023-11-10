//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/secretium/secretium/internal/application"
	"github.com/secretium/secretium/internal/attachments"
	"github.com/secretium/secretium/internal/config"
	"github.com/secretium/secretium/internal/database"
)

// initializeApplication provides dependency injection process by the "google/wire" package.
func initializeApplication() (*application.Application, error) {
	panic(wire.Build(attachments.New, config.New, database.New, application.New))
}
