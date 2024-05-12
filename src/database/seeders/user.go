package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func UserSeeder() {
	var users []*models.User

	for i := 0; i < 9; i++ {
		phone := "08" + gofakeit.Numerify("##########")
		avatar := "/public/dummy/avatar_user.png"

		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    &phone,
			Password: "password",
			Avatar:   &avatar,
		})
	}

	users = append(users, &models.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "password",
	})

	users = append(users, &models.User{
		Name:     "Chandra",
		Email:    "chandra123@gmail.com",
		Password: "chandra123",
	})

	config.DB.Create(users)

	fmt.Println("User Seeder created")
}
