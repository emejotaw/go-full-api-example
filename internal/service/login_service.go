package service

import (
	"errors"
	"time"

	"github.com/emejotaw/product-api/internal/infra/database"
	"github.com/emejotaw/product-api/internal/infra/database/sqlite"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type LoginService struct {
	userRepository database.UserRepository
}

func NewLoginService(db *gorm.DB) *LoginService {

	userRepository := sqlite.NewUserRepository(db)
	return &LoginService{
		userRepository: userRepository,
	}
}

func (ls *LoginService) Login(email, password string) ([]byte, error) {

	user, err := ls.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if !user.ValidatePassword(password) {

		return nil, errors.New("username and/or password invalid")
	}

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
