package app

import (
	"uplink-api/config"
	"uplink-api/internal/router"
	"uplink-api/repository"
	"uplink-api/rules"
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
	endpointRepo := repository.NewEndpointRepository(db)
	workflowRepo := repository.NewWorkflowRepository(db)

	projectRules := rules.NewProjectRules(projectRepo)

	authenticateService := service.NewAuthenticateService(userRepo, projectRepo, cfg)
	userService := service.NewUserService()
	projectService := service.NewProjectService(projectRepo, userRepo, projectRules)
	endpointService := service.NewEndpointService(endpointRepo, projectRepo, userRepo)
	workflowService := service.NewWorkflowService(workflowRepo, projectRepo, userRepo)

	app := fiber.New()

	router.Setup(app, router.Dependencies{
		AuthenticateService: authenticateService,
		UserRepo:            userRepo,
		ProjectRepo:         projectRepo,
		UserService:         userService,
		ProjectService:      projectService,
		EndpointService:     endpointService,
		WorkflowService:     workflowService,
	})

	return &App{
		fiber: app,
		db:    db,
	}
}

func (a *App) Start(addr string) error {
	return a.fiber.Listen(addr)
}
