package routes

import (
	"github.com/emejotaw/product-api/internal/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConfigureLoginRoutes(app *fiber.App, db *gorm.DB) {

	loginController := controller.NewLoginController(db)
	app.Post("/login", loginController.Login)
}
