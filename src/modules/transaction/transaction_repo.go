package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func findCart(userId uint, cart *[]models.Cart, preload bool) error {
	tx := config.DB.Where("user_id = ?", userId)
	if preload {
		tx = tx.Preload("Product.Restaurant.Category").Preload("Product.Restaurant.Location")
	}

	return tx.Find(cart).Error
}

func findOneCart(cart *models.Cart, userId uint, productId uint) error {
	return config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).First(&cart).Error
}

func saveCart(cart *models.Cart) error {
	return config.DB.Save(cart).Error
}

func deleteCart(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.Cart{}).Error
}