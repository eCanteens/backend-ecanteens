package models

import "github.com/google/uuid"

type Wallet struct {
	Id
	Uuid         uuid.UUID `gorm:"type:uuid;unique;default:gen_random_uuid()" json:"uuid"`
	Balance      uint      `gorm:"type:int" json:"balance"`
	Timestamps
}
