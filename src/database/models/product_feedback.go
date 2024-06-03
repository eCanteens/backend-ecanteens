package models

type ProductFeedback struct {
	PK
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	UserId    uint `gorm:"type:bigint" json:"user_id"`
	IsLike    bool `gorm:"type:bool" json:"is_like"`
	Timestamps

	// Relations
	Product *Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product,omitempty"`
	User    *User    `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
}
