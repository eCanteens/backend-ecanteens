package routes

import (
	"github.com/eCanteens/backend-ecanteens/src/modules/admin"
	restoAuth "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/auth"
	restoProduct "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/product"
	restoTransaction "github.com/eCanteens/backend-ecanteens/src/modules/restaurant/transaction"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/auth"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/product"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/restaurant"
	"github.com/eCanteens/backend-ecanteens/src/modules/user/transaction"
	"github.com/gin-gonic/gin"
)

func Apiv1(router *gin.RouterGroup) {
	user := router.Group("/user")
	resto := router.Group("/restaurant")
	admin.Routes(router.Group("/admin"))

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
}
