package models

type FavoriteProduct struct {
	PK
	UserId    uint `gorm:"type:bigint" json:"user_id"`
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	Timestamps

	// Relations
	User    *User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user,omitempty"`
	Product *Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:product_id" json:"product,omitempty"`
}
