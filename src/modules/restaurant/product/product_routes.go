package product

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	productRepo       = NewRepository()
	productService    = NewService(productRepo)
	productController = NewController(productService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Resto)
	{
		authorized.POST("/", productController.createProduct)
		authorized.GET("/", productController.getAllProducts)
		authorized.PUT("/:id", productController.updateProduct)
		authorized.DELETE("/:id", productController.deleteProduct)
	}
}
