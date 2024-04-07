package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func RestaurantSeeder() {
	var restaurants []models.Restaurant

	for i := 0; i < 10; i++ {
		restaurants = append(restaurants, models.Restaurant{
			Name:       gofakeit.AppName(),
			Phone:      gofakeit.Numerify("08##########"),
			LocationId: 1,
			Avatar:     "/public/uploads/dummy/avatar_resto.png",
			Banner:     "/public/uploads/dummy/banner.jpeg",
			Balance:    (gofakeit.Number(100000, 2000000) / 100) * 100,
			CategoryId: 1,
		})
	}

	config.DB.Create(&restaurants)
	fmt.Println("Restaurant Seeder created")
}

func RestaurantCategorySeeder() {
	restaurantCategory := models.RestaurantCategory{
		Name: "Jajanan",
	}

	config.DB.Create(&restaurantCategory)
	fmt.Println("Restaurant Category Seeder created")
}
