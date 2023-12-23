package sqlite

import (
	"testing"

	"github.com/emejotaw/product-api/internal/entity"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getDatabase() (*gorm.DB, error) {

	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	return db, err
}

func autoMigrate(db *gorm.DB) {

	db.AutoMigrate(&entity.User{})
}

func TestCreate(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatal("could not get database connection")
	}

	autoMigrate(db)

	user, err := entity.New("jhon doe", "jhondoe@gmail.com", "123456")

	if err != nil {
		t.Fatal("could not map user")
	}

	repository := NewUser(db)
	err = repository.Create(user)

	assert.Nil(t, err)
}

func TestFindByEmail(t *testing.T) {

	db, err := getDatabase()

	if err != nil {
		t.Fatal("could not establish database connection")
	}

	repository := NewUser(db)
	user, err := repository.FindByEmail("jhondoe@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
