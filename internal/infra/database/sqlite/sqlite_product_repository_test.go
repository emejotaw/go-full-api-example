package sqlite

import (
	"errors"
	"testing"

	"github.com/emejotaw/product-api/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func autoMigrateProduct(db *gorm.DB) {

	db.AutoMigrate(&entity.Product{})

}

func TestCreateProduct(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatal("could not establish database connection")
	}

	product, err := entity.NewProduct("Rice", 20.48)

	if err != nil {
		t.Fatal("could not generate new product")
	}

	autoMigrateProduct(db)

	repository := NewProductRepository(db)
	err = repository.Create(product)
	defer repository.Delete(product.ID.String())

	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatal("could not establish database connection")
	}

	autoMigrate(db)

	type testCase struct {
		name  string
		price float32
	}

	tests := []testCase{
		{"Rice", 20.35},
		{"Sugar", 20.35},
		{"Bean", 10.35},
		{"Potato", 4.45},
		{"Milk", 1.36},
	}

	repository := NewProductRepository(db)

	for _, test := range tests {
		product, err := entity.NewProduct(test.name, test.price)

		if err != nil {
			t.Fatal("could not generate product")
		}

		repository.Create(product)
	}

	products, err := repository.FindAll(0, 10, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(products), len(tests))
}

func TestFindByID(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatalf("could not establish database connection, error: %v", err)
	}

	repository := NewProductRepository(db)
	product, err := entity.NewProduct("Potato", 4.46)

	if err != nil {
		t.Fatalf("could not generate product, error: %v", err)
	}

	err = repository.Create(product)

	if err != nil {
		t.Fatalf("could not create product, error: %v", err)
	}

	p, err := repository.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestUpdate(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatalf("could not establish database connection, error: %v", err)
	}

	repository := NewProductRepository(db)
	product, err := entity.NewProduct("Meat", 20.90)
	if err != nil {
		t.Fatalf("could not generate the product, error: %v", err)
	}

	err = repository.Create(product)

	if err != nil {
		t.Fatalf("could not create the product, error: %v", err)
	}

	product.Name = "TOMATO"
	product.Price = 100

	err = repository.Update(product)

	assert.Nil(t, err)
	assert.Equal(t, "TOMATO", product.Name)
	assert.Equal(t, float32(100), product.Price)
}

func TestDelete(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatalf("could not establish database connection, error: %v", err)
	}

	repository := NewProductRepository(db)
	product, err := entity.NewProduct("Pen", 2.45)

	if err != nil {
		t.Fatalf("could not generate the product, error: %v", err)
	}

	err = repository.Create(product)

	if err != nil {
		t.Fatalf("could not create the product, error: %v", err)
	}

	type testCase struct {
		productID     string
		expectedError error
	}

	tests := []testCase{
		{product.ID.String(), nil},
		{uuid.New().String(), errors.New("record not found")},
	}

	for _, test := range tests {

		err = repository.Delete(test.productID)
		assert.Equal(t, test.expectedError, err)
	}
}
