package models

import (
	"database/sql/driver"
	"fmt"
)

type TransactionType string
type TransactionStatus string

const (
	PAY      TransactionType = "PAY"
	TOPUP    TransactionType = "TOPUP"
	WITHDRAW TransactionType = "WITHDRAW"
)

func (ct *TransactionType) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = TransactionType(v)
	case string:
		*ct = TransactionType(v)
	default:
		return fmt.Errorf("unsupported Scan type for TransactionType: %T", value)
	}
	return nil
}

func (ct TransactionType) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	INPROGRESS TransactionStatus = "INPROGRESS"
	SUCCESS    TransactionStatus = "SUCCESS"
	CANCELED   TransactionStatus = "CANCELED"
)

func (ct *TransactionStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = TransactionStatus(v)
	case string:
		*ct = TransactionStatus(v)
	default:
		return fmt.Errorf("unsupported Scan type for TransactionStatus: %T", value)
	}
	return nil
}

func (ct TransactionStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Transaction struct {
	PK
	TransactionId string            `gorm:"type:varchar(255);unique" json:"transaction_id"`
	UserId        uint              `gorm:"type:bigint" json:"user_id"`
	Type          TransactionType   `gorm:"type:Transaction_type" json:"type"`
	Status        TransactionStatus `gorm:"type:Transaction_status" json:"status"`
	Amount        uint              `gorm:"type:int" json:"amount"`
	Items         string            `gorm:"type:json" json:"items"`
	Timestamps

	// Relation
	User *User `gorm:"foreignKey:user_id" json:"user,omitempty"`
}
