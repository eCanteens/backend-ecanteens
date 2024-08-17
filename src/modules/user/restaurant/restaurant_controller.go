package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/gin-gonic/gin"
)

func getFavorite(ctx *gin.Context) {
	var query paginationQS
	user, _ := ctx.Get("user")

	ctx.ShouldBindQuery(&query)

	data, err := getFavoriteService(*user.(models.User).Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func getAll(ctx *gin.Context) {
	var query paginationQS

	ctx.ShouldBindQuery(&query)

	data, err := getAllService(&query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func getReviews(ctx *gin.Context) {
	id := ctx.Param("id")
	var query reviewQS

	ctx.ShouldBindQuery(&query)

	data, err := getReviewsService(id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func getDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := getDetailService(id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func getRestosProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	var query paginationQS

	ctx.ShouldBindQuery(&query)

	data, err := getRestosProductsService(id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := addFavoriteService(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil ditambahkan ke favorit"})
}

func removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFavoriteService(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil dihapus dari favorit"})
}
