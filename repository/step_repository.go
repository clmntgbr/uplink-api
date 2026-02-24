package repository

import (
	"context"

	"uplink-api/domain"
	"uplink-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) *StepRepository {
	return &StepRepository{db: db}
}

func (r *StepRepository) Create(ctx context.Context, step *domain.Step) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(step).Error
	})
}

func (r *StepRepository) Update(ctx context.Context, step *domain.Step) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(step).Updates(step).Error
	})
}

func (r *StepRepository) FindAllByWorkflowID(ctx context.Context, workflowID uuid.UUID, q dto.PaginateQuery) ([]domain.Step, int64, error) {
	var steps []domain.Step
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Step{}).
		Where("workflow_id = ?", workflowID)

	if q.Search != "" {
		db = db.Where("name ILIKE ?", "%"+q.Search+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "position"
	if q.SortBy != "" {
		sortBy = q.SortBy
	}

	orderBy := "asc"
	if q.OrderBy == "desc" {
		orderBy = "desc"
	}

	err := db.
		Preload("Endpoint").
		Order(sortBy + " " + orderBy).
		Limit(q.Limit).
		Offset(q.Offset()).
		Find(&steps).Error

	if err != nil {
		return nil, 0, err
	}
	return steps, total, nil
}

func (r *StepRepository) FindByWorkflowIDAndStepID(ctx context.Context, workflowID uuid.UUID, stepID uuid.UUID) (domain.Step, error) {
	var step domain.Step
	err := r.db.WithContext(ctx).
		Where("workflow_id = ? AND id = ?", workflowID, stepID).
		First(&step).Error
	if err != nil {
		return domain.Step{}, err
	}
	return step, nil
}
