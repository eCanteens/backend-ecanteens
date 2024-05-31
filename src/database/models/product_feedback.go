package models

type ProductFeedback struct {
	PK
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	UserId    uint `gorm:"type:bigint" json:"user_id"`
	IsLike    bool `gorm:"type:bool" json:"is_like"`
	Timestamps

	// Relations
	Product *Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:product_id" json:"product,omitempty"`
	User    *User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user,omitempty"`
}
