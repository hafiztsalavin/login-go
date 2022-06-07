package auth

import (
	"github.com/golang-jwt/jwt"
)

// result as an auth token
type Claims struct {
	Id    uint
	Email string
	Role  string
	jwt.StandardClaims
}

type AccessClaims struct {
	AccessToken string
	jwt.StandardClaims
}

type Payload struct {
	Id    uint
	Email string
	Role  string
}
