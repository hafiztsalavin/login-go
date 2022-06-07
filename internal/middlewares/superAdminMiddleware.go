package middlewares

import (
	"login-go/internal/auth"
	"login-go/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuperAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, _ := auth.ExtractToken(c)

		if user.Role != "super admin" {
			return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizeResponse())
		}

		c.Set("id", user.Id)
		c.Set("email", user.Email)
		c.Set("role", user.Role)
		return next(c)
	}
}