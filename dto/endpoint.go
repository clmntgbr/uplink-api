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
	Header  map[string]any `json:"header"`
	Body    map[string]any `json:"body"`
	Query   map[string]any `json:"query"`
}

type CreateEndpointInput struct {
	Name    string `json:"name" validate:"required"`
	BaseURI string `json:"baseUri" validate:"required"`
	Path    string `json:"path" validate:"required"`
	Method  string `json:"method" validate:"required"`
	Timeout int    `json:"timeout" validate:"required"`
	Header  map[string]any `json:"header" validate:"required"`
	Body    map[string]any `json:"body" validate:"required"`
	Query   map[string]any `json:"query" validate:"required"`
}

func NewEndpointOutput(endpoint domain.Endpoint) EndpointOutput {
	return EndpointOutput{
		ID:      endpoint.ID.String(),
		Name:    endpoint.Name,
		BaseURI: endpoint.BaseURI,
		Path:    endpoint.Path,
		Method:  endpoint.Method,
		Timeout: endpoint.Timeout,
		Header:  endpoint.Header,
		Body:    endpoint.Body,
		Query:   endpoint.Query,
	}
}

func NewEndpointsOutput(endpoints []domain.Endpoint) []EndpointOutput {
	outputs := make([]EndpointOutput, len(endpoints))
	for i, endpoint := range endpoints {
		outputs[i] = NewEndpointOutput(endpoint)
	}
	return outputs
}
