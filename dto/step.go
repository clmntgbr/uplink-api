package dto

import (
	"time"
	"uplink-api/domain"
)

type StepOutput struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateStepInput struct {
	Position   int    `json:"position" validate:"required,min=1,max=1000,number"`
	EndpointID string `json:"endpointId" validate:"required,uuid"`
}

func NewStepOutput(step domain.Step) StepOutput {
	return StepOutput{
		ID:        step.ID.String(),
		CreatedAt: step.CreatedAt,
		UpdatedAt: step.UpdatedAt,
	}
}

func NewStepsOutput(steps []domain.Step) []StepOutput {
	outputs := make([]StepOutput, len(steps))
	for i, step := range steps {
		outputs[i] = NewStepOutput(step)
	}
	return outputs
}
