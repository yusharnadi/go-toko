package controller

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	app.Get("/auth/login", login)
}

func login(c *fiber.Ctx) error {
	return c.Render("auth.login", nil)
}
