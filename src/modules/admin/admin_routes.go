package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	adminRepo       = NewRepository()
	adminService    = NewService(adminRepo)
	adminController = NewController(adminService)
)

func Routes(router *gin.RouterGroup) {
	// auth
	router.POST("/login", adminController.adminLogin)

	authorized := router.Group("")
	authorized.Use(middleware.Admin)

	{
		// dashboard
		authorized.GET("/dashboard", adminController.dashoard)

		// check wallet
		authorized.GET("/wallet/:phone", adminController.checkWallet)

		// topup
		authorized.POST("/topup/:phone", adminController.topup)

		// withdraw
		authorized.POST("/withdraw/:phone", adminController.withdraw)

		// transaction
		authorized.GET("/transaction/:id", adminController.transaction)

		// mutasi
		authorized.GET("/mutasi", adminController.mutasi)

		// profile
		authorized.GET("/profile", adminController.adminProfile)
		authorized.PUT("/profile", adminController.updateAdminProfile)
		authorized.PUT("/password", adminController.updateAdminPassword)
	}
}
