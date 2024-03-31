package models

type RestaurantCategory struct {
	Id
	Name string `gorm:"type:varchar(50)" json:"name"`
	Timestamps
}
