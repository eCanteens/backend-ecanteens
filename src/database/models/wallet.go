package models

import "github.com/google/uuid"

type Wallet struct {
	Id
	Uuid    uuid.UUID `gorm:"type:uuid;unique;default:gen_random_uuid()" json:"uuid"`
	Pin     string    `gorm:"type:varchar(255)" json:"pin,omitempty"`
	Balance uint      `gorm:"type:int" json:"balance"`
	Timestamps
}
