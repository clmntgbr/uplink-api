package handler

import (
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
)

type ConnectionHandler struct {
	connectionService *service.ConnectionService
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		connectionService: connectionService,
	}
}

func (h *ConnectionHandler) CreateConnection(c fiber.Ctx) error {
	var req dto.CreateConnectionInput
	if err := bindAndValidate(c, &req); err != nil {
		return nil
	}

	connection, err := h.connectionService.CreateConnection(c.Context(), req)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(connection)
}

func (h *ConnectionHandler) DeleteConnection(c fiber.Ctx) error {
	connectionUUID, err := parseUUIDParam(c, "id", errors.ErrInvalidConnectionID)
	if err != nil {
		return err
	}

	connection, err := h.connectionService.DeleteConnection(c.Context(), connectionUUID)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(connection)
}
