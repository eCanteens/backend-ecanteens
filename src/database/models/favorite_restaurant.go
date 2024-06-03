package models

type FavoriteRestaurant struct {
	PK
	UserId       uint `gorm:"type:bigint" json:"user_id"`
	RestaurantId uint `gorm:"type:bigint" json:"restaurant_id"`
	Timestamps

	// Relations
	User       *User       `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"restaurant,omitempty"`
}
