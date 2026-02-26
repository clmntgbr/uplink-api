package domain

import (
	"time"

	"github.com/google/uuid"
)

type Step struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Position int       `gorm:"not null" json:"position"`

	EndpointID uuid.UUID `gorm:"type:uuid;not null" json:"endpoint_id"`
	WorkflowID uuid.UUID `gorm:"type:uuid;not null" json:"workflow_id"`

	Endpoint Endpoint `gorm:"foreignKey:EndpointID" json:"endpoint"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Step) TableName() string {
	return "steps"
}
