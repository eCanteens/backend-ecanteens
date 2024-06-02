package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func checkFeedback(userId uint, productId uint) (*[]models.ProductFeedback, error) {
	var feedbacks []models.ProductFeedback

	if err := config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return &feedbacks, nil
}

func updateFeedback(id uint, body *feedbackScheme) error {
	return config.DB.Model(&models.ProductFeedback{}).Where("id = ?", id).Update("is_like", *body.IsLike).Error
}

func createFeedback(feedback *models.ProductFeedback) error {
	return config.DB.Create(feedback).Error
}

func deleteFeedback(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.ProductFeedback{}).Error
}

func findFavorite(result *pagination.Pagination, userId uint, query *paginationQS) error {
	likeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = TRUE")

	dislikeCountSubquery := config.DB.Table("product_feedbacks").
		Select("COUNT(*)").
		Where("product_feedbacks.product_id = products.id AND product_feedbacks.is_like = FALSE")

	favoriteSubquery := config.DB.Table("favorite_products").
		Select("product_id").
		Where("user_id = ?", userId)

	q := config.DB.Table("products").
		Select("products.*, (?) AS like, (?) AS dislike", likeCountSubquery, dislikeCountSubquery).
		Where("products.id IN (?)", favoriteSubquery).
		Where("name ILIKE ?", "%"+query.Search+"%").
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

func checkFavorite(userId uint, ProductId uint) *[]models.FavoriteProduct {
	var favorites []models.FavoriteProduct

	config.DB.Where("user_id = ?", userId).Where("product_id = ?", ProductId).Find(&favorites)

	return &favorites
}

func createFavorite(favorite *models.FavoriteProduct) error {
	return config.DB.Create(favorite).Error
}

func deleteFavorite(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.FavoriteProduct{}).Error
}
