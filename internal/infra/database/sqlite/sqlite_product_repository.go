package sqlite

import (
	"github.com/emejotaw/product-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, error) {

	products := []entity.Product{}

	err := p.db.Offset((page - 1) * limit).Limit(limit).Find(&products).Error

	return products, err
}

func (p *ProductRepository) FindByID(id string) (*entity.Product, error) {

	product := new(entity.Product)
	err := p.db.First(product, "id = ?", id).Error
	return product, err
}

func (p *ProductRepository) Update(product *entity.Product) error {

	_, err := p.FindByID(product.ID.String())

	if err != nil {
		return err
	}

	return p.db.Save(product).Error
}

func (p *ProductRepository) Delete(id string) error {

	product, err := p.FindByID(id)

	if err != nil {
		return err
	}

	return p.db.Delete(product).Error
}
