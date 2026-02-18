package repository

import (
	"context"

	"uplink-api/domain"
	"uplink-api/dto"

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

func (r *EndpointRepository) FindAllByProjectID(ctx context.Context, projectID uuid.UUID, q dto.PaginateQuery) ([]domain.Endpoint, int64, error) {
	var endpoints []domain.Endpoint

	db := r.db.WithContext(ctx).Model(&domain.Endpoint{}).
		Where("project_id = ?", projectID)

	if q.Search != "" {
		db = db.Where("name ILIKE ?", "%"+q.Search+"%")
	}

	db, total, err := Paginate(db, q)
	if err != nil {
		return nil, 0, err
	}

	err = db.Find(&endpoints).Error
	if err != nil {
		return nil, 0, err
	}

	return endpoints, total, nil
}
