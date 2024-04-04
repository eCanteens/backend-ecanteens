package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/modules/auth"
	"github.com/eCanteens/backend-ecanteens/src/modules/example"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB(true)
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Docs on /api"})
	})

	router.GET("/api", func(c *gin.Context) {
		c.Redirect(301, "https://documenter.getpostman.com")
	})

	// routes
	example.Routes(router.Group("/api/example"))
	auth.Routes(router.Group("/api/auth"))

	router.Run()
}
