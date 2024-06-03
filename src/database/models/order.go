package models

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/constants/order"
)

type Order struct {
	PK
	UserId          uint              `gorm:"type:bigint" json:"user_id"`
	RestaurantId    uint              `gorm:"type:bigint" json:"restaurant_id"`
	Notes           string            `gorm:"type:varchar(255)" json:"notes"`
	Amount          uint              `gorm:"type:int" json:"amount"`
	Status          order.OrderStatus `gorm:"type:varchar(20);default:WAITING" json:"status"` // [inprogress, success, canceled]
	IsPreorder      bool              `gorm:"type:bool" json:"is_preorder"`
	FullfilmentDate *time.Time        `gorm:"type:timestamptz" json:"fullfilment_date"`
	TransactionId   uint              `gorm:"type:bigint" json:"transaction_id"`
	Timestamps

	// Relation
	User       *User       `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Restaurant *Restaurant `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"restaurant,omitempty"`
	Items      []OrderItem `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
}
