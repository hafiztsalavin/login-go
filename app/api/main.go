package main

import (
	"fmt"
	"login-go/internal/config"
	cadm "login-go/internal/controller/admin"
	ca "login-go/internal/controller/auth"
	cu "login-go/internal/controller/users"
	"login-go/internal/middlewares"

	"login-go/internal/repository/postgres"
	adr "login-go/internal/repository/postgres/admin"
	ar "login-go/internal/repository/postgres/auth"
	ur "login-go/internal/repository/postgres/user"
	"login-go/internal/validate"

	"login-go/internal/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	e := echo.New()

	e.Use(middlewares.Logging)
	e.HTTPErrorHandler = middlewares.ErrorHandler

	e.Validator = &validate.Validator{Validator: validator.New()}

	userRepo := ur.NewUserRepository(db)
	userController := cu.NewUserController(userRepo)

	authRepo := ar.NewUserAuthRepository(db)
	authController := ca.NewUserAuthController(authRepo)

	adminRepo := adr.NewAdminRepository(db)
	adminController := cadm.NewAdminController(adminRepo)

	routes.RegisterPath(e, userController)
	routes.UserAuthPath(e, authController)
	routes.AdminPath(e, adminController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
