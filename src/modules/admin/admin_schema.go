package admin

import (
	"mime/multipart"
)

type AdminLoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type CheckWalletScheme struct {
	WalletId string `binding:"required" mod:"trim" json:"wallet_id"`
}

type TopupWithdrawScheme struct {
	Amount   uint       `binding:"required" mod:"trim" json:"amount"`
}

type UpdateAdminProfileScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Email  string                `binding:"required,email" mod:"trim" json:"email" form:"email"`
}

type UpdateAdminPasswordScheme struct {
	OldPassword string `binding:"required" mod:"trim" json:"old_password"`
	NewPassword string `binding:"required,min=8" mod:"trim" json:"new_password"`
}