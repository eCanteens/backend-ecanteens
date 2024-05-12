package models

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"gorm.io/gorm"
)

type User struct {
	Id
	Name     string  `gorm:"type:varchar(50);not null" json:"name"`
	Email    string  `gorm:"type:varchar(50);not null;unique" json:"email"`
	Phone    *string `gorm:"type:varchar(20);unique" json:"phone"`
	Password string  `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Pin      string  `gorm:"type:varchar(255)" json:"pin,omitempty"`
	Avatar   *string `gorm:"type:varchar(255)" json:"avatar"`
	WalletId uint    `gorm:"type:bigint" json:"wallet_id"`
	Timestamps

	// Relations
	Wallet              *Wallet      `gorm:"foreignKey:wallet_id" json:"wallet,omitempty"`
	FavoriteRestaurants []Restaurant `gorm:"many2many:favorite_restaurants;" json:"favorite_restaurant,omitempty"`
	FavoriteProducts    []Product    `gorm:"many2many:favorite_products;" json:"favorite_products,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	wallet := Wallet{}
	config.DB.Create(&wallet)
	u.WalletId = *wallet.Id.Id
	return
}
