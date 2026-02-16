package router

import (
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}