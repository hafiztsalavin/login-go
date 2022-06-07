package main

import (
	"fmt"
	"news-be/internal/config"
	"news-be/internal/controller/users"

	"news-be/internal/repository/postgres"
	"news-be/internal/repository/postgres/user"
	"news-be/internal/validate"

	"news-be/internal/routes"

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
	userRepo := user.NewUserRepository(db)

	// controller
	userController := users.NewUserController(userRepo)

	// routes
	routes.RegisterPath(e, userController)

	address := fmt.Sprintf(":%d", cfg.Port)
	e.Logger.Fatal(e.Start(address))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
