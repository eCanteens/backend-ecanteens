package transaction

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func getOrderService(restaurantId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	var result = pagination.New(models.Order{})

	if err := findOrder(result, restaurantId, query); err != nil {
		return nil, err
	}

	return result, nil
}

func updateOrderService(id string, user *models.User, body *updateOrderScheme) error {
	var order models.Order

	if err := findOrderById(id, *user.Restaurant.Id, &order); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("pesanan tidak ditemukan")
		}

		return err
	}

	switch body.Status {
	case "INPROGRESS":
		if order.Status == "WAITING" {
			if order.Transaction.PaymentMethod == transaction.ECANTEENSPAY && !order.IsPreorder {
				user.Wallet.Balance += order.Transaction.Amount

				if order.User.Wallet.Balance < order.Transaction.Amount {
					return errors.New("saldo pembeli tidak cukup")
				}

				order.User.Wallet.Balance -= order.Transaction.Amount

				order.Transaction.Status = transaction.SUCCESS

				if err := transferBalance(order.User.Wallet, user.Wallet, order.Transaction); err != nil {
					return err
				}
			}

			return updateOrderStatus(id, *user.Restaurant.Id, body.Status)
		}
	case "READY":
		if order.Status == "INPROGRESS" {
			return updateOrderStatus(id, *user.Restaurant.Id, body.Status)
		}
	case "CANCELED":
		if order.Status == "WAITING" {
			if order.Transaction.PaymentMethod == transaction.ECANTEENSPAY && order.IsPreorder {
				order.User.Wallet.Balance += order.Transaction.Amount

				if user.Wallet.Balance < order.Transaction.Amount {
					return errors.New("saldo anda tidak cukup untuk mengembalikan saldo pembeli")
				}

				user.Wallet.Balance -= order.Transaction.Amount

				order.Transaction.Status = transaction.CANCELED

				if err := transferBalance(user.Wallet, order.User.Wallet, order.Transaction); err != nil {
					return err
				}
			} else {
				order.Transaction.Status = transaction.CANCELED
				if err := update(order.Transaction); err != nil {
					return err
				}
			}

			return updateOrderStatus(id, *user.Restaurant.Id, body.Status)
		}
	default:
		return errors.New("status tidak diketahui")
	}

	return errors.New("pesanan gagal diperbarui")
}