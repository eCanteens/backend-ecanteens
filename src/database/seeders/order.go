package seeders

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/constants/order"
	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
)

func OrderSeeder() {
	var users []models.User
	var orders []*models.Order

	config.DB.Select("id").Where("role_id", 2).Find(&users)

	// Loop user
	for _, user := range users {
		// Loop Order
		for i := 0; i < 20; i++ {
			restaurantId := gofakeit.RandomUint([]uint{1, 2, 3, 4, 5})
			var ordItems []models.OrderItem
			var products []models.Product
			var amount uint

			config.DB.Where("restaurant_id = ?", restaurantId).Find(&products)

			productIds := helpers.Map(&products, func(p *models.Product) *uint {
				return p.Id
			})

			// Loop Order Item
			for j := 0; j < 3; j++ {
				productId := gofakeit.RandomUint(*productIds)
				quantity := gofakeit.UintRange(1, 5)
				product := helpers.Find(&products, func(t *models.Product) bool {
					return *t.Id == productId
				})

				ordItems = append(ordItems, models.OrderItem{
					ProductId: productId,
					Quantity:  quantity,
					Price:     product.Price,
				})

				amount += product.Price * quantity
			}

			ord := models.Order{
				UserId:       *user.Id,
				RestaurantId: restaurantId,
				Notes:        gofakeit.Comment(),
				Status:       order.OrderStatus(gofakeit.RandomString([]string{"WAITING", "INPROGRESS", "READY", "SUCCESS", "CANCELED"})),
				IsPreorder:   gofakeit.Bool(),
				Amount:       amount,
				Items:        ordItems,
				Transaction: &models.Transaction{
					TransactionCode: fmt.Sprintf("EC-%d-%d", gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()).Unix(), 3),
					UserId:          *user.Id,
					Type:            transaction.PAY,
					Amount:          amount,
					PaymentMethod:   transaction.TransactionPaymentMethod(gofakeit.RandomString([]string{"CASH", "ECANTEENSPAY"})),
				},
			}

			// Set fullfilment date
			if ord.IsPreorder && ord.Status == order.WAITING {
				ord.FullfilmentDate = helpers.PointerTo(gofakeit.DateRange(time.Now(), time.Now().AddDate(0, 0, 5)))
			} else if ord.IsPreorder {
				ord.FullfilmentDate = helpers.PointerTo(gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()))
			}

			// Set transaction status
			if (ord.Status == order.WAITING && ord.IsPreorder && ord.Transaction.PaymentMethod == transaction.ECANTEENSPAY) || ord.Status == order.SUCCESS {
				ord.Transaction.Status = transaction.SUCCESS
			} else if ord.Status == order.CANCELED {
				ord.Transaction.Status = transaction.CANCELED
			} else {
				ord.Transaction.Status = transaction.INPROGRESS
			}

			orders = append(orders, &ord)
		}
	}

	config.DB.Create(orders)

	fmt.Println("Order Seeder created")
}
