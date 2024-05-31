package models

import (
	"time"

	"gorm.io/gorm"
)

type PK struct {
	Id *uint `gorm:"type:bigint;primaryKey" json:"id"`
}

type Timestamps struct {
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
