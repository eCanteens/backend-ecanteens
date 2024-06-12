package transaction

import (
	"errors"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func update[T any](data *T) error {
	return config.DB.Save(data).Error
}

func findOrder(result *pagination.Pagination[models.Order], restaurantId uint, query *getOrderQS) error {
	tx := config.DB.Where("restaurant_id = ?", restaurantId).
		Preload("Items").
		Preload("Transaction").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, phone")
		})

	if query.Filter == "1" {
		// Masuk
		tx.Where(
			config.DB.
				Where("status = ?", "WAITING").
				Or(config.DB.Where("status = ?", "INPROGRESS").Where("is_preorder = ?", true).Where("fullfilment_date > ?", time.Now())),
		)
	} else if query.Filter == "2" {
		// Berlangsung
		tx.Where(
			config.DB.
				Where(config.DB.Where("status = ?", "INPROGRESS").Where("is_preorder = ?", false)).
				Or("status = ?", "READY").
				Or(config.DB.Where("status = ?", "INPROGRESS").Where("is_preorder = ?", true).Where("fullfilment_date <= ?", time.Now())),
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

func findOrderById(id string, restaurantId uint, order *models.Order) error {
	return config.DB.
		Where("id = ?", id).
		Where("restaurant_id = ?", restaurantId).
		Preload("Transaction").
		Preload("User.Wallet").
		First(order).Error
}

func updateOrderStatus(id string, restaurantId uint, status string) error {
	if affected := config.DB.Model(&models.Order{}).
		Where("id = ?", id).
		Where("restaurant_id = ?", restaurantId).
		Update("status", status).
		RowsAffected; affected == 0 {
		return errors.New("pesanan gagal diperbarui")
	}

	return nil
}

func transferBalance(src *models.Wallet, dst *models.Wallet, trx *models.Transaction, status enums.TransactionStatus) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		src.Balance -= trx.Amount
		dst.Balance += trx.Amount
		trx.Status = status

		if err := tx.Updates(src).Error; err != nil {
			return err
		}

		if err := tx.Updates(dst).Error; err != nil {
			return err
		}

		if err := tx.Updates(trx).Error; err != nil {
			return err
		}

		return nil
	})
}
