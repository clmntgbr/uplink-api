package handler

import (
	"uplink-api/ctxutil"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(c fiber.Ctx) error {
	user, err := ctxutil.GetUser(c)
	if err != nil {
		return sendUnauthorized(c)
	}

	output, err := h.userService.GetUser(user)
	if err != nil {
		return sendInternalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
