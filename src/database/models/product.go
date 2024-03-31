package models

type Product struct {
	Id
	RestaurantId uint   `gorm:"type:bigint unsigned" json:"-"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Description  string `gorm:"type:text" json:"description"`
	Image        string `gorm:"type:varchar(255)" json:"image"`
	CategoryId   uint   `gorm:"type:bigint unsigned" json:"-"`
	Price        uint   `gorm:"type:int(10)" json:"price"`
	Stock        uint   `gorm:"type:int(10)" json:"stock"`
	Sold         uint   `gorm:"type:int(10)" json:"sold"`
	Like         uint   `gorm:"type:int(10)" json:"like"`
	Dislike      uint   `gorm:"type:int(10)" json:"dislike"`
	Timestamps

	// Relations
	Restaurant *Restaurant      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:restaurant_id" json:"restaurant"`
	Category   *ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category"`
}
