package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWTToken(userId uint32, email, role string) (string, error) {
	expirationAccessToken := time.Now().Add(time.Hour * 1)
	claims := &Claims{
		Id:    userId,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationAccessToken.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_ACCESS_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateRefreshToken(accessToken string) (string, error) {
	expirationRefreshToken := time.Now().Add(time.Hour * 6)
	claimsRefresh := &AccessClaims{
		AccessToken: accessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationRefreshToken.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_KEY")))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
