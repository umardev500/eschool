package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenPayload struct {
	Username string
	Duration time.Duration
}

// GenerateToken generates a JWT token based on the given payload.
//
// The payload parameter is a struct that contains the necessary data for generating the token.
//
// The function returns a pointer to a string and an error. The string pointer points to the generated token,
// while the error indicates if there was an error during the token generation process.
func GenerateToken(payload GenerateTokenPayload) (*string, error) {
	secret := os.Getenv("JWT_SECRET")
	claim := jwt.MapClaims{
		"username": payload.Username,
		"exp":      time.Now().Add(payload.Duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(secret))

	return &t, err
}
