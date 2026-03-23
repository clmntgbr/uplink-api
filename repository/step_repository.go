package repository

import (
	"context"

	"uplink-api/domain"

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

func (r *StepRepository) Delete(ctx context.Context, step *domain.Step) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Delete(step).Error
	})
}
