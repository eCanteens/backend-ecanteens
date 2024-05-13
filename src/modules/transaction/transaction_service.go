package transaction

import "github.com/eCanteens/backend-ecanteens/src/database/models"

func getCartService(user *models.User) (*[]models.Cart, error) {
	var cart []models.Cart

	if err := findCart(*user.Id.Id, &cart); err != nil {
		return nil, err
	}

	return &cart, nil
}

func addCartService(user *models.User, body *AddCartScheme) error {
	var cart models.Cart

	if err := findOneCart(*user.Id.Id, body.ProductId, &cart); err != nil {
		return createCart(&models.Cart{
			UserId: *user.Id.Id,
			ProductId: body.ProductId,
			Quantity: body.Quantity,
			Amount: body.Amount,
			Notes: body.Notes,
		})
	}else {
		return updateCart(&cart)
	}
}