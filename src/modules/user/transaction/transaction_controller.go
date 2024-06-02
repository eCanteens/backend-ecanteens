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

	if err := updateCartService(id, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return	
	}

	ctx.JSON(201, helpers.SuccessResponse("Catatan berhasil ditambahkan"))
}

func addCart(ctx *gin.Context) {
	var body addCartScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := addCartService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		return
	}

	if *body.Quantity == 0 {
		ctx.JSON(201, helpers.SuccessResponse("Produk berhasil dihapus dari keranjang"))
	} else {
		ctx.JSON(201, helpers.SuccessResponse("Produk berhasil ditambahkan ke keranjang"))
	}
}

func order(ctx *gin.Context) {

}
