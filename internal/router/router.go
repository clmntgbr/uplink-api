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

	authenticateMiddleware := middleware.NewAuthenticateMiddleware(deps.AuthenticateService, deps.UserRepo)

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
	api.Post("/projects/activate", projectHandler.ActivateProject)
}

func setupEndpointRoutes(api fiber.Router, deps Dependencies) {
	endpointHandler := handler.NewEndpointHandler(deps.EndpointService)

	api.Get("/endpoints", endpointHandler.GetEndpoints)
	api.Post("/endpoints", endpointHandler.CreateEndpoint)
}

func setupWorkflowRoutes(api fiber.Router, deps Dependencies) {
	workflowHandler := handler.NewWorkflowHandler(deps.WorkflowService)

	api.Get("/workflows", workflowHandler.GetWorkflows)
	api.Post("/workflows", workflowHandler.CreateWorkflow)
}
