package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func createJWT() (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		Issuer:    "test",
		Audience:  []string{"Platform-DEV"},
		Subject:   "5d1710b1-6bb4-44b3-bd4b-f9edd50b1c10",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}
