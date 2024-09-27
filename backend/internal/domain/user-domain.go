package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID             bson.ObjectID `bson:"_id"`
	Name           string        `bson:"name"`
	Email          string        `bson:"email"`
	Password       string        `bson:"password" json:"-"`
	ProfilePicture string        `bson:"profilepicture"`
}

type UserSignupRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
}

type UserSigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	Data         User   `json:"data"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID string
	jwt.RegisteredClaims
}

type UserRepository interface {
	CreateUser(c context.Context, user UserSignupRequest) (User, error)
	FindByID(c context.Context, id string) (User, error)
	FindByEmail(c context.Context, email string) (User, error)
}

type UserUsecase interface {
	SignupUser(c context.Context, user UserSignupRequest) (User, error)
	SigninUser(c context.Context, user UserSigninRequest) (User, error)
}
