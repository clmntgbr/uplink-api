package dto

import (
	"time"
	"uplink-api/domain"

	"github.com/google/uuid"
)

type StepOutput struct {
	ID        string          `json:"id"`
	Position  int             `json:"position"`
	Endpoint  *EndpointOutput `json:"endpoint"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

type CreateStepInput struct {
	Position   int    `json:"position" validate:"required,min=1,max=1000,number"`
	EndpointID string `json:"endpointId" validate:"required,uuid"`
}

type UpdateStepPositionInput struct {
	Steps []StepPosition `json:"steps" validate:"required,min=1,max=1000,dive"`
}

type StepPosition struct {
	StepID   string `json:"stepId" validate:"required,uuid"`
	Position int    `json:"position" validate:"required,min=1,max=1000,number"`
}

func NewStepOutput(step domain.Step) StepOutput {
	var endpoint *EndpointOutput
	if step.Endpoint.ID != uuid.Nil {
		endpointOutput := NewEndpointOutput(step.Endpoint)
		endpoint = &endpointOutput
	}

	return StepOutput{
		ID:        step.ID.String(),
		Position:  step.Position,
		Endpoint:  endpoint,
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
