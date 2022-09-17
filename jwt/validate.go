package jwt

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v4"
)

func Validar(tokens string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (any, error) {
		// verificamos se houve erro no método de criptografia
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, errors.New("algorítimo inválido")
		}
		return Secret, nil
	})

	// verificamos se houve erro na validação do algorítimo
	if err != nil {
		return nil, err
	}

	// pegamos os campos/claims do token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

