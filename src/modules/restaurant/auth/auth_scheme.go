package auth

import "mime/multipart"

type checkRegisterScheme struct {
	Email string `binding:"required,email" json:"email"`
	Phone string `binding:"required,numeric,min=11,max=13" json:"phone"`
}

type registerScheme struct {
	Avatar           *multipart.FileHeader `binding:"required" form:"avatar"`
	Name             string                `binding:"required" form:"name" json:"name" mod:"trim"`
	Phone            string                `binding:"required,numeric,min=11,max=13" form:"phone" json:"phone" mod:"trim"`
	Email            string                `binding:"required,email" form:"email" json:"email" mod:"trim"`
	Password         string                `binding:"required,min=8" form:"password" json:"password" mod:"trim"`
	RestaurantAvatar *multipart.FileHeader `binding:"required" form:"restaurant_avatar" json:"restaurant_avatar" mod:"trim"`
	RestaurantName   string                `binding:"required" form:"restaurant_name" json:"restaurant_name" mod:"trim"`
	CategoryId       uint                  `binding:"required,numeric" form:"category_id" json:"category_id" mod:"trim"`
	Banner           *multipart.FileHeader `binding:"required" form:"banner"`
}

type loginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type refreshScheme struct {
	RefreshToken string `binding:"required" json:"refresh_token"`
}

type updateProfileScheme struct {
	Avatar *multipart.FileHeader `form:"avatar"`
	Name   string                `binding:"required" json:"name" form:"name"`
	Phone  string                `binding:"required,numeric,min=11,max=13" json:"phone" form:"phone"`
	Email  string                `binding:"required,email" json:"email" form:"email"`
}
