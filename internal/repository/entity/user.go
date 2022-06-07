package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Email    string
	Password string
	Role     string `gorm:"default:user"`
}
