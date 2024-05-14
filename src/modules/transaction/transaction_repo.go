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

func updateCart(userId uint, productId uint, cart *models.Cart) error {
	return config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).Updates(cart).Error
}

func updateManyCart(cart *[]*models.Cart) error {
	tx := config.DB.Begin()
	
	for _, c := range *cart {
		tx.Save(c)
	}

	return tx.Commit().Error
}

func deleteCart(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.Cart{}).Error
}