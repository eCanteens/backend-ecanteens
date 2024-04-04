package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/forgot-password", Forgot)
	router.PATCH("/reset-password/:token", Reset)

	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/profile", Profile)
	}
}
