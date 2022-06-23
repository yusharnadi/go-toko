package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return ProductController{productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/product/", controller.Create)
	app.Get("/product/", controller.GetAll)
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

func (controller *ProductController) GetAll(c *fiber.Ctx) error {
	var products []model.CreateProductResponse

	res, _ := controller.productService.GetAll()
	for _, v := range res {
		productRes := model.CreateProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		}

		products = append(products, productRes)
	}
	// return c.JSON(products)
	return c.Render("index", fiber.Map{"Title": "hello", "data": products}, "layouts/main")
}
