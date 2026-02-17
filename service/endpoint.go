package service

import (
	"context"
	"errors"
	"uplink-api/domain"
	"uplink-api/dto"
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
	project, err := s.projectRepo.FindByProjectID(ctx, projectID)
	if err != nil {
		return dto.EndpointOutput{}, errors.New("project not found")
	}

	endpoint := &domain.Endpoint{
		Name:      req.Name,
		ProjectID: project.ID,
		BaseURI:   req.BaseURI,
		Path:      req.Path,
		Method:    req.Method,
		Timeout:   req.Timeout,
	}

	if err := s.endpointRepo.Create(ctx, endpoint); err != nil {
		return dto.EndpointOutput{}, err
	}

	return dto.NewEndpointOutput(*endpoint), nil
}

func (s *EndpointService) GetEndpoints(ctx context.Context, projectID uuid.UUID) ([]dto.EndpointOutput, error) {
	project, err := s.projectRepo.FindByProjectID(ctx, projectID)
	if err != nil {
		return nil, errors.New("project not found")
	}

	endpoints, err := s.endpointRepo.FindAllByProjectID(ctx, project.ID)
	if err != nil {
		return nil, errors.New("endpoints not found")
	}

	output := dto.NewEndpointsOutput(endpoints)
	return output, nil
}
