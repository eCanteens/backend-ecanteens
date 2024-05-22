package middleware

import (
	"errors"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(401, helpers.ErrorResponse("Anda belum login!"))
		return
	}

	token = strings.Split(token, " ")[1]

	claim, err := helpers.ParseJwt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, helpers.ErrorResponse(err.Error()))
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", claim["id"]).Preload("Wallet").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		} else {
			ctx.AbortWithStatusJSON(500, helpers.ErrorResponse(err.Error()))
		}
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
