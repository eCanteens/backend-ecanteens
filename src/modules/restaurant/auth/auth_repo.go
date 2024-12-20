package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"gorm.io/gorm"
)

type Repository interface {
	createResto(data *models.Restaurant) error
	updateResto(data *models.Restaurant) error

	findByEmail(user *models.User, email string) error
	checkEmailAndPhone(email string, phone string, id ...uint) *[]models.User
	updateUser(data *models.User) error

	findToken(model *models.Token, token string) error
	createToken(data *models.Token) error
	updateToken(data *models.Token) error
	deleteToken(token string) error
	deleteTokenById(data *models.Token) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) createResto(data *models.Restaurant) error {
	return config.DB.Create(data).Error
}

func (r *repository) updateResto(data *models.Restaurant) error {
	return config.DB.Save(data).Error
}

func (r *repository) findByEmail(user *models.User, email string) error {
	return config.DB.Where("email = ?", email).Where("role_id = ?", 3).Preload("Wallet").Preload("Restaurant").First(user).Error
}

func (r *repository) checkEmailAndPhone(email string, phone string, id ...uint) *[]models.User {
	var sameUser []models.User

	query := config.DB.Where(
		config.DB.Where("email = ?", email).Or("phone = ?", phone),
	)

	if len(id) > 0 {
		query = query.Not("id = ?", id[0])
	}

	query.Limit(2).Find(&sameUser)

	return &sameUser
}

func (r *repository) updateUser(data *models.User) error {
	return config.DB.Save(data).Error
}

func (r *repository) findToken(model *models.Token, token string) error {
	return config.DB.Where("token = ?", token).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id = ?", 3)
	}).First(model).Error
}

func (r *repository) createToken(data *models.Token) error {
	return config.DB.Create(data).Error
}

func (r *repository) updateToken(data *models.Token) error {
	return config.DB.Save(data).Error
}

func (r *repository) deleteToken(token string) error {
	if affected := config.DB.Unscoped().Where("token = ?", token).Delete(&models.Token{}).RowsAffected; affected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) deleteTokenById(data *models.Token) error {
	return config.DB.Unscoped().Delete(data).Error
}
