package auth

type LoginSchema struct {
	Email    string `binding:"required"`
	Password string	`binding:"required"`
}

type ForgotSchema struct {
	Email    string `binding:"required"`
}