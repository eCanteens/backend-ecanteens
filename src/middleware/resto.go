package middleware

import (
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/gin-gonic/gin"
)

func Resto(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		response.Error(ctx, "Anda belum login!", 401)
		return
	}

	token = strings.Split(token, " ")[1]

	claim, err := jwt.Parse(token)
	if err != nil {
		response.Error(ctx, "Anda belum login!", 401)
		return
	}

	if claim["type"].(string) != "access" {
		response.Error(ctx, "Token tidak valid", 401)
		return
	}

	if uint(claim["role"].(float64)) != 3 {
		response.Error(ctx, "Pengguna tidak ditemukan", 404)
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", claim["sub"]).Preload("Wallet").Preload("Restaurant").First(&user).Error; err != nil {
		custErr := customerror.GormError(err, "Pengguna")
		response.Error(ctx, custErr.Message, custErr.StatusCode)
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
