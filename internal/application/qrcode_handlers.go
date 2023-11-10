package application

import (
	"fmt"
	"image/png"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/secretium/secretium/internal/helpers"
)

// QRCodeGenerationHandler generates a QR code image from a given text (GET).
func (a *Application) QRCodeGenerationHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get key from the URL.
	key := params.ByName("key")

	// Check, if the current URL has a 'key' parameter with a valid secret key.
	if err := helpers.IsSecretKeyValid(key, 16); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Build the URL with the given key.
	shareURL := url.URL{
		Scheme: a.Config.DomainSchema,
		Host:   a.Config.Domain,
		Path:   fmt.Sprintf("get/%s", key),
	}

	// Generate a QR code image with the given text and a size of 156 pixels.
	qrCode, err := helpers.GenerateQRCode(shareURL.String(), 156)
	if err != nil {
		// If there was an error generating the QR code, wrap the error and return
		// an HTTP internal server error response with the error message.
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Set the Content-Type header of the response to "image/png".
	w.Header().Set("Content-Type", "image/png")

	// Write the generated QR code image to the response body.
	if err := png.Encode(w, qrCode); err != nil {
		helpers.WrapHTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
