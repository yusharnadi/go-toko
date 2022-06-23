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
	app.Get("/product/create", controller.Create)
	app.Post("/product/store", controller.Store)
	app.Get("/product/", controller.GetAll)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	return c.Render("product.create", nil, "layouts/main")
}

func (controller *ProductController) Store(c *fiber.Ctx) error {

	var newProduct *model.CreateProductRequest

	err := c.BodyParser(&newProduct)
	if err != nil {
		return c.Render("product.create", fiber.Map{"message": err}, "layouts/main")
	}

	err = controller.productService.Insert(newProduct)

	return c.Redirect("/product")
}

func (controller *ProductController) GetAll(c *fiber.Ctx) error {
	var products []model.CreateProductResponse

	res, _ := controller.productService.GetAll()
	for _, v := range *res {
		productRes := model.CreateProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		}

		products = append(products, productRes)
	}
	// return c.JSON(products)
	return c.Render("index", fiber.Map{"data": products}, "layouts/main")
}
