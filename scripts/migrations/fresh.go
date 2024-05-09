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
	tables := []interface{}{
		&models.User{},
		&models.Location{},
		&models.RestaurantCategory{},
		&models.Restaurant{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Review{},
		&models.Favorite{},
		&models.ProductFeedback{},
	}

	config.DB.Migrator().DropTable(tables...)

	fmt.Println("Tables Dropped")

	config.DB.Migrator().CreateTable(tables...)

	fmt.Println("Tables Created")
}
