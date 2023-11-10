package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/secretium/secretium/internal/messages"
)

// Claims represents the JWT claims.
type Claims struct {
	jwt.RegisteredClaims
}

// ExtractClaims extracts the JWT claims.
type ExtractClaims struct {
	ExpiresAt int64
}

// GenerateJWT generates a new JWT.
func GenerateJWT(secretKey string) (string, error) {
	// Create a new JWT claims.
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	// Create a new JWT with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// ExtractJWT extracts the all data from a JWT.
func ExtractJWT(tokenFromHeader, secretKey string) (*ExtractClaims, error) {
	// Get 'Authorization' header.
	jwtFromAuthHeader := strings.Split(tokenFromHeader, " ")

	// Check, if the request has a 'Authorization' header.
	if len(jwtFromAuthHeader) == 0 {
		return nil, errors.New(messages.ErrJWTHeaderNotValid)
	}

	// Check, if the request has a 'Authorization' header with 'Bearer' prefix.
	if len(jwtFromAuthHeader) > 0 && jwtFromAuthHeader[0] != "Bearer" {
		return nil, errors.New(messages.ErrJWTHeaderNotValid)
	}

	// Parse the JWT.
	token, err := parseJWT(jwtFromAuthHeader[1], secretKey)
	if err != nil {
		return nil, err
	}

	// Get the expiration time from the JWT.
	expiresAt, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, errors.New(messages.ErrJWTClaimsNotValid)
	}

	// Extract the expiration time from the JWT.
	return &ExtractClaims{
		ExpiresAt: expiresAt.Unix(),
	}, nil
}

// parseJWT parses a JWT.
func parseJWT(tokenFromHeader, secretKey string) (*jwt.Token, error) {
	// Parse the JWT and check for errors.
	return jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}
