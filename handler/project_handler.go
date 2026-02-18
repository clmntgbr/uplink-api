package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) GetProjects(c fiber.Ctx) error {
	userID, err := ctxutil.GetUserID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	projects, err := h.projectService.GetProjects(c.Context(), userID)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}

func (h *ProjectHandler) GetProjectByID(c fiber.Ctx) error {
	userID, err := ctxutil.GetUserID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	projectID := c.Params("id")
	if projectID == "" {
		return sendBadRequest(c, errors.ErrInvalidProjectID)
	}

	projectUUID, err := uuid.Parse(projectID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidProjectID)
	}

	project, err := h.projectService.GetProjectByID(c.Context(), userID, projectUUID)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	userID, err := ctxutil.GetUserID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var req dto.CreateProjectInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	project, err := h.projectService.CreateProject(c.Context(), userID, req)
	if err != nil {
		return sendBadRequest(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}

func (h *ProjectHandler) ActivateProject(c fiber.Ctx) error {
	userID, err := ctxutil.GetUserID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var req dto.ActivateProjectInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	projectUUID, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidProjectID)
	}

	activated, err := h.projectService.ActivateProject(c.Context(), userID, projectUUID)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"activated": activated,
	})
}
