package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	transactionRepo       = NewRepository()
	transactionService    = NewService(transactionRepo)
	transactionController = NewController(transactionService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/carts", transactionController.getCart)
		authorized.PUT("/carts/:id", transactionController.updateCart)
		authorized.POST("/carts", transactionController.addCart)
		authorized.GET("/orders", transactionController.getOrder)
		authorized.POST("/orders", transactionController.createOrder)
		authorized.POST("/orders/:id/review", transactionController.postReview)
		authorized.PUT("/orders/:id", transactionController.updateOrder)
	}
}