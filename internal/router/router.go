package router

import (
	"uplink-api/handler"
	"uplink-api/repository"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Dependencies struct {
	AuthenticateService *service.AuthenticateService
	UserRepo            *repository.UserRepository
	ProjectRepo         *repository.ProjectRepository
}

func Setup(app *fiber.App, deps Dependencies) {
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	setupHealthChecks(app)
	setupAPIRoutes(app, deps)
}

func setupHealthChecks(app *fiber.App) {
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())
}

func setupAPIRoutes(app *fiber.App, deps Dependencies) {
	api := app.Group("/api")

	authHandler := handler.NewAuthenticateHandler(deps.AuthenticateService)

	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
}
