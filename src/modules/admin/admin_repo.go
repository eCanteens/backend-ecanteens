package admin

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func findAdminEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 1).Preload("Wallet").First(user).Error
}

func count(table string, count *int64) error {
	return config.DB.Table(table).Count(count).Error
}

func checkEmail(user *models.User, id ...*uint) *[]models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", user.Email),
	)

	if len(id) > 0 {
		query = query.Not("id = ?", *id[0])
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
	if tipe == "topup" {
		user.Wallet.Balance += amount
	} else if tipe == "withdraw" {
		if amount > user.Wallet.Balance {
			return errors.New("balance tidak mencukupi")
		}
		user.Wallet.Balance -= amount
	}

	return config.DB.Save(user.Wallet).Error
}