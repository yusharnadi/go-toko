package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/yusharnadi/go-toko/controller"
	"github.com/yusharnadi/go-toko/repository"
	"github.com/yusharnadi/go-toko/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	if errCfg := godotenv.Load(); errCfg != nil {
		log.Fatal(errCfg)
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")

	dsn := db_user + ":" + db_pass + "@tcp(127.0.0.1:3306)/go_toko?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a new engine
	engine := html.New("./views", ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false
	store := session.New()

	// Setup repository
	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)
	authRepository := repository.NewAuthRepository(db)

	//setup Service
	productService := service.NewProductService(productRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(authRepository)

	//setup Controller
	productController := controller.NewProductController(productService, store)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService, store)

	// Setup Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/", "./assets")

	// Setup Routing
	productController.Route(app)
	userController.Route(app)
	authController.Route(app)

	// Start App
	errors := app.Listen(":3000")
	if errors != nil {
		log.Fatal(errors.Error())
	}
}
