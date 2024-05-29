package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func getFavorite(ctx *gin.Context) {
	query := map[string]string{}
	user, _ := ctx.Get("user")

	ctx.ShouldBindQuery(query)

	data, err := getFavoriteService(*user.(models.User).Id.Id, query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func getAll(ctx *gin.Context) {
	query := map[string]string{}

	ctx.ShouldBindQuery(query)

	data, err := getAllService(query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

func getDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := getDetailService(id)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func getRestosProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	query := map[string]string{}

	ctx.ShouldBindQuery(query)

	data, err := getRestosProductsService(id, query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

func addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := addFavoriteService(*user.(models.User).Id.Id, id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Restoran berhasil ditambahkan ke favorit"))
}

func removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFavoriteService(*user.(models.User).Id.Id, id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Restoran berhasil dihapus dari favorit"))
}
