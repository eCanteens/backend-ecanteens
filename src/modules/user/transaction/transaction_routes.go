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
		authorized.GET("/orders", getOrder)
		authorized.POST("/orders", handleOrder)
		authorized.POST("/orders/:id/review", handlePostReview)
		authorized.PUT("/orders/:id", handleUpdateOrder)
	}
}