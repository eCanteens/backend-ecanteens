package enums

import (
	"database/sql/driver"
	"fmt"
)

type OrderCancelBy string

const (
	OrderCancelByResto OrderCancelBy = "Resto"
	OrderCancelByUser  OrderCancelBy = "User"
)

func (ct *OrderCancelBy) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*ct = OrderCancelBy(v)
	case string:
		*ct = OrderCancelBy(v)
	default:
		return fmt.Errorf("unsupported Scan type for OrderCancelBy: %T", value)
	}
	return nil
}

func (ct OrderCancelBy) Value() (driver.Value, error) {
	return string(ct), nil
}
