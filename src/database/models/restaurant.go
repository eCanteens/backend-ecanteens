package models

import (
	"strconv"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id
	Name       string `gorm:"type:varchar(30)" json:"name" binding:"required"`
	Phone      string `gorm:"type:varchar(20);unique" json:"phone" binding:"required"`
	LocationId uint   `gorm:"type:bigint" json:"location_id"`
	Banner     string `gorm:"type:varchar(255)" json:"banner"`
	Avatar     string `gorm:"type:varchar(255)" json:"avatar"`
	Balance    uint    `gorm:"type:int" json:"balance"`
	CategoryId uint   `gorm:"type:bigint" json:"category_id"`
	Timestamps

	// Relation
	Location *Location           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:location_id" json:"location,omitempty"`
	Category *RestaurantCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`
	Reviews  []*Review           `gorm:"foreignKey:restaurant_id" json:"reviews,omitempty"`

	// Extra
	Rating float32 `gorm:"-" json:"rating"`
}

func (r *Restaurant) AfterFind(tx *gorm.DB) (err error) {
	type Result struct {
		Average string
	}

	var result Result

	config.DB.Model(&Review{}).Select("AVG(rating) as average").Where("restaurant_id = ?", r.Id.Id).Group("restaurant_id").Scan(&result)

	avgFloat, _ := strconv.ParseFloat(result.Average, 32)

	r.Rating = float32(avgFloat)
	return nil
}
