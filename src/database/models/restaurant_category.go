package models

type RestaurantCategory struct {
	PK
	Name string `gorm:"type:varchar(50)" json:"name"`
	Timestamps
}
