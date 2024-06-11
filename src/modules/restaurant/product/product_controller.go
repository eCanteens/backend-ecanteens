package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func handleCreateProduct(ctx *gin.Context) {
	var body createProduct

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := createProductService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Menu berhasil ditambahkan"))
}

func handleGetAllProduct(ctx *gin.Context) {
	var query productQs

	ctx.ShouldBindQuery(&query)

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := getAllProductService(&query, &_user)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

func handleUpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateProduct

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProductService(&_user, &body, id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Menu berhasil diupdate"))
}

func handleDeleteProduct(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	id := ctx.Param("id")

	if err := deleteProductService(&_user, id); err != nil {
		ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Menu berhasil dihapus"))
}
