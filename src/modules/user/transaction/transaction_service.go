package transaction

import (
	"strconv"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/datatypes"
)

type Service interface {
	getCarts(user *models.User) (*[]models.Cart, error)
	getRestaurantCart(restaurantId uint, user *models.User) (*models.Cart, error)
	updateCart(id string, body *updateCartNoteScheme, userId uint) error
	addCart(user *models.User, body *addCartScheme) error
	getOrder(userId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error)
	order(body *orderScheme, user *models.User) (*models.Order, error)
	updateOrder(body *updateOrderScheme, id string, user *models.User) error
	postReview(body *postReviewScheme, id string, userId uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getCarts(user *models.User) (*[]models.Cart, error) {
	var carts []models.Cart

	if err := s.repo.findCart(*user.Id, &carts, true); err != nil {
		return nil, customerror.GormError(err, "Keranjang")
	}

	return &carts, nil
}
func (s *service) getRestaurantCart(restaurantId uint, user *models.User) (*models.Cart, error) {
	var cart models.Cart

	if err := s.repo.findCartByRestoId(restaurantId, &cart, *user.Id, true); err != nil {
		return nil, customerror.GormError(err, "Keranjang")
	}

	return &cart, nil
}

func (s *service) updateCart(id string, body *updateCartNoteScheme, userId uint) error {
	if err := s.repo.updateCartNote(id, userId, body.Notes); err != nil {
		return customerror.GormError(err, "Keranjang")
	}

	return nil
}

func (s *service) addCart(user *models.User, body *addCartScheme) error {
	var product models.Product
	var carts []models.Cart

	if err := s.repo.findOneProduct(&product, body.ProductId); err != nil {
		return customerror.GormError(err, "produk")
	}

	if product.Stock == 0 {
		return customerror.New("Stok produk habis", 400)
	}

	if !product.Restaurant.IsOpen {
		return customerror.New("Restoran sedang tutup", 400)
	}

	if err := s.repo.findCart(*user.Id, &carts, false); err != nil {
		return customerror.GormError(err, "Keranjang")
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

			if err := s.repo.createCart(cart); err != nil {
				return customerror.GormError(err, "Keranjang")
			}

			return nil
		}

		// but quantity is 0 then return error
		return customerror.New("Produk tidak ditemukan di keranjang", 404)
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

			if err := s.repo.createCartItem(cartItem); err != nil {
				return customerror.GormError(err, "Keranjang")
			}
			return nil
		}

		// and quantity is 0 then return error
		return customerror.New("Produk tidak ditemukan di keranjang", 404)
	} else {
		// if cart & cart item found
		if *body.Quantity > 0 {
			// if quantity not 0 then update cart item
			cartItem.Quantity = *body.Quantity
			if err := s.repo.updateCartItem(cartItem); err != nil {
				return customerror.GormError(err, "Keranjang")
			}

			return nil
		} else {
			// if quantity is 0
			if len(cart.Items) > 1 {
				// if cart items more than 1 then delete cart item
				if err := s.repo.deleteCartItem(cartItem); err != nil {
					return customerror.GormError(err, "Keranjang")
				}

				return nil
			} else {
				// if cart items just 1 then delete cart
				if err := s.repo.deleteCart(cart); err != nil {
					return customerror.GormError(err, "Keranjang")
				}

				return nil
			}
		}
	}
}

func (s *service) getOrder(userId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	result := pagination.New(models.Order{})

	if err := s.repo.findOrder(result, userId, query); err != nil {
		return nil, customerror.GormError(err, "Pesanan")
	}

	return result, nil
}

func (s *service) order(body *orderScheme, user *models.User) (*models.Order, error) {
	var cart models.Cart

	// Find cart
	if err := s.repo.findCartById(body.CartId, &cart, *user.Id, true); err != nil {
		return nil, customerror.GormError(err, "Keranjang")
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
		return nil, customerror.New("Restoran sedang tutup", 404)
	}

	var fullfilmentDate *time.Time

	// Validate fullfilment date if preorder
	if *body.IsPreorder {
		date, err := time.Parse(time.DateTime, body.FullfilmentDate)
		if err != nil {
			return nil, customerror.New("Format waktu tidak valid", 400)
		}

		fullfilmentDate = &date
	}

	// Create order
	order := models.Order{
		UserId:          cart.UserId,
		RestaurantId:    cart.RestaurantId,
		Notes:           body.Notes,
		Status:          enums.OrderStatusWaiting,
		IsPreorder:      *body.IsPreorder,
		FullfilmentDate: fullfilmentDate,
	}

	// Insert cart items into order items & calculate amount
	for _, item := range cart.Items {
		if item.Product.Stock == 0 {
			return nil, customerror.New("Stok produk habis", 400)
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
	if transaction.PaymentMethod == enums.TrxPaymentEcanteensPay {
		if user.Wallet.Balance < transaction.Amount {
			return nil, customerror.New("Saldo anda tidak mencukupi", 400)
		}

		user.Wallet.Balance -= transaction.Amount
	}

	if err := s.repo.createOrder(user.Wallet, &cart, &order); err != nil {
		return nil, customerror.New("Gagal saat memproses pesanan", 500)
	}

	return &order, nil
}

func (s *service) updateOrder(body *updateOrderScheme, id string, user *models.User) error {
	var order models.Order

	if err := s.repo.findOrderById(&order, id, *user.Id, []string{"Transaction", "Restaurant.Owner.Wallet"}); err != nil {
		return customerror.GormError(err, "Pesanan")
	}

	switch body.Status {
	case "SUCCESS":
		if order.Status == enums.OrderStatusReady {
			order.Status = enums.OrderStatusSuccess

			order.Transaction.Status = enums.TrxStatusSuccess

			// Return balance to user
			user.Wallet.Balance += order.Transaction.Amount

			if err := s.repo.updateOrderTransaction(&order, user.Wallet); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	case "CANCELED":
		if order.Status == enums.OrderStatusWaiting {
			order.Status = enums.OrderStatusCanceled
			order.CancelBy = helpers.PointerTo(enums.OrderCancelByUser)
			order.CancelReason = &body.Reason

			order.Transaction.Status = enums.TrxStatusCanceled

			// Release balance to resto
			order.Restaurant.Owner.Wallet.Balance += order.Transaction.Amount

			if err := s.repo.updateOrderTransaction(&order, order.Restaurant.Owner.Wallet); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	default:
		return customerror.New("Status tidak diketahui", 400)
	}

	return customerror.New("Pesanan gagal diperbarui", 400)
}

func (s *service) postReview(body *postReviewScheme, id string, userId uint) error {
	var order models.Order
	if err := s.repo.findOrderById(&order, id, userId, []string{"Review"}); err != nil {
		return customerror.GormError(err, "Pesanan")
	}

	if order.Status != enums.OrderStatusSuccess {
		return customerror.New("Tidak bisa mengirim ulasan jika pesanan belum selesai", 400)
	}

	if order.Review != nil {
		return customerror.New("Anda sudah mengirim ulasan", 400)
	}

	orderId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return customerror.New("Id pesanan tidak valid", 400)
	}

	review := models.Review{
		OrderId: uint(orderId),
		Rating:  body.Rating,
		Tags:    datatypes.NewJSONType(body.Tags),
		Comment: body.Comment,
	}

	if err := s.repo.createReview(&review); err != nil {
		return customerror.GormError(err, "Ulasan")
	}

	return nil
}
