package helpers

import (
	"log/slog"
	"net/http"
)

// WrapHTTPError wraps HTTP errors.
func WrapHTTPError(w http.ResponseWriter, r *http.Request, status int, err string) {
	// Log error.
	slog.Error(
		err,
		"method", r.Method, "status", status,
		"path", r.URL.Path, "query", r.URL.Query().Encode(),
	)

	// Set response status code.
	switch status {
	case 0:
		// If you don't want to write HTTP headers, set the status to 0.
	case http.StatusUnauthorized, http.StatusForbidden:
		// Set the 'HX-Trigger' header.
		w.Header().Set("HX-Trigger", `{"jwtRemoveFromLocalStorage": ""}`)
		// Redirect to the index page.
		w.Header().Set("HX-Redirect", "/") // HTTP 401 and 403
	case http.StatusNotFound:
		http.NotFound(w, r) // HTTP 404
	default:
		http.Error(w, err, status) // other
	}
}
