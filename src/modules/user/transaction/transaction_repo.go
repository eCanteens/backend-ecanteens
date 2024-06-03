package transaction

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func create[T any](data *T) error {
	return config.DB.Create(data).Error
}

func findCart(userId uint, cart *[]models.Cart, preload bool) error {
	tx := config.DB.Where("user_id = ?", userId).Preload("Items")
	if preload {
		tx = tx.Preload("Restaurant.Category").Preload("Items.Product")
	}

	return tx.Find(cart).Error
}

func findOneProduct(product *models.Product, id uint) error {
	return config.DB.Where("id = ?", id).Preload("Restaurant").First(product).Error
}

func saveRecord[T any](data *T) error {
	return config.DB.Save(data).Error
}

func deleteRecord[T any](data *T) error {
	return config.DB.Unscoped().Delete(data).Error
}

func updateCartNote(id, notes string) error {
	tx := config.DB.Model(&models.Cart{}).Where("id = ?", id).Update("notes", notes)

	if tx.RowsAffected == 0 {
		return errors.New("keranjang tidak ditemukan")
	}

	return tx.Error
}
