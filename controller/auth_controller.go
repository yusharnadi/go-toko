package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/service"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) authController {
	return authController{authService}
}

func (controller *authController) Route(app *fiber.App) {
	app.Get("/auth/login", controller.login)
	app.Post("/auth/login", controller.doLogin)
}

func (controller *authController) login(c *fiber.Ctx) error {
	return c.Render("auth.login", nil)
}

func (controller *authController) doLogin(c *fiber.Ctx) error {
	type AuthField struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	var authField AuthField

	err := c.BodyParser(&authField)
	if err != nil {
		return err
	}
	user, err := controller.authService.FindByEmail(authField.Email, &entity.User{})
	if user.ID == 0 {
		fmt.Print("User not Found")
		return nil
	}
	ok := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authField.Password))
	if ok != nil {
		fmt.Print(ok)
		return nil
	}
	fmt.Print("all done, let create session")
	return nil
}
