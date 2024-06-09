package seeders

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"golang.org/x/crypto/bcrypt"
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

	for i := 0; i < 5; i++ {
		pw, _ := bcrypt.GenerateFromPassword([]byte("password-admin"), bcrypt.DefaultCost)

		restaurants = append(restaurants, &models.Restaurant{
			Name:       gofakeit.AppName(),
			Avatar:     os.Getenv("BASE_URL") + "/public/dummy/avatar-resto.png",
			Banner:     os.Getenv("BASE_URL") + "/public/dummy/banner.jpeg",
			CategoryId: 1,
			IsOpen:     true,
			Owner: &models.User{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Phone:    helpers.PointerTo("08" + gofakeit.Numerify("##########")),
				Password: string(pw),
				Avatar:   os.Getenv("BASE_URL") + "/public/assets/avatar-user.jpg",
				RoleId:   3,
			},
		})
	}

	config.DB.Create(restaurants)

	fmt.Println("Restaurant Seeder created")
}
