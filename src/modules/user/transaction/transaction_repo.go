package transaction

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	findOrder(result *pagination.Pagination[models.Order], userId uint, query *getOrderQS) error
	findOrderById(order *models.Order, id string, userId uint, preloads []string) error
	createOrder(wallet *models.Wallet, cart *models.Cart, order *models.Order) error
	updateOrderTransaction(order *models.Order, amountDst *models.Wallet) error

	findCart(userId uint, cart *[]models.Cart, preload bool) error
	findCartById(id uint, cart *models.Cart, userId uint, preload bool) error
	findCartByRestoId(id uint, cart *models.Cart, userId uint, preload bool) error
	createCart(data *models.Cart) error
	updateCartNote(id string, userId uint, notes string) error
	deleteCart(data *models.Cart) error

	createCartItem(data *models.CartItem) error
	updateCartItem(data *models.CartItem) error
	deleteCartItem(data *models.CartItem) error

	createReview(data *models.Review) error
	findOneProduct(product *models.Product, id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) findOrder(result *pagination.Pagination[models.Order], userId uint, query *getOrderQS) error {
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

func (r *repository) findOrderById(order *models.Order, id string, userId uint, preloads []string) error {
	tx := config.DB.Where("id = ?", id).Where("user_id = ?", userId)

	for _, preload := range preloads {
		tx.Preload(preload)
	}

	return tx.First(&order).Error
}

func (r *repository) createOrder(wallet *models.Wallet, cart *models.Cart, order *models.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if order.Transaction.PaymentMethod == enums.TrxPaymentEcanteensPay {
			// Update Buyer Balance
			if err := tx.Save(wallet).Error; err != nil {
				return err
			}
		}

		// Write order and transaction data into db
		if err := config.DB.Create(order).Error; err != nil {
			return err
		}

		// Delete cart & cart items data
		if err := r.deleteCart(cart); err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) updateOrderTransaction(order *models.Order, amountDst *models.Wallet) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// Update wallet
		if err := tx.Save(amountDst).Error; err != nil {
			return err
		}

		// Update order
		if err := tx.Save(order).Error; err != nil {
			return err
		}

		// Update transaction
		if err := tx.Save(order.Transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) findCart(userId uint, cart *[]models.Cart, preload bool) error {
	tx := config.DB.Where("user_id = ?", userId).Preload("Items")
	if preload {
		tx.Preload("Restaurant.Category").Preload("Items.Product")
	}

	return tx.Find(cart).Error
}

func (r *repository) findCartById(id uint, cart *models.Cart, userId uint, preload bool) error {
	tx := config.DB.Where("id = ?", id).Where("user_id = ?", userId).Preload("Items")
	if preload {
		tx.Preload("Restaurant.Category").Preload("Restaurant.Owner.Wallet").Preload("Items.Product")
	}

	return tx.First(cart).Error
}

func (r *repository) findCartByRestoId(id uint, cart *models.Cart, userId uint, preload bool) error {
	tx := config.DB.Where("restaurant_id = ?", id).Where("user_id = ?", userId).Preload("Items")
	if preload {
		tx.Preload("Restaurant.Category").Preload("Restaurant.Owner.Wallet").Preload("Items.Product")
	}

	return tx.First(cart).Error
}

func (r *repository) createCart(data *models.Cart) error {
	return config.DB.Create(data).Error
}

func (r *repository) updateCartNote(id string, userId uint, notes string) error {
	tx := config.DB.Model(&models.Cart{}).Where("id = ?", id).Where("user_id = ?", userId).Update("notes", notes)

	if tx.RowsAffected == 0 {
		return customerror.New("Keranjang tidak ditemukan", 404)
	}

	return tx.Error
}

func (r *repository) deleteCart(data *models.Cart) error {
	return config.DB.Unscoped().Delete(data).Error
}

func (r *repository) createCartItem(data *models.CartItem) error {
	return config.DB.Create(data).Error
}

func (r *repository) updateCartItem(data *models.CartItem) error {
	return config.DB.Save(data).Error
}

func (r *repository) deleteCartItem(data *models.CartItem) error {
	return config.DB.Unscoped().Delete(data).Error
}

func (r *repository) findOneProduct(product *models.Product, id uint) error {
	return config.DB.Where("id = ?", id).Preload("Restaurant").First(product).Error
}

func (r *repository) createReview(data *models.Review) error {
	return config.DB.Create(data).Error
}
