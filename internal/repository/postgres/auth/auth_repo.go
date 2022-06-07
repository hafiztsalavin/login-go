package auth

import (
	"login-go/internal/repository/entity"

	"gorm.io/gorm"
)

type AuthInterface interface {
	GetUser(id uint) (entity.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ur *AuthRepository) GetUser(id uint) (entity.User, error) {
	var user entity.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
