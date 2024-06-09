package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"gorm.io/datatypes"
)

func ReviewSeeder() {
	var orders []models.Order
	var reviews []*models.Review

	config.DB.Where("status = ?", "SUCCESS").Find(&orders)

	for _, ord := range orders {
		reviews = append(reviews, &models.Review{
			OrderId: *ord.Id,
			Tags:    datatypes.NewJSONType([]string{"Kemasan", "Kebersihan"}),
			Rating:  helpers.RandomElement([]float32{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}),
			Comment: gofakeit.Comment(),
		})
	}

	config.DB.Create(reviews)
	fmt.Println("Review Seeder created")
}
