package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateUserToken(Autor string, Usuario string, ID uint, permanecerConectado bool) (string, error) {
	var token string
	expiration := time.Now().Add(5 * time.Minute)
	if permanecerConectado {
		expiration = time.Now().Add(30 * time.Hour)
	}
	claim := &User{
		Usuario: Usuario,
		ID:      ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   Autor,
		},
	}
	token_bruto := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := token_bruto.SignedString(Secret)

	if err != nil {
		return "", err
	}
	return token, nil
}

