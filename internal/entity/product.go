package entity

import (
	"errors"
	"time"

	"github.com/emejotaw/product-api/pkg/entity"
)

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt string    `json:"createdAt"`
}

func NewProduct(name string, price float32) (*Product, error) {

	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().String(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrPriceIsInvalid
	}

	return nil
}
