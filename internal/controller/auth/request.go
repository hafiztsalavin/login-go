package auth

type UpdateUserRequest struct {
	Name     string `json:"name" gorm:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
}

type AdminRequest struct {
	Name     string `json:"name" gorm:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
}
type DeleteUserRequest struct {
	Name  string `json:"name" gorm:"omitempty"`
	Email string `json:"email" validate:"omitempty,email"`
}
