package ctxutil

import (
	"uplink-api/domain"
	"uplink-api/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	ActiveProjectIDKey = "activeProjectID"
	UserKey            = "user"
)

func GetActiveProjectID(c fiber.Ctx) (uuid.UUID, error) {
	activeProjectID, ok := c.Locals(ActiveProjectIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.ErrActiveProjectNotFound
	}
	return activeProjectID, nil
}

func GetUser(c fiber.Ctx) (*domain.User, error) {
	user, ok := c.Locals(UserKey).(*domain.User)
	if !ok {
		return nil, errors.ErrUserNotAuthenticated
	}
	return user, nil
}
