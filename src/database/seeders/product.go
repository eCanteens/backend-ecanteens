package seeders

import (
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func productCategorySeeder() {
	productCategory := []*models.ProductCategory{
		{ Name: "Makanan Berat" },
		{ Name: "Minuman" },
		{ Name: "Jajanan" },
		{ Name: "Makanan Pendamping" },
		{ Name: "Lainnya" },
	}

	config.DB.Create(&productCategory)
}

func ProductSeeder() {
	productCategorySeeder()

	var products []*models.Product

	for i := 0; i < 50; i++ {
		randomInt := gofakeit.IntRange(0, 5)
		productName := []string{gofakeit.Breakfast(), gofakeit.Lunch(), gofakeit.Drink(), gofakeit.Snack(), gofakeit.Fruit(), gofakeit.Dessert()}
		var category uint = 1

		if randomInt == 2 {
			category = 2
		} else if randomInt == 3 {
			category = 3
		} else if randomInt == 4 {
			category = 4
		} else {
			category = 5
		}

		products = append(products, &models.Product{
			RestaurantId: gofakeit.UintRange(1, 6),
			Name:         productName[randomInt],
			Description:  gofakeit.ProductDescription(),
			Image:        os.Getenv("BASE_URL") + "/public/dummy/product.png",
			CategoryId:   category,
			Price:        (gofakeit.UintRange(1_000, 10_000) / 100) * 100,
			Stock:        gofakeit.UintRange(0, 20),
			Sold:         gofakeit.UintRange(0, 200),
		})
	}

	config.DB.Create(products)

	fmt.Println("Product Seeder created")
}
