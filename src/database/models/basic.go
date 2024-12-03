package models

import (
	"time"

	"gorm.io/gorm"
)

type PK struct {
	Id *uint `gorm:"type:bigint;primaryKey" json:"id"`
}

type Timestamps struct {
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
