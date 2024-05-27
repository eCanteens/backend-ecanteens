package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

// login
func handleAdminLogin(ctx *gin.Context) {
	var body AdminLoginScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, token, err := adminLoginService(&body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Login berhasil", helpers.Data{"token": token, "data": data}))
}

// dashboard
func handleDashoard(ctx *gin.Context) {
	data, err := dashboardService()

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// wallet
func handleCheckWallet(ctx *gin.Context) {
	var body CheckWalletScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := checkWalletService(&body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Wallet Ditemukan"))
}

func handleGetWallet(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := helpers.Bind(ctx, id); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, err := getWalletService(id)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// topup
func handleTopup(ctx *gin.Context) {
	id := ctx.Param("id")
	var body TopupWithdrawScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, err := topupWithdrawService(id, &body, "topup")

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Top Up Berhasil", helpers.Data{"data": data}))
}

// withdraw
func handleWithdraw(ctx *gin.Context) {
	id := ctx.Param("id")
	var body TopupWithdrawScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, err := topupWithdrawService(id, &body, "withdraw")

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Withdraw Berhasil", helpers.Data{"data": data}))
}

// mutasi
func handleMutasi(ctx *gin.Context) {}

// profile
func handleAdminProfile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""
	_user.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func handleUpdateAdminProfile(ctx *gin.Context) {
	var body UpdateAdminProfileScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	__user, err := updateAdminProfileService(ctx, &_user, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": __user}))
}

func handleUpdateAdminPassword(ctx *gin.Context) {
	var body UpdateAdminPasswordScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	if err := updateAdminPasswordService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Password berhasil diperbarui"))
}