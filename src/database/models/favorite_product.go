package models

type FavoriteProduct struct {
	PK
	UserId    uint `gorm:"type:bigint" json:"user_id"`
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	Timestamps

	// Relations
	User    *User    `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Product *Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product,omitempty"`
}
