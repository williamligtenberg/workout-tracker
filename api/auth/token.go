package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

func Init(secret string) {
	jwtKey = []byte(secret)
}

func GenerateToken(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "workout-tracker",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
