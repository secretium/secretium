package session

import (
	"net/http"
	"testing"
	"time"

	"github.com/secretium/secretium/internal/config"
)

func TestNew(t *testing.T) {
	// Test setting session options
	m := New(&config.Config{
		Domain:       "example.com",
		DomainSchema: "https",
	})

	if m.Manager.Lifetime != 1*time.Hour {
		t.Errorf("unexpected session lifetime, got: %v, want: %v", m.Manager.Lifetime, 1*time.Hour)
	}

	if m.Manager.IdleTimeout != 30*time.Minute {
		t.Errorf("unexpected session idle timeout, got: %v, want: %v", m.Manager.IdleTimeout, 30*time.Minute)
	}

	if m.Manager.Cookie.Name != "secretium_session_id" {
		t.Errorf("unexpected session cookie name, got: %v, want: %v", m.Manager.Cookie.Name, "secretium_session_id")
	}

	if m.Manager.Cookie.Domain != "example.com" {
		t.Errorf("unexpected session cookie domain, got: %v, want: %v", m.Manager.Cookie.Domain, "example.com")
	}

	if m.Manager.Cookie.Secure != true {
		t.Errorf("unexpected session cookie secure flag, got: %v, want: %v", m.Manager.Cookie.Secure, true)
	}

	if m.Manager.Cookie.SameSite != http.SameSiteStrictMode {
		t.Errorf("unexpected session cookie SameSite attribute, got: %v, want: %v", m.Manager.Cookie.SameSite, http.SameSiteStrictMode)
	}

	if m.Manager.Cookie.HttpOnly != true {
		t.Errorf("unexpected session cookie HttpOnly flag, got: %v, want: %v", m.Manager.Cookie.HttpOnly, true)
	}

	if m.Manager.Cookie.Path != "/" {
		t.Errorf("unexpected session cookie path, got: %v, want: %v", m.Manager.Cookie.Path, "/")
	}
}
