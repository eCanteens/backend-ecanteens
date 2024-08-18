package main

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/eCanteens/backend-ecanteens/src/middleware"
	"github.com/eCanteens/backend-ecanteens/src/modules/admin"
	restoAuth "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/auth"
	restoProduct "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/product"
	restoTransaction "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/transaction"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/auth"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/product"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/restaurant"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/transaction"
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

	user := router.Group("/api/user")
	resto := router.Group("/api/restaurant")
	admin.Routes(router.Group("/api/admin"))

	// User routes
	{
		auth.Routes(user.Group("/auth"))
		restaurant.Routes(user.Group("/restaurants"))
		product.Routes(user.Group("/products"))
		transaction.Routes(user.Group("/transactions"))
	}
	// Resto routes
	{
		restoAuth.Routes(resto.Group("/auth"))
		restoProduct.Routes(resto.Group("/products"))
		restoTransaction.Routes(resto.Group("/transactions"))
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "API Docs on /api"})
	})

	router.Run()
}
