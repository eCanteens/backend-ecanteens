package product

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	checkFeedback(userId uint, productId uint) (*models.ProductFeedback, error)
	updateFeedback(id uint, body *feedbackScheme) error
	createFeedback(feedback *models.ProductFeedback) error
	deleteFeedback(userId uint, productId uint) error
	findFavorite(result *pagination.Pagination[models.Product], userId uint, query *paginationQS) error
	checkFavorite(userId uint, ProductId uint) (*models.FavoriteProduct, error)
	createFavorite(favorite *models.FavoriteProduct) error
	deleteFavorite(userId uint, productId uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) checkFeedback(userId uint, productId uint) (*models.ProductFeedback, error) {
	var feedbacks models.ProductFeedback

	if err := config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).First(&feedbacks).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &feedbacks, nil
}

func (r *repository) updateFeedback(id uint, body *feedbackScheme) error {
	return config.DB.Model(&models.ProductFeedback{}).Where("id = ?", id).Update("is_like", *body.IsLike).Error
}

func (r *repository) createFeedback(feedback *models.ProductFeedback) error {
	return config.DB.Create(feedback).Error
}

func (r *repository) deleteFeedback(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.ProductFeedback{}).Error
}

func (r *repository) findFavorite(result *pagination.Pagination[models.Product], userId uint, query *paginationQS) error {
	q := config.DB.Table("products").
		Joins("LEFT JOIN product_feedbacks pf ON pf.product_id = products.id").
		Joins("LEFT JOIN favorite_products fp ON fp.product_id = products.id").
		Select("products.*, SUM(CASE WHEN pf.is_like = TRUE THEN 1 ELSE 0 END) AS like, SUM(CASE WHEN pf.is_like = FALSE THEN 1 ELSE 0 END) AS dislike").
		Group("products.id").
		Where("products.name ILIKE ?", "%"+query.Search+"%").
		Where("fp.user_id", userId).
		Preload("Category")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      query.Page,
		Limit:     query.Limit,
		Order:     query.Order,
		Direction: query.Direction,
	})
}

func (r *repository) checkFavorite(userId uint, ProductId uint) (*models.FavoriteProduct, error) {
	var favorites models.FavoriteProduct

	if err := config.DB.Where("user_id = ?", userId).Where("product_id = ?", ProductId).First(&favorites).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &favorites, nil
}

func (r *repository) createFavorite(favorite *models.FavoriteProduct) error {
	return config.DB.Create(favorite).Error
}

func (r *repository) deleteFavorite(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.FavoriteProduct{}).Error
}
