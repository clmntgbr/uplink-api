package service

import (
	"context"
	"fmt"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type WorkflowService struct {
	workflowRepo   *repository.WorkflowRepository
	stepRepo       *repository.StepRepository
	connectionRepo *repository.ConnectionRepository
}

func NewWorkflowService(workflowRepo *repository.WorkflowRepository, stepRepo *repository.StepRepository, connectionRepo *repository.ConnectionRepository) *WorkflowService {
	return &WorkflowService{
		workflowRepo:   workflowRepo,
		stepRepo:       stepRepo,
		connectionRepo: connectionRepo,
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

	if err := s.upsertSteps(ctx, workflow.ID, req.Steps); err != nil {
		return dto.WorkflowOutput{}, err
	}

	if err := s.upsertConnections(ctx, workflow.ID, req.Connections); err != nil {
		return dto.WorkflowOutput{}, err
	}

	updatedWorkflow, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrWorkflowNotFound
	}

	return dto.NewWorkflowOutput(*updatedWorkflow), nil
}

func (s *WorkflowService) upsertSteps(ctx context.Context, workflowID uuid.UUID, stepsInput []dto.UpdateStepInput) error {
	if len(stepsInput) == 0 {
		return s.stepRepo.DeleteByWorkflowID(ctx, workflowID)
	}

	if err := s.stepRepo.DeleteByWorkflowID(ctx, workflowID); err != nil {
		return err
	}

	stepsToCreate := make([]*domain.Step, 0, len(stepsInput))
	for _, stepInput := range stepsInput {
		endpointUUID, err := uuid.Parse(stepInput.EndpointID)
		if err != nil {
			return errors.ErrInvalidRequestBody
		}

		index := 0
		fmt.Sscanf(stepInput.Index, "%d", &index)

		step := &domain.Step{
			Name:        stepInput.Name,
			Description: stepInput.Description,
			Position:    domain.Position{X: stepInput.Position.X, Y: stepInput.Position.Y},
			Index:       index,
			EndpointID:  endpointUUID,
			WorkflowID:  workflowID,
		}

		if stepInput.ID != "" {
			stepUUID, err := uuid.Parse(stepInput.ID)
			if err == nil {
				step.ID = stepUUID
			}
		}

		stepsToCreate = append(stepsToCreate, step)
	}

	return s.stepRepo.CreateBatch(ctx, stepsToCreate)
}

func (s *WorkflowService) upsertConnections(ctx context.Context, workflowID uuid.UUID, connectionsInput []dto.UpdateConnectionInput) error {
	if len(connectionsInput) == 0 {
		return s.connectionRepo.DeleteByWorkflowID(ctx, workflowID)
	}

	if err := s.connectionRepo.DeleteByWorkflowID(ctx, workflowID); err != nil {
		return err
	}

	connectionsToCreate := make([]*domain.Connection, 0, len(connectionsInput))
	for _, connInput := range connectionsInput {
		fromUUID, err := uuid.Parse(connInput.From)
		if err != nil {
			return errors.ErrInvalidRequestBody
		}

		toUUID, err := uuid.Parse(connInput.To)
		if err != nil {
			return errors.ErrInvalidRequestBody
		}

		connection := &domain.Connection{
			FromStepID: fromUUID,
			ToStepID:   toUUID,
			WorkflowID: workflowID,
		}

		if connInput.ID != "" {
			connUUID, err := uuid.Parse(connInput.ID)
			if err == nil {
				connection.ID = connUUID
			}
		}

		connectionsToCreate = append(connectionsToCreate, connection)
	}

	return s.connectionRepo.CreateBatch(ctx, connectionsToCreate)
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
