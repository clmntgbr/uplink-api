package service

import (
	"context"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type StepService struct {
	stepRepo     *repository.StepRepository
	workflowRepo *repository.WorkflowRepository
}

func NewStepService(stepRepo *repository.StepRepository, workflowRepo *repository.WorkflowRepository) *StepService {
	return &StepService{
		stepRepo:     stepRepo,
		workflowRepo: workflowRepo,
	}
}

func (s *StepService) UpdateStep(ctx context.Context, projectID uuid.UUID, stepID uuid.UUID, req dto.UpdateStepDetailsInput) (dto.StepOutput, error) {
	workflowUUID, err := uuid.Parse(req.WorkflowID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrInvalidWorkflowID
	}

	workflow, err := s.workflowRepo.FindByProjectIDAndWorkflowID(ctx, projectID, workflowUUID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrWorkflowNotFound
	}

	step, err := s.stepRepo.FindByIDAndWorkflowID(ctx, stepID, workflow.ID)
	if err != nil {
		return dto.StepOutput{}, errors.ErrStepNotFound
	}

	step.Name = req.Name
	step.Description = req.Description

	if req.EndpointID != "" {
		endpointUUID, err := uuid.Parse(req.EndpointID)
		if err != nil {
			return dto.StepOutput{}, errors.ErrInvalidRequestBody
		}
		step.EndpointID = endpointUUID
	}

	if err := s.stepRepo.Update(ctx, step); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(*step), nil
}
