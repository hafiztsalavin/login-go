package routes

import (
	"login-go/internal/controller/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *users.UserController) {
	e.POST("/register", userController.Create)
	e.POST("/login", userController.Login)
	e.POST("/refresh", userController.Refresh)
	e.GET("/ping", userController.HomePage)
}
