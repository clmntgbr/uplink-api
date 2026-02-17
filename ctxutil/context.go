package ctxutil

import (
	"uplink-api/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	ActiveProjectIDKey = "activeProjectID"
	UserIDKey          = "userID"
	UserEmailKey       = "userEmail"
)

func GetActiveProjectID(c fiber.Ctx) (uuid.UUID, error) {
	activeProjectID, ok := c.Locals(ActiveProjectIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.ErrActiveProjectNotFound
	}
	return activeProjectID, nil
}

func GetUserID(c fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals(UserIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.ErrUserNotAuthenticated
	}
	return userID, nil
}
