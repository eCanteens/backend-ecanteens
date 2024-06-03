package transaction

import (
	"database/sql/driver"
	"fmt"
)

type TransactionType string

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