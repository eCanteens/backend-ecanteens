package transaction

import (
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

func transferBalance(src *models.Wallet, dst *models.Wallet, trx *models.Transaction, status enums.TransactionStatus) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		src.Balance -= trx.Amount
		dst.Balance += trx.Amount
		trx.Status = status

		if err := tx.Save(src).Error; err != nil {
			return err
		}

		if err := tx.Save(dst).Error; err != nil {
			return err
		}

		if err := tx.Save(trx).Error; err != nil {
			return err
		}

		return nil
	})
}

func updateOrderWithTransfer(seller *models.User, buyer *models.User, order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := config.DB.Save(&order).Error; err != nil {
			return err
		}

		switch order.Status {
		case enums.OrderStatusInProgress:
			return transferBalance(buyer.Wallet, seller.Wallet, order.Transaction, enums.TrxStatusSuccess)
		case enums.OrderStatusCanceled:
			return transferBalance(seller.Wallet, buyer.Wallet, order.Transaction, enums.TrxStatusCanceled)
		}
		return nil
	})
}

func updateOrderTransaction(order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := config.DB.Save(&order).Error; err != nil {
			return err
		}

		if err := config.DB.Save(&order.Transaction).Error; err != nil {
			return err
		}

		return nil
	})
}