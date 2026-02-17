package handler

import (
	"uplink-api/errors"
	"uplink-api/validator"

	"github.com/gofiber/fiber/v3"
)

func bindAndValidate(c fiber.Ctx, req interface{}) error {
	if err := c.Bind().JSON(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.ErrInvalidRequestBody.Error(),
		})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.ErrInvalidRequestBody.Error(),
			"errors":  validator.FormatValidationErrors(err),
		})
	}

	return nil
}

func sendUnauthorized(c fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": errors.ErrUserNotAuthenticated.Error(),
	})
}

func sendInternalError(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func sendBadRequest(c fiber.Ctx, message error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message.Error(),
	})
}
