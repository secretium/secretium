package application

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/messages"
)

// MiddlewareUserAuth checks, if the user is authenticated in the session cookie.
func (a *Application) MiddlewareUserAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Check, if the user is authenticated.
		if !a.Session.Manager.GetBool(r.Context(), "authenticated") {
			slog.Error(
				messages.ErrSessionUserNotAuthenticated,
				"method", r.Method, "status", http.StatusUnauthorized, "path", r.URL.Path,
				"client_ip", r.RemoteAddr,
			)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// Call the next handler.
		next(w, r, params)
	}
}

// MiddlewareUserAuth checks, if the user is authenticated in the session cookie.
func (a *Application) MiddlewareUserAuthWithHTMXRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Check, if the request has a 'HX-Request' header.
		if r.Header.Get("HX-Request") == "" || r.Header.Get("HX-Request") != "true" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Check, if the user is authenticated.
		if !a.Session.Manager.GetBool(r.Context(), "authenticated") {
			slog.Error(
				messages.ErrSessionUserNotAuthenticated,
				"method", r.Method, "status", http.StatusUnauthorized, "path", r.URL.Path,
				"client_ip", r.RemoteAddr,
			)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// Call the next handler.
		next(w, r, params)
	}
}

// MiddlewareHTMXRequest checks, if request is a valid HTMX request.
func (a *Application) MiddlewareHTMXRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Check, if the request has a 'HX-Request' header.
		if r.Header.Get("HX-Request") == "" || r.Header.Get("HX-Request") != "true" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Call the next handler.
		next(w, r, params)
	}
}
