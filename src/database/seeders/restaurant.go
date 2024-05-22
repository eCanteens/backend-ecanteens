package seeders

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func restaurantCategorySeeder() {
	restaurantCategory := models.RestaurantCategory{
		Name: "Jajanan",
	}

	config.DB.Create(&restaurantCategory)
}

func RestaurantSeeder() {
	restaurantCategorySeeder()

	var restaurants []*models.Restaurant

	for i := 0; i < 10; i++ {
		restaurants = append(restaurants, &models.Restaurant{
			Name:       gofakeit.AppName(),
			Phone:      "08" + gofakeit.Numerify("##########"),
			Email:      gofakeit.Email(),
			LocationId: 1,
			Avatar:     os.Getenv("BASE_URL") + "/public/dummy/avatar_resto.png",
			Banner:     os.Getenv("BASE_URL") + "/public/dummy/banner.jpeg",
			CategoryId: 1,
		})
	}

	config.DB.Create(restaurants)

	fmt.Println("Restaurant Seeder created")
}
