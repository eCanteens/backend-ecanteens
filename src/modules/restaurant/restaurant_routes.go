package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/favorites", GetFavoriteResto)
		authorized.GET("/", GetAllResto)
		authorized.GET("/:id/products", GetRestosProducts)
		authorized.GET("/:id", GetDetailResto)
	}
}