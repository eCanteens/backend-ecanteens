package seeders

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func LocationSeeder() {
	location := models.Location{
		Name: "SMK Negeri 69 Jakarta",
		Address: "Jl. Swadaya, RT.7/RW.7, Jatinegara, Kec. Cakung, Kota Jakarta Timur, Daerah Khusus Ibukota Jakarta 13930",
	}

	config.DB.Create(&location)

	fmt.Println("Location Seeder created")
}