package dashboard

import (
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	dashboardRepo       = NewRepository()
	dashboardService    = NewService(dashboardRepo)
	dashboardController = NewController(dashboardService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Resto)
	{
		authorized.GET("/", dashboardController.dashboard)
		authorized.GET("/open", dashboardController.toggleOpen)
	}
}