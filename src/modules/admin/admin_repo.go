package admin

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/constants/transaction"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

func findAdminEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 1).First(user).Error
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

func createTransaction(user *models.User, amount uint, tipe transaction.TransactionType) (*models.Transaction, error) {
	transaction := models.Transaction{
		TransactionCode: fmt.Sprintf("EC-%d-%d", time.Now().Unix(), *user.Id),
		UserId:          *user.Id,
		Type:            tipe,
		Status:          transaction.SUCCESS,
		Amount:          amount,
	}
	return &transaction, config.DB.Create(&transaction).Error
}

func findTransaction(transaction *models.Transaction, id string) (*models.Transaction, error) {
	return transaction, config.DB.Where("transaction_id = ?", id).Preload("User.Wallet").First(transaction).Error
}

func findMutasi(result *pagination.Pagination[models.Transaction], query *mutationQS) error {
	search := query.Search
	page := query.Page
	order := query.Order
	direction := query.Direction

	db := config.DB.Model(&models.Transaction{}).
		Joins("JOIN users ON users.id = transactions.user_id").
		Preload("User.Wallet")

	if search != "" {
		searchPattern := "%" + search + "%"
		db = db.Where(
			config.DB.Where("users.name ILIKE ?", searchPattern).
				Or("users.email ILIKE ?", searchPattern).
				Or("CAST(transactions.amount AS TEXT) ILIKE ?", searchPattern).
				Or("CAST(transactions.created_at AS TEXT) ILIKE ?", searchPattern),
		)
	}

	if query.Type != "" {
		typeFilter := strings.ToUpper(query.Type)
		db = db.Where("transactions.type = ?", typeFilter)
	}

	params := &pagination.Params{
		Query:     db,
		Page:      page,
		Limit:     "25",
		Order:     order,
		Direction: direction,
	}

	return result.Execute(params)
}
