package transaction

import (
	"errors"
	"fmt"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/constants/order"
	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func getCartService(user *models.User) (*[]models.Cart, error) {
	var carts []models.Cart

	if err := findCart(*user.Id, &carts, true); err != nil {
		return nil, err
	}

	return &carts, nil
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

			return create(cart)
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

			return create(cartItem)
		}

		// and quantity is 0 then return error
		return errors.New("produk tidak ditemukan di keranjang")
	} else {
		// if cart & cart item found
		if *body.Quantity > 0 {
			// if quantity not 0 then update cart item
			cartItem.Quantity = *body.Quantity
			return update(cartItem)
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

func getOrderService(userId uint, query *getOrderQS) (*pagination.Pagination, error) {
	var result pagination.Pagination

	if err := findOrder(&result, userId, query); err != nil {
		return nil, err
	}

	return &result, nil
}

func orderService(body *orderScheme, user *models.User) (*models.Order, error) {
	var cart models.Cart

	// Find cart
	if err := findCartById(body.CartId, &cart, true); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("keranjang tidak ditemukan")
		}
		return nil, err
	}

	// Make transaction
	trx := models.Transaction{
		TransactionCode: fmt.Sprintf("EC-%d-%d", time.Now().Unix(), *user.Id),
		UserId:          *user.Id,
		Type:            transaction.PAY,
		Status:          transaction.INPROGRESS,
		PaymentMethod:   transaction.TransactionPaymentMethod(body.PaymentMethod),
	}

	// Validate restaurant open
	if !cart.Restaurant.IsOpen {
		return nil, errors.New("restoran sedang tutup")
	}

	var fullfilmentDate *time.Time

	// Validate fullfilment date if preorder
	if *body.IsPreorder {
		date, err := time.Parse(time.RFC3339, body.FullfilmentDate)
		if err != nil {
			return nil, errors.New("format waktu tidak valid")
		}

		fullfilmentDate = &date
	}

	// Create order
	ord := models.Order{
		UserId:          cart.UserId,
		RestaurantId:    cart.RestaurantId,
		Notes:           cart.Notes,
		Status:          order.WAITING,
		IsPreorder:      *body.IsPreorder,
		FullfilmentDate: fullfilmentDate,
	}

	// Insert cart items into order items & calculate amount
	for _, item := range cart.Items {
		if item.Product.Stock == 0 {
			return nil, errors.New("stok produk habis")
		}
		ord.Items = append(ord.Items, models.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})

		ord.Amount += item.Quantity * item.Product.Price
		trx.Amount += item.Quantity * item.Product.Price
	}

	ord.Transaction = &trx

	// Validate wallet balance
	if trx.PaymentMethod == transaction.ECANTEENSPAY && user.Wallet.Balance < trx.Amount {
		return nil, errors.New("saldo anda tidak mencukupi")
	}

	// Write order and transaction data into db
	if err := create(&ord); err != nil {
		return nil, err
	}

	// Delete cart & cart items data
	if err := deleteRecord(&cart); err != nil {
		return nil, err
	}

	// if preorder & payment method with ecanteenspay then update user balance & resto owner balance
	if ord.IsPreorder && trx.PaymentMethod == transaction.ECANTEENSPAY {
		user.Wallet.Balance -= trx.Amount

		if err := update(user.Wallet); err != nil {
			return nil, err
		}

		cart.Restaurant.Owner.Wallet.Balance += trx.Amount

		if err := update(cart.Restaurant.Owner.Wallet); err != nil {
			return nil, err
		}
	}

	return &ord, nil
}

func cancelOrderService(id string) error {
	return cancelOrderById(id)
}
