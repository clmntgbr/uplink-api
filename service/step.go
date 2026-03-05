package service

import (
	"context"
	"encoding/json"
	"strings"
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
		WorkflowID:   workflowID,
		Position:     int(count + 1),
		EndpointID:   endpoint.ID,
		Name:         req.Name,
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

	if req.UpdateVariables != nil {
		if err := s.UpdateStepVariables(ctx, workflowID, req.UpdateVariables); err != nil {
			return dto.StepOutput{}, err
		}
	}

	step.Name = req.Name
	step.Header = req.Header
	step.Body = req.Body
	step.Query = req.Query
	step.SetVariables = req.SetVariables

	if err := s.stepRepo.Update(ctx, &step); err != nil {
		return dto.StepOutput{}, err
	}

	return dto.NewStepOutput(step), nil
}

func (s *StepService) UpdateStepVariables(ctx context.Context, workflowID uuid.UUID, updateVariables datatypes.JSON) error {
	steps, _, err := s.stepRepo.FindAllByWorkflowID(ctx, workflowID, dto.PaginateQuery{
		Limit: 1000,
		Page:  1,
	})

	if err != nil {
		return err
	}

	var variablesMappings []map[string]string
	if err := json.Unmarshal(updateVariables, &variablesMappings); err != nil {
		return err
	}

	for _, currentStep := range steps {
		updated := false

		headerStr := string(currentStep.Header)
		for _, mapping := range variablesMappings {
			for oldKey, newKey := range mapping {
				oldPattern := "{{" + oldKey + "}}"
				newPattern := "{{" + newKey + "}}"
				if strings.Contains(headerStr, oldPattern) {
					headerStr = strings.ReplaceAll(headerStr, oldPattern, newPattern)
					updated = true
				}
			}
		}
		if updated {
			currentStep.Header = datatypes.JSON(headerStr)
		}

		bodyStr := string(currentStep.Body)
		bodyUpdated := false
		for _, mapping := range variablesMappings {
			for oldKey, newKey := range mapping {
				oldPattern := "{{" + oldKey + "}}"
				newPattern := "{{" + newKey + "}}"
				if strings.Contains(bodyStr, oldPattern) {
					bodyStr = strings.ReplaceAll(bodyStr, oldPattern, newPattern)
					bodyUpdated = true
				}
			}
		}
		if bodyUpdated {
			currentStep.Body = datatypes.JSON(bodyStr)
			updated = true
		}

		queryStr := string(currentStep.Query)
		queryUpdated := false
		for _, mapping := range variablesMappings {
			for oldKey, newKey := range mapping {
				oldPattern := "{{" + oldKey + "}}"
				newPattern := "{{" + newKey + "}}"
				if strings.Contains(queryStr, oldPattern) {
					queryStr = strings.ReplaceAll(queryStr, oldPattern, newPattern)
					queryUpdated = true
				}
			}
		}
		if queryUpdated {
			currentStep.Query = datatypes.JSON(queryStr)
			updated = true
		}

		if updated {
			if err := s.stepRepo.Update(ctx, &currentStep); err != nil {
				return err
			}
		}
	}

	return nil
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
