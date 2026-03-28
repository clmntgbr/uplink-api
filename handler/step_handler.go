package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
)

type StepHandler struct {
	stepService *service.StepService
}

func NewStepHandler(stepService *service.StepService) *StepHandler {
	return &StepHandler{
		stepService: stepService,
	}
}

func (h *StepHandler) UpdateStep(c fiber.Ctx) error {
	var req dto.UpdateStepDetailsInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	stepUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidStepID)
	if err != nil {
		return err
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	step, err := h.stepService.UpdateStep(c.Context(), activeProject.ID, stepUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(step)
}
