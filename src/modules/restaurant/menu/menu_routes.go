package menu

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/menu", handleGetAllMenu)
		authorized.POST("/menu", handleCreateMenu)
		authorized.GET("/menu/:id", handleGetOneMenu)
		authorized.PUT("/menu/:id", handleUpdateMenu)
		authorized.DELETE("/menu/:id", handleDeleteMenu)
	}
}
