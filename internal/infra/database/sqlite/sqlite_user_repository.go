package sqlite

import (
	"github.com/emejotaw/product-api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(user *entity.User) error {

	return u.db.Create(user).Error
}

func (p *User) FindByEmail(email string) (*entity.User, error) {

	user := &entity.User{}
	p.db.Where("email = ?", email).Find(user)
	return user, nil
}
