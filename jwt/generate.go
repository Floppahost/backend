package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateUserToken(Autor string, Username string) (string, error) {
	var token string

	// todo: change expiration
	expiration := time.Now().Add(5 * time.Minute)

	claim := &User{
		User: Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   Autor,
		},
	}

	// the token itself
	token_bruto := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// string token
	token, err := token_bruto.SignedString(Secret)

	if err != nil {
		return "", err
	}
	return token, nil
}
