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
	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.GET("/orders", transactionController.getOrder)
		authorized.PUT("/orders/:id", transactionController.updateOrder)
	}
}