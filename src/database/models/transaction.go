package models

import "database/sql/driver"

type transactionType string
type transactionStatus string

const (
	OUTCOME transactionType = "OUTCOME"
	INCOME  transactionType = "INCOME"
)

func (ct *transactionType) Scan(value interface{}) error {
	*ct = transactionType(value.([]byte))
	return nil
}

func (ct transactionType) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	INPROGRESS transactionStatus = "INPROGRESS"
	SUCCESS    transactionStatus = "SUCCESS"
	CANCELED   transactionStatus = "CANCELED"
)

func (ct *transactionStatus) Scan(value interface{}) error {
	*ct = transactionStatus(value.([]byte))
	return nil
}

func (ct transactionStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Transaction struct {
	Id
	UserId uint              `gorm:"type:bigint" json:"user_id"`
	Type   transactionType   `gorm:"type:transaction_type" json:"type"`
	Status transactionStatus `gorm:"type:transaction_status" json:"status"`
	Amount uint              `gorm:"type:int" json:"amount"`
	Items  string            `gorm:"type:json" json:"items"`
	Timestamps

	// Relation
	User *User `gorm:"foreignKey:user_id" json:"user,omitempty"`
}
