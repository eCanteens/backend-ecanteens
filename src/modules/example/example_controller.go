package example

import (
	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Test(ctx *gin.Context) {
	var users []models.User
	var _pagination models.Pagination
	var page = ctx.Query("page")
	var limit = ctx.Query("limit")

	if err := config.DB.Scopes(pagination.Paginate(&users, &_pagination, &pagination.Params{
		Page:  page,
		Limit: limit,
	})).Where("username ILIKE ?", "%"+ctx.Query("search")+"%").Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, _pagination)
}

func TestPOST(ctx *gin.Context) {
	var user models.User

	// Validate body
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	if err := config.DB.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"data": user,
	})
}
