package product

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/favorites", getFavorite)

		authorized.POST("/:id/feedback", addFeedback)
		authorized.DELETE("/:id/feedback", removeFeedback)

		authorized.POST("/:id/favorite", addFavorite)
		authorized.DELETE("/:id/favorite", removeFavorite)
	}
}