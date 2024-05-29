package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"gorm.io/gorm"
)

func findFavorite(user *models.User, id uint, query map[string]string) error {
	return config.DB.Where("id = ?", id).Preload("FavoriteRestaurants", func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category").Order(query["order"] + " " + query["direction"])
	}).Find(user).Error
}

func find(pagination *helpers.Pagination, restaurants *[]models.Restaurant, query map[string]string) error {
	return pagination.Paginate(restaurants, &helpers.Params{
		Query:     config.DB.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category"),
		Page:      query["page"],
		Limit:     query["limit"],
		Order:     query["order"],
		Direction: query["direction"],
	})
}

func findOne(restaurant *models.Restaurant, id string) error {
	return config.DB.Where("id = ?", id).Preload("Category").First(restaurant).Error
}

func findRestosProducts(pagination *helpers.Pagination, products *[]models.Product, id string, query map[string]string) error {
	return pagination.Paginate(products, &helpers.Params{
		Query:     config.DB.Where("restaurant_id = ?", id),
		Page:      query["page"],
		Limit:     query["limit"],
		Order:     query["order"],
		Direction: query["direction"],
	})
}

func checkFavorite(userId uint, restaurantId uint) *[]models.FavoriteRestaurant {
	var favorites []models.FavoriteRestaurant

	config.DB.Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).Find(&favorites)

	return &favorites
}

func createFavorite(favorite *models.FavoriteRestaurant) error {
	return config.DB.Create(favorite).Error
}

func deleteFavorite(userId uint, restaurantId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).Delete(&models.FavoriteRestaurant{}).Error
}