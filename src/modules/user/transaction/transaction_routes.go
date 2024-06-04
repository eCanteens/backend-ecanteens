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
		authorized.PUT("/carts/:id", updateCart)
		authorized.POST("/carts", addCart)
		authorized.GET("/order", getOrder)
		authorized.POST("/order", handleOrder)
		authorized.PUT("/order/cancel", handleCancelOrder)
	}
}