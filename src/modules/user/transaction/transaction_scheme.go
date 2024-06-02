package transaction

type addCartScheme struct {
	ProductId uint  `binding:"required,numeric" json:"product_id"`
	Quantity  *uint `binding:"required,numeric" json:"quantity"`
}

type updateCartNoteScheme struct {
	Notes string `json:"notes"`
}