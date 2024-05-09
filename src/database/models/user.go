package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id
	Name     string  `gorm:"type:varchar(30);not null" json:"name" binding:"required"`
	Email    string  `gorm:"type:varchar(50);not null;unique" json:"email" binding:"required,email"`
	Phone    *string `gorm:"type:varchar(20);unique" json:"phone"`
	Password string  `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required,min=8"`
	Avatar   *string `gorm:"type:varchar(255)" json:"avatar"`
	Balance  uint    `gorm:"type:int; not null" json:"balance"`
	Timestamps

	// Relations
	FavoriteRestaurant []Restaurant `gorm:"many2many:favorites;" json:"favorite_restaurant,omitempty"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return
}
