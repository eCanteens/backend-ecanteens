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
	return config.DB.Save(data).Error
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

func findCartById(id uint, cart *models.Cart, userId uint, preload bool) error {
	tx := config.DB.Where("id = ?", id).Where("user_id = ?", userId).Preload("Items")
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

func findOrderById(order *models.Order, id string, userId uint) error {
	return config.DB.Where("id = ?", id).Where("user_id = ?", userId).Preload("Transaction").First(&order).Error
}

func updateStatusOrder(order *models.Order, body *updateOrderScheme) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		status := enums.OrderStatus(body.Status)
		order.Status = status

		if status == enums.OrderStatusCanceled {
			order.CancelBy = helpers.PointerTo(enums.OrderCancelByUser)
			order.CancelReason = &body.Reason
		}

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		if order.Transaction.PaymentMethod == enums.TrxPaymentCash {
			if err := tx.Model(&models.Transaction{}).Where("id = ?", order.TransactionId).Update("status", body.Status).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func orderRepo(user *models.User, cart *models.Cart, order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if order.IsPreorder && order.Transaction.PaymentMethod == enums.TrxPaymentEcanteensPay {
			user.Wallet.Balance -= order.Transaction.Amount
			cart.Restaurant.Owner.Wallet.Balance += order.Transaction.Amount
			order.Transaction.Status = enums.TrxStatusSuccess

			// Update Buyer Wallet
			if err := tx.Updates(user.Wallet).Error; err != nil {
				return err
			}

			// Update Seller Wallet
			if err := tx.Updates(cart.Restaurant.Owner.Wallet).Error; err != nil {
				return err
			}
		}

		// Write order and transaction data into db
		if err := create(&order); err != nil {
			return err
		}

		// Delete cart & cart items data
		if err := deleteRecord(&cart); err != nil {
			return err
		}

		return nil
	})
}