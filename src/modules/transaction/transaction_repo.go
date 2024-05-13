package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func findCart(userId uint, cart *[]models.Cart) error {
	return config.DB.Where("user_id = ?", userId).Preload("Product.Restaurant.Category").Preload("Product.Restaurant.Location").Find(cart).Error
}

func findOneCart(userId uint, productId uint, cart *models.Cart) error {
	return config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).First(cart).Error
}

func updateCart(cart *models.Cart) error {
	return config.DB.Updates(cart).Error
}

func createCart(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}
