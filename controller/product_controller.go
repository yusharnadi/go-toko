package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{*productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/product", controller.Create)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.CreateProductRequest

	err := c.BodyParser(&request)

	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.productService.Insert(request)

	return c.JSON(response)
}
