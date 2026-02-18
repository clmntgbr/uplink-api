package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type EndpointService struct {
	endpointRepo *repository.EndpointRepository
	projectRepo  *repository.ProjectRepository
	userRepo     *repository.UserRepository
}

func NewEndpointService(endpointRepo *repository.EndpointRepository, projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository) *EndpointService {
	return &EndpointService{
		endpointRepo: endpointRepo,
		projectRepo:  projectRepo,
		userRepo:     userRepo,
	}
}

func (s *EndpointService) CreateEndpoint(ctx context.Context, projectID uuid.UUID, req dto.CreateEndpointInput) (dto.EndpointOutput, error) {
	endpoint := &domain.Endpoint{
		Name:      req.Name,
		ProjectID: projectID,
		BaseURI:   req.BaseURI,
		Path:      req.Path,
		Method:    req.Method,
		Timeout:   req.Timeout,
		Header:    req.Header,
		Body:      req.Body,
		Query:     req.Query,
	}

	if err := s.endpointRepo.Create(ctx, endpoint); err != nil {
		return dto.EndpointOutput{}, err
	}

	return dto.NewEndpointOutput(*endpoint), nil
}

func (s *EndpointService) GetEndpoints(ctx context.Context, projectID uuid.UUID, query dto.PaginateQuery) (dto.PaginateResponse, error) {
	endpoints, total, err := s.endpointRepo.FindAllByProjectID(ctx, projectID, query)
	if err != nil {
		return dto.PaginateResponse{}, errors.ErrEndpointsNotFound
	}

	outputs := dto.NewEndpointsOutput(endpoints)
	return dto.NewPaginateResponse(outputs, int(total), query), nil
}
