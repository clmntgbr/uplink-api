package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type WorkflowHandler struct {
	workflowService *service.WorkflowService
}

func NewWorkflowHandler(workflowService *service.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{
		workflowService: workflowService,
	}
}

func (h *WorkflowHandler) CreateWorkflow(c fiber.Ctx) error {
	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var req dto.CreateWorkflowInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	workflow, err := h.workflowService.CreateWorkflow(c.Context(), activeProject.ID, req)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(workflow)
}

func (h *WorkflowHandler) UpdateWorkflow(c fiber.Ctx) error {
	var req dto.UpdateWorkflowInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	workflowID := c.Params("id")
	if workflowID == "" {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	workflowUUID, err := uuid.Parse(workflowID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	workflow, err := h.workflowService.UpdateWorkflow(c.Context(), activeProject.ID, workflowUUID, req)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(workflow)
}

func (h *WorkflowHandler) GetWorkflowByID(c fiber.Ctx) error {
	workflowID := c.Params("id")
	if workflowID == "" {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	workflowUUID, err := uuid.Parse(workflowID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	project, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	workflow, err := h.workflowService.GetWorkflowByID(c.Context(), project.ID, workflowUUID)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflow)
}

func (h *WorkflowHandler) GetWorkflows(c fiber.Ctx) error {
	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var query dto.PaginateQuery
	if err := c.Bind().Query(&query); err != nil {
		return sendBadRequest(c, errors.ErrInvalidQueryParams)
	}

	query.Normalize()

	workflows, err := h.workflowService.GetWorkflows(c.Context(), activeProject.ID, query)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflows)
}
