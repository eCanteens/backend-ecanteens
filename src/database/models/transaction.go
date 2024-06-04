package models

import "github.com/eCanteens/backend-ecanteens/src/constants/transaction"

type Transaction struct {
	PK
	TransactionCode string                               `gorm:"type:varchar(255);unique" json:"transaction_id"`
	UserId          uint                                 `gorm:"type:bigint" json:"user_id"`
	Type            transaction.TransactionType          `gorm:"type:varchar(20)" json:"type"`   // [PAY, TOPUP, WITHDRAW]
	Status          transaction.TransactionStatus        `gorm:"type:varchar(20)" json:"status"` // [INPROGRESS, SUCCESS, CANCELED]
	Amount          uint                                 `gorm:"type:int" json:"amount"`
	PaymentMethod   transaction.TransactionPaymentMethod `gorm:"type:varchar(20)" json:"payment_method"`
	Timestamps

	// Relation
	User   *User   `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Orders []Order `gorm:"foreignKey:transaction_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order,omitempty"`
}
