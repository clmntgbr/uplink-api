package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
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
	activeProjectID, err := ctxutil.GetActiveProjectID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var req dto.CreateWorkflowInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	workflow, err := h.workflowService.CreateWorkflow(c.Context(), activeProjectID, req)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(workflow)
}

func (h *WorkflowHandler) GetWorkflows(c fiber.Ctx) error {
	activeProjectID, err := ctxutil.GetActiveProjectID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var query dto.PaginateQuery
	if err := c.Bind().Query(&query); err != nil {
		return sendBadRequest(c, errors.ErrInvalidQueryParams)
	}

	query.Normalize()

	workflows, err := h.workflowService.GetWorkflows(c.Context(), activeProjectID, query)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflows)
}
