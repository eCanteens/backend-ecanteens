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
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Anda belum login!",
		})
		return
	}

	token = strings.Split(token, " ")[1]

	claim, err := helpers.ParseJwt(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Anda belum login!",
		})
		return
	}

	var user models.User
	config.DB.Where("email = ?", claim["email"]).First(&user)
	user.Password = ""

	ctx.Set("user", user)
	ctx.Next()
}