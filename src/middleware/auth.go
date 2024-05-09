package middleware

import (
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"

	"github.com/eCanteens/backend-ecanteens/src/helpers"

	"github.com/gin-gonic/gin"
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
	config.DB.Where("email = ?", claim["email"]).Preload("Wallet").First(&user)

	ctx.Set("user", user)
	ctx.Next()
}
