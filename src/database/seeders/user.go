package seeders

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func UserSeeder() {
	var users []models.User

	for i := 0; i < 10; i++ {
		phone := gofakeit.Numerify("08##########")
		avatar := "/public/uploads/dummy/avatar_user.png"

		users = append(users, models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    &phone,
			Password: "password",
			Avatar:   &avatar,
			Balance:    (gofakeit.Number(5000, 200000) / 100) * 100,
		})

		config.DB.Create(&users)
	}
}
