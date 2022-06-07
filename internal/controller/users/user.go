package users

import (
	"net/http"
	"news-be/internal/repository/entity"
	"news-be/internal/repository/postgres/user"
	"news-be/internal/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repository user.UserInterface
}

// init
func NewUserController(userInterface user.UserInterface) *UserController {
	return &UserController{Repository: userInterface}
}

func (uc UserController) Create(c echo.Context) error {
	var userReq CreateUserRequest

	c.Bind(&userReq)
	if err := c.Validate(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	password, _ := utils.HashPassword(userReq.Password)
	user := entity.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: password,
	}

	userDB, err := uc.Repository.Register(user)
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
