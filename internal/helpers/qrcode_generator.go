package helpers

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// GenerateQRCode generates a QR code image from a given text and size.
func GenerateQRCode(text string, size int) (barcode.Barcode, error) {
	// Create the QR code.
	qrCode, _ := qr.Encode(text, qr.M, qr.Auto)

	// Scale the barcode and return it.
	return barcode.Scale(qrCode, size, size)
}
