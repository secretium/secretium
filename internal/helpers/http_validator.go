package helpers

import (
	"net/http"
)

// IsRequestValid returns nil if the given HTTP request is valid.
func IsRequestValid(r *http.Request) bool {
	// Check, if the request has a 'HX-Request' header.
	if r.Header.Get("HX-Request") == "" || r.Header.Get("HX-Request") != "true" {
		return false
	}

	return true
}
