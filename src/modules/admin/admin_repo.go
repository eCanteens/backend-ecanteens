package admin

import (
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type Repository interface {
	findAdminEmail(user *models.User, email string) error
	count(table string, count *int64) error
	checkEmail(email string, id ...uint) *[]models.User
	save(user *models.User) error
	findUser(user *models.User, phone string) error
	topupWithdraw(amount uint, user *models.User, tipe string) error
	createTransaction(user *models.User, amount uint, tipe enums.TransactionType) (*models.Transaction, error)
	findTransaction(transaction *models.Transaction, id string) (*models.Transaction, error)
	findMutasi(result *pagination.Pagination[models.Transaction], query *mutationQS) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) findAdminEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 1).First(user).Error
}

func (r *repository) count(table string, count *int64) error {
	return config.DB.Table(table).Count(count).Error
}

func (r *repository) checkEmail(email string, id ...uint) *[]models.User {
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

func (r *repository) save(user *models.User) error {
	return config.DB.Save(user).Error
}

func (r *repository) findUser(user *models.User, phone string) error {
	return config.DB.Where("phone = ?", phone).Preload("Wallet").First(user).Error
}

func (r *repository) topupWithdraw(amount uint, user *models.User, tipe string) error {
	if tipe == "TOPUP" {
		user.Wallet.Balance += amount
	} else if tipe == "WITHDRAW" {
		if amount > user.Wallet.Balance {
			return customerror.New("Saldo tidak mencukupi", 400)
		}
		user.Wallet.Balance -= amount
	}

	return config.DB.Save(user.Wallet).Error
}

func (r *repository) createTransaction(user *models.User, amount uint, tipe enums.TransactionType) (*models.Transaction, error) {
	transaction := models.Transaction{
		TransactionCode: helpers.GenerateTrxCode(*user.Id),
		UserId:          *user.Id,
		Type:            tipe,
		Status:          enums.TrxStatusSuccess,
		Amount:          amount,
	}
	return &transaction, config.DB.Create(&transaction).Error
}

func (r *repository) findTransaction(transaction *models.Transaction, id string) (*models.Transaction, error) {
	return transaction, config.DB.Where("transaction_id = ?", id).Preload("User.Wallet").First(transaction).Error
}

func (r *repository) findMutasi(result *pagination.Pagination[models.Transaction], query *mutationQS) error {
	search := query.Search
	page := query.Page
	order := query.Order
	direction := query.Direction

	db := config.DB.Model(&models.Transaction{}).
		Joins("LEFT JOIN users ON users.id = transactions.user_id").
		Preload("User.Wallet")

	if search != "" {
		searchPattern := "%" + search + "%"
		db.Where(
			config.DB.Where("users.name ILIKE ?", searchPattern).
				Or("users.email ILIKE ?", searchPattern).
				Or("CAST(transactions.amount AS TEXT) ILIKE ?", searchPattern).
				Or("CAST(transactions.created_at AS TEXT) ILIKE ?", searchPattern),
		)
	}

	if query.Type == "" {
		db.Where(
			config.DB.Where("transactions.type = ?", "TOPUP").Or("transactions.type = ?", "WITHDRAW"),
		)
	} else {
		typeFilter := strings.ToUpper(query.Type)
		db.Where("transactions.type = ?", typeFilter)
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
