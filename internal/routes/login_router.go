package routes

import (
	"github.com/emejotaw/product-api/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigureLoginRoutes(app *fiber.App) {

	loginController := controller.NewLoginController()
	app.Post("/login", loginController.Login)
}
