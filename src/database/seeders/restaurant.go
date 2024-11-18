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
	restaurantCategories := []*models.RestaurantCategory{
		{ Name: "Makanan" },
		{ Name: "Minuman" },
		{ Name: "Jajanan" },
	}

	config.DB.Create(&restaurantCategories)
}

func RestaurantSeeder() {
	restaurantCategorySeeder()

	var restaurants []*models.Restaurant
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	restaurants = append(restaurants, &models.Restaurant{
		Name:       "Resto",
		Avatar:     os.Getenv("BASE_URL") + "/public/dummy/avatar-resto.png",
		Banner:     os.Getenv("BASE_URL") + "/public/dummy/banner.jpeg",
		CategoryId: 1,
		IsOpen:     true,
		Owner: &models.User{
			Name:     "Resto Owner",
			Email:    "resto@gmail.com",
			Phone:    helpers.PointerTo("081234567890"),
			Password: string(pw),
			Avatar:   os.Getenv("BASE_URL") + "/public/assets/avatar-user.jpg",
			RoleId:   3,
		},
	})

	for i := 0; i < 20; i++ {
		restaurants = append(restaurants, &models.Restaurant{
			Name:       gofakeit.AppName(),
			Avatar:     os.Getenv("BASE_URL") + "/public/dummy/avatar-resto.png",
			Banner:     os.Getenv("BASE_URL") + "/public/dummy/banner.jpeg",
			CategoryId: gofakeit.UintRange(1, 3),
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
