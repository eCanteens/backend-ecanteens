package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	authRepo       = NewRepository()
	authService    = NewService(authRepo)
	authController = NewController(authService)
)

func Routes(router *gin.RouterGroup) {
	router.POST("/register", authController.register)
	router.POST("/login", authController.login)
	router.POST("/logout", authController.logout)
	router.POST("/google", authController.google)
	router.POST("/refresh", authController.refresh)
	router.POST("/forgot-password", authController.forgot)
	router.PUT("/new-password", authController.reset)
	
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/google/setup", authController.setup)
		authorized.GET("/profile", authController.profile)
		authorized.PUT("/profile", authController.updateProfile)
		authorized.PUT("/password", authController.updatePassword)
		authorized.POST("/check-pin", authController.checkPin)
		authorized.PUT("/pin", authController.updatePin)
	}
}
