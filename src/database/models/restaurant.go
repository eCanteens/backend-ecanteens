package models

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"gorm.io/gorm"
)

type Rating struct {
	Average float32
	Count   uint
}

type Restaurant struct {
	Id
	Name       string    `gorm:"type:varchar(50)" json:"name"`
	Avatar     string    `gorm:"type:varchar(255)" json:"avatar"`
	Banner     string    `gorm:"type:varchar(255)" json:"banner"`
	CategoryId uint      `gorm:"type:bigint" json:"category_id"`
	OwnerId    uint      `gorm:"type:bigint" json:"owner_id"`
	IsOpen     bool      `gorm:"type:boolean;default:false" json:"is_open"`
	Timestamps

	// Relation
	Owner    *User               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:owner_id" json:"owner,omitempty"`
	Category *RestaurantCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`
	Reviews  []*Review           `gorm:"foreignKey:restaurant_id" json:"reviews,omitempty"`

	// Extra
	Rating Rating `gorm:"-" json:"rating"`
}

func (r *Restaurant) AfterFind(tx *gorm.DB) (err error) {
	var result Rating

	config.DB.Model(&Review{}).Select("AVG(rating) as average, COUNT(*) as count").Where("restaurant_id = ?", r.Id.Id).Group("restaurant_id").Scan(&result)

	r.Rating = result
	return
}
