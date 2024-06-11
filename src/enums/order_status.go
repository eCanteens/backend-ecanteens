package enums

import (
	"database/sql/driver"
	"fmt"
)

type OrderStatus string

const (
	OrderStatusWaiting    OrderStatus = "WAITING"
	OrderStatusInProgress OrderStatus = "INPROGRESS"
	OrderStatusReady      OrderStatus = "READY"
	OrderStatusSuccess    OrderStatus = "SUCCESS"
	OrderStatusCanceled   OrderStatus = "CANCELED"
)

func (ct *OrderStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = OrderStatus(v)
	case string:
		*ct = OrderStatus(v)
	default:
		return fmt.Errorf("unsupported Scan type for OrderStatus: %T", value)
	}
	return nil
}

func (ct OrderStatus) Value() (driver.Value, error) {
	return string(ct), nil
}
