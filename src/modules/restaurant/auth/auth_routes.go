package auth

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/register/check", handleCheckRegister)
	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)
	router.POST("/refresh", handleRefresh)
}