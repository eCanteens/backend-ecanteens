package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.POST("/register/check", handleCheckRegister)
	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)
	router.POST("/refresh", handleRefresh)

	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.GET("/profile", handleProfile)
		authorized.PUT("/profile", handleUpdateProfile)
		authorized.PUT("/restaurant", handleUpdateResto)
		authorized.PUT("/password", handleUpdatePassword)
	}
}