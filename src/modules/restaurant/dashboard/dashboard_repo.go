package dashboard

import (
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/gorm"
)

type Repository interface {
	toggleOpen(restoId uint, isOpen bool) error
	summary(result *summaryDto, restoId uint) error
	latestHistory(result *models.Order, restoId uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) toggleOpen(restoId uint, isOpen bool) error {
	return config.DB.Model(models.Restaurant{}).Where("id = ?", restoId).Update("is_open", isOpen).Error
}

func (r *repository) summary(result *summaryDto, restoId uint) error {
	return config.DB.Table("orders").
		Select("COUNT(id) AS sum_today_trx, SUM(amount) AS sum_today_income").
		Where("restaurant_id = ?", restoId).
		Where("status = ?", "SUCCESS").
		Where("DATE(updated_at) = ?", time.Now().Format("2006-01-02")).
		Scan(result).Error
}

func (r *repository) latestHistory(result *models.Order, restoId uint) error {
	return config.DB.Where("restaurant_id = ?", restoId).
		Where("status = ?", "SUCCESS").
		Where("DATE(updated_at) = ?", time.Now().Format("2006-01-02")).
		Preload("Items").
		Preload("Transaction").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, avatar")
		}).
		Find(result).Error
}
