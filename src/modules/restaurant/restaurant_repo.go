package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func GetFavoriteRestoRepo(favorites *[]models.Favorite, userId uint) error {
	return config.DB.Where("user_id = ?", userId).Preload("Restaurant.Category").Preload("Restaurant.Location").Find(favorites).Error
}