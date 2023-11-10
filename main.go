package main

import (
	"log/slog"
)

func main() {
	// Initialize application.
	app, err := initializeApplication()
	if err != nil {
		slog.Error("failed to initialize application", "details", err.Error())
		return
	}

	// Make sure to close the DB connection when the application exits.
	defer app.Database.Connection.Close()

	// Run application.
	if err := app.Run(); err != nil {
		slog.Error("failed to run application", "details", err.Error())
		return
	}
}
