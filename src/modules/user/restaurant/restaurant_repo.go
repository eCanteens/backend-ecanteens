package restaurant

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	findPopular(result *[]models.Restaurant) error
	findRecentResto(result *[]models.Restaurant, userId uint) error

	findFavorite(result *pagination.Pagination[models.Restaurant], userId uint, query *paginationQS) error
	find(result *pagination.Pagination[models.Restaurant], query *paginationQS, categoryId uint) error
	findReviews(reviews *[]models.Review, restaurantId string, query *reviewQS) error
	findOne(restaurant *models.Restaurant, id string, userId uint) error
	findRestosProducts(result *pagination.Pagination[models.Product], id string, query *paginationQS, categoryId uint, userId uint) error
	checkFavorite(userId uint, restaurantId uint) (*models.FavoriteRestaurant, error)
	createFavorite(favorite *models.FavoriteRestaurant) error
	deleteFavorite(userId uint, restaurantId uint) error

	findProductCategories(categories *[]models.ProductCategory, categoryId string) error
	findRestoCategories(categories *[]models.RestaurantCategory, categoryId string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) findPopular(result *[]models.Restaurant) error {
	return config.DB.
		Joins("LEFT JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count").
		Group("restaurants.id").
		Order("COUNT(orders.id) desc").
		Limit(5).
		Find(result).Error;
}

func (r *repository) findRecentResto(result *[]models.Restaurant, userId uint) error {
	return config.DB.
		Joins("LEFT JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.id) AS rating_count, SUM(CASE WHEN orders.user_id = ? THEN 1 ELSE 0 END) AS user_order_count", userId).
		Group("restaurants.id").
		Order("user_order_count DESC").
		Limit(2).
		Find(result).Error;
}

func (r *repository) findFavorite(result *pagination.Pagination[models.Restaurant], userId uint, query *paginationQS) error {
	q := config.DB.
		Joins("LEFT JOIN favorite_restaurants fr ON fr.restaurant_id = restaurants.id").
		Joins("LEFT JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count").
		Group("restaurants.id").
		Where("fr.user_id", userId).
		Where("restaurants.name ILIKE ?", "%"+query.Search+"%").
		Preload("Category")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func (r *repository) find(result *pagination.Pagination[models.Restaurant], query *paginationQS, categoryId uint) error {
	q := config.DB.
		Joins("LEFT JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count").
		Group("restaurants.id").
		Where("restaurants.name ILIKE ?", "%"+query.Search+"%").
		Where("restaurants.category_id = ?", categoryId)

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func (r *repository) findReviews(reviews *[]models.Review, restaurantId string, query *reviewQS) error {
	tx := config.DB.
		Joins("LEFT JOIN orders ON orders.id = reviews.order_id").
		Preload("Order", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, created_at, user_id")
		}).
		Preload("Order.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, avatar")
		}).
		Where("orders.restaurant_id = ?", restaurantId)

	if query.Filter != "" {
		tx.Where("reviews.rating = ?", query.Filter)
	}

	return tx.Find(reviews).Error
}

func (r *repository) findOne(restaurant *models.Restaurant, id string, userId uint) error {
	return config.DB.Table("restaurants").
		Joins("LEFT JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Joins("LEFT JOIN favorite_restaurants fr ON fr.restaurant_id = restaurants.id AND fr.user_id = ?", userId).
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count, (fr.id IS NOT NULL) AS is_favorited").
		Where("restaurants.id = ?", id).
		Group("restaurants.id").
		Group("fr.id").
		Preload("Category").
		Preload("Owner", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, phone")
		}).
		First(restaurant).Error
}

func (r *repository) findRestosProducts(result *pagination.Pagination[models.Product], id string, query *paginationQS, categoryId uint, userId uint) error {
	q := config.DB.
		Joins("LEFT JOIN product_feedbacks pf ON pf.product_id = products.id").
		Select("products.*, SUM(CASE WHEN pf.is_like = TRUE THEN 1 ELSE 0 END) AS like, SUM(CASE WHEN pf.is_like = FALSE THEN 1 ELSE 0 END) AS dislike, "+
			"(SELECT CASE "+
			"WHEN EXISTS (SELECT 1 FROM product_feedbacks WHERE product_id = products.id AND user_id = ? AND is_like = TRUE) THEN TRUE "+
			"WHEN EXISTS (SELECT 1 FROM product_feedbacks WHERE product_id = products.id AND user_id = ? AND is_like = FALSE) THEN FALSE "+
			"ELSE NULL END) AS is_liked", userId, userId).
		Where("products.restaurant_id = ?", id).
		Where("products.category_id = ?", categoryId).
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Where("products.is_active = ?", true).
		Group("products.id")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func (r *repository) checkFavorite(userId uint, restaurantId uint) (*models.FavoriteRestaurant, error) {
	var favorites models.FavoriteRestaurant

	if err := config.DB.Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).First(&favorites).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &favorites, nil
}

func (r *repository) createFavorite(favorite *models.FavoriteRestaurant) error {
	return config.DB.Create(favorite).Error
}

func (r *repository) deleteFavorite(userId uint, restaurantId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).Delete(&models.FavoriteRestaurant{}).Error
}

func (r *repository) findProductCategories(categories *[]models.ProductCategory, categoryId string) error {
	if categoryId == "" {
		return config.DB.Find(categories).Error
	} else {
		return config.DB.Where("id = ?", categoryId).Find(categories).Error
	}
}

func (r *repository) findRestoCategories(categories *[]models.RestaurantCategory, categoryId string) error {
	if categoryId == "" {
		return config.DB.Find(categories).Error
	} else {
		return config.DB.Where("id = ?", categoryId).Find(categories).Error
	}
}
