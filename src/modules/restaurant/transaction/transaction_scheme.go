package transaction

type getOrderQS struct {
	Page      string `form:"page"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
	Filter    string `form:"filter"`
}

type updateOrderScheme struct {
	Status string `binding:"required,oneof=INPROGRESS READY CANCELED" json:"status"`
	Reason string `binding:"required_if=Status CANCELED" json:"reason"`
}