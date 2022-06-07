package auth

import (
	"login-go/internal/repository/entity"
	"login-go/internal/repository/postgres/auth"
	"login-go/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Repository auth.AuthInterface
}

// init
func NewUserAuthController(userAuthInterface auth.AuthInterface) *AuthController {
	return &AuthController{Repository: userAuthInterface}
}

func (ac AuthController) UserDetails(c echo.Context) error {
	idUser := c.Get("id").(uint)

	user, err := ac.Repository.GetUser(uint(idUser))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	response := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

func (ac AuthController) UserUpdate(c echo.Context) error {
	idUser := c.Get("id").(uint)

	var reqUpdate UpdateUserRequest
	c.Bind(&reqUpdate)
	if err := c.Validate(&reqUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	var password string
	if reqUpdate.Password != "" {
		password, _ = utils.HashPassword(reqUpdate.Password)
	}
	user := entity.User{
		Name:     reqUpdate.Name,
		Email:    reqUpdate.Email,
		Password: password,
	}

	user, err := ac.Repository.UpdateUser(idUser, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (ac AuthController) UserDelete(c echo.Context) error {
	var delUser DeleteUserRequest
	c.Bind(&delUser)
	if err := c.Validate(&delUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	_, err := ac.Repository.DeleteUser(delUser.Name, delUser.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundResponse())
	}

	return c.JSON(http.StatusOK, utils.NewSuccessOperationResponse())
}

func (ac AuthController) CreateAdmin(c echo.Context) error {
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
