package handler

import (
	"uplink-api/validator"

	"github.com/gofiber/fiber/v3"
)

func bindAndValidate(c fiber.Ctx, req interface{}) error {
	if err := c.Bind().JSON(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validator.FormatValidationErrors(err),
		})
	}

	return nil
}

func sendUnauthorized(c fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func sendInternalError(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func sendBadRequest(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}
