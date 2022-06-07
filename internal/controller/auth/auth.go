package auth

import (
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
