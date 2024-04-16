package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	DB = database
}