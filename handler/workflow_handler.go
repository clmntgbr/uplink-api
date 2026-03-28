package handler

import (
	"fmt"
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
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(workflow)
}

func (h *WorkflowHandler) UpdateWorkflow(c fiber.Ctx) error {
	var req dto.UpdateWorkflowInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	workflowUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidWorkflowID)
	if err != nil {
		return err
	}

	fmt.Println("workflowUUID", workflowUUID)

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	workflow, err := h.workflowService.UpdateWorkflow(c.Context(), activeProject.ID, workflowUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(workflow)
}

func (h *WorkflowHandler) GetWorkflowByID(c fiber.Ctx) error {
	workflowUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidWorkflowID)
	if err != nil {
		return err
	}

	project, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	workflow, err := h.workflowService.GetWorkflowByID(c.Context(), project.ID, workflowUUID)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflow)
}

func (h *WorkflowHandler) GetWorkflows(c fiber.Ctx) error {
	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	query, err := bindPaginateQuery(c)
	if err != nil {
		return err
	}

	workflows, err := h.workflowService.GetWorkflows(c.Context(), activeProject.ID, query)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflows)
}
