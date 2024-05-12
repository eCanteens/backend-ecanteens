package auth

import "mime/multipart"

type RegisterScheme struct {
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

type LoginScheme struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required" json:"password"`
}

type ForgotScheme struct {
	Email string `binding:"required,email" json:"email"`
}

type ResetScheme struct {
	Password string `binding:"required,min=8" json:"password"`
}

type UpdateScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" json:"name" form:"name"`
	Email  string                `binding:"required,email" json:"email" form:"email"`
	Phone  *string               `json:"phone" form:"phone"`
}

type UpdatePasswordScheme struct {
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required,min=8" json:"new_password"`
}

type CheckPinScheme struct {
	Pin string `binding:"required" json:"pin"`
}

type UpdatePinScheme struct {
	Pin string `binding:"required,numeric,len=6" json:"pin"`
}
