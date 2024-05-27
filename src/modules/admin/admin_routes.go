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

		// check wallet
		authorized.POST("/wallet", handleCheckWallet)
		authorized.GET("/wallet/:id", handleGetWallet)

		// topup
		authorized.POST("/topup/:id", handleTopup)

		// withdraw
		authorized.POST("/withdraw/:id", handleWithdraw)

		// mutasi
		authorized.GET("/mutasi", handleMutasi)

		// profile
		authorized.GET("/profile", handleAdminProfile)
		authorized.PUT("/profile", handleUpdateAdminProfile)
		authorized.PUT("/password", handleUpdateAdminPassword)
	}
}