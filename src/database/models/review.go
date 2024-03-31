package models

type Review struct {
	Id
	Rating       uint `gorm:"type:int" json:"rating"`
	UserId       uint `gorm:"type:bigint unsigned" json:"-"`
	RestaurantId uint `gorm:"type:bigint unsigned" json:"-"`
	Timestamps

	// Relation
	User       *User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user"`
	Restaurant *Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant"`
}
