package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.GET("/orders", handleGetOrder)
	}
}