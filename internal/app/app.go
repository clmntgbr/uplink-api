package app

import (
	"uplink-api/config"
	"uplink-api/internal/router"
	"uplink-api/repository"
	"uplink-api/service"

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

	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	authenticateService := service.NewAuthenticateService(userRepo, projectRepo, cfg)
	userService := service.NewUserService(userRepo, cfg)
	projectService := service.NewProjectService(projectRepo, cfg)

	app := fiber.New()

	router.Setup(app, router.Dependencies{
		AuthenticateService: authenticateService,
		UserRepo:            userRepo,
		UserService:         userService,
		ProjectService:      projectService,
	})

	return &App{
		fiber: app,
		db:    db,
	}
}

func (a *App) Start(addr string) error {
	return a.fiber.Listen(addr)
}
