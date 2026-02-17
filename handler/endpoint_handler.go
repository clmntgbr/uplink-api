package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/service"
	"uplink-api/validator"

	"github.com/gofiber/fiber/v3"
)

type EndpointHandler struct {
	endpointService *service.EndpointService
}

func NewEndpointHandler(endpointService *service.EndpointService) *EndpointHandler {
	return &EndpointHandler{
		endpointService: endpointService,
	}
}

func (h *EndpointHandler) CreateEndpoint(c fiber.Ctx) error {
	activeProjectID, err := ctxutil.GetActiveProjectID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var req dto.CreateEndpointInput

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validator.FormatValidationErrors(err),
		})
	}

	endpoint, err := h.endpointService.CreateEndpoint(c.Context(), activeProjectID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(endpoint)
}

func (h *EndpointHandler) GetEndpoints(c fiber.Ctx) error {
	activeProjectID, err := ctxutil.GetActiveProjectID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	endpoints, err := h.endpointService.GetEndpoints(c.Context(), activeProjectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(endpoints)
}
