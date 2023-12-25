package routes

import (
	_ "github.com/emejotaw/product-api/docs"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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

	app.Get("/docs/*", swagger.HandlerDefault)

	ConfigureLoginRoutes(app, db)
	ConfigureUserRoutes(app, db)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	ConfigureProductRoutes(app, db)

	app.Listen(":8000")
}
