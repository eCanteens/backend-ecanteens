package models

type Review struct {
	Id
	Rating       uint `gorm:"type:int" json:"rating"`
	UserId       uint `gorm:"type:bigint" json:"user_id"`
	RestaurantId uint `gorm:"type:bigint" json:"restaurant_id"`
	Timestamps

	// Relation
	User       *User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant,omitempty"`
}
