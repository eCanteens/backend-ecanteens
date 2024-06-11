package enums

import (
	"database/sql/driver"
	"fmt"
)

type TransactionStatus string

const (
	TrxStatusInProgress TransactionStatus = "INPROGRESS"
	TrxStatusSuccess    TransactionStatus = "SUCCESS"
	TrxStatusCanceled   TransactionStatus = "CANCELED"
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