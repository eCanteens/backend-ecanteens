package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func addFeedback(ctx *gin.Context) {
	var body feedbackScheme
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := addFeedbackService(&body, *user.(models.User).Id, id); err != nil {
		ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		return
	}

	msg := "Produk berhasil di"
	if *body.IsLike {
		msg += "like"
	} else {
		msg += "dislike"
	}

	ctx.JSON(200, helpers.SuccessResponse(msg))
}

func removeFeedback(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFeedbackService(*user.(models.User).Id, id); err != nil {
		ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Produk berhasil diunlike/undislike"))
}

func getFavorite(ctx *gin.Context) {
	var query paginationQS
	user, _ := ctx.Get("user")

	ctx.ShouldBindQuery(&query)

	data, err := getFavoriteService(*user.(models.User).Id, &query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

func addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := addFavoriteService(*user.(models.User).Id, id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Produk berhasil ditambahkan ke favorit"))
}

func removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFavoriteService(*user.(models.User).Id, id); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Produk berhasil dihapus dari favorit"))
}
