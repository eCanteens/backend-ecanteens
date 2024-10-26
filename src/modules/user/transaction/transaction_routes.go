package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	trxRepo       = NewRepository()
	trxService    = NewService(trxRepo)
	trxController = NewController(trxService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/cart", trxController.getRestaurantCart)
		authorized.GET("/carts", trxController.getCarts)
		authorized.PUT("/carts/:id", trxController.updateCart)
		authorized.POST("/carts", trxController.addCart)
		authorized.GET("/orders", trxController.getOrder)
		authorized.POST("/orders", trxController.createOrder)
		authorized.POST("/orders/:id/review", trxController.postReview)
		authorized.PUT("/orders/:id", trxController.updateOrder)
		authorized.GET("/history", trxController.getTrxHistory)
	}
}