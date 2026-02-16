package main

import (
	"github.com/gofiber/fiber/v3"
	"uplink-api/config"
)

func main() {
    cfg := config.Load()

	db := config.ConnectDatabase(cfg)
	config.AutoMigrate(db)

    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
