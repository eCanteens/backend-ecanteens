package models

import "gorm.io/datatypes"

type Review struct {
	PK
	Rating  float32                      `gorm:"type:float" json:"rating"`
	OrderId uint                         `gorm:"type:bigint" json:"order_id"`
	Tags    datatypes.JSONType[[]string] `gorm:"type:json" json:"tags"` // [Rasa, Kebersihan, Porsi, Kemasan]
	Comment string                       `gorm:"type:text" json:"comment"`
	Timestamps

	// Relation
	Order *Order `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"order,omitempty"`
}
