package transaction

import "github.com/eCanteens/backend-ecanteens/src/database/models"

type addCartScheme struct {
	ProductId uint  `binding:"required" json:"product_id"`
	Quantity  *uint `binding:"required" json:"quantity"`
}

type updateCartNoteScheme struct {
	Notes string `json:"notes"`
}

type orderScheme struct {
	CartId          uint   `binding:"required" json:"cart_id"`
	PaymentMethod   string `binding:"required,oneof=CASH ECANTEENSPAY" json:"payment_method"`
	IsPreorder      *bool  `binding:"required" json:"is_preorder"`
	FullfilmentDate string `binding:"required_if=IsPreorder true" json:"fullfilment_date"`
	Notes           string `json:"notes"`
}

type updateOrderScheme struct {
	Status string `binding:"required,oneof=SUCCESS CANCELED" json:"status"`
	Reason string `binding:"required_if=Status CANCELED" json:"reason"`
}

type postReviewScheme struct {
	Rating  float32  `binding:"required" json:"rating"`
	Tags    []string `binding:"required" json:"tags"`
	Comment string   `json:"comment"`
}

type getOrderQS struct {
	Page      string `form:"page"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
	Filter    string `form:"filter"`
}

type getTrxHistoryQS struct {
	Page      string `form:"page"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type restaurantCartQS struct {
	RestaurantId uint `form:"restaurant_id"`
}

type trxHistoryDataTo struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type trxHistoryData struct {
	*models.Transaction
	To *trxHistoryDataTo `json:"to"`
}
