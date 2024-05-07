package auth

type RegisterSchema struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

type LoginSchema struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type ForgotSchema struct {
	Email string `binding:"required,email"`
}

type ResetSchema struct {
	Password string `binding:"required,min=8"`
}

type UpdateSchema struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Phone string `binding:"min=10,max=14,numeric"`
}
