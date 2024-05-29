package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	// auth
	router.POST("/login", handleAdminLogin)

	authorized := router.Group("/")
	authorized.Use(middleware.Admin)

	{
		// dashboard
		authorized.GET("/dashboard", handleDashoard)

		// check wallet
		authorized.GET("/wallet/:phone", handleCheckWallet)

		// topup
		authorized.POST("/topup/:phone", handleTopup)

		// withdraw
		authorized.POST("/withdraw/:phone", handleWithdraw)

		// transaction
		authorized.GET("/transaction/:id", handleTransaction)

		// mutasi
		authorized.GET("/mutasi", handleMutasi)

		// profile
		authorized.GET("/profile", handleAdminProfile)
		authorized.PUT("/profile", handleUpdateAdminProfile)
		authorized.PUT("/password", handleUpdateAdminPassword)
	}
}