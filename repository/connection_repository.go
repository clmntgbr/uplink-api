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

func (r *ConnectionRepository) CreateBatch(ctx context.Context, connections []*domain.Connection) error {
	if len(connections) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&connections).Error
	})
}

func (r *ConnectionRepository) DeleteByWorkflowID(ctx context.Context, workflowID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("workflow_id = ?", workflowID).Delete(&domain.Connection{}).Error
	})
}
