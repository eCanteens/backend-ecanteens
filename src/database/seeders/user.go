package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	var users []*models.User

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	chandra, _ := bcrypt.GenerateFromPassword([]byte("nadia123"), bcrypt.DefaultCost)
	admin, _ := bcrypt.GenerateFromPassword([]byte("password-admin"), bcrypt.DefaultCost)

	users = append(users, &models.User{
		Name:     "Admin",
		Email:    "admin@ecanteens.com",
		Password: string(admin),
		RoleId:   1,
	})

	for i := 0; i < 9; i++ {
		phone := "08" + gofakeit.Numerify("##########")
		avatar := "/public/dummy/avatar_user.png"

		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    &phone,
			Password: string(password),
			Avatar:   &avatar,
		})
	}

	users = append(users, &models.User{
		Name:     "Chandra",
		Email:    "mdutchand@gmail.com",
		Password: string(chandra),
	})

	config.DB.Create(users)

	fmt.Println("User Seeder created")
}
