package service

import (
	"context"
	"errors"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/repository"
	"uplink-api/rules"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectRepo  *repository.ProjectRepository
	userRepo     *repository.UserRepository
	projectRules *rules.ProjectRules
}

func NewProjectService(projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository, projectRules *rules.ProjectRules) *ProjectService {
	return &ProjectService{
		projectRepo:  projectRepo,
		userRepo:     userRepo,
		projectRules: projectRules,
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

func (s *ProjectService) GetProjectByID(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (dto.ProjectOutput, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return dto.ProjectOutput{}, errors.New("user not found")
	}

	project, err := s.projectRepo.FindByUserIDAndProjectID(ctx, projectID, user)
	if err != nil {
		return dto.ProjectOutput{}, errors.New("project not found")
	}

	output := dto.NewProjectOutput(*project, user.ActiveProjectID)
	return output, nil
}

func (s *ProjectService) CreateProject(ctx context.Context, userID uuid.UUID, input dto.CreateProjectInput) (dto.ProjectOutput, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return dto.ProjectOutput{}, errors.New("user not found")
	}

	if err := s.projectRules.MaxProjectsPerUser(ctx, userID); err != nil {
		return dto.ProjectOutput{}, err
	}

	project := &domain.Project{
		Name: input.Name,
		Users: []domain.User{
			{
				ID: user.ID,
			},
		},
	}

	if err := s.projectRepo.Create(ctx, project); err != nil {
		return dto.ProjectOutput{}, err
	}

	return dto.NewProjectOutput(*project, userID), nil
}

func (s *ProjectService) ActivateProject(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (bool, error) {
	err := s.projectRepo.ActivateProject(ctx, userID, projectID)
	if err != nil {
		return false, err
	}

	return true, nil
}
