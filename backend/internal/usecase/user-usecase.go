package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/naol86/addis-software-starter/project/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
	Timeout        time.Duration
}

// SigninUser implements domain.UserUsecase.
func (u *UserUsecase) SigninUser(c context.Context, user domain.UserSigninRequest) (domain.User, error) {
	userData, err := u.UserRepository.FindByEmail(c, user.Email)
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, errors.New("Invalid password")
	}
	return userData, nil
}

// SignupUser implements domain.UserUsecase.
func (u *UserUsecase) SignupUser(c context.Context, user domain.UserSignupRequest) (domain.User, error) {
	_, err := u.UserRepository.FindByEmail(c, user.Email)
	if err == nil {
		return domain.User{}, errors.New("User already exists")
	}

	return u.UserRepository.CreateUser(c, user)
}

func NewUserUseCase(timeout time.Duration, userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		Timeout:        timeout,
	}
}
