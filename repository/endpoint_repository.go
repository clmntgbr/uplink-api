package repository

import (
	"context"

	"uplink-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *EndpointRepository {
	return &EndpointRepository{db: db}
}

func (r *EndpointRepository) Create(ctx context.Context, endpoint *domain.Endpoint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(endpoint).Error
	})
}

func (r *EndpointRepository) FindAllByProjectID(ctx context.Context, projectID uuid.UUID) ([]domain.Endpoint, error) {
	var endpoints []domain.Endpoint

	err := r.db.WithContext(ctx).
		Where("endpoints.project_id = ?", projectID).
		Find(&endpoints).Error

	if err != nil {
		return nil, err
	}

	return endpoints, nil
}
