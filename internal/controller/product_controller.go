package controller

import (
	"net/http"
	"strconv"

	"github.com/emejotaw/product-api/internal/service"
	"github.com/emejotaw/product-api/internal/types"
	"github.com/gofiber/fiber"
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

func (pc *ProductController) CreateProduct(ctx *fiber.Ctx) {

	productDTO := new(types.ProductDTO)
	err := ctx.BodyParser(productDTO)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = pc.productService.Create(productDTO)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (pc *ProductController) FindAll(ctx *fiber.Ctx) {

	page, errPage := strconv.Atoi(ctx.Query("page"))

	if errPage != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	size, errSize := strconv.Atoi(ctx.Query("size"))

	if errSize != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	sort := ctx.Query("sort")

	products, err := pc.productService.FindAll(page, size, sort)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(products)
}

func (pc *ProductController) FindByID(ctx *fiber.Ctx) {

	productId := ctx.Query("productId")

	product, err := pc.productService.FindByID(productId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(product)
}

func (pc *ProductController) Update(ctx *fiber.Ctx) {

	productId := ctx.Query("productId")
	productDTO := new(types.ProductDTO)
	err := ctx.BodyParser(productDTO)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = pc.productService.Update(productId, productDTO)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (pc *ProductController) Delete(ctx *fiber.Ctx) {

	productID := ctx.Query("productId")

	err := pc.productService.Delete(productID)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
