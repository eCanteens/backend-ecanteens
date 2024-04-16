package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func GetFavoriteRestoService(userId uint) (*[]models.Restaurant, error) {
	var favorites []models.Favorite

	if err := GetFavoriteRestoRepo(&favorites, userId); err != nil {
		return nil, err
	}

	var restaurants []models.Restaurant

	for _, v := range favorites {
		restaurants = append(restaurants, *v.Restaurant)
	}

	return &restaurants, nil
}