package auth

type RegisterScheme struct {
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

type LoginScheme struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type ForgotScheme struct {
	Email string `binding:"required,email"`
}

type ResetScheme struct {
	Password string `binding:"required,min=8"`
}

type UpdateScheme struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Phone string `binding:"min=10,max=14,numeric"`
}
