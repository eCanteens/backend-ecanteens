package seeders

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func TransactionSeeder() {
	var transactions []*models.Transaction

	for i := 0; i < 21; i++ {
		for j := 0; j < 2; j++ {
			transactions = append(transactions, &models.Transaction{
				TransactionId: fmt.Sprintf("EC-%d-%d", gofakeit.DateRange(time.Now(), time.Now().AddDate(0, 0, 1)).Unix(), i+1),
				UserId:        uint(i) + 1,
				Type:          transaction.TransactionType(gofakeit.RandomString([]string{"TOPUP", "WITHDRAW"})),
				Status:        transaction.TransactionStatus(gofakeit.RandomString([]string{"INPROGRESS", "SUCCESS", "CANCELED"})),
				Amount:        (gofakeit.UintRange(1_000, 20_000) / 100) * 100,
			})
		}
	}

	config.DB.Create(transactions)

	fmt.Println("Transaction Seeder created")
}
