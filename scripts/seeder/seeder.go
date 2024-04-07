package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/seeders"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB(false)
}

func main() {
	seeders.LocationSeeder()
	seeders.RestaurantCategorySeeder()
	seeders.RestaurantSeeder()
	seeders.UserSeeder()
}