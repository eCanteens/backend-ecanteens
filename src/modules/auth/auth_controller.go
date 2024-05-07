package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var body RegisterSchema

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	user, err := RegisterService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Register berhasil", helpers.Data{"data": &user}))
}

func Login(ctx *gin.Context) {
	var body LoginSchema

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	token, err := LoginService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{"token": token}))
}

func Forgot(ctx *gin.Context) {
	var body ForgotSchema

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := ForgotService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Email telah dikirim"))
}

func Reset(ctx *gin.Context) {
	var body ResetSchema
	token := ctx.Param("token")

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := ResetService(&body, token); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil direset"))
}

func Profile(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func UpdateProfile(ctx *gin.Context) {
	var body UpdateSchema
	user, _ := ctx.Get("user")

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	_user, err := UpdateProfileService(user.(models.User).Id.Id, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": _user}))
}
