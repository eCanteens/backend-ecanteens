package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

// login
func handleAdminLogin(ctx *gin.Context) {
	var body adminLoginScheme

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

// check wallet
func handleCheckWallet(ctx *gin.Context) {
	phone := ctx.Param("phone")

	data, err := checkWalletService(phone)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	data.Password = ""
	data.Wallet.Pin = ""

	ctx.JSON(200, helpers.SuccessResponse("User ditemukan", helpers.Data{"data": data}))
}

// topup
func handleTopup(ctx *gin.Context) {
	phone := ctx.Param("phone")
	var body topupWithdrawScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, err := topupWithdrawService(phone, &body, "TOPUP")

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Top Up Berhasil", helpers.Data{"data": data}))
}

// withdraw
func handleWithdraw(ctx *gin.Context) {
	phone := ctx.Param("phone")
	var body topupWithdrawScheme

	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	data, err := topupWithdrawService(phone, &body, "WITHDRAW")

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Withdraw Berhasil", helpers.Data{"data": data}))
}

// transaction
func handleTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := transactionService(id)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"data": data})
}

// mutasi
func handleMutasi(ctx *gin.Context) {
	var query mutationQS

	ctx.ShouldBindQuery(&query)

	data, err := mutasiService(&query)

	if err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, data)
}

// profile
func handleAdminProfile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func handleUpdateAdminProfile(ctx *gin.Context) {
	var body updateAdminProfileScheme
	if err := helpers.Bind(ctx, &body); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateAdminProfileService(&_user, &body); err != nil {
		ctx.AbortWithStatusJSON(400, helpers.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("Profil berhasil diperbarui", helpers.Data{"data": _user}))
}

func handleUpdateAdminPassword(ctx *gin.Context) {
	var body updateAdminPasswordScheme
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
