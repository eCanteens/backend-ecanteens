package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/gorm"
)

func create[T any](data *T) error {
	return config.DB.Create(data).Error
}

func deleteById[T any](data *T) error {
	return config.DB.Unscoped().Delete(data).Error
}

func deleteToken(token string) error {
	if affected := config.DB.Unscoped().Where("token = ?", token).Delete(&models.Token{}).RowsAffected; affected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func checkEmailAndPhone(email string, phone string, id ...uint) *[]models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", email).Or("phone = ?", phone),
	)

	if len(id) > 0 {
		query.Not("id = ?", id[0])
	}

	query.Limit(2).Find(&sameUser)

	return &sameUser
}

func findByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 2).Preload("Wallet").First(user).Error
}

func findToken(model *models.Token, token string) error {
	return config.DB.Where("token = ?", token).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id = ?", 2)
	}).First(model).Error
}

func update[T any](model *T) error {
	return config.DB.Save(model).Error
}

func updatePassword(id uint, user *models.User) error {
	return config.DB.Select("password").Where("id = ?", id).Updates(user).Error
}
