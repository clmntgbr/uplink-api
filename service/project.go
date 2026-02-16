package service

import (
	"context"
	"errors"
	"uplink-api/config"
	"uplink-api/dto"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
	config      *config.Config
}

func NewProjectService(projectRepo *repository.ProjectRepository, cfg *config.Config) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		config:      cfg,
	}
}

func (s *ProjectService) GetProjects(ctx context.Context, userID uuid.UUID) ([]dto.ProjectOutput, error) {
	projects, err := s.projectRepo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("projects not found")
	}

	output := dto.NewProjectsOutput(projects)
	return output, nil
}
