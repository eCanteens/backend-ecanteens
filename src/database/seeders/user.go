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

func UserSeeder() {
	var users []*models.User

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	chandra, _ := bcrypt.GenerateFromPassword([]byte("chandra123"), bcrypt.DefaultCost)
	admin, _ := bcrypt.GenerateFromPassword([]byte("password-admin"), bcrypt.DefaultCost)

	avatar := os.Getenv("BASE_URL") + "/public/assets/avatar-user.jpg"

	fakeUsers := []*models.User{
		{
			Name:     "Admin",
			Email:    "admin@ecanteens.com",
			Password: string(admin),
			Phone:    helpers.PointerTo("-"),
			RoleId:   1,
			Avatar:   os.Getenv("BASE_URL") + "/public/assets/avatar-admin.jpg",
		},
		{
			Name:     "Chandra",
			Email:    "mdutchand@gmail.com",
			Phone:    helpers.PointerTo("085797175262"),
			Password: string(chandra),
			Avatar:   avatar,
		},
		{
			Name:     "Muhajir",
			Email:    "amuhajir.syamlan@gmail.com",
			Phone:    helpers.PointerTo("088289570068"),
			Password: string(password),
			Avatar:   avatar,
		},
	}

	users = append(users, fakeUsers...)

	for i := 0; i < 10; i++ {
		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Phone:    helpers.PointerTo("08" + gofakeit.Numerify("##########")),
			Password: string(password),
			Avatar:   avatar,
		})
	}

	config.DB.Create(users)

	fmt.Println("User Seeder created")
}
