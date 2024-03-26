package models

import "time"

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(30);unique" json:"username" binding:"required"`
	Password  string    `gorm:"type:varchar(255)" json:"password,omitempty" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}