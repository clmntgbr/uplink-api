package ctxutil

import (
	"uplink-api/domain"
	"uplink-api/errors"

	"github.com/gofiber/fiber/v3"
)

const (
	UserKey          = "user"
	ActiveProjectKey = "activeProject"
)

func GetUser(c fiber.Ctx) (*domain.User, error) {
	user, ok := c.Locals(UserKey).(*domain.User)
	if !ok {
		return nil, errors.ErrUserNotAuthenticated
	}
	return user, nil
}

func SetUser(c fiber.Ctx, user domain.User) {
	c.Locals(UserKey, &user)
}

func GetActiveProject(c fiber.Ctx) (*domain.Project, error) {
	project, ok := c.Locals(ActiveProjectKey).(*domain.Project)
	if !ok {
		return nil, errors.ErrProjectNotFound
	}
	return project, nil
}

func SetActiveProject(c fiber.Ctx, project domain.Project) {
	c.Locals(ActiveProjectKey, &project)
}
