package models

import "time"

type Order struct {
	PK
	UserId          uint      `gorm:"type:bigint" json:"user_id"`
	RestaurantId    int       `gorm:"type:bigint" json:"restaurant_id"`
	Notes           string    `gorm:"type:varchar(255)" json:"notes"`
	Amount          int       `gorm:"type:int" json:"amount"`
	Status          string    `gorm:"type:varchar(20)" json:"status"` // [inprogress, success, canceled]
	IsPreorder      bool      `gorm:"type:bool" json:"is_preorder"`
	FullfilmentDate time.Time `gorm:"type:timestamptz" json:"fullfilment_date"`
	Timestamps

	// Relation
	User       *User       `gorm:"foreignKey:user_id" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"foreignKey:restaurant_id" json:"restaurant,omitempty"`
}
