package transaction

type AddCartScheme struct {
	ProductId uint   `binding:"required,numeric" json:"product_id"`
	Quantity  uint   `binding:"required,numeric" json:"quantity"`
	Amount    uint   `binding:"required,numeric" json:"amount"`
	Notes     string `binding:"required" json:"notes"`
}
