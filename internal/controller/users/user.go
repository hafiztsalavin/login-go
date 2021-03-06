package users

import (
	"login-go/internal/auth"
	"login-go/internal/repository/entity"
	"login-go/internal/repository/postgres/user"
	"login-go/internal/utils"
	"net/http"

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

func (uc UserController) Login(c echo.Context) error {
	var loginReq LoginUserRequest

	c.Bind(&loginReq)
	if err := c.Validate(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	userDB, err := uc.Repository.Login(loginReq.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "Not Found"))
	}

	passTrue := utils.CheckPassword(userDB.Password, loginReq.Password)
	if !passTrue {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, "wrong password or email"))
	}

	var accessToken string
	var refreshToken string

	if passTrue {
		accessToken, _ = auth.CreateJWT(userDB.ID, userDB.Email, userDB.Role)
		refreshToken, _ = auth.CreateRefreshToken(accessToken)
	}

	responseData := DataResponse{
		Name: userDB.Name,
		Role: userDB.Role,
	}

	responseToken := TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, utils.SuccessLoginResponse(responseData, responseToken))
}

func (uc UserController) Refresh(c echo.Context) error {
	var refreshReq RefreshTokenRequest

	c.Bind(&refreshReq)
	if err := c.Validate(&refreshReq); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestResponse())
	}

	user, err := auth.ValidateRefresh(refreshReq.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizeResponse())
	}
	userDB, err := uc.Repository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, "Not Found"))
	}

	accessToken, _ := auth.CreateJWT(userDB.ID, userDB.Email, userDB.Role)
	refreshToken, _ := auth.CreateRefreshToken(accessToken)

	responseData := DataResponse{
		Name: userDB.Name,
		Role: userDB.Role,
	}

	responseToken := TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, utils.SuccessLoginResponse(responseData, responseToken))
}

func (uc UserController) HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
