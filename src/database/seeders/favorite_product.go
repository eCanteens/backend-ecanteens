package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func FavoriteProductSeeder() {
	var favorites []*models.FavoriteProduct

	for i := 0; i < 11; i++ {
		favorites = append(favorites, &models.FavoriteProduct{
			UserId:    uint(i) + 2,
			ProductId: gofakeit.UintRange(1, 50),
		})
	}

	config.DB.Create(favorites)
	fmt.Println("Favorite Restaurant Seeder created")
}
