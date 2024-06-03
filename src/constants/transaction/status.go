package transaction

import (
	"database/sql/driver"
	"fmt"
)

type TransactionStatus string

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