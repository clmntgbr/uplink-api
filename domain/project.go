package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`

	Endpoints []Endpoint `gorm:"foreignKey:ProjectID" json:"endpoints"`
	Workflows []Workflow `gorm:"foreignKey:ProjectID" json:"workflows"`
	Users     []User     `gorm:"many2many:user_projects" json:"users,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	IsActive bool `gorm:"-"`
}

func (Project) TableName() string {
	return "projects"
}
