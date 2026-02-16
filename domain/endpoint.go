package domain

import (
	"time"

	"gorm.io/datatypes"
	"github.com/google/uuid"
)

type Endpoint struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	BaseUri string    `gorm:"not null" json:"baseUri"`
	Path string    `gorm:"not null" json:"path"`
	Method string    `gorm:"not null" json:"method"`
	Timeout int `gorm:"not null" json:"timeout"`

	Header datatypes.JSON `json:"header" gorm:"type:json"`
	Body   datatypes.JSON `json:"body" gorm:"type:json"`
	Query  datatypes.JSON `json:"query" gorm:"type:json"`

	ProjectID uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Endpoint) TableName() string {
	return "endpoints"
}