package service

import (
	"uplink-api/dto"
	"uplink-api/errors"
	"uplink-api/repository"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(userID uuid.UUID) (*dto.UserOutput, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	output := dto.NewUserOutput(*user)
	return &output, nil
}
