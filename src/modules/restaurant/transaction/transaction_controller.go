package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func handleGetOrder(ctx *gin.Context) {
	var query getOrderQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := getOrderService(*_user.Restaurant.Id, &query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}