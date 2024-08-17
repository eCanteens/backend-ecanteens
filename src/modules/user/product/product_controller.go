package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func addFeedback(ctx *gin.Context) {
	var body feedbackScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := addFeedbackService(&body, *_user.Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	msg := "Produk berhasil di"
	if *body.IsLike {
		msg += "like"
	} else {
		msg += "dislike"
	}

	response.Success(ctx, 200, gin.H{"message": msg})
}

func removeFeedback(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFeedbackService(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Produk berhasil diunlike/undislike"})
}

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

func addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := addFavoriteService(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Produk berhasil ditambahkan ke favorit"})
}

func removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := removeFavoriteService(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Produk berhasil dihapus dari favorit"})
}
