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

func handleGetOneProduct(ctx *gin.Context) {}

func handleUpdateProduct(ctx *gin.Context) {}

func handleDeleteProduct(ctx *gin.Context) {}
