package repository

import (
	"context"

	"uplink-api/domain"

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

func (r *StepRepository) CreateBatch(ctx context.Context, steps []*domain.Step) error {
	if len(steps) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&steps).Error
	})
}

func (r *StepRepository) Update(ctx context.Context, step *domain.Step) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(step).Updates(step).Error
	})
}

func (r *StepRepository) DeleteByWorkflowID(ctx context.Context, workflowID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("workflow_id = ?", workflowID).Delete(&domain.Step{}).Error
	})
}
