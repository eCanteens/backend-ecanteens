package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func checkEmailAndPhone(email string, phone string, id ...uint) *[]models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", email).Or("phone = ?", phone),
	)

	if len(id) > 0 {
		query = query.Not("id = ?", id[0])
	}

	query.Find(&sameUser)

	return &sameUser
}

func findByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 3).Preload("Wallet").First(user).Error
}

func findById(user *models.User, id uint) error {
	return config.DB.Where("id = ?", id).Where("role_id = ?", 3).First(user).Error
}

func create(user *models.User) error {
	return config.DB.Create(user).Error
}