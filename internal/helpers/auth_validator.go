package helpers

import (
	"net/http"
	"strings"
)

// IsRequestWithAuthorizationHeader checks, if the request has a 'Authorization' header with 'Bearer' prefix.
func IsRequestWithAuthorizationHeader(r *http.Request) bool {
	// Get 'Authorization' header.
	header := strings.Split(r.Header.Get("Authorization"), " ")

	// Check, if the request has a 'Authorization' header.
	if len(header) == 0 {
		return false
	}

	// Check, if the request has a 'Authorization' header with 'Bearer' prefix.
	if len(header) > 0 && header[0] != "Bearer" {
		return false
	}

	return true
}
