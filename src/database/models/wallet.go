package models

type Wallet struct {
	Id
	Pin     string    `gorm:"type:varchar(255)" json:"pin,omitempty"`
	Balance uint      `gorm:"type:int" json:"balance"`
	Timestamps
}
