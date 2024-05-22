package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func handleRegister(ctx *gin.Context) {
	var body RegisterScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, err := registerService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(201, helpers.SuccessResponse("Register berhasil", helpers.Data{"data": &user}))
}

func handleLogin(ctx *gin.Context) {
	var body LoginScheme

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
	var body GoogleScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, token, isPasswordSet, err := googleService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{
		"token":           token,
		"data":            data,
		"is_password_set": isPasswordSet,
	}))
}

func handleForgot(ctx *gin.Context) {
	var body ForgotScheme

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
	var body ResetScheme

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
	var body UpdateScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	__user, err := updateProfileService(ctx, &_user, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": __user}))
}

func handleUpdatePassword(ctx *gin.Context) {
	var body UpdatePasswordScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := updatePasswordService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil diperbarui"))
}

func handleCheckPin(ctx *gin.Context) {
	var body CheckPinScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := checkPinService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Pin benar"))
}

func handleUpdatePin(ctx *gin.Context) {
	var body UpdatePinScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := updatePinService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Pin berhasil diperbarui"))
}
