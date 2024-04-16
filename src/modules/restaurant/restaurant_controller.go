package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/gin-gonic/gin"
)

func GetFavoriteResto(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	data, _ := GetFavoriteRestoService(*user.(models.User).Id.Id)

	ctx.JSON(200, gin.H{
		"data": data,
	})
}