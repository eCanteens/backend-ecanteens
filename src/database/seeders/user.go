package seeders

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	var users []*models.User

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	chandra, _ := bcrypt.GenerateFromPassword([]byte("chandra123"), bcrypt.DefaultCost)
	admin, _ := bcrypt.GenerateFromPassword([]byte("password-admin"), bcrypt.DefaultCost)

	avatar := os.Getenv("BASE_URL") + "/public/dummy/avatar_user.png"

	users = append(users, &models.User{
		Name:     "Admin",
		Email:    "admin@ecanteens.com",
		Password: string(admin),
		Phone:    "-",
		RoleId:   1,
		Avatar:   &avatar,
	})

	for i := 0; i < 9; i++ {

		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    "08" + gofakeit.Numerify("##########"),
			Password: string(password),
			Avatar:   &avatar,
		})
	}

	users = append(users, &models.User{
		Name:     "Chandra",
		Email:    "mdutchand@gmail.com",
		Phone:    "085797175262",
		Password: string(chandra),
	})

	config.DB.Create(users)

	fmt.Println("User Seeder created")
}
