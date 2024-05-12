package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func create(user *models.User) error {
	return config.DB.Create(user).Error
}

func checkEmailAndPhone(user *models.User, id ...*uint) []models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", user.Email).Or("phone = ?", user.Phone),
	)

	if len(id) > 0 {
		query = query.Not("id = ?", *id[0])
	}

	query.Find(&sameUser)

	return sameUser
}

func findByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Preload("Wallet").First(user).Error
}

func save(user *models.User) error {
	return config.DB.Save(user).Error
}

func updatePassword(id uint, user *models.User) error {
	return config.DB.Model("password").Where("id = ?", id).Updates(user).Error
}