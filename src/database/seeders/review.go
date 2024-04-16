package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func ReviewSeeder() {
	var reviews []*models.Review

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			reviews = append(reviews, &models.Review{
				Rating:       gofakeit.UintRange(1, 5),
				UserId:       uint(j + 1),
				RestaurantId: uint(i + 1),
			})
		}
	}

	config.DB.Create(reviews)
	fmt.Println("Review Seeder created")
}