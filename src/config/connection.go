package config

import (
	"fmt"
	"os"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(autoMigrate bool) {
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	if autoMigrate {
		database.AutoMigrate(
			&models.User{},
			&models.Location{},
			&models.RestaurantCategory{},
			&models.Restaurant{},
			&models.ProductCategory{},
			&models.Product{},
			&models.Review{},
			&models.Favorite{},
		)
	}

	DB = database
}