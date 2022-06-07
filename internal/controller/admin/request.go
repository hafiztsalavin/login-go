package admin

type AdminRequest struct {
	Name     string `json:"name" gorm:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8"`
}
type DeleteAdminRequest struct {
	Name  string `json:"name" gorm:"omitempty"`
	Email string `json:"email" validate:"omitempty,email"`
}
