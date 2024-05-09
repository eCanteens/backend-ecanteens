package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func checkFeedback(userId uint, productId uint) (*[]models.ProductFeedback, error) {
	var feedbacks []models.ProductFeedback

	if err := config.DB.Where("user_id = ?", userId).Where("product_id = ?", productId).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return &feedbacks, nil
}

func updateFeedback(id uint, body *FeedbackScheme) error {
	return config.DB.Model(&models.ProductFeedback{}).Where("id = ?", id).Update("like", *body.Like).Error
}

func createFeedback(feedback *models.ProductFeedback) error {
	return config.DB.Create(feedback).Error
}

func deleteFeedback(userId uint, productId uint) error {
	return config.DB.Unscoped().Where("user_id = ?", userId).Where("product_id = ?", productId).Delete(&models.ProductFeedback{}).Error
}

func CountLike(productId uint, count *int64) error {
	return config.DB.Where("product_id = ?", productId).Where("like = ?", true).Count(count).Error
}

func CountDislike(productId uint, count *int64) error {
	return config.DB.Where("product_id = ?", productId).Where("like = ?", false).Count(count).Error
}