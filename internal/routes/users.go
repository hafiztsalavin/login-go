package routes

import (
	"news-be/internal/controller/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *users.UserController) {
	e.POST("/register", userController.Create)
}
