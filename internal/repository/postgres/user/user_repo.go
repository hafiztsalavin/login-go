package user

import (
	"news-be/internal/repository/entity"

	"gorm.io/gorm"
)

type UserInterface interface {
	Register(user entity.User) (entity.User, error)
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
