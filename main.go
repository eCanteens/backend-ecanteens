package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading env")
	}
}

func main() {
	router := gin.Default()
	config.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Docs on /api"})
	})

	router.GET("/api", func(c *gin.Context) {
		c.Redirect(301, "https://documenter.getpostman.com")
	})

	router.Use(cors.Default())
	router.Run()
}