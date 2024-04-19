package models

type Favorite struct {
	Id
	UserId       uint `gorm:"type:bigint" json:"user_id"`
	RestaurantId uint `gorm:"type:bigint" json:"restaurant_id"`
	Timestamps

	// Relations
	User       *User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant,omitempty"`
}
