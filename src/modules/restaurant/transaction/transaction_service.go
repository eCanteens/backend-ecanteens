package transaction

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
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
			order.Status = enums.OrderStatusInProgress
			return update(&order)
		}
	case "READY":
		if order.Status == "INPROGRESS" {
			order.Status = enums.OrderStatusReady
			return update(&order)
		}
	case "CANCELED":
		if order.Status == "WAITING" {
			order.Status = enums.OrderStatusCanceled
			order.CancelReason = &body.Reason
			order.CancelBy = helpers.PointerTo(enums.OrderCancelByResto)
			order.Transaction.Status = enums.TrxStatusCanceled

			if order.Transaction.PaymentMethod == enums.TrxPaymentEcanteensPay {
				return updateOrderWithReturn(&order)
			}

			return updateOrderTransaction(&order)
		}
	default:
		return errors.New("status tidak diketahui")
	}

	return errors.New("pesanan gagal diperbarui")
}
