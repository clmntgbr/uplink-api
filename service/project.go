package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
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

func (s *ProjectService) GetProjects(ctx context.Context, user *domain.User, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	projects, total, err := s.projectRepo.FindAllByUserID(ctx, user.ID, query)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrProjectsNotFound
	}

	outputs := dto.NewProjectsOutput(projects, user.ActiveProjectID)
	return dto.NewPaginateResponse(outputs, int(total), query), nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, user *domain.User, projectID uuid.UUID) (dto.ProjectOutput, error) {
	project, err := s.projectRepo.FindByUserIDAndProjectID(ctx, projectID, user.ID)
	if err != nil {
		return dto.ProjectOutput{}, errors.ErrProjectNotFound
	}

	return dto.NewProjectOutput(*project, user.ActiveProjectID), nil
}

func (s *ProjectService) CreateProject(ctx context.Context, user *domain.User, input dto.CreateProjectInput) (dto.ProjectOutput, error) {
	if err := s.projectRules.MaxProjectsPerUser(ctx, user.ID); err != nil {
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

	return dto.NewProjectOutput(*project, user.ID), nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, user *domain.User, projectID uuid.UUID, input dto.UpdateProjectInput) (dto.ProjectOutput, error) {
	project, err := s.projectRepo.FindByUserIDAndProjectID(ctx, projectID, user.ID)
	if err != nil {
		return dto.ProjectOutput{}, errors.ErrProjectNotFound
	}

	project.Name = input.Name

	if err := s.projectRepo.Update(ctx, project); err != nil {
		return dto.ProjectOutput{}, err
	}

	return dto.NewProjectOutput(*project, user.ActiveProjectID), nil
}

func (s *ProjectService) ActivateProject(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (domain.Project, error) {
	project, err := s.projectRepo.ActivateProject(ctx, userID, projectID)
	if err != nil {
		return domain.Project{}, err
	}

	return *project, nil
}
