package dto

import (
	"time"
	"uplink-api/domain"
)

type ConnectionOutput struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateConnectionInput struct {
	WorkflowID string `json:"workflowId" validate:"required,uuid"`
	From       string `json:"from" validate:"required,uuid"`
	To         string `json:"to" validate:"required,uuid"`
}

type DeleteConnectionInput struct {
	WorkflowID string `json:"workflowId" validate:"required,uuid"`
	ID         string `json:"id" validate:"required,uuid"`
}

func NewConnectionOutput(connection domain.Connection) ConnectionOutput {
	return ConnectionOutput{
		ID:        connection.ID.String(),
		From:      connection.FromStepID.String(),
		To:        connection.ToStepID.String(),
		CreatedAt: connection.CreatedAt,
		UpdatedAt: connection.UpdatedAt,
	}
}

func NewConnectionsOutput(connections []domain.Connection) []ConnectionOutput {
	outputs := make([]ConnectionOutput, len(connections))
	for i, connection := range connections {
		outputs[i] = NewConnectionOutput(connection)
	}
	return outputs
}
