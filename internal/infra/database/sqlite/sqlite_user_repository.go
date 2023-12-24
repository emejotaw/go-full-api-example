package sqlite

import (
	"github.com/emejotaw/product-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *entity.User) error {

	return u.db.Create(user).Error
}

func (p *UserRepository) FindByEmail(email string) (*entity.User, error) {

	user := &entity.User{}
	p.db.Where("email = ?", email).Find(user)
	return user, nil
}
