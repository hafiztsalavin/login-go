package routes

import (
	"login-go/internal/controller/auth"
	"login-go/internal/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserAuthPath(e *echo.Echo, authController *auth.AuthController) {
	auth := e.Group("/auth")
	auth.GET("/profile", authController.UserDetails, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.UserRole)
	auth.PUT("/update", authController.UserUpdate, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.UserRole)

	admin := e.Group("/admin")
	admin.GET("/profile", authController.UserDetails, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.AdminRole)
	auth.PUT("/delete", authController.UserDelete, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.AdminRole)
}
