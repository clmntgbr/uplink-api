package handler

import (
	"uplink-api/dto"
	"uplink-api/service"

	"github.com/gofiber/fiber/v3"
)

type AuthenticateHandler struct {
	authenticateService *service.AuthenticateService
}

func NewAuthenticateHandler(authService *service.AuthenticateService) *AuthenticateHandler {
	return &AuthenticateHandler{
		authenticateService: authService,
	}
}

func (h *AuthenticateHandler) Login(c fiber.Ctx) error {
	var req dto.LoginInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	loginOutput, err := h.authenticateService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(loginOutput)
}

func (h *AuthenticateHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterInput
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	registerOutput, err := h.authenticateService.Register(c, req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(registerOutput)
}
