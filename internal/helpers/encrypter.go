package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// EncryptString encrypts a string using AES encryption with a given secret key.
// It generates a random initialization vector (IV), encrypts the text using AES in CBC mode,
// appends the IV to the ciphertext, and returns the ciphertext as a hexadecimal string.
func EncryptString(secretKey, text string) (string, error) {
	// Create a new AES cipher using the key.
	block, err := aes.NewCipher([]byte(fmt.Sprintf("%-16s", secretKey)))
	if block == nil || err != nil {
		return "", err
	}

	// Generate a random initialization vector (IV).
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// Pad the text with bytes representing the number of padding bytes.
	paddedText := []byte(text)
	blockSize := block.BlockSize()
	padding := blockSize - len(paddedText)%blockSize
	paddedText = append(paddedText, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Create a buffer to hold the ciphertext with room for the IV.
	ciphertext := make([]byte, aes.BlockSize+len(paddedText))

	// Encrypt the padded text using AES in CBC mode.
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedText)

	// Copy the IV to the beginning of the ciphertext.
	copy(ciphertext[:aes.BlockSize], iv)

	// Return the ciphertext as a hexadecimal string.
	return hex.EncodeToString(ciphertext), nil
}
