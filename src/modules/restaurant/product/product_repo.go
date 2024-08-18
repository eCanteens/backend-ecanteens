package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Repository interface {
	create(product *models.Product) error
	findAll(result *pagination.Pagination[models.Product], query *productQs, user *models.User) error
	update(product *models.Product, id string) error
	delete(productId uint, restaurantId uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *repository) findAll(result *pagination.Pagination[models.Product], query *productQs, user *models.User) error {
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

func (r *repository) update(product *models.Product, id string) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *repository) delete(productId uint, restaurantId uint) error {
	// return config.DB.Where("id = ?", productId).Where("restaurant_id = ?", restaurantId).Delete(&models.Product{}).Error

	if affected := config.DB.Where("id = ?", productId).Where("restaurant_id = ?", restaurantId).Delete(&models.Product{}).RowsAffected; affected == 0 {
		return customerror.New("Menu tidak ditemukan", 404)
	}
	return nil
}
