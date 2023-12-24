package service

import (
	"log"

	"github.com/emejotaw/product-api/internal/entity"
	"github.com/emejotaw/product-api/internal/infra/database"
	"github.com/emejotaw/product-api/internal/infra/database/sqlite"
	"github.com/emejotaw/product-api/internal/types"
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

func (ps *ProductService) Create(productDTO *types.ProductDTO) error {

	product, err := entity.NewProduct(productDTO.Name, productDTO.Price)

	if err != nil {
		log.Printf("could not generate the product, error: %v", err)
		return err
	}

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
