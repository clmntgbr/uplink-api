package dto

import (
	"time"
	"uplink-api/domain"
)

type UserOutput struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUserOutput(user domain.User) UserOutput {
	return UserOutput{
		ID:       user.ID.String(),
		Email:    user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
