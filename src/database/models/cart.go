package models

type Cart struct {
	Id
	UserId    uint   `gorm:"type:bigint" json:"user_id"`
	ProductId uint   `gorm:"type:bigint" json:"product_id"`
	Quantity  uint   `gorm:"type:int" json:"quantity"`
	Amount    uint   `gorm:"type:int" json:"amount"`
	Notes     string `gorm:"type:varchar(255)" json:"notes"`
	Timestamps

	// Relation
	User    *User    `gorm:"foreignKey:user_id" json:"user,omitempty"`
	Product *Product `gorm:"foreignKey:product_id" json:"product,omitempty"`
}
