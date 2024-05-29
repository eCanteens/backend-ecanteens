package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)
	router.POST("/google", handleGoogle)
	router.POST("/refresh", handleRefresh)
	router.POST("/forgot-password", handleForgot)
	router.PUT("/new-password", handleReset)
	
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/google/setup", handleSetup)
		authorized.GET("/profile", handleProfile)
		authorized.PUT("/profile", handleUpdateProfile)
		authorized.PUT("/password", handleUpdatePassword)
		authorized.POST("/check-pin", handleCheckPin)
		authorized.PUT("/pin", handleUpdatePin)
	}
}
