package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	type testCase struct {
		name     string
		email    string
		password string
	}
	tests := []testCase{
		{name: "Jhon Doe", email: "jhondoe@gmail.com", password: "12345678"},
		{name: "Sara Doe", email: "saradoe@gmail.com", password: "87654321"},
		{name: "Peter Doe", email: "peterdoe@gmail.com", password: "87654321"},
		{name: "Mike Doe", email: "mikedoe@gmail.com", password: "12345678"},
		{name: "Ana Doe", email: "anadoe@gmail.com", password: "12345678"},
		{name: "Lana Doe", email: "lanadoe@gmail.com", password: "12345678"},
		{name: "Elizabeth Doe", email: "elizabethdoe@gmail.com", password: "12345678"},
		{name: "Walter Doe", email: "walterdoe@gmail.com", password: "12345678"},
		{name: "Jordan Doe", email: "jordandoe@gmail.com", password: "12345678"},
		{name: "Luke Doe", email: "lukedoe@gmail.com", password: "12345678"},
	}

	for _, test := range tests {

		_, err := New(test.name, test.email, test.password)

		assert.Nil(t, err)
	}
}

func TestValidatePassword(t *testing.T) {

	type testCase struct {
		name     string
		email    string
		password string
	}
	tests := []testCase{
		{name: "Jhon Doe", email: "jhondoe@gmail.com", password: "12345678"},
		{name: "Sara Doe", email: "saradoe@gmail.com", password: "87654321"},
		{name: "Peter Doe", email: "peterdoe@gmail.com", password: "87654321"},
		{name: "Mike Doe", email: "mikedoe@gmail.com", password: "12345678"},
		{name: "Ana Doe", email: "anadoe@gmail.com", password: "12345678"},
		{name: "Lana Doe", email: "lanadoe@gmail.com", password: "12345678"},
		{name: "Elizabeth Doe", email: "elizabethdoe@gmail.com", password: "12345678"},
		{name: "Walter Doe", email: "walterdoe@gmail.com", password: "12345678"},
		{name: "Jordan Doe", email: "jordandoe@gmail.com", password: "12345678"},
		{name: "Luke Doe", email: "lukedoe@gmail.com", password: "12345678"},
	}

	for _, test := range tests {

		user, _ := New(test.name, test.email, test.password)
		isAuthenticated := user.ValidatePassword(test.password)
		assert.True(t, isAuthenticated)
	}
}
