package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id
	Name     string `gorm:"type:varchar(30)" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(50);unique" json:"email" binding:"required,email"`
	Phone    string `gorm:"type:varchar(20);unique" json:"phone"`
	Password string `gorm:"type:varchar(255)" json:"password,omitempty" binding:"required"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
	Balance  int    `gorm:"type:int" json:"balance"`
	Timestamps
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return
}
