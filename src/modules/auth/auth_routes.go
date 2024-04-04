package auth

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/forgot", Forgot)
}