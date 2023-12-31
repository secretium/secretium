package application

import (
	"github.com/secretium/secretium/internal/attachments"
	"github.com/secretium/secretium/internal/config"
	"github.com/secretium/secretium/internal/database"
	"github.com/secretium/secretium/internal/session"
)

// Application contains DB connection and other dependencies for application.
type Application struct {
	Attachments *attachments.Attachments
	Config      *config.Config
	Database    *database.Database
	Session     *session.Session
}

// New returns a new instance of Application.
func New(a *attachments.Attachments, c *config.Config, d *database.Database, s *session.Session) *Application {
	return &Application{
		Attachments: a,
		Config:      c,
		Database:    d,
		Session:     s,
	}
}
