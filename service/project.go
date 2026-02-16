package service

import (
	"context"
	"errors"
	"uplink-api/dto"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
	userRepo    *repository.UserRepository
}

func NewProjectService(projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (s *ProjectService) GetProjects(ctx context.Context, userID uuid.UUID) ([]dto.ProjectOutput, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	projects, err := s.projectRepo.FindAllByUserID(ctx, user)
	if err != nil {
		return nil, errors.New("projects not found")
	}

	output := dto.NewProjectsOutput(projects, user.ActiveProjectID)
	return output, nil
}

func (s *ProjectService) GetProjectById(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (dto.ProjectOutput, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return 	dto.ProjectOutput{}, errors.New("user not found")
	}

	project, err := s.projectRepo.FindByID(ctx, projectID, user)
	if err != nil {
		return dto.ProjectOutput{}, errors.New("project not found")
	}

	output := dto.NewProjectOutput(*project, user.ActiveProjectID)
	return output, nil
}
