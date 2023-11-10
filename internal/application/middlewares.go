package application

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/helpers"
	"github.com/secretium/secretium/internal/messages"
)

// MiddlewareUserAuth checks, if the request has a 'Authorization' header.
func MiddlewareUserAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Check, if the request has a 'Authorization' header with 'Bearer' prefix.
		if !helpers.IsRequestWithAuthorizationHeader(r) {
			// Send a 403 forbidden status to logs and redirect to the index page.
			helpers.WrapHTTPError(w, r, http.StatusForbidden, messages.ErrJWTHeaderNotValid)
			return
		}

		// Call the next handler.
		next(w, r, params)
	}
}
