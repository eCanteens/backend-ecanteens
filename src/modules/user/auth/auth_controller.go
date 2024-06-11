package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func handleRegister(ctx *gin.Context) {
	var body registerScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	err := registerService(&body)

	if err != nil {
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

func handleGoogle(ctx *gin.Context) {
	var body googleScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, token, err := googleService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{
		"token":     token,
		"data":      data,
		"is_setted": data.Phone != nil,
	}))
}

func handleSetup(ctx *gin.Context) {
	var body setupScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := setupGoogleService(&body, &_user); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Nomor telepon dan Kata sandi berhasil disimpan", helpers.Data{
		"data": _user,
	}))
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

func handleForgot(ctx *gin.Context) {
	var body forgotScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := forgotService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Email telah dikirim"))
}

func handleReset(ctx *gin.Context) {
	var body resetScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := resetService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil direset"))
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
	var body updateScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProfileService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": _user}))
}

func handleUpdatePassword(ctx *gin.Context) {
	var body updatePasswordScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updatePasswordService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil diperbarui"))
}

func handleCheckPin(ctx *gin.Context) {
	var body checkPinScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := checkPinService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Pin benar"))
}

func handleUpdatePin(ctx *gin.Context) {
	var body updatePinScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updatePinService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Pin berhasil diperbarui"))
}
