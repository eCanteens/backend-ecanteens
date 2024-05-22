package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	// auth
	router.POST("/login", handleAdminLogin)

	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		// dashboard
		authorized.GET("/dashboard", handleDashoard)

		// topup
		authorized.POST("/wallet", handleCheckWallet)
		authorized.POST("/topup", handleTopup)

		// withdraw
		authorized.POST("/withdraw", handleWithdraw)

		// mutasi
		authorized.GET("/mutasi", handleMutasi)

		// profile
		authorized.GET("/profile", handleAdminProfile)
		authorized.PUT("/profile", handleAdminUpdateProfile)
		authorized.PUT("/password", handleUpdatePassword)
	}
}