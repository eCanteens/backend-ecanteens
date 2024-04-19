package seeders

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func productCategorySeeder() {
	productCategory := models.ProductCategory{
		Name: "Jajanan",
	}

	config.DB.Create(&productCategory)
}

func ProductSeeder() {
	productCategorySeeder()

	var products []*models.Product

	for i := 0; i < 50; i++ {
		products = append(products, &models.Product{
			RestaurantId: gofakeit.UintRange(1, 10),
			Name: gofakeit.ProductName(),
			Description: gofakeit.ProductDescription(),
			Image: "/public/uploads/dummy/product.png",
			CategoryId: 1,
			Price: (gofakeit.UintRange(1_000, 20_000) / 100) * 100,
			Stock: gofakeit.UintRange(0, 20),
			Sold: gofakeit.UintRange(0, 200),
			Like: gofakeit.UintRange(0, 200),
			Dislike: gofakeit.UintRange(0, 200),
		})
	}

	config.DB.Create(products)

	fmt.Println("Product Seeder created")
}