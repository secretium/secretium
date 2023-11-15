package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/secretium/secretium/internal/config"
)

// Session contains session manager.
type Session struct {
	Manager *scs.SessionManager
}

// New creates a new session manager.
func New(c *config.Config) *Session {
	// Create a new session manager.
	m := scs.New()

	// Set the session options.
	m.Lifetime = 1 * time.Hour                  // set the lifetime of the session to 1 hour
	m.IdleTimeout = 30 * time.Minute            // set the idle timeout of the session to 30 minutes
	m.Cookie.Name = "secretium_session_id"      // set the name of the session cookie
	m.Cookie.Domain = c.Domain                  // set the domain of the session cookie to the value specified in the config
	m.Cookie.Secure = c.DomainSchema == "https" // set the secure flag of the session cookie based on the domain schema in the config
	m.Cookie.SameSite = http.SameSiteStrictMode // set the SameSite attribute of the session cookie to strict mode
	m.Cookie.HttpOnly = true                    // set the HttpOnly flag of the session cookie to true
	m.Cookie.Path = "/"                         // set the path of the session cookie to '/'

	return &Session{
		Manager: m,
	}
}
