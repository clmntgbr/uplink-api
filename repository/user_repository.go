package repository

import (
	"uplink-api/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
