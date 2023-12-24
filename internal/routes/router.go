package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type router struct {
}

func NewRouter() *router {
	return &router{}
}

func (r *router) Run(db *gorm.DB) {

	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {

		c.SendString("UP")
		return nil
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	productRouter := NewProductRouter()
	productRouter.Configure(app, db)

	app.Listen(":8000")
}
