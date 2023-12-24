package routes

import (
	"github.com/emejotaw/product-api/internal/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type productRouter struct {
}

func NewProductRouter() *productRouter {
	return &productRouter{}
}

func (r *productRouter) Configure(server *fiber.App, db *gorm.DB) {

	productController := controller.NewProductController(db)

	server.Post("/api/v1/products", productController.CreateProduct)
	server.Get("/api/v1/products", productController.FindAll)
	server.Get("/api/v1/products/get", productController.FindByID)
	server.Put("/api/v1/products", productController.Update)
	server.Delete("/api/v1/products", productController.Delete)
}
