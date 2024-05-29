package admin

import (
	"errors"
	"fmt"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func findAdminEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 1).Preload("Wallet").First(user).Error
}

func count(table string, count *int64) error {
	return config.DB.Table(table).Count(count).Error
}

func checkEmail(email string, id ...uint) *[]models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", email),
	)

	if len(id) > 0 {
		query = query.Not("id = ?", id[0])
	}

	query.Find(&sameUser)

	return &sameUser
}

func save(user *models.User) error {
	return config.DB.Save(user).Error
}

func findUser(user *models.User, phone string) error {
	return config.DB.Where("phone = ?", phone).Preload("Wallet").First(user).Error
}

func topupWithdraw(amount uint, user *models.User, tipe string) error {
	if tipe == "TOPUP" {
		user.Wallet.Balance += amount
	} else if tipe == "WITHDRAW" {
		if amount > user.Wallet.Balance {
			return errors.New("saldo tidak mencukupi")
		}
		user.Wallet.Balance -= amount
	}

	return config.DB.Save(user.Wallet).Error
}

func createTransaction(user *models.User, amount uint, tipe models.TransactionType) (*models.Transaction, error) {
	transaction := models.Transaction{
		TransactionId: fmt.Sprintf("EC-%d-%d", time.Now().Unix(), *user.Id.Id),
		UserId:        *user.Id.Id,
		Type:          tipe,
		Status:        models.SUCCESS,
		Amount:        amount,
		Items:         "[]",
	}
	return &transaction, config.DB.Create(&transaction).Error
}

func findTransaction(transaction *models.Transaction, id string) (*models.Transaction, error) {
	return transaction, config.DB.Where("transaction_id = ?", id).Preload("User.Wallet").First(transaction).Error
}
