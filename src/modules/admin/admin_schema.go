package admin

type LoginScheme struct {
	Email    string `binding:"required,email" mod:"trim" json:"email"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type CheckWalletScheme struct {
	WalletId uint `binding:"required" mod:"trim" json:"wallet_id"`
}