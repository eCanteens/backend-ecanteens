package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func getOrderService(restaurantId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	var result = pagination.New(models.Order{})

	if err := findOrder(result, restaurantId, query); err != nil {
		return nil, customerror.GormError(err, "Pesanan")
	}

	return result, nil
}

func updateOrderService(id string, user *models.User, body *updateOrderScheme) error {
	var order models.Order

	if err := findOrderById(id, *user.Restaurant.Id, &order); err != nil {
		return customerror.GormError(err, "Pesanan")
	}

	switch body.Status {
	case "INPROGRESS":
		if order.Status == "WAITING" {
			order.Status = enums.OrderStatusInProgress
			if err := update(&order); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	case "READY":
		if order.Status == "INPROGRESS" {
			order.Status = enums.OrderStatusReady
			if err := update(&order); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	case "CANCELED":
		if order.Status == "WAITING" {
			order.Status = enums.OrderStatusCanceled
			order.CancelReason = &body.Reason
			order.CancelBy = helpers.PointerTo(enums.OrderCancelByResto)
			order.Transaction.Status = enums.TrxStatusCanceled

			if order.Transaction.PaymentMethod == enums.TrxPaymentEcanteensPay {
				if err := updateOrderWithReturn(&order); err != nil {
					return customerror.GormError(err, "Pesanan")
				}

				return nil
			}

			if err := updateOrderTransaction(&order); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	default:
		return customerror.New("Status tidak diketahui", 400)
	}

	return customerror.New("Pesanan gagal diperbarui", 400)
}
