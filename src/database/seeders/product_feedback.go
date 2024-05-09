package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func ProductFeedbackSeeder() {
	var productFeedbacks []*models.ProductFeedback

	for i := 0; i < 50; i++ {
		for j := 0; j < 11; j++ {
			productFeedbacks = append(productFeedbacks, &models.ProductFeedback{
				ProductId: uint(i + 1),
				UserId:    uint(j + 1),
				IsLike:    gofakeit.Bool(),
			})
		}
	}

	config.DB.Create(productFeedbacks)
	fmt.Println("Product Feedback Seeder created")
}
