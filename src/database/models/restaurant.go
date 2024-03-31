package models

type Restaurant struct {
	Id
	Name       string `gorm:"type:varchar(30)" json:"name" binding:"required"`
	Phone      string `gorm:"type:varchar(20);unique" json:"phone" binding:"required"`
	LocationId uint   `gorm:"type:bigint" json:"-"`
	Banner     string `gorm:"type:varchar(255)" json:"banner"`
	Balance    int    `gorm:"type:int" json:"balance"`
	CategoryId uint   `gorm:"type:bigint" json:"-"`
	Timestamps

	// Relation
	Location *Location           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:location_id" json:"location"`
	Category *RestaurantCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category"`
}
