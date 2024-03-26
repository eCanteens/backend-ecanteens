package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(30)" json:"name" binding:"required"`
	Email     string    `gorm:"type:varchar(50);unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"type:varchar(255)" json:"password,omitempty" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return
}
