package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func getCart(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := getCartService(&_user)

	if err != nil {
		ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func updateCart(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateCartNoteScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateCartService(id, &body, *_user.Id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Catatan berhasil ditambahkan"))
}

func addCart(ctx *gin.Context) {
	var body addCartScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := addCartService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	if *body.Quantity == 0 {
		ctx.JSON(201, helpers.SuccessResponse("Produk berhasil dihapus dari keranjang"))
	} else {
		ctx.JSON(201, helpers.SuccessResponse("Produk berhasil ditambahkan ke keranjang"))
	}
}

func getOrder(ctx *gin.Context) {
	var query getOrderQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := getOrderService(*_user.Id, &query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

func handleOrder(ctx *gin.Context) {
	var body orderScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := orderService(&body, &_user)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Pesanan berhasil dibuat", helpers.Data{
		"data": data,
	}))
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

	if err := updateOrderService(&body, id, *_user.Id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	msg := "Pesanan berhasil"

	switch body.Status {
	case "SUCCESS":
		msg += " dikonfirmasi"
	case "CANCELED":
		msg += " dibatalkan"
	}

	ctx.JSON(200, helpers.SuccessResponse(msg))
}
