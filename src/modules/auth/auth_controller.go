package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User

	if response := helpers.Bind(ctx, &user); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := RegisterService(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Register berhasil",
		"data":    &user,
	})
}

func Login(ctx *gin.Context) {
	var body LoginSchema

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	token, err := LoginService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}
