package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func createCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}