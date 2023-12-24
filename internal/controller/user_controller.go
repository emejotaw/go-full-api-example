package controller

import (
	"net/http"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/emejotaw/product-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {

	userService := service.NewUserService(db)
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Create(c *fiber.Ctx) error {

	userDTO := new(types.UserDTO)
	err := c.BodyParser(userDTO)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}

	err = uc.userService.Create(userDTO)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}

	c.Status(http.StatusCreated)
	return nil
}
