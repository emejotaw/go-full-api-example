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

// Create product godoc
// @Summary Create product
// @Description Endpoint to create a product
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "access token"
// @Param request body types.ProductDTO true "product.request"
// @Success 201
// @Failure 400
// @Failure 500
// @Router  /api/v1/products [post]
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

// Find all products godoc
// @Summary Find all products
// @Description List all the products paginated
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "access token"
// @Param page query int true "1"
// @Param size query int true "10"
// @Param sort query string false "asc|desc"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/products [get]
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

// Get product godoc
// @Summary Get product by id
// @Description Receives an id and returns a product by given id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "access token"
// @Param productId query string true "1"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/products/get [get]
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

// Update product godoc
// @Summary Update product
// @Description Endpoint to update a product
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "access token"
// @Param request body types.ProductDTO true "product.request"
// @Success 204
// @Failure 400
// @Failure 500
// @Router  /api/v1/products [put]
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

// Delete product godoc
// @Summary Delete product by id
// @Description Receives an id and delete a product by given id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "access token"
// @Param productId query string true "1"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/v1/products [delete]
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
