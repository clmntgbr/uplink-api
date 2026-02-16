package repository

import (
	"context"

	"uplink-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, project *domain.Project) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(project).Error
	})
}

func (r *ProjectRepository) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Project, error) {
	var projects []domain.Project

	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("user_projects.user_id = ?", userID).
		Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, projectID uuid.UUID, userID uuid.UUID) (*domain.Project, error) {
	var project domain.Project

	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", projectID, userID).
		First(&project).Error

	if err != nil {
		return nil, err
	}
	return &project, nil
}
