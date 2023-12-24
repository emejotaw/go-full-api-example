package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (ls *LoginService) Login(username, password string) ([]byte, error) {

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	return []byte(signedToken), nil
}
