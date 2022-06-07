package routes

import (
	"login-go/internal/controller/auth"
	"login-go/internal/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserAuthPath(e *echo.Echo, authController *auth.AuthController) {
	// auth := e.Group("auth")
	e.GET("/profile", authController.UserDetails, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.UserRole)
}
