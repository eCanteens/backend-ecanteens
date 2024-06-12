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

func handleUpdateOrder(ctx *gin.Context) {
	var body updateOrderScheme
	id := ctx.Param("id")

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err:= updateOrderService(id, &_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	msg := "Pesanan berhasil"

	switch body.Status {
	case "INPROGRESS":
		msg += " diterima"
	case "CANCELED":
		msg += " ditolak"
	case "READY":
		msg = "Status pesanan berhasil diperbarui"
	}

	ctx.JSON(200, helpers.SuccessResponse(msg))
}