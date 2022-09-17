package jwt

import (
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
)

var Secret = []byte(os.Getenv("JWT_SECRET"))

type User struct {
	Usuario string `json:"usuario"`
	ID      uint   `json:"ID"`
	jwt.RegisteredClaims
}

type TrocarSenha struct {
	ID string `json:"ID"`
	jwt.RegisteredClaims
}

