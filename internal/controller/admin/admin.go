package admin

import (
	"login-go/internal/repository/entity"
	"login-go/internal/repository/postgres/admin"
	"login-go/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	Repository admin.AdminInterface
}

// init
func NewAdminController(adminInterface admin.AdminInterface) *AdminController {
	return &AdminController{Repository: adminInterface}
}

func (ac AdminController) CreateAdmin(c echo.Context) error {
	var adminReq AdminRequest

	c.Bind(&adminReq)
	if err := c.Validate(&adminReq); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	password, _ := utils.HashPassword(adminReq.Password)
	user := entity.User{
		Name:     adminReq.Name,
		Email:    adminReq.Email,
		Password: password,
	}

	userDB, err := ac.Repository.RegisterAdmin(user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, utils.ErrorResponse(http.StatusNotAcceptable, "Email already exist"))
	}
	response := UserResponse{
		ID:    userDB.ID,
		Name:  userDB.Name,
		Email: userDB.Email,
		Role:  userDB.Role,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (ac AdminController) DeleteAdmin(c echo.Context) error {
	var delUser DeleteAdminRequest
	c.Bind(&delUser)
	if err := c.Validate(&delUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	_, err := ac.Repository.DeleteUserAdmin(delUser.Name, delUser.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}
