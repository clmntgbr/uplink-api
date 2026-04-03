package repository

import (
	"context"

	"uplink-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConnectionRepository struct {
	db *gorm.DB
}

func NewConnectionRepository(db *gorm.DB) *ConnectionRepository {
	return &ConnectionRepository{db: db}
}

func (r *ConnectionRepository) Create(ctx context.Context, connection *domain.Connection) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(connection).Error
	})
}

func (r *ConnectionRepository) Delete(ctx context.Context, connectionID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Connection{}, connectionID).Error
}

func (r *ConnectionRepository) FindByFromTo(ctx context.Context, fromStepID uuid.UUID, toStepID uuid.UUID) ([]domain.Connection, error) {
	var connections []domain.Connection
	err := r.db.WithContext(ctx).Where("from_step_id = ? AND to_step_id = ?", fromStepID, toStepID).Find(&connections).Error
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func (r *ConnectionRepository) FindByID(ctx context.Context, connectionID uuid.UUID) (*domain.Connection, error) {
	var connection domain.Connection
	err := r.db.WithContext(ctx).Where("id = ?", connectionID).First(&connection).Error
	if err != nil {
		return nil, err
	}
	return &connection, nil
}
