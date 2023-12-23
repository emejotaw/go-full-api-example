package database

import "github.com/emejotaw/product-api/internal/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll(page, limit, sort int) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
