package models

import "time"

type Token struct {
	PK
	UserId    uint          `gorm:"type:bigint" json:"user_id"`
	Token     string        `gorm:"type:varchar(255);not null;unique" json:"token"`
	LastUsed  time.Time     `gorm:"not null;index" json:"last_used"`
	Timestamps

	// Relations
	User *User `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
