package dto

import (
	"time"
	"uplink-api/domain"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type StepOutput struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Position     int             `json:"position"`
	Endpoint     *EndpointOutput `json:"endpoint"`
	Header       datatypes.JSON  `json:"header"`
	Body         datatypes.JSON  `json:"body"`
	Query        datatypes.JSON  `json:"query"`
	SetVariables datatypes.JSON  `json:"setVariables"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}

type CreateStepInput struct {
	Name       string `json:"name" validate:"required,min=2,max=255"`
	EndpointID string `json:"endpointId" validate:"required,uuid"`

	Header       datatypes.JSON `json:"header" validate:"required,json"`
	Body         datatypes.JSON `json:"body" validate:"required,json"`
	Query        datatypes.JSON `json:"query" validate:"required,json"`
	SetVariables datatypes.JSON `json:"setVariables" validate:"required,json"`
}

type UpdateStepInput struct {
	Name string `json:"name" validate:"required,min=2,max=255"`

	Header       datatypes.JSON `json:"header" validate:"required,json"`
	Body         datatypes.JSON `json:"body" validate:"required,json"`
	Query        datatypes.JSON `json:"query" validate:"required,json"`
	SetVariables datatypes.JSON `json:"setVariables" validate:"required,json"`
}

type UpdateStepPositionInput struct {
	Steps []StepPosition `json:"steps" validate:"required,min=1,max=1000,dive"`
}

type StepPosition struct {
	StepID       string `json:"stepId" validate:"required,uuid"`
	EndpointName string `json:"endpointName" validate:"required,min=2,max=255"`
	Position     int    `json:"position" validate:"required,min=1,max=1000,number"`
}

func NewStepOutput(step domain.Step) StepOutput {
	var endpoint *EndpointOutput
	if step.Endpoint.ID != uuid.Nil {
		endpointOutput := NewEndpointOutput(step.Endpoint)
		endpoint = &endpointOutput
	}

	return StepOutput{
		ID:           step.ID.String(),
		Name:         step.Name,
		Position:     step.Position,
		Endpoint:     endpoint,
		Header:       step.Header,
		Body:         step.Body,
		Query:        step.Query,
		SetVariables: step.SetVariables,
		CreatedAt:    step.CreatedAt,
		UpdatedAt:    step.UpdatedAt,
	}
}

func NewStepsOutput(steps []domain.Step) []StepOutput {
	outputs := make([]StepOutput, len(steps))
	for i, step := range steps {
		outputs[i] = NewStepOutput(step)
	}
	return outputs
}
