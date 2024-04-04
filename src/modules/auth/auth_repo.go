package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func RegisterRepo(user *models.User) error {
	return config.DB.Create(user).Error
}

func CheckEmailAndPhone(user *models.User) []models.User {
	var sameUser []models.User

	config.DB.Where(
		config.DB.Where("email = ?", user.Email),
	).Or(
		config.DB.Where("phone = ?", user.Phone),
	).Find(&sameUser)

	return sameUser
}

func FindByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).First(user).Error
}