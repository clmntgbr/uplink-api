package dto

import (
	"time"
	"uplink-api/domain"
)

type WorkflowOutput struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	StepsCount  int                `json:"stepsCount"`
	Steps       []StepOutput       `json:"steps"`
	Connections []ConnectionOutput `json:"connections"`
}

type CreateWorkflowInput struct {
	Name        string `json:"name" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"omitempty,min=2,max=255"`
}

type UpdateWorkflowInput struct {
	Name        string            `json:"name" validate:"required,min=2,max=255"`
	Description string            `json:"description" validate:"omitempty,min=2,max=255"`
	Steps       []UpdateStepInput `json:"steps" validate:"omitempty,dive"`
}

func NewWorkflowOutput(workflow domain.Workflow) WorkflowOutput {
	stepsCount := workflow.StepsCount
	if stepsCount == 0 {
		stepsCount = len(workflow.Steps)
	}
	return WorkflowOutput{
		ID:          workflow.ID.String(),
		Name:        workflow.Name,
		Description: workflow.Description,
		CreatedAt:   workflow.CreatedAt,
		UpdatedAt:   workflow.UpdatedAt,
		StepsCount:  stepsCount,
		Steps:       NewStepsOutput(workflow.Steps),
		Connections: NewConnectionsOutput(workflow.Connections),
	}
}

func NewWorkflowsOutput(workflows []domain.Workflow) []WorkflowOutput {
	outputs := make([]WorkflowOutput, len(workflows))
	for i, workflow := range workflows {
		outputs[i] = NewWorkflowOutput(workflow)
	}
	return outputs
}
