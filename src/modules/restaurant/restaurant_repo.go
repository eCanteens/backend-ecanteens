package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func FindFavoriteResto(user *models.User, id uint, query map[string]string) error {
	return config.DB.Where("id = ?", id).Preload("FavoriteRestaurant", func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category").Order(query["order"] + " " + query["direction"])
	}).Find(user).Error
}

func FindResto(_pagination *pagination.Pagination, restaurants *[]models.Restaurant, query map[string]string) error {
	return _pagination.Paginate(restaurants, &pagination.Params{
		Query:     config.DB.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category"),
		Page:      query["page"],
		Limit:     query["limit"],
		Order:     query["order"],
		Direction: query["direction"],
	})
}

func FindOneResto(restaurant *models.Restaurant, id string) error {
	return config.DB.Where("id = ?", id).Preload("Category").Preload("Location").First(restaurant).Error
}

func FindRestosProducts(_pagination *pagination.Pagination, products *[]models.Product, id string, query map[string]string) error {
	return _pagination.Paginate(products, &pagination.Params{
		Query:     config.DB.Where("restaurant_id = ?", id),
		Page:      query["page"],
		Limit:     query["limit"],
		Order:     query["order"],
		Direction: query["direction"],
	})
}