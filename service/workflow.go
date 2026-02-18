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
	workflow := &domain.Workflow{
		Name:        req.Name,
		ProjectID:   projectID,
		Description: req.Description,
	}

	if err := s.workflowRepo.Create(ctx, workflow); err != nil {
		return dto.WorkflowOutput{}, err
	}

	return dto.NewWorkflowOutput(*workflow), nil
}

func (s *WorkflowService) GetWorkflows(ctx context.Context, projectID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	workflows, total, err := s.workflowRepo.FindAllByProjectID(ctx, projectID, query)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrWorkflowsNotFound
	}

	outputs := dto.NewWorkflowsOutput(workflows)
	return dto.NewPaginateResponse(outputs, int(total), query), nil
}
