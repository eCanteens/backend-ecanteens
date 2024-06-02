package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

func findFavorite(result *pagination.Pagination, userId uint, query *paginationQS) error {
	avgSubquery := config.DB.Table("restaurant_reviews").
		Select("COALESCE(AVG(rating), 0)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	countSubquery := config.DB.Table("restaurant_reviews").
		Select("COUNT(*)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	favoriteSubquery := config.DB.Table("favorite_restaurants").
		Select("restaurant_id").
		Where("user_id = ?", userId)

	q := config.DB.Table("restaurants").
		Select("restaurants.*, (?) AS rating_avg, (?) AS rating_count", avgSubquery, countSubquery).
		Where("restaurants.id IN (?)", favoriteSubquery).
		Where("name ILIKE ?", "%"+query.Search+"%").
		Preload("Category")

	return result.New(&pagination.Params{
		Query:     q,
		Model:     &[]models.Restaurant{},
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func find(result *pagination.Pagination, query *paginationQS) error {
	avgSubquery := config.DB.Table("restaurant_reviews").
		Select("COALESCE(AVG(rating), 0)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	countSubquery := config.DB.Table("restaurant_reviews").
		Select("COUNT(*)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	q := config.DB.Table("restaurants").
		Select("restaurants.*, (?) AS rating_avg, (?) AS rating_count", avgSubquery, countSubquery).
		Where("restaurants.name ILIKE ?", "%"+query.Search+"%").
		Preload("Category")

	return result.New(&pagination.Params{
		Query:     q,
		Model:     &[]models.Restaurant{},
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func findReviews(reviews *[]models.RestaurantReview, restaurantId string, query *reviewQS) error {
	tx := config.DB.Where("restaurant_id = ?", restaurantId)

	if query.Filter != "" {
		tx = tx.Where("rating = ?", query.Filter)
	}

	return tx.Find(reviews).Error
}

func findOne(restaurant *models.Restaurant, id string) error {
	avgSubquery := config.DB.Table("restaurant_reviews").
		Select("COALESCE(AVG(rating), 0)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	countSubquery := config.DB.Table("restaurant_reviews").
		Select("COUNT(*)").
		Where("restaurant_reviews.restaurant_id = restaurants.id")

	return config.DB.Select("restaurants.*, (?) AS rating_avg, (?) AS rating_count", avgSubquery, countSubquery).
		Where("id = ?", id).
		Preload("Category").
		Preload("Owner", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, phone")
		}).
		First(restaurant).Error
}

func findRestosProducts(result *pagination.Pagination, id string, query *paginationQS) error {
	likeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = TRUE")

	dislikeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = FALSE")

	q := config.DB.Table("products").
		Select("products.*, (?) AS like, (?) AS dislike", likeCountSubquery, dislikeCountSubquery).
		Where("products.restaurant_id = ?", id).
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Preload("Category")

	return result.New(&pagination.Params{
		Query:     q,
		Model:     &[]models.Product{},
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
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
