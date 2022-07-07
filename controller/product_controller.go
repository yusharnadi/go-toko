package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/service"
)

type ProductController struct {
	productService service.ProductService
	store          *session.Store
}

func NewProductController(productService service.ProductService, store *session.Store) ProductController {
	return ProductController{productService, store}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Get("/", controller.Home)
	app.Get("/product/create", controller.Create)
	app.Get("/product/:id", controller.Edit)
	app.Post("/product/:id/update", controller.Update)
	app.Get("/product/:id/delete", controller.Delete)
	app.Post("/product/store", controller.Store)
	app.Get("/product/", controller.GetAll)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	return c.Render("product.create", nil, "layouts/main")
}

func (controller *ProductController) Home(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	if sess.Get("secret") == nil {
		return c.Redirect("auth/login")
	}
	return c.Render("home", fiber.Map{"session": sess.Get("email")}, "layouts/main")
}

func (controller *ProductController) Store(c *fiber.Ctx) error {

	var newProduct model.CreateProductRequest

	err := c.BodyParser(&newProduct)
	if err != nil {
		return err
	}

	errors := model.ValidateStruct(newProduct)
	if errors != nil {
		return c.Render("product.create", fiber.Map{"error": errors}, "layouts/main")

	}
	product := entity.Product{
		Name:  newProduct.Name,
		Price: newProduct.Price,
		Stock: newProduct.Stock,
	}

	err = controller.productService.Insert(&product)

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
	return c.Render("product.index", fiber.Map{"data": products}, "layouts/main")
}

func (controller *ProductController) Edit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}

	product, err := controller.productService.FindId(id)
	if err != nil {
		return c.SendStatus(404)
	}
	return c.Render("product.edit", fiber.Map{"data": product}, "layouts/main")
}

func (controller *ProductController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}

	var newproduct model.UpdateProductRequest

	newproduct.UpdatedAt = time.Now()

	err = c.BodyParser(&newproduct)
	if err != nil {
		return err
	}

	var Product entity.Product

	Product.Name = newproduct.Name
	Product.Price = newproduct.Price
	Product.Stock = newproduct.Stock
	Product.UpdatedAt = newproduct.UpdatedAt

	err = controller.productService.Update(&Product, id)

	if err != nil {
		return err
	}

	return c.Redirect("/product")

}

func (controller *ProductController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(404)
	}

	if errDel := controller.productService.Delete(id); errDel != nil {
		return errDel
	}
	return c.Redirect("/product")
}
