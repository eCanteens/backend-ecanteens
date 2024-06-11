package enums

import (
	"database/sql/driver"
	"fmt"
)

type TransactionType string

const (
	TrxTypePay      TransactionType = "PAY"
	TrxTypeTopUp    TransactionType = "TOPUP"
	TrxTypeWithdraw TransactionType = "WITHDRAW"
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