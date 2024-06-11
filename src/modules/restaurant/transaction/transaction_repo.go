package transaction

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func findOrder(result *pagination.Pagination[models.Order], restaurantId uint, query *getOrderQS) error {
	tx := config.DB.Debug().Where("restaurant_id = ?", restaurantId).
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
