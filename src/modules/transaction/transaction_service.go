package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func getCartService(user *models.User) (*[]models.Cart, error) {
	var cart []models.Cart

	if err := findCart(*user.Id.Id, &cart, true); err != nil {
		return nil, err
	}

	return &cart, nil
}

func addCartService(user *models.User, body *[]AddUpdateCartScheme) error {
	var existedCart []models.Cart
	var updateCart []*models.Cart

	if err := findCart(*user.Id.Id, &existedCart, false); err != nil {
		return err
	}

	for _, e := range existedCart {
		for _, b := range *body {
			if e.ProductId == b.ProductId {
				updateCart = append(updateCart, &models.Cart{
					Id: e.Id,
					UserId: *user.Id.Id,
					ProductId: b.ProductId,
					Quantity: e.Quantity + b.Quantity,
					Amount: b.Amount,
					Notes: b.Notes,
				})

				updatedBody := helpers.Filter(*body, func(scheme AddUpdateCartScheme) bool {
					return scheme.ProductId != e.ProductId
				})

				body = &updatedBody

				break
			}
		}
	}

	for _, b := range *body {
		updateCart = append(updateCart, &models.Cart{
			UserId: *user.Id.Id,
			ProductId: b.ProductId,
			Quantity: b.Quantity,
			Amount: b.Amount,
			Notes: b.Notes,
		})
	}

	return updateManyCart(&updateCart)
}

func updateCartService(user *models.User, body *AddUpdateCartScheme) error {
	if body.Quantity == 0 {
		return deleteCart(*user.Id.Id, body.ProductId)
	} else {
		cart := models.Cart{
			UserId:    *user.Id.Id,
			ProductId: body.ProductId,
			Quantity:  body.Quantity,
			Amount:    body.Amount,
			Notes:     body.Notes,
		}
	
		return updateCart(*user.Id.Id, body.ProductId, &cart)
	}
}
