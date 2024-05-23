package admin

import (
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

func findWallet(wallet *models.Wallet, walletId string) (*models.Wallet, error) {
	return wallet, config.DB.Where("uuid = ?", walletId).First(wallet).Error
}

func findUserByWalletId(user *models.User, walletId ...*uint) (*models.User, error) {
	return user, config.DB.Where("wallet_id = ?", walletId).Preload("Wallet").First(user).Error
}

func topup(amount uint, id string) (*models.Wallet, error) {
	var wallet models.Wallet

	if err := config.DB.Model(&models.Wallet{}).Where("uuid = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}

	wallet.Balance += amount
	return &wallet, config.DB.Save(&wallet).Error
}