package auth

type LoginSchema struct {
	Email    string `binding:"required"`
	Password string	`binding:"required"`
}