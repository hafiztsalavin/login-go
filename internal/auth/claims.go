package auth

import (
	"github.com/golang-jwt/jwt"
)

// result as an auth token
type Claims struct {
	Id    uint32
	Email string
	Role  string
	jwt.StandardClaims
}

type AccessClaims struct {
	AccessToken string
	jwt.StandardClaims
}
