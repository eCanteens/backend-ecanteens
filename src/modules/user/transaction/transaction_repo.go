package transaction

import (
	"errors"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func create[T any](data *T) error {
	return config.DB.Create(data).Error
}

func update[T any](data *T) error {
	return config.DB.Updates(data).Error
}

func deleteRecord[T any](data *T) error {
	return config.DB.Unscoped().Delete(data).Error
}

func findCart(userId uint, cart *[]models.Cart, preload bool) error {
	tx := config.DB.Where("user_id = ?", userId).Preload("Items")
	if preload {
		tx.Preload("Restaurant.Category").Preload("Items.Product")
	}

	return tx.Find(cart).Error
}

func findCartById(id uint, cart *models.Cart, preload bool) error {
	tx := config.DB.Where("id = ?", id).Preload("Items")
	if preload {
		tx.Preload("Restaurant.Category").Preload("Restaurant.Owner.Wallet").Preload("Items.Product")
	}

	return tx.First(cart).Error
}

func findOneProduct(product *models.Product, id uint) error {
	return config.DB.Where("id = ?", id).Preload("Restaurant").First(product).Error
}

func findOrder(result *pagination.Pagination[models.Order], userId uint, query *getOrderQS) error {
	tx := config.DB.Where("user_id = ?", userId).Preload("Items").Preload("Transaction").Preload("Restaurant")

	if query.Filter == "1" {
		// Berlangsung
		tx.Where(
			config.DB.
				Where(config.DB.Where("status = ?", "WAITING").Where("is_preorder = ?", false)).
				Or(config.DB.Where("status = ?", "INPROGRESS").Where("is_preorder = ?", false)).
				Or(config.DB.Where("status = ?", "READY")).
				Or(config.DB.Where("is_preorder = ?", true).Where("fullfilment_date <= ?", time.Now()).Not("status = ?", "SUCCESS").Not("status = ?", "CANCELED")),
		)
	} else if query.Filter == "2" {
		// Dijadwalkan
		tx.Where(
			config.DB.
				Where("is_preorder = ?", true).
				Where("fullfilment_date > ?", time.Now()).
				Not("status = ?", "SUCCESS").
				Not("status = ?", "CANCELED"),
		)
	} else if query.Filter == "3" {
		// Riwayat
		tx.Where(
			config.DB.Where("status = ?", "SUCCESS").Or("status = ?", "CANCELED"),
		)
	}

	return result.Execute(&pagination.Params{
		Query:     tx,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func updateCartNote(id string, userId uint, notes string) error {
	tx := config.DB.Model(&models.Cart{}).Where("id = ?", id).Where("user_id = ?", userId).Update("notes", notes)

	if tx.RowsAffected == 0 {
		return errors.New("keranjang tidak ditemukan")
	}

	return tx.Error
}

func cancelOrderById(reason, id string, userId uint) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var order models.Order

		if err := tx.Where("id = ?", id).Where("user_id = ?", userId).Where("status = ?", enums.OrderStatusWaiting).First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("pesanan tidak ditemukan")
			}

			return err
		}

		order.Status = enums.OrderStatusCanceled
		order.CancelBy = helpers.PointerTo(enums.OrderCancelByUser)
		order.CancelReason = &reason

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Transaction{}).Where("id = ?", order.TransactionId).Update("status", enums.TrxStatusCanceled).Error; err != nil {
			return err
		}

		return nil
	})
}
