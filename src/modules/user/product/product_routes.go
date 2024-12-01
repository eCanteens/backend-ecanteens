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
	authorized := router.Group("")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/favorites", productController.getFavorite)

		authorized.GET("/:id/feedback", productController.checkFeedback)
		authorized.POST("/:id/feedback", productController.addFeedback)

		authorized.POST("/:id/favorite", productController.addFavorite)
		authorized.DELETE("/:id/favorite", productController.removeFavorite)
	}
}
