package transaction

import (
	"errors"
	"fmt"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/constants/order"
	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
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
		return errors.New("stok produk habis")
	}

	if !product.Restaurant.IsOpen {
		return errors.New("restoran sedang tutup")
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

func getOrderService(userId uint) (*[]models.Order, error) {
	var orders []models.Order

	if err := findOrder(&orders, userId); err != nil {
		return nil, err
	}

	return &orders, nil
}

func orderService(body *orderScheme, user *models.User) error {
	var carts []models.Cart
	if err := findCart(*user.Id, &carts, true); err != nil {
		return err
	}

	if len(carts) == 0 {
		return errors.New("keranjang masih kosong")
	}

	trx := models.Transaction{
		TransactionCode: fmt.Sprintf("EC-%d-%d", time.Now().Unix(), *user.Id),
		UserId:          *user.Id,
		Type:            transaction.PAY,
		Status:          transaction.INPROGRESS,
		PaymentMethod:   transaction.TransactionPaymentMethod(body.PaymentMethod),
	}

	var fullfilmentDate *time.Time

	if *body.IsPreorder {
		date, err := time.Parse(time.RFC3339, body.FullfilmentDate)
		if err != nil {
			return errors.New("format waktu tidak valid")
		}

		fullfilmentDate = &date
	}

	for _, cart := range carts {
		if !cart.Restaurant.IsOpen {
			return errors.New("restoran sedang tutup")
		}
		ord := models.Order{
			UserId:          cart.UserId,
			RestaurantId:    cart.RestaurantId,
			Notes:           cart.Notes,
			Status:          order.WAITING,
			IsPreorder:      *body.IsPreorder,
			FullfilmentDate: fullfilmentDate,
		}

		for _, item := range cart.Items {
			if item.Product.Stock == 0 {
				return errors.New("stok produk habis")
			}
			ord.Items = append(ord.Items, models.OrderItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
			})

			ord.Amount += item.Quantity * item.Product.Price
			trx.Amount += item.Quantity * item.Product.Price
		}

		trx.Orders = append(trx.Orders, ord)
	}

	if trx.PaymentMethod == transaction.ECANTEENSPAY && user.Wallet.Balance < trx.Amount {
		return errors.New("saldo anda tidak mencukupi")
	}

	if err := create(&trx); err != nil {
		return err
	}

	if err := deleteRecord(&carts); err != nil {
		return err
	}

	return nil
}

func cancelOrderService(userId uint) error {
	return cancelOrder(userId)
}
