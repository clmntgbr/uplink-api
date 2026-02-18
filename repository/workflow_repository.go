package repository

import (
	"context"

	"uplink-api/domain"
	"uplink-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowRepository struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) *WorkflowRepository {
	return &WorkflowRepository{db: db}
}

func (r *WorkflowRepository) Create(ctx context.Context, workflow *domain.Workflow) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(workflow).Error
	})
}

func (r *WorkflowRepository) Update(ctx context.Context, workflow *domain.Workflow) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(workflow).Updates(workflow).Error
	})
}

func (r *WorkflowRepository) FindAllByProjectID(ctx context.Context, projectID uuid.UUID, q dto.PaginateQuery) ([]domain.Workflow, int64, error) {
	var workflows []domain.Workflow
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Workflow{}).
		Where("project_id = ?", projectID)

	if q.Search != "" {
		db = db.Where("name ILIKE ?", "%"+q.Search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	if q.SortBy != "" {
		sortBy = q.SortBy
	}

	orderBy := "desc"
	if q.OrderBy == "asc" {
		orderBy = "asc"
	}

	err := db.
		Order(sortBy + " " + orderBy).
		Limit(q.Limit).
		Offset(q.Offset()).
		Find(&workflows).Error

	if err != nil {
		return nil, 0, err
	}

	return workflows, total, nil
}

func (r *WorkflowRepository) FindByProjectIDAndWorkflowID(ctx context.Context, projectID uuid.UUID, workflowID uuid.UUID) (*domain.Workflow, error) {
	var workflow domain.Workflow

	err := r.db.WithContext(ctx).
		Where("project_id = ? AND id = ?", projectID, workflowID).
		First(&workflow).Error

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}
