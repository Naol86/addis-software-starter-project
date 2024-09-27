package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naol86/addis-software-starter/project/backend/config"
	"github.com/naol86/addis-software-starter/project/backend/internal/domain"
	"github.com/naol86/addis-software-starter/project/backend/package/tokens"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env *config.Env
}

func (uc *UserController) Signin(c *gin.Context) {
	var user domain.UserSigninRequest
	var response domain.UserResponse
	if err := c.ShouldBind(&user); err != nil {
		response.Message = "Invalid request"
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userData, err := uc.UserUsecase.SigninUser(c, user)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	accessToken, err := tokens.CreateAccessToken(&userData, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		response.Message = "Failed to create access token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	refreshToken, err := tokens.CreateRefreshToken(&userData, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpiryHour)
	if err != nil {
		response.Message = "Failed to create refresh token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "User signed in successfully"
	response.Success = true
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.Data = userData
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) Signup(c *gin.Context) {
	var user domain.UserSignupRequest
	var response domain.UserResponse
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Message = "Invalid request"
		response.Success = false
		c.JSON(http.StatusBadRequest, response)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Message = "Failed to hash password"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user.Password = string(hashedPassword)
	new_user, err := uc.UserUsecase.SignupUser(c, user)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		c.JSON(http.StatusConflict, response)
		return
	}

	accessToken, err := tokens.CreateAccessToken(&new_user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		response.Message = "Failed to create access token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	refreshToken, err := tokens.CreateRefreshToken(&new_user, uc.Env.RefreshTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		response.Message = "Failed to create access token"
		response.Success = false
		c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "User created successfully"
	response.Success = true
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.Data = new_user
	c.JSON(http.StatusCreated, response)
}