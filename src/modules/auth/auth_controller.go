package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	var body RegisterScheme

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	user, err := registerService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Register berhasil", helpers.Data{"data": &user}))
}

func login(ctx *gin.Context) {
	var body LoginScheme

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	token, err := loginService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{"token": token}))
}

func forgot(ctx *gin.Context) {
	var body ForgotScheme

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := forgotService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Email telah dikirim"))
}

func reset(ctx *gin.Context) {
	var body ResetScheme
	token := ctx.Param("token")

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := resetService(&body, token); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil direset"))
}

func profile(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func updateProfile(ctx *gin.Context) {
	var body UpdateScheme
	user, _ := ctx.Get("user")

	if response := helpers.Bind(ctx, &body); response != nil {
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	_user, err := updateProfileService(user.(models.User).Id.Id, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": _user}))
}
