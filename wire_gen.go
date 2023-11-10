// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/secretium/secretium/internal/application"
	"github.com/secretium/secretium/internal/attachments"
	"github.com/secretium/secretium/internal/config"
	"github.com/secretium/secretium/internal/database"
)

// Injectors from wire.go:

// initializeApplication provides dependency injection process by the "google/wire" package.
func initializeApplication() (*application.Application, error) {
	attachmentsAttachments := attachments.New()
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	databaseDatabase, err := database.New(configConfig)
	if err != nil {
		return nil, err
	}
	applicationApplication := application.New(attachmentsAttachments, configConfig, databaseDatabase)
	return applicationApplication, nil
}
