package product

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.POST("/:id/feedback", addFeedback)
		authorized.DELETE("/:id/feedback", removeFeedback)
	}
}