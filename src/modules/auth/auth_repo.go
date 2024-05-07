package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
)

func Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func CheckEmailAndPhone(user *models.User, id ...*uint) []models.User {
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

func FindByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).First(user).Error
}

func UpdatePassword(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Updates(user).Error
}

func Update(id *uint, user *models.User) error {
	return config.DB.Where("id = ?", id).Updates(user).Error
}
