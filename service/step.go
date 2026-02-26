package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type StepService struct {
	workflowRepo *repository.WorkflowRepository
	stepRepo     *repository.StepRepository
	endpointRepo *repository.EndpointRepository
}

func NewStepService(workflowRepo *repository.WorkflowRepository, stepRepo *repository.StepRepository, endpointRepo *repository.EndpointRepository) *StepService {
	return &StepService{
		workflowRepo: workflowRepo,
		stepRepo:     stepRepo,
		endpointRepo: endpointRepo,
	}
}

func (s *StepService) GetStepsByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
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

func (s *StepService) CreateStepByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, req dto.CreateStepInput) (dto.StepOutput, error) {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrWorkflowNotFound
	}

	endpointID, err := uuid.Parse(req.EndpointID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrInvalidEndpointID
	}

	endpoint, err := s.endpointRepo.FindByProjectIDAndEndpointID(ctx, projectID, endpointID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrEndpointNotFound
	}

	count, err := s.stepRepo.CountByWorkflowID(ctx, workflowID)
	if err != nil {
		return dto.StepOutput{}, err
	}

	step := &domain.Step{
		WorkflowID: workflowID,
		Position:   int(count + 1),
		EndpointID: endpoint.ID,
		Name:       req.Name,
	}

	if err := s.stepRepo.Create(ctx, step); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(*step), nil
}

func (s *StepService) UpdateStepPosition(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, req dto.UpdateStepPositionInput) error {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return errors.ErrWorkflowNotFound
	}

	for _, item := range req.Steps {
		itemID, err := uuid.Parse(item.StepID)
		if err != nil {
			return errors.ErrInvalidStepID
		}

		step, err := s.stepRepo.FindByWorkflowIDAndStepID(ctx, workflowID, itemID)
		if err != nil {
			return errors.ErrStepNotFound
		}

		step.Position = item.Position
		if err := s.stepRepo.Update(ctx, &step); err != nil {
			return err
		}
	}

	return nil
}
