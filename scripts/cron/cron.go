package main

import (
	"log"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	log.Println("Start cleaning up...")
	for {
		now := time.Now()

		// Perbarui semua token yang kedaluwarsa
		config.DB.Unscoped().
			Where("last_used < ?", now.Add(-config.App.Auth.AccessTokenExpiresIn)).
			Delete(&models.Token{})

		log.Println("Cleanup: expired tokens deactivated")
		time.Sleep(24 * time.Hour) // Jalankan setiap 24 jam
	}
}
