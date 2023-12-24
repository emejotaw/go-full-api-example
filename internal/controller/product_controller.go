package controller

import (
	"net/http"
	"strconv"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/emejotaw/product-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {

	productService := service.NewProductService(db)
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) CreateProduct(ctx *fiber.Ctx) error {

	productDTO := new(types.ProductDTO)
	err := ctx.BodyParser(productDTO)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	err = pc.productService.Create(productDTO)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (pc *ProductController) FindAll(ctx *fiber.Ctx) error {

	page, errPage := strconv.Atoi(ctx.Query("page"))

	if errPage != nil {
		ctx.Status(http.StatusBadRequest)
		return errPage
	}

	size, errSize := strconv.Atoi(ctx.Query("size"))

	if errSize != nil {
		ctx.Status(http.StatusBadRequest)
		return errSize
	}

	sort := ctx.Query("sort")

	products, err := pc.productService.FindAll(page, size, sort)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(products)
	return nil
}

func (pc *ProductController) FindByID(ctx *fiber.Ctx) error {

	productId := ctx.Query("productId")

	product, err := pc.productService.FindByID(productId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(product)
	return nil
}

func (pc *ProductController) Update(ctx *fiber.Ctx) error {

	productId := ctx.Query("productId")
	productDTO := new(types.ProductDTO)
	err := ctx.BodyParser(productDTO)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return err
	}

	err = pc.productService.Update(productId, productDTO)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

func (pc *ProductController) Delete(ctx *fiber.Ctx) error {

	productID := ctx.Query("productId")

	err := pc.productService.Delete(productID)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}
