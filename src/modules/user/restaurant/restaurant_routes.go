package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	restaurantRepo       = NewRepository()
	restaurantService    = NewService(restaurantRepo)
	restaurantController = NewController(restaurantService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/favorites", restaurantController.getFavorite)
		authorized.GET("", restaurantController.getAll)
		authorized.GET("/:id/reviews", restaurantController.getReviews)
		authorized.GET("/:id/products", restaurantController.getRestosProducts)
		authorized.GET("/:id", restaurantController.getDetail)
		
		authorized.POST("/:id/favorite", restaurantController.addFavorite)
		authorized.DELETE("/:id/favorite", restaurantController.removeFavorite)

	}
}