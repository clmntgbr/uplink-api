package router

import (
	"uplink-api/handler"
	"uplink-api/middleware"
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
	StepRepo            *repository.StepRepository
	UserService         *service.UserService
	ProjectService      *service.ProjectService
	EndpointService     *service.EndpointService
	WorkflowService     *service.WorkflowService
}

func Setup(app *fiber.App, deps Dependencies) {
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	setupHealthChecks(app)
	setupAPIRoutes(app, deps)
}

func setupAPIRoutes(app *fiber.App, deps Dependencies) {
	api := app.Group("/api")

	authenticateMiddleware := middleware.NewAuthenticateMiddleware(deps.AuthenticateService, deps.UserRepo, deps.ProjectRepo)

	setupAuthRoutes(api, deps)

	api.Use(authenticateMiddleware.Protected())

	setupUserRoutes(api, deps)
	setupProjectRoutes(api, deps)
	setupWorkflowRoutes(api, deps)
	setupEndpointRoutes(api, deps)
}

func setupHealthChecks(app *fiber.App) {
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())
}

func setupAuthRoutes(api fiber.Router, deps Dependencies) {
	authenticateHandler := handler.NewAuthenticateHandler(deps.AuthenticateService)

	api.Post("/login", authenticateHandler.Login)
	api.Post("/register", authenticateHandler.Register)
}

func setupUserRoutes(api fiber.Router, deps Dependencies) {
	userHandler := handler.NewUserHandler(deps.UserService)

	api.Get("/user", userHandler.GetUser)
}

func setupProjectRoutes(api fiber.Router, deps Dependencies) {
	projectHandler := handler.NewProjectHandler(deps.ProjectService)

	api.Get("/projects", projectHandler.GetProjects)
	api.Get("/projects/:id", projectHandler.GetProjectByID)
	api.Post("/projects", projectHandler.CreateProject)
	api.Put("/projects/:id", projectHandler.UpdateProject)
	api.Post("/projects/activate", projectHandler.ActivateProject)
}

func setupEndpointRoutes(api fiber.Router, deps Dependencies) {
	endpointHandler := handler.NewEndpointHandler(deps.EndpointService)

	api.Get("/endpoints", endpointHandler.GetEndpoints)
	api.Post("/endpoints", endpointHandler.CreateEndpoint)
	api.Get("/endpoints/:id", endpointHandler.GetEndpointByID)
	api.Put("/endpoints/:id", endpointHandler.UpdateEndpoint)
}

func setupWorkflowRoutes(api fiber.Router, deps Dependencies) {
	workflowHandler := handler.NewWorkflowHandler(deps.WorkflowService)

	api.Get("/workflows", workflowHandler.GetWorkflows)
	api.Post("/workflows", workflowHandler.CreateWorkflow)
	api.Put("/workflows/:id", workflowHandler.UpdateWorkflow)
	api.Get("/workflows/:id", workflowHandler.GetWorkflowByID)
	api.Get("/workflows/:id/steps", workflowHandler.GetStepsByWorkflowID)
	api.Post("/workflows/:id/steps", workflowHandler.CreateStepByWorkflowID)
}
