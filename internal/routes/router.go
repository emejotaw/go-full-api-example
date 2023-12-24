package routes

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type router struct {
}

func NewRouter() *router {
	return &router{}
}

func (r *router) Run(db *gorm.DB) {

	productRouter := NewProductRouter()
	server := fiber.New()

	productRouter.Configure(server, db)

	server.Listen(":8000")
}
