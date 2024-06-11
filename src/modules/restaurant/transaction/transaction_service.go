package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func getOrderService(restaurantId uint, query *getOrderQS) (*pagination.Pagination[models.Order], error) {
	var result = pagination.New(models.Order{})

	if err := findOrder(result, restaurantId, query); err != nil {
		return nil, err
	}

	return result, nil
}