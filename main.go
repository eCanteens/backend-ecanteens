package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/eCanteens/backend-ecanteens/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	// helpers.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = config.App.MaxMultipartMemory
	router.Use(middleware.Cors)
	router.Use(middleware.RateLimiter)
	router.Use(middleware.Gzip)

	customValidator := validation.NewCustomValidator()
	binding.Validator = customValidator

	router.Static("/public", "./public")
	router.Static("/.well-known", "./.well-known")

	router.GET("/api", func(c *gin.Context) {
		c.Redirect(301, "https://documenter.getpostman.com/view/34881046/2sA3s9CTPU")
	})

	routes.Apiv1(router.Group("/api/v1"))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "API Docs on /api"})
	})

	router.Run()
}
