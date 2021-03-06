package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateJWT(userId uint, email, role string) (string, error) {
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

func ExtractToken(e echo.Context) (Claims, error) {
	user := e.Get("user").(*jwt.Token)
	claims := &Claims{}
	jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
	})

	return *claims, errors.New("invalid token")
}

func ValidateRefresh(refreshToken string) (Claims, error) {
	accessClaims := &AccessClaims{}
	oldRefreshToken, _ := jwt.ParseWithClaims(refreshToken, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
	})

	if oldRefreshToken.Valid {

		// get Payload from Old Access Token
		claims := &Claims{}
		jwtToken, err := jwt.ParseWithClaims(accessClaims.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
		})

		if jwtToken.Valid {
			return *claims, errors.New("your login sesion still valid")
		}

		if err == nil {
			if err != jwt.ErrSignatureInvalid {
				return *claims, errors.New("your login sesion still valid")
			}
			return *claims, errors.New("your login sesion still valid")
		}

		return *claims, nil
	}

	return Claims{}, errors.New("invalid token")
}
