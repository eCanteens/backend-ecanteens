package models

type OrderItem struct {
	PK
	OrderId   uint `gorm:"type:bigint" json:"order_id"`
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	Quantity  uint `gorm:"type:int" json:"quantity"`
	Price     uint `gorm:"type:int" json:"price"`
	Timestamps

	// Relation
	Order   *Order   `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"order,omitempty"`
	Product *Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
}
