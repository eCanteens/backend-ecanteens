package admin

import "mime/multipart"

type AdminLoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type CheckWalletScheme struct {
	WalletId uint `binding:"required" mod:"trim" json:"wallet_id"`
}

type UpdateAdminProfileScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Email  string                `binding:"required,email" mod:"trim" json:"email" form:"email"`
}