package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/service"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Get("/user", controller.GetUser)
	app.Get("/user/create", controller.Create)
	app.Post("/user/store", controller.Store)
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	var users []model.GetUserResponse

	res, _ := controller.userService.GetAll()
	for _, v := range *res {
		user := model.GetUserResponse{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
		users = append(users, user)
	}
	return c.Render("user.index", fiber.Map{"user": users}, "layouts/main")
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	return c.Render("user.create", nil, "layouts/main")
}

func (controller *UserController) Store(c *fiber.Ctx) error {

	var newUser model.CreateUserRequest

	err := c.BodyParser(&newUser)
	if err != nil {
		return err
	}

	// errors := model.ValidateStruct(newProduct)
	// if errors != nil {
	// 	return c.Render("product.create", fiber.Map{"error": errors}, "layouts/main")

	// modified newuser struct to hash password field
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return err
	}

	newUser.Password = string(hashed)

	user := entity.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	err = controller.userService.Insert(&user)

	if err != nil {
		return err
	}

	return c.Redirect("/user")
}
