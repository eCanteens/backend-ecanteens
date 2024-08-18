package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Service interface {
	getOrder(restaurantId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error)
	updateOrder(id string, user *models.User, body *updateOrderScheme) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) getOrder(restaurantId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	var result = pagination.New(models.Order{})

	if err := s.repo.findOrder(result, restaurantId, query); err != nil {
		return nil, customerror.GormError(err, "Pesanan")
	}

	return result, nil
}

func (s *service) updateOrder(id string, user *models.User, body *updateOrderScheme) error {
	var order models.Order

	if err := s.repo.findOrderById(id, *user.Restaurant.Id, &order); err != nil {
		return customerror.GormError(err, "Pesanan")
	}

	switch body.Status {
	case "INPROGRESS":
		if order.Status == "WAITING" {
			order.Status = enums.OrderStatusInProgress
			if err := s.repo.updateOrder(&order); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	case "READY":
		if order.Status == "INPROGRESS" {
			order.Status = enums.OrderStatusReady
			if err := s.repo.updateOrder(&order); err != nil {
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
				if err := s.repo.updateOrderWithReturn(&order); err != nil {
					return customerror.GormError(err, "Pesanan")
				}

				return nil
			}

			if err := s.repo.updateOrderTransaction(&order); err != nil {
				return customerror.GormError(err, "Pesanan")
			}

			return nil
		}
	default:
		return customerror.New("Status tidak diketahui", 400)
	}

	return customerror.New("Pesanan gagal diperbarui", 400)
}
