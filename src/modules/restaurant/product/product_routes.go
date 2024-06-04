package product

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.POST("/", handleCreateProduct)
		authorized.GET("/", handleGetAllProduct)
		authorized.GET("/:id", handleGetOneProduct)
		authorized.PUT("/:id", handleUpdateProduct)
		authorized.DELETE("/:id", handleDeleteProduct)
	}
}
