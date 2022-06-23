package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
	"github.com/yusharnadi/go-toko/controller"
	"github.com/yusharnadi/go-toko/repository"
	"github.com/yusharnadi/go-toko/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/go_toko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a new engine
	engine := handlebars.New("./views", ".hbs")

	// Setup repository
	productRepository := repository.NewProductRepository(db)

	//setup Service
	productService := service.NewProductService(productRepository)

	//setup Controller
	productController := controller.NewProductController(productService)

	// Setup Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static/", "./assets")

	// Setup Routing
	productController.Route(app)

	// Start App
	errors := app.Listen(":3000")
	if errors != nil {
		log.Fatal(errors.Error())
	}
}
