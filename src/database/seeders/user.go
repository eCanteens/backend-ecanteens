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
		phone := gofakeit.Numerify("08##########")
		avatar := "/public/uploads/dummy/avatar_user.png"

		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    &phone,
			Password: "password",
			Avatar:   &avatar,
			Balance:  (gofakeit.Number(5_000, 200_000) / 100) * 100,
		})
	}

	phone := "081234567890"
	avatar := "/public/uploads/dummy/avatar_user.png"

	users = append(users, &models.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Phone:    &phone,
		Password: "password",
		Avatar:   &avatar,
		Balance:  100_000,
	})

	config.DB.Create(users)

	fmt.Println("User Seeder created")
}