package service

import (
	"uplink-api/domain"
	"uplink-api/dto"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(user *domain.User) (*dto.UserOutput, error) {
	output := dto.NewUserOutput(*user)
	return &output, nil
}
