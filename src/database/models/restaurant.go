package models

type Restaurant struct {
	PK
	Name       string `gorm:"type:varchar(50)" json:"name"`
	Avatar     string `gorm:"type:varchar(255)" json:"avatar"`
	Banner     string `gorm:"type:varchar(255)" json:"banner"`
	CategoryId uint   `gorm:"type:bigint" json:"category_id"`
	OwnerId    uint   `gorm:"type:bigint" json:"owner_id"`
	IsOpen     bool   `gorm:"type:boolean;default:false;not null" json:"is_open"`
	Timestamps

	// Relation
	Owner    *User               `gorm:"foreignKey:owner_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"owner,omitempty"`
	Category *RestaurantCategory `gorm:"foreignKey:category_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category,omitempty"`
	Reviews  []RestaurantReview  `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"reviews,omitempty"`

	// Extra
	RatingAvg   float64 `json:"rating_avg"`
	RatingCount int64   `json:"rating_count"`
}
