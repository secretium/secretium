package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// DecryptString decrypts a string that was encrypted using AES encryption with a given secret key.
func DecryptString(secretKey, encryptedText string) (string, error) {
	// Create a new AES cipher block using the key.
	block, err := aes.NewCipher([]byte(fmt.Sprintf("%-16s", secretKey)))
	if err != nil {
		return "", err
	}

	// Decode the encrypted text from hexadecimal to bytes.
	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Get the initialization vector (IV) from the beginning of the ciphertext.
	iv := ciphertext[:aes.BlockSize]
	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a buffer to hold the decrypted text.
	decryptedText := make([]byte, len(ciphertext))

	// Create a new CBC decrypt with the block and IV.
	mode := cipher.NewCBCDecrypter(block, iv)
	// Decrypt the ciphertext and store the result in the decryptedText buffer.
	mode.CryptBlocks(decryptedText, ciphertext)

	// Get the padding value from the last byte of the decrypted text.
	padding := int(decryptedText[len(decryptedText)-1])
	// Remove the padding from the decrypted text
	decryptedText = decryptedText[:len(decryptedText)-padding]

	// Return the decrypted text as a string.
	return string(decryptedText), nil
}
