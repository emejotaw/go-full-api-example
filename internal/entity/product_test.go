package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	type testCase struct {
		name        string
		price       float32
		expectedErr error
	}

	tests := []testCase{
		{name: "Rice", price: 4.99, expectedErr: nil},
		{name: "Bean", price: 2.99, expectedErr: nil},
		{name: "", price: 5.466, expectedErr: ErrNameIsRequired},
		{name: "Potato", price: 0, expectedErr: ErrPriceIsRequired},
		{name: "Tomato", price: -3.54, expectedErr: ErrPriceIsInvalid},
	}

	for _, test := range tests {

		_, err := NewProduct(test.name, test.price)

		assert.Equal(t, test.expectedErr, err)
	}
}
