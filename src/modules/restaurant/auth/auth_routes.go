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
	router.POST("/register/check", authController.checkRegister)
	router.POST("/register", authController.register)
	router.POST("/login", authController.login)
	router.POST("/logout", authController.logout)
	router.POST("/refresh", authController.refresh)

	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.GET("/profile", authController.profile)
		authorized.PUT("/profile", authController.updateProfile)
		authorized.PUT("/restaurant", authController.updateResto)
		authorized.PUT("/password", authController.updatePassword)
	}
}