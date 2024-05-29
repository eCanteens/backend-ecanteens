package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/favorites", getFavorite)
		authorized.GET("/", getAll)
		authorized.GET("/:id/products", getRestosProducts)
		authorized.GET("/:id", getDetail)
		
		authorized.POST("/:id/favorite", addFavorite)
		authorized.DELETE("/:id/favorite", removeFavorite)
	}
}