package dto

import (
	"uplink-api/domain"

	"github.com/google/uuid"
)

type ProjectOutput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

type CreateProjectInput struct {
	Name string `json:"name" validate:"required"`
}

type UpdateProjectInput struct {
	Name string `json:"name" validate:"required"`
}

type ActivateProjectInput struct {
	ProjectID string `json:"projectId" validate:"required"`
}

func NewProjectOutput(project domain.Project, activeProjectID uuid.UUID) ProjectOutput {
	isActive := false
	if activeProjectID != uuid.Nil && activeProjectID == project.ID {
		isActive = true
	}

	return ProjectOutput{
		ID:       project.ID.String(),
		Name:     project.Name,
		IsActive: isActive,
	}
}

func NewProjectsOutput(projects []domain.Project, activeProjectID uuid.UUID) []ProjectOutput {
	outputs := make([]ProjectOutput, len(projects))
	for i, project := range projects {
		outputs[i] = NewProjectOutput(project, activeProjectID)
	}
	return outputs
}
