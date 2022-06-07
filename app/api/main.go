package main

import (
	"fmt"
	"login-go/internal/config"
	cadm "login-go/internal/controller/admin"
	ca "login-go/internal/controller/auth"
	cu "login-go/internal/controller/users"

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

	// Initialize news config
	cfg := config.NewConfig()

	// Initialize DB repositories
	db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	checkErr(err)

	e := echo.New()

	// logging
	// middlewares.LogMiddleware(e)

	// validator
	e.Validator = &validate.Validator{Validator: validator.New()}

	// repository
	userRepo := ur.NewUserRepository(db)
	authRepo := ar.NewUserAuthRepository(db)
	adminRepo := adr.NewAdminRepository(db)

	// controller
	userController := cu.NewUserController(userRepo)
	authController := ca.NewUserAuthController(authRepo)
	adminController := cadm.NewAdminController(adminRepo)

	// routes
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
