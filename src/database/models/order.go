package models

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/enums"
)

type Order struct {
	PK
	UserId          uint                 `gorm:"type:bigint" json:"user_id,omitempty"`
	RestaurantId    uint                 `gorm:"type:bigint" json:"restaurant_id,omitempty"`
	Notes           string               `gorm:"type:varchar(255)" json:"notes"`
	Amount          uint                 `gorm:"type:int" json:"amount,omitempty"`
	Status          enums.OrderStatus    `gorm:"type:varchar(20);default:WAITING" json:"status,omitempty"` // [WAITING, INPROGRESS, READY, SUCCESS, CANCELED]
	IsPreorder      bool                 `gorm:"type:bool" json:"is_preorder"`
	FullfilmentDate *time.Time           `gorm:"type:timestamptz" json:"fullfilment_date,omitempty"`
	TransactionId   uint                 `gorm:"type:bigint" json:"transaction_id,omitempty"`
	CancelReason    *string              `gorm:"type:varchar(255)" json:"cancel_reason,omitempty"`
	CancelBy        *enums.OrderCancelBy `gorm:"type:varchar(30)" json:"cancel_by,omitempty"` // [RESTO, USER]
	Timestamps

	// Relation
	User        *User        `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Restaurant  *Restaurant  `gorm:"foreignKey:restaurant_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"restaurant,omitempty"`
	Items       []OrderItem  `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items,omitempty"`
	Transaction *Transaction `gorm:"foreignKey:transaction_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaction,omitempty"`
	Review      *Review      `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"review,omitempty"`
}
