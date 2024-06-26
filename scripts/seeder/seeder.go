package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/seeders"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	seeders.UserSeeder()
	seeders.RestaurantSeeder()
	seeders.ProductSeeder()
	seeders.ProductFeedbackSeeder()
	seeders.FavoriteRestaurantSeeder()
	seeders.FavoriteProductSeeder()
	seeders.TransactionSeeder()
	seeders.OrderSeeder()
	seeders.ReviewSeeder()
}