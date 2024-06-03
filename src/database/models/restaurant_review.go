package models

type RestaurantReview struct {
	PK
	Rating       float32 `gorm:"type:float" json:"rating"`
	UserId       uint    `gorm:"type:bigint" json:"user_id"`
	RestaurantId uint    `gorm:"type:bigint" json:"restaurant_id"`
	Comment      string  `gorm:"type:text" json:"comment"`
	Timestamps

	// Relation
	User       *User       `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"restaurant,omitempty"`
}
