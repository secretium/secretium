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
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Set the Content-Type header of the response to "image/png".
	w.Header().Set("Content-Type", "image/png")

	// Write the generated QR code image to the response body.
	if err := png.Encode(w, qrCode); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
