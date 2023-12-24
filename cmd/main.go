package main

import (
	"log"

	"github.com/emejotaw/product-api/config"
	"github.com/emejotaw/product-api/internal/entity"
	"github.com/emejotaw/product-api/internal/routes"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {

	_ = config.New(".")
	dsn := "database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("could not establish database connection, error: %v", err)
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	router := routes.NewRouter()
	router.Run(db)
}
