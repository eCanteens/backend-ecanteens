package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/forgot-password", forgot)
	router.PUT("/reset-password/:token", reset)

	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/profile", profile)
		authorized.PUT("/profile", updateProfile)
	}
}
