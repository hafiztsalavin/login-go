package main

import (
	"fmt"
	"log"
	"news-be/internal/config"
	"news-be/internal/repository/entity"
	"news-be/internal/repository/postgres"
	"news-be/internal/utils"

	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := AdminSeed(db); err != nil {
		log.Fatalf("Error when create admin, %v", err)
	}
}

func AdminSeed(db *gorm.DB) error {
	pass, _ := utils.HashPassword("passwordadmin")
	admin := entity.User{
		Name:     "tsalavin",
		Email:    "admintsalavin@gmail.com",
		Role:     "admin",
		Password: pass,
	}

	db.Create(&admin)

	fmt.Println("Admin(s) created")

	return nil
}
