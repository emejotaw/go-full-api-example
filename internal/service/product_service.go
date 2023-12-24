package service

import (
	"github.com/emejotaw/product-api/internal/entity"
	"github.com/emejotaw/product-api/internal/infra/database"
	"github.com/emejotaw/product-api/internal/infra/database/sqlite"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepository database.ProductRepository
}

func NewProductService(db *gorm.DB) *ProductService {

	repository := sqlite.NewProductRepository(db)
	return &ProductService{
		productRepository: repository,
	}
}

func (ps *ProductService) Create(product *entity.Product) error {

	return ps.productRepository.Create(product)
}

func (ps *ProductService) FindAll(page, size int, sort string) ([]entity.Product, error) {

	return ps.productRepository.FindAll(page, size, sort)
}

func (ps *ProductService) FindByID(productId string) (*entity.Product, error) {
	return ps.productRepository.FindByID(productId)
}

func (ps *ProductService) Update(product *entity.Product) error {
	return ps.productRepository.Update(product)
}

func (ps *ProductService) Delete(productId string) error {
	return ps.productRepository.Delete(productId)
}
