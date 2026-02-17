package rules

import (
	"context"
	"uplink-api/domain"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type ProjectRules struct {
	projectRepo *repository.ProjectRepository
}

func NewProjectRules(projectRepo *repository.ProjectRepository) *ProjectRules {
	return &ProjectRules{
		projectRepo: projectRepo,
	}
}
func (p *ProjectRules) MaxProjectsPerUser(ctx context.Context, userID uuid.UUID) error {
	count, err := p.projectRepo.CountProjectsByUserID(ctx, userID)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if count >= 10 {
		return domain.ErrMaxProjectsReached
	}

	return nil
}
