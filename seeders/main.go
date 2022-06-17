package main

import (
	"fmt"
	"log"
	"login-go/internal/config"
	"login-go/internal/repository/entity"
	"login-go/internal/repository/postgres"
	"login-go/internal/utils"

	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgresRepo(&cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	seedersFunc := []func(*gorm.DB) error{
		UserSeed, AdminSeed, SuperAdminSeed}

	for _, function := range seedersFunc {
		if err := function(db); err != nil {
			log.Fatal(err)
		}
	}
}

func UserSeed(db *gorm.DB) error {

	pass, _ := utils.HashPassword("userbiasa")
	var users []entity.User = []entity.User{
		{Name: "hafizts", Email: "hafizts@gmail.com", Password: pass},
		{Name: "tsalavin", Email: "tsalavin@gmail.com", Password: pass},
		{Name: "alvin", Email: "alvin@gmail.com", Password: pass},
	}

	for _, user := range users {
		user := entity.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		}
		db.Create(&user)
	}

	fmt.Println("User(s) created")
	return nil
}

func AdminSeed(db *gorm.DB) error {

	pass, _ := utils.HashPassword("adminbiasa")
	roleAdmin := "admin"
	var admins []entity.User = []entity.User{
		{Name: "admin hafiz", Email: "adminhafizt@gmail.com", Password: pass, Role: roleAdmin},
		{Name: "admin tsalavin", Email: "admintsalavin@gmail.com", Password: pass, Role: roleAdmin},
		{Name: "admin alvin", Email: "adminalvin@gmail.com", Password: pass, Role: roleAdmin},
	}

	for _, admin := range admins {
		admin := entity.User{
			Name:     admin.Name,
			Email:    admin.Email,
			Password: admin.Password,
			Role:     admin.Role,
		}
		db.Create(&admin)
	}
	fmt.Println("Admin(s) created")
	return nil
}

func SuperAdminSeed(db *gorm.DB) error {

	pass, _ := utils.HashPassword("superadmin")
	roleSuper := "super admin"
	var supers []entity.User = []entity.User{
		{Name: "super hafiz", Email: "superhafizt@gmail.com", Password: pass, Role: roleSuper},
		{Name: "super tsalavin", Email: "supertsalavin@gmail.com", Password: pass, Role: roleSuper},
		{Name: "super alvin", Email: "superalvin@gmail.com", Password: pass, Role: roleSuper},
	}

	for _, super := range supers {
		superAdmin := entity.User{
			Name:     super.Name,
			Email:    super.Email,
			Password: super.Password,
			Role:     super.Role,
		}
		db.Create(&superAdmin)
	}

	fmt.Println("Super admin(s) created")
	return nil
}
