package enums

import (
	"database/sql/driver"
	"fmt"
)

type TransactionPaymentMethod string

const (
	TrxPaymentCash         TransactionPaymentMethod = "CASH"
	TrxPaymentEcanteensPay TransactionPaymentMethod = "ECANTEENSPAY"
)

func (ct *TransactionPaymentMethod) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = TransactionPaymentMethod(v)
	case string:
		*ct = TransactionPaymentMethod(v)
	default:
		return fmt.Errorf("unsupported Scan type for PaymentMethod: %T", value)
	}
	return nil
}

func (ct TransactionPaymentMethod) Value() (driver.Value, error) {
	return string(ct), nil
}