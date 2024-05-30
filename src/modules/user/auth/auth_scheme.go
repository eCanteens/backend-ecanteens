package auth

import "mime/multipart"

type registerScheme struct {
	Name     string `binding:"required" mod:"trim" json:"name"`
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Phone    string `binding:"required,numeric,min=11,max=13" mod:"trim" json:"phone"`
	Password string `binding:"required,min=8" mod:"trim" json:"password"`
}

type loginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type googleScheme struct {
	IdToken string `binding:"required" mod:"trim" json:"id_token"`
}

type setupScheme struct {
	Phone    string `binding:"required,numeric,min=11,max=13" mod:"trim" json:"phone"`
	Password string `binding:"required,min=8" mod:"trim" json:"password"`
}

type refreshScheme struct {
	RefreshToken string `binding:"required" json:"refresh_token"`
}

type forgotScheme struct {
	Email string `binding:"required,email" mod:"trim" json:"email"`
}

type resetScheme struct {
	Token    string `binding:"required" mod:"trim" json:"token"`
	Password string `binding:"required,min=8" mod:"trim" json:"password"`
}

type updateScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Email  string                `binding:"required,email" mod:"trim" json:"email" form:"email"`
	Phone  string                `binding:"required,numeric,min=11,max=13" mod:"trim" json:"phone" form:"phone"`
}

type updatePasswordScheme struct {
	OldPassword string `binding:"required" mod:"trim" json:"old_password"`
	NewPassword string `binding:"required,min=8" mod:"trim" json:"new_password"`
}

type checkPinScheme struct {
	Pin string `binding:"required" mod:"trim" json:"pin"`
}

type updatePinScheme struct {
	Pin string `binding:"required,numeric,len=6" mod:"trim" json:"pin"`
}
