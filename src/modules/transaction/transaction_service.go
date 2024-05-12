package transaction

import "github.com/eCanteens/backend-ecanteens/src/database/models"

func AddCartService(user *models.User, body *AddCartScheme) error {
	cart := &models.Cart{
		UserId: *user.Id.Id,
		ProductId: body.ProductId,
		Quantity: body.Quantity,
		Notes: body.Notes,
		Amount: body.Amount,
	}

	return createCart(cart)
}