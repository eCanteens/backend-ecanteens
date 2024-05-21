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
		router.GET("/dashboard", handleDashoard)

		// topup
		router.POST("/wallet", handleCheckWallet)
		router.POST("/topup", handleTopup)

		// withdraw
		router.POST("/withdraw", handleWithdraw)

		// mutasi
		router.GET("/mutasi", handleMutasi)

		// profile
		authorized.GET("/profile", handleProfile)
		authorized.PUT("/profile", handleUpdateProfile)
		authorized.PUT("/password", handleUpdatePassword)
	}
}