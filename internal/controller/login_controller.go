package controller

import (
	"net/http"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController() *LoginController {

	loginService := service.NewLoginService()
	return &LoginController{
		loginService: loginService,
	}
}

func (lc *LoginController) Login(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	tokenBytes, err := lc.loginService.Login(username, password)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}

	return c.JSON(fiber.Map{
		"accessToken": string(tokenBytes),
	})
}
