package admin

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type AdminLoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type CheckWalletScheme struct {
	WalletId uuid.UUID `binding:"required" mod:"trim" json:"wallet_id"`
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