package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func FavoriteSeeder() {
	var favorites []*models.Favorite

	for i := 0; i < 11; i++ {
		favorites = append(favorites, &models.Favorite{
			UserId:       uint(i) + 1,
			RestaurantId: gofakeit.UintRange(1, 10),
		})
	}

	config.DB.Create(favorites)
	fmt.Println("Favorite Seeder created")
}