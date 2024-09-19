package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	findFavorite(result *pagination.Pagination[models.Restaurant], userId uint, query *paginationQS) error
	find(result *pagination.Pagination[models.Restaurant], query *paginationQS) error
	findReviews(reviews *[]models.Review, restaurantId string, query *reviewQS) error
	findOne(restaurant *models.Restaurant, id string) error
	findRestosProducts(result *pagination.Pagination[models.Product], id string, query *paginationQS, categoryId uint) error
	checkFavorite(userId uint, restaurantId uint) *[]models.FavoriteRestaurant
	createFavorite(favorite *models.FavoriteRestaurant) error
	deleteFavorite(userId uint, restaurantId uint) error

	findProductCategories(categories *[]models.ProductCategory, categoryId string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) findFavorite(result *pagination.Pagination[models.Restaurant], userId uint, query *paginationQS) error {
	q := config.DB.
		Joins("JOIN favorite_restaurants fr ON fr.restaurant_id = restaurants.id").
		Joins("JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("JOIN reviews ON reviews.order_id = orders.id").
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

func (r *repository) find(result *pagination.Pagination[models.Restaurant], query *paginationQS) error {
	q := config.DB.
		Joins("JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count").
		Group("restaurants.id").
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

func (r *repository) findReviews(reviews *[]models.Review, restaurantId string, query *reviewQS) error {
	tx := config.DB.
		Joins("JOIN orders ON orders.id = reviews.order_id").
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

func (r *repository) findOne(restaurant *models.Restaurant, id string) error {
	return config.DB.Table("restaurants").
		Joins("JOIN orders ON orders.restaurant_id = restaurants.id").
		Joins("JOIN reviews ON reviews.order_id = orders.id").
		Select("restaurants.*, COALESCE(AVG(reviews.rating), 0) AS rating_avg, COUNT(reviews.*) AS rating_count").
		Where("restaurants.id = ?", id).
		Group("restaurants.id").
		Preload("Category").
		Preload("Owner", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, email, phone")
		}).
		First(restaurant).Error
}

func (r *repository) findRestosProducts(result *pagination.Pagination[models.Product], id string, query *paginationQS, categoryId uint) error {
	q := config.DB.
		Joins("JOIN product_feedbacks pf ON pf.product_id = products.id").
		Select("products.*, SUM(CASE WHEN pf.is_like = TRUE THEN 1 ELSE 0 END) AS like, SUM(CASE WHEN pf.is_like = FALSE THEN 1 ELSE 0 END) AS dislike").
		Where("products.restaurant_id = ?", id).
		Where("products.category_id = ?", categoryId).
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Group("products.id").
		Preload("Category")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func (r *repository) checkFavorite(userId uint, restaurantId uint) *[]models.FavoriteRestaurant {
	var favorites []models.FavoriteRestaurant

	config.DB.Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).Find(&favorites)

	return &favorites
}

func (r *repository) createFavorite(favorite *models.FavoriteRestaurant) error {
	return config.DB.Create(favorite).Error
}

func (r *repository) deleteFavorite(userId uint, restaurantId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("restaurant_id = ?", restaurantId).Delete(&models.FavoriteRestaurant{}).Error
}

func (r *repository) findProductCategories(categories *[]models.ProductCategory, categoryId string) error {
	if(categoryId == "") {
		return config.DB.Find(categories).Error
	} else {
		return config.DB.Where("id = ?", categoryId).Find(categories).Error
	}
}