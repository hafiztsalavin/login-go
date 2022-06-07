package routes

import (
	"login-go/internal/controller/admin"
	"login-go/internal/middlewares"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AdminPath(e *echo.Echo, adminController *admin.AdminController) {
	root := e.Group("/root")
	root.POST("/register", adminController.CreateAdmin, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.SuperAdminRole)
	root.PUT("/delete", adminController.DeleteAdmin, middleware.JWT([]byte(os.Getenv("JWT_ACCESS_KEY"))), middlewares.SuperAdminRole)

}
