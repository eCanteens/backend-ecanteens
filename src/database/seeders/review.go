package seeders

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func ReviewSeeder() {
	var reviews []*models.Review

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			reviews = append(reviews, &models.Review{
				Rating:       helpers.RandomElement([]float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}),
				UserId:       uint(j + 1),
				RestaurantId: uint(i + 1),
			})
		}
	}

	config.DB.Create(reviews)
	fmt.Println("Review Seeder created")
}