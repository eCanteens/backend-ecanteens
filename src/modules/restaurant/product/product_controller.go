package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func handleCreateProduct(ctx *gin.Context) {
	var body createProduct

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := createProductService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{"message": "Menu berhasil ditambahkan"})
}

func handleGetAllProduct(ctx *gin.Context) {
	var query productQs

	ctx.ShouldBindQuery(&query)

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := getAllProductService(&query, &_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func handleUpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateProduct

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProductService(&_user, &body, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Menu berhasil diupdate"})
}

func handleDeleteProduct(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	id := ctx.Param("id")

	if err := deleteProductService(&_user, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Menu berhasil dihapus"})
}
