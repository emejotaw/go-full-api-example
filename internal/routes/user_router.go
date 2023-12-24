package routes

import (
	"github.com/emejotaw/product-api/internal/controller"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConfigureUserRoutes(app *fiber.App, db *gorm.DB) {

	userController := controller.NewUserController(db)

	app.Post("/sign-up", userController.Create)
}
