package application

import (
	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/helpers"
)

// router returns a new mux instance with all the routes.
func (a *Application) router() *httprouter.Router {
	// Create a new mux.
	router := httprouter.New()

	// Add a file server from the embedded static files.
	router.NotFound = helpers.StaticFileServerHandler(a.Attachments.StaticFiles)

	/*
		Public routes.
	*/

	// Add a public set of HTML page handlers.
	router.GET("/", a.PageIndexHandler)          // handle the index page
	router.GET("/get/:key", a.PageSecretHandler) // handle the secret page

	// Add a public set of API handlers.
	router.POST("/api/secret/unlock/:key", a.MiddlewareHTMXRequest(a.APIUnlockSecretHandler))                     // handle the unlock secret request to the API
	router.PATCH("/api/secret/expire/:key", a.MiddlewareHTMXRequest(a.APIExpireSecretExpiresAtFieldByKeyHandler)) // handle the expire secret request to the API
	router.POST("/api/user/login", a.MiddlewareHTMXRequest(a.APIUserLoginHandler))                                // handle the user login request to the API

	/*
		Private routes.
	*/

	// Add a public set of HTML page handlers.
	router.GET("/dashboard", a.MiddlewareUserAuth(a.PageDashboardIndexHandler))                                 // handle the user dashboard page
	router.GET("/dashboard/add", a.MiddlewareUserAuthWithHTMXRequest(a.PageDashboardAddSecretHandler))          // handle the dashboard add secret page
	router.GET("/dashboard/share/:key", a.MiddlewareUserAuthWithHTMXRequest(a.PageDashboardShareSecretHandler)) // handle the dashboard share secret page

	// Add a set of API handlers.
	router.POST("/api/secret/add", a.MiddlewareUserAuthWithHTMXRequest(a.APIAddSecretHandler))                                   // handle the add secret request to the API
	router.PATCH("/api/secret/renew/:key", a.MiddlewareUserAuthWithHTMXRequest(a.APIRenewSecretExpiresAtFieldByKeyHandler))      // handle the renew secret request to the API
	router.PATCH("/api/secret/restore/:key", a.MiddlewareUserAuthWithHTMXRequest(a.APIRestoreSecretAccessCodeFieldByKeyHandler)) // handle the restore secret access code request to the API
	router.DELETE("/api/secret/delete/:key", a.MiddlewareUserAuthWithHTMXRequest(a.APIDeleteSecretByKeyHandler))                 // handle the delete secret request to the API
	router.GET("/api/dashboard/secrets/active", a.MiddlewareUserAuthWithHTMXRequest(a.APIDashboardActiveSecretsHandler))         // handle the get active secret request to the API
	router.GET("/api/dashboard/secrets/expired", a.MiddlewareUserAuthWithHTMXRequest(a.APIDashboardExpiredSecretsHandler))       // handle the get expired secret request to the API
	router.GET("/api/user/logout", a.MiddlewareUserAuthWithHTMXRequest(a.APIUserLogoutHandler))                                  // handle the user logout request to the API

	// Add a set of QR code generation handler.
	router.GET("/qr/generate/:key", a.MiddlewareUserAuth(a.QRCodeGenerationHandler)) // handle the request to generate a QR code

	return router
}
