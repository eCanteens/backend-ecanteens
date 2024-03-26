package example

import "github.com/gin-gonic/gin"

func Route(router *gin.RouterGroup) {
	router.GET("/", Test)
	router.POST("/", TestPOST)
}