package repository

import (
	"context"
	"errors"

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

func (r *ProjectRepository) FindAllByUserID(ctx context.Context, user *domain.User) ([]domain.Project, error) {
	var projects []domain.Project

	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("user_projects.user_id = ?", user.ID).
		Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, projectID uuid.UUID, user *domain.User) (*domain.Project, error) {
	var project domain.Project

	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", projectID, user.ID).
		First(&project).Error

	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) ActivateProject(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) error {
	var project domain.Project
	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", projectID, userID).
		First(&project).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("project not found or access denied")
		}
		return err
	}

	return r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Update("active_project_id", projectID).Error
}

