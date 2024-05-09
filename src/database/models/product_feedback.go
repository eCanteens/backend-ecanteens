package models

type ProductFeedback struct {
	Id
	ProductId uint `gorm:"type:bigint" json:"product_id"`
	UserId    uint `gorm:"type:bigint" json:"user_id"`
	Like      bool `gorm:"type:bool" json:"like"`
	Timestamps

	// Relations
	Product *Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:product_id" json:"product,omitempty"`
	User    *User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id" json:"user,omitempty"`
}
