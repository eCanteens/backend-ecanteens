package example

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	var users []models.User
	var _pagination pagination.Pagination

	if err := _pagination.Paginate(&users, &pagination.Params{
		Query: config.DB.Where("name ILIKE ?", "%"+ctx.Query("search")+"%"),
		Page:  ctx.Query("page"),
		Limit: ctx.Query("limit"),
		Order: ctx.Query("order"),
		Direction: ctx.Query("direction"),
	}).Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, _pagination)
}
