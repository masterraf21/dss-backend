package middleware

import (
	"github.com/golang-jwt/jwt"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	ID    uint32 `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
