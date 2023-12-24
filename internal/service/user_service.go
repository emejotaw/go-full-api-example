package service

import (
	"github.com/emejotaw/product-api/internal/entity"
	"github.com/emejotaw/product-api/internal/infra/database"
	"github.com/emejotaw/product-api/internal/infra/database/sqlite"
	"github.com/emejotaw/product-api/internal/types"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository database.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {

	userRepository := sqlite.NewUserRepository(db)
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) Create(userDTO *types.UserDTO) error {

	user, err := entity.New(userDTO.Name, userDTO.Email, userDTO.Password)

	if err != nil {
		return err
	}

	return us.userRepository.Create(user)
}

func (us *UserService) FindByEmail(email string) (*entity.User, error) {

	return us.userRepository.FindByEmail(email)
}
