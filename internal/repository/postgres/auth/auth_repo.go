package auth

import (
	"login-go/internal/repository/entity"

	"gorm.io/gorm"
)

type AuthInterface interface {
	GetUser(id uint) (entity.User, error)
	UpdateUser(id uint, updateUser entity.User) (entity.User, error)
	DeleteUser(name, email string) (entity.User, error)
	RegisterAdmin(entity.User) (entity.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) RegisterAdmin(user entity.User) (entity.User, error) {
	user.Role = "admin"
	if err := ar.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ar *AuthRepository) GetUser(id uint) (entity.User, error) {
	var user entity.User
	if err := ar.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ar *AuthRepository) UpdateUser(id uint, updateUser entity.User) (entity.User, error) {
	var user entity.User

	if err := ar.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	ar.db.Model(&user).Updates(updateUser)

	return updateUser, nil
}

func (ar *AuthRepository) DeleteUser(nama, email string) (entity.User, error) {
	var user entity.User

	if err := ar.db.Where("name = ? AND email = ?", nama, email).First(&user).Error; err != nil {
		return user, err
	}

	ar.db.Delete(&user)

	return user, nil
}
