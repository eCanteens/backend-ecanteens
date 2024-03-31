package models

type Favorite struct {
	Id
	UserId       uint `gorm:"type:bigint" json:"-"`
	RestaurantId uint `gorm:"type:bigint" json:"-"`
	Timestamps

	// Relations
	User       *User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user"`
	Restaurant *Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant"`
}
