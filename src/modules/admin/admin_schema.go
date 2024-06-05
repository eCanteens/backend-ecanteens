package admin

import (
	"mime/multipart"
)

type adminLoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type topupWithdrawScheme struct {
	Amount uint `binding:"required" mod:"trim" json:"amount"`
}

type updateAdminProfileScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Email  string                `binding:"required,email" mod:"trim" json:"email" form:"email"`
}

type updateAdminPasswordScheme struct {
	OldPassword string `binding:"required" mod:"trim" json:"old_password"`
	NewPassword string `binding:"required,min=8" mod:"trim" json:"new_password"`
}

type mutationQS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Order     string `form:"sort"`
	Direction string `form:"direction"`
	Type      string `form:"type"`
}
