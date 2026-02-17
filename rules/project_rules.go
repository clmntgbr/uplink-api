package rules

import (
	"context"
	"uplink-api/domain"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type ProjectRules struct {
	userRepo *repository.UserRepository
}

func NewProjectRules(userRepo *repository.UserRepository) *ProjectRules {
	return &ProjectRules{
		userRepo: userRepo,
	}
}
func (p *ProjectRules) MaxProjectsPerUser(ctx context.Context, userID uuid.UUID) error {
	count, err := p.userRepo.CountProjectsByUserID(ctx, userID)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if count >= 10 {
		return domain.ErrMaxProjectsReached
	}

	return nil
}
