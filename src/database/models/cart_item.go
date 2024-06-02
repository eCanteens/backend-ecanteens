package models

type CartItem struct {
	PK
	CartId    uint `gorm:"type:bigint" json:"user_id"`
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	Quantity  uint `gorm:"type:int" json:"quantity"`
	Timestamps

	// Relation
	Cart    *Cart    `gorm:"foreignKey:cart_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cart,omitempty"`
	Product *Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
