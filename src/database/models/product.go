package models

import (
// "gorm.io/gorm"
)

type Product struct {
	Id
	RestaurantId uint   `gorm:"type:bigint" json:"restaurant_id"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	Image        string `gorm:"type:varchar(255)" json:"image"`
	CategoryId   uint   `gorm:"type:bigint" json:"category_id"`
	Price        uint   `gorm:"type:int" json:"price"`
	Stock        uint   `gorm:"type:int" json:"stock"`
	Sold         uint   `gorm:"type:int" json:"sold"`
	Timestamps

	// Relations
	Restaurant *Restaurant      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant,omitempty"`
	Category   *ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`
	Feedbacks  []User           `gorm:"many2many:product_feedbacks;" json:"feedbacks,omitempty"`

	// Extra
	Like    uint `gorm:"-" json:"like"`
	Dislike uint `gorm:"-" json:"dislike"`
}

// func (p *Product) AfterFind(tx *gorm.DB) (err error) {
// 	tx.Preload("Feedback", func(db *gorm.DB) *gorm.DB {
// 		return db.Where("like = ?", true)
// 	})

// 	return nil
// }
