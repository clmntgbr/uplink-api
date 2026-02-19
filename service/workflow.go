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
	stepRepo     *repository.StepRepository
}

func NewWorkflowService(workflowRepo *repository.WorkflowRepository, projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository, stepRepo *repository.StepRepository) *WorkflowService {
	return &WorkflowService{
		workflowRepo: workflowRepo,
		projectRepo:  projectRepo,
		userRepo:     userRepo,
		stepRepo:     stepRepo,
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

func (s *WorkflowService) UpdateWorkflow(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, req dto.UpdateWorkflowInput) (dto.WorkflowOutput, error) {
	workflow, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrWorkflowNotFound
	}

	workflow.Name = req.Name
	workflow.Description = req.Description

	if err := s.workflowRepo.Update(ctx, workflow); err != nil {
		return dto.WorkflowOutput{}, err
	}

	return dto.NewWorkflowOutput(*workflow), nil
}

func (s *WorkflowService) GetWorkflowByID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID) (dto.WorkflowOutput, error) {
	workflow, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrWorkflowNotFound
	}

	return dto.NewWorkflowOutput(*workflow), nil
}

func (s *WorkflowService) GetWorkflows(ctx context.Context, projectID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	workflows, total, err := s.workflowRepo.FindAllByProjectID(ctx, projectID, query)
	if err != nil {
		return dto.PaginateResponse{}, nil
	}

	outputs := dto.NewWorkflowsOutput(workflows)
	return dto.NewPaginateResponse(outputs, int(total), query), nil
}

func (s *WorkflowService) GetStepsByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrWorkflowNotFound
	}

	steps, total, err := s.stepRepo.FindAllByWorkflowID(ctx, workflowID, query)
	if err != nil {
		return dto.PaginateResponse{}, nil
	}

	outputs := dto.NewStepsOutput(steps)
	return dto.NewPaginateResponse(outputs, int(total), query), nil
}
