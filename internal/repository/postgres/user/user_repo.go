package user

import (
	"login-go/internal/repository/entity"

	"gorm.io/gorm"
)

type UserInterface interface {
	Register(user entity.User) (entity.User, error)
	Login(email string) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Register(user entity.User) (entity.User, error) {

	if err := ur.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Login(email string) (entity.User, error) {
	var user entity.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
