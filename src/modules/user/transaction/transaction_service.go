package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func getCartService(user *models.User) (*[]models.Cart, error) {
	var cart []models.Cart

	if err := findCart(*user.Id, &cart, true); err != nil {
		return nil, err
	}

	return &cart, nil
}

func addCartService(user *models.User, body *addUpdateCartScheme) error {
	var cart models.Cart

	findOneCart(&cart, *user.Id, body.ProductId)

	if body.Quantity == 0 {
		return deleteCart(*user.Id, body.ProductId)
	}

	cart.UserId = *user.Id
	cart.ProductId = body.ProductId
	cart.Quantity = body.Quantity
	cart.Amount = body.Amount
	cart.Notes = body.Notes

	return saveCart(&cart)
}
