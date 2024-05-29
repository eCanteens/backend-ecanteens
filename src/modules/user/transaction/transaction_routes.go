package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/carts", getCart)
		authorized.POST("/carts", addCart)
		authorized.POST("/order", order)
	}
}