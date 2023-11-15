package helpers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// WrapHTTPError wraps HTTP errors.
func WrapHTTPError(w http.ResponseWriter, r *http.Request, status int, errTemplate templ.Component, errMsg string) {
	// Log error.
	slog.Error(errMsg, "method", r.Method, "status", status, "path", r.URL.Path)

	// Set response status code.
	switch status {
	case 0:
		// If you don't want to write HTTP headers, set the status to 0.
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden:
		w.Header().Set("HX-Retarget", "#errors")
		errTemplate.Render(r.Context(), w)
	case http.StatusNotFound:
		http.NotFound(w, r) // HTTP 404
	default:
		http.Error(w, errMsg, status) // other
	}
}
