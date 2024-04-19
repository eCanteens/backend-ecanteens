package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/gin-gonic/gin"
)

func GetFavoriteResto(ctx *gin.Context) {
	query := map[string]string{}
	user, _ := ctx.Get("user")

	ctx.BindQuery(query)

	data, err := GetFavoriteRestoService(*user.(models.User).Id.Id, query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func GetAllResto(ctx *gin.Context) {
	query := map[string]string{}

	ctx.BindQuery(query)

	data, err := GetAllRestoService(query)
	
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, data)
}

func GetDetailResto(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := GetDetailRestoService(id)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func GetRestosProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	query := map[string]string{}

	ctx.BindQuery(query)

	data, err := GetRestosProductsService(id, query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, data)
}