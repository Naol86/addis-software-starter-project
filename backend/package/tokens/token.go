package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/naol86/addis-software-starter/project/backend/internal/domain"
)

func CreateAccessToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)
	claims := domain.Claims{
		Email: user.Email,
		ID : user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)
	claims := domain.RefreshClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(token string, secret string) (bool, error) {

	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}
	if !tokenString.Valid {
		return false, errors.New("invalid token")
	}
	return true, nil
}

func GetEmail(token string, secret string) (string, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unexpected claims")
	}

	return claims["email"].(string), nil
}

func GetUserId(token string, secret string) (string, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unexpected claims")
	}

	return claims["id"].(string), nil
}

func GetUserClaims(token string, secret string) (jwt.MapClaims, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token)(interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok && !tokenString.Valid {
		return nil, errors.New("Invalid Token for working")
	}
	return claims, nil
}