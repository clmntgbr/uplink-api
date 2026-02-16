package dto

import (
	"uplink-api/domain"
)

type ProjectOutput struct {
	ID       string `json:"id"`
	Name    string `json:"name"`
}

func NewProjectOutput(project domain.Project) ProjectOutput {
	return ProjectOutput{
		ID:       project.ID.String(),
		Name:     project.Name,
	}
}

func NewProjectsOutput(projects []domain.Project) []ProjectOutput {
	outputs := make([]ProjectOutput, len(projects))
	for i, project := range projects {
		outputs[i] = NewProjectOutput(project)
	}
	return outputs
}

