package service

import (
	"context"
	"uplink-api/domain"
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type ConnectionService struct {
	connectionRepo *repository.ConnectionRepository
}

func NewConnectionService(connectionRepo *repository.ConnectionRepository) *ConnectionService {
	return &ConnectionService{
		connectionRepo: connectionRepo,
	}
}

func (s *ConnectionService) CreateConnection(ctx context.Context, req dto.CreateConnectionInput) (dto.ConnectionOutput, error) {

	workflowID, err := uuid.Parse(req.WorkflowID)
	if err != nil {
		return dto.ConnectionOutput{}, errors.ErrInvalidWorkflowID
	}

	fromStepID, err := uuid.Parse(req.From)
	if err != nil {
		return dto.ConnectionOutput{}, errors.ErrInvalidStepID
	}

	toStepID, err := uuid.Parse(req.To)
	if err != nil {
		return dto.ConnectionOutput{}, errors.ErrInvalidStepID
	}

	connections, err := s.connectionRepo.FindByFromTo(ctx, fromStepID, toStepID)
	if err != nil {
		return dto.ConnectionOutput{}, err
	}

	if len(connections) > 0 {
		return dto.ConnectionOutput{}, nil
	}

	connection := &domain.Connection{
		ID:         uuid.New(),
		WorkflowID: workflowID,
		FromStepID: fromStepID,
		ToStepID:   toStepID,
	}

	if err := s.connectionRepo.Create(ctx, connection); err != nil {
		return dto.ConnectionOutput{}, err
	}

	return dto.NewConnectionOutput(*connection), nil
}

func (s *ConnectionService) DeleteConnection(ctx context.Context, connectionID uuid.UUID) (dto.ConnectionOutput, error) {
	connection, err := s.connectionRepo.FindByID(ctx, connectionID)
	if err != nil {
		return dto.ConnectionOutput{}, errors.ErrInvalidRequestBody
	}

	if err := s.connectionRepo.Delete(ctx, connectionID); err != nil {
		return dto.ConnectionOutput{}, err
	}

	return dto.NewConnectionOutput(*connection), nil
}
