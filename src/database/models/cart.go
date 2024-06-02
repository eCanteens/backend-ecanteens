package models

type Cart struct {
	PK
	UserId       uint   `gorm:"type:bigint" json:"user_id"`
	RestaurantId uint   `gorm:"type:bigint" json:"restaurant_id"`
	Notes        string `gorm:"type:varchar(255)" json:"notes"`
	Timestamps

	// Relation
	User       *User       `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"restaurant,omitempty"`
	Items      []CartItem  `gorm:"foreignKey:cart_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
}
