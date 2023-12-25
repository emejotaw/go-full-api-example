package main

import (
	"log"

	"github.com/emejotaw/product-api/config"
	"github.com/emejotaw/product-api/internal/entity"
	"github.com/emejotaw/product-api/internal/routes"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// @title Api to create products
// @version 1.0
// @description Product management api
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
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
