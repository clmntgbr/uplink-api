package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
	"gorm.io/datatypes"
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
		WorkflowID:   workflowID,
		Position:     int(count + 1),
		EndpointID:   endpoint.ID,
		Endpoint:     endpoint,
		Name:         req.Name,
		URL:          req.URL,
		Header:       req.Header,
		Body:         req.Body,
		Query:        req.Query,
		SetVariables: req.SetVariables,
	}

	if err := s.stepRepo.Create(ctx, step); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(*step), nil
}

func (s *StepService) UpdateStepByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, stepID uuid.UUID, req dto.UpdateStepInput) (dto.StepOutput, error) {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrWorkflowNotFound
	}

	step, err := s.stepRepo.FindByWorkflowIDAndStepID(ctx, workflowID, stepID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrStepNotFound
	}

	step.Name = req.Name
	step.URL = req.URL
	step.Header = req.Header
	step.Body = req.Body
	step.Query = req.Query

	if err := s.stepRepo.Update(ctx, &step); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(step), nil
}

func (s *StepService) DeleteStepByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, stepID uuid.UUID) error {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return errors.ErrWorkflowNotFound
	}

	step, err := s.stepRepo.FindByWorkflowIDAndStepID(ctx, workflowID, stepID)
	if err != nil {
		return errors.ErrStepNotFound
	}

	if err := s.stepRepo.Delete(ctx, &step); err != nil {
		return err
	}

	return nil
}

func (s *StepService) DuplicateStepByWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, stepID uuid.UUID) (dto.StepOutput, error) {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrWorkflowNotFound
	}

	step, err := s.stepRepo.FindByWorkflowIDAndStepID(ctx, workflowID, stepID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrStepNotFound
	}

	count, err := s.stepRepo.CountByWorkflowID(ctx, workflowID)
	if err != nil {
		return dto.StepOutput{}, err
	}

	dupeStep := &domain.Step{
		WorkflowID:   workflowID,
		Position:     int(count + 1),
		EndpointID:   step.EndpointID,
		Name:         step.Name + " (Copy)",
		Header:       step.Header,
		Body:         step.Body,
		Query:        step.Query,
		SetVariables: datatypes.JSON(nil),
	}

	if err := s.stepRepo.Create(ctx, dupeStep); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(*dupeStep), nil
}

func (s *StepService) UpdateStepPosition(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID, req dto.UpdateStepPositionInput) (dto.WorkflowOutput, error) {
	_, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrWorkflowNotFound
	}

	for _, item := range req.Steps {
		itemID, err := uuid.Parse(item.StepID)
		if err != nil {
			return dto.WorkflowOutput{}, errors.ErrInvalidStepID
		}

		step, err := s.stepRepo.FindByWorkflowIDAndStepID(ctx, workflowID, itemID)
		if err != nil {
			return dto.WorkflowOutput{}, errors.ErrStepNotFound
		}

		step.Position = item.Position
		if err := s.stepRepo.Update(ctx, &step); err != nil {
			return dto.WorkflowOutput{}, err
		}
	}

	workflow, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowID)
	if err != nil {
		return dto.WorkflowOutput{}, errors.ErrWorkflowNotFound
	}

	return dto.NewWorkflowOutput(*workflow), nil
}
