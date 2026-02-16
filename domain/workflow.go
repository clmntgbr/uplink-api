package domain

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`

	Steps []Step `gorm:"foreignKey:WorkflowID" json:"steps"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Workflow) TableName() string {
	return "workflows"
}