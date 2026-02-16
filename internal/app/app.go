package app

import (
	"uplink-api/config"
	"uplink-api/internal/router"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type App struct {
	fiber *fiber.App
	db    *gorm.DB
}

func New(cfg *config.Config) *App {
	db := config.ConnectDatabase(cfg)
	config.AutoMigrate(db)

	app := fiber.New()

	router.Setup(app)

	return &App{
		fiber: app,
		db:    db,
	}
}

func (a *App) Start(addr string) error {
	return a.fiber.Listen(addr)
}