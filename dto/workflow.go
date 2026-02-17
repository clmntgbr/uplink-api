package dto

import (
	"uplink-api/domain"
)

type WorkflowOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateWorkflowInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func NewWorkflowOutput(workflow domain.Workflow) WorkflowOutput {
	return WorkflowOutput{
		ID:          workflow.ID.String(),
		Name:        workflow.Name,
		Description: workflow.Description,
	}
}

func NewWorkflowsOutput(workflows []domain.Workflow) []WorkflowOutput {
	outputs := make([]WorkflowOutput, len(workflows))
	for i, workflow := range workflows {
		outputs[i] = NewWorkflowOutput(workflow)
	}
	return outputs
}
