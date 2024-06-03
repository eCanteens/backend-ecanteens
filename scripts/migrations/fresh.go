package main

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	tables, _ := config.DB.Migrator().GetTables()

	_tables := []interface{}{}
	for _, v := range tables {
		_tables = append(_tables, v)
	}
	config.DB.Migrator().DropTable(_tables...)

	fmt.Println("Tables Dropped")

	config.DB.Migrator().CreateTable(
		&models.Wallet{},
		&models.User{},
		&models.RestaurantCategory{},
		&models.Restaurant{},
		&models.RestaurantReview{},
		&models.FavoriteRestaurant{},
		&models.ProductCategory{},
		&models.Product{},
		&models.ProductFeedback{},
		&models.FavoriteProduct{},
		&models.Cart{},
		&models.CartItem{},
		&models.Transaction{},
		&models.Order{},
		&models.OrderItem{},
	)

	fmt.Println("Tables Created")
}
