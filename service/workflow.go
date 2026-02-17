package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type WorkflowService struct {
	workflowRepo *repository.WorkflowRepository
	projectRepo  *repository.ProjectRepository
	userRepo     *repository.UserRepository
}

func NewWorkflowService(workflowRepo *repository.WorkflowRepository, projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository) *WorkflowService {
	return &WorkflowService{
		workflowRepo: workflowRepo,
		projectRepo:  projectRepo,
		userRepo:     userRepo,
	}
}

func (s *WorkflowService) CreateWorkflow(ctx context.Context, projectID uuid.UUID, req dto.CreateWorkflowInput) (dto.WorkflowOutput, error) {
	project, err := s.projectRepo.FindByProjectID(ctx, projectID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrProjectNotFound
	}

	workflow := &domain.Workflow{
		Name:        req.Name,
		ProjectID:   project.ID,
		Description: req.Description,
	}

	if err := s.workflowRepo.Create(ctx, workflow); err != nil {
		return dto.WorkflowOutput{}, err
	}

	return dto.NewWorkflowOutput(*workflow), nil
}

func (s *WorkflowService) GetWorkflows(ctx context.Context, projectID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	project, err := s.projectRepo.FindByProjectID(ctx, projectID)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrProjectNotFound
	}

	workflows, total, err := s.workflowRepo.FindAllByProjectID(ctx, project.ID, query)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrWorkflowsNotFound
	}

	return dto.NewPaginateResponse(workflows, int(total), query), nil
}
