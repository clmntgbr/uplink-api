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
	stepService     *service.StepService
}

func NewWorkflowHandler(workflowService *service.WorkflowService, stepService *service.StepService) *WorkflowHandler {
	return &WorkflowHandler{
		workflowService: workflowService,
		stepService:     stepService,
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
		return handleError(c, err)
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
		return handleError(c, err)
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
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflows)
}

func (h *WorkflowHandler) CreateStepByWorkflowID(c fiber.Ctx) error {
	var req dto.CreateStepInput
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

	step, err := h.stepService.CreateStepByWorkflowID(c.Context(), activeProject.ID, workflowUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(step)
}

func (h *WorkflowHandler) UpdateStepByWorkflowID(c fiber.Ctx) error {
	var req dto.UpdateStepInput
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

	stepID := c.Params("stepId")
	if stepID == "" {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	stepUUID, err := uuid.Parse(stepID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	step, err := h.stepService.UpdateStepByWorkflowID(c.Context(), activeProject.ID, workflowUUID, stepUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(step)
}

func (h *WorkflowHandler) DeleteStepByWorkflowID(c fiber.Ctx) error {
	workflowID := c.Params("id")
	if workflowID == "" {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	workflowUUID, err := uuid.Parse(workflowID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	stepID := c.Params("stepId")
	if stepID == "" {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	stepUUID, err := uuid.Parse(stepID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	err = h.stepService.DeleteStepByWorkflowID(c.Context(), activeProject.ID, workflowUUID, stepUUID)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Step deleted successfully",
	})
}

func (h *WorkflowHandler) DuplicateStepByWorkflowID(c fiber.Ctx) error {
	workflowID := c.Params("id")
	if workflowID == "" {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	workflowUUID, err := uuid.Parse(workflowID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidWorkflowID)
	}

	stepID := c.Params("stepId")
	if stepID == "" {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	stepUUID, err := uuid.Parse(stepID)
	if err != nil {
		return sendBadRequest(c, errors.ErrInvalidStepID)
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	step, err := h.stepService.DuplicateStepByWorkflowID(c.Context(), activeProject.ID, workflowUUID, stepUUID)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(step)
}

func (h *WorkflowHandler) UpdateReorderSteps(c fiber.Ctx) error {
	var req dto.UpdateReorderStepsInput
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

	workflow, err := h.stepService.UpdateReorderSteps(c.Context(), activeProject.ID, workflowUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(workflow)
}
