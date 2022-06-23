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
	app.Get("/product/:id", controller.Edit)
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
	var products []model.GetProductResponse

	res, _ := controller.productService.GetAll()
	for _, v := range *res {
		productRes := model.GetProductResponse{
			ID:    int(v.ID),
			Name:  v.Name,
			Price: v.Price,
			Stock: v.Stock,
		}

		products = append(products, productRes)
	}
	// return c.JSON(products)
	return c.Render("index", fiber.Map{"data": products}, "layouts/main")
}

func (controller *ProductController) Edit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}

	product, err := controller.productService.FindId(id)
	return c.Render("product.edit", fiber.Map{"data": product}, "layouts/main")
}
