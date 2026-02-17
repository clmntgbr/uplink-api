package dto

import (
	"uplink-api/domain"
)

type EndpointOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	BaseURI string `json:"baseUri"`
	Path    string `json:"path"`
	Method  string `json:"method"`
	Timeout int    `json:"timeout"`
}

type CreateEndpointInput struct {
	Name    string `json:"name" validate:"required"`
	BaseURI string `json:"baseUri" validate:"required"`
	Path    string `json:"path" validate:"required"`
	Method  string `json:"method" validate:"required"`
	Timeout int    `json:"timeout" validate:"required"`
}

func NewEndpointOutput(endpoint domain.Endpoint) EndpointOutput {
	return EndpointOutput{
		ID:      endpoint.ID.String(),
		Name:    endpoint.Name,
		BaseURI: endpoint.BaseURI,
		Path:    endpoint.Path,
		Method:  endpoint.Method,
		Timeout: endpoint.Timeout,
	}
}

func NewEndpointsOutput(endpoints []domain.Endpoint) []EndpointOutput {
	outputs := make([]EndpointOutput, len(endpoints))
	for i, endpoint := range endpoints {
		outputs[i] = NewEndpointOutput(endpoint)
	}
	return outputs
}
