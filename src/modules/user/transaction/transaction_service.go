package transaction

import (
	"errors"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
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

func updateCartService(id string, body *updateCartNoteScheme, userId uint) error {
	return updateCartNote(id, userId, body.Notes)
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

func getOrderService(userId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	result := pagination.New(models.Order{})

	if err := findOrder(result, userId, query); err != nil {
		return nil, err
	}

	return result, nil
}

func orderService(body *orderScheme, user *models.User) (*models.Order, error) {
	var cart models.Cart

	// Find cart
	if err := findCartById(body.CartId, &cart, *user.Id, true); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("keranjang tidak ditemukan")
		}
		return nil, err
	}

	// Make transaction
	transaction := models.Transaction{
		TransactionCode: helpers.GenerateTrxCode(*user.Id),
		UserId:          *user.Id,
		Type:            enums.TrxTypePay,
		Status:          enums.TrxStatusInProgress,
		PaymentMethod:   enums.TransactionPaymentMethod(body.PaymentMethod),
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
	order := models.Order{
		UserId:          cart.UserId,
		RestaurantId:    cart.RestaurantId,
		Notes:           cart.Notes,
		Status:          enums.OrderStatusWaiting,
		IsPreorder:      *body.IsPreorder,
		FullfilmentDate: fullfilmentDate,
	}

	// Insert cart items into order items & calculate amount
	for _, item := range cart.Items {
		if item.Product.Stock == 0 {
			return nil, errors.New("stok produk habis")
		}
		order.Items = append(order.Items, models.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})

		order.Amount += item.Quantity * item.Product.Price
		transaction.Amount += item.Quantity * item.Product.Price
	}

	order.Transaction = &transaction

	// Validate wallet balance
	if transaction.PaymentMethod == enums.TrxPaymentEcanteensPay && user.Wallet.Balance < transaction.Amount {
		return nil, errors.New("saldo anda tidak mencukupi")
	}

	if err := orderRepo(user, &cart, &order); err != nil {
		return nil, errors.New("gagal saat memproses pesanan")
	}

	return &order, nil
}

func updateOrderService(body *updateOrderScheme, id string, userId uint) error {
	var order models.Order

	if err := findOrderById(&order, id, userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("pesanan tidak ditemukan")
		}

		return err
	}

	switch body.Status {
	case "SUCCESS":
		if order.Status == enums.OrderStatusReady {
			return updateStatusOrder(&order, body)
		}
	case "CANCELED":
		if order.Status == enums.OrderStatusWaiting {
			return updateStatusOrder(&order, body)
		}
	default:
		return errors.New("status tidak diketahui")
	}

	return errors.New("pesanan gagal diperbarui")
}
