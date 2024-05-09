package models

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	Average float32
	Count   uint
}

type Restaurant struct {
	Id
	Uuid       uuid.UUID `gorm:"type:uuid;unique;default:gen_random_uuid()" json:"uuid"`
	Name       string    `gorm:"type:varchar(30)" json:"name" binding:"required"`
	Phone      string    `gorm:"type:varchar(20);unique" json:"phone" binding:"required"`
	Email      string    `gorm:"type:varchar(30);unique" json:"email" binding:"required"`
	Banner     string    `gorm:"type:varchar(255)" json:"banner"`
	LocationId uint      `gorm:"type:bigint" json:"location_id"`
	CategoryId uint      `gorm:"type:bigint" json:"category_id"`
	Avatar     string    `gorm:"type:varchar(255)" json:"avatar"`
	Timestamps

	// Relation
	Location *Location           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:location_id" json:"location,omitempty"`
	Category *RestaurantCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`
	Reviews  []*Review           `gorm:"foreignKey:restaurant_id" json:"reviews,omitempty"`

	// Extra
	Rating Rating `gorm:"-" json:"rating"`
}

func (r *Restaurant) AfterFind(tx *gorm.DB) (err error) {
	var result Rating

	config.DB.Model(&Review{}).Select("AVG(rating) as average, COUNT(*) as count").Where("restaurant_id = ?", r.Id.Id).Group("restaurant_id").Scan(&result)

	r.Rating = result
	return nil
}
