package transaction

type addCartScheme struct {
	ProductId uint  `binding:"required,numeric" json:"product_id"`
	Quantity  *uint `binding:"required,numeric" json:"quantity"`
}

type updateCartNoteScheme struct {
	Notes string `json:"notes"`
}

type orderScheme struct {
	CartId          uint   `binding:"required,numeric" json:"cart_id"`
	PaymentMethod   string `binding:"required,oneof=CASH ECANTEENSPAY" json:"payment_method"`
	IsPreorder      *bool  `binding:"required" json:"is_preorder"`
	FullfilmentDate string `binding:"required_if=IsPreorder true" json:"fullfilment_date"`
}

type getOrderQS struct {
	Page      string `form:"page"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
	Filter    string `form:"filter"`
}
