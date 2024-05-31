package models

type Restaurant struct {
	PK
	Name       string `gorm:"type:varchar(50)" json:"name"`
	Avatar     string `gorm:"type:varchar(255)" json:"avatar"`
	Banner     string `gorm:"type:varchar(255)" json:"banner"`
	CategoryId uint   `gorm:"type:bigint" json:"category_id"`
	OwnerId    uint   `gorm:"type:bigint" json:"owner_id"`
	IsOpen     bool   `gorm:"type:boolean;default:false" json:"is_open"`
	Timestamps

	// Relation
	Owner    *User               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:owner_id" json:"owner,omitempty"`
	Category *RestaurantCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:category_id" json:"category,omitempty"`
	Reviews  []*Review           `gorm:"foreignKey:restaurant_id" json:"reviews,omitempty"`

	// Extra
	RatingAvg   float64 `json:"rating_avg"`
	RatingCount int64   `json:"rating_count"`
}
