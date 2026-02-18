package handler

import (
	"uplink-api/errors"
	"uplink-api/validator"

	"github.com/gofiber/fiber/v3"
)

func bindAndValidate(c fiber.Ctx, req interface{}) error {
	if err := c.Bind().JSON(req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.ErrInvalidRequestBody.Error(),
		})
		return errors.ErrValidationFailed
	}

	if err := validator.ValidateStruct(req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.ErrInvalidRequestBody.Error(),
			"errors":  validator.FormatValidationErrors(err),
		})
		return errors.ErrValidationFailed
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

func sendNotFound(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func handleError(c fiber.Ctx, err error) error {
	switch err {
	case errors.ErrWorkflowNotFound, errors.ErrProjectNotFound, errors.ErrUserNotFound:
		return sendNotFound(c, err)
	case errors.ErrUserNotAuthenticated:
		return sendUnauthorized(c)
	case errors.ErrInvalidWorkflowID, errors.ErrInvalidProjectID, errors.ErrInvalidQueryParams, errors.ErrInvalidRequestBody:
		return sendBadRequest(c, err)
	default:
		return sendInternalError(c, err)
	}
}
