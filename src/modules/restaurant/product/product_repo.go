package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Repository interface {
	create(product *models.Product) error
	find(result *pagination.Pagination[models.Product], query *paginationQS, categoryId uint, restoId uint) error
	update(product *models.Product, id string) error
	delete(productId uint, restaurantId uint) error

	findProductCategories(categories *[]models.ProductCategory, categoryId string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *repository) find(result *pagination.Pagination[models.Product], query *paginationQS, categoryId uint, restoId uint) error {
	q := config.DB.
		Joins("LEFT JOIN product_feedbacks pf ON pf.product_id = products.id").
		Select("products.*, SUM(CASE WHEN pf.is_like = TRUE THEN 1 ELSE 0 END) AS like, SUM(CASE WHEN pf.is_like = FALSE THEN 1 ELSE 0 END) AS dislike").
		Where("products.restaurant_id = ?", restoId).
		Where("products.category_id = ?", categoryId).
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Group("products.id")

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

func (r *repository) findProductCategories(categories *[]models.ProductCategory, categoryId string) error {
	if categoryId == "" {
		return config.DB.Find(categories).Error
	} else {
		return config.DB.Where("id = ?", categoryId).Find(categories).Error
	}
}