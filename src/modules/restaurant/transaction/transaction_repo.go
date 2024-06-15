package transaction

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
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

func updateOrderWithReturn(order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// Update order
		if err := config.DB.Save(order).Error; err != nil {
			return err
		}

		// Update transaction
		if err := config.DB.Save(order.Transaction).Error; err != nil {
			return err
		}

		// Return balance to buyer
		order.User.Wallet.Balance += order.Transaction.Amount
		if err := config.DB.Save(order.User.Wallet).Error; err != nil {
			return err
		}

		return nil
	})
}

func updateOrderTransaction(order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := config.DB.Save(order).Error; err != nil {
			return err
		}

		if err := config.DB.Save(order.Transaction).Error; err != nil {
			return err
		}

		return nil
	})
}