package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func findFavorite(user *models.User, id uint, query map[string]string) error {
	return config.DB.Where("id = ?", id).Preload("FavoriteRestaurants", func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category").Order(query["order"] + " " + query["direction"])
	}).Find(user).Error
}

func find(result *pagination.Pagination, query map[string]string) error {
	return result.New(&pagination.Params{
		Query:     config.DB.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category"),
		Model:     &[]models.Restaurant{},
		Page:      query["page"],
		Limit:     query["limit"],
		Order:     query["order"],
		Direction: query["direction"],
	})
}

func findOne(restaurant *models.Restaurant, id string) error {
	return config.DB.Where("id = ?", id).Preload("Category").First(restaurant).Error
}

func findRestosProducts(result *pagination.Pagination, id string, query map[string]string) error {
	return result.New(&pagination.Params{
		Query:     config.DB.Where("restaurant_id = ?", id),
		Model:     &[]models.Product{},
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
