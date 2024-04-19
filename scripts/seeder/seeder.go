package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/seeders"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	seeders.LocationSeeder()
	seeders.RestaurantSeeder()
	seeders.UserSeeder()
	seeders.ReviewSeeder()
	seeders.FavoriteSeeder()
	seeders.ProductSeeder()
}