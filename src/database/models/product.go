package models

type Product struct {
	PK
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
	Restaurant *Restaurant      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:restaurant_id" json:"restaurant,omitempty"`
	Category   *ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`

	// Extra
	Like    int64 `json:"like"`
	Dislike int64 `json:"dislike"`
}
