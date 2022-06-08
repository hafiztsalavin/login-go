package admin

import (
	"errors"
	"login-go/internal/repository/entity"

	"gorm.io/gorm"
)

type AdminInterface interface {
	RegisterAdmin(entity.User) (entity.User, error)
	DeleteUserAdmin(nama, email string) (entity.User, error)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (ar *AdminRepository) RegisterAdmin(user entity.User) (entity.User, error) {
	user.Role = "admin"
	if err := ar.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ar *AdminRepository) DeleteUserAdmin(nama, email string) (entity.User, error) {
	var user entity.User

	if err := ar.db.Where("name = ? AND email = ?", nama, email).First(&user).Error; err != nil {
		return user, err
	}
	if user.Role == "super admin" {
		return user, errors.New("admin user not allowed to delete")
	}
	ar.db.Delete(&user)

	return user, nil
}
