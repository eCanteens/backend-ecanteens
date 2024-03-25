package config

import (
	"fmt"
	"os"

	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))

	helpers.ErrorPanic(err)

	fmt.Print("Database connected!")

	database.AutoMigrate()

	DB = database
}