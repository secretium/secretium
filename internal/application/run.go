package application

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

// Run runs the application.
func (a *Application) Run() error {
	// Create a new DB schema, if it does not exist.
	if err := a.Database.Migrate("sql_queries/init.sql"); err != nil {
		return err
	}

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.Config.Server.Port),
		ReadTimeout:  time.Duration(a.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(a.Config.Server.WriteTimeout) * time.Second,
		Handler:      a.Session.Manager.LoadAndSave(a.router()), // use the HttpRouter instance with session manager
	}

	// Get the URL of the server.
	u := url.URL{
		Scheme: a.Config.DomainSchema,
		Host:   a.Config.Domain,
	}

	// Log the start of the application.
	slog.Info(
		"running application",
		"domain", u.String(),
		"port", a.Config.Server.Port,
		"read_timeout", a.Config.Server.ReadTimeout,
		"write_timeout", a.Config.Server.WriteTimeout,
	)

	// Start the server in a separate goroutine.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "error", err)
		}
	}()

	// Create a channel to listen for OS signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Listen for the OS signal to gracefully shutdown the server.
	<-stop

	// Create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully.
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to gracefully shutdown server", "error", err)
	}

	slog.Info("server gracefully stopped")

	return nil
}
