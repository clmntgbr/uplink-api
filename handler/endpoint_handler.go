package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/service"

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
		return sendUnauthorized(c)
	}

	var req dto.CreateEndpointInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	endpoint, err := h.endpointService.CreateEndpoint(c.Context(), activeProjectID, req)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.JSON(endpoint)
}

func (h *EndpointHandler) GetEndpoints(c fiber.Ctx) error {
	activeProjectID, err := ctxutil.GetActiveProjectID(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	endpoints, err := h.endpointService.GetEndpoints(c.Context(), activeProjectID)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.JSON(endpoints)
}
