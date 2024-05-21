package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

// login
func handleAdminLogin(ctx *gin.Context) {
	var body LoginScheme

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

// topup
func handleCheckWallet(ctx *gin.Context) {
}

func handleTopup(ctx *gin.Context) {}

// withdraw
func handleWithdraw(ctx *gin.Context) {}

// mutasi
func handleMutasi(ctx *gin.Context) {}

// profile
func handleProfile(ctx *gin.Context) {}

func handleUpdateProfile(ctx *gin.Context) {}

func handleUpdatePassword(ctx *gin.Context) {}