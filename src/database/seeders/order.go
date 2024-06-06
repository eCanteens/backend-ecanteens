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
	var orders []*models.Order

	for i := 0; i < 10; i++ {
		var ordItems []models.OrderItem
		var products []models.Product
		var amount uint

		config.DB.Where("restaurant_id = ?", i+1).Find(&products)

		productIds := helpers.Map(&products, func(p *models.Product) *uint {
			return p.Id
		})

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
			UserId:       3,
			RestaurantId: uint(i + 1),
			Notes:        gofakeit.Comment(),
			Status:       order.OrderStatus(gofakeit.RandomString([]string{"WAITING", "INPROGRESS", "READY", "SUCCESS", "CANCELED"})),
			IsPreorder:   gofakeit.Bool(),
			Amount:       amount,
			Items:        ordItems,
			Transaction: &models.Transaction{
				TransactionCode: fmt.Sprintf("EC-%d-%d", gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()).Unix(), 3),
				UserId:          3,
				Type:            transaction.PAY,
				Amount:          amount,
				PaymentMethod:   transaction.TransactionPaymentMethod(gofakeit.RandomString([]string{"CASH", "ECANTEENSPAY"})),
			},
		}

		if ord.IsPreorder && ord.Status == order.WAITING {
			ord.FullfilmentDate = helpers.PointerTo(gofakeit.DateRange(time.Now(), time.Now().AddDate(0, 0, 5)))
		} else if ord.IsPreorder {
			ord.FullfilmentDate = helpers.PointerTo(gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()))
		}

		if (ord.Status == order.WAITING && ord.IsPreorder && ord.Transaction.PaymentMethod == transaction.ECANTEENSPAY) || ord.Status == order.SUCCESS {
			ord.Transaction.Status = transaction.SUCCESS
		} else if ord.Status == order.CANCELED {
			ord.Transaction.Status = transaction.CANCELED
		} else {
			ord.Transaction.Status = transaction.INPROGRESS
		}

		orders = append(orders, &ord)
	}

	config.DB.Create(orders)

	fmt.Println("Order Seeder created")
}
