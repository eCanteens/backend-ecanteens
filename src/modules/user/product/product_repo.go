package product

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/gorm"
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

func findFavorite(user *models.User, userId uint, query map[string]string) error {
	return config.DB.Where("id = ?", userId).Preload("FavoriteProducts", func(db *gorm.DB) *gorm.DB {
		return db.Where("name ILIKE ?", "%"+query["search"]+"%").Preload("Category").Order(query["order"] + " " + query["direction"])
	}).Find(user).Error
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
