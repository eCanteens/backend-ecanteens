package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/eCanteens/backend-ecanteens/src/modules/admin"
	"github.com/eCanteens/backend-ecanteens/src/modules/auth"
	"github.com/eCanteens/backend-ecanteens/src/modules/product"
	"github.com/eCanteens/backend-ecanteens/src/modules/restaurant"
	"github.com/eCanteens/backend-ecanteens/src/modules/transaction"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	// config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	config.Upload(router)
	router.Use(cors.Default())
	router.Use(middleware.RateLimiter)

	customValidator := helpers.NewCustomValidator() 
	binding.Validator = customValidator

	router.Static("/public", "./public")
	router.Static("/.well-known", "./.well-known")

	router.GET("/api", func(c *gin.Context) {
		c.Redirect(301, "https://documenter.getpostman.com/view/34881046/2sA3JNb1JV")
	})

	// routes
	auth.Routes(router.Group("/api/auth"))
	restaurant.Routes(router.Group("/api/restaurants"))
	product.Routes(router.Group("/api/products"))
	transaction.Routes(router.Group("/api/transactions"))
	admin.Routes(router.Group("/api/admin"))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API Docs on /api"})
	})

	router.Run()
}
