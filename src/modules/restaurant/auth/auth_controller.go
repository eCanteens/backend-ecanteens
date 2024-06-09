package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func handleCheckRegister(ctx *gin.Context) {
	var body checkRegisterScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := checkUniqueService(body.Email, body.Phone); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Data berhasil divalidasi"))
}

func handleRegister(ctx *gin.Context) {
	var body registerScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := registerService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Register berhasil"))
}

func handleLogin(ctx *gin.Context) {
	var body loginScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, token, err := loginService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{"token": token, "data": data}))
}

func handleRefresh(ctx *gin.Context) {
	var body refreshScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	token, err := refreshService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func handleProfile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""
	isPinSet := _user.Wallet.Pin != ""
	_user.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data":       _user,
		"is_pin_set": isPinSet,
	})
}

func handleUpdateProfile(ctx *gin.Context) {
	var body updateProfileScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProfileService(&body, &_user); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui"))
}