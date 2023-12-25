package controller

import (
	"net/http"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(db *gorm.DB) *LoginController {

	loginService := service.NewLoginService(db)
	return &LoginController{
		loginService: loginService,
	}
}

// Login godoc
// @Summary      Endpoint to authenticate
// @Description  Receives an email and a password and returns a JWT
// @Tags         login
// @Accept       json
// @Produce      json
// @Param		 request body types.UserDTO true "user.request"
// @Success      201
// @Failure      400
// @Failure      500
// @Router       /login [post]
func (lc *LoginController) Login(c *fiber.Ctx) error {

	username := c.FormValue("email")
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
