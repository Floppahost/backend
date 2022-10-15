package jwt

import (
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
)

var Secret = []byte(os.Getenv("JWT_SECRET"))

// base user Struct; may change later
type User struct {
	User string `json:"usuario"`
	ID      uint   `json:"ID"`
	jwt.RegisteredClaims
}

// change password token
type ChangePassword struct {
	ID string `json:"ID"`
	jwt.RegisteredClaims
}

