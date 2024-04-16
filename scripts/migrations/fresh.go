package main

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.Migrator().DropTable(
		&models.User{},
		&models.Location{},
		&models.RestaurantCategory{},
		&models.Restaurant{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Review{},
		&models.Favorite{},
	)

	fmt.Println("Tables Dropped")

	config.DB.Migrator().CreateTable(
		&models.User{},
		&models.Location{},
		&models.RestaurantCategory{},
		&models.Restaurant{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Review{},
		&models.Favorite{},
	)

	fmt.Println("Tables Created")
}
