package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/dto"
	"uplink-api/errors"
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
	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	var req dto.CreateEndpointInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	endpoint, err := h.endpointService.CreateEndpoint(c.Context(), activeProject.ID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(endpoint)
}

func (h *EndpointHandler) GetEndpoints(c fiber.Ctx) error {
	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	query, err := bindPaginateQuery(c)
	if err != nil {
		return err
	}

	endpoints, err := h.endpointService.GetEndpoints(c.Context(), activeProject.ID, query)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(endpoints)
}

func (h *EndpointHandler) GetEndpointByID(c fiber.Ctx) error {
	endpointUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidEndpointID)
	if err != nil {
		return err
	}

	project, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	endpoint, err := h.endpointService.GetEndpointByID(c.Context(), project.ID, endpointUUID)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(endpoint)
}

func (h *EndpointHandler) UpdateEndpoint(c fiber.Ctx) error {
	var req dto.UpdateEndpointInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	endpointUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidEndpointID)
	if err != nil {
		return err
	}

	activeProject, err := ctxutil.GetActiveProject(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	endpoint, err := h.endpointService.UpdateEndpoint(c.Context(), activeProject.ID, endpointUUID, req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(endpoint)
}
