package models

type Location struct {
	Id
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Address string `gorm:"type:text" json:"address"`
	Timestamps
}
