package main

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading env")
	}
}

func main() {
	config.ConnectDB()

	config.DB.Migrator().DropTable(
		&models.User{},
		&models.Location{},
		&models.RestaurantCategory{},
		&models.Restaurant{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Review{},
	)

	fmt.Println("Tables Dropped")
}
