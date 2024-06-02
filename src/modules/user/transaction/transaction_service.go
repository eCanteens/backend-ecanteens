package transaction

import (
	"errors"
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"gorm.io/gorm"
)

func getCartService(user *models.User) (*[]models.Cart, error) {
	var cart []models.Cart

	if err := findCart(*user.Id, &cart, true); err != nil {
		return nil, err
	}

	return &cart, nil
}

func updateCartService(id string, body *updateCartNoteScheme) error {
	return updateCartNote(id, body.Notes)
}

func addCartService(user *models.User, body *addCartScheme) error {
	var product models.Product
	var carts []models.Cart

	if err := findOneProduct(&product, body.ProductId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("produk tidak ditemukan")
		}
		return err
	}

	if product.Stock == 0 {
		return errors.New("stok habis")
	}

	if !product.Restaurant.IsOpen {
		return errors.New("restoran tutup")
	}

	if err := findCart(*user.Id, &carts, false); err != nil {
		return err
	}

	cart := helpers.Find(&carts, func(cart *models.Cart) bool {
		return cart.RestaurantId == product.RestaurantId
	})

	if cart == nil {
		// if cart not found
		if *body.Quantity > 0 {
			// and quantity not 0 then create cart and cart item
			cart = &models.Cart{
				UserId:       *user.Id,
				RestaurantId: product.RestaurantId,
				Items: []models.CartItem{
					{
						ProductId: body.ProductId,
						Quantity:  *body.Quantity,
					},
				},
			}

			return saveRecord(cart)
		}

		// but quantity is 0 then return error
		return errors.New("produk tidak ditemukan di keranjang")
	}
	fmt.Println("atas", cart.Items)

	cartItem := helpers.Find(&cart.Items, func(item *models.CartItem) bool {
		return item.ProductId == body.ProductId
	})

	if cartItem == nil {
		// if cart found but cart item not found
		if *body.Quantity > 0 {
			// if quantity not 0 then create cart item
			cartItem = &models.CartItem{
				CartId:    *cart.Id,
				ProductId: body.ProductId,
				Quantity:  *body.Quantity,
			}

			return saveRecord(cartItem)
		}

		// and quantity is 0 then return error
		return errors.New("produk tidak ditemukan di keranjang")
	} else {
		// if cart & cart item found
		if *body.Quantity > 0 {
			// if quantity not 0 then update cart item
			cartItem.Quantity = *body.Quantity
			return saveRecord(cartItem)
		} else {
			// if quantity is 0
			if len(cart.Items) > 1 {
				// if cart items more than 1 then delete cart item
				return deleteRecord(cartItem)
			} else {
				// if cart items just 1 then delete cart
				return deleteRecord(cart)
			}
		}
	}
}
