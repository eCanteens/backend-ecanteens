package auth

import "mime/multipart"

type RegisterScheme struct {
	Name     string `binding:"required" mod:"trim" json:"name"`
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required,min=8" mod:"trim" json:"password"`
}

type LoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type GoogleScheme struct {
	Name   string `binding:"required" json:"name"`
	Email  string `binding:"required,email" json:"email"`
	Avatar string `binding:"required" json:"avatar"`
}

type ForgotScheme struct {
	Email string `binding:"required,email" mod:"trim" json:"email"`
}

type ResetScheme struct {
	Token    string `binding:"required" mod:"trim" json:"token"`
	Password string `binding:"required,min=8" mod:"trim" json:"password"`
}

type UpdateScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Email  string                `binding:"required,email" mod:"trim" json:"email" form:"email"`
	Phone  string                `mod:"trim" json:"phone" form:"phone"`
}

type UpdatePasswordScheme struct {
	OldPassword string `binding:"required" mod:"trim" json:"old_password"`
	NewPassword string `binding:"required,min=8" mod:"trim" json:"new_password"`
}

type CheckPinScheme struct {
	Pin string `binding:"required" mod:"trim" json:"pin"`
}

type UpdatePinScheme struct {
	Pin string `binding:"required,numeric,len=6" mod:"trim" json:"pin"`
}
