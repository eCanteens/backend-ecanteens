package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func findAllProduct(result *pagination.Pagination[models.Product], query *productQs, user *models.User) error {
	likeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = TRUE")

	dislikeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = FALSE")

	q := config.DB.Table("products").
		Select("products.*, (?) AS like, (?) AS dislike", likeCountSubquery, dislikeCountSubquery).
		Where("products.restaurant_id = ?", user.Restaurant.Id).
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Preload("Category")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}
