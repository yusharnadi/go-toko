package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/model"
	"github.com/yusharnadi/go-toko/service"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	authService service.AuthService
	// store *Store
	store *session.Store
}

func NewAuthController(authService service.AuthService, store *session.Store) authController {
	return authController{authService, store}
}

func (controller *authController) Route(app *fiber.App) {
	app.Get("/auth/login", controller.login)
	app.Get("/auth/register", controller.register)
	app.Get("/auth/logout", controller.logout)
	app.Post("/auth/login", controller.doLogin)
	app.Post("/auth/register", controller.doRegister)
}

func (controller *authController) login(c *fiber.Ctx) error {
	// Get session from storage
	sess, err := controller.store.Get(c)

	if err != nil {
		fmt.Println(err)
	}

	if sess.Get("secret") != nil {
		return c.Redirect("/")
	}

	return c.Render("auth.login", nil)
}

func (controller *authController) register(c *fiber.Ctx) error {
	// Get session from storage
	sess, err := controller.store.Get(c)

	if err != nil {
		fmt.Println(err)
	}

	if sess.Get("secret") != nil {
		return c.Redirect("/")
	}

	return c.Render("auth.register", nil)
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

		return c.Redirect("/auth/login")
	}

	sess, err := controller.store.Get(c)

	if err != nil {
		return err
	}

	// Set key/value
	sess.Set("secret", "authenticated")
	sess.Set("email", user.Email)

	if err := sess.Save(); err != nil {
		return err
	}

	return c.Redirect("/")
}

func (controller *authController) doRegister(c *fiber.Ctx) error {
	var newUser model.CreateUserRequest

	err := c.BodyParser(&newUser)
	if err != nil {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: string(hashed),
	}

	err = controller.authService.Register(&user)
	if err != nil {
		return err
	}

	return c.Redirect("/")
}

func (controller *authController) logout(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err != nil {
		return err
	}

	// Destroy session
	if err := sess.Destroy(); err != nil {
		return err
	}

	return c.Redirect("/auth/login")
}
